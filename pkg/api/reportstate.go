package api

type ReportState string

const (
	ReportStateNew                  ReportState = "new"
	ReportStatePendingProgramReview ReportState = "pending-program-review"
	ReportStateTriaged              ReportState = "triaged"
	ReportStateNeedsMoreInfo        ReportState = "needs-more-info"
	ReportStateResolved             ReportState = "resolved"
	ReportStateNotApplicable        ReportState = "not-applicable"
	ReportStateInformative          ReportState = "informative"
	ReportStateDuplicate            ReportState = "duplicate"
	ReportStateSpam                 ReportState = "spam"
	ReportStateRetesting            ReportState = "retesting"
)
