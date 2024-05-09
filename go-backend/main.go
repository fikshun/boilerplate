package main

import (
  "flag"
  "fmt"
  "github.com/fikshun/go-backend/internal/api"
  "github.com/fikshun/go-backend/internal/database"
  "github.com/fikshun/go-backend/internal/sec"
  "github.com/gin-gonic/gin"
  middleware "github.com/oapi-codegen/gin-middleware"
  "gorm.io/gorm"
  "log"
  "net"
  "net/http"
  "os"
)

var DB *gorm.DB

func NewBoilerplateServer(b *api.Boilerplate, port string) *http.Server {
  swagger, err := api.GetSwagger()
  if err != nil {
    _, err := fmt.Fprintf(os.Stderr, "Error loading swagger spec: %s\n", err)
    if err != nil {
      return nil
    }
    os.Exit(1)
  }

  // skip server validation
  swagger.Servers = nil

  // basic gin router
  r := gin.Default()

  // Use validation middleware to check all requests against
  // OpenAPI schema
  r.Use(middleware.OapiRequestValidator(swagger))
  // Use auth middleware
  r.Use(sec.TokenAuthMiddleWare())

  // register custom boilerplate handlers
  api.RegisterHandlers(r, b)

  s := &http.Server{
    Handler: r,
    Addr:    net.JoinHostPort("0.0.0.0", port),
  }

  return s
}

func main() {

  // Initialize database
  DB = database.InitDB()
  // Migrate database
  database.Migrate(DB)

  port := flag.String("port", "8080", "port to listen on")
  flag.Parse()
  // Create instance of boilerplate handlers that satisfies internal.api.ServerInterface
  b := api.NewBoilerplate(DB)
  // Create server
  s := NewBoilerplateServer(b, *port)
  // Start server
  log.Fatal(s.ListenAndServe())
}
