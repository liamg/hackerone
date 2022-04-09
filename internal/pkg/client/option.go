package client

type Option func(c *Client)

func OptionWithBaseURL(baseURL string) Option {
	return func(c *Client) {
		c.baseURL = baseURL
	}
}
