package hackers

import (
	"context"
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
	Data api.CreateReportInput `json:"data"`
}

type createReportResponse struct {
	Data api.Report `json:"data"`
}

// GetReport returns the report for the given id.
func (a *API) GetReport(ctx context.Context, id int) (*api.Report, error) {
	var response getReportResponse
	path := fmt.Sprintf("/hackers/reports/%d", id)
	if err := a.client.Get(ctx, path, &response); err != nil {
		return nil, err
	}
	return &response.Report, nil
}

// GetReports returns a list of reports made by the hacker. If there are further pages, nextPage will be >0.
func (a *API) GetReports(ctx context.Context, pageOptions *api.PageOptions) (reports []api.Report, nextPage int, err error) {
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
func (a *API) CreateReport(ctx context.Context, report *api.CreateReportInput) (*api.Report, error) {
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
