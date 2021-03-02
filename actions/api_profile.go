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
	User *models.User `json:"user"`
	//in:body
	Zgroup *models.Zgroup `json:"zgroup"`
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
	return c.Render(http.StatusOK, r.JSON(&ProfileResponse{
		User:   user,
		Zgroup: modelext.ZgroupForUser(user),
	}))
}
