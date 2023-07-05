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

func (s Server) ListEntries(ctx *gin.Context) {
	var paramId = ctx.Param("id")
	idConvert, err := strconv.Atoi(paramId)

	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, errorResponse(errors.New("internal server err")))
		return
	}

	id := int32(idConvert)

	limitQuery := ctx.Query("limit")
	limitConvert, err := strconv.Atoi(limitQuery)
	limit := int32(limitConvert)

	pageQuery := ctx.Query("page")
	pageConvert, err := strconv.Atoi(pageQuery)
	page := int32(pageConvert)

	entries, err := s.store.ListEntries(context.Background(), db.ListEntriesParams{
		AccountID: id,
		Limit:     limit,
		Offset:    page,
	})

	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusNotFound, errorResponse(errors.New("entries not found")))
		return
	}

	i := len(entries)

	if i < 0 {
		fmt.Println(err)
		ctx.JSON(http.StatusNotFound, errorResponse(errors.New("entries not found")))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"entries": entries,
	})
}
