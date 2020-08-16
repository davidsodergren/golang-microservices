package services

import (
	"davidsodergren/golang-microservices/mvc/domain"
	"davidsodergren/golang-microservices/mvc/utils"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)


var(
	userDaoMock     usersDaoMock
	getUserFunction func(userId int64) (*domain.User, *utils.ApplicationError)
	)

func init() {
	fmt.Println("Does this run??????")
	domain.UserDao = &userDaoMock
}

type usersDaoMock struct {}

func (m *usersDaoMock)GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return getUserFunction(userId)
}

func TestGetUserNotFoundInDatabase(t *testing.T) {
	getUserFunction = func(userId int64) (*domain.User, *utils.ApplicationError) {
		return nil, &utils.ApplicationError{
			StatusCode: http.StatusNotFound,
			Message: fmt.Sprintf("User with id %v was not found", userId),
		}
	}
	user, err := UserService.GetUser(0)
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, fmt.Sprintf("User with id %v was not found", 0), err.Message)
}

func TestGetUserNoError(t *testing.T) {
	getUserFunction = func(userId int64) (*domain.User, *utils.ApplicationError) {
		return &domain.User{
			Id: 123,
		}, nil
	}

	user, err := UserService.GetUser(123)
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, user.Id, 123)
}
