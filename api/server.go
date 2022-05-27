package api

import (
	db "github.com/aanhntm/restful-api/db/sqlc"
	"github.com/gin-gonic/gin"
)

type Server struct {
	record *db.Record
	router *gin.Engine
}

func NewServer(record *db.Record) *Server {
	server := &Server{record: record}
	router := gin.Default()

	router.POST("/order", server.CreateOrder)
	router.GET("/order", server.GetOrder)

	server.router = router
	return server
}
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}
}
