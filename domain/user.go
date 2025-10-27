package domain

type User struct {
	ID          int    `json:"id" db:"id"`
	FirstName   string `json:"firstName" db:"first_name"`
	LastName    string `json:"lastName" db:"last_name"`
	Email       string `json:"email" db:"email"`
	Password    string `json:"password" db:"password"`
	IsShopOwner bool   `json:"isShopOwner" db:"is_shop_owner"`
}
