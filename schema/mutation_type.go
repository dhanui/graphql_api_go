package schema

import (
  "github.com/graphql-go/graphql"
)

var mutationType = graphql.NewObject(graphql.ObjectConfig{
  Name: "Mutation",
  Fields: graphql.Fields{
    "createTodo": &graphql.Field{
      Description: "Create new to do",
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
      Description: "Update existing to do",
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
    "deleteTodo": &graphql.Field{
      Description: "Delete existing to do",
      Type: todoType,
      Args: graphql.FieldConfigArgument{
        "id": &graphql.ArgumentConfig{
          Type: graphql.NewNonNull(graphql.Int),
        },
      },
      Resolve: deleteTodo,
    },
    "createUser": &graphql.Field{
      Description: "Create new user",
      Type: userType,
      Args: graphql.FieldConfigArgument{
        "name": &graphql.ArgumentConfig{
          Type: graphql.NewNonNull(graphql.String),
        },
        "email": &graphql.ArgumentConfig{
          Type: graphql.NewNonNull(graphql.String),
        },
        "password": &graphql.ArgumentConfig{
          Type: graphql.NewNonNull(graphql.String),
        },
      },
      Resolve: createUser,
    },
    "updateUser": &graphql.Field{
      Description: "Update existing user",
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
    "deleteUser": &graphql.Field{
      Description: "Delete existing user",
      Type: userType,
      Args: graphql.FieldConfigArgument{
        "id": &graphql.ArgumentConfig{
          Type: graphql.NewNonNull(graphql.Int),
        },
      },
      Resolve: deleteUser,
    },
  },
})
