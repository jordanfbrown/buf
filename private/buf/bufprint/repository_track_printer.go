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
	"encoding/json"
	"fmt"
	"io"
	"time"

	registryv1alpha1 "github.com/bufbuild/buf/private/gen/proto/go/buf/alpha/registry/v1alpha1"
)

type repositoryTrackPrinter struct {
	writer io.Writer
}

func newRepositoryTrackPrinter(
	writer io.Writer,
) *repositoryTrackPrinter {
	return &repositoryTrackPrinter{
		writer: writer,
	}
}

func (p *repositoryTrackPrinter) PrintRepositoryTrack(format Format, message *registryv1alpha1.RepositoryTrack) error {
	outRepositoryTrack := registryTrackToOutputTrack(message)
	switch format {
	case FormatText:
		return p.printRepositoryTracksText([]outputRepositoryTrack{outRepositoryTrack})
	case FormatJSON:
		return json.NewEncoder(p.writer).Encode(outRepositoryTrack)
	default:
		return fmt.Errorf("unknown format: %v", format)
	}
}

func (p *repositoryTrackPrinter) PrintRepositoryTracks(format Format, nextPageToken string, messages ...*registryv1alpha1.RepositoryTrack) error {
	if len(messages) == 0 {
		return nil
	}
	outputRepositoryTracks := registryTracksToOutputTracks(messages)
	switch format {
	case FormatText:
		return p.printRepositoryTracksText(outputRepositoryTracks)
	case FormatJSON:
		return json.NewEncoder(p.writer).Encode(paginationWrapper{
			NextPage: nextPageToken,
			Results:  outputRepositoryTracks,
		})
	default:
		return fmt.Errorf("unknown format: %v", format)
	}
}

func (p *repositoryTrackPrinter) printRepositoryTracksText(outputRepositoryTracks []outputRepositoryTrack) error {
	return WithTabWriter(
		p.writer,
		[]string{
			"Name",
			"Commit",
			"Created",
		},
		func(tabWriter TabWriter) error {
			for _, ot := range outputRepositoryTracks {
				if err := tabWriter.Write(
					ot.Name,
					ot.CreateTime.Format(time.RFC3339),
				); err != nil {
					return err
				}
			}
			return nil
		},
	)
}

type outputRepositoryTrack struct {
	ID         string    `json:"id,omitempty"`
	Name       string    `json:"name,omitempty"`
	CreateTime time.Time `json:"create_time,omitempty"`
}

func registryTrackToOutputTrack(
	repositoryTrack *registryv1alpha1.RepositoryTrack,
) outputRepositoryTrack {
	return outputRepositoryTrack{
		Name:       repositoryTrack.Name,
		CreateTime: repositoryTrack.CreateTime.AsTime(),
	}
}

func registryTracksToOutputTracks(tracks []*registryv1alpha1.RepositoryTrack) []outputRepositoryTrack {
	outputRepositoryTracks := make([]outputRepositoryTrack, len(tracks))
	for i, track := range tracks {
		outputRepositoryTracks[i] = registryTrackToOutputTrack(track)
	}
	return outputRepositoryTracks
}
