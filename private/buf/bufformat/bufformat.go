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

package bufformat

import (
	"bytes"
	"context"
	"io"
	"strings"

	"github.com/bufbuild/buf/private/bufpkg/bufmodule"
	"github.com/bufbuild/buf/private/pkg/storage"
	"github.com/jhump/protocompile/ast"
	"github.com/jhump/protocompile/parser"
	"github.com/jhump/protocompile/reporter"
)

// Format formats all of the module files and writes them into
// the given bucket.
func Format(ctx context.Context, moduleFiles []bufmodule.ModuleFile, writeBucket storage.WriteBucket) error {
	reporter := reporter.NewHandler(nil)
	fileNodes := make([]*ast.FileNode, 0, len(moduleFiles))
	for _, moduleFile := range moduleFiles {
		fileNode, err := parser.Parse(moduleFile.ExternalPath(), moduleFile, reporter)
		if err != nil {
			return err
		}
		fileNodes = append(fileNodes, fileNode)
	}
	formatted := make(map[string]io.Reader, len(fileNodes))
	for i, fileNode := range fileNodes {
		reader, err := format(fileNode)
		if err != nil {
			return err
		}
		formatted[moduleFiles[i].ExternalPath()] = reader
	}
	for externalPath, reader := range formatted {
		writeObjectCloser, err := writeBucket.Put(ctx, externalPath)
		if err != nil {
			return err
		}
		if _, err = io.Copy(writeObjectCloser, reader); err != nil {
			return err
		}
	}
	return nil
}

// format rewrites the given FileNode into formatted file content, represented as a string.
func format(fileNode *ast.FileNode) (io.Reader, error) {
	buffer := bytes.NewBuffer(nil)
	if err := newFormatter(buffer, fileNode).Run(); err != nil {
		return nil, err
	}
	// This is a hack to remove trailing newlines before the file is
	// written. Clean this up.
	trimmed := bytes.NewBuffer(nil)
	if _, err := trimmed.WriteString(strings.TrimRight(buffer.String(), "\n")); err != nil {
		return nil, err
	}
	return trimmed, nil
}
