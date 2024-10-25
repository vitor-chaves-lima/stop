package repository

import (
	"context"
	"errors"

	"github.com/vitor-chaves-lima/stop/internal/data"
	"github.com/vitor-chaves-lima/stop/internal/data/entities"
)

// CategoryRepository defines the methods for interacting with category data.
type CategoryRepository interface {
	// Count returns the total number of categories in the repository.
	Count(c context.Context) (int, *data.Error)

	// GetAll retrieves all categories with pagination options.
	GetAll(c context.Context, paginationOptions *PaginationOptions) ([]*entities.Category, *data.Error)
}

// PaginationOptions holds pagination parameters for queries.
type PaginationOptions struct {
	Page  int // Page number (starting from 1) for pagination
	Limit int // Maximum number of results to return per page
}

// Validate checks the pagination options for valid values.
func (p *PaginationOptions) Validate() *data.Error {
	if p.Page < 1 {
		return data.NewError(errors.New("page must be greater than or equal to 1"), nil)
	}
	if p.Limit < 1 {
		return data.NewError(errors.New("limit must be greater than 0"), nil)
	}
	return nil
}
