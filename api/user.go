package api

import (
	"fmt"
	"net/http"
	"simplebank/db/gen"
	"simplebank/util"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type createUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Passwoed string `json:"password" binding:"required,min=6"`
	FullName string `json:"fullname" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	HashedPassword, err1 := util.HashPassword(req.Passwoed)
	if err1 != nil {
		ctx.JSON(http.StatusInternalServerError, err1.Error())
	}
	arg := gen.CreateUserParams{
		Username:       req.Username,
		HashedPassword: HashedPassword,
		FullName:       req.FullName,
		Email:          req.Email,
	}
	account, err := server.store.CreateUser(ctx, arg)
	fmt.Println(err.Error())
	if err != nil {
		if pqerr, ok := err.(*pq.Error); ok {
			fmt.Println(pqerr.Code.Name())
			switch pqerr.Code.Name() {
			case  "unique_violation":
				ctx.JSON(http.StatusForbidden, err.Error())
				return

			}

		}

		ctx.JSON(http.StatusInternalServerError, err.Error())
		return 
	}

	ctx.JSON(http.StatusOK, account)

}
