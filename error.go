package biteship

import (
	"encoding/json"
	"log"
	"net/http"
)

type Error struct {
	Status    int    `json:"status,omitempty"`
	ErrorCode string `json:"error_code,omitempty"`
	Message   string `json:"message,omitempty"`
	RawError  string `json:"error,omitempty"`
	Code      int    `json:"code,omitempty"`
}

func ErrorGo(err error) *Error {
	return &Error{
		Status:    http.StatusInternalServerError,
		ErrorCode: "Error Go",
		Message:   err.Error(),
	}
}

func ErrorRequestParam(err error) *Error {
	return &Error{
		Status:    http.StatusBadRequest,
		ErrorCode: "Bad Request",
		Message:   err.Error(),
	}
}

func ErrorHttp(status int, respBody []byte) *Error {
	var httpError *Error
	if err := json.Unmarshal(respBody, &httpError); err != nil {
		log.Println(err)
		return ErrorGo(err)
	}
	httpError.Status = status

	return httpError
}
