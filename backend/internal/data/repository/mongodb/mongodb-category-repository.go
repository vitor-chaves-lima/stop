package mongodb

import (
	"context"

	"github.com/vitor-chaves-lima/stop/internal/data"
	"github.com/vitor-chaves-lima/stop/internal/data/entities"
	"github.com/vitor-chaves-lima/stop/internal/data/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

type CategoryRepository struct {
	session    *mongo.Session
	collection *mongo.Collection
}

func NewCategoryRepository(session *mongo.Session, collection *mongo.Collection) *CategoryRepository {
	return &CategoryRepository{session: session, collection: collection}
}

func (r *CategoryRepository) Count(c context.Context) (int, *data.Error) {
	documents, err := r.collection.CountDocuments(c, nil)
	if err != nil {
		return 0, data.NewError(err, nil)
	}

	return int(documents), nil
}

func (r *CategoryRepository) GetAll(c context.Context, paginationOptions *repository.PaginationOptions) ([]*entities.Category, *data.Error) {
	if err := paginationOptions.Validate(); err != nil {
		return nil, err
	}

	cursor, err := r.collection.Find(c, nil)
	if err != nil {
		return nil, data.NewError(err, nil)
	}

	var categories []*entities.Category
	if err := cursor.All(c, &categories); err != nil {
		return nil, data.NewError(err, nil)
	}

	return categories, nil
}
