package models

import (
  "testing"

  "github.com/stretchr/testify/require"
)

var user User

func TestSuccessCreateUserAndSuccessGetUser(t *testing.T) {
  user = User{
    Name: "Test User",
    Email: "test@example.com",
  }
  err := user.Create("password")
  require.Nil(t, err)
  getUser, err := GetUser(user.Id)
  require.Nil(t, err)
  require.Equal(t, user, getUser)
}

func TestSuccessUpdateUserAndSuccessGetUser(t *testing.T) {
  user.Name = "Update Name"
  user.Email = "updated@example.com"
  err := user.Update()
  require.Nil(t, err)
  getUser, err := GetUser(user.Id)
  require.Nil(t, err)
  require.Equal(t, user, getUser)
}

func TestSuccessGetUserList(t *testing.T) {
  users, err := GetUserList()
  require.Nil(t, err)
  require.Equal(t, 2, len(users))
  require.Equal(t, user, users[1])
}

func TestSuccessAuthenticateUser(t *testing.T) {
  authUser, ok := AuthenticateUser("updated@example.com", "password")
  require.Equal(t, true, ok)
  require.Equal(t, user, authUser)
}

func TestFailAuthenticateUser(t *testing.T) {
  _, ok := AuthenticateUser("updated@example.com", "p455w0rd")
  require.Equal(t, false, ok)
}

func TestSuccessDeleteUserAndFailGetUser(t *testing.T) {
  err := user.Delete()
  require.Nil(t, err)
  _, err = GetUser(user.Id)
  require.NotNil(t, err)
}
