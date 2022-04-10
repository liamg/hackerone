package api

import "time"

type NewReport struct {
	Type       string `json:"type"`
	Attributes struct {
		TeamHandle               string `json:"team_handle"`
		Title                    string `json:"title"`
		VulnerabilityInformation string `json:"vulnerability_information"`
		Impact                   string `json:"impact"`
		SeverityRating           string `json:"severity_rating"`
		WeaknessId               int    `json:"weakness_id"`
		StructuredScopeId        int    `json:"structured_scope_id"`
	} `json:"attributes"`
}

type Report struct {
	Id         string `json:"id"`
	Type       string `json:"type"`
	Attributes struct {
		Title                         string      `json:"title"`
		State                         ReportState `json:"state"`
		CreatedAt                     time.Time   `json:"created_at"`
		VulnerabilityInformation      *string     `json:"vulnerability_information"`
		TriagedAt                     *time.Time  `json:"triaged_at"`
		ClosedAt                      *time.Time  `json:"closed_at"`
		LastReporterActivityAt        *time.Time  `json:"last_reporter_activity_at"`
		FirstProgramActivityAt        *time.Time  `json:"first_program_activity_at"`
		LastProgramActivityAt         *time.Time  `json:"last_program_activity_at"`
		BountyAwardedAt               *time.Time  `json:"bounty_awarded_at"`
		SwagAwardedAt                 *time.Time  `json:"swag_awarded_at"`
		DisclosedAt                   *time.Time  `json:"disclosed_at"`
		ReporterAgreedOnGoingPublicAt *time.Time  `json:"reporter_agreed_on_going_public_at"`
		LastPublicActivityAt          *time.Time  `json:"last_public_activity_at"`
		LastActivityAt                *time.Time  `json:"last_activity_at"`
		CVEIDs                        []string    `json:"cve_ids"`
	} `json:"attributes"`
	Relationships struct {
		Program struct {
			Data ProgramSmall `json:"data"`
		} `json:"program"`
		Attachments struct {
			Data []Attachment `json:"data"`
		} `json:"attachments"`
		Swag struct {
			Data []interface{} `json:"data"`
		} `json:"swag"`
		Reporter *struct {
			Data User `json:"data"`
		} `json:"reporter"`
		Assignee *struct { // TODO: this should support user and group
			Data User `json:"data"`
		} `json:"assignee"`
		Weakness *struct {
			Data Weakness `json:"data"`
		} `json:"weakness"`
		Severity *struct {
			Data Severity `json:"data"`
		} `json:"severity"`
		StructuredScope *struct {
			Data StructuredScope `json:"data"`
		} `json:"structured_scope"`
		Activities struct {
			Data []Activity `json:"data"`
		} `json:"activities"`
		Bounties struct {
			Data []interface{} `json:"data"`
		} `json:"bounties"`
		Summaries struct {
			Data []interface{} `json:"data"`
		} `json:"summaries"`
		TriggeredPreSubmissionTrigger *struct {
			Data Trigger `json:"data"`
		} `json:"triggered_pre_submission_trigger"`
		CustomFieldValues struct {
			Data []interface{} `json:"data"`
		} `json:"custom_field_values"`
		AutomatedRemediationGuidance *struct {
			Data AutomatedRemediationGuidance `json:"data"`
		} `json:"automated_remediation_guidance"`
		CustomRemediationGuidance *struct {
			Data CustomRemediationGuidance `json:"data"`
		} `json:"custom_remediation_guidance"`
	} `json:"relationships"`
}
