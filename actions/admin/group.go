package admin

import (
	"net/http"
	"strconv"

	"github.com/GracepointMinistries/hub/actions/render"
	"github.com/GracepointMinistries/hub/modelext"
	"github.com/gobuffalo/buffalo"
)

// Groups is the paginated view of groups
func Groups(c buffalo.Context) error {
	groups, pagination, err := modelext.PaginatedGroups(c)
	if err != nil {
		return c.Error(http.StatusInternalServerError, err)
	}
	c.Set("pagination", pagination)
	c.Set("groups", groups)
	return c.Render(http.StatusOK, render.HTML("admin/groups.html", "admin.html"))
}

// Group shows detailed information about a specific group
func Group(c buffalo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.Error(http.StatusInternalServerError, err)
	}
	group, err := modelext.FindGroup(c, id)
	if err != nil {
		return c.Error(http.StatusInternalServerError, err)
	}
	c.Set("group", group)
	c.Set("users", group.R.Users)
	return c.Render(http.StatusOK, render.HTML("admin/group.html", "admin.html"))
}
