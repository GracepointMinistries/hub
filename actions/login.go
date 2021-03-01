package actions

import (
	"github.com/gobuffalo/buffalo"
)

func adminLoginHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("admin/login.html", "empty.html"))
}

func loginHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("user/login.html", "empty.html"))
}
