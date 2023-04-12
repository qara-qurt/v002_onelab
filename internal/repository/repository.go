package repository

import (
	"v002_onelab/configs"
	"v002_onelab/internal/repository/postgres"
)

type Repository struct {
	Transaction postgres.ITransationRepository
}

func New(conf *configs.Config) (*Repository, error) {
	db, err := postgres.NewDatabasePSQL(conf)
	if err != nil {
		return nil, err
	}
	transaction := postgres.NewTransaction(db.DB)
	return &Repository{
		Transaction: transaction,
	}, nil
}
