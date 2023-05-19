package biteship

import (
	"fmt"
	"net/http"
)

func (bite *Impl) TrackOrder(orderId string) (*ResponseTrackingOrder, *Error) {
	resp := &ResponseTrackingOrder{}
	var url = fmt.Sprintf("%s/v1/trackings/%s", bite.Config.BiteshipUrl, orderId)

	errRequest := bite.HttpRequest.Call(http.MethodGet, url, bite.Config.SecretKey, nil, resp)
	if errRequest != nil {
		return resp, errRequest
	}

	return resp, nil
}

func (bite *Impl) TrackOrderByWaybill(waybillId string, courierCode string) (*ResponseTrackingOrder, *Error) {
	resp := &ResponseTrackingOrder{}
	var url = fmt.Sprintf("%s/v1/trackings/%s/couriers/%s", bite.Config.BiteshipUrl, waybillId, courierCode)

	errRequest := bite.HttpRequest.Call(http.MethodGet, url, bite.Config.SecretKey, nil, resp)
	if errRequest != nil {
		return resp, errRequest
	}

	return resp, nil
}
