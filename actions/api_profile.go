package actions

import (
	"net/http"

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
	User *modelext.UserWithZgroup `json:"user"`
}

// swagger:route GET /api/v1/profile user profile
// Returns the users profile.
// responses:
//   200: profileResponse
//	 400: apiErrorResponse
//	 401: apiErrorResponse
//	 403: apiErrorResponse
//	 422: apiErrorResponse
//	 500: apiErrorResponse
func apiProfile(c buffalo.Context) error {
	user, err := modelext.FindUser(c, c.Value("ID").(int))
	if err != nil {
		return c.Error(http.StatusInternalServerError, err)
	}
	return c.Render(http.StatusOK, r.JSON(&ProfileResponsePayload{
		User: user,
	}))
}
