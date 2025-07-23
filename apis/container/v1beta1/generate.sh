#!/bin/bash
#
# Copyright 2020 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -o errexit
set -o nounset
set -o pipefail

# The following solution for making go-to-protobuf work is taken from
# https://github.com/kubernetes/code-generator/issues/91#issuecomment-838213462
# See also https://github.com/kubernetes/code-generator/issues/91.
#
# The go-to-protobuf binary does not seem to be able to find the .proto files
# unless they are on the GOPATH.
# We create a temporary GOPATH and copy the necessary files there.
# The protobuf definitions are in
# k8s.io/api and k8s.io/apimachinery.
# The tool also needs the `go-to-protobuf` package at the root, so we also copy
# the `hack` directory.
# The `deepcopy-gen` and `defaulter-gen` tools also need to be on the GOPATH.
#
# The following variables are expected to be set:
# - SCRIPT_ROOT: the root of the repository
# - CODEGEN_PKG: the code-generator package
# - GOPATH: the user's GOPATH

# Create a temporary GOPATH
GOPATH_TMP="$(mktemp -d)"
# Copy the necessary files to the temporary GOPATH
mkdir -p "${GOPATH_TMP}/src/k8s.io"
cp -r "${GOPATH}/src/k8s.io/api" "${GOPATH_TMP}/src/k8s.io"
cp -r "${GOPATH}/src/k8s.io/apimachinery" "${GOPATH_TMP}/src/k8s.io"
cp -r "${CODEGEN_PKG}/hack" "${GOPATH_TMP}/src"
cp -r "${CODEGEN_PKG}/generate-groups.sh" "${GOPATH_TMP}/src"
cp -r "${CODEGEN_PKG}/cmd/deepcopy-gen" "${GOPATH_TMP}/src"
cp -r "${CODEGEN_PKG}/cmd/defaulter-gen" "${GOPATH_TMP}/src"
# Also copy the current repository to the temporary GOPATH
mkdir -p "${GOPATH_TMP}/src/github.com/GoogleCloudPlatform"
cp -r "${SCRIPT_ROOT}" "${GOPATH_TMP}/src/github.com/GoogleCloudPlatform/k8s-config-connector"

# Run the code generator
(
  # To make go-to-protobuf work, we need to be in the GOPATH
  cd "${GOPATH_TMP}/src"
  # We need to set the GOPATH to the temporary GOPATH
  export GOPATH="${GOPATH_TMP}"
  # We need to be in a directory that is a git repository
  # so that the code generator can find the boilerplate file.
  # We create a fake git repository.
  git init >/dev/null
  git config user.email "test@example.com"
  git config user.name "Test"
  git commit --allow-empty -m "Initial commit" >/dev/null
  # Run the code generator
  bash "${GOPATH_TMP}/src/generate-groups.sh" "all" \
    "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated" \
    "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis" \
    "container:v1beta1" \
    --go-header-file "${SCRIPT_ROOT}/hack/boilerplate.go.txt"
)

# Clean up the temporary GOPATH
rm -rf "${GOPATH_TMP}"
