package models

import "github.com/vitor-chaves-lima/stop/internal/data/entities"

type Category struct {
	ID string `json:"id"`
}

func ToCategoryModel(categoryEntity *entities.Category) *Category {
	return &Category{
		ID: categoryEntity.ID,
	}
}

func ToCategoryModels(categoryEntities []*entities.Category) []*Category {
	categoryModels := make([]*Category, len(categoryEntities))
	for i, categoryEntity := range categoryEntities {
		categoryModels[i] = ToCategoryModel(categoryEntity)
	}
	return categoryModels
}
