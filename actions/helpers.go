package actions

import (
	"reflect"

	"github.com/GracepointMinistries/hub/modelext"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/render"
)

// APIError wraps errors in a well-defined api payload
type APIError struct {
	Error string `json:"error"`
}

func apiError(message string) render.Renderer {
	return r.JSON(&APIError{
		Error: message,
	})
}

func isNil(c interface{}) bool {
	return c == nil || (reflect.ValueOf(c).Kind() == reflect.Ptr && reflect.ValueOf(c).IsNil())
}

func addHelpers(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		c.Set("isNil", isNil)
		c.Set("hasZoomLink", modelext.HasZoomLink)
		c.Set("zgroupFor", modelext.ZgroupForUser)
		c.Set("zgroupStatus", modelext.ZgroupStatus)
		c.Set("totalUsersIn", modelext.TotalUsersIn)
		c.Set("userStatus", modelext.UserStatus)
		c.Set("ctx", c)
		return next(c)
	}
}
