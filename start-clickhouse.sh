#!/bin/sh

echo "Starting ClickHouse"
docker run --name clickhouse -p 9000:9000 -p 8123:8123 -d clickhouse/clickhouse-server
