# golang-rest-api-template
Golang RestAPI Template

A Go-based backend rest api service template.

## Project Structure

```
.
├── cmd/                             # Entry points (main.go ของแต่ละ service/cli)
├── config/                          # Config loader (YAML/ENV)
├── pkg/                             # Public libraries (external consumers)
├── internal/
│   ├── core/
│   │   ├── domain/                  # Entities & Value Objects
│   │   ├── port/
│   │   │   ├── inbound/             # Inbound Ports Interfaces
│   │   │   └── outbound/            # Outbound Ports Interfaces
│   │   └── service/                 # Use-Case Services (Primary Adapters)
│   ├── handler/                     # Inbound Adapters (HTTP, gRPC, CLI)
│   │   ├── http/
│   │   └── kafka/
│   ├── repository/                  # Outbound Adapters (DB, Cache)
│   │   ├── postgres/
│   │   ├── redis/
│   │   └── mongo/
│   └── util/                        # Utility packages (logging, middleware)
├── protocol/                        # API definitions (GraphQL, gRPC, OpenAPI, JSON Schema)
│   ├── graphql/                     # GraphQL schema & codegen config
│   │   ├── schema.graphql
│   │   └── gqlgen.yml
│   ├── proto/                       # gRPC / Protobuf IDL
│   │   ├── user.proto
│   │   └── course.proto
│   ├── openapi/                     # OpenAPI/Swagger specs
│   │   ├── openapi.yaml
│   │   └── swagger-config.json
│   └── jsonschema/                  # JSON Schema definitions (ถ้ามี)
│       ├── user.json
│       └── course.json
├── Dockerfile
├── docker-compose.yml
├── go.mod
├── go.sum
└── init.sql
```

## Prerequisites

- Go 1.24 or later
- Docker and Docker Compose
- PostgreSQL

## Getting Started

### Set up environment variables:

   ```bash
   cp .env.example .env
   # Edit .env with your configuration
   ```

### . Run the application:
   ```bash
   go run cmd/server/main.go
   ```

## Development

### Database Migrations

The project uses Ent for database schema management. To create a new migration:

```bash
go generate ./ent
```

## Dependencies

Main dependencies:

- [Echo](https://github.com/labstack/echo) - Web framework
- [Ent](https://entgo.io/) - Entity framework
- [Viper](https://github.com/spf13/viper) - Configuration management
- [Zap](https://go.uber.org/zap) - Logging
- [PostgreSQL](https://www.postgresql.org/) - Database

## Hot Reload 
- (แนะนำ) [air](https://github.com/cosmtrek/air) สำหรับ hot reload

