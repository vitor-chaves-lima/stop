package logic

import (
	"errors"

	"github.com/vitor-chaves-lima/stop/internal/data/repository"
)

// PaginationInfo holds metadata about the pagination state.
type PaginationInfo struct {
	TotalCount   int // Total number of items available
	PageCount    int // Total number of pages available
	CurrentPage  int // The current page number
	ItemsPerPage int // Number of items per page
}

// NewPaginationInfo creates and returns a new PaginationInfo structure.
// It calculates the PageCount based on the TotalCount and ItemsPerPage.
func NewPaginationInfo(totalCount, currentPage, itemsPerPage int) *PaginationInfo {
	pageCount := (totalCount + itemsPerPage - 1) / itemsPerPage

	return &PaginationInfo{
		TotalCount:   totalCount,
		PageCount:    pageCount,
		CurrentPage:  currentPage,
		ItemsPerPage: itemsPerPage,
	}
}

// PaginationOptions holds pagination parameters for queries in the service layer.
type PaginationOptions struct {
	Page  int // Page number (starting from 1) for pagination
	Limit int // Maximum number of results to return per page
}

// Validate checks the pagination options for valid values in the service layer.
func (p *PaginationOptions) Validate() *Error {
	if p.Page < 1 {
		return NewError("OptionsError", errors.New("page must be greater than or equal to 1"))
	}
	if p.Limit < 1 {
		return NewError("OptionsError", errors.New("limit must be greater than 0"))
	}
	return nil
}

func (p *PaginationOptions) ToDataPaginationOptions() *repository.PaginationOptions {
	return &repository.PaginationOptions{
		Page:  p.Page,
		Limit: p.Limit,
	}
}
