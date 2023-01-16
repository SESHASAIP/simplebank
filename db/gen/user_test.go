package gen

import (
	"context"

	"simplebank/util"
	"testing"

	"github.com/stretchr/testify/require"
)

func CreateRandomUser(t *testing.T) User {
	password := util.RandomString(6)
	s, e := util.HashPassword(password)
	require.NoError(t, e)
	arg := CreateUserParams{
		Username:       util.RandomOwner(),
		HashedPassword: s,
		FullName:       util.RandomOwner(),
		Email:          util.RandomEmail(),
	}
	out, err := TestQueries.CreateUser(context.Background(), arg)
	if err != nil {
		require.NoError(t, err)
		require.NotEmpty(t, out)
		require.Equal(t, arg.Username, out.Username)
		require.Equal(t, arg.HashedPassword, out.HashedPassword)
		require.Equal(t, arg.FullName, out.FullName)
		require.Equal(t, arg.Email, out.Email)
		require.NotZero(t, out.CreatedAt)

	}
	return out

}
func TestCreateUser(t *testing.T) {
	CreateRandomUser(t)
}

func TestGetUser(t *testing.T) {
	account1 := CreateRandomUser(t)
	account2, err := TestQueries.GetUser(context.Background(), account1.Username)
	require.NoError(t, err)
	require.Equal(t, account1.Username, account2.Username)
	require.Equal(t, account1.Email, account2.Email)
	require.Equal(t, account1.HashedPassword, account2.HashedPassword)

}
