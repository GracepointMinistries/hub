package actions

import (
	"net/http"
	"strconv"

	"github.com/GracepointMinistries/hub/models"
	"github.com/gobuffalo/buffalo"
)

func impersonateHandler(c buffalo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.Error(http.StatusInternalServerError, err)
	}
	tx := getTx(c)
	user, err := models.FindUser(c, tx, id)
	if err != nil {
		return c.Error(http.StatusInternalServerError, err)
	}

	c.Session().Set("ID", user.ID)
	c.Session().Set("Email", user.Email)
	return c.Redirect(http.StatusSeeOther, "/")
}
