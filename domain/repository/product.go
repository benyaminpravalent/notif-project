package repository

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/project/notif-project/domain/model"
	"github.com/project/notif-project/pkg/database"
)

// ProductRepository manages database operations for product.
type ProductRepository interface {
	Create(product *model.Product) error
	GetBySKU(sku string) (*model.Product, error)
	GetByID(id int64) (*model.Product, error)
	GetByBrandID(brandID int64) ([]*model.Product, error)
}

type productRepoImpl struct {
	db *sqlx.DB
}

// NewProductRepository returns new instance of productRepoImpl
func NewProductRepository() *productRepoImpl {
	return &productRepoImpl{
		db: database.DB,
	}
}

func (r *productRepoImpl) scanRows(rows *sql.Rows) (items []*model.Product, err error) {
	items = make([]*model.Product, 0)
	for rows.Next() {
		res := &model.Product{}
		err = rows.Scan(&res.ID, &res.SKU, &res.BrandID, &res.Stock, &res.Price, &res.CreatedAt)
		if err != nil {
			return
		}
		items = append(items, res)
	}
	return
}

// Create creates a new product into the database.
func (r *productRepoImpl) Create(product *model.Product) error {
	res, err := r.db.Exec(`
		INSERT INTO product (sku, brand_id, stock, price)
		VALUES (?, ?, ?, ?)`, product.SKU, product.BrandID, product.Stock, product.Price)

	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	product.ID = id

	return err
}

// GetBySKU returns product's details by SKU.
func (r *productRepoImpl) GetBySKU(sku string) (*model.Product, error) {
	res := &model.Product{}
	err := r.db.Get(res, `
		SELECT *
		FROM product
		WHERE sku = ?`, sku)
	if err == sql.ErrNoRows {
		res = nil
		err = nil
	}
	return res, err
}

// GetByID returns product's details by ID.
func (r *productRepoImpl) GetByID(id int64) (*model.Product, error) {
	res := &model.Product{}
	err := r.db.Get(res, `
		SELECT * from skripsi.resource r where r.is_active = 1::bit and r.resource_id = 6`)
	if err == sql.ErrNoRows {
		res = nil
		err = nil
	}
	return res, err
}

// GetByBrandID returns produt's details by brand ID.
func (r *productRepoImpl) GetByBrandID(brandID int64) ([]*model.Product, error) {
	res, err := r.db.Query(`
		SELECT *
		FROM product
		WHERE brand_id = ?`, brandID)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	items, err := r.scanRows(res)
	return items, err
}
