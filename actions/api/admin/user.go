package admin

import (
	"net/http"
	"strconv"

	"github.com/GracepointMinistries/hub/actions/render"
	"github.com/GracepointMinistries/hub/modelext"
	"github.com/GracepointMinistries/hub/models"
	"github.com/gobuffalo/buffalo"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// UsersResponse returns the queried users
// swagger:response adminUsersResponse
type UsersResponse struct {
	//in:body
	Body UsersResponsePayload
}

// UsersResponsePayload contains the paginated users
type UsersResponsePayload struct {
	Users      []*modelext.UserWithGroup `json:"users"`
	Pagination *modelext.Pagination      `json:"pagination"`
}

// Users returns a list of paginated users
//
// swagger:route GET /api/v1/admin/users admin users
// Returns a paginated list of users.
// responses:
//   200: adminUsersResponse
//	 400: errorResponse
//	 401: errorResponse
//	 403: errorResponse
//	 422: errorResponse
//	 500: errorResponse
func Users(c buffalo.Context) error {
	users, pagination, err := modelext.PaginatedUsers(c, qm.Load(models.UserRels.Groups))
	if err != nil {
		return c.Error(http.StatusInternalServerError, err)
	}
	return c.Render(http.StatusOK, render.JSON(&UsersResponsePayload{
		Users:      users,
		Pagination: pagination,
	}))
}

// UserResponse returns the queried user
// swagger:response adminUserResponse
type UserResponse struct {
	//in:body
	Body UserResponsePayload
}

// UserResponsePayload contains the queried user
type UserResponsePayload struct {
	User *modelext.UserWithGroup `json:"user"`
}

// UserParameters documents the inbound parameters used
// for the User endpoint
// swagger:parameters user
type UserParameters struct {
	// in:path
	// required:true
	ID int `json:"id"`
}

// User returns a user and its group.
//
// swagger:route GET /api/v1/admin/users/{id} admin user
// Returns a user and its group.
// responses:
//   200: adminUserResponse
//	 400: errorResponse
//	 401: errorResponse
//	 403: errorResponse
//	 422: errorResponse
//	 500: errorResponse
func User(c buffalo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.Error(http.StatusBadRequest, err)
	}
	user, err := modelext.FindUser(c, id)
	if err != nil {
		return c.Error(http.StatusInternalServerError, err)
	}
	return c.Render(http.StatusOK, render.JSON(&UserResponsePayload{
		User: user,
	}))
}
