package schema

import (
  "github.com/graphql-go/graphql"

  "github.com/dhanui/graphql_api_go/models"
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
      Resolve: getUserFromTodo,
    },
    "created_at": &graphql.Field{
      Type: graphql.String,
    },
    "updated_at": &graphql.Field{
      Type: graphql.String,
    },
  },
})

func createTodo(params graphql.ResolveParams) (interface{}, error) {
  title, _ := params.Args["title"].(string)
  body, _ := params.Args["body"].(string)
  userId, _ := params.Args["user_id"].(int)
  newTodo := models.Todo{
    Title: title,
    Body: body,
    UserId: userId,
  }
  err := newTodo.Create()
  if err != nil {
    return nil, err
  } else {
    return newTodo, nil
  }
}

func updateTodo(params graphql.ResolveParams) (interface{}, error) {
  id, _ := params.Args["id"].(int)
  todo, err := models.GetTodo(id)
  if err != nil {
    return nil, err
  }
  title, ok := params.Args["title"].(string)
  if ok {
    todo.Title = title
  }
  body, ok := params.Args["body"].(string)
  if ok {
    todo.Body = body
  }
  userId, ok := params.Args["user_id"].(int)
  if ok {
    todo.UserId = userId
  }
  err = todo.Update()
  if err != nil {
    return nil, err
  } else {
    return todo, nil
  }
}

func deleteTodo(params graphql.ResolveParams) (interface{}, error) {
  id, _ := params.Args["id"].(int)
  todo, err := models.GetTodo(id)
  if err != nil {
    return nil, err
  }
  err = todo.Delete()
  if err != nil {
    return nil, err
  } else {
    return todo, nil
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
  userId, ok := params.Args["user_id"].(int)
  var todos []models.Todo
  var err error
  if ok {
    todos, err = models.GetTodoListFilteredByUserId(userId)
  } else {
    todos, err = models.GetTodoList()
  }
  if err != nil {
    return nil, err
  } else {
    return todos, nil
  }
}
