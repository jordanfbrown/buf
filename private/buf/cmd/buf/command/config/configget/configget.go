// Copyright 2020-2022 Buf Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package configget

import (
	"context"
	"fmt"
	"io"

	"github.com/bufbuild/buf/private/buf/bufcli"
	"github.com/bufbuild/buf/private/buf/bufprint"
	"github.com/bufbuild/buf/private/bufpkg/bufconfig"
	"github.com/bufbuild/buf/private/pkg/app/appcmd"
	"github.com/bufbuild/buf/private/pkg/app/appflag"
	"github.com/bufbuild/buf/private/pkg/encoding"
	"github.com/bufbuild/buf/private/pkg/storage/storageos"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

const formatFlagName = "format"

// NewCommand returns a new Command.
func NewCommand(
	name string,
	builder appflag.Builder,
) *appcmd.Command {
	flags := newFlags()
	return &appcmd.Command{
		Use:   name + " <source>",
		Short: "Outputs configuration.",
		Long:  bufcli.GetSourceLong(`the source`),
		Args:  cobra.MaximumNArgs(1),
		Run: builder.NewRunFunc(
			func(ctx context.Context, container appflag.Container) error {
				return run(ctx, container, flags)
			},
			bufcli.NewErrorInterceptor(),
		),
		BindFlags: flags.Bind,
	}
}

type flags struct {
	Format string
	// special
	InputHashtag string
}

func newFlags() *flags {
	return &flags{}
}

func (f *flags) Bind(flagSet *pflag.FlagSet) {
	bufcli.BindInputHashtag(flagSet, &f.InputHashtag)
	flagSet.StringVar(
		&f.Format,
		formatFlagName,
		bufprint.FormatText.String(),
		fmt.Sprintf(`The output format to use. Must be one of %s`, bufprint.AllFormatsString),
	)
}

func run(
	ctx context.Context,
	container appflag.Container,
	flags *flags,
) error {
	format, err := bufprint.ParseFormat(flags.Format)
	if err != nil {
		return appcmd.NewInvalidArgumentError(err.Error())
	}
	source, err := bufcli.GetInputValue(container, flags.InputHashtag, ".")
	if err != nil {
		return err
	}
	storageosProvider := storageos.NewProvider(storageos.ProviderWithSymlinks())
	readWriteBucket, err := storageosProvider.NewReadWriteBucket(
		source,
		storageos.ReadWriteBucketWithSymlinksIfSupported(),
	)
	if err != nil {
		return err
	}
	existingConfigFilePath, err := bufconfig.ExistingConfigFilePath(ctx, readWriteBucket)
	if err != nil {
		return err
	}
	if existingConfigFilePath == "" {
		return bufcli.ErrNoConfigFile
	}
	printer := bufprint.NewConfigPrinter(container.Stdout())
	configFile, ok, err := bufconfig.GetConfigReaderObjectCloserForBucket(ctx, readWriteBucket)
	if err != nil {
		return err
	}
	if !ok {
		// TODO: change to V1 when we make V1 the default
		return printer.PrintExternalConfigV1Beta1(format, bufconfig.ExternalConfigV1Beta1{}, []byte{})
	}
	data, err := io.ReadAll(configFile)
	if err != nil {
		return err
	}
	if err := configFile.Close(); err != nil {
		return err
	}

	var externalConfigVersion bufconfig.ExternalConfigVersion
	if err := encoding.UnmarshalJSONOrYAMLNonStrict(data, &externalConfigVersion); err != nil {
		return err
	}
	switch externalConfigVersion.Version {
	case bufconfig.V1Beta1Version:
		config := bufconfig.ExternalConfigV1Beta1{}
		if err := encoding.UnmarshalJSONOrYAMLStrict(data, &config); err != nil {
			return err
		}
		return printer.PrintExternalConfigV1Beta1(format, config, data)
	case bufconfig.V1Version:
		config := bufconfig.ExternalConfigV1{}
		if err := encoding.UnmarshalJSONOrYAMLStrict(data, &config); err != nil {
			return err
		}
		return printer.PrintExternalConfigV1(format, config, data)
	default:
		return fmt.Errorf(
			`%s has an invalid "version: %s" set. Please add "version: %s". See https://docs.buf.build/faq for more details`,
			configFile.ExternalPath(),
			externalConfigVersion.Version,
			bufconfig.V1Version,
		)
	}
}
