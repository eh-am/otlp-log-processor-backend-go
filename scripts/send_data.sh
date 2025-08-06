#!/usr/bin/env bash

set -euo pipefail


grpcurl -plaintext \
    -proto opentelemetry/proto/collector/logs/v1/logs_service.proto \
    localhost:4317 \
    opentelemetry.proto.collector.logs.v1.LogsService/Export
