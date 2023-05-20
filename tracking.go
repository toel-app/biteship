package biteship

import (
	"fmt"
	"net/http"
)

func (bite *Client) TrackOrder(orderId string) (*ResponseTrackingOrder, *Error) {
	resp := &ResponseTrackingOrder{}
	var url = fmt.Sprintf("%s/v1/trackings/%s", bite.BiteshipUrl, orderId)

	errRequest := bite.HttpRequest.Call(http.MethodGet, url, bite.SecretKey, nil, resp)
	if errRequest != nil {
		return resp, errRequest
	}

	return resp, nil
}

func (bite *Client) TrackOrderByWaybill(waybillId string, courierCode string) (*ResponseTrackingOrder, *Error) {
	resp := &ResponseTrackingOrder{}
	var url = fmt.Sprintf("%s/v1/trackings/%s/couriers/%s", bite.BiteshipUrl, waybillId, courierCode)

	errRequest := bite.HttpRequest.Call(http.MethodGet, url, bite.SecretKey, nil, resp)
	if errRequest != nil {
		return resp, errRequest
	}

	return resp, nil
}
