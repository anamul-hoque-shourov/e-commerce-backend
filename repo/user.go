package repo

import (
	"database/sql"
	"ecommerce/domain"
	"ecommerce/user"

	"github.com/jmoiron/sqlx"
)

type UserRepo interface {
	user.UserRepo
}

type userRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) UserRepo {
	return &userRepo{
		db: db,
	}
}

func (repo *userRepo) Create(user domain.User) (*domain.User, error) {
	query := `
		INSERT INTO users (
			first_name, 
			last_name, 
			email, 
			password, 
			is_shop_owner
		)
		VALUES (
			:first_name, 
			:last_name, 
			:email, 
			:password, 
			:is_shop_owner
		)
		RETURNING id
	`

	var id int
	rows, err := repo.db.NamedQuery(query, user)

	if err != nil {
		return nil, err
	}

	if rows.Next() {
		rows.Scan(&id)
	}

	user.Id = id

	return &user, nil
}

func (repo *userRepo) Get(email, password string) (*domain.User, error) {
	var user domain.User
	query := `
		SELECT 
			id, 
			first_name, 
			last_name, 
			email, 
			password, 
			is_shop_owner
		FROM users
		WHERE email = $1 AND password = $2
		LIMIT 1
	`
	err := repo.db.Get(&user, query, email, password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
