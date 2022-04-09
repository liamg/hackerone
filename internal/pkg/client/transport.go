package client

import (
	"net/http"
	"time"

	"github.com/avast/retry-go"
)

type transport struct {
	underlying http.Transport
	username   string
	apiKey     string
}

func (t *transport) RoundTrip(req *http.Request) (*http.Response, error) {
	var response *http.Response
	req.SetBasicAuth(t.username, t.apiKey)
	options := []retry.Option{
		retry.Delay(time.Millisecond * 100),
		retry.DelayType(retry.BackOffDelay),
		retry.Attempts(10),
		retry.LastErrorOnly(true),
	}
	if err := retry.Do(func() error {
		resp, err := t.underlying.RoundTrip(req)
		if err != nil {
			return err
		}
		if resp.StatusCode >= 400 {
			_ = resp.Body.Close()
			return &APIError{StatusCode: resp.StatusCode}
		}
		response = resp
		return nil
	}, options...); err != nil {
		return nil, err
	}
	return response, nil
}
