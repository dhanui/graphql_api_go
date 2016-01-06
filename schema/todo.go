package schema

import (
  "github.com/graphql-go/graphql"

  "../models"
)

var todoType = graphql.NewObject(graphql.ObjectConfig{
  Name: "Todo",
  Fields: graphql.Fields{
    "id": &graphql.Field{
      Type: graphql.Int,
    },
    "title": &graphql.Field{
      Type: graphql.String,
    },
    "body": &graphql.Field{
      Type: graphql.String,
    },
    "user": &graphql.Field{
      Type: userType,
      Resolve: func(params graphql.ResolveParams) (interface{}, error) {
        todo, _ := params.Source.(models.Todo)
        user, err := models.GetUser(todo.UserId)
        if err != nil {
          return nil, err
        } else {
          return user, nil
        }
      },
    },
    "created_at": &graphql.Field{
      Type: graphql.String,
    },
  },
})

func createTodo(params graphql.ResolveParams) (interface{}, error) {
  title, _ := params.Args["title"].(string)
  body, _ := params.Args["body"].(string)
  userId, _ := params.Args["user_id"].(int)

  newTodo, err := models.CreateTodo(title, body, userId)
  if err != nil {
    return nil, err
  } else {
    return newTodo, nil
  }
}

func getTodo(params graphql.ResolveParams) (interface{}, error) {
  id, _ := params.Args["id"].(int)
  todo, err := models.GetTodo(id)
  if err != nil {
    return nil, err
  } else {
    return todo, nil
  }
}

func getTodoList(params graphql.ResolveParams) (interface{}, error) {
  todos, err := models.GetTodoList()
  if err != nil {
    return nil, err
  } else {
    return todos, nil
  }
}
