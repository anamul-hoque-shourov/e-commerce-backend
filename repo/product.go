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

func (repo *productRepo) Create(product domain.Product) (*domain.Product, error) {
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
	row := repo.db.QueryRow(query, product.Title, product.Description, product.Price, product.ImageUrl)
	err := row.Scan(&product.Id)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (repo *productRepo) Get(productId int) (*domain.Product, error) {
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
	err := repo.db.Get(&product, query, productId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &product, nil
}

func (repo *productRepo) List(page, limit int) ([]*domain.Product, error) {
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
	err := repo.db.Select(&products, query, limit, offset)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return products, nil
}

func (repo *productRepo) Count() (int, error) {
	query := `
		SELECT
			COUNT(*)
		FROM products
	`
	var productsCount int
	err := repo.db.Get(&productsCount, query)
	if err != nil {
		return 0, err
	}
	return productsCount, nil
}

func (repo *productRepo) Update(product domain.Product) (*domain.Product, error) {
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
	err := repo.db.Get(&updatedProduct, query, product.Title, product.Description, product.Price, product.ImageUrl, product.Id)
	if err != nil {
		return nil, err
	}
	return &updatedProduct, nil
}

func (repo *productRepo) Delete(productId int) error {
	query := `
		DELETE FROM products
		WHERE id = $1
	`
	result, err := repo.db.Exec(query, productId)
	if err != nil {
		return err
	}
	_, err = result.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}
