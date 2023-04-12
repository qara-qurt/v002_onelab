package postgres

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"v002_onelab/internal/model"
)

type ITransationRepository interface {
	CreateTransaction(tran model.TransactionResponse) error
	GetAll(userID uint) ([]model.Transaction, error)
	GetByID(id uint) (model.Transaction, error)
}

type Transaction struct {
	db *sqlx.DB
}

func NewTransaction(db *sqlx.DB) *Transaction {
	return &Transaction{
		db: db,
	}
}

func (t Transaction) CreateTransaction(tran model.TransactionResponse) error {
	query, err := t.db.Preparex("INSERT INTO transactionBook(user_id,book_id,price) VALUES ($1,$2,$3)")
	if err != nil {
		return err
	}

	if _, err := query.Exec(tran.UserID, tran.BookID, tran.Price); err != nil {
		return err
	}
	return nil
}
func (t Transaction) GetAll(userID uint) ([]model.Transaction, error) {
	var buyBooks []model.Transaction
	query := "SELECT id, user_id, book_id, price FROM transactionBook WHERE user_id = $1"
	if err := t.db.Select(&buyBooks, query, userID); err != nil {
		return []model.Transaction{}, err
	}
	return buyBooks, nil
}

func (t Transaction) GetByID(id uint) (model.Transaction, error) {
	var buyBook model.Transaction
	query := "SELECT id, user_id, book_id, price FROM transactionBook WHERE id = $1"
	if err := t.db.Get(&buyBook, query, id); err != nil {
		if err == sql.ErrNoRows {
			// Return a custom error if there are no rows with the given ID
			return model.Transaction{}, model.ErrorNotFound
		}
		// Return the original error if it's not sql.ErrNoRows
		return model.Transaction{}, err
	}
	return buyBook, nil
}
