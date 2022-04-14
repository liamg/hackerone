package api

import "time"

type Attachment struct {
	Id         string `json:"id"`
	Type       string `json:"type"`
	Attributes struct {
		ExpiringUrl string    `json:"expiring_url"`
		CreatedAt   time.Time `json:"created_at"`
		FileName    string    `json:"file_name"`
		ContentType string    `json:"content_type"`
		FileSize    int       `json:"file_size"`
	} `json:"attributes"`
}
