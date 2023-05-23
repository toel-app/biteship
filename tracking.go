package biteship

import (
	"fmt"
	"net/http"
)

func (client *Client) TrackOrder(orderId string) (*ResponseTrackingOrder, *Error) {
	resp := &ResponseTrackingOrder{}
	var url = fmt.Sprintf("%s/v1/trackings/%s", client.BiteshipUrl, orderId)

	errRequest := client.HttpRequest.Call(http.MethodGet, url, client.SecretKey, nil, resp)
	if errRequest != nil {
		return resp, errRequest
	}

	return resp, nil
}

func (client *Client) TrackOrderByWaybill(waybillId string, courierCode string) (*ResponseTrackingOrder, *Error) {
	resp := &ResponseTrackingOrder{}
	var url = fmt.Sprintf("%s/v1/trackings/%s/couriers/%s", client.BiteshipUrl, waybillId, courierCode)

	errRequest := client.HttpRequest.Call(http.MethodGet, url, client.SecretKey, nil, resp)
	if errRequest != nil {
		return resp, errRequest
	}

	return resp, nil
}
