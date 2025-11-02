package model

type PaginationParam struct {
	Offset int
	Limit  int
}

type Pagination struct {
	Offset          int   `json:"offset" example:"1"`              // The current page
	Limit           int   `json:"limit" example:"10"`              // The size of the page
	TotalPages      int64 `json:"total_pages" example:"5"`         // The total number of pages
	TotalRowPerPage int64 `json:"total_row_per_page" example:"10"` // The total number of data per page
	TotalRows       int64 `json:"total_rows" example:"50"`         // The total number of data
}

type PaginationData[T any] struct {
	Offset          int   `json:"offset"`             // The current page
	Limit           int   `json:"limit"`              // The size of the page
	TotalPages      int64 `json:"total_pages"`        // The total number of pages
	TotalRowPerPage int64 `json:"total_row_per_page"` // The total number of data per page
	TotalRows       int64 `json:"total_rows"`         // The total number of data
	Data            []*T  `json:"data"`               // The actual data
}

func (t *PaginationData[T]) ToPagination() *Pagination {
	return &Pagination{
		Offset:          t.Offset,
		Limit:           t.Limit,
		TotalPages:      t.TotalPages,
		TotalRowPerPage: t.TotalRowPerPage,
		TotalRows:       t.TotalRows,
	}
}
