package biteship

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

const (
	secretKey        = "biteship_test.eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoidGVzLXBrZyIsInVzZXJJZCI6IjYxNjQzYmJiNzRkYWMxMzdjMDIyMjUxYyIsImlhdCI6MTY0MTc3OTY0Nn0.LA2Opjs1wNTHeSLDAZpD3W9CqMoMfZvAkOhvSYfIftk"
	invalidSecretKey = "biteship_test.eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoidGVzLXBrZyIsInVzZXJJZCI6IjYxNjQzYmJiNzRkYWMxMzdjMDIyMjUxYyIsImlhdCI6MTY0MTc3OTY0Nn0.LA2Opjs1wNTHeSLDAZpD3W9CqMoMfZvAkOhvSYfIftks"
)

var (
	orderIdConfirmed    = "61E02D940904D76428ADA74E"
	orderIdNotConfirmed = ""
	orderIdCancelled    = "61DBB6B1A4720916B2D1F576"
)

func TestGetCourier(t *testing.T) {
	biteship := New(WithSecret(secretKey))

	resp, _ := biteship.GetCouriers()
	assert.True(t, resp.Success)
	assert.Equal(t, resp.Object, "courier")
}

func TestGetCourierWithInvalidSecretKey(t *testing.T) {
	biteship := New(WithSecret(invalidSecretKey))

	resp, err := biteship.GetCouriers()
	assert.Equal(t, *resp, ResponseListCourier{})
	assert.Equal(t, err.Status, http.StatusBadRequest)
	assert.Equal(t, err.Code, 40000001)
}

func TestGetRatesCouriers(t *testing.T) {
	biteship := New(WithSecret(secretKey))
	var (
		destLat  = -6.2441792
		destLong = 106.783529000
	)

	var items []ItemCourierRate
	items = append(items, ItemCourierRate{
		Name:        "Shoes",
		Description: "Black colored size 45",
		Value:       199000,
		Length:      30,
		Width:       15,
		Height:      20,
		Weight:      200,
		Quantity:    2,
	})

	req := RequestCourierRates{
		DestinationLatitude:  &destLat,
		DestinationLongitude: &destLong,
		OriginPostalCode:     12440,
		Couriers:             fmt.Sprintf("%s,%s", CourierProviderJNE, CourierProviderTIKI),
		Items:                items,
	}

	resp, _ := biteship.GetCourierRates(&req)

	assert.True(t, resp.Success)
}

func TestGetFailRatesCouriers(t *testing.T) {
	biteship := New(WithSecret(invalidSecretKey))
	var (
		destLat  = -6.2441792
		destLong = 106.783529000
	)

	var items []ItemCourierRate
	items = append(items, ItemCourierRate{
		Name:        "Shoes",
		Description: "Black colored size 45",
		Value:       199000,
		Length:      30,
		Width:       15,
		Height:      20,
		Weight:      200,
		Quantity:    2,
	})

	req := RequestCourierRates{
		DestinationLatitude:  &destLat,
		DestinationLongitude: &destLong,
		OriginPostalCode:     12440,
		Couriers:             fmt.Sprintf("%s,%s", CourierProviderJNE, CourierProviderTIKI),
		Items:                items,
	}

	resp, err := biteship.GetCourierRates(&req)
	assert.Equal(t, *resp, ResponseListRatesCouriers{})
	assert.Equal(t, err.Status, http.StatusBadRequest)
	assert.Equal(t, err.Code, 40000001)
}

func TestGetMissingRequiredParamRatesCouriers(t *testing.T) {
	biteship := New(WithSecret(secretKey))
	var (
		destLat  = -6.2441792
		destLong = 106.783529000
	)

	var items []ItemCourierRate
	items = append(items, ItemCourierRate{
		Name:        "Shoes",
		Description: "Black colored size 45",
		Value:       199000,
		Length:      30,
		Width:       15,
		Height:      20,
		Weight:      200,
		Quantity:    2,
	})

	req := RequestCourierRates{
		DestinationLatitude:  &destLat,
		DestinationLongitude: &destLong,
		OriginPostalCode:     12440,
		Items:                items,
	}

	resp, err := biteship.GetCourierRates(&req)
	assert.Equal(t, *resp, ResponseListRatesCouriers{})
	assert.Equal(t, err.Status, http.StatusBadRequest)
}

func TestCreateAnOrderDirectConfirm(t *testing.T) {
	biteship := New(WithSecret(secretKey))

	var items []ProductItem
	items = append(items, ProductItem{
		Id:          "5db7ee67382e185bd6a14608",
		Name:        "Black L",
		Image:       "",
		Description: "White Shirt",
		Value:       165000,
		Quantity:    1,
		Height:      0,
		Length:      0,
		Width:       0,
		Weight:      0,
	})

	req := CreateOrderRequestParam{
		ShipperContactName:  "Amir",
		ShipperContactPhone: "081277882932",
		ShipperContactEmail: "biteship@test.com",
		ShipperOrganization: "Biteship Org Test",
		OriginContactName:   "Amir",
		OriginContactPhone:  "081740781720",
		OriginAddress:       "Plaza Senayan, Jalan Asia Afrik...",
		OriginNote:          "Deket pintu masuk STC",
		OriginPostalCode:    12440,
		OriginCoordinate: Coordinate{
			Latitude:  -6.2253114,
			Longitude: 106.7993735,
		},
		DestinationContactName:  "John Doe",
		DestinationContactPhone: "08170032123",
		DestinationContactEmail: "jon@test.com",
		DestinationAddress:      "Lebak Bulus MRT...",
		DestinationPostalCode:   12950,
		DestinationNote:         "Near the gas station",
		DestinationCoordinate: Coordinate{
			Latitude:  -6.28927,
			Longitude: 106.77492000000007,
		},
		DestinationCashOnDelivery:     nil,
		DestinationCashOnDeliveryType: nil,
		CourierCompany:                "jne",
		CourierType:                   "reg",
		CourierInsurance:              500000,
		DeliveryType:                  DeliveryTypeNow,
		DeliveryDate:                  "2022-01-14",
		DeliveryTime:                  "12:00",
		OrderNote:                     "Please be careful",
		Metadata:                      nil,
		Items:                         items,
	}

	resp, _ := biteship.CreateOrder(&req)
	orderIdConfirmed = resp.Id
	assert.True(t, resp.Success)
	assert.Equal(t, resp.Message, "Order successfully created")
	assert.Equal(t, resp.Object, "order")
}

func TestCreateAnOrderWithDeliveryLater(t *testing.T) {
	biteship := New(WithSecret(secretKey))

	var items []ProductItem
	items = append(items, ProductItem{
		Id:          "5db7ee67382e185bd6a14608",
		Name:        "Black L",
		Image:       "",
		Description: "White Shirt",
		Value:       165000,
		Quantity:    1,
		Height:      0,
		Length:      0,
		Width:       0,
		Weight:      0,
	})

	req := CreateOrderRequestParam{
		ShipperContactName:  "Amir",
		ShipperContactPhone: "081277882932",
		ShipperContactEmail: "biteship@test.com",
		ShipperOrganization: "Biteship Org Test",
		OriginContactName:   "Amir",
		OriginContactPhone:  "081740781720",
		OriginAddress:       "Plaza Senayan, Jalan Asia Afrik...",
		OriginNote:          "Deket pintu masuk STC",
		OriginPostalCode:    12440,
		OriginCoordinate: Coordinate{
			Latitude:  -6.2253114,
			Longitude: 106.7993735,
		},
		DestinationContactName:  "John Doe",
		DestinationContactPhone: "08170032123",
		DestinationContactEmail: "jon@test.com",
		DestinationAddress:      "Lebak Bulus MRT...",
		DestinationPostalCode:   12950,
		DestinationNote:         "Near the gas station",
		DestinationCoordinate: Coordinate{
			Latitude:  -6.28927,
			Longitude: 106.77492000000007,
		},
		DestinationCashOnDelivery:     nil,
		DestinationCashOnDeliveryType: nil,
		CourierCompany:                "jne",
		CourierType:                   "reg",
		CourierInsurance:              500000,
		DeliveryType:                  "later", // later or now
		DeliveryDate:                  "2025-12-14",
		DeliveryTime:                  "12:00",
		OrderNote:                     "Please be careful",
		Metadata:                      nil,
		Items:                         items,
	}

	resp, _ := biteship.CreateOrder(&req)
	orderIdNotConfirmed = resp.Id
	assert.True(t, resp.Success)
	assert.Equal(t, resp.Message, "Order successfully created")
	assert.Equal(t, resp.Object, "order")
}

func TestCreateAnOrderWithInvalidSecretKey(t *testing.T) {
	biteship := New(WithSecret(invalidSecretKey))

	var items []ProductItem
	items = append(items, ProductItem{
		Id:          "5db7ee67382e185bd6a14608",
		Name:        "Black L",
		Image:       "",
		Description: "White Shirt",
		Value:       165000,
		Quantity:    1,
		Height:      0,
		Length:      0,
		Width:       0,
		Weight:      0,
	})

	req := CreateOrderRequestParam{
		ShipperContactName:  "Amir",
		ShipperContactPhone: "081277882932",
		ShipperContactEmail: "biteship@test.com",
		ShipperOrganization: "Biteship Org Test",
		OriginContactName:   "Amir",
		OriginContactPhone:  "081740781720",
		OriginAddress:       "Plaza Senayan, Jalan Asia Afrik...",
		OriginNote:          "Deket pintu masuk STC",
		OriginPostalCode:    12440,
		OriginCoordinate: Coordinate{
			Latitude:  -6.2253114,
			Longitude: 106.7993735,
		},
		DestinationContactName:  "John Doe",
		DestinationContactPhone: "08170032123",
		DestinationContactEmail: "jon@test.com",
		DestinationAddress:      "Lebak Bulus MRT...",
		DestinationPostalCode:   12950,
		DestinationNote:         "Near the gas station",
		DestinationCoordinate: Coordinate{
			Latitude:  -6.28927,
			Longitude: 106.77492000000007,
		},
		DestinationCashOnDelivery:     nil,
		DestinationCashOnDeliveryType: nil,
		CourierCompany:                CourierProviderJNE,
		CourierType:                   "reg",
		CourierInsurance:              500000,
		DeliveryType:                  DeliveryTypeNow,
		DeliveryDate:                  "2022-01-11",
		DeliveryTime:                  "12:00",
		OrderNote:                     "Please be careful",
		Metadata:                      nil,
		Items:                         items,
	}

	resp, err := biteship.CreateOrder(&req)
	assert.Equal(t, *resp, ResponseCreateOrder{})
	assert.Equal(t, err.Status, http.StatusBadRequest)
	assert.Equal(t, err.Code, 40000001)
}

func TestRetrieveOrder(t *testing.T) {
	biteship := New(WithSecret(secretKey))

	resp, _ := biteship.RetrieveOrder(orderIdConfirmed)
	assert.True(t, resp.Success)
	assert.Equal(t, resp.Object, "order")
	assert.Equal(t, resp.Message, "Order successfully retrieved")
}

func TestRetrieveOrderWithInvalidKey(t *testing.T) {
	biteship := New(WithSecret(invalidSecretKey))

	resp, err := biteship.RetrieveOrder(orderIdConfirmed)
	assert.Equal(t, *resp, ResponseRetrieveOrder{})
	assert.Equal(t, err.Status, http.StatusBadRequest)
	assert.Equal(t, err.Code, 40000001)
}

func TestRetrieveOrderWithInvalidOrderId(t *testing.T) {
	biteship := New(WithSecret(secretKey))

	resp, err := biteship.RetrieveOrder("61dce8889a096b081e70198f")

	assert.Equal(t, *resp, ResponseRetrieveOrder{})
	assert.Equal(t, err.Status, http.StatusBadRequest)
	assert.Equal(t, err.Code, ErrOrderOrderNotFound)
}

func TestUpdateOrderBeforeConfirmed(t *testing.T) {
	biteship := New(WithSecret(secretKey))

	requestUpdate := struct {
		OriginAddress string `json:"origin_address"`
	}{OriginAddress: "Jalan Perubahan nomor 5"}

	resp, _ := biteship.UpdateOrder(orderIdNotConfirmed, requestUpdate)

	assert.True(t, resp.Success)
	assert.Equal(t, resp.Object, "order")
	assert.Equal(t, resp.Message, "Order has been updated")
}

func TestUpdateOrderThatWasConfirmed(t *testing.T) {
	biteship := New(WithSecret(secretKey))

	requestUpdate := struct {
		OriginAddress string `json:"origin_address"`
	}{OriginAddress: "Jalan Perubahan nomor 5"}

	_, err := biteship.UpdateOrder(orderIdConfirmed, requestUpdate)

	assert.Equal(t, err.Status, http.StatusBadRequest)
	assert.EqualValues(t, err.Code, ErrOrderCannotEditConfirmedOrder)
	assert.Equal(t, err.RawError, "Order has already been confirmed therefore cannot edit order. Please create new order instead.")
}

func TestUpdateOrderThatWasPassedTime(t *testing.T) {
	biteship := New(WithSecret(secretKey))

	requestUpdate := struct {
		OriginAddress string `json:"origin_address"`
	}{OriginAddress: "Jalan Perubahan nomor 5"}

	_, err := biteship.UpdateOrder(orderIdCancelled, requestUpdate)
	assert.Equal(t, err.Status, http.StatusBadRequest)
	assert.EqualValues(t, err.Code, ErrOrderTimeAlreadyPassed)
	assert.Equal(t, err.RawError, "Time already passed. Set new delivery time.")
}

func TestConfirmOrder(t *testing.T) {
	biteship := New(WithSecret(secretKey))

	resp, _ := biteship.ConfirmOrder(orderIdNotConfirmed)
	assert.True(t, resp.Success)
	assert.Equal(t, resp.Object, "order")
	assert.EqualValues(t, resp.Status, StatusConfirmed)
	assert.Equal(t, resp.Message, "Success: Order status updated to confirmed")
}

func TestConfirmOrderThatWasBeenConfirmed(t *testing.T) {
	biteship := New(WithSecret(secretKey))

	_, err := biteship.ConfirmOrder(orderIdConfirmed)
	assert.Equal(t, err.Status, http.StatusBadRequest)
	assert.EqualValues(t, err.Code, ErrOrderOrderAlreadyInNewStatus)
	assert.Equal(t, err.RawError, "Order has already been confirmed")
}

func TestConfirmOrderThatWasBeenCancelled(t *testing.T) {
	biteship := New(WithSecret(secretKey))

	_, err := biteship.ConfirmOrder(orderIdCancelled)
	assert.Equal(t, err.Status, http.StatusBadRequest)
	assert.EqualValues(t, err.Code, ErrOrderOrderAlreadyInNewStatus)
	assert.Equal(t, err.RawError, "Order has already been cancelled")
}

func TestCancelOrder(t *testing.T) {
	biteship := New(WithSecret(secretKey))

	reason := "Ingin mengganti kurir"

	resp, err := biteship.CancelOrder(orderIdNotConfirmed, reason)

	assert.Nil(t, err)
	assert.True(t, resp.Success)
	assert.Equal(t, resp.Object, "order")
	assert.EqualValues(t, resp.Status, StatusCancelled)
	assert.Equal(t, resp.Id, orderIdNotConfirmed)
	assert.Equal(t, resp.CancellationReason, reason)
}

func TestCancelOrderThatWasConfirmed(t *testing.T) {
	biteship := New(WithSecret(secretKey))

	reason := "Ingin mengganti kurir"

	resp, err := biteship.CancelOrder(orderIdConfirmed, reason)

	assert.Nil(t, err)
	assert.True(t, resp.Success)
	assert.Equal(t, resp.Object, "order")
	assert.EqualValues(t, resp.Status, StatusCancelled)
	assert.Equal(t, resp.Id, orderIdConfirmed)
	assert.Equal(t, resp.CancellationReason, reason)
}
