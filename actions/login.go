package actions

import (
	"github.com/gobuffalo/buffalo"
)

func loginHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("login.html", "empty.html"))
}
