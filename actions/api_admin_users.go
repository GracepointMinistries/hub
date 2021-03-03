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
	Body AdminUsersResponsePayload
}

// AdminUsersResponsePayload contains the paginated users
type AdminUsersResponsePayload struct {
	Users      []*models.User       `json:"users"`
	Pagination *modelext.Pagination `json:"pagination"`
}

// swagger:route GET /api/v1/admin/users adminUsers
// Returns a paginated list of users.
// responses:
//   200: adminUsersResponse
//	 400: apiErrorResponse
//	 401: apiErrorResponse
//	 403: apiErrorResponse
//	 422: apiErrorResponse
//	 500: apiErrorResponse
func apiAdminUsers(c buffalo.Context) error {
	users, pagination, err := modelext.PaginatedUsers(c, qm.Load(models.UserRels.Zgroups))
	if err != nil {
		return c.Error(http.StatusInternalServerError, err)
	}
	return c.Render(http.StatusOK, r.JSON(&AdminUsersResponsePayload{
		users,
		pagination,
	}))
}
