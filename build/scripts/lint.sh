#!/bin/bash

set -euxo pipefail

cd "${SRC_DIR}"

go version
golangci-lint --version

golangci-lint run -c "${GO_LINT_CONFIG}"
