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

// Oauth is an object representing the database table.
type Oauth struct {
	ID         int       `boil:"id" json:"id" toml:"id" yaml:"id"`
	Provider   string    `boil:"provider" json:"provider" toml:"provider" yaml:"provider"`
	ProviderID string    `boil:"provider_id" json:"providerID" toml:"providerID" yaml:"providerID"`
	UpdatedAt  time.Time `boil:"updated_at" json:"updatedAt" toml:"updatedAt" yaml:"updatedAt"`
	CreatedAt  time.Time `boil:"created_at" json:"createdAt" toml:"createdAt" yaml:"createdAt"`

	R *oauthR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L oauthL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var OauthColumns = struct {
	ID         string
	Provider   string
	ProviderID string
	UpdatedAt  string
	CreatedAt  string
}{
	ID:         "id",
	Provider:   "provider",
	ProviderID: "provider_id",
	UpdatedAt:  "updated_at",
	CreatedAt:  "created_at",
}

// Generated where

var OauthWhere = struct {
	ID         whereHelperint
	Provider   whereHelperstring
	ProviderID whereHelperstring
	UpdatedAt  whereHelpertime_Time
	CreatedAt  whereHelpertime_Time
}{
	ID:         whereHelperint{field: "\"oauths\".\"id\""},
	Provider:   whereHelperstring{field: "\"oauths\".\"provider\""},
	ProviderID: whereHelperstring{field: "\"oauths\".\"provider_id\""},
	UpdatedAt:  whereHelpertime_Time{field: "\"oauths\".\"updated_at\""},
	CreatedAt:  whereHelpertime_Time{field: "\"oauths\".\"created_at\""},
}

// OauthRels is where relationship names are stored.
var OauthRels = struct {
	Users string
}{
	Users: "Users",
}

// oauthR is where relationships are stored.
type oauthR struct {
	Users UserSlice `boil:"Users" json:"Users" toml:"Users" yaml:"Users"`
}

// NewStruct creates a new relationship struct
func (*oauthR) NewStruct() *oauthR {
	return &oauthR{}
}

// oauthL is where Load methods for each relationship are stored.
type oauthL struct{}

var (
	oauthAllColumns            = []string{"id", "provider", "provider_id", "updated_at", "created_at"}
	oauthColumnsWithoutDefault = []string{"provider", "provider_id"}
	oauthColumnsWithDefault    = []string{"id", "updated_at", "created_at"}
	oauthPrimaryKeyColumns     = []string{"id"}
)

type (
	// OauthSlice is an alias for a slice of pointers to Oauth.
	// This should generally be used opposed to []Oauth.
	OauthSlice []*Oauth

	oauthQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	oauthType                 = reflect.TypeOf(&Oauth{})
	oauthMapping              = queries.MakeStructMapping(oauthType)
	oauthPrimaryKeyMapping, _ = queries.BindMapping(oauthType, oauthMapping, oauthPrimaryKeyColumns)
	oauthInsertCacheMut       sync.RWMutex
	oauthInsertCache          = make(map[string]insertCache)
	oauthUpdateCacheMut       sync.RWMutex
	oauthUpdateCache          = make(map[string]updateCache)
	oauthUpsertCacheMut       sync.RWMutex
	oauthUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single oauth record from the query.
func (q oauthQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Oauth, error) {
	o := &Oauth{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for oauths")
	}

	return o, nil
}

// All returns all Oauth records from the query.
func (q oauthQuery) All(ctx context.Context, exec boil.ContextExecutor) (OauthSlice, error) {
	var o []*Oauth

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Oauth slice")
	}

	return o, nil
}

// Count returns the count of all Oauth records in the query.
func (q oauthQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count oauths rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q oauthQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if oauths exists")
	}

	return count > 0, nil
}

// Users retrieves all the user's Users with an executor.
func (o *Oauth) Users(mods ...qm.QueryMod) userQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.InnerJoin("\"oauth_users\" on \"users\".\"id\" = \"oauth_users\".\"user_id\""),
		qm.Where("\"oauth_users\".\"oauth_id\"=?", o.ID),
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
func (oauthL) LoadUsers(ctx context.Context, e boil.ContextExecutor, singular bool, maybeOauth interface{}, mods queries.Applicator) error {
	var slice []*Oauth
	var object *Oauth

	if singular {
		object = maybeOauth.(*Oauth)
	} else {
		slice = *maybeOauth.(*[]*Oauth)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &oauthR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &oauthR{}
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
		qm.Select("\"users\".id, \"users\".name, \"users\".email, \"users\".blocked, \"users\".updated_at, \"users\".created_at, \"a\".\"oauth_id\""),
		qm.From("\"users\""),
		qm.InnerJoin("\"oauth_users\" as \"a\" on \"users\".\"id\" = \"a\".\"user_id\""),
		qm.WhereIn("\"a\".\"oauth_id\" in ?", args...),
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

		err = results.Scan(&one.ID, &one.Name, &one.Email, &one.Blocked, &one.UpdatedAt, &one.CreatedAt, &localJoinCol)
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

	if singular {
		object.R.Users = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &userR{}
			}
			foreign.R.Oauths = append(foreign.R.Oauths, object)
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
				foreign.R.Oauths = append(foreign.R.Oauths, local)
				break
			}
		}
	}

	return nil
}

// AddUsers adds the given related objects to the existing relationships
// of the oauth, optionally inserting them as new records.
// Appends related to o.R.Users.
// Sets related.R.Oauths appropriately.
func (o *Oauth) AddUsers(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*User) error {
	var err error
	for _, rel := range related {
		if insert {
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		}
	}

	for _, rel := range related {
		query := "insert into \"oauth_users\" (\"oauth_id\", \"user_id\") values ($1, $2)"
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
		o.R = &oauthR{
			Users: related,
		}
	} else {
		o.R.Users = append(o.R.Users, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &userR{
				Oauths: OauthSlice{o},
			}
		} else {
			rel.R.Oauths = append(rel.R.Oauths, o)
		}
	}
	return nil
}

// SetUsers removes all previously related items of the
// oauth replacing them completely with the passed
// in related items, optionally inserting them as new records.
// Sets o.R.Oauths's Users accordingly.
// Replaces o.R.Users with related.
// Sets related.R.Oauths's Users accordingly.
func (o *Oauth) SetUsers(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*User) error {
	query := "delete from \"oauth_users\" where \"oauth_id\" = $1"
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

	removeUsersFromOauthsSlice(o, related)
	if o.R != nil {
		o.R.Users = nil
	}
	return o.AddUsers(ctx, exec, insert, related...)
}

// RemoveUsers relationships from objects passed in.
// Removes related items from R.Users (uses pointer comparison, removal does not keep order)
// Sets related.R.Oauths.
func (o *Oauth) RemoveUsers(ctx context.Context, exec boil.ContextExecutor, related ...*User) error {
	var err error
	query := fmt.Sprintf(
		"delete from \"oauth_users\" where \"oauth_id\" = $1 and \"user_id\" in (%s)",
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
	removeUsersFromOauthsSlice(o, related)
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

func removeUsersFromOauthsSlice(o *Oauth, related []*User) {
	for _, rel := range related {
		if rel.R == nil {
			continue
		}
		for i, ri := range rel.R.Oauths {
			if o.ID != ri.ID {
				continue
			}

			ln := len(rel.R.Oauths)
			if ln > 1 && i < ln-1 {
				rel.R.Oauths[i] = rel.R.Oauths[ln-1]
			}
			rel.R.Oauths = rel.R.Oauths[:ln-1]
			break
		}
	}
}

// Oauths retrieves all the records using an executor.
func Oauths(mods ...qm.QueryMod) oauthQuery {
	mods = append(mods, qm.From("\"oauths\""))
	return oauthQuery{NewQuery(mods...)}
}

// FindOauth retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindOauth(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*Oauth, error) {
	oauthObj := &Oauth{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"oauths\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, oauthObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from oauths")
	}

	return oauthObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Oauth) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no oauths provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.UpdatedAt.IsZero() {
			o.UpdatedAt = currTime
		}
		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	nzDefaults := queries.NonZeroDefaultSet(oauthColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	oauthInsertCacheMut.RLock()
	cache, cached := oauthInsertCache[key]
	oauthInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			oauthAllColumns,
			oauthColumnsWithDefault,
			oauthColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(oauthType, oauthMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(oauthType, oauthMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"oauths\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"oauths\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into oauths")
	}

	if !cached {
		oauthInsertCacheMut.Lock()
		oauthInsertCache[key] = cache
		oauthInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the Oauth.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Oauth) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	key := makeCacheKey(columns, nil)
	oauthUpdateCacheMut.RLock()
	cache, cached := oauthUpdateCache[key]
	oauthUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			oauthAllColumns,
			oauthPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update oauths, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"oauths\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, oauthPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(oauthType, oauthMapping, append(wl, oauthPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update oauths row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for oauths")
	}

	if !cached {
		oauthUpdateCacheMut.Lock()
		oauthUpdateCache[key] = cache
		oauthUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q oauthQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for oauths")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for oauths")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o OauthSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), oauthPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"oauths\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, oauthPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in oauth slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all oauth")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Oauth) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no oauths provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	nzDefaults := queries.NonZeroDefaultSet(oauthColumnsWithDefault, o)

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

	oauthUpsertCacheMut.RLock()
	cache, cached := oauthUpsertCache[key]
	oauthUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			oauthAllColumns,
			oauthColumnsWithDefault,
			oauthColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			oauthAllColumns,
			oauthPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert oauths, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(oauthPrimaryKeyColumns))
			copy(conflict, oauthPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"oauths\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(oauthType, oauthMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(oauthType, oauthMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert oauths")
	}

	if !cached {
		oauthUpsertCacheMut.Lock()
		oauthUpsertCache[key] = cache
		oauthUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single Oauth record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Oauth) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Oauth provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), oauthPrimaryKeyMapping)
	sql := "DELETE FROM \"oauths\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from oauths")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for oauths")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q oauthQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no oauthQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from oauths")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for oauths")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o OauthSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), oauthPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"oauths\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, oauthPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from oauth slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for oauths")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Oauth) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindOauth(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *OauthSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := OauthSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), oauthPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"oauths\".* FROM \"oauths\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, oauthPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in OauthSlice")
	}

	*o = slice

	return nil
}

// OauthExists checks if the Oauth row exists.
func OauthExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"oauths\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if oauths exists")
	}

	return exists, nil
}
