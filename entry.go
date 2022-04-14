package hackerone

import (
	"github.com/liamg/hackerone/pkg/client"
	"github.com/liamg/hackerone/pkg/hackers"
)

type API struct {
	Hackers *hackers.API
}

func New(username, apiKey string) *API {

	c := client.New(username, apiKey)

	return &API{
		Hackers: hackers.New(c),
	}
}
