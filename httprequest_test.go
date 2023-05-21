package biteship

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"testing"
)

func TestCall_EmptySecret(t *testing.T) {
	client := NewHttp()

	err := client.Call(http.MethodPost, "https://someUrl.com", "", nil, nil)

	assert.NotNil(t, err)
	assert.Equal(t, err.Status, http.StatusUnauthorized)
	assert.Equal(t, err.Message, "missing/invalid secret key")
}

func TestCall_InvalidURL(t *testing.T) {
	var (
		body   io.Reader = nil
		secret           = "somesecret"
		url              = "wttps://invalid-url.com.com"
	)

	client := NewHttp()

	err := client.Call(http.MethodPost, url, secret, body, nil)

	assert.NotNil(t, err)
	assert.Equal(t, err.Status, http.StatusInternalServerError)
	assert.Equal(t, err.Error(), fmt.Sprintf(`Post "%s": unsupported protocol scheme "wttps"`, url))
}

func TestCall_UnknownUrl(t *testing.T) {
	var (
		body   io.Reader = nil
		secret           = "somesecret"
		url              = "%escape"
	)

	client := NewHttp()

	err := client.Call(http.MethodPost, url, secret, body, nil)

	assert.NotNil(t, err)
	assert.Equal(t, err.Status, http.StatusInternalServerError)
	assert.Equal(t, err.Message, "Cannot create request")
	assert.Equal(t, err.RawError, `parse "%escape": invalid URL escape "%es"`)
}
