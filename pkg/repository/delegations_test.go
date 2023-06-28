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

func testDelegations(t *testing.T) {
	t.Parallel()

	query := Delegations()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testDelegationsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Delegation{}
	if err = randomize.Struct(seed, o, delegationDBTypes, true, delegationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Delegation struct: %s", err)
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

	count, err := Delegations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDelegationsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Delegation{}
	if err = randomize.Struct(seed, o, delegationDBTypes, true, delegationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Delegation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Delegations().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Delegations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDelegationsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Delegation{}
	if err = randomize.Struct(seed, o, delegationDBTypes, true, delegationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Delegation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := DelegationSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Delegations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDelegationsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Delegation{}
	if err = randomize.Struct(seed, o, delegationDBTypes, true, delegationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Delegation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := DelegationExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Delegation exists: %s", err)
	}
	if !e {
		t.Errorf("Expected DelegationExists to return true, but got false.")
	}
}

func testDelegationsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Delegation{}
	if err = randomize.Struct(seed, o, delegationDBTypes, true, delegationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Delegation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	delegationFound, err := FindDelegation(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if delegationFound == nil {
		t.Error("want a record, got nil")
	}
}

func testDelegationsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Delegation{}
	if err = randomize.Struct(seed, o, delegationDBTypes, true, delegationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Delegation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Delegations().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testDelegationsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Delegation{}
	if err = randomize.Struct(seed, o, delegationDBTypes, true, delegationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Delegation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Delegations().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testDelegationsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	delegationOne := &Delegation{}
	delegationTwo := &Delegation{}
	if err = randomize.Struct(seed, delegationOne, delegationDBTypes, false, delegationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Delegation struct: %s", err)
	}
	if err = randomize.Struct(seed, delegationTwo, delegationDBTypes, false, delegationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Delegation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = delegationOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = delegationTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Delegations().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testDelegationsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	delegationOne := &Delegation{}
	delegationTwo := &Delegation{}
	if err = randomize.Struct(seed, delegationOne, delegationDBTypes, false, delegationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Delegation struct: %s", err)
	}
	if err = randomize.Struct(seed, delegationTwo, delegationDBTypes, false, delegationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Delegation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = delegationOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = delegationTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Delegations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func delegationBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Delegation) error {
	*o = Delegation{}
	return nil
}

func delegationAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Delegation) error {
	*o = Delegation{}
	return nil
}

func delegationAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Delegation) error {
	*o = Delegation{}
	return nil
}

func delegationBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Delegation) error {
	*o = Delegation{}
	return nil
}

func delegationAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Delegation) error {
	*o = Delegation{}
	return nil
}

func delegationBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Delegation) error {
	*o = Delegation{}
	return nil
}

func delegationAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Delegation) error {
	*o = Delegation{}
	return nil
}

func delegationBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Delegation) error {
	*o = Delegation{}
	return nil
}

func delegationAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Delegation) error {
	*o = Delegation{}
	return nil
}

func testDelegationsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Delegation{}
	o := &Delegation{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, delegationDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Delegation object: %s", err)
	}

	AddDelegationHook(boil.BeforeInsertHook, delegationBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	delegationBeforeInsertHooks = []DelegationHook{}

	AddDelegationHook(boil.AfterInsertHook, delegationAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	delegationAfterInsertHooks = []DelegationHook{}

	AddDelegationHook(boil.AfterSelectHook, delegationAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	delegationAfterSelectHooks = []DelegationHook{}

	AddDelegationHook(boil.BeforeUpdateHook, delegationBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	delegationBeforeUpdateHooks = []DelegationHook{}

	AddDelegationHook(boil.AfterUpdateHook, delegationAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	delegationAfterUpdateHooks = []DelegationHook{}

	AddDelegationHook(boil.BeforeDeleteHook, delegationBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	delegationBeforeDeleteHooks = []DelegationHook{}

	AddDelegationHook(boil.AfterDeleteHook, delegationAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	delegationAfterDeleteHooks = []DelegationHook{}

	AddDelegationHook(boil.BeforeUpsertHook, delegationBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	delegationBeforeUpsertHooks = []DelegationHook{}

	AddDelegationHook(boil.AfterUpsertHook, delegationAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	delegationAfterUpsertHooks = []DelegationHook{}
}

func testDelegationsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Delegation{}
	if err = randomize.Struct(seed, o, delegationDBTypes, true, delegationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Delegation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Delegations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDelegationsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Delegation{}
	if err = randomize.Struct(seed, o, delegationDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Delegation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(delegationColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Delegations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDelegationsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Delegation{}
	if err = randomize.Struct(seed, o, delegationDBTypes, true, delegationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Delegation struct: %s", err)
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

func testDelegationsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Delegation{}
	if err = randomize.Struct(seed, o, delegationDBTypes, true, delegationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Delegation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := DelegationSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testDelegationsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Delegation{}
	if err = randomize.Struct(seed, o, delegationDBTypes, true, delegationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Delegation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Delegations().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	delegationDBTypes = map[string]string{`ID`: `bigint`, `Timestamp`: `timestamp without time zone`, `Amount`: `bigint`, `Delegator`: `character varying`, `Block`: `bigint`}
	_                 = bytes.MinRead
)

func testDelegationsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(delegationPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(delegationAllColumns) == len(delegationPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Delegation{}
	if err = randomize.Struct(seed, o, delegationDBTypes, true, delegationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Delegation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Delegations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, delegationDBTypes, true, delegationPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Delegation struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testDelegationsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(delegationAllColumns) == len(delegationPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Delegation{}
	if err = randomize.Struct(seed, o, delegationDBTypes, true, delegationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Delegation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Delegations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, delegationDBTypes, true, delegationPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Delegation struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(delegationAllColumns, delegationPrimaryKeyColumns) {
		fields = delegationAllColumns
	} else {
		fields = strmangle.SetComplement(
			delegationAllColumns,
			delegationPrimaryKeyColumns,
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

	slice := DelegationSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testDelegationsUpsert(t *testing.T) {
	t.Parallel()

	if len(delegationAllColumns) == len(delegationPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Delegation{}
	if err = randomize.Struct(seed, &o, delegationDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Delegation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Delegation: %s", err)
	}

	count, err := Delegations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, delegationDBTypes, false, delegationPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Delegation struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Delegation: %s", err)
	}

	count, err = Delegations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}