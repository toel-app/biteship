package biteship

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

type IHttpRequest interface {
	Call(method string, url string, secretKey string, body io.Reader, result interface{}) *Error
}

type HttpRequest struct{}

func NewHttp() IHttpRequest {
	return &HttpRequest{}
}

func (client *HttpRequest) Call(method string, url string, secretKey string, body io.Reader, result interface{}) *Error {
	if secretKey == "" {
		return &Error{
			Status:  http.StatusUnauthorized,
			Message: "missing/invalid secret key",
		}
	}

	req, errNewReq := http.NewRequest(method, url, body)
	if errNewReq != nil {
		return &Error{
			Status:   http.StatusInternalServerError,
			Message:  "Cannot create request",
			RawError: errNewReq.Error(),
		}
	}

	req.Header.Add("Authorization", secretKey)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	return client.doRequest(req, result)
}

func (client *HttpRequest) doRequest(req *http.Request, result interface{}) *Error {
	httpClient := &http.Client{}

	response, errRequest := httpClient.Do(req)
	if errRequest != nil {
		return ErrorGo(errRequest)
	}
	defer response.Body.Close()

	respBody, errRead := ioutil.ReadAll(response.Body)
	if errRead != nil {
		return ErrorGo(errRead)
	}

	if response.StatusCode < http.StatusOK || response.StatusCode >= http.StatusMultipleChoices {
		return ErrorHttp(response.StatusCode, respBody)
	}

	errUnmarshall := json.Unmarshal(respBody, &result)
	if errUnmarshall != nil {
		return ErrorGo(errUnmarshall)
	}

	return nil
}
