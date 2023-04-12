package service

import (
	"v002_onelab/internal/model"
	"v002_onelab/internal/repository/postgres"
)

type ITransaction interface {
	CreateTransaction(tran model.TransactionResponse) error
	GetAll(userID uint) ([]model.Transaction, error)
	GetByID(id uint) (model.Transaction, error)
}

type Transaction struct {
	repo postgres.ITransationRepository
}

func NewTransaction(repo postgres.ITransationRepository) *Transaction {
	return &Transaction{
		repo: repo,
	}
}

func (t Transaction) CreateTransaction(tran model.TransactionResponse) error {
	return t.repo.CreateTransaction(tran)
}

func (t Transaction) GetAll(userID uint) ([]model.Transaction, error) {
	return t.repo.GetAll(userID)
}

func (t Transaction) GetByID(id uint) (model.Transaction, error) {
	return t.repo.GetByID(id)
}
