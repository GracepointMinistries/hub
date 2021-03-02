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
	Body struct {
		User   *models.User   `json:"user"`
		Zgroup *models.Zgroup `json:"zgroup"`
	}
}

func profileResponse(user *models.User, zgroup *models.Zgroup) *ProfileResponse {
	response := &ProfileResponse{}
	response.Body.User = user
	response.Body.Zgroup = zgroup
	return response
}

// swagger:route GET /api/v1/profile profile
// Returns the users profile.
// responses:
//   200: profileResponse
func apiProfile(c buffalo.Context) error {
	user, err := modelext.FindProfile(c, c.Session().Get("ID").(int))
	if err != nil {
		return c.Error(http.StatusInternalServerError, err)
	}
	return c.Render(http.StatusOK, r.JSON(profileResponse(
		user,
		modelext.ZgroupForUser(user),
	).Body))
}
