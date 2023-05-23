package biteship

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"
	"reflect"
)

func (client *Client) GetCourierRates(request *RequestCourierRates) (*ResponseListRatesCouriers, *Error) {
	var (
		resp        = &ResponseListRatesCouriers{}
		url         = fmt.Sprintf("%s/v1/rates/couriers", client.BiteshipUrl)
		jsonRequest = []byte("")
		errMarshal  error
	)

	if errValidate := validator.New().Struct(request); errValidate != nil {
		return resp, ErrorRequestParam(errValidate)
	}

	isParamsNil := reflect.ValueOf(request).Kind() == reflect.Ptr && reflect.ValueOf(request).IsNil()

	if !isParamsNil {
		jsonRequest, errMarshal = json.Marshal(request)
		if errMarshal != nil {
			log.Println(errMarshal)
			return nil, ErrorGo(errMarshal)
		}
	}

	errRequest := client.HttpRequest.Call(http.MethodPost, url, client.SecretKey, bytes.NewBuffer(jsonRequest), resp)
	if errRequest != nil {
		return resp, errRequest
	}

	return resp, nil
}
