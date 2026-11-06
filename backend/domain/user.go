package domain

type User struct {
	ID          int    `json:"id" db:"id"`
	FirstName   string `json:"first_name" db:"first_name"`
	LastName    string `json:"last_name" db:"last_name"`
	Email       string `json:"email" db:"email"`
	Password    string `json:"password" db:"password"`
	IsShopOwner bool   `json:"is_shop_owner" db:"is_shop_owner"`
}
//=================
type Profile struct {
	ID              int     `json:"id" db:"id"`
	FirstName       string  `json:"first_name" db:"first_name"`
	LastName        string  `json:"last_name" db:"last_name"`
	Email           string  `json:"email" db:"email"`
	Bio             *string `json:"bio" db:"bio"`
	City            *string `json:"city" db:"city"`
	PersonalityType *string `json:"personality_type" db:"personality_type"`
}