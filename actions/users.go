package actions

import (
	"net/http"

	"github.com/GracepointMinistries/hub/models"
	"github.com/gobuffalo/buffalo"
)

func usersHandler(c buffalo.Context) error {
	users, pagination, err := models.GetUsersWithZGroup(c, getPagination(c))
	if err != nil {
		return c.Error(http.StatusInternalServerError, err)
	}

	if acceptContentType(c) == acceptJSON {
		return c.Render(http.StatusOK, r.JSON(&PaginatedResponse{
			Data:       users,
			Pagination: pagination,
		}))
	}

	c.Set("pagination", pagination)
	c.Set("users", users)
	return c.Render(http.StatusOK, r.HTML("admin/users.html", "admin.html"))
}
