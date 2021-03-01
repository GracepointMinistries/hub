package models

import (
	"context"
	"time"

	"github.com/gobuffalo/pop"
)

// ZGroup represents a zgroup
type ZGroup struct {
	ID        int       `db:"id" json:"id"`
	Name      string    `db:"name" json:"name"`
	Archived  string    `db:"archived" json:"archived"`
	CreatedAt time.Time `db:"created_at" json:"createdAt"`
}

// GetZGroup returns a zgroup with the given id
func GetZGroup(ctx context.Context, id int) (*ZGroup, error) {
	zgroup := &ZGroup{}
	db := ctx.Value("tx").(*pop.Connection)
	err := db.RawQuery(`SELECT * FROM zgroups WHERE id = ?`, id).First(zgroup)
	if err != nil {
		return nil, err
	}
	return zgroup, nil
}

// GetZGroups returns a paginated list of zgroups
func GetZGroups(ctx context.Context, includeArchived bool, pagination *Pagination) ([]*ZGroup, *Pagination, error) {
	zgroups := make([]*ZGroup, 0)
	db := ctx.Value("tx").(*pop.Connection)
	query := `SELECT * FROM zgroups WHERE id > ? AND archived = FALSE ORDER BY id ASC LIMIT ?`
	if includeArchived {
		query = `SELECT * FROM zgroups WHERE id > ? ORDER BY id ASC LIMIT ?`
	}
	err := db.RawQuery(query, pagination.Cursor, pagination.Limit).All(&zgroups)
	if err != nil {
		return nil, nil, err
	}
	next := &Pagination{
		Limit:  pagination.Limit,
		Cursor: pagination.Cursor,
	}
	if len(zgroups) > 0 {
		next.Cursor = zgroups[len(zgroups)-1].ID
	}

	return zgroups, next, nil
}
