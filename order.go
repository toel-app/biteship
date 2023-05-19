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

func (bite *Impl) CreateOrder(request *CreateOrderRequestParam) (*ResponseCreateOrder, *Error) {
	resp := &ResponseCreateOrder{}
	var url = fmt.Sprintf("%s/v1/orders", bite.Config.BiteshipUrl)
	var errMarshal error
	jsonRequest := []byte("")

	validate = validator.New()
	errValidate := validate.Struct(request)
	if errValidate != nil {
		return resp, ErrorRequestParam(errValidate)
	}

	isParamsNil := reflect.ValueOf(request).Kind() == reflect.Ptr && reflect.ValueOf(request).IsNil()

	if !isParamsNil {
		jsonRequest, errMarshal = json.Marshal(request)
		if errMarshal != nil {
			log.Println(errMarshal)
			return resp, ErrorGo(errMarshal)
		}
	}

	errRequest := bite.HttpRequest.Call(http.MethodPost, url, bite.Config.SecretKey, bytes.NewBuffer(jsonRequest), resp)
	if errRequest != nil {
		return resp, errRequest
	}

	return resp, nil
}

func (bite *Impl) RetrieveOrder(orderId string) (*ResponseRetrieveOrder, *Error) {

	resp := &ResponseRetrieveOrder{}
	var url = fmt.Sprintf("%s/v1/orders/%s", bite.Config.BiteshipUrl, orderId)

	errRequest := bite.HttpRequest.Call(http.MethodGet, url, bite.Config.SecretKey, nil, resp)
	if errRequest != nil {
		return resp, errRequest
	}

	return resp, nil
}

func (bite *Impl) UpdateOrder(orderId string, request interface{}) (*ResponseCreateOrder, *Error) {

	resp := &ResponseCreateOrder{}
	var url = fmt.Sprintf("%s/v1/orders/%s", bite.Config.BiteshipUrl, orderId)
	var errMarshal error
	jsonRequest := []byte("")

	isParamsNil := reflect.ValueOf(request).Kind() == reflect.Ptr && reflect.ValueOf(request).IsNil()

	if !isParamsNil {
		jsonRequest, errMarshal = json.Marshal(request)
		if errMarshal != nil {
			log.Println(errMarshal)
			return resp, ErrorGo(errMarshal)
		}
	}

	errRequest := bite.HttpRequest.Call(http.MethodPost, url, bite.Config.SecretKey, bytes.NewBuffer(jsonRequest), resp)
	if errRequest != nil {
		return resp, errRequest
	}
	log.Println(bytes.NewBuffer(jsonRequest))

	return resp, nil
}

func (bite *Impl) ConfirmOrder(orderId string) (*ResponseCreateOrder, *Error) {

	resp := &ResponseCreateOrder{}
	var url = fmt.Sprintf("%s/v1/orders/%s/confirm", bite.Config.BiteshipUrl, orderId)

	errRequest := bite.HttpRequest.Call(http.MethodPost, url, bite.Config.SecretKey, nil, resp)
	if errRequest != nil {
		return resp, errRequest
	}

	return resp, nil
}

func (bite *Impl) CancelOrder(orderId string, reason string) (*ResponseCancelOrder, *Error) {
	var body io.Reader
	resp := &ResponseCancelOrder{}
	var url = fmt.Sprintf("%s/v1/orders/%s", bite.Config.BiteshipUrl, orderId)
	var errMarshal error
	jsonRequest := []byte("")

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

	errRequest := bite.HttpRequest.Call(http.MethodDelete, url, bite.Config.SecretKey, body, resp)
	if errRequest != nil {
		return resp, errRequest
	}

	return resp, nil
}
