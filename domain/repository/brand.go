package repository

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/project/notif-project/domain/model"
	"github.com/project/notif-project/pkg/database"
)

// BrandRepository manages database operations for brand.
type BrandRepository interface {
	Create(brand string) (int64, error)
	GetByID(id int64) (*model.Brand, error)
}

type brandRepoImpl struct {
	db *sqlx.DB
}

// NewBrandRepository returns new instance of brandRepoImpl.
func NewBrandRepository() *brandRepoImpl {
	return &brandRepoImpl{
		db: database.DB,
	}
}

// Create creates a new brand into the database.
func (r *brandRepoImpl) Create(brand string) (int64, error) {
	res, err := r.db.Exec(`
		INSERT INTO brand (name)
		VALUES (?)`, brand)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	return id, err
}

// GetByID returns a brand's details by ID.
func (r *brandRepoImpl) GetByID(id int64) (*model.Brand, error) {
	res := &model.Brand{}
	err := r.db.Get(res, `
		SELECT *
		FROM brand
		WHERE id = ?`, id)
	if err == sql.ErrNoRows {
		res = nil
		err = nil
	}
	return res, err
}
