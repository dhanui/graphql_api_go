package models

import (
  "time"
)

type User struct {
  Id int `json:"id"`
  Name string `json:"name"`
  Email string `json:"email"`
  CreatedAt time.Time `json:"created_at"`
  UpdatedAt time.Time `json:"updated_at"`
}

func (user *User) Create() (err error) {
  tx, err := db.Begin()
  if err != nil {
    return
  }
  defer tx.Rollback()

  stmt, err := db.Prepare("INSERT INTO users(name, email, created_at, updated_at) VALUES(?, ?, ?, ?)")
  if err != nil {
    return
  }
  defer stmt.Close()

  user.CreatedAt = time.Now()
  user.UpdatedAt = user.CreatedAt
  res, err := stmt.Exec(user.Name, user.Email, user.CreatedAt, user.UpdatedAt)
  if err != nil {
    return
  }

  lastId, err := res.LastInsertId()
  if err != nil {
    return
  }
  user.Id = int(lastId)

  err = tx.Commit()
  return
}

func GetUser(id int) (user User, err error) {
  user = User{}
  err = db.QueryRow("SELECT id, name, email, created_at, updated_at FROM users WHERE id = ?", id).
    Scan(&user.Id, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt)
  return
}

func GetUserList() (users []User, err error) {
  rows, err := db.Query("SELECT id, name, email, created_at, updated_at FROM users")
  if err != nil {
    return
  }
  defer rows.Close()
  for rows.Next() {
    user := User{}
    err = rows.Scan(&user.Id, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt)
    if err != nil {
      return
    }
    users = append(users, user)
  }
  err = rows.Err()
  return
}
