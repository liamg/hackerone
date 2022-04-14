package api

import "time"

type AutomatedRemediationGuidance struct {
	Id         string `json:"id"`
	Type       string `json:"type"`
	Attributes struct {
		Reference string    `json:"reference"`
		CreatedAt time.Time `json:"created_at"`
	} `json:"attributes"`
}
