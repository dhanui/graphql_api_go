package schema

import (
  "github.com/graphql-go/graphql"
)

var rootQuery = graphql.NewObject(graphql.ObjectConfig{
  Name: "RootQuery",
  Fields: graphql.Fields{
    "todo": &graphql.Field{
      Type: todoType,
      Args: graphql.FieldConfigArgument{
        "id": &graphql.ArgumentConfig{
          Type: graphql.Int,
        },
      },
      Resolve: getTodo,
    },
    "todoList": &graphql.Field{
      Type: graphql.NewList(todoType),
      Description: "List of todos",
      Resolve: getTodoList,
    },
  },
})
