package actions

import (
	"net/http"
	"strconv"

	"github.com/GracepointMinistries/hub/models"
	"github.com/gobuffalo/buffalo"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func zgroupsHandler(c buffalo.Context) error {
	tx := getTx(c)
	zgroups, err := models.Zgroups(
		getPagination(c)...,
	).All(c, tx)
	if err != nil {
		return c.Error(http.StatusInternalServerError, err)
	}
	paginatedZgroups := paginateZgroups(c, zgroups)
	if acceptContentType(c) == acceptJSON {
		return c.Render(http.StatusOK, r.JSON(paginatedZgroups))
	}

	c.Set("pagination", paginatedZgroups.Pagination)
	c.Set("zgroups", zgroups)
	return c.Render(http.StatusOK, r.HTML("admin/zgroups.html", "admin.html"))
}

func zgroupHandler(c buffalo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.Error(http.StatusInternalServerError, err)
	}
	tx := getTx(c)
	zgroup, err := models.Zgroups(models.ZgroupWhere.ID.EQ(id), qm.Load(models.ZgroupRels.Users)).One(c, tx)
	if err != nil {
		return c.Error(http.StatusInternalServerError, err)
	}

	if acceptContentType(c) == acceptJSON {
		return c.Render(http.StatusOK, r.JSON(zgroup))
	}

	c.Set("zgroup", zgroup)
	c.Set("users", zgroup.R.Users)
	return c.Render(http.StatusOK, r.HTML("admin/zgroup.html", "admin.html"))
}
