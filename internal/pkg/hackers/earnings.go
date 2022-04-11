package hackers

import (
	"context"
	"fmt"

	"github.com/liamg/hackerone/internal/pkg/api"
)

type getEarningsResponse struct {
	Data  []api.Earning `json:"data"`
	Links api.Links     `json:"links"`
}

// GetEarnings returns a list of earnings made by the hacker. If there are further pages, nextPage will be >0.
func (a *API) GetEarnings(ctx context.Context, pageOptions *api.PageOptions) (earnings []api.Earning, nextPage int, err error) {
	var response getEarningsResponse
	path := fmt.Sprintf(
		"/hackers/payments/earnings?page[number]=%d&page[size]=%d",
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
