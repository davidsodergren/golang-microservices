package controllers

import (
	"davidsodergren/golang-microservices/mvc/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Respond(c *gin.Context, body interface{})  {
	if c.GetHeader("Accept") == "application/xml" {
		c.XML(http.StatusOK, body)
	}
	c.JSON(http.StatusOK, body)
}

func RespondError(c *gin.Context, err *utils.ApplicationError)  {
	if c.GetHeader("Accept") == "application/xml" {
		c.XML(err.StatusCode, err)
	}
	c.JSON(err.StatusCode, err)
}
