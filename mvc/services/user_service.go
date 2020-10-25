package services

import (
	"mvc/domain"
	"mvc/utils"
)

type userService struct {
	userDao domain.UserDaoInterface
}

var (
	UsersService userService
)

func (u *userService) GetUser(id int64) (*domain.User, *utils.ApplicationError){
	return u.userDao.GetUser(id)
}