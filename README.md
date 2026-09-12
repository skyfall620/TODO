# TODO API

Simple REST API for managing TODO tasks written in Go.

## Stack

- Go
- PostgreSQL
- pgx
- JWT
- Docker Compose
- Taskfile

## Project structure

```text
.
├── cmd/api          # application entry point
├── config           # configuration
├── internal         # application logic
├── pkg              # shared packages
├── schema           # database schema
├── task             # Taskfile tasks
├── docker-compose.yml
└── Taskfile.yml