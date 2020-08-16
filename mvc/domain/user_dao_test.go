package domain

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestGetUserNoUserFound(t *testing.T) {
	const userId int64 = 0
	user, err := UserDao.GetUser(userId)

	assert.Nil(t, user, "We were not expecting a user with id 0")
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, fmt.Sprintf("User with id %v was not found", userId), err.Message)
}

func TestGetUserNoError(t *testing.T)  {
	const userId uint64 = 123
	user, err := UserDao.GetUser(123)

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, userId, user.Id)

}