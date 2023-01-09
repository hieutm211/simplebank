package api

import (
	"context"
	"fmt"
	"net/http"
	db "simplebank/db/sqlc"

	"github.com/gin-gonic/gin"
)

type TransferMoneyParams struct {
	FromAccountID int64  `json:"from_account_id" binding:"required,min=1"`
	ToAccountID   int64  `json:"to_account_id" binding:"required,min=1"`
	Amount        int64  `json:"amount" binding:"required,min=1"`
	Currency      string `json:"currency" binding:"required,currency"`
}

func (s *Server) transferMoney(ctx *gin.Context) {
	var req TransferMoneyParams
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := s.validCurrency(ctx, req.FromAccountID, req.Currency); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	if err := s.validCurrency(ctx, req.ToAccountID, req.Currency); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	result, err := s.store.TransferTx(ctx, db.TransferTxParams{
		FromAccountID: req.FromAccountID,
		ToAccountID:   req.ToAccountID,
		Amount:        req.Amount,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (s *Server) validCurrency(ctx context.Context, accountID int64, currency string) error {
	account, err := s.store.GetAccount(ctx, accountID)
	if err != nil {
		return err
	}

	if account.Currency != currency {
		return fmt.Errorf("currency mismatch at account [%v]: %v vs %v", accountID, account.Currency, currency)
	}

	return nil
}
