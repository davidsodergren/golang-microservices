package controllers

import (
	"davidsodergren/golang-microservices/mvc/services"
	"davidsodergren/golang-microservices/mvc/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetUser(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		userErr := &utils.ApplicationError{
			Message: "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code: "bad_requst",
		}
		RespondError(c, userErr)
		return
	}
	user, apiErr := services.UserService.GetUser(userId)
	if apiErr != nil {
		RespondError(c, apiErr)
		return
	}
	Respond(c, user)
}
