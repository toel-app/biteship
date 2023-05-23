package biteship

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io"
	"net/http"
	"net/url"
	"testing"
)

type mapsTestHttpMock struct {
	mock.Mock
}

func (m *mapsTestHttpMock) Call(method string, url string, secretKey string, body io.Reader, result interface{}) *Error {
	args := m.Called(method, url, secretKey, body, result)
	stub := args.Get(0)
	errStruct, ok := stub.(*Error)

	if ok {
		return errStruct
	}

	return nil
}

func TestRetrieveArea_Success(t *testing.T) {
	var (
		countries = "ID"
		input     = "Jakarta Selatan"
		mockHttp  = new(mapsTestHttpMock)
	)

	client := Client{
		BiteshipUrl: DefaultUrl,
		HttpRequest: mockHttp,
	}

	v := url.Values{}
	v.Set("countries", countries)
	v.Set("input", input)

	mockHttp.On("Call", http.MethodGet, fmt.Sprintf("%s/v1/maps/areas?%s", DefaultUrl, v.Encode()), client.SecretKey, nil, mock.Anything).Return(nil).Once()
	resp, err := client.RetrieveArea(countries, input)

	mockHttp.AssertExpectations(t)
	assert.Nil(t, err)
	assert.NotNil(t, resp)
}

func TestRetrieveArea_ErrHttp(t *testing.T) {
	var (
		countries = "ID"
		input     = "Jakarta Selatan"
		mockHttp  = new(mapsTestHttpMock)
	)

	client := Client{
		BiteshipUrl: DefaultUrl,
		HttpRequest: mockHttp,
	}

	v := url.Values{}
	v.Set("countries", countries)
	v.Set("input", input)

	mockHttp.On("Call", http.MethodGet, fmt.Sprintf("%s/v1/maps/areas?%s", DefaultUrl, v.Encode()), client.SecretKey, nil, mock.Anything).Return(&Error{}).Once()
	resp, err := client.RetrieveArea(countries, input)

	mockHttp.AssertExpectations(t)
	assert.NotNil(t, err)
	assert.Nil(t, resp)
}

func TestRetrieveAreaByID_Success(t *testing.T) {
	var (
		areaId   = "IDNP6IDNC148IDND842IDZ12520"
		mockHttp = new(mapsTestHttpMock)
	)

	client := Client{
		BiteshipUrl: DefaultUrl,
		HttpRequest: mockHttp,
	}

	mockHttp.On("Call", http.MethodGet, fmt.Sprintf("%s/v1/maps/areas/%s", DefaultUrl, areaId), client.SecretKey, nil, mock.Anything).Return(nil).Once()
	resp, err := client.RetrieveAreaByID(areaId)

	mockHttp.AssertExpectations(t)
	assert.Nil(t, err)
	assert.NotNil(t, resp)
}

func TestRetrieveAreaByID_ErrHttp(t *testing.T) {
	var (
		areaId   = "IDNP6IDNC148IDND842IDZ12520"
		mockHttp = new(mapsTestHttpMock)
	)

	client := Client{
		BiteshipUrl: DefaultUrl,
		HttpRequest: mockHttp,
	}

	mockHttp.On("Call", http.MethodGet, fmt.Sprintf("%s/v1/maps/areas/%s", DefaultUrl, areaId), client.SecretKey, nil, mock.Anything).Return(&Error{}).Once()
	resp, err := client.RetrieveAreaByID(areaId)

	mockHttp.AssertExpectations(t)
	assert.NotNil(t, err)
	assert.Nil(t, resp)
}
