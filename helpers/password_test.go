package helpers

import (
  "testing"

  "github.com/stretchr/testify/require"
)

func TestSuccessCreateHash(t *testing.T) {
  passwordHash, err := CreateHash("password01")
  require.Nil(t, err)
  require.NotEqual(t, "", passwordHash)
}

func TestSuccessValidateHash(t *testing.T) {
  password := "password01"
  passwordHash, err := CreateHash(password)
  require.Nil(t, err)
  require.Equal(t, true, ValidateHash(password, passwordHash))
}

func TestFailValidateHash(t *testing.T) {
  password, passwordHash := "password01", "randomHash"
  require.Equal(t, false, ValidateHash(password, passwordHash))
}
