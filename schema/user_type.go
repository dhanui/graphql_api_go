package schema

import (
  "github.com/graphql-go/graphql"

  "github.com/dhanui/graphql_api_go/repository"
)

var userType = graphql.NewObject(graphql.ObjectConfig{
  Name: "User",
  Fields: graphql.Fields{
    "id": &graphql.Field{
      Type: graphql.Int,
    },
    "name": &graphql.Field{
      Type: graphql.String,
    },
    "email": &graphql.Field{
      Type: graphql.String,
    },
    "created_at": &graphql.Field{
      Type: graphql.String,
    },
    "updated_at": &graphql.Field{
      Type: graphql.String,
    },
  },
})

func getCurrentUser(params graphql.ResolveParams) (interface{}, error) {
  return params.Context.Value("currentUser"), nil
}

func createUser(params graphql.ResolveParams) (interface{}, error) {
  name, _ := params.Args["name"].(string)
  email, _ := params.Args["email"].(string)
  password, _ := params.Args["password"].(string)
  newUser := repository.User{
    Name: name,
    Email: email,
  }
  err := newUser.Create(password)
  if err != nil {
    return nil, err
  } else {
    return newUser, nil
  }
}

func updateUser(params graphql.ResolveParams) (interface{}, error) {
  id, _ := params.Args["id"].(int)
  user, err := repository.GetUser(id)
  if err != nil {
    return nil, err
  }
  name, ok := params.Args["name"].(string)
  if ok {
    user.Name = name
  }
  email, ok := params.Args["email"].(string)
  if ok {
    user.Email = email
  }
  err = user.Update()
  if err != nil {
    return nil, err
  } else {
    return user, nil
  }
}

func deleteUser(params graphql.ResolveParams) (interface{}, error) {
  id, _ := params.Args["id"].(int)
  user, err := repository.GetUser(id)
  if err != nil {
    return nil, err
  }
  err = user.Delete()
  if err != nil {
    return nil, err
  } else {
    return user, nil
  }
}

func getUser(params graphql.ResolveParams) (interface{}, error) {
  id, _ := params.Args["id"].(int)

  user, err := repository.GetUser(id)
  if err != nil {
    return nil, err
  } else {
    return user, nil
  }
}

func getUserFromTodo(params graphql.ResolveParams) (interface{}, error) {
  todo, _ := params.Source.(repository.Todo)
  user, err := repository.GetUser(todo.UserId)
  if err != nil {
    return nil, err
  } else {
    return user, nil
  }
}

func getUserList(params graphql.ResolveParams) (interface{}, error) {
  users, err := repository.GetUserList()
  if err != nil {
    return nil, err
  } else {
    return users, nil
  }
}
