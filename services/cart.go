package services

import (
	"context"
	"shopular/ent"
	"shopular/ent/cart"
	"shopular/ent/product"
	"shopular/ent/user"

	"github.com/gofiber/fiber/v2"
)

type CartService struct {
	client *ent.Client
}

func NewCartService(client *ent.Client) CartService {
	return CartService{client: client}
}

func (s CartService) FindByUserId(userId int) ([]*ent.Cart, error) {
	return s.client.Cart.
		Query().
		WithProduct().
		Where(cart.HasUserWith(user.ID(userId))).
		All(context.Background())
}

func (s CartService) RemoveFromCart(cartId int) error {
	existing := s.client.Cart.
		Query().
		Where(cart.ID(cartId)).
		FirstX(context.Background())
	if existing == nil {
		return fiber.NewError(fiber.StatusNotFound)
	}
	if existing.Quantity > 1 {
		s.client.Cart.Update().
			SetQuantity(existing.Quantity - 1).
			ExecX(context.Background())
		return nil
	}
	s.client.Cart.DeleteOneID(existing.ID).ExecX(context.Background())
	return nil
}

func (s CartService) AddToCart(userId int, productId int) error {
	existing := s.client.Cart.
		Query().
		Where(
			cart.And(cart.HasProductWith(product.ID(productId)), cart.HasUserWith(user.ID(userId))),
		).FirstX(context.Background())
	if existing == nil {
		s.client.Cart.
			Create().
			SetQuantity(1).
			AddProductIDs(productId).
			AddUserIDs(userId).
			ExecX(context.Background())
		return nil
	}
	next := (existing.Quantity + 1)
	s.client.Cart.Update().
		Where(cart.ID(existing.ID)).
		SetQuantity(next).
		SaveX(context.Background())
	return nil
}
