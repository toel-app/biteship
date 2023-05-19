package biteship

import "github.com/go-playground/validator/v10"

var validate *validator.Validate

type Biteship interface {
	GetCouriers() (*ResponseListCourier, *Error)
	// GetCourierRates List of available courier
	GetCourierRates(request *RequestCourierRates) (*ResponseListRatesCouriers, *Error)

	// CreateOrder Order objects are created to handle sellers’ shipment.
	// You can create, retrieve, and update individual orders. Orders are identified by a unique, random ID
	CreateOrder(request *CreateOrderRequestParam) (*ResponseCreateOrder, *Error)

	// RetrieveOrder Check your order history or tracking by orderId.
	// You can get the Order ID from the Order API request
	RetrieveOrder(orderId string) (*ResponseRetrieveOrder, *Error)

	// UpdateOrder When the order status has not been updated to confirmed, sellers can edit their order based on their needs.
	// For Example Let’s pretend if you want to change your origin address.
	// You can just simply send a JSON body with only the origin_address field. It will automatically change your current order details.
	UpdateOrder(orderId string, request interface{}) (*ResponseCreateOrder, *Error)
	ConfirmOrder(orderId string) (*ResponseCreateOrder, *Error)
	CancelOrder(orderId string, reason string) (*ResponseCancelOrder, *Error)

	// TrackOrder This endpoint can only be used when you order via our order API.
	// Biteship will generate tracking_id separately if you create an Order through Biteship API.
	TrackOrder(orderId string) (*ResponseTrackingOrder, *Error)

	// TrackOrderByWaybill (Public Tracking) This endpoint can be used to track any other waybill from any other source.
	// It requires the courier code which you can find the in Courier API
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
