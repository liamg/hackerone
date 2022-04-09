package hacker

import "github.com/liamg/hackerone/internal/pkg/client"

type API struct {
	client *client.Client
}

func New(c *client.Client) *API {
	return &API{
		client: c,
	}
}
