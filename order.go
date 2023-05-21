package biteship

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"io"
	"log"
	"net/http"
	"reflect"
)

type DeliveryType string

const (
	DeliveryTypeNow   DeliveryType = "now"
	DeliveryTypeLater DeliveryType = "later"
)

func (bite *Client) CreateOrder(request *CreateOrderRequestParam) (*ResponseCreateOrder, *Error) {
	var (
		resp        = &ResponseCreateOrder{}
		url         = fmt.Sprintf("%s/v1/orders", bite.BiteshipUrl)
		jsonRequest = []byte("")
		errMarshal  error
	)

	errValidate := validator.New().Struct(request)
	if errValidate != nil {
		return resp, ErrorRequestParam(errValidate)
	}

	isNilRequest := reflect.ValueOf(request).Kind() == reflect.Ptr && reflect.ValueOf(request).IsNil()

	if !isNilRequest {
		jsonRequest, errMarshal = json.Marshal(request)
		if errMarshal != nil {
			return resp, ErrorGo(errMarshal)
		}
	}

	errRequest := bite.HttpRequest.Call(http.MethodPost, url, bite.SecretKey, bytes.NewBuffer(jsonRequest), resp)
	return resp, errRequest
}

func (bite *Client) RetrieveOrder(orderId string) (*ResponseRetrieveOrder, *Error) {
	var (
		resp       = &ResponseRetrieveOrder{}
		url        = fmt.Sprintf("%s/v1/orders/%s", bite.BiteshipUrl, orderId)
		errRequest = bite.HttpRequest.Call(http.MethodGet, url, bite.SecretKey, nil, resp)
	)

	return resp, errRequest
}

func (bite *Client) UpdateOrder(orderId string, request interface{}) (*ResponseCreateOrder, *Error) {
	var (
		resp        = &ResponseCreateOrder{}
		url         = fmt.Sprintf("%s/v1/orders/%s", bite.BiteshipUrl, orderId)
		jsonRequest = []byte("")
		errMarshal  error
	)

	isNilPtr := reflect.ValueOf(request).Kind() == reflect.Ptr && reflect.ValueOf(request).IsNil()

	if !isNilPtr {
		if jsonRequest, errMarshal = json.Marshal(request); errMarshal != nil {
			log.Println(errMarshal)
			return resp, ErrorGo(errMarshal)
		}
	}

	return resp, bite.HttpRequest.Call(http.MethodPost, url, bite.SecretKey, bytes.NewBuffer(jsonRequest), resp)
}

func (bite *Client) ConfirmOrder(orderId string) (*ResponseCreateOrder, *Error) {
	var (
		resp = &ResponseCreateOrder{}
		url  = fmt.Sprintf("%s/v1/orders/%s/confirm", bite.
			BiteshipUrl, orderId)
	)

	errRequest := bite.HttpRequest.Call(http.MethodPost, url, bite.SecretKey, nil, resp)

	if errRequest != nil {
		return resp, errRequest
	}

	return resp, nil
}

func (bite *Client) CancelOrder(orderId string, reason string) (*ResponseCancelOrder, *Error) {
	var (
		body        io.Reader
		resp        = &ResponseCancelOrder{}
		url         = fmt.Sprintf("%s/v1/orders/%s", bite.BiteshipUrl, orderId)
		errMarshal  error
		jsonRequest = []byte("")
	)

	isParamsNil := reason == "" || (reflect.ValueOf(reason).Kind() == reflect.Ptr && reflect.ValueOf(reason).IsNil())

	if !isParamsNil {
		cancelReason := ReasonRequest{reason}
		jsonRequest, errMarshal = json.Marshal(cancelReason)
		if errMarshal != nil {
			log.Println(errMarshal)
			return resp, ErrorGo(errMarshal)
		}
		body = bytes.NewBuffer(jsonRequest)
	} else {
		body = nil
	}

	errRequest := bite.HttpRequest.Call(http.MethodDelete, url, bite.
		SecretKey, body, resp)
	if errRequest != nil {
		return resp, errRequest
	}

	return resp, nil
}
