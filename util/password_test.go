package util

import (
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestPassword(t *testing.T) {
	password := RandomPassword(12)
	hashedPassword1, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword1)

	err = VerifyPassword(hashedPassword1, password)
	require.NoError(t, err)

	incorrectPassword := RandomPassword(12)
	err = VerifyPassword(hashedPassword1, incorrectPassword)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())

	// Compare the bcrypt hash of the same password
	hashedPassword2, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword2)
	require.NotEqual(t, hashedPassword1, hashedPassword2)
}
