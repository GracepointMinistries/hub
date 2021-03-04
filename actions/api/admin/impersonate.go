package admin

import (
	"net/http"
	"strconv"

	"github.com/GracepointMinistries/hub/actions/api/common"
	"github.com/GracepointMinistries/hub/actions/render"
	"github.com/GracepointMinistries/hub/actions/utils"
	"github.com/gobuffalo/buffalo"
)

// ImpersonateParameters documents the inbound parameters used
// for the Impersonate endpoint
// swagger:parameters impersonate
type ImpersonateParameters struct {
	// in:path
	// required:true
	ID int `json:"id"`
}

// Impersonate allows an administrator to impersonate a user
//
// swagger:route GET /api/v1/admin/impersonate/{id} admin impersonate
// Gets an authentication token for an admin that allows scoping as a user
// responses:
//   200: tokenResponse
//	 400: errorResponse
//	 401: errorResponse
//	 403: errorResponse
//	 422: errorResponse
//	 500: errorResponse
func Impersonate(c buffalo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.Error(http.StatusBadRequest, err)
	}
	admin := c.Value("Admin").(string)
	sessionID := c.Value("SessionID").(int)
	token, err := utils.GenerateScopedToken(admin, id, sessionID)
	if err != nil {
		return c.Error(http.StatusBadRequest, err)
	}
	return c.Render(http.StatusOK, render.JSON(&common.TokenPayload{
		Token: token,
	}))
}
