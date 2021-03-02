// Code generated by SQLBoiler 4.4.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("Oauths", testOauths)
	t.Run("Users", testUsers)
	t.Run("Zgroups", testZgroups)
}

func TestDelete(t *testing.T) {
	t.Run("Oauths", testOauthsDelete)
	t.Run("Users", testUsersDelete)
	t.Run("Zgroups", testZgroupsDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("Oauths", testOauthsQueryDeleteAll)
	t.Run("Users", testUsersQueryDeleteAll)
	t.Run("Zgroups", testZgroupsQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("Oauths", testOauthsSliceDeleteAll)
	t.Run("Users", testUsersSliceDeleteAll)
	t.Run("Zgroups", testZgroupsSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("Oauths", testOauthsExists)
	t.Run("Users", testUsersExists)
	t.Run("Zgroups", testZgroupsExists)
}

func TestFind(t *testing.T) {
	t.Run("Oauths", testOauthsFind)
	t.Run("Users", testUsersFind)
	t.Run("Zgroups", testZgroupsFind)
}

func TestBind(t *testing.T) {
	t.Run("Oauths", testOauthsBind)
	t.Run("Users", testUsersBind)
	t.Run("Zgroups", testZgroupsBind)
}

func TestOne(t *testing.T) {
	t.Run("Oauths", testOauthsOne)
	t.Run("Users", testUsersOne)
	t.Run("Zgroups", testZgroupsOne)
}

func TestAll(t *testing.T) {
	t.Run("Oauths", testOauthsAll)
	t.Run("Users", testUsersAll)
	t.Run("Zgroups", testZgroupsAll)
}

func TestCount(t *testing.T) {
	t.Run("Oauths", testOauthsCount)
	t.Run("Users", testUsersCount)
	t.Run("Zgroups", testZgroupsCount)
}

func TestHooks(t *testing.T) {
	t.Run("Oauths", testOauthsHooks)
	t.Run("Users", testUsersHooks)
	t.Run("Zgroups", testZgroupsHooks)
}

func TestInsert(t *testing.T) {
	t.Run("Oauths", testOauthsInsert)
	t.Run("Oauths", testOauthsInsertWhitelist)
	t.Run("Users", testUsersInsert)
	t.Run("Users", testUsersInsertWhitelist)
	t.Run("Zgroups", testZgroupsInsert)
	t.Run("Zgroups", testZgroupsInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {
	t.Run("OauthToUsers", testOauthToManyUsers)
	t.Run("UserToOauths", testUserToManyOauths)
	t.Run("UserToZgroups", testUserToManyZgroups)
	t.Run("ZgroupToUsers", testZgroupToManyUsers)
}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {
	t.Run("OauthToUsers", testOauthToManyAddOpUsers)
	t.Run("UserToOauths", testUserToManyAddOpOauths)
	t.Run("UserToZgroups", testUserToManyAddOpZgroups)
	t.Run("ZgroupToUsers", testZgroupToManyAddOpUsers)
}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {
	t.Run("OauthToUsers", testOauthToManySetOpUsers)
	t.Run("UserToOauths", testUserToManySetOpOauths)
	t.Run("UserToZgroups", testUserToManySetOpZgroups)
	t.Run("ZgroupToUsers", testZgroupToManySetOpUsers)
}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {
	t.Run("OauthToUsers", testOauthToManyRemoveOpUsers)
	t.Run("UserToOauths", testUserToManyRemoveOpOauths)
	t.Run("UserToZgroups", testUserToManyRemoveOpZgroups)
	t.Run("ZgroupToUsers", testZgroupToManyRemoveOpUsers)
}

func TestReload(t *testing.T) {
	t.Run("Oauths", testOauthsReload)
	t.Run("Users", testUsersReload)
	t.Run("Zgroups", testZgroupsReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("Oauths", testOauthsReloadAll)
	t.Run("Users", testUsersReloadAll)
	t.Run("Zgroups", testZgroupsReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("Oauths", testOauthsSelect)
	t.Run("Users", testUsersSelect)
	t.Run("Zgroups", testZgroupsSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("Oauths", testOauthsUpdate)
	t.Run("Users", testUsersUpdate)
	t.Run("Zgroups", testZgroupsUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("Oauths", testOauthsSliceUpdateAll)
	t.Run("Users", testUsersSliceUpdateAll)
	t.Run("Zgroups", testZgroupsSliceUpdateAll)
}