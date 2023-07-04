package api

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"

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
		errorResponse(err)
		return
	}

	objBody.Balance = 0
	objBody.Owner = body.Owner
	objBody.Currency = body.Currency

	result, err := s.store.CreateAccount(context.Background(), objBody)

	if err != nil {
		ctx.Status(http.StatusBadRequest)
		errorResponse(err)
		return
	}

	res, err := result.LastInsertId()

	if err != nil {
		ctx.Status(http.StatusBadRequest)
		errorResponse(err)
		return
	}
	var convertId = int(res)

	var id = int32(convertId)

	account, err := s.store.GetAccount(context.Background(), id)

	if err != nil {
		ctx.Status(http.StatusBadRequest)
		errorResponse(err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"account": account,
	})
}

func (s Server) GetAccount(ctx *gin.Context) {
	var paramId = ctx.Param("id")

	idConvert, err := strconv.Atoi(paramId)

	if err != nil {
		ctx.Status(http.StatusBadRequest)
		errorResponse(err)
		return
	}

	id := int32(idConvert)

	account, err := s.store.GetAccount(context.Background(), id)

	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusNotFound, errorResponse(errors.New("account not found")))
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"account": account,
	})
}

func (s Server) ListAccounts(ctx *gin.Context) {
	limitQuery := ctx.Query("limit")
	limitConvert, err := strconv.Atoi(limitQuery)
	limit := int32(limitConvert)

	pageQuery := ctx.Query("page")
	pageConvert, err := strconv.Atoi(pageQuery)
	page := int32(pageConvert)

	accounts, err := s.store.ListAccounts(context.Background(), db.ListAccountsParams{
		Limit:  limit,
		Offset: page,
	})

	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusNotFound, errorResponse(errors.New("account not found")))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"accounts": accounts,
	})
}
