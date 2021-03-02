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
	Body struct {
		Zgroups    []*models.Zgroup     `json:"zgroups"`
		Pagination *modelext.Pagination `json:"pagination"`
	}
}

func adminZgroupsResponse(zgroups []*models.Zgroup, pagination *modelext.Pagination) *AdminZgroupsResponse {
	response := &AdminZgroupsResponse{}
	response.Body.Zgroups = zgroups
	response.Body.Pagination = pagination
	return response
}

// swagger:route GET /api/v1/admin/zgroups adminZgroups
// Returns a paginated list of zGroups.
// responses:
//   200: adminZgroupsResponse
func apiAdminZgroups(c buffalo.Context) error {
	zgroups, pagination, err := modelext.PaginatedZgroups(c)
	if err != nil {
		return c.Error(http.StatusInternalServerError, err)
	}
	return c.Render(http.StatusOK, r.JSON(adminZgroupsResponse(
		zgroups,
		pagination,
	).Body))
}

// AdminZgroupResponse returns the queried zgroups
// swagger:response adminZgroupResponse
type AdminZgroupResponse struct {
	//in:body
	Body struct {
		Zgroup *models.Zgroup `json:"zgroup"`
		Users  []*models.User `json:"users"`
	}
}

func adminZgroupResponse(zgroup *models.Zgroup, users []*models.User) *AdminZgroupResponse {
	response := &AdminZgroupResponse{}
	response.Body.Zgroup = zgroup
	response.Body.Users = users
	return response
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
func apiAdminZgroup(c buffalo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.Error(http.StatusInternalServerError, err)
	}
	zgroup, err := modelext.FindZgroup(c, id)
	if err != nil {
		return c.Error(http.StatusInternalServerError, err)
	}
	return c.Render(http.StatusOK, r.JSON(adminZgroupResponse(
		zgroup,
		zgroup.R.Users,
	).Body))
}
