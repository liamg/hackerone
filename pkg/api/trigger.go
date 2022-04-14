package api

type Trigger struct {
	Id         string `json:"id"`
	Type       string `json:"type"`
	Attributes struct {
		Title string `json:"title"`
	} `json:"attributes"`
}
