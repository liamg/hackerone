package client

import "strings"

type Option func(c *Client)

func OptionWithBaseURL(baseURL string) Option {
	return func(c *Client) {
		c.baseURL = strings.TrimSuffix(baseURL, "/")
	}
}
