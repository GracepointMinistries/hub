package actions

import (
	"net/http"

	"github.com/GracepointMinistries/hub/modelext"
	"github.com/GracepointMinistries/hub/models"
	"github.com/gobuffalo/buffalo"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func adminUsersPage(c buffalo.Context) error {
	users, pagination, err := modelext.PaginatedUsers(c, qm.Load(models.UserRels.Zgroups))
	if err != nil {
		return c.Error(http.StatusInternalServerError, err)
	}
	c.Set("pagination", pagination)
	c.Set("users", users)
	return c.Render(http.StatusOK, r.HTML("admin/users.html", "admin.html"))
}
