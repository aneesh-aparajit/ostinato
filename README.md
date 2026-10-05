# ostinato

A Postgres-backed job scheduler written in Go. Supports one-off tasks and cron schedules (with timezones, timeouts and misfire policies).

## Requirements

- Go 1.25+
- PostgreSQL 17
- [sqlc](https://sqlc.dev) and [golang-migrate](https://github.com/golang-migrate/migrate) for development

## Getting started

```bash
docker compose up -d postgres
migrate -path pkg/adapters/postgres/migrations -database "postgres://postgres:postgres@localhost:5432/ostinato?sslmode=disable" up
make run
```

Configuration lives in `resources/properties.yaml`.

## Development

| Command             | Description                                   |
| ------------------- | --------------------------------------------- |
| `make build`        | Build the binary to `bin/ostinato`            |
| `make run`          | Build and run                                 |
| `make sql-generate` | Regenerate Go code from `queries/` with sqlc  |

## Layout

```
main.go
pkg/
  config/              config loading (viper)
  logger/              logging (zap)
  adapters/postgres/   migrations, sqlc queries and generated code
resources/             runtime configuration
```
