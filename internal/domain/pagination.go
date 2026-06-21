package domain

import "github.com/hamidgh01/Go-Blog-API/internal/domain/entity"

type PaginationQueryParams struct {
	Size int `form:"size"`
	Page int `form:"page"`
}

func (p *PaginationQueryParams) GetOffset() int {
	return (p.GetPage() - 1) * p.GetSize()
}

func (p *PaginationQueryParams) GetSize() int {
	if p.Size <= 0 {
		p.Size = 10
	}

	if p.Size >= 100 {
		p.Size = 100
	}

	return p.Size
}

func (p *PaginationQueryParams) GetPage() int {
	if p.Page <= 0 {
		p.Page = 1
	}
	return p.Page
}

// -----------------------------------------------------------------------------

type PagedList[TEntity entity.TDBEntities] struct {
	Page            int        `json:"page"`
	Size            int        `json:"size"`
	TotalRows       int        `json:"total_rows"`
	TotalPages      int        `json:"total_pages"`
	HasPreviousPage bool       `json:"has_previous_page"`
	PreviousPage    int        `json:"previous_page,omitempty"`
	HasNextPage     bool       `json:"has_next_page"`
	NextPage        int        `json:"next_page,omitempty"`
	Items           []*TEntity `json:"items"`
}

func Paginate[TEntity entity.TDBEntities](items []*TEntity, totalRows int, page int, size int) *PagedList[TEntity] {
	pl := &PagedList[TEntity]{
		Page:      page,
		Size:      size,
		TotalRows: totalRows,
		Items:     items,
	}

	pl.TotalPages = max((totalRows+size-1)/size, 1)
	// e.g. totalRows=30, Size = 10 ->  max(39 / 10, 1) ->  max (3, 1) ->  TotalPages=3
	// e.g. totalRows=31, Size = 10 ->  max(40 / 10, 1) ->  max (4, 1) ->  TotalPages=4
	// e.g. totalRows=8, Size = 10  ->  max(17 / 10, 1) ->  max (1, 1) ->  TotalPages=1
	// e.g. totalRows=0, Size = 10  ->  max(9 / 10, 1)  ->  max (0, 1) ->  TotalPages=0

	pl.HasPreviousPage = pl.Page > 1
	if pl.HasPreviousPage {
		pl.PreviousPage = pl.Page - 1
	}

	pl.HasNextPage = pl.Page < pl.TotalPages
	if pl.HasPreviousPage {
		pl.PreviousPage = pl.Page + 1
	}

	return pl
}
