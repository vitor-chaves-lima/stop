package services

import (
	"context"

	"github.com/vitor-chaves-lima/stop/internal/data/repository"
	"github.com/vitor-chaves-lima/stop/internal/logic"
	"github.com/vitor-chaves-lima/stop/internal/logic/models"
)

type CategoryServiceInterface interface {
	ListCategories(c context.Context, options logic.PaginationOptions) ([]*models.Category, *logic.PaginationInfo, *logic.Error)
}

type CategoryService struct {
	categoryRepository repository.Category
}

func NewCategoryService(categoryRepository repository.Category) *CategoryService {
	return &CategoryService{
		categoryRepository: categoryRepository,
	}
}

func (s *CategoryService) ListCategories(c context.Context, options logic.PaginationOptions) ([]*models.Category, *logic.PaginationInfo, *logic.Error) {
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

	categories := models.ToCategoryModels(categoryEntities)
	return categories, logic.NewPaginationInfo(totalCount, options.Page, options.Limit), nil
}
