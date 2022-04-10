package api

import "time"

type ProgramSmall struct {
	Id         string `json:"id"`
	Type       string `json:"type"`
	Attributes struct {
		Handle    string     `json:"handle"`
		CreatedAt *time.Time `json:"created_at"`
		UpdatedAt *time.Time `json:"updated_at"`
	} `json:"attributes"`
}

type Program struct {
	Id         string `json:"id"`
	Type       string `json:"type"`
	Attributes struct {
		Handle        string `json:"handle"`
		Relationships struct {
			StructuredScopes struct {
				Data []StructuredScope `json:"data"`
			} `json:"structured_scopes"`
		} `json:"relationships"`
	} `json:"attributes"`
}
