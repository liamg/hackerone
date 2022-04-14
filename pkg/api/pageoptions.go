package api

// PageOptions allows the consumer to request a specific page number and/or size.
type PageOptions struct {
	// PageNumber is the page number of results to return. Starts at (and defaults to) 1.
	PageNumber int
	// PageSize is the number of records to return per page. 1-100 allowed. Defaults to 100.
	PageSize int
}

// GetPageNumber returns the input page number with corrections applied
func (p *PageOptions) GetPageNumber() int {
	if p == nil || p.PageNumber <= 0 {
		return 1
	}
	return p.PageNumber
}

// GetPageSize returns the input page size with corrections applied
func (p *PageOptions) GetPageSize() int {
	if p == nil || p.PageSize <= 0 || p.PageSize > 100 {
		return 100
	}
	return p.PageSize
}
