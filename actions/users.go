package actions

import (
	"net/http"

	"github.com/GracepointMinistries/hub/models"
	"github.com/gobuffalo/buffalo"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func usersHandler(c buffalo.Context) error {
	tx := getTx(c)
	clauses := append(getPagination(c), qm.Load(models.UserRels.Zgroups))
	users, err := models.Users(
		clauses...,
	).All(c, tx)
	if err != nil {
		return c.Error(http.StatusInternalServerError, err)
	}

	paginatedUsers := paginateUsers(c, users)
	if acceptContentType(c) == acceptJSON {
		return c.Render(http.StatusOK, r.JSON(paginatedUsers))
	}

	c.Set("pagination", paginatedUsers.Pagination)
	c.Set("users", users)
	return c.Render(http.StatusOK, r.HTML("admin/users.html", "admin.html"))
}
