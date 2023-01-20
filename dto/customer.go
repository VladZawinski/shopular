package dto

type CreateCustomer struct {
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Address   *string `json:"address,omitempty"`
	Phone     *string `json:"phone,omitempty"`
	Email     *string `json:"email,omitempty"`
}
