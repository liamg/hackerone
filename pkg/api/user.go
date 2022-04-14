package api

import "time"

type User struct {
	Id         string `json:"id"`
	Type       string `json:"type"`
	Attributes struct {
		Username       string    `json:"username"`
		Name           string    `json:"name"`
		Disabled       bool      `json:"disabled"`
		CreatedAt      time.Time `json:"created_at"`
		ProfilePicture struct {
			X62  string `json:"62x62"`
			X82  string `json:"82x82"`
			X110 string `json:"110x110"`
			X260 string `json:"260x260"`
		} `json:"profile_picture"`
	} `json:"attributes"`
}
