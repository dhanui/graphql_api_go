package models

import (
  "time"
)

type User struct {
  Id int `json:"id"`
  Name string `json:"name"`
  Email string `json:"email"`
  CreatedAt time.Time `json:"created_at"`
}

func CreateUser(name string, email string) (newUser User, err error) {
  newUser = User{
    Name: name,
    Email: email,
    CreatedAt: time.Now(),
  }

  tx, err := db.Begin()
  if err != nil {
    return
  }
  defer tx.Rollback()

  stmt, err := db.Prepare("INSERT INTO users(name, email, created_at) VALUES(?, ?, ?)")
  if err != nil {
    return
  }
  defer stmt.Close()

  res, err := stmt.Exec(newUser.Name, newUser.Email, newUser.CreatedAt)
  if err != nil {
    return
  }

  lastId, err := res.LastInsertId()
  if err != nil {
    return
  }
  newUser.Id = int(lastId)

  _, err = res.RowsAffected()
  if err != nil {
    return
  }

  err = tx.Commit()
  return
}

func GetUser(id int) (user User, err error) {
  user = User{}
  err = db.QueryRow("SELECT id, name, email, created_at FROM users WHERE id = ?", id).
    Scan(&user.Id, &user.Name, &user.Email, &user.CreatedAt)
  return
}

func GetUserList() (users []User, err error) {
  rows, err := db.Query("SELECT id, name, email, created_at FROM users")
  if err != nil {
    return
  }
  defer rows.Close()
  for rows.Next() {
    user := User{}
    err = rows.Scan(&user.Id, &user.Name, &user.Email, &user.CreatedAt)
    if err != nil {
      return
    }
    users = append(users, user)
  }
  err = rows.Err()
  return
}
