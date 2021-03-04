package admin

import (
	"net/http"

	"github.com/GracepointMinistries/hub/actions/render"
	"github.com/GracepointMinistries/hub/modelext"
	"github.com/gobuffalo/buffalo"
)

// Users displays a page of users
func Users(c buffalo.Context) error {
	users, pagination, err := modelext.PaginatedUsers(c)
	if err != nil {
		return c.Error(http.StatusInternalServerError, err)
	}
	c.Set("pagination", pagination)
	c.Set("users", users)
	return c.Render(http.StatusOK, render.HTML("admin/users.html", "admin.html"))
}
