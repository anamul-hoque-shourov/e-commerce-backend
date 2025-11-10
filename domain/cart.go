package domain

type Cart struct {
	Id     int        `json:"id" db:"id"`
	UserId int        `json:"userId" db:"user_id"`
	Items  []CartItem `json:"items" db:"items"`
}

type CartItem struct {
	Id        int     `json:"id" db:"id"`
	CartId    int     `json:"cartId" db:"cart_id"`
	ProductId int     `json:"productId" db:"product_id"`
	Quantity  int     `json:"quantity" db:"quantity"`
	Price     float64 `json:"price" db:"price"`
}
