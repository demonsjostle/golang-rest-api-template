# golang-rest-api-template
Golang RestAPI Template

A Go-based backend rest api service template.

## Project Structure

```
.
├── cmd/            # Application entry points
├── config/         # Configuration files and environment variables
├── ent/           # Entity definitions and database schema
├── internal/      # Private application code
├── infrastructure/# Infrastructure layer (database, external services)
├── pkg/           # Public libraries that can be used by external applications
├── protocol/      # API protocol definitions and interfaces
├── Dockerfile     # Container definition
├── docker-compose.yml # Local development environment setup
├── go.mod         # Go module definition
├── go.sum         # Go module checksums
└── init.sql       # Database initialization script
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

