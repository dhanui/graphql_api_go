package helpers

import (
  "golang.org/x/crypto/bcrypt"
)

func CreateHash(password string) (string, error) {
  passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
  if err != nil {
    return "", err
  } else {
    return string(passwordHash), nil
  }
}

func ValidateHash(password string, passwordHash string) bool {
  return bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password)) == nil
}
