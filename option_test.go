package biteship

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"io"
	"testing"
)

func TestWithSecret_Apply(t *testing.T) {
	someSecret := "alien exists"
	client := Client{
		SecretKey:   "",
		BiteshipUrl: "",
		HttpRequest: nil,
	}

	WithSecret(someSecret).Apply(&client)

	assert.Equal(t, client.SecretKey, someSecret)
	assert.Empty(t, client.BiteshipUrl)
	assert.Nil(t, client.HttpRequest)
}

func TestWithUrl_Apply(t *testing.T) {
	someUrl := "https://google.com"
	client := Client{
		SecretKey:   "",
		BiteshipUrl: "",
		HttpRequest: nil,
	}

	WithUrl(someUrl).Apply(&client)

	assert.Empty(t, client.SecretKey)
	assert.Equal(t, client.BiteshipUrl, someUrl)
	assert.Nil(t, client.HttpRequest)
}

type httpMock struct{}

func (m httpMock) Call(method string, url string, secretKey string, body io.Reader, result interface{}) *Error {
	return ErrorGo(errors.New("someerror"))
}

func TestWithHttpRequest_Apply(t *testing.T) {
	mockHttp := httpMock{}
	client := Client{
		SecretKey:   "",
		BiteshipUrl: "",
		HttpRequest: nil,
	}

	WithHttpRequest(mockHttp).Apply(&client)

	assert.NotNil(t, client.HttpRequest)
}
