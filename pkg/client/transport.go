package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/liamg/hackerone/pkg/api"

	"github.com/avast/retry-go"
)

type transport struct {
	underlying http.Transport
	username   string
	apiKey     string
}

func (t *transport) RoundTrip(req *http.Request) (response *http.Response, err error) {
	req.SetBasicAuth(t.username, t.apiKey)
	options := []retry.Option{
		retry.Delay(time.Millisecond * 100),
		retry.DelayType(retry.BackOffDelay),
		retry.Attempts(5),
		retry.LastErrorOnly(true),
	}
	err = retry.Do(func() error {
		resp, err := t.underlying.RoundTrip(req)
		if err != nil {
			return err
		}
		if resp.StatusCode >= 400 {
			defer resp.Body.Close()
			var apiErrors api.Errors
			err = json.NewDecoder(resp.Body).Decode(&apiErrors)
			if err != nil {
				err = fmt.Errorf("server error: status %d, %s", resp.StatusCode, resp.Body)
			} else {
				err = apiErrors.Error()
			}
			if resp.StatusCode < 500 {
				// Client error
				return retry.Unrecoverable(err)
			}
			// Server error
			return err
		}
		response = resp
		return nil
	}, options...)
	return
}
