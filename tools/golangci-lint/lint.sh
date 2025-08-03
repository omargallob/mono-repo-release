#!/usr/bin/env bash

set -euo pipefail

# Run  golangci-lint
exec bazel run //tools/golangci-lint:golangci-lint -- "$@"