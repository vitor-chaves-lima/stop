package services

import (
	"context"
	"errors"

	"github.com/vitor-chaves-lima/stop/internal/data/repository"
	"github.com/vitor-chaves-lima/stop/internal/logic"
	"github.com/vitor-chaves-lima/stop/internal/logic/models"
)

type CategoryServiceInterface interface {
	ListCategories(c context.Context, options PaginationOptions) ([]*models.Category, *logic.PaginationInfo, *logic.Error)
}

type CategoryService struct {
	categoryRepository repository.Category
}

func NewCategoryService(categoryRepository repository.Category) *CategoryService {
	return &CategoryService{
		categoryRepository: categoryRepository,
	}
}

func (s *CategoryService) ListCategories(c context.Context, options PaginationOptions) ([]*models.Category, *logic.PaginationInfo, *logic.Error) {
	totalCount, err := s.categoryRepository.Count(c)
	if err != nil {
		return nil, nil, logic.NewError("ServiceError", err)
	}

	if totalCount == 0 {
		return []*models.Category{}, logic.NewPaginationInfo(totalCount, options.Page, options.Limit), nil
	}

	categoryEntities, err := s.categoryRepository.GetAll(c, options.ToDataPaginationOptions())
	if err != nil {
		return nil, nil, logic.NewError("ServiceError", err)
	}

	categories := make([]*models.Category, len(categoryEntities))
	for i, categoryEntity := range categoryEntities {
		categories[i] = models.ToCategoryModel(categoryEntity)
	}
	return categories, nil, nil
}

// PaginationOptions holds pagination parameters for queries in the service layer.
type PaginationOptions struct {
	Page  int // Page number (starting from 1) for pagination
	Limit int // Maximum number of results to return per page
}

// Validate checks the pagination options for valid values in the service layer.
func (p *PaginationOptions) Validate() *logic.Error {
	if p.Page < 1 {
		return logic.NewError("OptionsError", errors.New("page must be greater than or equal to 1"))
	}
	if p.Limit < 1 {
		return logic.NewError("OptionsError", errors.New("limit must be greater than 0"))
	}
	return nil
}

func (p *PaginationOptions) ToDataPaginationOptions() *repository.PaginationOptions {
	return &repository.PaginationOptions{
		Page:  p.Page,
		Limit: p.Limit,
	}
}
