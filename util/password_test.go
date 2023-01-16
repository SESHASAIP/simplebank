package util

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestPassword(t *testing.T) {
	password := RandomString(6)
	hashed, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashed)

	err = CheckHashedPassword(hashed, password)
	require.NoError(t, err)

	WrongPassword := RandomString(6)
	err = CheckHashedPassword(hashed, WrongPassword)
	fmt.Println(err.Error())
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())

	hashed2, err2 := HashPassword(password)
	require.NoError(t, err2)
	require.NotEmpty(t, hashed2)
	require.NotEqual(t, hashed, hashed2)

}
