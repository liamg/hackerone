package api

import "time"

type ProgramSmall struct {
	Id         string `json:"id"`
	Type       string `json:"type"`
	Attributes struct {
		Handle    string     `json:"handle"`
		CreatedAt *time.Time `json:"created_at"`
		UpdatedAt *time.Time `json:"updated_at"`
	} `json:"attributes"`
}

type Program struct {
	Id         string `json:"id"`
	Type       string `json:"type"`
	Attributes struct {
		Handle                          string     `json:"handle"`
		Name                            string     `json:"name"`
		Currency                        string     `json:"currency"`
		ProfilePicture                  string     `json:"profile_picture"`
		SubmissionState                 string     `json:"submission_state"`
		TriageActive                    bool       `json:"triage_active"`
		State                           string     `json:"state"`
		StartedAcceptingAt              time.Time  `json:"started_accepting_at"`
		NumberOfReportsForUser          int        `json:"number_of_reports_for_user"`
		NumberOfValidReportsForUser     int        `json:"number_of_valid_reports_for_user"`
		BountyEarnedForUser             float64    `json:"bounty_earned_for_user"`
		LastInvitationAcceptedAtForUser *time.Time `json:"last_invitation_accepted_at_for_user"`
		Bookmarked                      bool       `json:"bookmarked"`
		AllowsBountySplitting           bool       `json:"allows_bounty_splitting"`
		OffersBounties                  bool       `json:"offers_bounties"`
		Relationships                   struct {
			StructuredScopes struct {
				Data []StructuredScope `json:"data"`
			} `json:"structured_scopes"`
		} `json:"relationships"`
	} `json:"attributes"`
}
