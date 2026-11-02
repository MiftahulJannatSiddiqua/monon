package domain

//model or entity or existance
//business logic
type ProductType struct {
	ID          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
	Price       int    `json:"price" db:"price"`
	ImageUrl    string `json:"imageUrl" db:"img_url"`
}