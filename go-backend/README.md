# boilerplate

The goal of boilerplate is to simply prototype/quick start projects while 
writing as little code as possible.

## primary packages
    - gin-gonic/gin
    - gorm.io/gorm
    - spf13/viper
    - deepmap/oapi-codegen

## Features
    - Build gin server interface/structs easily with openapi spec
    - Use gorm for database operations and migrations
    - Use viper for configuration
    - Use postgres as database driver
    - Docker/docker-compose for containerization

## Usage

1. Clone this repository
2. Run `go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v2.1.0`
3. Update the openapi spec `openapi.yaml`
4. Run `make generate` to create api server interface/structs
5. Implement the generated interface/structs in `internal/api/boilerplate.go`
6. Start the server with `make dev-up`

## Notes
- The `Makefile` contains useful commands for development and build
- The `docker-compose.yml` file contains the service definitions for the application
- Should change registry in Makefile/docker-compose.yml to your own 
- oapi-codegen generated output can be customized in server.cfg.yaml and types.cfg.yaml
- Migrations run when the server starts