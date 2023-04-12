package http

import (
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	_ "v002_onelab/docs"
	"v002_onelab/internal/model"
)

// BuyBook @Summary CreateTransaction(buy-book)
// @Tags transaction
// @Security ApiKeyAuth
// @Description buy book
// @ID transacton
// @Accept json
// @Produce json
// @Param  input body model.TransactionInput  true "book info and price"
// @Success 201
// @Failure 500 {object} model.ErrorResponse
// @Failure 400 {object} model.ErrorResponse
// @Router /buy-book/ [post]
func (h *Handler) BuyBook(c echo.Context) error {
	userID := c.Get("id")
	var transactionInp model.TransactionInput
	if err := c.Bind(&transactionInp); err != nil {
		return c.JSON(http.StatusBadRequest, model.NewError(err.Error()))
	}

	var transaction model.TransactionResponse
	transaction.UserID = userID.(uint)
	transaction.BookID = transactionInp.BookID
	transaction.Price = transactionInp.Price
	if err := h.Transaction.CreateTransaction(transaction); err != nil {
		return c.JSON(http.StatusInternalServerError, model.NewError(err.Error()))
	}

	return c.NoContent(http.StatusCreated)
}

// GetBuyBooks @Summary GetBuyBooks(buy-book)
// @Tags transaction
// @Security ApiKeyAuth
// @Description get buy books
// @ID get-buy-books
// @Accept json
// @Produce json
// @Success 200 {object} []model.Transaction
// @Failure 500 {object} model.ErrorResponse
// @Router /buy-book/ [get]
func (h *Handler) GetBuyBooks(c echo.Context) error {
	userID := c.Get("id")
	res, err := h.Transaction.GetAll(userID.(uint))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.NewError(err.Error()))
	}

	return c.JSON(http.StatusOK, res)
}

// GetBuyBook @Summary GetBuyBook(buy-book by id)
// @Tags transaction
// @Security ApiKeyAuth
// @Description get buy books
// @ID get-buy-book
// @Accept json
// @Produce json
// @Param  id path int  true "transaction ID"
// @Success 200 {object} []model.Transaction
// @Failure 500 {object} model.ErrorResponse
// @Failure 400 {object} model.ErrorResponse
// @Router /buy-book/{id} [get]
func (h *Handler) GetBuyBook(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.NewError(err.Error()))
	}

	res, err := h.Transaction.GetByID(uint(id))
	if err != nil {
		if errors.Is(model.ErrorNotFound, err) {
			return c.JSON(http.StatusBadRequest, model.NewError(err.Error()))
		}
		return c.JSON(http.StatusInternalServerError, model.NewError(err.Error()))
	}

	return c.JSON(http.StatusCreated, res)
}
