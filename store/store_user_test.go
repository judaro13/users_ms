package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
	// _ "github.com/lib/pq"
)

func TestValidateUser(t *testing.T) {
	user := new(User)
	assert.Equal(t, ErrNameRequired, ValidateUser(user))

	user.Name = "name"
	assert.Equal(t, ErrEmailRequired, ValidateUser(user))

	user.Email = "email"
	assert.Equal(t, ErrPasswordRequired, ValidateUser(user))

	user.Password = "asdf45qwfas4"
	assert.Equal(t, ErrPhoneRequired, ValidateUser(user))

	user.PhoneNumber = "54321"
	assert.Equal(t, nil, ValidateUser(user))
}
