// Code generated by SQLBoiler 4.14.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package repository

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("Delegations", testDelegations)
	t.Run("DelegationsStats", testDelegationsStats)
	t.Run("Delegators", testDelegators)
}

func TestDelete(t *testing.T) {
	t.Run("Delegations", testDelegationsDelete)
	t.Run("DelegationsStats", testDelegationsStatsDelete)
	t.Run("Delegators", testDelegatorsDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("Delegations", testDelegationsQueryDeleteAll)
	t.Run("DelegationsStats", testDelegationsStatsQueryDeleteAll)
	t.Run("Delegators", testDelegatorsQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("Delegations", testDelegationsSliceDeleteAll)
	t.Run("DelegationsStats", testDelegationsStatsSliceDeleteAll)
	t.Run("Delegators", testDelegatorsSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("Delegations", testDelegationsExists)
	t.Run("DelegationsStats", testDelegationsStatsExists)
	t.Run("Delegators", testDelegatorsExists)
}

func TestFind(t *testing.T) {
	t.Run("Delegations", testDelegationsFind)
	t.Run("DelegationsStats", testDelegationsStatsFind)
	t.Run("Delegators", testDelegatorsFind)
}

func TestBind(t *testing.T) {
	t.Run("Delegations", testDelegationsBind)
	t.Run("DelegationsStats", testDelegationsStatsBind)
	t.Run("Delegators", testDelegatorsBind)
}

func TestOne(t *testing.T) {
	t.Run("Delegations", testDelegationsOne)
	t.Run("DelegationsStats", testDelegationsStatsOne)
	t.Run("Delegators", testDelegatorsOne)
}

func TestAll(t *testing.T) {
	t.Run("Delegations", testDelegationsAll)
	t.Run("DelegationsStats", testDelegationsStatsAll)
	t.Run("Delegators", testDelegatorsAll)
}

func TestCount(t *testing.T) {
	t.Run("Delegations", testDelegationsCount)
	t.Run("DelegationsStats", testDelegationsStatsCount)
	t.Run("Delegators", testDelegatorsCount)
}

func TestHooks(t *testing.T) {
	t.Run("Delegations", testDelegationsHooks)
	t.Run("DelegationsStats", testDelegationsStatsHooks)
	t.Run("Delegators", testDelegatorsHooks)
}

func TestInsert(t *testing.T) {
	t.Run("Delegations", testDelegationsInsert)
	t.Run("Delegations", testDelegationsInsertWhitelist)
	t.Run("DelegationsStats", testDelegationsStatsInsert)
	t.Run("DelegationsStats", testDelegationsStatsInsertWhitelist)
	t.Run("Delegators", testDelegatorsInsert)
	t.Run("Delegators", testDelegatorsInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {
	t.Run("DelegationToDelegatorUsingDelegationDelegator", testDelegationToOneDelegatorUsingDelegationDelegator)
	t.Run("DelegationsStatToDelegatorUsingDelegatorAddressDelegator", testDelegationsStatToOneDelegatorUsingDelegatorAddressDelegator)
}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {
	t.Run("DelegatorToDelegations", testDelegatorToManyDelegations)
	t.Run("DelegatorToDelegatorAddressDelegationsStats", testDelegatorToManyDelegatorAddressDelegationsStats)
}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {
	t.Run("DelegationToDelegatorUsingDelegations", testDelegationToOneSetOpDelegatorUsingDelegationDelegator)
	t.Run("DelegationsStatToDelegatorUsingDelegatorAddressDelegationsStats", testDelegationsStatToOneSetOpDelegatorUsingDelegatorAddressDelegator)
}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {
	t.Run("DelegationsStatToDelegatorUsingDelegatorAddressDelegationsStats", testDelegationsStatToOneRemoveOpDelegatorUsingDelegatorAddressDelegator)
}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {
	t.Run("DelegatorToDelegations", testDelegatorToManyAddOpDelegations)
	t.Run("DelegatorToDelegatorAddressDelegationsStats", testDelegatorToManyAddOpDelegatorAddressDelegationsStats)
}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {
	t.Run("DelegatorToDelegatorAddressDelegationsStats", testDelegatorToManySetOpDelegatorAddressDelegationsStats)
}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {
	t.Run("DelegatorToDelegatorAddressDelegationsStats", testDelegatorToManyRemoveOpDelegatorAddressDelegationsStats)
}

func TestReload(t *testing.T) {
	t.Run("Delegations", testDelegationsReload)
	t.Run("DelegationsStats", testDelegationsStatsReload)
	t.Run("Delegators", testDelegatorsReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("Delegations", testDelegationsReloadAll)
	t.Run("DelegationsStats", testDelegationsStatsReloadAll)
	t.Run("Delegators", testDelegatorsReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("Delegations", testDelegationsSelect)
	t.Run("DelegationsStats", testDelegationsStatsSelect)
	t.Run("Delegators", testDelegatorsSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("Delegations", testDelegationsUpdate)
	t.Run("DelegationsStats", testDelegationsStatsUpdate)
	t.Run("Delegators", testDelegatorsUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("Delegations", testDelegationsSliceUpdateAll)
	t.Run("DelegationsStats", testDelegationsStatsSliceUpdateAll)
	t.Run("Delegators", testDelegatorsSliceUpdateAll)
}
