package github_provider

import (
	"davidsodergren/golang-microservices/src/api/clients/restclient"
	"davidsodergren/golang-microservices/src/api/domain/github"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	headerAuthorization = "Authorization"
	headerAuthorizationFormat = "token %v"

	urlCreateRepo = "https://api.github.com/user/repos"
)

func getAuthorizationHeader(accessToken string) string {
	return fmt.Sprintf(headerAuthorizationFormat, accessToken)
}

func CreateRepo(headers http.Header, request github.CreateRepoRequest) (*github.CreateRepoResponse, error) {

	response, err := restclient.Post(urlCreateRepo, request, headers)
	if err != nil {
		return nil, err
	}

	httpResponseInBytes, err := httpResponseToBytes(response)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if err := statusCodeOk(response, httpResponseInBytes); err != nil {
		return nil, err
	}

	var result github.CreateRepoResponse

	if err := json.Unmarshal(httpResponseInBytes, &result); err != nil {
		log.Println(fmt.Sprintf("error when trying to unmarshal create repo successful response: %s", err.Error()))
		return nil, &github.ErrorResponse{StatusCode: http.StatusInternalServerError, Message: "error when trying to unmarshal github create repo response "}
	}

	return &result, nil
}

func createHeader(accessToken string) http.Header {
	headers := http.Header{}
	headers.Set(headerAuthorization, getAuthorizationHeader(accessToken))
	return headers
}

func httpResponseToBytes(response *http.Response) ([]byte, error){
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(fmt.Sprintf("error when trying to unmarshal create repo successful response: %s", err.Error()))
		return nil, &github.ErrorResponse{StatusCode: http.StatusInternalServerError, Message: "Invalid response body"}
	}
	return bytes, nil
}

func statusCodeOk(response *http.Response, httpResponseInBytes []byte ) error {
	if response.StatusCode > 299 {
		var errResponse github.ErrorResponse
		if err := json.Unmarshal(httpResponseInBytes, &errResponse); err != nil {
			return &github.ErrorResponse{StatusCode: http.StatusInternalServerError, Message: "invalid json response body"}
		}
		errResponse.StatusCode = response.StatusCode
		return &errResponse
	}
}
