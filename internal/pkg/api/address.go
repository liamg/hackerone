package api

import "time"

type Address struct {
	Id         string `json:"id"`
	Type       string `json:"type"`
	Attributes struct {
		Name        string    `json:"name"`
		Street      string    `json:"street"`
		City        string    `json:"city"`
		PostalCode  string    `json:"postal_code"`
		State       string    `json:"state"`
		Country     string    `json:"country"`
		CreatedAt   time.Time `json:"created_at"`
		TShirtSize  string    `json:"tshirt_size"`
		PhoneNumber string    `json:"phone_number"`
	} `json:"attributes"`
}
