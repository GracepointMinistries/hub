package actions

import (
	"net/http"
	"strconv"

	"github.com/GracepointMinistries/hub/models"
	"github.com/gobuffalo/buffalo"
)

func zgroupsHandler(c buffalo.Context) error {
	zgroups, pagination, err := models.GetZGroups(c, true, getPagination(c))
	if err != nil {
		return c.Error(http.StatusInternalServerError, err)
	}

	if acceptContentType(c) == acceptJSON {
		return c.Render(http.StatusOK, r.JSON(&PaginatedResponse{
			Data:       zgroups,
			Pagination: pagination,
		}))
	}

	c.Set("pagination", pagination)
	c.Set("zgroups", zgroups)
	return c.Render(http.StatusOK, r.HTML("admin/zgroups.html", "admin.html"))
}

func zgroupHandler(c buffalo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.Error(http.StatusInternalServerError, err)
	}
	zgroup, err := models.GetZGroup(c, id)
	if err != nil {
		return c.Error(http.StatusInternalServerError, err)
	}

	if acceptContentType(c) == acceptJSON {
		return c.Render(http.StatusOK, r.JSON(zgroup))
	}

	c.Set("zgroup", zgroup)
	return c.Render(http.StatusOK, r.HTML("admin/zgroup.html", "admin.html"))
}
