#!/bin/bash

INPUT_VERSION="${INPUT_VERSION:-latest}"
INPUT_BIND_NATIVE="${INPUT_BIND_NATIVE:-9000}"
INPUT_BIND_HTTP="${INPUT_BIND_HTTP:-8123}"

echo "Starting ClickHouse"
docker run --name clickhouse -p "$INPUT_BIND_NATIVE":9000 -p "$INPUT_BIND_HTTP":8123 -d clickhouse/clickhouse-server:"$INPUT_VERSION"
