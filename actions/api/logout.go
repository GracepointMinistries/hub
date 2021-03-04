package api

import (
	"net/http"

	"github.com/GracepointMinistries/hub/actions/api/common"
	"github.com/GracepointMinistries/hub/actions/render"
	"github.com/GracepointMinistries/hub/actions/utils"
	"github.com/GracepointMinistries/hub/modelext"
	"github.com/gobuffalo/buffalo"
)

// Logout ends the session of the user or scoped admin
//
// swagger:route DELETE /api/v1/logout user logout
// Log out of the user account.
// responses:
//   200: tokenResponse
//	 400: errorResponse
//	 401: errorResponse
//	 403: errorResponse
//	 422: errorResponse
//	 500: errorResponse
func Logout(c buffalo.Context) error {
	sessionID := c.Value("SessionID").(int)
	admin := c.Value("Admin")
	if admin != nil {
		// "log out" by unscoping out auth token
		token, err := utils.GenerateScopedToken(admin.(string), 0, sessionID)
		if err != nil {
			return c.Error(http.StatusBadRequest, err)
		}
		return c.Render(http.StatusOK, render.JSON(&common.TokenPayload{
			Token: token,
		}))
	}
	if err := modelext.DeleteUserSession(c, sessionID); err != nil {
		return c.Error(http.StatusInternalServerError, err)
	}
	return c.Render(http.StatusOK, render.JSON(&common.TokenPayload{
		Token: "",
	}))
}
