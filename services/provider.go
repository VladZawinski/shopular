package services

import "shopular/ent"

type Provider struct {
	Category    CategoryService
	SubCategory SubCategoryService
	Product     ProductService
	User        UserService
	Cart        CartService
	Customer    CustomerService
}

func NewServiceProvider(client *ent.Client) Provider {
	return Provider{
		Category:    NewCategoryService(client),
		SubCategory: NewSubCategoryService(client),
		Product:     NewProductService(client),
		User:        NewUserService(*client),
		Cart:        NewCartService(client),
		Customer:    NewCustomerService(client),
	}
}
