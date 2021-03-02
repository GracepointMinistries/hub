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
	Users []*models.User `json:"users"`
	//in:body
	Pagination *modelext.Pagination `json:"pagination"`
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
	return c.Render(http.StatusOK, r.JSON(&AdminUsersResponse{
		Users:      users,
		Pagination: pagination,
	}))
}
