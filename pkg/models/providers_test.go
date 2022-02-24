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

func testProviders(t *testing.T) {
	t.Parallel()

	query := Providers()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testProvidersDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Provider{}
	if err = randomize.Struct(seed, o, providerDBTypes, true, providerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Provider struct: %s", err)
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

	count, err := Providers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testProvidersQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Provider{}
	if err = randomize.Struct(seed, o, providerDBTypes, true, providerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Provider struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Providers().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Providers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testProvidersSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Provider{}
	if err = randomize.Struct(seed, o, providerDBTypes, true, providerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Provider struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ProviderSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Providers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testProvidersExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Provider{}
	if err = randomize.Struct(seed, o, providerDBTypes, true, providerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Provider struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := ProviderExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Provider exists: %s", err)
	}
	if !e {
		t.Errorf("Expected ProviderExists to return true, but got false.")
	}
}

func testProvidersFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Provider{}
	if err = randomize.Struct(seed, o, providerDBTypes, true, providerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Provider struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	providerFound, err := FindProvider(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if providerFound == nil {
		t.Error("want a record, got nil")
	}
}

func testProvidersBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Provider{}
	if err = randomize.Struct(seed, o, providerDBTypes, true, providerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Provider struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Providers().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testProvidersOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Provider{}
	if err = randomize.Struct(seed, o, providerDBTypes, true, providerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Provider struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Providers().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testProvidersAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	providerOne := &Provider{}
	providerTwo := &Provider{}
	if err = randomize.Struct(seed, providerOne, providerDBTypes, false, providerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Provider struct: %s", err)
	}
	if err = randomize.Struct(seed, providerTwo, providerDBTypes, false, providerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Provider struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = providerOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = providerTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Providers().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testProvidersCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	providerOne := &Provider{}
	providerTwo := &Provider{}
	if err = randomize.Struct(seed, providerOne, providerDBTypes, false, providerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Provider struct: %s", err)
	}
	if err = randomize.Struct(seed, providerTwo, providerDBTypes, false, providerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Provider struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = providerOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = providerTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Providers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func providerBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Provider) error {
	*o = Provider{}
	return nil
}

func providerAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Provider) error {
	*o = Provider{}
	return nil
}

func providerAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Provider) error {
	*o = Provider{}
	return nil
}

func providerBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Provider) error {
	*o = Provider{}
	return nil
}

func providerAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Provider) error {
	*o = Provider{}
	return nil
}

func providerBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Provider) error {
	*o = Provider{}
	return nil
}

func providerAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Provider) error {
	*o = Provider{}
	return nil
}

func providerBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Provider) error {
	*o = Provider{}
	return nil
}

func providerAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Provider) error {
	*o = Provider{}
	return nil
}

func testProvidersHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Provider{}
	o := &Provider{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, providerDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Provider object: %s", err)
	}

	AddProviderHook(boil.BeforeInsertHook, providerBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	providerBeforeInsertHooks = []ProviderHook{}

	AddProviderHook(boil.AfterInsertHook, providerAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	providerAfterInsertHooks = []ProviderHook{}

	AddProviderHook(boil.AfterSelectHook, providerAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	providerAfterSelectHooks = []ProviderHook{}

	AddProviderHook(boil.BeforeUpdateHook, providerBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	providerBeforeUpdateHooks = []ProviderHook{}

	AddProviderHook(boil.AfterUpdateHook, providerAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	providerAfterUpdateHooks = []ProviderHook{}

	AddProviderHook(boil.BeforeDeleteHook, providerBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	providerBeforeDeleteHooks = []ProviderHook{}

	AddProviderHook(boil.AfterDeleteHook, providerAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	providerAfterDeleteHooks = []ProviderHook{}

	AddProviderHook(boil.BeforeUpsertHook, providerBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	providerBeforeUpsertHooks = []ProviderHook{}

	AddProviderHook(boil.AfterUpsertHook, providerAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	providerAfterUpsertHooks = []ProviderHook{}
}

func testProvidersInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Provider{}
	if err = randomize.Struct(seed, o, providerDBTypes, true, providerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Provider struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Providers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testProvidersInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Provider{}
	if err = randomize.Struct(seed, o, providerDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Provider struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(providerColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Providers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testProviderToOnePeerUsingRemote(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Provider
	var foreign Peer

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, providerDBTypes, false, providerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Provider struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, peerDBTypes, false, peerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Peer struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.RemoteID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Remote().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := ProviderSlice{&local}
	if err = local.L.LoadRemote(ctx, tx, false, (*[]*Provider)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Remote == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Remote = nil
	if err = local.L.LoadRemote(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Remote == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testProviderToOneRetrievalUsingRetrieval(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Provider
	var foreign Retrieval

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, providerDBTypes, false, providerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Provider struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, retrievalDBTypes, false, retrievalColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Retrieval struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.RetrievalID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Retrieval().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := ProviderSlice{&local}
	if err = local.L.LoadRetrieval(ctx, tx, false, (*[]*Provider)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Retrieval == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Retrieval = nil
	if err = local.L.LoadRetrieval(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Retrieval == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testProviderToOneSetOpPeerUsingRemote(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Provider
	var b, c Peer

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, providerDBTypes, false, strmangle.SetComplement(providerPrimaryKeyColumns, providerColumnsWithoutDefault)...); err != nil {
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
		err = a.SetRemote(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Remote != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.RemoteProviders[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.RemoteID != x.ID {
			t.Error("foreign key was wrong value", a.RemoteID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.RemoteID))
		reflect.Indirect(reflect.ValueOf(&a.RemoteID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.RemoteID != x.ID {
			t.Error("foreign key was wrong value", a.RemoteID, x.ID)
		}
	}
}
func testProviderToOneSetOpRetrievalUsingRetrieval(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Provider
	var b, c Retrieval

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, providerDBTypes, false, strmangle.SetComplement(providerPrimaryKeyColumns, providerColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, retrievalDBTypes, false, strmangle.SetComplement(retrievalPrimaryKeyColumns, retrievalColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, retrievalDBTypes, false, strmangle.SetComplement(retrievalPrimaryKeyColumns, retrievalColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Retrieval{&b, &c} {
		err = a.SetRetrieval(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Retrieval != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Providers[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.RetrievalID != x.ID {
			t.Error("foreign key was wrong value", a.RetrievalID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.RetrievalID))
		reflect.Indirect(reflect.ValueOf(&a.RetrievalID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.RetrievalID != x.ID {
			t.Error("foreign key was wrong value", a.RetrievalID, x.ID)
		}
	}
}

func testProvidersReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Provider{}
	if err = randomize.Struct(seed, o, providerDBTypes, true, providerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Provider struct: %s", err)
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

func testProvidersReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Provider{}
	if err = randomize.Struct(seed, o, providerDBTypes, true, providerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Provider struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ProviderSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testProvidersSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Provider{}
	if err = randomize.Struct(seed, o, providerDBTypes, true, providerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Provider struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Providers().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	providerDBTypes = map[string]string{`ID`: `integer`, `RetrievalID`: `integer`, `RemoteID`: `integer`, `MultiAddressIds`: `ARRAYinteger`, `FoundAt`: `timestamp with time zone`}
	_               = bytes.MinRead
)

func testProvidersUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(providerPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(providerAllColumns) == len(providerPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Provider{}
	if err = randomize.Struct(seed, o, providerDBTypes, true, providerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Provider struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Providers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, providerDBTypes, true, providerPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Provider struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testProvidersSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(providerAllColumns) == len(providerPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Provider{}
	if err = randomize.Struct(seed, o, providerDBTypes, true, providerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Provider struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Providers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, providerDBTypes, true, providerPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Provider struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(providerAllColumns, providerPrimaryKeyColumns) {
		fields = providerAllColumns
	} else {
		fields = strmangle.SetComplement(
			providerAllColumns,
			providerPrimaryKeyColumns,
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

	slice := ProviderSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testProvidersUpsert(t *testing.T) {
	t.Parallel()

	if len(providerAllColumns) == len(providerPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Provider{}
	if err = randomize.Struct(seed, &o, providerDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Provider struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Provider: %s", err)
	}

	count, err := Providers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, providerDBTypes, false, providerPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Provider struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Provider: %s", err)
	}

	count, err = Providers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}