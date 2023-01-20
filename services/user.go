package services

import (
	"context"
	"shopular/dto"
	"shopular/ent"
	"shopular/ent/user"
)

type UserService struct {
	client *ent.Client
}

func NewUserService(client ent.Client) UserService {
	return UserService{client: &client}
}

func (s UserService) CreateNewUser(input dto.RegisterUser) *ent.User {
	return s.client.User.
		Create().
		SetUsername(input.Username).
		SetPassword(input.Password).
		SetRole(user.RoleUser).
		SaveX(context.Background())
}

func (s UserService) FindAll() ([]*ent.User, error) {
	return s.client.User.Query().All(context.Background())
}

func (s UserService) FindByUsername(username string) *ent.User {
	u, err := s.client.User.Query().Where(user.Username(username)).First(context.Background())
	if err != nil {
		return nil
	}
	return u
}
