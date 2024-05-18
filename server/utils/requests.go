package utils

import (
	"gotodo/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func LimitOffsetParams(context *gin.Context, paginationParams *models.OffsetPaginationParams) error {
	pageQuery := context.DefaultQuery("page", "1")
	limitQuery := context.DefaultQuery("limit", "10")

	page, err := strconv.ParseInt(pageQuery, 10, 64)
	if err != nil {
		return err
	}

	limit, err := strconv.ParseInt(limitQuery, 10, 64)
	if err != nil {
		return err
	}

	paginationParams.Page = int(page)
	paginationParams.Limit = int(limit)

	return err
}
