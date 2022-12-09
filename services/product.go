package services

import (
	"context"
	"shopular/dto"
	"shopular/ent"
)

type ProductService struct {
	db *ent.Client
}

func NewProductService(client *ent.Client) ProductService {
	return ProductService{db: client}
}

func (s *ProductService) GetAll() ([]*ent.Product, error) {
	return s.db.Product.Query().WithSub().All(context.Background())
}

func (s *ProductService) Create(input *dto.CreateProduct) *ent.Product {
	return s.db.Product.
		Create().
		SetTitle(input.Title).
		SetSummary(input.Summary).
		SetImageUrl(input.ImageUrl).
		SetQuantity(input.Quantity).
		SetPrice(float64(input.Price)).
		AddSubIDs(input.SubId).
		SaveX(context.Background())
}
