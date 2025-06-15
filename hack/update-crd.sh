#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

${GOPATH}/bin/controller-gen crd paths=./pkg/apis/... output:crd:dir=./pkg/crd