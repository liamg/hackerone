package client

import "fmt"

type APIError struct {
	StatusCode int
}

func (e *APIError) Error() string {
	return fmt.Sprintf("api error: status code %d", e.StatusCode)
}
