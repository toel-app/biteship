package biteship

import "net/http"

type ResponseWithMap map[string]interface{}

type Shipper struct {
	Name         string `json:"name"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	Organization string `json:"organization"`
}

type Origin struct {
	Coordinate   Coordinate `json:"coordinate"`
	PostalCode   uint       `json:"postal_code"`
	ContactName  string     `json:"contact_name"`
	ContactPhone string     `json:"contact_phone"`
	Address      string     `json:"address"`
	Note         string     `json:"note"`
}

type Destination struct {
	Coordinate   Coordinate `json:"coordinate"`
	PostalCode   uint       `json:"postal_code"`
	ContactName  string     `json:"contact_name"`
	ContactPhone string     `json:"contact_phone"`
	ContactEmail string     `json:"contact_email"`
	Address      string     `json:"address,omitempty"`
	Location     string     `json:"location,omitempty"`
	Note         string     `json:"note"`
}

type Courier struct {
	TrackingId *string `json:"tracking_id"`
	WaybillId  *string `json:"waybill_id"`
	Company    string  `json:"company"`
	Name       *string `json:"name"`
	Phone      *string `json:"phone"`
	Type       string  `json:"type"`
	Link       *string `json:"link"`
	Status     *string `json:"status,omitempty"`
	History    []struct {
		ServiceType string `json:"service_type"`
		Status      string `json:"status"`
		Note        string `json:"note"`
		UpdatedAt   string `json:"updated_at"`
	} `json:"history,omitempty"`
}

type ResponseCreateOrder struct {
	Success     bool          `json:"success"`
	Message     string        `json:"message"`
	Object      string        `json:"object"`
	Id          string        `json:"id"`
	Shipper     Shipper       `json:"shipper"`
	Origin      Origin        `json:"origin"`
	Destination Destination   `json:"destination"`
	Courier     Courier       `json:"courier"`
	Items       []ProductItem `json:"items"`
	Price       uint          `json:"price"`
	Note        string        `json:"note"`
	Status      string        `json:"status"`
}

type Delivery struct {
	Type     string  `json:"type"`
	Datetime string  `json:"datetime"`
	Note     *string `json:"note"`
}

// ResponseRetrieveOrder RESPONSE RETRIEVE ORDER
type ResponseRetrieveOrder struct {
	Success     bool          `json:"success"`
	Message     string        `json:"message"`
	Object      string        `json:"object"`
	Id          string        `json:"id"`
	Shipper     Shipper       `json:"shipper"`
	Origin      Origin        `json:"origin"`
	Delivery    Delivery      `json:"delivery"`
	Destination Destination   `json:"destination"`
	Courier     Courier       `json:"courier"`
	Items       []ProductItem `json:"items"`
	Price       uint          `json:"price"`
	Note        string        `json:"note"`
	Status      string        `json:"status"`
}

type ResponseCancelOrder struct {
	Success            bool   `json:"success"`
	Message            string `json:"message"`
	Object             string `json:"object"`
	Id                 string `json:"id"`
	Status             string `json:"status"`
	CancellationReason string `json:"cancellation_reason"`
}

// PricingRate RATES COURIER
type PricingRate struct {
	Company               string `json:"company"`
	CourierName           string `json:"courier_name"`
	CourierCode           string `json:"courier_code"`
	CourierServiceName    string `json:"courier_service_name"`
	CourierServiceCode    string `json:"courier_service_code"`
	Type                  string `json:"type"`
	Description           string `json:"description"`
	Duration              string `json:"duration"`
	ShipmentDurationRange string `json:"shipment_duration_range"`
	ShipmentDurationUnit  string `json:"shipment_duration_unit"`
	ServiceType           string `json:"service_type"`
	ShippingType          string `json:"shipping_type"`
	Price                 uint   `json:"price"`
}

type OriginRate struct {
	Latitude                         float64 `json:"latitude"`
	Longitude                        float64 `json:"longitude"`
	PostalCode                       uint    `json:"postal_code"`
	CountryName                      string  `json:"country_name"`
	CountryCode                      string  `json:"country_code"`
	AdministrativeDivisionLevel1Name string  `json:"administrative_division_level_1_name"`
	AdministrativeDivisionLevel1Type string  `json:"administrative_division_level_1_type"`
	AdministrativeDivisionLevel2Name string  `json:"administrative_division_level_2_name"`
	AdministrativeDivisionLevel2Type string  `json:"administrative_division_level_2_type"`
	AdministrativeDivisionLevel3Name string  `json:"administrative_division_level_3_name"`
	AdministrativeDivisionLevel3Type string  `json:"administrative_division_level_3_type"`
	AdministrativeDivisionLevel4Name string  `json:"administrative_division_level_4_name"`
	AdministrativeDivisionLevel4Type string  `json:"administrative_division_level_4_type"`
}

type DestinationRate struct {
	Latitude                         float64 `json:"latitude"`
	Longitude                        float64 `json:"longitude"`
	PostalCode                       uint    `json:"postal_code"`
	CountryName                      string  `json:"country_name"`
	CountryCode                      string  `json:"country_code"`
	AdministrativeDivisionLevel1Name string  `json:"administrative_division_level_1_name"`
	AdministrativeDivisionLevel1Type string  `json:"administrative_division_level_1_type"`
	AdministrativeDivisionLevel2Name string  `json:"administrative_division_level_2_name"`
	AdministrativeDivisionLevel2Type string  `json:"administrative_division_level_2_type"`
	AdministrativeDivisionLevel3Name string  `json:"administrative_division_level_3_name"`
	AdministrativeDivisionLevel3Type string  `json:"administrative_division_level_3_type"`
	AdministrativeDivisionLevel4Name string  `json:"administrative_division_level_4_name"`
	AdministrativeDivisionLevel4Type string  `json:"administrative_division_level_4_type"`
}

type ResponseListRatesCouriers struct {
	Success     bool            `json:"success"`
	Message     string          `json:"message"`
	Origin      OriginRate      `json:"origin"`
	Destination DestinationRate `json:"destination"`
	Pricing     []PricingRate   `json:"pricing"`
}

type ResponseCourier struct {
	AvailableForCashOnDelivery   bool   `json:"available_for_cash_on_delivery"`
	AvailableForProofOfDelivery  bool   `json:"available_for_proof_of_delivery"`
	AvailableForInstantWaybillId bool   `json:"available_for_instant_waybill_id"`
	CourierName                  string `json:"courier_name"`
	CourierCode                  string `json:"courier_code"`
	CourierServiceName           string `json:"courier_service_name"`
	CourierServiceCode           string `json:"courier_service_code"`
	Tier                         string `json:"tier"`
	Description                  string `json:"description"`
	ServiceType                  string `json:"service_type"`
	ShippingType                 string `json:"shipping_type"`
	ShipmentDurationRange        string `json:"shipment_duration_range"`
	ShipmentDurationUnit         string `json:"shipment_duration_unit"`
}

type ResponseListCourier struct {
	Success  bool              `json:"success"`
	Object   string            `json:"object"`
	Couriers []ResponseCourier `json:"couriers"`
}

type ResponseTrackingOrder struct {
	Success   bool   `json:"success"`
	Message   string `json:"message"`
	Object    string `json:"object"`
	Id        string `json:"id"`
	WaybillId string `json:"waybill_id"`
	Courier   struct {
		Company string  `json:"company"`
		Name    *string `json:"name"`
		Phone   *string `json:"phone"`
	} `json:"courier"`
	Origin struct {
		ContactName string `json:"contact_name"`
		Address     string `json:"address"`
	} `json:"origin"`
	Destination struct {
		ContactName string `json:"contact_name"`
		Address     string `json:"address"`
	}
	History []struct {
		Note      string `json:"note"`
		Status    string `json:"status"`
		UpdatedAt string `json:"updated_at"`
	} `json:"history"`
	Link    string  `json:"link"`
	OrderId *string `json:"order_id"`
	Status  string  `json:"status"`
}

type ApiResponse struct {
	Status     string
	StatusCode int
	Proto      string
	Header     http.Header
	RawBody    []byte
	Request    *http.Request
}
