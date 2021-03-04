package actions

import (
	"net/http"

	"github.com/GracepointMinistries/hub/modelext"
	"github.com/gobuffalo/buffalo"
)

// swagger:route DELETE /api/v1/logout user logout
// Log out of the user account.
// responses:
//   200: tokenResponse
//	 400: apiErrorResponse
//	 401: apiErrorResponse
//	 403: apiErrorResponse
//	 422: apiErrorResponse
//	 500: apiErrorResponse
func apiLogout(c buffalo.Context) error {
	sessionID := c.Value("SessionID").(int)
	admin := c.Value("Admin")
	if admin != nil {
		// "log out" by unscoping out auth token
		token, err := generateScopedToken(admin.(string), 0, sessionID)
		if err != nil {
			return c.Error(http.StatusBadRequest, err)
		}
		return c.Render(http.StatusOK, r.JSON(&TokenPayload{
			Token: token,
		}))
	}
	if err := modelext.DeleteUserSession(c, sessionID); err != nil {
		return c.Error(http.StatusInternalServerError, err)
	}
	return c.Render(http.StatusOK, r.JSON(&TokenPayload{
		Token: "",
	}))
}
