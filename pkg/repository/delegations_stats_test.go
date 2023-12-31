// Code generated by SQLBoiler 4.14.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package repository

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

func testDelegationsStats(t *testing.T) {
	t.Parallel()

	query := DelegationsStats()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testDelegationsStatsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DelegationsStat{}
	if err = randomize.Struct(seed, o, delegationsStatDBTypes, true, delegationsStatColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DelegationsStat struct: %s", err)
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

	count, err := DelegationsStats().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDelegationsStatsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DelegationsStat{}
	if err = randomize.Struct(seed, o, delegationsStatDBTypes, true, delegationsStatColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DelegationsStat struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := DelegationsStats().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := DelegationsStats().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDelegationsStatsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DelegationsStat{}
	if err = randomize.Struct(seed, o, delegationsStatDBTypes, true, delegationsStatColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DelegationsStat struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := DelegationsStatSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := DelegationsStats().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDelegationsStatsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DelegationsStat{}
	if err = randomize.Struct(seed, o, delegationsStatDBTypes, true, delegationsStatColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DelegationsStat struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := DelegationsStatExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if DelegationsStat exists: %s", err)
	}
	if !e {
		t.Errorf("Expected DelegationsStatExists to return true, but got false.")
	}
}

func testDelegationsStatsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DelegationsStat{}
	if err = randomize.Struct(seed, o, delegationsStatDBTypes, true, delegationsStatColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DelegationsStat struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	delegationsStatFound, err := FindDelegationsStat(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if delegationsStatFound == nil {
		t.Error("want a record, got nil")
	}
}

func testDelegationsStatsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DelegationsStat{}
	if err = randomize.Struct(seed, o, delegationsStatDBTypes, true, delegationsStatColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DelegationsStat struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = DelegationsStats().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testDelegationsStatsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DelegationsStat{}
	if err = randomize.Struct(seed, o, delegationsStatDBTypes, true, delegationsStatColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DelegationsStat struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := DelegationsStats().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testDelegationsStatsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	delegationsStatOne := &DelegationsStat{}
	delegationsStatTwo := &DelegationsStat{}
	if err = randomize.Struct(seed, delegationsStatOne, delegationsStatDBTypes, false, delegationsStatColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DelegationsStat struct: %s", err)
	}
	if err = randomize.Struct(seed, delegationsStatTwo, delegationsStatDBTypes, false, delegationsStatColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DelegationsStat struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = delegationsStatOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = delegationsStatTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := DelegationsStats().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testDelegationsStatsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	delegationsStatOne := &DelegationsStat{}
	delegationsStatTwo := &DelegationsStat{}
	if err = randomize.Struct(seed, delegationsStatOne, delegationsStatDBTypes, false, delegationsStatColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DelegationsStat struct: %s", err)
	}
	if err = randomize.Struct(seed, delegationsStatTwo, delegationsStatDBTypes, false, delegationsStatColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DelegationsStat struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = delegationsStatOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = delegationsStatTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DelegationsStats().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func delegationsStatBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *DelegationsStat) error {
	*o = DelegationsStat{}
	return nil
}

func delegationsStatAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *DelegationsStat) error {
	*o = DelegationsStat{}
	return nil
}

func delegationsStatAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *DelegationsStat) error {
	*o = DelegationsStat{}
	return nil
}

func delegationsStatBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *DelegationsStat) error {
	*o = DelegationsStat{}
	return nil
}

func delegationsStatAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *DelegationsStat) error {
	*o = DelegationsStat{}
	return nil
}

func delegationsStatBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *DelegationsStat) error {
	*o = DelegationsStat{}
	return nil
}

func delegationsStatAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *DelegationsStat) error {
	*o = DelegationsStat{}
	return nil
}

func delegationsStatBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *DelegationsStat) error {
	*o = DelegationsStat{}
	return nil
}

func delegationsStatAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *DelegationsStat) error {
	*o = DelegationsStat{}
	return nil
}

func testDelegationsStatsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &DelegationsStat{}
	o := &DelegationsStat{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, delegationsStatDBTypes, false); err != nil {
		t.Errorf("Unable to randomize DelegationsStat object: %s", err)
	}

	AddDelegationsStatHook(boil.BeforeInsertHook, delegationsStatBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	delegationsStatBeforeInsertHooks = []DelegationsStatHook{}

	AddDelegationsStatHook(boil.AfterInsertHook, delegationsStatAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	delegationsStatAfterInsertHooks = []DelegationsStatHook{}

	AddDelegationsStatHook(boil.AfterSelectHook, delegationsStatAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	delegationsStatAfterSelectHooks = []DelegationsStatHook{}

	AddDelegationsStatHook(boil.BeforeUpdateHook, delegationsStatBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	delegationsStatBeforeUpdateHooks = []DelegationsStatHook{}

	AddDelegationsStatHook(boil.AfterUpdateHook, delegationsStatAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	delegationsStatAfterUpdateHooks = []DelegationsStatHook{}

	AddDelegationsStatHook(boil.BeforeDeleteHook, delegationsStatBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	delegationsStatBeforeDeleteHooks = []DelegationsStatHook{}

	AddDelegationsStatHook(boil.AfterDeleteHook, delegationsStatAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	delegationsStatAfterDeleteHooks = []DelegationsStatHook{}

	AddDelegationsStatHook(boil.BeforeUpsertHook, delegationsStatBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	delegationsStatBeforeUpsertHooks = []DelegationsStatHook{}

	AddDelegationsStatHook(boil.AfterUpsertHook, delegationsStatAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	delegationsStatAfterUpsertHooks = []DelegationsStatHook{}
}

func testDelegationsStatsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DelegationsStat{}
	if err = randomize.Struct(seed, o, delegationsStatDBTypes, true, delegationsStatColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DelegationsStat struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DelegationsStats().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDelegationsStatsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DelegationsStat{}
	if err = randomize.Struct(seed, o, delegationsStatDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DelegationsStat struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(delegationsStatColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := DelegationsStats().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDelegationsStatToOneDelegatorUsingDelegatorAddressDelegator(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local DelegationsStat
	var foreign Delegator

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, delegationsStatDBTypes, true, delegationsStatColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DelegationsStat struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, delegatorDBTypes, false, delegatorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Delegator struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&local.DelegatorAddress, foreign.Address)
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.DelegatorAddressDelegator().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if !queries.Equal(check.Address, foreign.Address) {
		t.Errorf("want: %v, got %v", foreign.Address, check.Address)
	}

	ranAfterSelectHook := false
	AddDelegatorHook(boil.AfterSelectHook, func(ctx context.Context, e boil.ContextExecutor, o *Delegator) error {
		ranAfterSelectHook = true
		return nil
	})

	slice := DelegationsStatSlice{&local}
	if err = local.L.LoadDelegatorAddressDelegator(ctx, tx, false, (*[]*DelegationsStat)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.DelegatorAddressDelegator == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.DelegatorAddressDelegator = nil
	if err = local.L.LoadDelegatorAddressDelegator(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.DelegatorAddressDelegator == nil {
		t.Error("struct should have been eager loaded")
	}

	if !ranAfterSelectHook {
		t.Error("failed to run AfterSelect hook for relationship")
	}
}

func testDelegationsStatToOneSetOpDelegatorUsingDelegatorAddressDelegator(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a DelegationsStat
	var b, c Delegator

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, delegationsStatDBTypes, false, strmangle.SetComplement(delegationsStatPrimaryKeyColumns, delegationsStatColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, delegatorDBTypes, false, strmangle.SetComplement(delegatorPrimaryKeyColumns, delegatorColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, delegatorDBTypes, false, strmangle.SetComplement(delegatorPrimaryKeyColumns, delegatorColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Delegator{&b, &c} {
		err = a.SetDelegatorAddressDelegator(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.DelegatorAddressDelegator != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.DelegatorAddressDelegationsStats[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if !queries.Equal(a.DelegatorAddress, x.Address) {
			t.Error("foreign key was wrong value", a.DelegatorAddress)
		}

		zero := reflect.Zero(reflect.TypeOf(a.DelegatorAddress))
		reflect.Indirect(reflect.ValueOf(&a.DelegatorAddress)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if !queries.Equal(a.DelegatorAddress, x.Address) {
			t.Error("foreign key was wrong value", a.DelegatorAddress, x.Address)
		}
	}
}

func testDelegationsStatToOneRemoveOpDelegatorUsingDelegatorAddressDelegator(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a DelegationsStat
	var b Delegator

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, delegationsStatDBTypes, false, strmangle.SetComplement(delegationsStatPrimaryKeyColumns, delegationsStatColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, delegatorDBTypes, false, strmangle.SetComplement(delegatorPrimaryKeyColumns, delegatorColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = a.SetDelegatorAddressDelegator(ctx, tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemoveDelegatorAddressDelegator(ctx, tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.DelegatorAddressDelegator().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.DelegatorAddressDelegator != nil {
		t.Error("R struct entry should be nil")
	}

	if !queries.IsValuerNil(a.DelegatorAddress) {
		t.Error("foreign key value should be nil")
	}

	if len(b.R.DelegatorAddressDelegationsStats) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testDelegationsStatsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DelegationsStat{}
	if err = randomize.Struct(seed, o, delegationsStatDBTypes, true, delegationsStatColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DelegationsStat struct: %s", err)
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

func testDelegationsStatsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DelegationsStat{}
	if err = randomize.Struct(seed, o, delegationsStatDBTypes, true, delegationsStatColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DelegationsStat struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := DelegationsStatSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testDelegationsStatsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DelegationsStat{}
	if err = randomize.Struct(seed, o, delegationsStatDBTypes, true, delegationsStatColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DelegationsStat struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := DelegationsStats().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	delegationsStatDBTypes = map[string]string{`ID`: `bigint`, `Timestamp`: `timestamp without time zone`, `Kind`: `enum.stat_kind('TOP10VALIDATORS','TOP100VALIDATORS','DAILYVOLUME','WEEKLYVOLUME','MONTHLYVOLUME','YEARLYVOLUME')`, `Value`: `jsonb`, `DelegatorAddress`: `character varying`}
	_                      = bytes.MinRead
)

func testDelegationsStatsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(delegationsStatPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(delegationsStatAllColumns) == len(delegationsStatPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &DelegationsStat{}
	if err = randomize.Struct(seed, o, delegationsStatDBTypes, true, delegationsStatColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DelegationsStat struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DelegationsStats().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, delegationsStatDBTypes, true, delegationsStatPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DelegationsStat struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testDelegationsStatsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(delegationsStatAllColumns) == len(delegationsStatPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &DelegationsStat{}
	if err = randomize.Struct(seed, o, delegationsStatDBTypes, true, delegationsStatColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DelegationsStat struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DelegationsStats().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, delegationsStatDBTypes, true, delegationsStatPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DelegationsStat struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(delegationsStatAllColumns, delegationsStatPrimaryKeyColumns) {
		fields = delegationsStatAllColumns
	} else {
		fields = strmangle.SetComplement(
			delegationsStatAllColumns,
			delegationsStatPrimaryKeyColumns,
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

	slice := DelegationsStatSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testDelegationsStatsUpsert(t *testing.T) {
	t.Parallel()

	if len(delegationsStatAllColumns) == len(delegationsStatPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := DelegationsStat{}
	if err = randomize.Struct(seed, &o, delegationsStatDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DelegationsStat struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert DelegationsStat: %s", err)
	}

	count, err := DelegationsStats().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, delegationsStatDBTypes, false, delegationsStatPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DelegationsStat struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert DelegationsStat: %s", err)
	}

	count, err = DelegationsStats().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
