package services

import (
	"davidsodergren/golang-microservices/domain"
	"davidsodergren/golang-microservices/utils"
)

type userService struct {
}

var (
	UserService userService
)

func (u *userService) GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return domain.UserDao.GetUser(userId)
}