package hackers

import (
	"context"
	"fmt"

	api2 "github.com/liamg/hackerone/pkg/api"
)

type getProgramsResponse struct {
	Data  []api2.Program `json:"data"`
	Links api2.Links     `json:"links"`
}

type getProgramResponse struct {
	Data api2.Program `json:"data"`
}

type getProgramWeaknessesResponse struct {
	Data  []api2.Weakness `json:"data"`
	Links api2.Links      `json:"links"`
}

// GetPrograms returns a list of programs received by the hacker. If there are further pages, nextPage will be >0.
func (a *API) GetPrograms(ctx context.Context, pageOptions *api2.PageOptions) (programs []api2.Program, nextPage int, err error) {
	var response getProgramsResponse
	path := fmt.Sprintf(
		"/hackers/programs?page[number]=%d&page[size]=%d",
		pageOptions.GetPageNumber(),
		pageOptions.GetPageSize(),
	)
	if err := a.client.Get(ctx, path, &response); err != nil {
		return nil, 0, err
	}
	if response.Links.Next != "" {
		nextPage = pageOptions.GetPageNumber() + 1
	}
	return response.Data, nextPage, nil
}

// GetProgram returns the program for the given id.
func (a *API) GetProgram(ctx context.Context, handle string) (*api2.Program, error) {
	var response getProgramResponse
	path := fmt.Sprintf("/hackers/programs/%s", handle)
	if err := a.client.Get(ctx, path, &response); err != nil {
		return nil, err
	}
	return &response.Data, nil
}

// GetProgramWeaknesses returns a list of weaknesses for a given program. If there are further pages, nextPage will be >0.
func (a *API) GetProgramWeaknesses(ctx context.Context, handle string, pageOptions *api2.PageOptions) (programs []api2.Weakness, nextPage int, err error) {
	var response getProgramWeaknessesResponse
	path := fmt.Sprintf(
		"/hackers/programs/%s/weaknesses?page[number]=%d&page[size]=%d",
		handle,
		pageOptions.GetPageNumber(),
		pageOptions.GetPageSize(),
	)
	if err := a.client.Get(ctx, path, &response); err != nil {
		return nil, 0, err
	}
	if response.Links.Next != "" {
		nextPage = pageOptions.GetPageNumber() + 1
	}
	return response.Data, nextPage, nil
}
