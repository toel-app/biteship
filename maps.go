package biteship

import (
	"fmt"
	"net/http"
	"net/url"
)

type Area struct {
	ID                               string `json:"id"`
	PostalCode                       uint   `json:"postal_code"`
	CountryName                      string `json:"country_name"`
	CountryCode                      string `json:"country_code"`
	AdministrativeDivisionLevel1Name string `json:"administrative_division_level_1_name"`
	AdministrativeDivisionLevel1Type string `json:"administrative_division_level_1_type"`
	AdministrativeDivisionLevel2Name string `json:"administrative_division_level_2_name"`
	AdministrativeDivisionLevel2Type string `json:"administrative_division_level_2_type"`
	AdministrativeDivisionLevel3Name string `json:"administrative_division_level_3_name"`
	AdministrativeDivisionLevel3Type string `json:"administrative_division_level_3_type"`
	AdministrativeDivisionLevel4Name string `json:"administrative_division_level_4_name"`
	AdministrativeDivisionLevel4Type string `json:"administrative_division_level_4_type"`
}

type RetrieveAreaResponse struct {
	Success bool   `json:"success"`
	Object  string `json:"object"`
	Areas   []Area `json:"areas"`
}

func (client *Client) RetrieveArea(countries, input string) (*RetrieveAreaResponse, *Error) {
	var response RetrieveAreaResponse

	v := url.Values{}
	v.Set("countries", countries)
	v.Set("input", input)

	if err := client.HttpRequest.Call(http.MethodGet, fmt.Sprintf("%s/v1/maps/areas?%s", client.BiteshipUrl, v.Encode()), client.SecretKey, nil, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (client *Client) RetrieveAreaByID(id string) (*RetrieveAreaResponse, *Error) {
	var response RetrieveAreaResponse

	if err := client.HttpRequest.Call(http.MethodGet, fmt.Sprintf("%s/v1/maps/areas/%s", client.BiteshipUrl, id), client.SecretKey, nil, &response); err != nil {
		return nil, err
	}

	return &response, nil
}
