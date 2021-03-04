package modelext

import (
	"strconv"

	"github.com/gobuffalo/buffalo"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

const defaultLimit = 10

// PaginationParameters wrap the inbound parameters
// swagger:parameters users groups
type PaginationParameters struct {
	// in:query
	Limit int `json:"limit"`
	// in:query
	Cursor int `json:"cursor"`
	// in:query
	Filter string `json:"filter"`
}

// Pagination holds query data for pagination
type Pagination struct {
	Limit  int    `json:"limit"`
	Cursor int    `json:"cursor"`
	Filter string `json:"filter"`
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
	whereClause := qm.Where("id > ?", cursor)
	filter := c.Param("filter")
	if filter != "" {
		whereClause = qm.Where("id > ? AND name ILIKE '%' || ? || '%'", cursor, filter)
	}
	clauses := []qm.QueryMod{whereClause}
	if limit >= 1 {
		// if we have a positive limit, set it
		clauses = append(clauses, qm.Limit(limit))
	}
	return &Pagination{
		Limit:  limit,
		Cursor: cursor,
		Filter: filter,
	}, append(clauses, qm.OrderBy("id ASC"))
}
