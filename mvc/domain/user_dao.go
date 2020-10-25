package domain

import (
	"fmt"
	"mvc/utils"
	"net/http"
)

type UserDaoInterface interface {
	GetUser(id int64) (*User, *utils.ApplicationError)
}

type userDao struct{}

var (
	users = map[int64]*User{
		123: &User{
			Id:        123,
			FirstName: "Alex",
			LastName:  "Higgins",
			Email:     "TheAlexanderHiggins@gmail.com",
		},
	}
	UserDao UserDaoInterface
)

func init() {
	UserDao = userDao{}
}

func (u userDao) GetUser(id int64) (*User, *utils.ApplicationError) {
	if user := users[id]; user != nil {
		return user, nil
	}
	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("user %v not found", id),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}
}
