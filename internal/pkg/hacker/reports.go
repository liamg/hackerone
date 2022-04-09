package hacker

import (
	"fmt"
	"time"
)

type getReportResponse struct {
	Report Report `json:"data"`
}

type Report struct {
	Id         string `json:"id"`
	Type       string `json:"type"`
	Attributes struct {
		Title                    string      `json:"title"`
		State                    string      `json:"state"`
		CreatedAt                time.Time   `json:"created_at"`
		VulnerabilityInformation string      `json:"vulnerability_information"`
		TriagedAt                interface{} `json:"triaged_at"`
		ClosedAt                 interface{} `json:"closed_at"`
		LastReporterActivityAt   interface{} `json:"last_reporter_activity_at"`
		FirstProgramActivityAt   interface{} `json:"first_program_activity_at"`
		LastProgramActivityAt    interface{} `json:"last_program_activity_at"`
		BountyAwardedAt          interface{} `json:"bounty_awarded_at"`
		SwagAwardedAt            interface{} `json:"swag_awarded_at"`
		DisclosedAt              interface{} `json:"disclosed_at"`
		Source                   interface{} `json:"source"`
	} `json:"attributes"`
	Relationships struct {
		Reporter struct {
			Data struct {
				Id         string `json:"id"`
				Type       string `json:"type"`
				Attributes struct {
					Username       string    `json:"username"`
					Name           string    `json:"name"`
					Disabled       bool      `json:"disabled"`
					CreatedAt      time.Time `json:"created_at"`
					ProfilePicture struct {
						X62  string `json:"62x62"`
						X82  string `json:"82x82"`
						X110 string `json:"110x110"`
						X260 string `json:"260x260"`
					} `json:"profile_picture"`
				} `json:"attributes"`
			} `json:"data"`
		} `json:"reporter"`
		Assignee struct {
			Data struct {
				Id         string `json:"id"`
				Type       string `json:"type"`
				Attributes struct {
					Username       string    `json:"username"`
					Name           string    `json:"name"`
					Disabled       bool      `json:"disabled"`
					CreatedAt      time.Time `json:"created_at"`
					ProfilePicture struct {
						X62  string `json:"62x62"`
						X82  string `json:"82x82"`
						X110 string `json:"110x110"`
						X260 string `json:"260x260"`
					} `json:"profile_picture"`
				} `json:"attributes"`
			} `json:"data"`
		} `json:"assignee"`
		Program struct {
			Data struct {
				Id         string `json:"id"`
				Type       string `json:"type"`
				Attributes struct {
					Handle    string    `json:"handle"`
					CreatedAt time.Time `json:"created_at"`
					UpdatedAt time.Time `json:"updated_at"`
				} `json:"attributes"`
			} `json:"data"`
		} `json:"program"`
		Severity struct {
			Data struct {
				Id         string `json:"id"`
				Type       string `json:"type"`
				Attributes struct {
					Rating             string    `json:"rating"`
					AuthorType         string    `json:"author_type"`
					UserId             int       `json:"user_id"`
					CreatedAt          time.Time `json:"created_at"`
					Score              float64   `json:"score"`
					AttackComplexity   string    `json:"attack_complexity"`
					AttackVector       string    `json:"attack_vector"`
					Availability       string    `json:"availability"`
					Confidentiality    string    `json:"confidentiality"`
					Integrity          string    `json:"integrity"`
					PrivilegesRequired string    `json:"privileges_required"`
					UserInteraction    string    `json:"user_interaction"`
					Scope              string    `json:"scope"`
				} `json:"attributes"`
			} `json:"data"`
		} `json:"severity"`
		Swag struct {
			Data []interface{} `json:"data"`
		} `json:"swag"`
		Attachments struct {
			Data []interface{} `json:"data"`
		} `json:"attachments"`
		Weakness struct {
			Data struct {
				Id         string `json:"id"`
				Type       string `json:"type"`
				Attributes struct {
					Name        string    `json:"name"`
					Description string    `json:"description"`
					ExternalId  string    `json:"external_id"`
					CreatedAt   time.Time `json:"created_at"`
				} `json:"attributes"`
			} `json:"data"`
		} `json:"weakness"`
		StructuredScope struct {
			Data struct {
				Id         string `json:"id"`
				Type       string `json:"type"`
				Attributes struct {
					AssetIdentifier            string      `json:"asset_identifier"`
					AssetType                  string      `json:"asset_type"`
					ConfidentialityRequirement string      `json:"confidentiality_requirement"`
					IntegrityRequirement       string      `json:"integrity_requirement"`
					AvailabilityRequirement    string      `json:"availability_requirement"`
					MaxSeverity                string      `json:"max_severity"`
					CreatedAt                  time.Time   `json:"created_at"`
					UpdatedAt                  time.Time   `json:"updated_at"`
					Instruction                interface{} `json:"instruction"`
					EligibleForBounty          bool        `json:"eligible_for_bounty"`
					EligibleForSubmission      bool        `json:"eligible_for_submission"`
					Reference                  string      `json:"reference"`
				} `json:"attributes"`
			} `json:"data"`
		} `json:"structured_scope"`
		Activities struct {
			Data []struct {
				Type       string `json:"type"`
				Id         string `json:"id"`
				Attributes struct {
					Message   string    `json:"message"`
					CreatedAt time.Time `json:"created_at"`
					UpdatedAt time.Time `json:"updated_at"`
					Internal  bool      `json:"internal"`
				} `json:"attributes"`
				Relationships struct {
					Actor struct {
						Data struct {
							Id         string `json:"id"`
							Type       string `json:"type"`
							Attributes struct {
								Username       string    `json:"username"`
								Name           string    `json:"name"`
								Disabled       bool      `json:"disabled"`
								CreatedAt      time.Time `json:"created_at"`
								ProfilePicture struct {
									X62  string `json:"62x62"`
									X82  string `json:"82x82"`
									X110 string `json:"110x110"`
									X260 string `json:"260x260"`
								} `json:"profile_picture"`
								Signal           interface{} `json:"signal"`
								Impact           interface{} `json:"impact"`
								Reputation       interface{} `json:"reputation"`
								Bio              interface{} `json:"bio"`
								Website          interface{} `json:"website"`
								Location         interface{} `json:"location"`
								HackeroneTriager bool        `json:"hackerone_triager,omitempty"`
							} `json:"attributes"`
						} `json:"data"`
					} `json:"actor"`
					Attachments struct {
						Data []struct {
							Id         string `json:"id"`
							Type       string `json:"type"`
							Attributes struct {
								ExpiringUrl string    `json:"expiring_url"`
								CreatedAt   time.Time `json:"created_at"`
								FileName    string    `json:"file_name"`
								ContentType string    `json:"content_type"`
								FileSize    int       `json:"file_size"`
							} `json:"attributes"`
						} `json:"data"`
					} `json:"attachments,omitempty"`
				} `json:"relationships"`
			} `json:"data"`
		} `json:"activities"`
		Bounties struct {
			Data []interface{} `json:"data"`
		} `json:"bounties"`
		Summaries struct {
			Data []interface{} `json:"data"`
		} `json:"summaries"`
		TriggeredPreSubmissionTrigger struct {
			Data struct {
				Id         string `json:"id"`
				Type       string `json:"type"`
				Attributes struct {
					Title string `json:"title"`
				} `json:"attributes"`
			} `json:"data"`
		} `json:"triggered_pre_submission_trigger"`
		CustomFieldValues struct {
			Data []interface{} `json:"data"`
		} `json:"custom_field_values"`
		AutomatedRemediationGuidance struct {
			Data struct {
				Id         string `json:"id"`
				Type       string `json:"type"`
				Attributes struct {
					Reference string    `json:"reference"`
					CreatedAt time.Time `json:"created_at"`
				} `json:"attributes"`
			} `json:"data"`
		} `json:"automated_remediation_guidance"`
		CustomRemediationGuidance struct {
			Data struct {
				Id         string `json:"id"`
				Type       string `json:"type"`
				Attributes struct {
					Message   string    `json:"message"`
					CreatedAt time.Time `json:"created_at"`
				} `json:"attributes"`
				Relationships struct {
					Author struct {
						Data struct {
							Id         string `json:"id"`
							Type       string `json:"type"`
							Attributes struct {
								Username       string `json:"username"`
								Name           string `json:"name"`
								Disabled       bool   `json:"disabled"`
								CreatedAt      string `json:"created_at"`
								ProfilePicture struct {
									X62  string `json:"62x62"`
									X82  string `json:"82x82"`
									X110 string `json:"110x110"`
									X260 string `json:"260x260"`
								} `json:"profile_picture"`
							} `json:"attributes"`
						} `json:"data"`
					} `json:"author"`
				} `json:"relationships"`
			} `json:"data"`
		} `json:"custom_remediation_guidance"`
	} `json:"relationships"`
}

type getReportsResponse struct {
	Reports []ReportSummary `json:"data"`
	Links   links           `json:"links"`
}

type ReportSummary struct {
	Id            string                 `json:"id"`
	Type          string                 `json:"type"`
	Attributes    map[string]interface{} `json:"attributes"`
	Relationships map[string]struct {
		Data struct {
			Id         string                 `json:"id"`
			Type       string                 `json:"type"`
			Attributes map[string]interface{} `json:"attributes"`
		} `json:"data"`
	} `json:"relationships"`
}

type links struct {
	Prev string `json:"prev"`
	Self string `json:"self"`
	Next string `json:"next"`
}

func (a *API) GetReport(id int) (*Report, error) {
	var response getReportResponse
	path := fmt.Sprintf("hackers/reports/%d", id)
	if err := a.client.Get(path, &response); err != nil {
		return nil, err
	}
	return &response.Report, nil
}

func (a *API) GetReports() ([]ReportSummary, error) {
	var reports []ReportSummary
	var response getReportsResponse
	path := "hackers/me/reports"
	for {
		if err := a.client.Get(path, &response); err != nil {
			return nil, err
		}
		reports = append(reports, response.Reports...)
		if response.Links.Next == "" || len(response.Reports) == 0 {
			break
		}
		path = response.Links.Next
	}
	return reports, nil
}
