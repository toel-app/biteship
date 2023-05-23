package biteship

import (
	"fmt"
	"net/http"
)

const (
	CourierProviderGrab      CourierProvider = "grab"
	CourierProviderGojek     CourierProvider = "gojek"
	CourierProviderJNE       CourierProvider = "jne"
	CourierProviderTIKI      CourierProvider = "tiki"
	CourierProviderJET       CourierProvider = "jet"
	CourierProviderJNT       CourierProvider = "jnt"
	CourierProviderSicepat   CourierProvider = "sicepat"
	CourierProviderWahana    CourierProvider = "wahana"
	CourierProviderPos       CourierProvider = "pos"
	CourierProviderLion      CourierProvider = "lion"
	CourierProviderNinja     CourierProvider = "ninja"
	CourierProviderAnteraja  CourierProvider = "anteraja"
	CourierProviderRPX       CourierProvider = "rpx"
	CourierProviderPaxel     CourierProvider = "paxel"
	CourierProviderMrSpeedy  CourierProvider = "mrspeedy"
	CourierProviderLalamove  CourierProvider = "lalamove"
	CourierProviderDeliveree CourierProvider = "deliveree"
	CourierProviderSAP       CourierProvider = "sap"
)

type CourierProvider string

func (client *Client) GetCouriers() (*ResponseListCourier, *Error) {
	resp := &ResponseListCourier{}
	var url = fmt.Sprintf("%s/v1/couriers", client.
		BiteshipUrl)

	errRequest := client.HttpRequest.Call(http.MethodGet, url, client.
		SecretKey, nil, resp)
	if errRequest != nil {
		return resp, errRequest
	}

	return resp, nil
}
