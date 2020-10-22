package services

import (
	"mvc/domain"
	"mvc/utils"
)
type itemService struct {}

var (
	ItemService itemService
)

func (i itemService) GetItems(id string) (*domain.Item, *utils.ApplicationError) {
	return nil, &utils.ApplicationError{
		Message:    "IMPLEMENTE_ME",
		StatusCode: 0,
		Code:       "foo",
	}
}
