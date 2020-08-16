package restclient

import (
	bytes2 "bytes"
	"davidsodergren/golang-microservices/src/api/domain/github"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func Post(url string, body interface{}, headers http.Header) (*http.Response, error) {
	jsonBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest(http.MethodPost, url, bytes2.NewReader(jsonBytes))
	if err != nil {
		log.Println(fmt.Sprintf("Error when trying to create a new request: %s", err.Error()))
		return nil, &github.ErrorResponse{ StatusCode: http.StatusInternalServerError, Message: err.Error(),
		}
	}
	request.Header = headers
	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Println(fmt.Sprintf("Error when trying to create a new repo in github: %s", err.Error()))
		return nil, &github.ErrorResponse{ StatusCode: http.StatusInternalServerError, Message: err.Error(),
		}
	}
	return response, nil
}
