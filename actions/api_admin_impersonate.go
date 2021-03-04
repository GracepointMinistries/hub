package actions

import (
	"net/http"
	"strconv"

	"github.com/gobuffalo/buffalo"
)

// AdminImpersonateParameters documents the inbound parameters used
// for the apiAdminImpersonate endpoint
// swagger:parameters impersonate
type AdminImpersonateParameters struct {
	// in:path
	// required:true
	ID int `json:"id"`
}

// swagger:route POST /api/v1/admin/impersonate/{id} admin impersonate
// Gets an authentication token for an admin that allows scoping as a user
// responses:
//   200: tokenResponse
//	 400: apiErrorResponse
//	 401: apiErrorResponse
//	 403: apiErrorResponse
//	 422: apiErrorResponse
//	 500: apiErrorResponse
func apiAdminImpersonate(c buffalo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.Error(http.StatusBadRequest, err)
	}
	admin := c.Value("Admin").(string)
	sessionID := c.Value("SessionID").(int)
	token, err := generateScopedToken(admin, id, sessionID)
	if err != nil {
		return c.Error(http.StatusBadRequest, err)
	}
	return c.Render(http.StatusOK, r.JSON(&TokenPayload{
		Token: token,
	}))
}
