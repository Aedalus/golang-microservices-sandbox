package services

import (
	"mvc/domain"
	"mvc/utils"
)

func GetUser(id int64) (*domain.User, *utils.ApplicationError){
	return domain.GetUser(id)
}