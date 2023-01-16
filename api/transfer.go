package api

import (
	"database/sql"
	"fmt"
	"net/http"
	"simplebank/db/store"

	"github.com/gin-gonic/gin"
)

type transferRequest struct {
	FromAccountID int32  `json:"from_account_id" binding:"required,min=1"`
	ToAccountID   int32  `json:"to_account_id" binding:"required,min=1"`
	Amount        int64  `json:"amount" binding:"required,gt=0"`
	Currency      string `json:"currency" binding:"required,oneof=USD EUR CAD"`
}

func (server *Server) createTransfer(ctx *gin.Context) {
	var req transferRequest
	if !server.ValidAccount(ctx, int64(req.FromAccountID), req.Currency) {
		ctx.JSON(http.StatusBadRequest, "not a valid account ")
		return
	}
	if !server.ValidAccount(ctx, int64(req.ToAccountID), req.Currency) {
		ctx.JSON(http.StatusBadRequest, "not a valid account ")
		return

	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	arg := store.TransferTxParams{
		FromAccountID: req.FromAccountID,
		ToAccountID:   req.ToAccountID,
		Amount:        int32(req.Amount),
	}
	account, err := server.store.TransferTx(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
	} else {
		ctx.JSON(http.StatusOK, account)
	}

}

func (server *Server) ValidAccount(ctx *gin.Context, accountID int64, currency string) bool {
	account, err := server.store.GetAccount(ctx, int64(accountID))
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, err.Error())
			return false
		}
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return false
	}

	if account.Currency != currency {
		err = fmt.Errorf("the currency type missmatch for account %d and currency %s", account.ID, currency)
		ctx.JSON(http.StatusBadRequest, err.Error())
		return false
	}
	return true
}
