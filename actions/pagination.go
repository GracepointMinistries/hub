package actions

import (
	"strconv"

	"github.com/GracepointMinistries/hub/models"
	"github.com/gobuffalo/buffalo"
)

const defaultLimit = 10

// PaginatedResponse adds pagination to any payload
type PaginatedResponse struct {
	Data       interface{}        `json:"data"`
	Pagination *models.Pagination `json:"pagination"`
}

func getPagination(c buffalo.Context) *models.Pagination {
	pagination := &models.Pagination{
		Limit: defaultLimit,
	}
	limitParam := c.Param("limit")
	if limitParam != "" {
		if limit, err := strconv.Atoi(limitParam); err == nil {
			pagination.Limit = limit
		}
	}
	cursorParam := c.Param("cursor")
	if cursorParam != "" {
		if cursorParam, err := strconv.Atoi(cursorParam); err == nil {
			pagination.Cursor = cursorParam
		}
	}
	return pagination
}
