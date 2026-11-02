package repo

import (
	"database/sql"
	"ecommerce/domain"
	"ecommerce/product"

	"github.com/jmoiron/sqlx"
)

type ProductRepo interface {
	product.ProductRepo
}

type productRepo struct {
	db *sqlx.DB
}

// constractor function
func NewProductRepo(db *sqlx.DB) ProductRepo {
	return &productRepo{
		db: db,
	}
}

func (r *productRepo) Create(p domain.ProductType) (*domain.ProductType, error) {
	// p.ID = len(r.productList) + 1
	// r.productList = append(r.productList, &p)
	// return &p, nil

	query := `
	INSERT INTO products(
	title,
	description,
	price,
	img_url
) VALUES (
	$1,
	$2,
	$3,
	$4
)
	RETURNING id
`
	row := r.db.QueryRow(query, p.Title, p.Description, p.Price, p.ImageUrl)
	err := row.Scan(&p.ID)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *productRepo) Get(id int) (*domain.ProductType, error) {
	var prd domain.ProductType

	query := `
	SELECT 
	id, 
	title, 
	description, 
	price, 
	img_url 
	FROM products
	WHERE id = $1
	`
	err := r.db.Get(&prd, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &prd, nil

	// for _, product := range r.productList {
	// 	if product.ID == productId {
	// 		return product, nil
	// 	}
	// }
	// return nil, nil
}

func (r *productRepo) List(page, limit int64) ([]*domain.ProductType, error) {
	offset := (page - 1) * limit
	var prd []*domain.ProductType
	query := `
	SELECT id, 
	title, 
	description, 
	price, 
	img_url 
	FROM products
	LIMIT $1
	OFFSET $2;
	`
	err := r.db.Select(&prd, query, limit, offset)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return prd, nil
}

func (r *productRepo) Count() (int64, error) {
	query := `
	SELECT COUNT(*) 
	FROM products;
	`
	var count int
	err := r.db.QueryRow(query).Scan(&count)
	if err != nil {
		return 0, err
	}
	return int64(count), nil
}

func (r *productRepo) Update(p domain.ProductType) (*domain.ProductType, error) {

	query := `
	UPDATE products SET
	title= $1,
	description= $2,
	price=$3,
	img_url= $4
    WHERE id = $5;
	`
	row := r.db.QueryRow(query, p.Title, p.Description, p.Price, p.ImageUrl, p.ID)
	err := row.Err()
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *productRepo) Delete(id int) error {
	query := ` 
	DELETE FROM products
    WHERE ID = $1;
	`
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}