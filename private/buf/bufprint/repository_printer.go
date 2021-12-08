// Copyright 2020-2021 Buf Technologies, Inc.
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
	"context"
	"encoding/json"
	"fmt"
	"io"
	"time"

	"github.com/bufbuild/buf/private/gen/proto/apiclient/buf/alpha/registry/v1alpha1/registryv1alpha1apiclient"
	registryv1alpha1 "github.com/bufbuild/buf/private/gen/proto/go/buf/alpha/registry/v1alpha1"
)

type repositoryPrinter struct {
	address     string
	writer      io.Writer
	remote      string
	apiProvider registryv1alpha1apiclient.OwnerServiceProvider
}

func newRepositoryPrinter(
	address string,
	writer io.Writer,
	remote string,
	apiProvider registryv1alpha1apiclient.OwnerServiceProvider,
) *repositoryPrinter {
	return &repositoryPrinter{
		address:     address,
		writer:      writer,
		remote:      remote,
		apiProvider: apiProvider,
	}
}

func (p *repositoryPrinter) getRepositoryOwnerNames(
	ctx context.Context,
	repositories ...*registryv1alpha1.Repository,
) (map[string]string, error) {
	service, err := p.apiProvider.NewOwnerService(ctx, p.remote)
	if err != nil {
		return nil, err
	}
	ownerIDs := make([]string, len(repositories))
	for i, repository := range repositories {
		switch owner := repository.Owner.(type) {
		case *registryv1alpha1.Repository_OrganizationId:
			ownerIDs[i] = owner.OrganizationId
		case *registryv1alpha1.Repository_UserId:
			ownerIDs[i] = owner.UserId
		}
	}
	owners, err := service.GetOwnersByID(ctx, ownerIDs)
	if err != nil {
		return nil, err
	}
	ownerNames := make(map[string]string, len(owners))
	for _, o := range owners {
		switch owner := o.Owner.(type) {
		case *registryv1alpha1.Owner_User:
			ownerNames[owner.User.Id] = owner.User.Username
		case *registryv1alpha1.Owner_Organization:
			ownerNames[owner.Organization.Id] = owner.Organization.Name
		}
	}
	return ownerNames, nil
}

func (p *repositoryPrinter) PrintRepository(ctx context.Context, format Format, message *registryv1alpha1.Repository) error {
	outputRepositories, err := p.registryRepositoriesToOutRepositories(ctx, message)
	if err != nil {
		return err
	}
	if len(outputRepositories) != 1 {
		return fmt.Errorf("error converting repositories: expected 1 got %d", len(outputRepositories))
	}
	switch format {
	case FormatText:
		return p.printRepositoriesText(outputRepositories)
	case FormatJSON:
		return json.NewEncoder(p.writer).Encode(outputRepositories[0])
	default:
		return fmt.Errorf("unknown format: %v", format)
	}
}

func (p *repositoryPrinter) PrintRepositories(ctx context.Context, format Format, nextPageToken string, messages ...*registryv1alpha1.Repository) error {
	if len(messages) == 0 {
		return nil
	}
	outputRepositories, err := p.registryRepositoriesToOutRepositories(ctx, messages...)
	if err != nil {
		return err
	}
	switch format {
	case FormatText:
		return p.printRepositoriesText(outputRepositories)
	case FormatJSON:
		return json.NewEncoder(p.writer).Encode(paginationWrapper{
			NextPage: nextPageToken,
			Results:  outputRepositories,
		})
	default:
		return fmt.Errorf("unknown format: %v", format)
	}
}

func (p *repositoryPrinter) registryRepositoriesToOutRepositories(
	ctx context.Context,
	messages ...*registryv1alpha1.Repository,
) ([]outputRepository, error) {
	ownerNames, err := p.getRepositoryOwnerNames(ctx, messages...)
	if err != nil {
		return nil, err
	}
	var outputRepositories []outputRepository
	for _, repository := range messages {
		var ownerID string
		switch owner := repository.Owner.(type) {
		case *registryv1alpha1.Repository_OrganizationId:
			ownerID = owner.OrganizationId
		case *registryv1alpha1.Repository_UserId:
			ownerID = owner.UserId
		}
		owner, ok := ownerNames[ownerID]
		if !ok {
			return nil, fmt.Errorf("missing owner name for id: %s", ownerID)
		}
		outputRepositories = append(outputRepositories, outputRepository{
			ID:         repository.Id,
			Remote:     p.address,
			Owner:      owner,
			Name:       repository.Name,
			CreateTime: repository.CreateTime.AsTime(),
		})
	}
	return outputRepositories, nil
}

func (p *repositoryPrinter) printRepositoriesText(outputRepositories []outputRepository) error {
	return WithTabWriter(
		p.writer,
		[]string{
			"Full name",
			"Created",
		},
		func(tabWriter TabWriter) error {
			for _, outputRepository := range outputRepositories {
				if err := tabWriter.Write(
					outputRepository.Remote+"/"+outputRepository.Owner+"/"+outputRepository.Name,
					outputRepository.CreateTime.Format(time.RFC3339),
				); err != nil {
					return err
				}
			}
			return nil
		},
	)
}

type outputRepository struct {
	ID         string    `json:"id,omitempty"`
	Remote     string    `json:"remote,omitempty"`
	Owner      string    `json:"owner,omitempty"`
	Name       string    `json:"name,omitempty"`
	CreateTime time.Time `json:"create_time,omitempty"`
}
