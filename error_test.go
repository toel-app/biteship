package biteship

import (
	"encoding/json"
	"errors"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"testing"
)

func TestErrorGo(t *testing.T) {
	err := errors.New("myerror")
	errorGo := ErrorGo(err)

	assert.Equal(t, errorGo.Status, http.StatusInternalServerError)
	assert.Equal(t, errorGo.ErrorCode, "Error Go")
	assert.Equal(t, errorGo.Message, err.Error())
}

func TestErrorGo_Panics(t *testing.T) {
	assert.Panics(t, func() {
		ErrorGo(nil)
	})
}

func TestErrorRequestParam(t *testing.T) {
	err := errors.New("testerr")
	errRequestParam := ErrorRequestParam(err)

	assert.Equal(t, errRequestParam.Status, http.StatusBadRequest)
	assert.Equal(t, errRequestParam.ErrorCode, "Bad Request")
	assert.Equal(t, errRequestParam.Message, err.Error())
}

func TestErrorRequestParams_Panics(t *testing.T) {
	assert.Panics(t, func() {
		ErrorRequestParam(nil)
	})
}

func TestErrorHttp(t *testing.T) {
	mockError := Error{
		Status:    0,
		ErrorCode: "",
		Message:   "",
		RawError:  "",
		Code:      0,
	}

	validJson, err := json.Marshal(mockError)
	if err != nil {
		log.Println("Failure :")
		log.Fatalln(err)
	} else {
		err := ErrorHttp(http.StatusBadRequest, validJson)
		assert.Equal(t, err.Status, http.StatusBadRequest)
		assert.Equal(t, err.Code, mockError.Code)
		assert.Equal(t, err.Message, mockError.Message)
		assert.Equal(t, err.RawError, mockError.RawError)
		assert.Equal(t, err.ErrorCode, mockError.ErrorCode)
	}
}

func TestErrorHttp_InvalidJson(t *testing.T) {
	invalidJson := `{"some", "data", "invalid_property"}`

	err := ErrorHttp(http.StatusBadRequest, []byte(invalidJson))

	assert.Equal(t, err.Status, http.StatusInternalServerError)
	assert.Equal(t, err.ErrorCode, "Error Go")
	assert.NotEqual(t, err.Message, "")
}
