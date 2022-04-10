package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/liamg/hackerone/internal/pkg/api"

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
			defer func() { _ = resp.Body.Close() }()
			var apiError api.Error
			if err := json.NewDecoder(resp.Body).Decode(&apiError); err != nil || apiError.Status == 0 {
				return fmt.Errorf("server error: status %d", resp.StatusCode)
			}
			return &apiError
		}
		response = resp
		return nil
	}, options...); err != nil {
		return nil, err
	}
	return response, nil
}
