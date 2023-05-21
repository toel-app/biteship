package biteship

import (
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io"
	"net/http"
	"testing"
)

type orderTestHttpMock struct {
	mock.Mock
}

func (m *orderTestHttpMock) Call(method string, url string, secretKey string, body io.Reader, result interface{}) *Error {
	args := m.Called(method, url, secretKey, body, result)
	stub := args.Get(0)
	errStruct, isErrorStruct := stub.(*Error)

	if isErrorStruct {
		return errStruct
	}

	return nil
}

func TestCreateOrder(t *testing.T) {
	var (
		mockHttp = new(orderTestHttpMock)
		client   = Client{
			SecretKey:   "",
			BiteshipUrl: "",
			HttpRequest: mockHttp,
		}
		createOrderRequestPayload = new(CreateOrderRequestParam)
	)

	mockHttp.On("Call", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()

	response, err := client.CreateOrder(createOrderRequestPayload)

	mockHttp.AssertExpectations(t)
	assert.Nil(t, err)
	assert.NotNil(t, response)
}

func TestCreateOrder_ErrValidate(t *testing.T) {
	var (
		mockHttp = new(orderTestHttpMock)
		client   = Client{
			SecretKey:   "",
			BiteshipUrl: "",
			HttpRequest: mockHttp,
		}
		nilRequestParam *CreateOrderRequestParam
	)

	response, err := client.CreateOrder(nil)

	mockHttp.AssertNotCalled(t, "Call")
	assert.NotNil(t, err)
	assert.NotNil(t, response)
	assert.IsType(t, err, &Error{})
	assert.Equal(t, err.Status, http.StatusBadRequest)
	assert.Equal(t, err.ErrorCode, "Bad Request")
	assert.Equal(t, err.Message, validator.New().Struct(nilRequestParam).Error())
}

func TestRetrieveOrder_Success(t *testing.T) {
	var (
		mockHttp = new(orderTestHttpMock)
		client   = Client{
			SecretKey:   "",
			BiteshipUrl: "",
			HttpRequest: mockHttp,
		}
	)

	mockHttp.On("Call", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()

	resp, err := client.RetrieveOrder("orderId")

	assert.Nil(t, err)
	assert.NotNil(t, resp)
}

func TestRetrieveOrder_ErrHttpCall(t *testing.T) {
	var (
		mockHttp = new(orderTestHttpMock)
		client   = Client{
			SecretKey:   "",
			BiteshipUrl: "",
			HttpRequest: mockHttp,
		}
		mockErr = Error{
			Status: ErrOrderOrderAlreadyInNewStatus,
		}
	)

	mockHttp.On("Call", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&mockErr).Once()

	resp, err := client.RetrieveOrder("orderId")

	assert.NotNil(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, err.Status, ErrOrderOrderAlreadyInNewStatus)
}

func TestConfirmOrder_Success(t *testing.T) {
	var (
		mockHttp = new(orderTestHttpMock)
		client   = Client{
			SecretKey:   "",
			BiteshipUrl: "",
			HttpRequest: mockHttp,
		}
	)

	mockHttp.On("Call", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()

	resp, err := client.ConfirmOrder("orderId")

	assert.NotNil(t, resp)
	assert.Nil(t, err)
}

func TestConfirmOrder_Error(t *testing.T) {
	var (
		mockHttp = new(orderTestHttpMock)
		client   = Client{
			SecretKey:   "",
			BiteshipUrl: "",
			HttpRequest: mockHttp,
		}
		mockErr = Error{
			Status: ErrOrderOrderAlreadyInNewStatus,
		}
	)

	mockHttp.On("Call", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&mockErr).Once()

	resp, err := client.ConfirmOrder("orderId")

	assert.NotNil(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, err.Status, ErrOrderOrderAlreadyInNewStatus)
}

func TestCancelOrder_EmptyBody(t *testing.T) {
	var (
		mockHttp = new(orderTestHttpMock)
		client   = Client{
			HttpRequest: mockHttp,
		}
	)

	mockHttp.On("Call", mock.Anything, mock.Anything, mock.Anything, nil, mock.Anything).Return(nil).Once()
	resp, err := client.CancelOrder("orderId", "")

	mockHttp.AssertExpectations(t)
	assert.Nil(t, err)
	assert.NotNil(t, resp)
}

func TestCancelOrder_ErrHttpCall(t *testing.T) {
	var (
		mockHttp = new(orderTestHttpMock)
		client   = Client{
			HttpRequest: mockHttp,
		}
		mockErr = Error{
			Status: ErrOrderCannotEditConfirmedOrder,
		}
	)

	mockHttp.On("Call", mock.Anything, mock.Anything, mock.Anything, nil, mock.Anything).Return(&mockErr).Once()
	resp, err := client.CancelOrder("orderId", "")

	mockHttp.AssertExpectations(t)
	assert.NotNil(t, err)
	assert.NotNil(t, resp)
}
