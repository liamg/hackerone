package api

type Earning struct {
	Id         int    `json:"id"`
	Type       string `json:"type"`
	Attributes struct {
		Amount int `json:"amount"`
	} `json:"attributes"`
	Relationships struct {
		Team *struct {
			Data Program `json:"data"`
		} `json:"team"`
		Bounty *struct {
			Data Bounty `json:"data"`
		} `json:"bounty"`
		Pentester *struct {
			Data User `json:"data"`
		} `json:"pentester"`
		ReportRetestUser *struct {
			Data User `json:"data"`
		} `json:"report_retest_user"`
	} `json:"relationships"`
}
