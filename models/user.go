package models

import (
  "time"

  "../helpers"
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
  res, err := createUserStmt.Exec(user.Name, user.Email, passwordHash, user.CreatedAt, user.UpdatedAt)
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
  _, err = updateUserStmt.Exec(user.Name, user.Email, user.UpdatedAt, user.Id)
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
  _, err = deleteUserStmt.Exec(time.Now(), user.Id)
  if err != nil {
    return
  }
  err = tx.Commit()
  return
}

func GetUser(id int) (user User, err error) {
  user = User{}
  err = getUserStmt.QueryRow(id).Scan(&user.Id, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt)
  return
}

func GetUserList() (users []User, err error) {
  rows, err := getAllUserStmt.Query()
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

func AuthenticateUser(username string, password string) bool {
  var passwordHash string
  err := getUserByEmailStmt.QueryRow(username).Scan(&passwordHash)
  if err != nil {
    return false
  } else {
    return helpers.ValidateHash(password, passwordHash)
  }
}
