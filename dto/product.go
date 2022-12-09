package dto

type CreateProduct struct {
	Title    string  `json:"title"`
	ImageUrl string  `json:"imageUrl"`
	Summary  string  `json:"summary"`
	Price    float32 `json:"price"`
	Quantity int     `json:"quantity"`
	SubId    int     `json:"subCategoryId"`
}
