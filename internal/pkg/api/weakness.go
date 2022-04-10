package api

import "time"

type Weakness struct {
	Id         string `json:"id"`
	Type       string `json:"type"`
	Attributes struct {
		Name        string    `json:"name"`
		Description string    `json:"description"`
		ExternalId  string    `json:"external_id"`
		CreatedAt   time.Time `json:"created_at"`
	} `json:"attributes"`
}
