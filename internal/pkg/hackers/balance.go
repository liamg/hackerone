package hackers

import "context"

type getBalanceResponse struct {
	Data struct {
		Balance float64 `json:"balance"`
	} `json:"data"`
}

// GetBalance returns your account balance.
func (a *API) GetBalance(ctx context.Context) (float64, error) {
	var response getBalanceResponse
	path := "/hackers/payments/balance"
	if err := a.client.Get(ctx, path, &response); err != nil {
		return 0, err
	}
	return response.Data.Balance, nil
}
