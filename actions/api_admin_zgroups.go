package actions

import (
	"net/http"
	"strconv"

	"github.com/GracepointMinistries/hub/modelext"
	"github.com/GracepointMinistries/hub/models"
	"github.com/gobuffalo/buffalo"
)

// AdminZgroupsResponse returns the queried zgroups
// swagger:response adminZgroupsResponse
type AdminZgroupsResponse struct {
	//in:body
	Body AdminZgroupsResponsePayload
}

// AdminZgroupsResponsePayload contains paginated zgroups
type AdminZgroupsResponsePayload struct {
	Zgroups    []*models.Zgroup     `json:"zgroups"`
	Pagination *modelext.Pagination `json:"pagination"`
}

// swagger:route GET /api/v1/admin/zgroups adminZgroups
// Returns a paginated list of zGroups.
// responses:
//   200: adminZgroupsResponse
//	 400: apiErrorResponse
//	 401: apiErrorResponse
//	 403: apiErrorResponse
//	 422: apiErrorResponse
//	 500: apiErrorResponse
func apiAdminZgroups(c buffalo.Context) error {
	zgroups, pagination, err := modelext.PaginatedZgroups(c)
	if err != nil {
		return c.Error(http.StatusInternalServerError, err)
	}
	return c.Render(http.StatusOK, r.JSON(&AdminZgroupsResponsePayload{
		Zgroups:    zgroups,
		Pagination: pagination,
	}))
}

// AdminZgroupResponse returns the queried zgroups
// swagger:response adminZgroupResponse
type AdminZgroupResponse struct {
	//in:body
	Body AdminZgroupResponsePayload
}

// AdminZgroupResponsePayload contains the queried zgroup and its users
type AdminZgroupResponsePayload struct {
	Zgroup *models.Zgroup `json:"zgroup"`
	Users  []*models.User `json:"users"`
}

// AdminZgroupParameters documents the inbound parameters used
// for the apiAdminZgroup endpoint
// swagger:parameters adminZgroup
type AdminZgroupParameters struct {
	// in:path
	ID int `json:"id"`
}

// swagger:route GET /api/v1/admin/zgroups/{id} adminZgroup
// Returns a zGroup and its users.
// responses:
//   200: adminZgroupResponse
//	 400: apiErrorResponse
//	 401: apiErrorResponse
//	 403: apiErrorResponse
//	 422: apiErrorResponse
//	 500: apiErrorResponse
func apiAdminZgroup(c buffalo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.Error(http.StatusInternalServerError, err)
	}
	zgroup, err := modelext.FindZgroup(c, id)
	if err != nil {
		return c.Error(http.StatusInternalServerError, err)
	}
	return c.Render(http.StatusOK, r.JSON(&AdminZgroupResponsePayload{
		Zgroup: zgroup,
		Users:  zgroup.R.Users,
	}))
}
