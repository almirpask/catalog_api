package service

import (
	"github.com/almirpask/go_api/internal/database"
	"github.com/almirpask/go_api/internal/entity"
)

type CategoryService struct {
	CategoryDb database.CategoryDB
}

func NewCategoryService(categoryDb database.CategoryDB) *CategoryService {
	return &CategoryService{
		CategoryDb: categoryDb,
	}
}

func (cs *CategoryService) GetCategories() ([]*entity.Category, error) {
	categories, err := cs.CategoryDb.GetCategories()

	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (cs *CategoryService) CreateCategory(name string) (*entity.Category, error) {
	category := entity.NewCategory(name)

	_, err := cs.CategoryDb.CreateCategory(category)

	if err != nil {
		return nil, err
	}

	return category, nil
}

func (cs *CategoryService) GetCategory(id string) (*entity.Category, error) {
	category, err := cs.CategoryDb.GetCategory(id)

	if err != nil {
		return nil, err
	}

	return category, nil
}
