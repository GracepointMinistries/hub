package actions

import (
	"github.com/GracepointMinistries/hub/models"
	"github.com/gobuffalo/buffalo"
)

func zgroupFor(user *models.User) *models.Zgroup {
	zgroupsLen := len(user.R.Zgroups)
	if zgroupsLen == 0 {
		return nil
	}
	return user.R.Zgroups[zgroupsLen-1]
}

func zgroupStatus(zgroup *models.Zgroup) string {
	if zgroup.Archived {
		return "Archived"
	}
	return "Active"
}

func totalUsersIn(c buffalo.Context, zgroup *models.Zgroup) (int64, error) {
	tx := getTx(c)
	return zgroup.Users().Count(c, tx)
}

func userStatus(user *models.User) string {
	if user.Blocked {
		return "Blocked"
	}
	return "Active"
}

func addHelpers(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		c.Set("zgroupFor", zgroupFor)
		c.Set("zgroupStatus", zgroupStatus)
		c.Set("totalUsersIn", totalUsersIn)
		c.Set("userStatus", userStatus)
		c.Set("ctx", c)
		return next(c)
	}
}
