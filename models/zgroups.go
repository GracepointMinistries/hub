// Code generated by SQLBoiler 4.4.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Zgroup is an object representing the database table.
type Zgroup struct {
	ID        int       `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name      string    `boil:"name" json:"name" toml:"name" yaml:"name"`
	Archived  bool      `boil:"archived" json:"archived" toml:"archived" yaml:"archived"`
	CreatedAt time.Time `boil:"created_at" json:"createdAt" toml:"createdAt" yaml:"createdAt"`

	R *zgroupR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L zgroupL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ZgroupColumns = struct {
	ID        string
	Name      string
	Archived  string
	CreatedAt string
}{
	ID:        "id",
	Name:      "name",
	Archived:  "archived",
	CreatedAt: "created_at",
}

// Generated where

var ZgroupWhere = struct {
	ID        whereHelperint
	Name      whereHelperstring
	Archived  whereHelperbool
	CreatedAt whereHelpertime_Time
}{
	ID:        whereHelperint{field: "\"zgroups\".\"id\""},
	Name:      whereHelperstring{field: "\"zgroups\".\"name\""},
	Archived:  whereHelperbool{field: "\"zgroups\".\"archived\""},
	CreatedAt: whereHelpertime_Time{field: "\"zgroups\".\"created_at\""},
}

// ZgroupRels is where relationship names are stored.
var ZgroupRels = struct {
	Users string
}{
	Users: "Users",
}

// zgroupR is where relationships are stored.
type zgroupR struct {
	Users UserSlice `boil:"Users" json:"Users" toml:"Users" yaml:"Users"`
}

// NewStruct creates a new relationship struct
func (*zgroupR) NewStruct() *zgroupR {
	return &zgroupR{}
}

// zgroupL is where Load methods for each relationship are stored.
type zgroupL struct{}

var (
	zgroupAllColumns            = []string{"id", "name", "archived", "created_at"}
	zgroupColumnsWithoutDefault = []string{"name"}
	zgroupColumnsWithDefault    = []string{"id", "archived", "created_at"}
	zgroupPrimaryKeyColumns     = []string{"id"}
)

type (
	// ZgroupSlice is an alias for a slice of pointers to Zgroup.
	// This should generally be used opposed to []Zgroup.
	ZgroupSlice []*Zgroup
	// ZgroupHook is the signature for custom Zgroup hook methods
	ZgroupHook func(context.Context, boil.ContextExecutor, *Zgroup) error

	zgroupQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	zgroupType                 = reflect.TypeOf(&Zgroup{})
	zgroupMapping              = queries.MakeStructMapping(zgroupType)
	zgroupPrimaryKeyMapping, _ = queries.BindMapping(zgroupType, zgroupMapping, zgroupPrimaryKeyColumns)
	zgroupInsertCacheMut       sync.RWMutex
	zgroupInsertCache          = make(map[string]insertCache)
	zgroupUpdateCacheMut       sync.RWMutex
	zgroupUpdateCache          = make(map[string]updateCache)
	zgroupUpsertCacheMut       sync.RWMutex
	zgroupUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var zgroupBeforeInsertHooks []ZgroupHook
var zgroupBeforeUpdateHooks []ZgroupHook
var zgroupBeforeDeleteHooks []ZgroupHook
var zgroupBeforeUpsertHooks []ZgroupHook

var zgroupAfterInsertHooks []ZgroupHook
var zgroupAfterSelectHooks []ZgroupHook
var zgroupAfterUpdateHooks []ZgroupHook
var zgroupAfterDeleteHooks []ZgroupHook
var zgroupAfterUpsertHooks []ZgroupHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Zgroup) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range zgroupBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Zgroup) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range zgroupBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Zgroup) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range zgroupBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Zgroup) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range zgroupBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Zgroup) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range zgroupAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Zgroup) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range zgroupAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Zgroup) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range zgroupAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Zgroup) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range zgroupAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Zgroup) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range zgroupAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddZgroupHook registers your hook function for all future operations.
func AddZgroupHook(hookPoint boil.HookPoint, zgroupHook ZgroupHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		zgroupBeforeInsertHooks = append(zgroupBeforeInsertHooks, zgroupHook)
	case boil.BeforeUpdateHook:
		zgroupBeforeUpdateHooks = append(zgroupBeforeUpdateHooks, zgroupHook)
	case boil.BeforeDeleteHook:
		zgroupBeforeDeleteHooks = append(zgroupBeforeDeleteHooks, zgroupHook)
	case boil.BeforeUpsertHook:
		zgroupBeforeUpsertHooks = append(zgroupBeforeUpsertHooks, zgroupHook)
	case boil.AfterInsertHook:
		zgroupAfterInsertHooks = append(zgroupAfterInsertHooks, zgroupHook)
	case boil.AfterSelectHook:
		zgroupAfterSelectHooks = append(zgroupAfterSelectHooks, zgroupHook)
	case boil.AfterUpdateHook:
		zgroupAfterUpdateHooks = append(zgroupAfterUpdateHooks, zgroupHook)
	case boil.AfterDeleteHook:
		zgroupAfterDeleteHooks = append(zgroupAfterDeleteHooks, zgroupHook)
	case boil.AfterUpsertHook:
		zgroupAfterUpsertHooks = append(zgroupAfterUpsertHooks, zgroupHook)
	}
}

// One returns a single zgroup record from the query.
func (q zgroupQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Zgroup, error) {
	o := &Zgroup{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for zgroups")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Zgroup records from the query.
func (q zgroupQuery) All(ctx context.Context, exec boil.ContextExecutor) (ZgroupSlice, error) {
	var o []*Zgroup

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Zgroup slice")
	}

	if len(zgroupAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Zgroup records in the query.
func (q zgroupQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count zgroups rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q zgroupQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if zgroups exists")
	}

	return count > 0, nil
}

// Users retrieves all the user's Users with an executor.
func (o *Zgroup) Users(mods ...qm.QueryMod) userQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.InnerJoin("\"zgroup_members\" on \"users\".\"id\" = \"zgroup_members\".\"user_id\""),
		qm.Where("\"zgroup_members\".\"zgroup_id\"=?", o.ID),
	)

	query := Users(queryMods...)
	queries.SetFrom(query.Query, "\"users\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"users\".*"})
	}

	return query
}

// LoadUsers allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (zgroupL) LoadUsers(ctx context.Context, e boil.ContextExecutor, singular bool, maybeZgroup interface{}, mods queries.Applicator) error {
	var slice []*Zgroup
	var object *Zgroup

	if singular {
		object = maybeZgroup.(*Zgroup)
	} else {
		slice = *maybeZgroup.(*[]*Zgroup)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &zgroupR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &zgroupR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.Select("\"users\".id, \"users\".name, \"users\".email, \"users\".blocked, \"users\".created_at, \"a\".\"zgroup_id\""),
		qm.From("\"users\""),
		qm.InnerJoin("\"zgroup_members\" as \"a\" on \"users\".\"id\" = \"a\".\"user_id\""),
		qm.WhereIn("\"a\".\"zgroup_id\" in ?", args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load users")
	}

	var resultSlice []*User

	var localJoinCols []int
	for results.Next() {
		one := new(User)
		var localJoinCol int

		err = results.Scan(&one.ID, &one.Name, &one.Email, &one.Blocked, &one.CreatedAt, &localJoinCol)
		if err != nil {
			return errors.Wrap(err, "failed to scan eager loaded results for users")
		}
		if err = results.Err(); err != nil {
			return errors.Wrap(err, "failed to plebian-bind eager loaded slice users")
		}

		resultSlice = append(resultSlice, one)
		localJoinCols = append(localJoinCols, localJoinCol)
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on users")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for users")
	}

	if len(userAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.Users = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &userR{}
			}
			foreign.R.Zgroups = append(foreign.R.Zgroups, object)
		}
		return nil
	}

	for i, foreign := range resultSlice {
		localJoinCol := localJoinCols[i]
		for _, local := range slice {
			if local.ID == localJoinCol {
				local.R.Users = append(local.R.Users, foreign)
				if foreign.R == nil {
					foreign.R = &userR{}
				}
				foreign.R.Zgroups = append(foreign.R.Zgroups, local)
				break
			}
		}
	}

	return nil
}

// AddUsers adds the given related objects to the existing relationships
// of the zgroup, optionally inserting them as new records.
// Appends related to o.R.Users.
// Sets related.R.Zgroups appropriately.
func (o *Zgroup) AddUsers(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*User) error {
	var err error
	for _, rel := range related {
		if insert {
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		}
	}

	for _, rel := range related {
		query := "insert into \"zgroup_members\" (\"zgroup_id\", \"user_id\") values ($1, $2)"
		values := []interface{}{o.ID, rel.ID}

		if boil.IsDebug(ctx) {
			writer := boil.DebugWriterFrom(ctx)
			fmt.Fprintln(writer, query)
			fmt.Fprintln(writer, values)
		}
		_, err = exec.ExecContext(ctx, query, values...)
		if err != nil {
			return errors.Wrap(err, "failed to insert into join table")
		}
	}
	if o.R == nil {
		o.R = &zgroupR{
			Users: related,
		}
	} else {
		o.R.Users = append(o.R.Users, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &userR{
				Zgroups: ZgroupSlice{o},
			}
		} else {
			rel.R.Zgroups = append(rel.R.Zgroups, o)
		}
	}
	return nil
}

// SetUsers removes all previously related items of the
// zgroup replacing them completely with the passed
// in related items, optionally inserting them as new records.
// Sets o.R.Zgroups's Users accordingly.
// Replaces o.R.Users with related.
// Sets related.R.Zgroups's Users accordingly.
func (o *Zgroup) SetUsers(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*User) error {
	query := "delete from \"zgroup_members\" where \"zgroup_id\" = $1"
	values := []interface{}{o.ID}
	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, query)
		fmt.Fprintln(writer, values)
	}
	_, err := exec.ExecContext(ctx, query, values...)
	if err != nil {
		return errors.Wrap(err, "failed to remove relationships before set")
	}

	removeUsersFromZgroupsSlice(o, related)
	if o.R != nil {
		o.R.Users = nil
	}
	return o.AddUsers(ctx, exec, insert, related...)
}

// RemoveUsers relationships from objects passed in.
// Removes related items from R.Users (uses pointer comparison, removal does not keep order)
// Sets related.R.Zgroups.
func (o *Zgroup) RemoveUsers(ctx context.Context, exec boil.ContextExecutor, related ...*User) error {
	var err error
	query := fmt.Sprintf(
		"delete from \"zgroup_members\" where \"zgroup_id\" = $1 and \"user_id\" in (%s)",
		strmangle.Placeholders(dialect.UseIndexPlaceholders, len(related), 2, 1),
	)
	values := []interface{}{o.ID}
	for _, rel := range related {
		values = append(values, rel.ID)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, query)
		fmt.Fprintln(writer, values)
	}
	_, err = exec.ExecContext(ctx, query, values...)
	if err != nil {
		return errors.Wrap(err, "failed to remove relationships before set")
	}
	removeUsersFromZgroupsSlice(o, related)
	if o.R == nil {
		return nil
	}

	for _, rel := range related {
		for i, ri := range o.R.Users {
			if rel != ri {
				continue
			}

			ln := len(o.R.Users)
			if ln > 1 && i < ln-1 {
				o.R.Users[i] = o.R.Users[ln-1]
			}
			o.R.Users = o.R.Users[:ln-1]
			break
		}
	}

	return nil
}

func removeUsersFromZgroupsSlice(o *Zgroup, related []*User) {
	for _, rel := range related {
		if rel.R == nil {
			continue
		}
		for i, ri := range rel.R.Zgroups {
			if o.ID != ri.ID {
				continue
			}

			ln := len(rel.R.Zgroups)
			if ln > 1 && i < ln-1 {
				rel.R.Zgroups[i] = rel.R.Zgroups[ln-1]
			}
			rel.R.Zgroups = rel.R.Zgroups[:ln-1]
			break
		}
	}
}

// Zgroups retrieves all the records using an executor.
func Zgroups(mods ...qm.QueryMod) zgroupQuery {
	mods = append(mods, qm.From("\"zgroups\""))
	return zgroupQuery{NewQuery(mods...)}
}

// FindZgroup retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindZgroup(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*Zgroup, error) {
	zgroupObj := &Zgroup{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"zgroups\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, zgroupObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from zgroups")
	}

	return zgroupObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Zgroup) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no zgroups provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(zgroupColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	zgroupInsertCacheMut.RLock()
	cache, cached := zgroupInsertCache[key]
	zgroupInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			zgroupAllColumns,
			zgroupColumnsWithDefault,
			zgroupColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(zgroupType, zgroupMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(zgroupType, zgroupMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"zgroups\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"zgroups\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into zgroups")
	}

	if !cached {
		zgroupInsertCacheMut.Lock()
		zgroupInsertCache[key] = cache
		zgroupInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Zgroup.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Zgroup) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	zgroupUpdateCacheMut.RLock()
	cache, cached := zgroupUpdateCache[key]
	zgroupUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			zgroupAllColumns,
			zgroupPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update zgroups, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"zgroups\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, zgroupPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(zgroupType, zgroupMapping, append(wl, zgroupPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update zgroups row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for zgroups")
	}

	if !cached {
		zgroupUpdateCacheMut.Lock()
		zgroupUpdateCache[key] = cache
		zgroupUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q zgroupQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for zgroups")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for zgroups")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ZgroupSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), zgroupPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"zgroups\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, zgroupPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in zgroup slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all zgroup")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Zgroup) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no zgroups provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(zgroupColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	zgroupUpsertCacheMut.RLock()
	cache, cached := zgroupUpsertCache[key]
	zgroupUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			zgroupAllColumns,
			zgroupColumnsWithDefault,
			zgroupColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			zgroupAllColumns,
			zgroupPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert zgroups, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(zgroupPrimaryKeyColumns))
			copy(conflict, zgroupPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"zgroups\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(zgroupType, zgroupMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(zgroupType, zgroupMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert zgroups")
	}

	if !cached {
		zgroupUpsertCacheMut.Lock()
		zgroupUpsertCache[key] = cache
		zgroupUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Zgroup record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Zgroup) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Zgroup provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), zgroupPrimaryKeyMapping)
	sql := "DELETE FROM \"zgroups\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from zgroups")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for zgroups")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q zgroupQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no zgroupQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from zgroups")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for zgroups")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ZgroupSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(zgroupBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), zgroupPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"zgroups\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, zgroupPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from zgroup slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for zgroups")
	}

	if len(zgroupAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Zgroup) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindZgroup(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ZgroupSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ZgroupSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), zgroupPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"zgroups\".* FROM \"zgroups\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, zgroupPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in ZgroupSlice")
	}

	*o = slice

	return nil
}

// ZgroupExists checks if the Zgroup row exists.
func ZgroupExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"zgroups\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if zgroups exists")
	}

	return exists, nil
}
