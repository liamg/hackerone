package hackers

import (
	"fmt"

	"github.com/liamg/hackerone/internal/pkg/api"
)

type getPayoutsResponse struct {
	Data  []api.Payout `json:"data"`
	Links api.Links    `json:"links"`
}

// GetPayouts returns a list of payouts received by the hacker. If there are further pages, nextPage will be >0.
func (a *API) GetPayouts(pageOptions *api.PageOptions) (payouts []api.Payout, nextPage int, err error) {
	var response getPayoutsResponse
	path := fmt.Sprintf(
		"/hackers/payments/payouts?page[number]=%d&page[size]=%d",
		pageOptions.GetPageNumber(),
		pageOptions.GetPageSize(),
	)
	if err := a.client.Get(path, &response); err != nil {
		return nil, 0, err
	}
	if response.Links.Next != "" {
		nextPage = pageOptions.GetPageNumber() + 1
	}
	return response.Data, nextPage, nil
}
