package modelext

import (
	"strconv"

	"github.com/gobuffalo/buffalo"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

const defaultLimit = 10

// PaginationParameters wrap the inbound parameters
// swagger:parameters adminUsers adminZgroups
type PaginationParameters struct {
	// in: query
	Limit int `json:"limit"`
	// in: query
	Cursor int `json:"cursor"`
}

// Pagination holds query data for pagination
type Pagination struct {
	Limit  int `json:"limit"`
	Cursor int `json:"cursor"`
}

func getPagination(c buffalo.Context) (*Pagination, []qm.QueryMod) {
	limit := defaultLimit
	limitParam := c.Param("limit")
	if limitParam != "" {
		if l, err := strconv.Atoi(limitParam); err == nil {
			limit = l
		}
	}
	cursor := 0
	cursorParam := c.Param("cursor")
	if cursorParam != "" {
		if c, err := strconv.Atoi(cursorParam); err == nil {
			cursor = c
		}
	}
	return &Pagination{
		Limit:  limit,
		Cursor: cursor,
	}, []qm.QueryMod{qm.Where("id > ?", cursor), qm.Limit(limit), qm.OrderBy("id ASC")}
}
