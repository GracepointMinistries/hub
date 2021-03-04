package admin

import (
	"net/http"
	"strconv"

	"github.com/GracepointMinistries/hub/actions/render"
	"github.com/GracepointMinistries/hub/modelext"
	"github.com/GracepointMinistries/hub/models"
	"github.com/gobuffalo/buffalo"
)

// GroupsResponse returns the queried groups
// swagger:response adminGroupsResponse
type GroupsResponse struct {
	//in:body
	Body GroupsResponsePayload
}

// GroupsResponsePayload contains paginated groups
type GroupsResponsePayload struct {
	Groups     []*models.Group      `json:"groups"`
	Pagination *modelext.Pagination `json:"pagination"`
}

// Groups returns a paginated list of groups
//
// swagger:route GET /api/v1/admin/groups admin groups
// Returns a paginated list of groups.
// responses:
//   200: adminGroupsResponse
//	 400: errorResponse
//	 401: errorResponse
//	 403: errorResponse
//	 422: errorResponse
//	 500: errorResponse
func Groups(c buffalo.Context) error {
	groups, pagination, err := modelext.PaginatedGroups(c)
	if err != nil {
		return c.Error(http.StatusInternalServerError, err)
	}
	return c.Render(http.StatusOK, render.JSON(&GroupsResponsePayload{
		Groups:     groups,
		Pagination: pagination,
	}))
}

// GroupResponse returns the queried zgroups
// swagger:response adminGroupResponse
type GroupResponse struct {
	//in:body
	Body GroupResponsePayload
}

// GroupResponsePayload contains the queried group and its users
type GroupResponsePayload struct {
	Group *models.Group  `json:"group"`
	Users []*models.User `json:"users"`
}

// GroupParameters documents the inbound parameters used
// for the Group endpoint
// swagger:parameters group
type GroupParameters struct {
	// in:path
	// required:true
	ID int `json:"id"`
}

// Group returns a group and its users
//
// swagger:route GET /api/v1/admin/groups/{id} admin group
// Returns a group and its users.
// responses:
//   200: adminGroupResponse
//	 400: errorResponse
//	 401: errorResponse
//	 403: errorResponse
//	 422: errorResponse
//	 500: errorResponse
func Group(c buffalo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.Error(http.StatusInternalServerError, err)
	}
	group, err := modelext.FindGroup(c, id)
	if err != nil {
		return c.Error(http.StatusInternalServerError, err)
	}
	return c.Render(http.StatusOK, render.JSON(&GroupResponsePayload{
		Group: group,
		Users: group.R.Users,
	}))
}
