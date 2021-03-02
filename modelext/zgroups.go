package modelext

import (
	"github.com/GracepointMinistries/hub/models"
	"github.com/gobuffalo/buffalo"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// PaginatedZgroups adds pagination to a multi-zgroup query
func PaginatedZgroups(c buffalo.Context, queries ...qm.QueryMod) ([]*models.Zgroup, *Pagination, error) {
	pagination, clauses := getPagination(c)
	clauses = append(clauses, queries...)
	zgroups, err := models.Zgroups(
		clauses...,
	).All(c, getTx(c))
	if err != nil {
		return nil, nil, err
	}
	pagination.Cursor = -1
	if len(zgroups) > 0 {
		pagination.Cursor = zgroups[len(zgroups)-1].ID
	}
	return zgroups, pagination, nil
}

// FindZgroup finds the zgroup and eager-loads its users
func FindZgroup(c buffalo.Context, id int) (*models.Zgroup, error) {
	return models.Zgroups(models.ZgroupWhere.ID.EQ(id), qm.Load(models.ZgroupRels.Users)).One(c, getTx(c))
}

// ZgroupForUser returns the associated Zgroup of a user
func ZgroupForUser(user *models.User) *models.Zgroup {
	zgroupsLen := len(user.R.Zgroups)
	if zgroupsLen == 0 {
		return nil
	}
	return user.R.Zgroups[zgroupsLen-1]
}

// ZgroupStatus returns a human readable status for the Zgroup
func ZgroupStatus(zgroup *models.Zgroup) string {
	if zgroup.Archived {
		return "Archived"
	}
	return "Active"
}

// TotalUsersIn returns the count of users in a Zgroup
func TotalUsersIn(c buffalo.Context, zgroup *models.Zgroup) (int64, error) {
	return zgroup.Users().Count(c, getTx(c))
}

// HasZoomLink returns whether the zgroup has a published zoom link
func HasZoomLink(zgroup *models.Zgroup) bool {
	return zgroup.Published && zgroup.ZoomLink != ""
}
