package api

import "time"

type Bounty struct {
	Id         string `json:"id"`
	Type       string `json:"type"`
	Attributes struct {
		Amount             *string   `json:"amount"`
		BonusAmount        *string   `json:"bonus_amount"`
		AwardedAmount      *string   `json:"awarded_amount"`
		AwardedBonusAmount *string   `json:"awarded_bonus_amount"`
		AwardedCurrency    *string   `json:"awarded_currency"`
		CreatedAt          time.Time `json:"created_at"`
		Invitations        []struct {
			Id             string  `json:"id"`
			RecipientEmail *string `json:"recipient"`
			ClaimUrl       *string `json:"claim_url"`
		} `json:"invitations"`
	} `json:"attributes"`
}
