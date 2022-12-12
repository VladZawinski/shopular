package seed

import (
	"context"
	"shopular/ent"
	"shopular/ent/user"
)

func SeedPredefinedUser(c *ent.Client) {
	c.User.Create().
		SetUsername("admin").
		SetPassword("12345").
		SetRole(user.RoleAdmin).
		SaveX(context.Background())
}
