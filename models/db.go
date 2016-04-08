package models

import (
  "encoding/json"
  "os"
  "fmt"
  "database/sql"

  "github.com/mattes/migrate/migrate"
  _ "github.com/go-sql-driver/mysql"
  _ "github.com/mattes/migrate/driver/mysql"
)

type DBConf struct {
  Host string `json:"host"`
  Port string `json:"port"`
  User string `json:"user"`
  Pass string `json:"pass"`
  Name string `json:"name"`
}

var db *sql.DB

var (
  createTodoStmt *sql.Stmt
  updateTodoStmt *sql.Stmt
  deleteTodoStmt *sql.Stmt
  getTodoStmt *sql.Stmt
  getTodosByUserIdStmt *sql.Stmt
  getAllTodosStmt *sql.Stmt
)

var (
  createUserStmt *sql.Stmt
  updateUserStmt *sql.Stmt
  deleteUserStmt *sql.Stmt
  getUserStmt *sql.Stmt
  getAllUserStmt *sql.Stmt
  getUserByEmailStmt *sql.Stmt
)

func InitDBConnection() (err error) {
  connString, err := loadConnectionString()
  if err != nil {
    return
  }
  db, err = sql.Open("mysql", connString)
  if err != nil {
    return
  }
  fmt.Printf("Configuration loaded\nTesting database connection...\n")
  err = db.Ping()
  if err != nil {
    return
  }
  fmt.Printf("Connection successful\nPreparing statements...\n")
  err = prepareTodoStmts()
  if err != nil {
    return
  }
  err = prepareUserStmts()
  if err != nil {
    return
  }
  fmt.Printf("Statements preparation successful\n")
  return
}

func Migrate() (errors []error) {
  connString, err := loadConnectionString()
  if err != nil {
    return append(errors, err)
  }
  url := fmt.Sprintf("mysql://%s", connString)
  errors, _ = migrate.UpSync(url, "./migrations")
  return
}

func Rollback() (errors []error) {
  connString, err := loadConnectionString()
  if err != nil {
    return append(errors, err)
  }
  url := fmt.Sprintf("mysql://%s", connString)
  errors, _ = migrate.DownSync(url, "./migrations")
  return
}

func loadConnectionString() (connString string, err error) {
  fmt.Printf("Loading database configuration...\n")
  file, err := os.Open("database.json")
  if err != nil {
    return
  }
  decoder := json.NewDecoder(file)
  dbConf := DBConf{}
  err = decoder.Decode(&dbConf)
  if err != nil {
    return
  }
  return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
    dbConf.User, dbConf.Pass, dbConf.Host, dbConf.Port, dbConf.Name), nil
}

func prepareTodoStmts() (err error) {
  createTodoStmt, err = db.Prepare("INSERT INTO todos(title, body, user_id, created_at, updated_at) VALUES(?, ?, ?, ?, ?)")
  if err != nil {
    return
  }
  updateTodoStmt, err = db.Prepare("UPDATE todos SET title = ?, body = ?, user_id = ?, updated_at = ? WHERE id = ? AND deleted_at IS NULL")
  if err != nil {
    return
  }
  deleteTodoStmt, err = db.Prepare("UPDATE todos SET deleted_at = ? WHERE id = ? AND deleted_at IS NULL")
  if err != nil {
    return
  }
  getTodoStmt, err = db.Prepare("SELECT id, title, body, user_id, created_at, updated_at FROM todos WHERE id = ? AND deleted_at IS NULL LIMIT 1")
  if err != nil {
    return
  }
  getTodosByUserIdStmt, err = db.Prepare("SELECT id, title, body , user_id, created_at, updated_at FROM todos WHERE user_id = ? AND deleted_at IS NULL")
  if err != nil {
    return
  }
  getAllTodosStmt, err = db.Prepare("SELECT id, title, body, user_id, created_at, updated_at FROM todos WHERE deleted_at IS NULL")
  return
}

func prepareUserStmts() (err error) {
  createUserStmt, err = db.Prepare("INSERT INTO users(name, email, password_hash, created_at, updated_at) VALUES(?, ?, ?, ?, ?)")
  if err != nil {
    return
  }
  updateUserStmt, err = db.Prepare("UPDATE users SET name = ?, email = ?, updated_at = ? WHERE id = ? AND deleted_at IS NULL")
  if err != nil {
    return
  }
  deleteUserStmt, err = db.Prepare("UPDATE users SET deleted_at = ? WHERE id = ? AND deleted_at IS NULL")
  if err != nil {
    return
  }
  getUserStmt, err = db.Prepare("SELECT id, name, email, created_at, updated_at FROM users WHERE id = ? AND deleted_at IS NULL LIMIT 1")
  if err != nil {
    return
  }
  getAllUserStmt, err = db.Prepare("SELECT id, name, email, created_at, updated_at FROM users WHERE deleted_at IS NULL")
  if err != nil {
    return
  }
  getUserByEmailStmt, err = db.Prepare("SELECT id, name, email, created_at, updated_at, password_hash FROM users WHERE email = ? LIMIT 1")
  return
}
