package actions

import (
	"reflect"

	"github.com/GracepointMinistries/hub/modelext"
	"github.com/gobuffalo/buffalo"
)

func isNil(c interface{}) bool {
	return c == nil || (reflect.ValueOf(c).Kind() == reflect.Ptr && reflect.ValueOf(c).IsNil())
}

func addHelpers(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		c.Set("isNil", isNil)
		c.Set("hasZoomLink", modelext.HasZoomLink)
		c.Set("groupStatus", modelext.GroupStatus)
		c.Set("totalUsersIn", modelext.TotalUsersIn)
		c.Set("userStatus", modelext.UserStatus)
		c.Set("ctx", c)
		return next(c)
	}
}
