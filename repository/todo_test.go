package repository

import (
  "testing"

  "github.com/stretchr/testify/require"
)

var todo Todo

func TestSuccessCreateTodoAndSuccessGetTodo(t *testing.T) {
  todo = Todo{
    Title: "Create To Do",
    Body: "To do for creation testing",
    UserId: 1,
  }
  err := todo.Create()
  require.Nil(t, err)
  getTodo, err := GetTodo(todo.Id)
  require.Nil(t, err)
  require.Equal(t, todo, getTodo)
}

func TestSuccessUpdateTodoAndSuccessGetTodo(t *testing.T) {
  todo.Title = "Updated title"
  todo.Body = "Updated body"
  err := todo.Update()
  require.Nil(t, err)
  getTodo, err := GetTodo(todo.Id)
  require.Nil(t, err)
  require.Equal(t, todo, getTodo)
}

func TestSuccessGetTodoListFilteredByUserId(t *testing.T) {
  todos, err := GetTodoListFilteredByUserId(1)
  require.Nil(t, err)
  require.Equal(t, 1, len(todos))
  require.Equal(t, todo, todos[0])
}

func TestSuccessGetTodoList(t *testing.T) {
  todos, err := GetTodoList()
  require.Nil(t, err)
  require.Equal(t, 1, len(todos))
  require.Equal(t, todo, todos[0])
}

func TestSuccessDeleteTodoAndFailGetTodo(t *testing.T) {
  err := todo.Delete()
  _, err = GetTodo(todo.Id)
  require.NotNil(t, err)
}
