package services

import (
	"mvc/domain"
	"mvc/utils"
)

type userService struct {}

var (
	UserService userService
)

func (u *userService) GetUser(id int64) (*domain.User, *utils.ApplicationError){
	return domain.UserDao.GetUser(id)
}