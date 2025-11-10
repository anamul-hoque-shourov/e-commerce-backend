package repo

import (
	"database/sql"
	"ecommerce/cart"
	"ecommerce/domain"
	"errors"

	"github.com/jmoiron/sqlx"
)

type CartRepo interface {
	cart.CartRepo
}

type cartRepo struct {
	db *sqlx.DB
}

func NewCartRepo(db *sqlx.DB) CartRepo {
	return &cartRepo{
		db: db,
	}
}

func (r *cartRepo) GetByUserId(userId int) (*domain.Cart, error) {
	var cart domain.Cart
	query := `
		SELECT *
		FROM carts 
		WHERE user_id = $1
	`
	err := r.db.Get(&cart, query, userId)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	var items []domain.CartItem
	query = `
		SELECT * 
		FROM cart_items 
		WHERE cart_id = $1
	`
	err = r.db.Select(&items, query, cart.Id)
	if err != nil {
		return nil, err
	}

	cart.Items = items
	return &cart, nil
}

func (r *cartRepo) AddItem(userId, productId, quantity int) error {
	var cartID int
	err := r.db.Get(&cartID, `SELECT id FROM carts WHERE user_id=$1`, userId)
	if err == sql.ErrNoRows {
		err = r.db.Get(&cartID, `INSERT INTO carts (user_id) VALUES ($1) RETURNING id`, userId)
		if err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	// Check if item exists
	var exists bool
	err = r.db.Get(&exists, `SELECT EXISTS(SELECT 1 FROM cart_items WHERE cart_id=$1 AND product_id=$2)`, cartID, productId)
	if err != nil {
		return err
	}

	if exists {
		_, err = r.db.Exec(`
			UPDATE cart_items
			SET quantity = quantity + $1, updated_at = NOW()
			WHERE cart_id = $2 AND product_id = $3
		`, quantity, cartID, productId)
	} else {
		// Get product price
		var price float64
		err = r.db.Get(&price, `SELECT price FROM products WHERE id = $1`, productId)
		if err != nil {
			return err
		}

		_, err = r.db.Exec(`
			INSERT INTO cart_items (cart_id, product_id, quantity, price)
			VALUES ($1, $2, $3, $4)
		`, cartID, productId, quantity, price)
	}

	if err != nil {
		return err
	}

	return nil
}

func (r *cartRepo) UpdateItemQuantity(userId, productId, quantity int) error {
	query := `
		UPDATE cart_items
		SET quantity = $1, updated_at = NOW()
		WHERE cart_id = (SELECT id FROM carts WHERE user_id = $2)
		AND product_id = $3
	`
	_, err := r.db.Exec(query, quantity, userId, productId)
	return err
}

func (r *cartRepo) RemoveItem(userId, productId int) error {
	query := `
		DELETE FROM cart_items
		WHERE cart_id = (SELECT id FROM carts WHERE user_id = $1)
		AND product_id = $2
	`
	_, err := r.db.Exec(query, userId, productId)
	return err
}

func (r *cartRepo) ClearCart(userId int) error {
	query := `
		DELETE FROM cart_items
		WHERE cart_id = (SELECT id FROM carts WHERE user_id = $1)
	`
	_, err := r.db.Exec(query, userId)
	return err
}
