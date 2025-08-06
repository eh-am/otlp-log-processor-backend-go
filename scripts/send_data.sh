#!/usr/bin/env bash

set -euo pipefail

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )

grpcurl -plaintext \
    -d @ \
    -proto opentelemetry/proto/collector/logs/v1/logs_service.proto \
    localhost:4317 \
    opentelemetry.proto.collector.logs.v1.LogsService/Export < "$SCRIPT_DIR/logs.json"
