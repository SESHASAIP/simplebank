package api

import (
	"database/sql"
	"fmt"
	"net/http"

	"simplebank/db/gen"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type createAccountRequest struct {
	Owner    string `json:"owner" binding:"required"`
	Currency string `json:"currency" binding:"required,oneof=USD EUR CAD"`
}

func (server *Server) createAccount(ctx *gin.Context) {
	var req createAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	arg := gen.CreateAuthorParams{
		Owner:    req.Owner,
		Currency: req.Currency,
		Balance:  0,
	}
	account, err := server.store.CreateAuthor(ctx, arg)
	fmt.Println(err.Error())
	if err != nil {
		if pqerr, ok := err.(*pq.Error); ok {
			fmt.Println(pqerr.Code.Name())
			switch pqerr.Code.Name() {
			case "forign_key_violation", "unique_violation":
				ctx.JSON(http.StatusForbidden, err.Error())

			}

		}

		ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	ctx.JSON(http.StatusOK, account)

}

type getAccountByIdRequest struct {
	Id int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getAccountById(ctx *gin.Context) {
	var req getAccountByIdRequest
	err := ctx.ShouldBindUri(&req)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return

	}
	fmt.Println(req.Id)
	//err3:= strconv.Atoi(req.id)
	account, err1 := server.store.GetAccount(ctx, req.Id)
	if err1 != nil {
		if err1 == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, err1.Error())
			return
		}
		ctx.JSON(http.StatusInternalServerError, err1.Error())
		return
	}
	ctx.JSON(http.StatusOK, account)

}
