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

func Paginate[TEntity entity.TDBEntities](items []*TEntity, totalRows, page, size, totalPages int) *PagedList[TEntity] {
	pl := &PagedList[TEntity]{
		Page:       page,
		Size:       size,
		TotalRows:  totalRows,
		TotalPages: totalPages,
		Items:      items,
	}

	if pl.HasPreviousPage = pl.Page > 1; pl.HasPreviousPage {
		pl.PreviousPage = pl.Page - 1
	}

	if pl.HasNextPage = pl.Page < pl.TotalPages; pl.HasNextPage {
		pl.NextPage = pl.Page + 1
	}

	return pl
}
