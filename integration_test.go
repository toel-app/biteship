package biteship

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

const secretKey = "biteship_test.eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoidGVzLXBrZyIsInVzZXJJZCI6IjYxNjQzYmJiNzRkYWMxMzdjMDIyMjUxYyIsImlhdCI6MTY0MTc3OTY0Nn0.LA2Opjs1wNTHeSLDAZpD3W9CqMoMfZvAkOhvSYfIftk"
const invalidSecretKey = "biteship_test.eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoidGVzLXBrZyIsInVzZXJJZCI6IjYxNjQzYmJiNzRkYWMxMzdjMDIyMjUxYyIsImlhdCI6MTY0MTc3OTY0Nn0.LA2Opjs1wNTHeSLDAZpD3W9CqMoMfZvAkOhvSYfIftks"

//	CHANGE ORDER ID BELOW WITH YOUR ORDER ID
var orderIdConfirmed = "61E02D940904D76428ADA74E"
var orderIdNotConfirmed = ""
var orderIdCancelled = "61DBB6B1A4720916B2D1F576"

//	TEST GET COURIER
func TestGetCourier(t *testing.T) {
	biteship := New(secretKey)

	resp, _ := biteship.GetCouriers()
	assert.Equal(t, resp.Success, true)
	assert.Equal(t, resp.Object, "courier")
}

//	TEST FAIL GET COURIER WITH INVALID SECRET KEY
func TestGetCourierWithInvalidSecretKey(t *testing.T) {
	biteship := New(invalidSecretKey)

	resp, err := biteship.GetCouriers()
	assert.Equal(t, *resp, ResponseListCourier{})
	assert.Equal(t, err.Status, 400)
	assert.Equal(t, err.Code, 40000001)
}

//	TEST CHECK RATES OF COURIER
func TestGetRatesCouriers(t *testing.T) {
	biteship := New(secretKey)

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
		DestinationLatitude:  -6.2441792,
		DestinationLongitude: 106.783529000,
		OriginPostalCode:     12440,
		Couriers:             "jne,tiki",
		Items:                items,
	}

	resp, _ := biteship.GetCourierRates(&req)
	log.Println(resp)
	assert.Equal(t, resp.Success, true)
}

func TestGetFailRatesCouriers(t *testing.T) {
	biteship := New(invalidSecretKey)

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
		//OriginLatitude:       -6.3031123,
		//OriginLongitude:      106.7794934999,
		DestinationLatitude:  -6.2441792,
		DestinationLongitude: 106.783529000,
		OriginPostalCode:     12440,
		//DestinationPostalCode: 12240,
		Couriers: "jne,tiki",
		Items:    items,
	}

	resp, err := biteship.GetCourierRates(&req)
	assert.Equal(t, *resp, ResponseListRatesCouriers{})
	assert.Equal(t, err.Status, 400)
	assert.Equal(t, err.Code, 40000001)
}

func TestGetMissingRequiredParamRatesCouriers(t *testing.T) {
	biteship := New(secretKey)

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
		//OriginLatitude:       -6.3031123,
		//OriginLongitude:      106.7794934999,
		DestinationLatitude:  -6.2441792,
		DestinationLongitude: 106.783529000,
		OriginPostalCode:     12440,
		//DestinationPostalCode: 12240,
		//Couriers: "jne,tiki", // required
		Items: items,
	}

	resp, err := biteship.GetCourierRates(&req)
	assert.Equal(t, *resp, ResponseListRatesCouriers{})
	assert.Equal(t, err.Status, 400)
}

func TestCreateAnOrderDirectConfirm(t *testing.T) {
	biteship := New(secretKey)

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
		DeliveryType:                  "now", // later or now
		DeliveryDate:                  "2022-01-14",
		DeliveryTime:                  "12:00",
		OrderNote:                     "Please be carefull",
		Metadata:                      nil,
		Items:                         items,
	}

	resp, _ := biteship.CreateOrder(&req)
	orderIdConfirmed = resp.Id
	assert.Equal(t, resp.Success, true)
	assert.Equal(t, resp.Message, "Order successfully created")
	assert.Equal(t, resp.Object, "order")
}

func TestCreateAnOrderWithDeliveryLater(t *testing.T) {
	biteship := New(secretKey)

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
		OrderNote:                     "Please be carefull",
		Metadata:                      nil,
		Items:                         items,
	}

	resp, _ := biteship.CreateOrder(&req)
	orderIdNotConfirmed = resp.Id
	assert.Equal(t, resp.Success, true)
	assert.Equal(t, resp.Message, "Order successfully created")
	assert.Equal(t, resp.Object, "order")
}

func TestCreateAnOrderWithInvalidSecretKey(t *testing.T) {
	biteship := New(invalidSecretKey)

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
		DeliveryType:                  "now",
		DeliveryDate:                  "2022-01-11",
		DeliveryTime:                  "12:00",
		OrderNote:                     "Please be carefull",
		Metadata:                      nil,
		Items:                         items,
	}

	resp, err := biteship.CreateOrder(&req)
	assert.Equal(t, *resp, ResponseCreateOrder{})
	assert.Equal(t, err.Status, 400)
	assert.Equal(t, err.Code, 40000001)
}

func TestRetrieveOrder(t *testing.T) {
	biteship := New(secretKey)

	resp, _ := biteship.RetrieveOrder(orderIdConfirmed)
	assert.Equal(t, resp.Success, true)
	assert.Equal(t, resp.Object, "order")
	assert.Equal(t, resp.Message, "Order successfully retrieved")
}

func TestRetrieveOrderWithInvalidKey(t *testing.T) {
	biteship := New(invalidSecretKey)

	resp, err := biteship.RetrieveOrder(orderIdConfirmed)
	assert.Equal(t, *resp, ResponseRetrieveOrder{})
	assert.Equal(t, err.Status, 400)
	assert.Equal(t, err.Code, 40000001)
}

func TestRetrieveOrderWithInvalidOrderId(t *testing.T) {
	biteship := New(secretKey)

	resp, err := biteship.RetrieveOrder("61dce8889a096b081e70198f")
	assert.Equal(t, *resp, ResponseRetrieveOrder{})
	assert.Equal(t, err.Status, 400)
	assert.Equal(t, err.Code, 40002057)
}

func TestUpdateOrderBeforeConfirmed(t *testing.T) {
	biteship := New(secretKey)

	requestUpdate := struct {
		OriginAddress string `json:"origin_address"`
	}{OriginAddress: "Jalan Perubahan nomor 5"}

	resp, _ := biteship.UpdateOrder(orderIdNotConfirmed, requestUpdate)
	assert.Equal(t, resp.Success, true)
	assert.Equal(t, resp.Object, "order")
	assert.Equal(t, resp.Message, "Order has been updated")
}

func TestUpdateOrderThatWasConfirmed(t *testing.T) {
	biteship := New(secretKey)

	requestUpdate := struct {
		OriginAddress string `json:"origin_address"`
	}{OriginAddress: "Jalan Perubahan nomor 5"}

	_, err := biteship.UpdateOrder(orderIdConfirmed, requestUpdate)
	log.Println(err)
	assert.Equal(t, err.Status, 400)
	assert.Equal(t, err.Code, 40002044)
	assert.Equal(t, err.RawError, "Order has already been confirmed therefore cannot edit order. Please create new order instead.")
}

func TestUpdateOrderThatWasPassedTime(t *testing.T) {
	biteship := New(secretKey)

	requestUpdate := struct {
		OriginAddress string `json:"origin_address"`
	}{OriginAddress: "Jalan Perubahan nomor 5"}

	_, err := biteship.UpdateOrder(orderIdCancelled, requestUpdate)
	log.Println(err)
	assert.Equal(t, err.Status, 400)
	assert.Equal(t, err.Code, 40002013)
	assert.Equal(t, err.RawError, "Time already passed. Set new delivery time.")
}

func TestConfirmOrder(t *testing.T) {
	biteship := New(secretKey)

	resp, _ := biteship.ConfirmOrder(orderIdNotConfirmed)
	assert.Equal(t, resp.Success, true)
	assert.Equal(t, resp.Object, "order")
	assert.Equal(t, resp.Status, "confirmed")
	assert.Equal(t, resp.Message, "Success: Order status updated to confirmed")
}

func TestConfirmOrderThatWasBeenConfirmed(t *testing.T) {
	biteship := New(secretKey)

	_, err := biteship.ConfirmOrder(orderIdConfirmed)
	assert.Equal(t, err.Status, 400)
	assert.Equal(t, err.Code, 40002050)
	assert.Equal(t, err.RawError, "Order has already been confirmed")
}

func TestConfirmOrderThatWasBeenCancelled(t *testing.T) {
	biteship := New(secretKey)

	_, err := biteship.ConfirmOrder(orderIdCancelled)
	assert.Equal(t, err.Status, 400)
	assert.Equal(t, err.Code, 40002050)
	assert.Equal(t, err.RawError, "Order has already been cancelled")
}

func TestCancelOrder(t *testing.T) {
	biteship := New(secretKey)

	reason := "Ingin mengganti kurir"

	resp, err := biteship.CancelOrder(orderIdNotConfirmed, reason)
	log.Println(err)
	log.Println(resp)
	assert.Equal(t, resp.Success, true)
	assert.Equal(t, resp.Object, "order")
	assert.Equal(t, resp.Status, "cancelled")
	assert.Equal(t, resp.Id, orderIdNotConfirmed)
	assert.Equal(t, resp.CancellationReason, reason)
}

func TestCancelOrderThatWasConfirmed(t *testing.T) {
	biteship := New(secretKey)

	reason := "Ingin mengganti kurir"

	resp, err := biteship.CancelOrder(orderIdConfirmed, reason)
	log.Println(err)
	log.Println(resp)
	assert.Equal(t, resp.Success, true)
	assert.Equal(t, resp.Object, "order")
	assert.Equal(t, resp.Status, "cancelled")
	assert.Equal(t, resp.Id, orderIdConfirmed)
	assert.Equal(t, resp.CancellationReason, reason)
}
