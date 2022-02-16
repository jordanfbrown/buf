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

package bufencoding

import (
	"context"
	"fmt"
	"path/filepath"
	"strings"

	"go.opencensus.io/trace"
)

const (
	// MessageEncodingBin is the binary image encoding.
	MessageEncodingBin MessageEncoding = iota + 1
	// MessageEncodingJSON is the JSON image encoding.
	MessageEncodingJSON
)

// MessageEncoding is the encoding of the message
type MessageEncoding int

type MessageEncodingRef interface {
	Path() string
	MessageEncoding() MessageEncoding
}

func NewMessageEncodingRef(
	ctx context.Context,
	value string,
) (MessageEncodingRef, error) {
	ctx, span := trace.StartSpan(ctx, "new_message_encoding_ref")
	defer span.End()
	path, messageEncoding, err := getPathAndMessageEncoding(ctx, value)
	if err != nil {
		return nil, err
	}
	return newMessageEncodingRef(path, messageEncoding), nil
}

func getPathAndMessageEncoding(
	ctx context.Context,
	value string,
) (string, MessageEncoding, error) {
	path, options, err := getRawPathAndOptions(value)
	if err != nil {
		return "", 0, err
	}
	messageEncoding := parseMessageEncodingExt(filepath.Ext(path))
	if format, ok := options["format"]; ok {
		messageEncoding, err = parseMessageEncodingFormat(format)
		if err != nil {
			return "", 0, err
		}
	}
	return path, messageEncoding, nil
}

func parseMessageEncodingExt(ext string) MessageEncoding {
	switch ext {
	case ".bin":
		return MessageEncodingBin
	case ".json":
		return MessageEncodingJSON
	default:
		// this allows the filename without extension use the default encoding type based on the message encoding type of input.
		return 0
	}
}

func parseMessageEncodingFormat(format string) (MessageEncoding, error) {
	switch format {
	case "bin":
		return MessageEncodingBin, nil
	case "json":
		return MessageEncodingJSON, nil
	default:
		return 0, fmt.Errorf("invalid format for message: %q", format)
	}
}

// rawPath will be non-empty
func getRawPathAndOptions(value string) (string, map[string]string, error) {
	value = strings.TrimSpace(value)
	if value == "" {
		return "", nil, fmt.Errorf("value is required")
	}
	switch splitValue := strings.Split(value, "#"); len(splitValue) {
	case 1:
		return value, nil, nil
	case 2:
		path := strings.TrimSpace(splitValue[0])
		optionsString := strings.TrimSpace(splitValue[1])
		if path == "" {
			return "", nil, fmt.Errorf("%q starts with # which is invalid", value)
		}
		if optionsString == "" {
			return "", nil, fmt.Errorf("%q ends with # which is invalid", value)
		}
		options := make(map[string]string)
		for _, pair := range strings.Split(optionsString, ",") {
			split := strings.Split(pair, "=")
			if len(split) != 2 {
				return "", nil, fmt.Errorf("invalid options: %q", optionsString)
			}
			key := strings.TrimSpace(split[0])
			value := strings.TrimSpace(split[1])
			if key == "" || value == "" {
				return "", nil, fmt.Errorf("invalid options: %q", optionsString)
			}
			if _, ok := options[key]; ok {
				return "", nil, fmt.Errorf("duplicate options key: %q", key)
			}
			options[key] = value
		}
		return path, options, nil
	default:
		return "", nil, fmt.Errorf("%q has multiple #s which is invalid", value)
	}
}
