package actions

import (
	"net/http"

	"github.com/GracepointMinistries/hub/models"
	"github.com/gobuffalo/buffalo"
)

func profileHandler(c buffalo.Context) error {
	profile, err := models.GetUser(c, c.Session().Get("ID").(int))
	if err != nil {
		return c.Error(http.StatusInternalServerError, err)
	}

	if acceptContentType(c) == acceptJSON {
		return c.Render(http.StatusOK, r.JSON(profile))
	}

	c.Set("profile", profile)
	return c.Render(http.StatusOK, r.HTML("profile.html"))
}
