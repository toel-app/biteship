package biteship

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io"
	"testing"
)

type trackingTestHttpMock struct {
	mock.Mock
}

func (m *trackingTestHttpMock) Call(method string, url string, secretKey string, body io.Reader, result interface{}) *Error {
	args := m.Called(method, url, secretKey, body, result)
	stub := args.Get(0)
	errStruct, isErrorStruct := stub.(*Error)

	if isErrorStruct {
		return errStruct
	}

	return nil
}

func TestTrackOrder(t *testing.T) {
	var (
		mockHttp    = new(trackingTestHttpMock)
		mockOrderId = "orderId"
		mockClient  = Client{
			SecretKey:   "",
			BiteshipUrl: "",
			HttpRequest: mockHttp,
		}
	)

	mockHttp.On("Call", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()

	response, err := mockClient.TrackOrder(mockOrderId)

	mockHttp.AssertExpectations(t)
	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.IsType(t, response, &ResponseTrackingOrder{})
}

func TestTrackOrder_Error(t *testing.T) {
	var (
		mockHttp    = new(trackingTestHttpMock)
		mockOrderId = "orderId"
		mockClient  = Client{
			SecretKey:   "",
			BiteshipUrl: "",
			HttpRequest: mockHttp,
		}
	)

	mockError := Error{}
	mockHttp.On("Call", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&mockError).Once()

	response, err := mockClient.TrackOrder(mockOrderId)

	mockHttp.AssertExpectations(t)
	assert.NotNil(t, err)
	assert.NotNil(t, response)
	assert.IsType(t, err, &Error{})
}

func TestTrackOrderByWaybill(t *testing.T) {
	var (
		mockHttp        = new(trackingTestHttpMock)
		mockWaybillId   = "waybill_id"
		mockCourierCode = CourierProviderJNE
		mockClient      = Client{
			SecretKey:   "",
			BiteshipUrl: "",
			HttpRequest: mockHttp,
		}
	)

	mockError := Error{}
	mockHttp.On("Call", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&mockError).Once()

	response, err := mockClient.TrackOrderByWaybill(mockWaybillId, string(mockCourierCode))

	mockHttp.AssertExpectations(t)
	assert.NotNil(t, err)
	assert.NotNil(t, response)
	assert.IsType(t, err, &Error{})
}

func TestTrackOrderByWaybill_Error(t *testing.T) {
	var (
		mockHttp        = new(trackingTestHttpMock)
		mockWaybillId   = "waybill_id"
		mockCourierCode = CourierProviderJNE
		mockClient      = Client{
			SecretKey:   "",
			BiteshipUrl: "",
			HttpRequest: mockHttp,
		}
	)

	mockError := Error{}
	mockHttp.On(
		"Call",
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(&mockError).Once()

	response, err := mockClient.TrackOrderByWaybill(mockWaybillId, string(mockCourierCode))

	mockHttp.AssertExpectations(t)
	assert.NotNil(t, err)
	assert.NotNil(t, response)
	assert.IsType(t, err, &Error{})
}
