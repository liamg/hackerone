package hackers

import (
	"context"
	"fmt"

	api2 "github.com/liamg/hackerone/pkg/api"
)

type getReportResponse struct {
	Report api2.Report `json:"data"`
}

type getReportsResponse struct {
	Reports []api2.Report `json:"data"`
	Links   api2.Links    `json:"links"`
}

type createReportRequest struct {
	Data api2.CreateReportInput `json:"data"`
}

type createReportResponse struct {
	Data api2.Report `json:"data"`
}

// GetReport returns the report for the given id.
func (a *API) GetReport(ctx context.Context, id int) (*api2.Report, error) {
	var response getReportResponse
	path := fmt.Sprintf("/hackers/reports/%d", id)
	if err := a.client.Get(ctx, path, &response); err != nil {
		return nil, err
	}
	return &response.Report, nil
}

// GetReports returns a list of reports made by the hacker. If there are further pages, nextPage will be >0.
func (a *API) GetReports(ctx context.Context, pageOptions *api2.PageOptions) (reports []api2.Report, nextPage int, err error) {
	var response getReportsResponse
	path := fmt.Sprintf(
		"/hackers/me/reports?page[number]=%d&page[size]=%d",
		pageOptions.GetPageNumber(),
		pageOptions.GetPageSize(),
	)
	if err := a.client.Get(ctx, path, &response); err != nil {
		return nil, 0, err
	}
	if response.Links.Next != "" {
		nextPage = pageOptions.GetPageNumber() + 1
	}
	return response.Reports, nextPage, nil
}

// CreateReport creates a new report. See https://api.hackerone.com/hacker-resources/#reports-create-report
func (a *API) CreateReport(ctx context.Context, report *api2.CreateReportInput) (*api2.Report, error) {
	if report == nil {
		return nil, fmt.Errorf("report input cannot be nil")
	}
	var output createReportResponse
	if err := a.client.Post(ctx, "/hackers/reports", &output, &createReportRequest{
		Data: *report,
	}); err != nil {
		return nil, err
	}
	return &output.Data, nil
}
