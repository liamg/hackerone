package hackerone

import (
	"github.com/liamg/hackerone/internal/pkg/client"
	"github.com/liamg/hackerone/internal/pkg/hackers"
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
