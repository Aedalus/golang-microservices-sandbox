package services

import (
	"github.com/stretchr/testify/assert"
	"mvc/domain"
	"mvc/utils"
	"net/http"
	"testing"
)

type usersDaoMock struct{}

var userDaoMock usersDaoMock
var getUserFunction func(id int64) (*domain.User, *utils.ApplicationError)

func (u usersDaoMock) GetUser(id int64) (*domain.User, *utils.ApplicationError) {
	return getUserFunction(id)
}

func init() {
	domain.UserDao = &usersDaoMock{}
}

func TestUserService_GetUser_NotFoundInDatabase(t *testing.T) {

	// Stub get user function
	getUserFunction = func(id int64) (*domain.User, *utils.ApplicationError) {
		return nil, &utils.ApplicationError{
			Message:    "user 0 does not exist",
			StatusCode: http.StatusNotFound,
			Code:       "not_found",
		}
	}

	// Set up user
	user, err := UsersService.GetUser(0)

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "user 0 does not exist", err.Message)
}

func TestUserService_GetUser_FoundInDatabase(t *testing.T) {

	// Stub get user function
	getUserFunction = func(id int64) (*domain.User, *utils.ApplicationError) {
		return &domain.User{
			Id:        123,
			FirstName: "Josh",
			LastName:  "Bevers",
			Email:     "JoshBevers@gmail.com",
		}, nil
	}

	// Set up user
	user, err := UsersService.GetUser(0)

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, 123, user.Id)
	assert.EqualValues(t, "Josh", user.FirstName)
	assert.EqualValues(t, "Bevers", user.LastName)
	assert.EqualValues(t, "JoshBevers@gmail.com", user.Email)
}
