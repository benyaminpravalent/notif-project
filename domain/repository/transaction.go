package repository

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/richardsahvic/jamtangan/domain/model"
	"github.com/richardsahvic/jamtangan/pkg/database"
)

// TransactionRepository manages database operations for transaction.
type TransactionRepository interface {
	InsertList(order []model.Transaction) error
	GetDetail(orderID string) ([]*model.Transaction, error)
}

type transactionRepoImpl struct {
	db *sqlx.DB
}

// NewTransactionRepository returns new instance of transactionRepoImpl
func NewTransactionRepository() *transactionRepoImpl {
	return &transactionRepoImpl{
		db: database.DB,
	}
}

func (r *transactionRepoImpl) scanRows(rows *sql.Rows) (items []*model.Transaction, err error) {
	items = make([]*model.Transaction, 0)
	for rows.Next() {
		res := &model.Transaction{}
		err = rows.Scan(&res.ID, &res.SKU, &res.Quantity, &res.OrderID,
			&res.CreatedAt, &res.UpdatedAt, &res.DeletedAt, &res.Subtotal)
		if err != nil {
			return
		}
		items = append(items, res)
	}
	return
}

// InsertList inserts new list of transaction.
func (r *transactionRepoImpl) InsertList(transaction []model.Transaction) error {
	inserts := make([]string, len(transaction))
	params := make([]interface{}, 0, 4*len(transaction))

	for index, item := range transaction {
		values := make([]string, 0, 4)

		values = append(values, "?", "?", "?", "?")
		params = append(params, item.SKU, item.Quantity, item.OrderID, item.Subtotal)

		inserts[index] = fmt.Sprintf("(%s)", strings.Join(values, ", "))
	}

	_, err := r.db.Exec(fmt.Sprintf(`
		INSERT INTO transaction (
			sku, quantity, order_id, subtotal
		)
		VALUES %s`, strings.Join(inserts, ", ")), params...)
	return err
}

// GetDetail returns transaction's details by order ID.
func (r *transactionRepoImpl) GetDetail(orderID string) ([]*model.Transaction, error) {
	res, err := r.db.Query(`
		SELECT *
		FROM transaction
		WHERE order_id = ?`, orderID)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	items, err := r.scanRows(res)
	return items, err
}
