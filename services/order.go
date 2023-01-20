package services

import "shopular/ent"

type OrderService struct {
	client *ent.Client
}

func NewOrderService(client *ent.Client) CartService {
	return CartService{client: client}
}

func (os OrderService) CheckOut() {

}
