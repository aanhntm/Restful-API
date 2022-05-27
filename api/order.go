package api

import (
	"net/http"

	db "github.com/aanhntm/restful-api/db/sqlc"
	"github.com/gin-gonic/gin"
)

type CreateOrderRequest struct {
	UserName    string `json:"user" binding:"required"`
	ProductName string `json:"product" binding:"required"`
	Amount      int32  `json:"amount" binding:"required"`
}

func (server *Server) CreateOrder(ctx *gin.Context) {
	var req CreateOrderRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateOrderParams{
		UserName:    req.UserName,
		ProductName: req.ProductName,
		Amount:      req.Amount,
	}

	order, err := server.record.CreateOrder(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, order)
}

func (server *Server) GetOrder(ctx *gin.Context) {

	arg := db.GetManyOrdersParams{
		Limit:  5,
		Offset: 0,
	}

	order, err := server.record.GetManyOrders(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.IndentedJSON(http.StatusOK, order)
}
