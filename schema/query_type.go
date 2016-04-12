package schema

import (
  "github.com/graphql-go/graphql"
)

var queryType = graphql.NewObject(graphql.ObjectConfig{
  Name: "Query",
  Fields: graphql.Fields{
    "me": &graphql.Field{
      Description: "Get current user",
      Type: userType,
      Resolve: getCurrentUser,
    },
    "todo": &graphql.Field{
      Description: "Get a todo",
      Type: todoType,
      Args: graphql.FieldConfigArgument{
        "id": &graphql.ArgumentConfig{
          Type: graphql.NewNonNull(graphql.Int),
        },
      },
      Resolve: getTodo,
    },
    "todoList": &graphql.Field{
      Description: "Get list of todo",
      Type: graphql.NewList(todoType),
      Args: graphql.FieldConfigArgument{
        "user_id": &graphql.ArgumentConfig{
          Type: graphql.Int,
        },
      },
      Resolve: getTodoList,
    },
    "user": &graphql.Field{
      Description: "Get a user",
      Type: userType,
      Args: graphql.FieldConfigArgument{
        "id": &graphql.ArgumentConfig{
          Type: graphql.NewNonNull(graphql.Int),
        },
      },
      Resolve: getUser,
    },
    "userList": &graphql.Field{
      Description: "Get list of user",
      Type: graphql.NewList(userType),
      Resolve: getUserList,
    },
  },
})
