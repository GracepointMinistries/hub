package api

import (
	"net/http"

	"github.com/GracepointMinistries/hub/actions/render"
	"github.com/GracepointMinistries/hub/modelext"
	"github.com/gobuffalo/buffalo"
)

// ProfileResponse returns the user profile information
// swagger:response profileResponse
type ProfileResponse struct {
	//in:body
	Body ProfileResponsePayload
}

// ProfileResponsePayload contains user profile information
type ProfileResponsePayload struct {
	User *modelext.UserWithGroup `json:"user"`
}

// Profile returns the current user's profile
//
// swagger:route GET /api/v1/profile user profile
// Returns the users profile.
// responses:
//   200: profileResponse
//	 400: errorResponse
//	 401: errorResponse
//	 403: errorResponse
//	 422: errorResponse
//	 500: errorResponse
func Profile(c buffalo.Context) error {
	user, err := modelext.FindUser(c, c.Value("ID").(int))
	if err != nil {
		return c.Error(http.StatusInternalServerError, err)
	}
	return c.Render(http.StatusOK, render.JSON(&ProfileResponsePayload{
		User: user,
	}))
}
