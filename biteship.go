package biteship

import "github.com/go-playground/validator/v10"

var validate *validator.Validate

type Biteship interface {
	GetCouriers() (*ResponseListCourier, *Error)
	GetCourierRates(request *RequestCourierRates) (*ResponseListRatesCouriers, *Error)
	CreateOrder(request *CreateOrderRequestParam) (*ResponseCreateOrder, *Error)
	RetrieveOrder(orderId string) (*ResponseRetrieveOrder, *Error)
	UpdateOrder(orderId string, request interface{}) (*ResponseCreateOrder, *Error)
	ConfirmOrder(orderId string) (*ResponseCreateOrder, *Error)
	CancelOrder(orderId string, reason string) (*ResponseCancelOrder, *Error)
	TrackOrder(orderId string) (*ResponseTrackingOrder, *Error)
	TrackOrderByWaybill(waybillId string, courierCode string) (*ResponseTrackingOrder, *Error)
}

type Impl struct {
	Config      *ConfigOption
	HttpRequest *HttpRequestImpl
}

func New(key string, config ...ConfigOption) Biteship {
	defaultConfig := DefaultConfig(key)

	if len(config) > 0 {
		defaultConfig = &config[0]
		defaultConfig.SecretKey = key
	}

	return &Impl{
		Config: defaultConfig,
	}
}
