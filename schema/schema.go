package schema

import (
  "fmt"

  "github.com/graphql-go/graphql"
  "golang.org/x/net/context"

  "github.com/dhanui/graphql_api_go/repository"
)

var schema, _ = graphql.NewSchema(graphql.SchemaConfig{
  Query: queryType,
  Mutation: mutationType,
})

func ExecuteQuery(query string, user repository.User) *graphql.Result {
  result := graphql.Do(graphql.Params{
    Schema: schema,
    RequestString: query,
    Context: context.WithValue(context.Background(), "currentUser", user),
  })
  if len(result.Errors) > 0 {
    fmt.Printf("Request payload:\n%s\n", query)
    fmt.Printf("Errors:\n")
    for i := 0; i < len(result.Errors); i++ {
      fmt.Printf("* %s\n", result.Errors[i].Error())
    }
  }
  return result
}
