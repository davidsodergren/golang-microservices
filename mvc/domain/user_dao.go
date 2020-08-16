package domain

import (
	"davidsodergren/golang-microservices/mvc/utils"
	"fmt"
	"log"
	"net/http"
)

var(
	users = map[int64]*User{
		123: &User{Id: 123, FirstName: "Fede", LastName: "Leon", Email: "myemail@gmail.com"},
	}
	UserDao userDaoInterface
)

type userDaoInterface interface {
	GetUser(int64) (*User, *utils.ApplicationError)
}

func init()  {
	UserDao = &userDao{}
}


type userDao struct {

}



func (u *userDao)GetUser(userId int64) (*User, *utils.ApplicationError) {
	log.Println("We are accessing the database")
	if user := users[userId]; user != nil {
		return user, nil
	}
		return nil, &utils.ApplicationError{
			Message: fmt.Sprintf("User with id %v was not found", userId),
			StatusCode: http.StatusNotFound,
			Code: "not_found",
		}
}