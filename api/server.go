package api

import (
	"net/http"
	"simplebank/db/store"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type Server struct {
	store  store.Store
	router *gin.Engine
}

func NewServer(store store.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}
	router.POST("/user", server.createUser)
	router.POST("/accounts", server.createAccount)
	router.POST("/transfers", server.createTransfer)
	router.GET("/accounts/:id", server.getAccountById)
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusAccepted, "jai hind")
	})
	server.router = router

	return server

}
func (s *Server) Startserver(r string) error {
	return s.router.Run(r)
}

// func errResponse(err error) gin.H {
// 	return gin.H{
// 		"error":err.Error()

// 	}

// }
