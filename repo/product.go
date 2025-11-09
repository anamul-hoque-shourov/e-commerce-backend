package repo

import (
	"database/sql"
	"ecommerce/domain"
	"ecommerce/product"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type ProductRepo interface {
	product.ProductRepo
}

type productRepo struct {
	db *sqlx.DB
}

func NewProductRepo(db *sqlx.DB) ProductRepo {
	return &productRepo{
		db: db,
	}
}

func (r *productRepo) Create(product domain.Product) (*domain.Product, error) {
	query := `
		INSERT INTO products (
			title,
			description,
			price,
			image_url
		) VALUES (
			$1,
			$2,
			$3,
			$4
		)
		RETURNING id
	`
	row := r.db.QueryRow(query, product.Title, product.Description, product.Price, product.ImageUrl)
	err := row.Scan(&product.ID)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *productRepo) Get(productId int) (*domain.Product, error) {
	query := `
		SELECT 
			id,
			title,
			description,
			price,
			image_url
		FROM products
		WHERE id = $1
	`
	var product domain.Product
	err := r.db.Get(&product, query, productId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &product, nil
}

func (r *productRepo) List(page int, limit int) ([]*domain.Product, error) {
	offset := ((page - 1) * limit)
	query := `
		SELECT 
			id,
			title,
			description,
			price,
			image_url
		FROM products
		ORDER BY id
		LIMIT $1
		OFFSET $2
	`
	var products []*domain.Product
	err := r.db.Select(&products, query, limit, offset)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return products, nil
}

func (r *productRepo) Update(product domain.Product) (*domain.Product, error) {
	query := `
		UPDATE products
		SET
			title = $1,
			description = $2,
			price = $3,
			image_url = $4
		WHERE id = $5
		RETURNING id, title, description, price, image_url
	`
	var updatedProduct domain.Product
	err := r.db.Get(&updatedProduct, query, product.Title, product.Description, product.Price, product.ImageUrl, product.ID)
	if err != nil {
		return nil, err
	}
	return &updatedProduct, nil
}

func (r *productRepo) Delete(productId int) error {
	query := `
		DELETE FROM products
		WHERE id = $1
	`
	result, err := r.db.Exec(query, productId)
	if err != nil {
		return err
	}
	_, err = result.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}
