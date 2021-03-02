// Code generated by SQLBoiler 4.4.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testZgroups(t *testing.T) {
	t.Parallel()

	query := Zgroups()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testZgroupsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Zgroup{}
	if err = randomize.Struct(seed, o, zgroupDBTypes, true, zgroupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Zgroup struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Zgroups().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testZgroupsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Zgroup{}
	if err = randomize.Struct(seed, o, zgroupDBTypes, true, zgroupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Zgroup struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Zgroups().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Zgroups().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testZgroupsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Zgroup{}
	if err = randomize.Struct(seed, o, zgroupDBTypes, true, zgroupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Zgroup struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ZgroupSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Zgroups().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testZgroupsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Zgroup{}
	if err = randomize.Struct(seed, o, zgroupDBTypes, true, zgroupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Zgroup struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := ZgroupExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Zgroup exists: %s", err)
	}
	if !e {
		t.Errorf("Expected ZgroupExists to return true, but got false.")
	}
}

func testZgroupsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Zgroup{}
	if err = randomize.Struct(seed, o, zgroupDBTypes, true, zgroupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Zgroup struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	zgroupFound, err := FindZgroup(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if zgroupFound == nil {
		t.Error("want a record, got nil")
	}
}

func testZgroupsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Zgroup{}
	if err = randomize.Struct(seed, o, zgroupDBTypes, true, zgroupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Zgroup struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Zgroups().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testZgroupsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Zgroup{}
	if err = randomize.Struct(seed, o, zgroupDBTypes, true, zgroupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Zgroup struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Zgroups().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testZgroupsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	zgroupOne := &Zgroup{}
	zgroupTwo := &Zgroup{}
	if err = randomize.Struct(seed, zgroupOne, zgroupDBTypes, false, zgroupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Zgroup struct: %s", err)
	}
	if err = randomize.Struct(seed, zgroupTwo, zgroupDBTypes, false, zgroupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Zgroup struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = zgroupOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = zgroupTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Zgroups().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testZgroupsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	zgroupOne := &Zgroup{}
	zgroupTwo := &Zgroup{}
	if err = randomize.Struct(seed, zgroupOne, zgroupDBTypes, false, zgroupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Zgroup struct: %s", err)
	}
	if err = randomize.Struct(seed, zgroupTwo, zgroupDBTypes, false, zgroupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Zgroup struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = zgroupOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = zgroupTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Zgroups().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testZgroupsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Zgroup{}
	if err = randomize.Struct(seed, o, zgroupDBTypes, true, zgroupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Zgroup struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Zgroups().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testZgroupsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Zgroup{}
	if err = randomize.Struct(seed, o, zgroupDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Zgroup struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(zgroupColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Zgroups().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testZgroupToManyUsers(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Zgroup
	var b, c User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, zgroupDBTypes, true, zgroupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Zgroup struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, userDBTypes, false, userColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, userDBTypes, false, userColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	_, err = tx.Exec("insert into \"zgroup_members\" (\"zgroup_id\", \"user_id\") values ($1, $2)", a.ID, b.ID)
	if err != nil {
		t.Fatal(err)
	}
	_, err = tx.Exec("insert into \"zgroup_members\" (\"zgroup_id\", \"user_id\") values ($1, $2)", a.ID, c.ID)
	if err != nil {
		t.Fatal(err)
	}

	check, err := a.Users().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.ID == b.ID {
			bFound = true
		}
		if v.ID == c.ID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := ZgroupSlice{&a}
	if err = a.L.LoadUsers(ctx, tx, false, (*[]*Zgroup)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Users); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Users = nil
	if err = a.L.LoadUsers(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Users); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testZgroupToManyAddOpUsers(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Zgroup
	var b, c, d, e User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, zgroupDBTypes, false, strmangle.SetComplement(zgroupPrimaryKeyColumns, zgroupColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*User{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*User{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddUsers(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if first.R.Zgroups[0] != &a {
			t.Error("relationship was not added properly to the slice")
		}
		if second.R.Zgroups[0] != &a {
			t.Error("relationship was not added properly to the slice")
		}

		if a.R.Users[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Users[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Users().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testZgroupToManySetOpUsers(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Zgroup
	var b, c, d, e User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, zgroupDBTypes, false, strmangle.SetComplement(zgroupPrimaryKeyColumns, zgroupColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*User{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.SetUsers(ctx, tx, false, &b, &c)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.Users().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	err = a.SetUsers(ctx, tx, true, &d, &e)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.Users().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	// The following checks cannot be implemented since we have no handle
	// to these when we call Set(). Leaving them here as wishful thinking
	// and to let people know there's dragons.
	//
	// if len(b.R.Zgroups) != 0 {
	// 	t.Error("relationship was not removed properly from the slice")
	// }
	// if len(c.R.Zgroups) != 0 {
	// 	t.Error("relationship was not removed properly from the slice")
	// }
	if d.R.Zgroups[0] != &a {
		t.Error("relationship was not added properly to the slice")
	}
	if e.R.Zgroups[0] != &a {
		t.Error("relationship was not added properly to the slice")
	}

	if a.R.Users[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.Users[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testZgroupToManyRemoveOpUsers(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Zgroup
	var b, c, d, e User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, zgroupDBTypes, false, strmangle.SetComplement(zgroupPrimaryKeyColumns, zgroupColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*User{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.AddUsers(ctx, tx, true, foreigners...)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.Users().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 4 {
		t.Error("count was wrong:", count)
	}

	err = a.RemoveUsers(ctx, tx, foreigners[:2]...)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.Users().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if len(b.R.Zgroups) != 0 {
		t.Error("relationship was not removed properly from the slice")
	}
	if len(c.R.Zgroups) != 0 {
		t.Error("relationship was not removed properly from the slice")
	}
	if d.R.Zgroups[0] != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.Zgroups[0] != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}

	if len(a.R.Users) != 2 {
		t.Error("should have preserved two relationships")
	}

	// Removal doesn't do a stable deletion for performance so we have to flip the order
	if a.R.Users[1] != &d {
		t.Error("relationship to d should have been preserved")
	}
	if a.R.Users[0] != &e {
		t.Error("relationship to e should have been preserved")
	}
}

func testZgroupsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Zgroup{}
	if err = randomize.Struct(seed, o, zgroupDBTypes, true, zgroupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Zgroup struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testZgroupsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Zgroup{}
	if err = randomize.Struct(seed, o, zgroupDBTypes, true, zgroupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Zgroup struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ZgroupSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testZgroupsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Zgroup{}
	if err = randomize.Struct(seed, o, zgroupDBTypes, true, zgroupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Zgroup struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Zgroups().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	zgroupDBTypes = map[string]string{`ID`: `integer`, `Name`: `character varying`, `ZoomLink`: `character varying`, `Archived`: `boolean`, `UpdatedAt`: `timestamp with time zone`, `CreatedAt`: `timestamp with time zone`}
	_             = bytes.MinRead
)

func testZgroupsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(zgroupPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(zgroupAllColumns) == len(zgroupPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Zgroup{}
	if err = randomize.Struct(seed, o, zgroupDBTypes, true, zgroupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Zgroup struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Zgroups().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, zgroupDBTypes, true, zgroupPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Zgroup struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testZgroupsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(zgroupAllColumns) == len(zgroupPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Zgroup{}
	if err = randomize.Struct(seed, o, zgroupDBTypes, true, zgroupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Zgroup struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Zgroups().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, zgroupDBTypes, true, zgroupPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Zgroup struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(zgroupAllColumns, zgroupPrimaryKeyColumns) {
		fields = zgroupAllColumns
	} else {
		fields = strmangle.SetComplement(
			zgroupAllColumns,
			zgroupPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := ZgroupSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testZgroupsUpsert(t *testing.T) {
	t.Parallel()

	if len(zgroupAllColumns) == len(zgroupPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Zgroup{}
	if err = randomize.Struct(seed, &o, zgroupDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Zgroup struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Zgroup: %s", err)
	}

	count, err := Zgroups().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, zgroupDBTypes, false, zgroupPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Zgroup struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Zgroup: %s", err)
	}

	count, err = Zgroups().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
