package model

import "errors"

type Transaction struct {
	ID     uint    `json:"id" db:"id"`
	UserID uint    `json:"user_id" db:"user_id"`
	BookID uint    `json:"book_id" db:"book_id"`
	Price  float64 `json:"price" db:"price"`
}

type TransactionInput struct {
	BookID uint    `json:"book_id"`
	Price  float64 `json:"price"`
}

type TransactionResponse struct {
	UserID uint    `json:"user_id"`
	BookID uint    `json:"book_id"`
	Price  float64 `json:"price"`
}

type ErrorResponse struct {
	error string
}

var ErrorNotFound = errors.New("transaction with this id not found")

func NewError(err string) *ErrorResponse {
	return &ErrorResponse{
		error: err,
	}
}
