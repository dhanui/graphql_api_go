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

func GetUser(id int) (user User, err error) {
  user = User{}
  err = db.QueryRow("SELECT * FROM users WHERE id = ?", id).
    Scan(&user.Id, &user.Name, &user.Email, &user.CreatedAt)
  return
}

func GetUserList() (users []User, err error) {
  rows, err := db.Query("SELECT * FROM users")
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
