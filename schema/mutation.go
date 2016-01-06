package schema

import (
  "github.com/graphql-go/graphql"
)

var rootMutation = graphql.NewObject(graphql.ObjectConfig{
  Name: "RootMutation",
  Fields: graphql.Fields{
    "createTodo": &graphql.Field{
      Type: todoType,
      Args: graphql.FieldConfigArgument{
        "title": &graphql.ArgumentConfig{
          Type: graphql.NewNonNull(graphql.String),
        },
        "body": &graphql.ArgumentConfig{
          Type: graphql.NewNonNull(graphql.String),
        },
      },
      Resolve: createTodo,
    },
    "createUser": &graphql.Field{
      Type: userType,
      Args: graphql.FieldConfigArgument{
        "name": &graphql.ArgumentConfig{
          Type: graphql.NewNonNull(graphql.String),
        },
        "email": &graphql.ArgumentConfig{
          Type: graphql.NewNonNull(graphql.String),
        },
      },
      Resolve: createUser,
    },
  },
})
