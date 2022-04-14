package api

import "time"

type Activity struct {
	Type       string `json:"type"`
	Id         string `json:"id"`
	Attributes struct {
		Message   *string   `json:"message"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		Internal  bool      `json:"internal"`
	} `json:"attributes"`
	Relationships struct {
		// TODO add user/program etc.
		Attachments struct {
			Data []Attachment `json:"data"`
		} `json:"attachments,omitempty"`
	} `json:"relationships"`
}
