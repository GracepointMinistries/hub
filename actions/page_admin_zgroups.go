package actions

import (
	"net/http"
	"strconv"

	"github.com/GracepointMinistries/hub/modelext"
	"github.com/gobuffalo/buffalo"
)

func adminZgroupsPage(c buffalo.Context) error {
	zgroups, pagination, err := modelext.PaginatedZgroups(c)
	if err != nil {
		return c.Error(http.StatusInternalServerError, err)
	}
	c.Set("pagination", pagination)
	c.Set("zgroups", zgroups)
	return c.Render(http.StatusOK, r.HTML("admin/zgroups.html", "admin.html"))
}

func adminZgroupPage(c buffalo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.Error(http.StatusInternalServerError, err)
	}
	zgroup, err := modelext.FindZgroup(c, id)
	if err != nil {
		return c.Error(http.StatusInternalServerError, err)
	}
	c.Set("zgroup", zgroup)
	c.Set("users", zgroup.R.Users)
	return c.Render(http.StatusOK, r.HTML("admin/zgroup.html", "admin.html"))
}
