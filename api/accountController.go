package api

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/juliofilizzola/bank/internal/db/sqlc"
)

type CreateAccountParamsBody struct {
	Owner    string `json:"owner"`
	Currency string `json:"currency"`
}

func (s Server) CreateAccount(ctx *gin.Context) {
	var body CreateAccountParamsBody
	var objBody db.CreateAccountParams

	err := ctx.Bind(&body)

	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	objBody.Balance = 0
	objBody.Owner = body.Owner
	objBody.Currency = body.Currency

	result, err := s.store.CreateAccount(context.Background(), objBody)

	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	res, err := result.LastInsertId()

	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}
	var convertId = int(res)

	var id = int32(convertId)

	account, err := s.store.GetAccount(context.Background(), id)

	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"account": account,
	})
}
