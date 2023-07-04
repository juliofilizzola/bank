package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/juliofilizzola/bank/internal/db/sqlc"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()
	router.POST("account", server.CreateAccount)
	router.GET("account", server.ListAccounts)
	router.GET("account/:id", server.GetAccount)
	router.POST("account/:id", server.AddCash)

	router.GET("entries/:id", server.ListEntries)

	router.POST("transfer", server.CreateTransfer)
	server.router = router
	return server
}

func (s Server) Start(address string) error {
	return s.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
