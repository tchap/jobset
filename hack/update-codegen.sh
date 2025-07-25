#!/usr/bin/env bash

# Copyright 2023 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -o errexit
set -o nounset
set -o pipefail

cd "$(dirname "${0}")/.."
GO_CMD=${1:-go}
CODEGEN_PKG=${2:-bin}
REPO_ROOT="$(git rev-parse --show-toplevel)"

echo "GOPATH=$(go env GOPATH)"

source "${CODEGEN_PKG}/kube_codegen.sh"

kube::codegen::gen_helpers \
    --boilerplate "${REPO_ROOT}/hack/boilerplate.go.txt" \
    "${REPO_ROOT}/api"

kube::codegen::gen_client \
    --with-watch \
    --with-applyconfig \
    --applyconfig-externals "k8s.io/api/batch/v1.JobTemplateSpec:k8s.io/client-go/applyconfigurations/batch/v1" \
    --output-dir "${REPO_ROOT}/client-go" \
    --output-pkg sigs.k8s.io/jobset/client-go \
    --boilerplate "${REPO_ROOT}/hack/boilerplate.go.txt" \
    "${REPO_ROOT}/api"
