package models

import (
  "time"

  "github.com/dhanui/graphql_api_go/helpers"
)

type User struct {
  Id int `json:"id"`
  Name string `json:"name"`
  Email string `json:"email"`
  CreatedAt time.Time `json:"created_at"`
  UpdatedAt time.Time `json:"updated_at"`
}

func (user *User) Create(password string) (err error) {
  passwordHash, err := helpers.CreateHash(password)
  if err != nil {
    return
  }
  tx, err := db.Begin()
  if err != nil {
    return
  }
  defer tx.Rollback()
  user.CreatedAt = time.Now()
  user.UpdatedAt = user.CreatedAt
  res, err := tx.Exec("INSERT INTO users(name, email, password_hash, created_at, updated_at) VALUES(?, ?, ?, ?, ?)", user.Name, user.Email, passwordHash, user.CreatedAt, user.UpdatedAt)
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

func (user *User) Update() (err error) {
  tx, err := db.Begin()
  if err != nil {
    return
  }
  defer tx.Rollback()
  user.UpdatedAt = time.Now()
  _, err = tx.Exec("UPDATE users SET name = ?, email = ?, updated_at = ? WHERE id = ? AND deleted_at IS NULL", user.Name, user.Email, user.UpdatedAt, user.Id)
  if err != nil {
    return
  }
  err = tx.Commit()
  return
}

func (user *User) Delete() (err error) {
  tx, err := db.Begin()
  if err != nil {
    return
  }
  defer tx.Rollback()
  _, err = tx.Exec("UPDATE users SET deleted_at = ? WHERE id = ? AND deleted_at IS NULL", time.Now(), user.Id)
  if err != nil {
    return
  }
  err = tx.Commit()
  return
}

func GetUser(id int) (user User, err error) {
  user = User{}
  err = db.QueryRow("SELECT id, name, email, created_at, updated_at FROM users WHERE id = ? AND deleted_at IS NULL LIMIT 1", id).Scan(&user.Id, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt)
  return
}

func GetUserList() (users []User, err error) {
  rows, err := db.Query("SELECT id, name, email, created_at, updated_at FROM users WHERE deleted_at IS NULL")
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

func AuthenticateUser(username string, password string) (user User, ok bool) {
  var passwordHash string
  user = User{}
  err := db.QueryRow("SELECT id, name, email, created_at, updated_at, password_hash FROM users WHERE email = ? LIMIT 1", username).Scan(&user.Id, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt, &passwordHash)
  if err != nil {
    return user, false
  } else {
    return user, helpers.ValidateHash(password, passwordHash)
  }
}
