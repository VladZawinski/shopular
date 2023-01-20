package services

import (
	"context"
	"shopular/dto"
	"shopular/ent"
)

type CustomerService struct {
	client *ent.Client
}

func NewCustomerService(client *ent.Client) CustomerService {
	return CustomerService{client: client}
}

func (cs CustomerService) Create(body dto.CreateCustomer) (*ent.Customer, error) {
	return cs.client.Customer.
		Create().
		SetFirstName(body.FirstName).
		SetLastName(body.LastName).
		SetAddress(body.Address).
		SetPhone(body.Phone).
		SetEmail(body.Email).
		Save(context.Background())
}

func (cs CustomerService) FindAll() ([]*ent.Customer, error) {
	return cs.client.Customer.Query().All(context.Background())
}

func (cs CustomerService) DeleteOne(id int) error {
	return cs.client.Customer.DeleteOneID(id).Exec(context.Background())
}

func UpdateOne() {

}
