package models

import (
  "time"
)

type Todo struct {
  Id int `json:"id"`
  Title string `json:"title"`
  Body string `json:"body"`
  UserId int
  CreatedAt time.Time `json:"created_at"`
  UpdatedAt time.Time `json:"updated_at"`
}

func (todo *Todo) Create() (err error) {
  tx, err := db.Begin()
  if err != nil {
    return
  }
  defer tx.Rollback()
  todo.CreatedAt = time.Now()
  todo.UpdatedAt = todo.CreatedAt
  res, err := createTodoStmt.Exec(todo.Title, todo.Body, todo.UserId, todo.CreatedAt, todo.UpdatedAt)
  if err != nil {
    return
  }
  lastId, err := res.LastInsertId()
  if err != nil {
    return
  }
  todo.Id = int(lastId)
  err = tx.Commit()
  return
}

func (todo *Todo) Update() (err error) {
  tx, err := db.Begin()
  if err != nil {
    return
  }
  defer tx.Rollback()
  todo.UpdatedAt = time.Now()
  _, err = updateTodoStmt.Exec(todo.Title, todo.Body, todo.UserId, todo.UpdatedAt, todo.Id)
  if err != nil {
    return
  }
  err = tx.Commit()
  return
}

func (todo *Todo) Delete() (err error) {
  tx, err := db.Begin()
  if err != nil {
    return
  }
  defer tx.Rollback()
  _, err = deleteTodoStmt.Exec(time.Now(), todo.Id)
  if err != nil {
    return
  }
  err = tx.Commit()
  return
}

func GetTodo(id int) (todo Todo, err error) {
  todo = Todo{}
  err = getTodoStmt.QueryRow(id).Scan(&todo.Id, &todo.Title, &todo.Body, &todo.UserId, &todo.CreatedAt, &todo.UpdatedAt)
  return
}

func GetTodoListFilteredByUserId(userId int) (todos []Todo, err error) {
  rows, err := getTodosByUserIdStmt.Query(userId)
  if err != nil {
    return
  }
  defer rows.Close()
  for rows.Next() {
    todo := Todo{}
    err = rows.Scan(&todo.Id, &todo.Title, &todo.Body, &todo.UserId, &todo.CreatedAt, &todo.UpdatedAt)
    if err != nil {
      return
    }
    todos = append(todos, todo)
  }
  err = rows.Err()
  return
}

func GetTodoList() (todos []Todo, err error) {
  rows, err := getAllTodosStmt.Query()
  if err != nil {
    return
  }
  defer rows.Close()
  for rows.Next() {
    todo := Todo{}
    err = rows.Scan(&todo.Id, &todo.Title, &todo.Body, &todo.UserId, &todo.CreatedAt, &todo.UpdatedAt)
    if err != nil {
      return
    }
    todos = append(todos, todo)
  }
  err = rows.Err()
  return
}
