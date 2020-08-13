package services

import (
	"davidsodergren/golang-microservices/domain"
	"davidsodergren/golang-microservices/utils"
	"net/http"
)

type itemService struct{}

var (
	ItemService itemService
)

func (i *itemService) GetItem(itemId string) (*domain.Item, *utils.ApplicationError) {
	return nil, &utils.ApplicationError{
		Message: "Implement me",
		StatusCode: http.StatusInternalServerError,
		Code: "0",
	}

}
