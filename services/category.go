package services

import (
	"context"
	"shopular/dto"
	"shopular/ent"
)

type CategoryService struct {
	db *ent.Client
}

func NewCategoryService(client *ent.Client) CategoryService {
	return CategoryService{
		db: client,
	}
}

func (s *CategoryService) GetAll() ([]*ent.Category, error) {
	return s.db.Category.Query().All(context.Background())
}

func (s *CategoryService) Create(input *dto.CreateCategory) *ent.Category {
	return s.db.Category.
		Create().
		SetTitle(input.Name).
		SetDescription(input.Description).
		SaveX(context.Background())
}
