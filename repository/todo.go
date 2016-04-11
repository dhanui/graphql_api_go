package repository

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
  todo.CreatedAt = time.Now().UTC().Round(time.Second)
  todo.UpdatedAt = todo.CreatedAt
  res, err := tx.Exec("INSERT INTO todos(title, body, user_id, created_at, updated_at) VALUES(?, ?, ?, ?, ?)", todo.Title, todo.Body, todo.UserId, todo.CreatedAt, todo.UpdatedAt)
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
  todo.UpdatedAt = time.Now().UTC().Round(time.Second)
  _, err = tx.Exec("UPDATE todos SET title = ?, body = ?, user_id = ?, updated_at = ? WHERE id = ? AND deleted_at IS NULL", todo.Title, todo.Body, todo.UserId, todo.UpdatedAt, todo.Id)
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
  _, err = tx.Exec("UPDATE todos SET deleted_at = ? WHERE id = ? AND deleted_at IS NULL", time.Now().UTC().Round(time.Second), todo.Id)
  if err != nil {
    return
  }
  err = tx.Commit()
  return
}

func GetTodo(id int) (todo Todo, err error) {
  todo = Todo{}
  err = db.QueryRow("SELECT id, title, body, user_id, created_at, updated_at FROM todos WHERE id = ? AND deleted_at IS NULL LIMIT 1", id).Scan(&todo.Id, &todo.Title, &todo.Body, &todo.UserId, &todo.CreatedAt, &todo.UpdatedAt)
  return
}

func GetTodoListFilteredByUserId(userId int) (todos []Todo, err error) {
  rows, err := db.Query("SELECT id, title, body , user_id, created_at, updated_at FROM todos WHERE user_id = ? AND deleted_at IS NULL", userId)
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
  rows, err := db.Query("SELECT id, title, body, user_id, created_at, updated_at FROM todos WHERE deleted_at IS NULL")
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
