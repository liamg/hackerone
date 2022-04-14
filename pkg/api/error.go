package api

import "fmt"

type Error struct {
	Status int    `json:"status"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
	Source struct {
		Parameter string `json:"parameter"`
	} `json:"source"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s: %s", e.Title, e.Detail)
}
