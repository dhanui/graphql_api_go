package models

import (
  "time"
)

type Todo struct {
  Id int `json:"id"`
  Title string `json:"title"`
  Body string `json:"body"`
  AddedOn time.Time `json:"added_on"`
}

func CreateTodo(title string, body string) (*Todo, error) {
  tx, err := db.Begin()
  if err != nil {
    return nil, err
  }
  defer tx.Rollback()

  stmt, err := db.Prepare("INSERT INTO todos(title, body, added_on) VALUES(?, ?, ?)")
  if err != nil {
    return nil, err
  }
  defer stmt.Close()

  addedOn := time.Now()
  res, err := stmt.Exec(title, body, addedOn)

  lastId, err := res.LastInsertId()
  if err != nil {
    return nil, err
  }

  _, err = res.RowsAffected()
  if err != nil {
    return nil, err
  }

  err = tx.Commit()
  if err != nil {
    return nil, err
  }

  newTodo := Todo{
    Id: int(lastId),
    Title: title,
    Body: body,
    AddedOn: addedOn,
  }

  return &newTodo, nil
}

func GetTodo(id int) (*Todo, error) {
  todo := Todo{}

  err := db.QueryRow("SELECT * FROM todos WHERE id = ?", id).Scan(&todo.Id, &todo.Title, &todo.Body, &todo.AddedOn)
  if err != nil {
    return nil, err
  }

  return &todo, nil
}

func GetTodoList() ([]Todo, error) {
  var todos []Todo

  rows, err := db.Query("SELECT * FROM todos")
  if err != nil {
    return nil, err
  }
  defer rows.Close()
  for rows.Next() {
    todo := Todo{}
    err := rows.Scan(&todo.Id, &todo.Title, &todo.Body, &todo.AddedOn)
    if err != nil {
      return nil, err
    }
    todos = append(todos, todo)
  }
  err = rows.Err()
  if err != nil {
    return nil, err
  }

  return todos, nil
}
