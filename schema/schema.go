package schema

import (
  "fmt"

  "github.com/graphql-go/graphql"
)

var schema, _ = graphql.NewSchema(graphql.SchemaConfig{
  Query: rootQuery,
  Mutation: rootMutation,
})

func ExecuteQuery(query string) *graphql.Result {
  result := graphql.Do(graphql.Params{
    Schema: schema,
    RequestString: query,
  })
  if len(result.Errors) > 0 {
    fmt.Printf("Unexpected errors: %v\n", result.Errors)
  }
  return result
}
