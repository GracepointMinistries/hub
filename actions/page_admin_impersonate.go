package actions

import (
	"net/http"
	"strconv"

	"github.com/GracepointMinistries/hub/modelext"
	"github.com/gobuffalo/buffalo"
)

func adminImpersonatePage(c buffalo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.Error(http.StatusInternalServerError, err)
	}
	user, err := modelext.FindUser(c, id)
	if err != nil {
		return c.Error(http.StatusInternalServerError, err)
	}

	c.Session().Set("ID", user.ID)
	c.Session().Set("Email", user.Email)
	return c.Redirect(http.StatusSeeOther, "/")
}
