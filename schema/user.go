package schema

import (
  "github.com/graphql-go/graphql"

  "../models"
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
  },
})

func createUser(params graphql.ResolveParams) (interface{}, error) {
  name, _ := params.Args["name"].(string)
  email, _ := params.Args["email"].(string)

  newUser, err := models.CreateUser(name, email)
  if err != nil {
    return nil, err
  } else {
    return newUser, nil
  }
}

func getUser(params graphql.ResolveParams) (interface{}, error) {
  id, _ := params.Args["id"].(int)

  user, err := models.GetUser(id)
  if err != nil {
    return nil, err
  } else {
    return user, nil
  }
}

func getUserList(params graphql.ResolveParams) (interface{}, error) {
  users, err := models.GetUserList()
  if err != nil {
    return nil, err
  } else {
    return users, nil
  }
}