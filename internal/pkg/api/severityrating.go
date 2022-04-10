package api

type SeverityRating string

const (
	SeverityRatingNone     SeverityRating = "none"
	SeverityRatingLow      SeverityRating = "low"
	SeverityRatingMedium   SeverityRating = "medium"
	SeverityRatingHigh     SeverityRating = "high"
	SeverityRatingCritical SeverityRating = "critical"
)
