package main

import (
  "fmt"
  "flag"
  "net/http"

  "github.com/dhanui/graphql_api_go/repository"
  "github.com/dhanui/graphql_api_go/handlers"
)

var dbConfPath = flag.String("C", "./config/database.json", "path to database config file")
var migrationPath = flag.String("m", "./migrations", "path to migration directory")
var port = flag.Int("p", 8080, "HTTP port")

func printErrors(errors []error) {
  for i := 0; i < len(errors); i++ {
    fmt.Printf("* %s\n", errors[i].Error())
  }
}

func main() {
  flag.Parse()
  args := flag.Args()
  if len(args) == 1 {
    switch args[0] {
    case "migrate":
      errors := repository.Migrate(*dbConfPath, *migrationPath, false)
      if len(errors) > 0 {
        fmt.Println("Migration errors:")
        printErrors(errors)
      }
    case "rollback":
      errors := repository.Rollback(*dbConfPath, *migrationPath, false)
      if len(errors) > 0 {
        fmt.Println("Rollback errors:")
        printErrors(errors)
      }
    case "server":
      err := repository.InitDbConnection(*dbConfPath, false)
      if (err != nil) {
        fmt.Printf("Error initializing database connection: %s\n", err.Error())
        return
      }
      http.HandleFunc("/graphql", handlers.GraphqlHandler)
      fmt.Printf("HTTP server listening on port %d...\n", *port)
      err = http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
      if err != nil {
        fmt.Printf("Error starting HTTP server: %s\n", err.Error())
      }
    default:
      fmt.Printf("Unknown command: %s\n", args[0])
    }
  } else {
    fmt.Println("Usage: graphql_api_go [flags] [command]")
    fmt.Println()
    fmt.Println("Flags:")
    fmt.Println("  -C string    path to database config file (default: ./config/database.json)")
    fmt.Println("  -m string    path to migration directory (default: ./migrations)")
    fmt.Println("  -p int       HTTP port (default: 8080)")
    fmt.Println()
    fmt.Println("Commands:")
    fmt.Println("  migrate      migrate database")
    fmt.Println("  rollback     rollback database")
    fmt.Println("  server       start HTTP server")
  }
}
