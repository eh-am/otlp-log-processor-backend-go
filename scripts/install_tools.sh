#!/usr/bin/env bash

set -euo pipefail

echo "Installing grpccurl"

go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
grpccurl
