// Code generated by SQLBoiler 4.6.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

func testProvides(t *testing.T) {
	t.Parallel()

	query := Provides()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testProvidesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Provide{}
	if err = randomize.Struct(seed, o, provideDBTypes, true, provideColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Provide struct: %s", err)
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

	count, err := Provides().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testProvidesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Provide{}
	if err = randomize.Struct(seed, o, provideDBTypes, true, provideColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Provide struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Provides().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Provides().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testProvidesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Provide{}
	if err = randomize.Struct(seed, o, provideDBTypes, true, provideColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Provide struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ProvideSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Provides().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testProvidesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Provide{}
	if err = randomize.Struct(seed, o, provideDBTypes, true, provideColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Provide struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := ProvideExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Provide exists: %s", err)
	}
	if !e {
		t.Errorf("Expected ProvideExists to return true, but got false.")
	}
}

func testProvidesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Provide{}
	if err = randomize.Struct(seed, o, provideDBTypes, true, provideColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Provide struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	provideFound, err := FindProvide(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if provideFound == nil {
		t.Error("want a record, got nil")
	}
}

func testProvidesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Provide{}
	if err = randomize.Struct(seed, o, provideDBTypes, true, provideColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Provide struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Provides().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testProvidesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Provide{}
	if err = randomize.Struct(seed, o, provideDBTypes, true, provideColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Provide struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Provides().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testProvidesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	provideOne := &Provide{}
	provideTwo := &Provide{}
	if err = randomize.Struct(seed, provideOne, provideDBTypes, false, provideColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Provide struct: %s", err)
	}
	if err = randomize.Struct(seed, provideTwo, provideDBTypes, false, provideColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Provide struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = provideOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = provideTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Provides().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testProvidesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	provideOne := &Provide{}
	provideTwo := &Provide{}
	if err = randomize.Struct(seed, provideOne, provideDBTypes, false, provideColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Provide struct: %s", err)
	}
	if err = randomize.Struct(seed, provideTwo, provideDBTypes, false, provideColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Provide struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = provideOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = provideTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Provides().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func provideBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Provide) error {
	*o = Provide{}
	return nil
}

func provideAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Provide) error {
	*o = Provide{}
	return nil
}

func provideAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Provide) error {
	*o = Provide{}
	return nil
}

func provideBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Provide) error {
	*o = Provide{}
	return nil
}

func provideAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Provide) error {
	*o = Provide{}
	return nil
}

func provideBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Provide) error {
	*o = Provide{}
	return nil
}

func provideAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Provide) error {
	*o = Provide{}
	return nil
}

func provideBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Provide) error {
	*o = Provide{}
	return nil
}

func provideAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Provide) error {
	*o = Provide{}
	return nil
}

func testProvidesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Provide{}
	o := &Provide{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, provideDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Provide object: %s", err)
	}

	AddProvideHook(boil.BeforeInsertHook, provideBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	provideBeforeInsertHooks = []ProvideHook{}

	AddProvideHook(boil.AfterInsertHook, provideAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	provideAfterInsertHooks = []ProvideHook{}

	AddProvideHook(boil.AfterSelectHook, provideAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	provideAfterSelectHooks = []ProvideHook{}

	AddProvideHook(boil.BeforeUpdateHook, provideBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	provideBeforeUpdateHooks = []ProvideHook{}

	AddProvideHook(boil.AfterUpdateHook, provideAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	provideAfterUpdateHooks = []ProvideHook{}

	AddProvideHook(boil.BeforeDeleteHook, provideBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	provideBeforeDeleteHooks = []ProvideHook{}

	AddProvideHook(boil.AfterDeleteHook, provideAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	provideAfterDeleteHooks = []ProvideHook{}

	AddProvideHook(boil.BeforeUpsertHook, provideBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	provideBeforeUpsertHooks = []ProvideHook{}

	AddProvideHook(boil.AfterUpsertHook, provideAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	provideAfterUpsertHooks = []ProvideHook{}
}

func testProvidesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Provide{}
	if err = randomize.Struct(seed, o, provideDBTypes, true, provideColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Provide struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Provides().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testProvidesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Provide{}
	if err = randomize.Struct(seed, o, provideDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Provide struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(provideColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Provides().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testProvideToOnePeerUsingProvider(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Provide
	var foreign Peer

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, provideDBTypes, false, provideColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Provide struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, peerDBTypes, false, peerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Peer struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.ProviderID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Provider().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := ProvideSlice{&local}
	if err = local.L.LoadProvider(ctx, tx, false, (*[]*Provide)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Provider == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Provider = nil
	if err = local.L.LoadProvider(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Provider == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testProvideToOneSetOpPeerUsingProvider(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Provide
	var b, c Peer

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, provideDBTypes, false, strmangle.SetComplement(providePrimaryKeyColumns, provideColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, peerDBTypes, false, strmangle.SetComplement(peerPrimaryKeyColumns, peerColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, peerDBTypes, false, strmangle.SetComplement(peerPrimaryKeyColumns, peerColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Peer{&b, &c} {
		err = a.SetProvider(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Provider != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.ProviderProvides[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.ProviderID != x.ID {
			t.Error("foreign key was wrong value", a.ProviderID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.ProviderID))
		reflect.Indirect(reflect.ValueOf(&a.ProviderID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.ProviderID != x.ID {
			t.Error("foreign key was wrong value", a.ProviderID, x.ID)
		}
	}
}

func testProvidesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Provide{}
	if err = randomize.Struct(seed, o, provideDBTypes, true, provideColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Provide struct: %s", err)
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

func testProvidesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Provide{}
	if err = randomize.Struct(seed, o, provideDBTypes, true, provideColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Provide struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ProvideSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testProvidesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Provide{}
	if err = randomize.Struct(seed, o, provideDBTypes, true, provideColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Provide struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Provides().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	provideDBTypes = map[string]string{`ID`: `integer`, `ProviderID`: `integer`, `ContentID`: `text`, `InitialRoutingTableID`: `integer`, `FinalRoutingTableID`: `integer`, `StartedAt`: `timestamp with time zone`, `UpdatedAt`: `timestamp with time zone`, `CreatedAt`: `timestamp with time zone`}
	_              = bytes.MinRead
)

func testProvidesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(providePrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(provideAllColumns) == len(providePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Provide{}
	if err = randomize.Struct(seed, o, provideDBTypes, true, provideColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Provide struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Provides().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, provideDBTypes, true, providePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Provide struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testProvidesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(provideAllColumns) == len(providePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Provide{}
	if err = randomize.Struct(seed, o, provideDBTypes, true, provideColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Provide struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Provides().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, provideDBTypes, true, providePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Provide struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(provideAllColumns, providePrimaryKeyColumns) {
		fields = provideAllColumns
	} else {
		fields = strmangle.SetComplement(
			provideAllColumns,
			providePrimaryKeyColumns,
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

	slice := ProvideSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testProvidesUpsert(t *testing.T) {
	t.Parallel()

	if len(provideAllColumns) == len(providePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Provide{}
	if err = randomize.Struct(seed, &o, provideDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Provide struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Provide: %s", err)
	}

	count, err := Provides().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, provideDBTypes, false, providePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Provide struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Provide: %s", err)
	}

	count, err = Provides().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
