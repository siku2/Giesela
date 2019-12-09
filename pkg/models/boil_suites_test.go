// Code generated by SQLBoiler 3.6.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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
	t.Run("RoleTargets", testRoleTargets)
	t.Run("Roles", testRoles)
}

func TestDelete(t *testing.T) {
	t.Run("RoleTargets", testRoleTargetsDelete)
	t.Run("Roles", testRolesDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("RoleTargets", testRoleTargetsQueryDeleteAll)
	t.Run("Roles", testRolesQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("RoleTargets", testRoleTargetsSliceDeleteAll)
	t.Run("Roles", testRolesSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("RoleTargets", testRoleTargetsExists)
	t.Run("Roles", testRolesExists)
}

func TestFind(t *testing.T) {
	t.Run("RoleTargets", testRoleTargetsFind)
	t.Run("Roles", testRolesFind)
}

func TestBind(t *testing.T) {
	t.Run("RoleTargets", testRoleTargetsBind)
	t.Run("Roles", testRolesBind)
}

func TestOne(t *testing.T) {
	t.Run("RoleTargets", testRoleTargetsOne)
	t.Run("Roles", testRolesOne)
}

func TestAll(t *testing.T) {
	t.Run("RoleTargets", testRoleTargetsAll)
	t.Run("Roles", testRolesAll)
}

func TestCount(t *testing.T) {
	t.Run("RoleTargets", testRoleTargetsCount)
	t.Run("Roles", testRolesCount)
}

func TestHooks(t *testing.T) {
	t.Run("RoleTargets", testRoleTargetsHooks)
	t.Run("Roles", testRolesHooks)
}

func TestInsert(t *testing.T) {
	t.Run("RoleTargets", testRoleTargetsInsert)
	t.Run("RoleTargets", testRoleTargetsInsertWhitelist)
	t.Run("Roles", testRolesInsert)
	t.Run("Roles", testRolesInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {
	t.Run("RoleTargetToRoleUsingRole", testRoleTargetToOneRoleUsingRole)
}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {
	t.Run("RoleToRoleTargets", testRoleToManyRoleTargets)
}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {
	t.Run("RoleTargetToRoleUsingRoleTargets", testRoleTargetToOneSetOpRoleUsingRole)
}

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
	t.Run("RoleToRoleTargets", testRoleToManyAddOpRoleTargets)
}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {}

func TestReload(t *testing.T) {
	t.Run("RoleTargets", testRoleTargetsReload)
	t.Run("Roles", testRolesReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("RoleTargets", testRoleTargetsReloadAll)
	t.Run("Roles", testRolesReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("RoleTargets", testRoleTargetsSelect)
	t.Run("Roles", testRolesSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("RoleTargets", testRoleTargetsUpdate)
	t.Run("Roles", testRolesUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("RoleTargets", testRoleTargetsSliceUpdateAll)
	t.Run("Roles", testRolesSliceUpdateAll)
}