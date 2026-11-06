package repo

import (
	"database/sql"
	"ecommerce/domain"
	"ecommerce/user"
	"fmt"

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

func (r *userRepo) Create(user domain.User) (*domain.User, error) {
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
			:is_shop_owner)
		RETURNING id
	`
	var userID int
	rows, err := r.db.NamedQuery(query, user)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	if rows.Next() {
		rows.Scan(&userID)
	}
	user.ID = userID

	return &user, nil
}

func (r *userRepo) Find(email, pass string) (*domain.User, error) {
	var user domain.User

	query := `
		SELECT id, 
		first_name, last_name, email, password, is_shop_owner
		FROM users
		WHERE email = $1 AND password = $2
	`

	err := r.db.Get(&user, query, email, pass)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil //no matching user
		}
		return nil, err
	}

	return &user, nil
}

//===================
func (r *userRepo) GetProfile(id int) (*domain.Profile, error) {
	var p domain.Profile

	query := `
		SELECT id, first_name, last_name, email, bio, city, personality_type
		FROM users
		WHERE id = $1
	`

	err := r.db.Get(&p, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &p, nil
}

func (r *userRepo) UpdateProfile(id int, bio, city *string) error {
	query := `
		UPDATE users
		SET bio = COALESCE($2, bio),
		    city = COALESCE($3, city),
		    updated_at = CURRENT_TIMESTAMP
		WHERE id = $1
	`

	_, err := r.db.Exec(query, id, bio, city)
	return err
}