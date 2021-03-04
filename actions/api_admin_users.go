package actions

import (
	"net/http"
	"strconv"

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
	Users      []*modelext.UserWithZgroup `json:"users"`
	Pagination *modelext.Pagination       `json:"pagination"`
}

// swagger:route GET /api/v1/admin/users admin users
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

// AdminUserResponse returns the queried user
// swagger:response adminUserResponse
type AdminUserResponse struct {
	//in:body
	Body AdminUserResponsePayload
}

// AdminUserResponsePayload contains the queried user
type AdminUserResponsePayload struct {
	User *modelext.UserWithZgroup `json:"user"`
}

// AdminUserParameters documents the inbound parameters used
// for the apiAdminUser endpoint
// swagger:parameters user
type AdminUserParameters struct {
	// in:path
	// required:true
	ID int `json:"id"`
}

// swagger:route GET /api/v1/admin/users/{id} admin user
// Returns a user and its users.
// responses:
//   200: adminUserResponse
//	 400: apiErrorResponse
//	 401: apiErrorResponse
//	 403: apiErrorResponse
//	 422: apiErrorResponse
//	 500: apiErrorResponse
func apiAdminUser(c buffalo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.Error(http.StatusBadRequest, err)
	}
	user, err := modelext.FindUser(c, id)
	if err != nil {
		return c.Error(http.StatusInternalServerError, err)
	}
	return c.Render(http.StatusOK, r.JSON(&ProfileResponsePayload{
		User: user,
	}))
}
