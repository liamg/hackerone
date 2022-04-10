package hacker

import (
	"fmt"

	"github.com/liamg/hackerone/internal/pkg/api"
)

type getReportResponse struct {
	Report api.Report `json:"data"`
}

type getReportsResponse struct {
	Reports []api.Report `json:"data"`
	Links   api.Links    `json:"links"`
}

type createReportRequest struct {
	Data api.NewReport `json:"data"`
}

type createReportResponse struct {
	Data api.Report `json:"data"`
}

func (a *API) GetReport(id int) (*api.Report, error) {
	var response getReportResponse
	path := fmt.Sprintf("/hackers/reports/%d", id)
	if err := a.client.Get(path, &response); err != nil {
		return nil, err
	}
	return &response.Report, nil
}

func (a *API) GetReports() ([]api.Report, error) {
	var reports []api.Report
	var response getReportsResponse
	path := "/hackers/me/reports"
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

func (a *API) CreateReport(report *api.NewReport) (*api.Report, error) {
	var output createReportResponse
	if err := a.client.Post("", &output, &createReportRequest{
		Data: *report,
	}); err != nil {
		return nil, err
	}
	return &output.Data, nil
}
