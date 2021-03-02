package actions

import (
	"net/http"

	"github.com/GracepointMinistries/hub/models"
	"github.com/gobuffalo/buffalo"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func profileHandler(c buffalo.Context) error {
	tx := getTx(c)
	id := c.Session().Get("ID").(int)
	user, err := models.Users(
		models.UserWhere.ID.EQ(id),
		qm.Load(models.UserRels.Zgroups, models.ZgroupWhere.Archived.EQ(false)),
	).One(c, tx)
	if err != nil {
		return c.Error(http.StatusInternalServerError, err)
	}

	if acceptContentType(c) == acceptJSON {
		return c.Render(http.StatusOK, r.JSON(user))
	}

	c.Set("user", user)
	return c.Render(http.StatusOK, r.HTML("user/profile.html"))
}
