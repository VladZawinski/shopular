package services

import (
	"context"
	"shopular/dto"
	"shopular/ent"
)

type SubCategoryService struct {
	db *ent.Client
}

func NewSubCategoryService(client *ent.Client) SubCategoryService {
	return SubCategoryService{
		db: client,
	}
}

func (s *SubCategoryService) GetAll() ([]*ent.SubCategory, error) {
	return s.db.SubCategory.Query().All(context.Background())
}

func (s *SubCategoryService) Create(input *dto.CreateSubCategory) *ent.SubCategory {
	return s.db.SubCategory.
		Create().
		SetTitle(input.Title).
		SetDescription(input.Description).
		SetCategoryID(input.CategoryId).
		SaveX(context.Background())
}
