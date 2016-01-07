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
        "user_id": &graphql.ArgumentConfig{
          Type: graphql.NewNonNull(graphql.Int),
        },
      },
      Resolve: createTodo,
    },
    "updateTodo": &graphql.Field{
      Type: todoType,
      Args: graphql.FieldConfigArgument{
        "id": &graphql.ArgumentConfig{
          Type: graphql.NewNonNull(graphql.Int),
        },
        "title": &graphql.ArgumentConfig{
          Type: graphql.String,
        },
        "body": &graphql.ArgumentConfig{
          Type: graphql.String,
        },
        "user_id": &graphql.ArgumentConfig{
          Type: graphql.Int,
        },
      },
      Resolve: updateTodo,
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
    "updateUser": &graphql.Field{
      Type: userType,
      Args: graphql.FieldConfigArgument{
        "id": &graphql.ArgumentConfig{
          Type: graphql.NewNonNull(graphql.Int),
        },
        "name": &graphql.ArgumentConfig{
          Type: graphql.String,
        },
        "email": &graphql.ArgumentConfig{
          Type: graphql.String,
        },
      },
      Resolve: updateUser,
    },
  },
})
