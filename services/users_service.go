package services

import (
	"davidsodergren/golang-microservices/domain"
	"davidsodergren/golang-microservices/utils"
)

func GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userId)
}