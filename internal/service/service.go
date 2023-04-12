package service

import "v002_onelab/internal/repository"

type Service struct {
	Transaction ITransaction
}

func New(repo *repository.Repository) *Service {
	transaction := NewTransaction(repo.Transaction)
	return &Service{
		Transaction: transaction,
	}
}
