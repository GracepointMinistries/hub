package actions

import (
	"net/http"

	"github.com/GracepointMinistries/hub/modelext"
	"github.com/GracepointMinistries/hub/models"
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
	User   *models.User   `json:"user"`
	Zgroup *models.Zgroup `json:"zgroup"`
}

// swagger:route GET /api/v1/profile profile
// Returns the users profile.
// responses:
//   200: profileResponse
//	 400: apiErrorResponse
//	 401: apiErrorResponse
//	 403: apiErrorResponse
//	 422: apiErrorResponse
//	 500: apiErrorResponse
func apiProfile(c buffalo.Context) error {
	user, err := modelext.FindProfile(c, c.Session().Get("ID").(int))
	if err != nil {
		return c.Error(http.StatusInternalServerError, err)
	}
	return c.Render(http.StatusOK, r.JSON(&ProfileResponsePayload{
		User:   user,
		Zgroup: modelext.ZgroupForUser(user),
	}))
}
