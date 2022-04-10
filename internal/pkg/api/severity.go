package api

import "time"

type Severity struct {
	Id         string `json:"id"`
	Type       string `json:"type"`
	Attributes struct {
		Rating             SeverityRating `json:"rating"`
		AuthorType         string         `json:"author_type"`
		UserId             int            `json:"user_id"`
		CreatedAt          time.Time      `json:"created_at"`
		Score              *float64       `json:"score"`
		AttackComplexity   string         `json:"attack_complexity"`
		AttackVector       *string        `json:"attack_vector"`
		Availability       string         `json:"availability"`
		Confidentiality    string         `json:"confidentiality"`
		Integrity          string         `json:"integrity"`
		PrivilegesRequired string         `json:"privileges_required"`
		UserInteraction    string         `json:"user_interaction"`
		Scope              *string        `json:"scope"`
	} `json:"attributes"`
}
