# ClickHouse GitHub Action [![Tests](https://github.com/EpicStep/clickhouse-github-action/actions/workflows/ci.yml/badge.svg)](https://github.com/EpicStep/clickhouse-github-action/actions/workflows/ci.yml)

This action sets up a ClickHouse.

# Usage

```yaml
steps:
  - uses: EpicStep/clickhouse-github-action@v1.1.1
    with:
      version: 22 # default: latest
      bind_native: 9090 # default: 9000 (https://clickhouse.com/docs/en/interfaces/tcp)
      bind_http: 8080 # default: 8123 (https://clickhouse.com/docs/en/interfaces/http)
```

# License

This project released under the [MIT License](LICENSE)
