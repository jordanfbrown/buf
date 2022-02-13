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

package bufprint

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/bufbuild/buf/private/bufpkg/bufconfig"
)

type configPrinter struct {
	writer io.Writer
}

func newConfigPrinter(writer io.Writer) *configPrinter {
	return &configPrinter{
		writer: writer,
	}
}

func (p *configPrinter) PrintExternalConfigV1Beta1(
	format Format,
	config bufconfig.ExternalConfigV1Beta1,
	raw []byte,
) error {
	switch format {
	case FormatText:
		return p.printText(raw)
	case FormatJSON:
		return p.printJSON(config)
	default:
		return fmt.Errorf("unknown format: %v", format)
	}
}

func (p *configPrinter) PrintExternalConfigV1(
	format Format,
	config bufconfig.ExternalConfigV1,
	raw []byte,
) error {
	switch format {
	case FormatText:
		return p.printText(raw)
	case FormatJSON:
		return p.printJSON(config)
	default:
		return fmt.Errorf("unknown format: %v", format)
	}
}

func (p *configPrinter) printText(raw []byte) error {
	_, err := p.writer.Write(raw)
	return err
}

func (p *configPrinter) printJSON(val interface{}) error {
	encoder := json.NewEncoder(p.writer)
	encoder.SetIndent("", "  ")
	return encoder.Encode(val)
}
