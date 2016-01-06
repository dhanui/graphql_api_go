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
}

func CreateTodo(title string, body string) (newTodo Todo, err error) {
  newTodo = Todo{
    Title: title,
    Body: body,
    CreatedAt: time.Now(),
  }

  tx, err := db.Begin()
  if err != nil {
    return
  }
  defer tx.Rollback()

  stmt, err := db.Prepare("INSERT INTO todos(title, body, created_at) VALUES(?, ?, ?)")
  if err != nil {
    return
  }
  defer stmt.Close()

  res, err := stmt.Exec(newTodo.Title, newTodo.Body, newTodo.CreatedAt)

  lastId, err := res.LastInsertId()
  if err != nil {
    return
  }
  newTodo.Id = int(lastId)

  _, err = res.RowsAffected()
  if err != nil {
    return
  }

  err = tx.Commit()
  return
}

func GetTodo(id int) (todo Todo, err error) {
  todo = Todo{}
  err = db.QueryRow("SELECT id, title, body, user_id, created_at FROM todos WHERE id = ?", id).
    Scan(&todo.Id, &todo.Title, &todo.Body, &todo.UserId, &todo.CreatedAt)
  return
}

func GetTodoList() (todos []Todo, err error) {
  rows, err := db.Query("SELECT id, title, body, user_id, created_at FROM todos")
  if err != nil {
    return
  }
  defer rows.Close()
  for rows.Next() {
    todo := Todo{}
    err = rows.Scan(&todo.Id, &todo.Title, &todo.Body, &todo.UserId, &todo.CreatedAt)
    if err != nil {
      return
    }
    todos = append(todos, todo)
  }
  err = rows.Err()
  return
}
