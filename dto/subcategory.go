package dto

type CreateSubCategory struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	CategoryId  int    `json:"categoryId"`
}
