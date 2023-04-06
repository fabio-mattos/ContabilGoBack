package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("81094965987", "Fabio Mattos", "softvision.fpolis@gmail.com", "123456", 1)
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.IDUsuario)
	assert.NotEmpty(t, user.DeSenha)
	assert.Equal(t, "Fabio Mattos", user.NmCompleto)
	assert.Equal(t, "softvision.fpolis@gmail.com", user.DeEmail)
	assert.Equal(t, "81094965987", user.DeLogin)
}

func TestUser_ValidatePassword(t *testing.T) {
	user, err := NewUser("81094965987", "Fabio Mattos", "softvision.fpolis@gmail.com", "123456", 1)
	assert.Nil(t, err)
	assert.True(t, user.ValidatePassword("123456"))
	assert.False(t, user.ValidatePassword("1234567"))
	assert.NotEqual(t, "123456", user.DeEmail)
}
