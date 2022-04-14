package hackers

import (
	"context"
	"fmt"

	api2 "github.com/liamg/hackerone/pkg/api"
)

type getPayoutsResponse struct {
	Data  []api2.Payout `json:"data"`
	Links api2.Links    `json:"links"`
}

// GetPayouts returns a list of payouts received by the hacker. If there are further pages, nextPage will be >0.
func (a *API) GetPayouts(ctx context.Context, pageOptions *api2.PageOptions) (payouts []api2.Payout, nextPage int, err error) {
	var response getPayoutsResponse
	path := fmt.Sprintf(
		"/hackers/payments/payouts?page[number]=%d&page[size]=%d",
		pageOptions.GetPageNumber(),
		pageOptions.GetPageSize(),
	)
	if err := a.client.Get(ctx, path, &response); err != nil {
		return nil, 0, err
	}
	if response.Links.Next != "" {
		nextPage = pageOptions.GetPageNumber() + 1
	}
	return response.Data, nextPage, nil
}
