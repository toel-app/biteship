package biteship

type Metadata interface{}

type Coordinate struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type ProductItem struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Value       uint   `json:"value"`
	Quantity    uint   `json:"quantity" validate:"gte=1"`
	Height      uint   `json:"height" validate:"gte=1"`
	Width       uint   `json:"width" validate:"gte=1"`
	Length      uint   `json:"length" validate:"gte=1"`
	Weight      uint   `json:"weight"`
}

type CreateOrderRequestParam struct {
	ShipperContactName            string          `json:"shipper_contact_name"`
	ShipperContactPhone           string          `json:"shipper_contact_phone"`
	ShipperContactEmail           string          `json:"shipper_contact_email"`
	ShipperOrganization           string          `json:"shipper_organization"`
	OriginContactName             string          `json:"origin_contact_name"`
	OriginContactPhone            string          `json:"origin_contact_phone"`
	OriginAddress                 string          `json:"origin_address"`
	OriginNote                    string          `json:"origin_note"`
	OriginPostalCode              uint32          `json:"origin_postal_code"`
	OriginCoordinate              Coordinate      `json:"origin_coordinate"`
	DestinationContactName        string          `json:"destination_contact_name"`
	DestinationContactPhone       string          `json:"destination_contact_phone"`
	DestinationContactEmail       string          `json:"destination_contact_email"`
	DestinationAddress            string          `json:"destination_address"`
	DestinationPostalCode         uint32          `json:"destination_postal_code"`
	DestinationNote               string          `json:"destination_note"`
	DestinationCoordinate         Coordinate      `json:"destination_coordinate"`
	DestinationCashOnDelivery     *uint           `json:"destination_cash_on_delivery"` // Optional
	DestinationCashOnDeliveryType *string         `json:"destination_cash_on_delivery_type"`
	CourierCompany                CourierProvider `json:"courier_company" binding:"required"`
	CourierType                   string          `json:"courier_type" binding:"required"`
	CourierInsurance              uint            `json:"courier_insurance"`
	DeliveryType                  DeliveryType    `json:"delivery_type" binding:"required"` // "later" or "now"
	DeliveryDate                  string          `json:"delivery_date"`                    // yyyy-mm-dd
	DeliveryTime                  string          `json:"delivery_time"`                    // hh:mm
	PaymentType                   string          `json:"payment_type" binding:"required"`  // Set to be 'online'
	OrderNote                     string          `json:"order_note"`
	Metadata                      Metadata        `json:"metadata"` // Optional

	Items []ProductItem `json:"items"`
}

type ItemCourierRate struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Value       uint   `json:"value"`
	Length      uint   `json:"length"`
	Width       uint   `json:"width"`
	Height      uint   `json:"height"`
	Weight      uint   `json:"weight"`
	Quantity    uint   `json:"quantity"`
}

type RequestCourierRates struct {
	OriginLatitude        *float64          `json:"origin_latitude,omitempty"`
	OriginLongitude       *float64          `json:"origin_longitude,omitempty"`
	DestinationLatitude   *float64          `json:"destination_latitude,omitempty"`
	DestinationLongitude  *float64          `json:"destination_longitude,omitempty"`
	OriginPostalCode      uint              `json:"origin_postal_code,omitempty"`
	DestinationPostalCode uint              `json:"destination_postal_code,omitempty"`
	Couriers              string            `json:"couriers" validate:"required"` // "grab, gojek, jne, tiki, jet, jnt, sicepat, wahana, pos, lion, ninja, anteraja, rpx, paxel, mrspeedy, lalamove, deliveree, sap"
	Items                 []ItemCourierRate `json:"items"`
}

type ReasonRequest struct {
	CancellationReason string `json:"cancellation_reason"`
}
