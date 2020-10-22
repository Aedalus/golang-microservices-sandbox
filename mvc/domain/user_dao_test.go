package domain

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestGetUserUserNotFound(t *testing.T) {
	user, err := GetUser(0)

	assert.Nil(t, user)
	assert.NotNil(t,err)
	assert.EqualValues(t,err.StatusCode, http.StatusNotFound)
	assert.EqualValues(t,err.Code,"not_found")
	assert.EqualValues(t, err.Message, "user 0 not found")
}

func TestGetUserFound(t *testing.T) {
	user, err := GetUser(123)

	assert.Nil(t,err)
	assert.NotNil(t,user)
	assert.EqualValues(t, user.Id, uint64(123))
	assert.EqualValues(t, user.FirstName, "Alex")
	assert.EqualValues(t, user.LastName, "Higgins")
	assert.EqualValues(t, user.Email, "TheAlexanderHiggins@gmail.com")
}