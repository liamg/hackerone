package client

import (
	"bytes"
	"encoding/json"
	"io"
	"net"
	"net/http"
	"strings"
	"time"
)

const defaultBaseURL = "https://api.hackerone.com/v1"

type Client struct {
	baseURL    string
	underlying *http.Client
}

func New(username, apiKey string, options ...Option) *Client {
	c := &Client{
		baseURL: defaultBaseURL,
		underlying: &http.Client{
			Transport: &transport{
				username: username,
				apiKey:   apiKey,
				underlying: http.Transport{
					Proxy: http.ProxyFromEnvironment,
					DialContext: (&net.Dialer{
						Timeout:   30 * time.Second,
						KeepAlive: 30 * time.Second,
					}).DialContext,
					ForceAttemptHTTP2:     true,
					MaxIdleConns:          100,
					IdleConnTimeout:       90 * time.Second,
					TLSHandshakeTimeout:   10 * time.Second,
					ExpectContinueTimeout: 1 * time.Second,
				},
			},
		},
	}
	for _, opt := range options {
		opt(c)
	}
	return c
}

func (c *Client) url(path string) string {
	if strings.Contains(path, "://") {
		return path
	}
	return c.baseURL + path
}

func (c *Client) do(method, path string, target interface{}, body io.Reader) error {
	req, err := http.NewRequest(method, c.url(path), body)
	if err != nil {
		return err
	}
	req.Header.Set("Accept", "application/json")
	resp, err := c.underlying.Do(req)
	if err != nil {
		return err
	}
	defer func() { _ = resp.Body.Close() }()
	if target == nil {
		return nil
	}
	return json.NewDecoder(resp.Body).Decode(target)
}

func (c *Client) Get(path string, target interface{}) error {
	return c.do(http.MethodGet, path, target, nil)
}

func (c *Client) Post(path string, target interface{}, body interface{}) error {
	buffer := bytes.NewBuffer([]byte{})
	if err := json.NewEncoder(buffer).Encode(body); err != nil {
		return err
	}
	return c.do(http.MethodPost, path, target, buffer)
}

func (c *Client) Delete(path string) error {
	return c.do(http.MethodDelete, path, nil, nil)
}

func (c *Client) Patch(path string, target interface{}, body interface{}) error {
	buffer := bytes.NewBuffer([]byte{})
	if err := json.NewEncoder(buffer).Encode(body); err != nil {
		return err
	}
	return c.do(http.MethodPatch, path, target, buffer)
}
