package hackerone

import (
	"github.com/liamg/hackerone/internal/pkg/client"
	"github.com/liamg/hackerone/internal/pkg/hacker"
)

type API struct {
	//Customers customer.Client
	Hackers *hacker.API
}

func New(username, apiKey string) *API {

	c := client.New(username, apiKey)

	return &API{
		Hackers: hacker.New(c),
	}
}
