package api

import "time"

type StructuredScope struct {
	Id         string `json:"id"`
	Type       string `json:"type"`
	Attributes struct {
		AssetIdentifier            string    `json:"asset_identifier"`
		AssetType                  string    `json:"asset_type"`
		ConfidentialityRequirement *string   `json:"confidentiality_requirement"`
		IntegrityRequirement       *string   `json:"integrity_requirement"`
		AvailabilityRequirement    *string   `json:"availability_requirement"`
		MaxSeverity                string    `json:"max_severity"`
		CreatedAt                  time.Time `json:"created_at"`
		UpdatedAt                  time.Time `json:"updated_at"`
		Instruction                *string   `json:"instruction"`
		EligibleForBounty          bool      `json:"eligible_for_bounty"`
		EligibleForSubmission      bool      `json:"eligible_for_submission"`
		Reference                  *string   `json:"reference"`
	} `json:"attributes"`
}
