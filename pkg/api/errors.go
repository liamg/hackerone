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

type Errors struct {
	Errors []Error `json:"errors"`
}

func (e *Errors) Error() error {
	if len(e.Errors) > 0 {
		return fmt.Errorf("%d: %s", e.Errors[0].Status, e.Errors[0].Detail)
	}
	return fmt.Errorf("No errors")
}
