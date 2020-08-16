package app

import (
	"davidsodergren/golang-microservices/mvc/controllers"
)

func mapUrls()  {
	router.GET("/users/:user_id", controllers.GetUser)
}
