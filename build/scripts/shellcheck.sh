#!/bin/sh

set -euxo

cd "${SRC_DIR}"

shellcheck --version
# shellcheck disable=SC2046
shellcheck $(find "${SHELLCHECK_SOURCEPATH}" -type f -name '*.sh')
