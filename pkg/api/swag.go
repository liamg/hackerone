package api

import "time"

type Swag struct {
	Id         string `json:"id"`
	Type       string `json:"type"`
	Attributes struct {
		Sent      bool      `json:"sent"`
		CreatedAt time.Time `json:"created_at"`
	} `json:"attributes"`
	Relationships struct {
		User struct {
			Data User `json:"data"`
		} `json:"user"`
		Address struct {
			Data Address `json:"data"`
		} `json:"address"`
	} `json:"relationships"`
}
