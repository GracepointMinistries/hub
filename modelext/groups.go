package modelext

import (
	"database/sql"
	"errors"
	"log"

	"github.com/GracepointMinistries/hub/models"
	"github.com/gobuffalo/buffalo"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// PaginatedGroups adds pagination to a multi-group query
func PaginatedGroups(c buffalo.Context, queries ...qm.QueryMod) ([]*models.Group, *Pagination, error) {
	pagination, clauses := getPagination(c)
	clauses = append(clauses, queries...)
	groups, err := models.Groups(
		clauses...,
	).All(c, GetTx(c))
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		log.Println(err)
		return nil, nil, err
	}
	pagination.Cursor = -1
	if len(groups) > 0 {
		pagination.Cursor = groups[len(groups)-1].ID
	}
	return groups, pagination, nil
}

// FindGroup finds the group and eager-loads its users
func FindGroup(c buffalo.Context, id int) (*models.Group, error) {
	return models.Groups(models.GroupWhere.ID.EQ(id), qm.Load(models.GroupRels.Users)).One(c, GetTx(c))
}

// GroupForUser returns the associated group of a user
func GroupForUser(user *models.User) *models.Group {
	groupsLen := len(user.R.Groups)
	if groupsLen == 0 {
		return nil
	}
	return user.R.Groups[groupsLen-1]
}

// GroupStatus returns a human readable status for the Group
func GroupStatus(group *models.Group) string {
	if group.Archived {
		return "Archived"
	}
	return "Active"
}

// TotalUsersIn returns the count of users in a group
func TotalUsersIn(c buffalo.Context, group *models.Group) (int64, error) {
	return group.Users().Count(c, GetTx(c))
}

// HasZoomLink returns whether the group has a published zoom link
func HasZoomLink(group *models.Group) bool {
	return group.Published && group.ZoomLink != ""
}
