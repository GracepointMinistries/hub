package actions

import (
	"net/http"

	"github.com/GracepointMinistries/hub/modelext"
	"github.com/GracepointMinistries/hub/models"
	"github.com/gobuffalo/buffalo"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// AdminUsersResponse returns the queried users
// swagger:response adminUsersResponse
type AdminUsersResponse struct {
	//in:body
	Body struct {
		Users      []*models.User       `json:"users"`
		Pagination *modelext.Pagination `json:"pagination"`
	}
}

func adminUsersResponse(users []*models.User, pagination *modelext.Pagination) *AdminUsersResponse {
	response := &AdminUsersResponse{}
	response.Body.Users = users
	response.Body.Pagination = pagination
	return response
}

// swagger:route GET /api/v1/admin/users adminUsers
// Returns a paginated list of users.
// responses:
//   200: adminUsersResponse
func apiAdminUsers(c buffalo.Context) error {
	users, pagination, err := modelext.PaginatedUsers(c, qm.Load(models.UserRels.Zgroups))
	if err != nil {
		return c.Error(http.StatusInternalServerError, err)
	}
	return c.Render(http.StatusOK, r.JSON(adminUsersResponse(
		users,
		pagination,
	).Body))
}
