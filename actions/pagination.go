package actions

import (
	"strconv"

	"github.com/GracepointMinistries/hub/models"
	"github.com/gobuffalo/buffalo"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

const defaultLimit = 10

// Pagination holds query data for pagination
type Pagination struct {
	Limit  int `json:"limit"`
	Cursor int `json:"cursor"`
}

// PaginatedResponse adds pagination to any payload
type PaginatedResponse struct {
	Data       interface{} `json:"data"`
	Pagination *Pagination `json:"pagination"`
}

func getPagination(c buffalo.Context) []qm.QueryMod {
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
	return []qm.QueryMod{qm.Where("id > ?", cursor), qm.Limit(limit), qm.OrderBy("id ASC")}
}

func paginateUsers(c buffalo.Context, users []*models.User) *PaginatedResponse {
	limit := defaultLimit
	limitParam := c.Param("limit")
	if limitParam != "" {
		if l, err := strconv.Atoi(limitParam); err == nil {
			limit = l
		}
	}
	cursor := -1
	if len(users) > 0 {
		cursor = users[len(users)-1].ID
	}
	return &PaginatedResponse{
		Data: users,
		Pagination: &Pagination{
			Limit:  limit,
			Cursor: cursor,
		},
	}
}

func paginateZgroups(c buffalo.Context, zgroups []*models.Zgroup) *PaginatedResponse {
	limit := defaultLimit
	limitParam := c.Param("limit")
	if limitParam != "" {
		if l, err := strconv.Atoi(limitParam); err == nil {
			limit = l
		}
	}
	cursor := -1
	if len(zgroups) > 0 {
		cursor = zgroups[len(zgroups)-1].ID
	}
	return &PaginatedResponse{
		Data: zgroups,
		Pagination: &Pagination{
			Limit:  limit,
			Cursor: cursor,
		},
	}
}
