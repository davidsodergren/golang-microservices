package controllers

import (
	"davidsodergren/golang-microservices/services"
	"davidsodergren/golang-microservices/utils"
	"encoding/json"
	"net/http"
	"strconv"
)

func GetUser(resp http.ResponseWriter, req *http.Request)  {
	userId, err := strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64)
	if err != nil {
		userErr := &utils.ApplicationError{
			Message: "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code: "bad_requst",
		}
		jsonValue, _ := json.Marshal(userErr)
		resp.WriteHeader(userErr.StatusCode)
		resp.Write(jsonValue)
		return
	}
	user, apiErr := services.UserService.GetUser(userId)
	if apiErr != nil {
		jsonValue, _ := json.Marshal(apiErr)
		resp.WriteHeader(http.StatusNotFound)
		resp.Write(jsonValue)
		return
	}
	jsonValue, _ := json.Marshal(user)
	resp.Write(jsonValue)
}
