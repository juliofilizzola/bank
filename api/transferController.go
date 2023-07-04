package api

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/juliofilizzola/bank/internal/db/sqlc"
)

type TransferBody struct {
	ToAccountId   int32 `json:"to_account_id"`
	FromAccountID int32 `json:"from_account_id"`
	Amount        int64 `json:"amount"`
}

func (s Server) CreateTransfer(ctx *gin.Context) {
	var body TransferBody

	err := ctx.Bind(&body)

	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusNotFound, errorResponse(errors.New("account not found")))
		return
	}

	res, err := s.store.TransferTx(context.Background(), db.TransferTxParams{
		FromAccountID: body.FromAccountID,
		ToAccountId:   body.ToAccountId,
		Amount:        body.Amount,
	})
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusNotFound, errorResponse(errors.New("account not found")))
		return
	}
	// fmt.Println(&res)
	ctx.JSON(http.StatusCreated, gin.H{
		"account": res,
	})

}
