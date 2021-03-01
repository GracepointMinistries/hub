package actions

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
)

func adminHandler(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("admin/index.html", "admin.html"))
}
