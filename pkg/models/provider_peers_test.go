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

func testProviderPeers(t *testing.T) {
	t.Parallel()

	query := ProviderPeers()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testProviderPeersDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ProviderPeer{}
	if err = randomize.Struct(seed, o, providerPeerDBTypes, true, providerPeerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ProviderPeer struct: %s", err)
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

	count, err := ProviderPeers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testProviderPeersQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ProviderPeer{}
	if err = randomize.Struct(seed, o, providerPeerDBTypes, true, providerPeerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ProviderPeer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := ProviderPeers().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ProviderPeers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testProviderPeersSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ProviderPeer{}
	if err = randomize.Struct(seed, o, providerPeerDBTypes, true, providerPeerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ProviderPeer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ProviderPeerSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ProviderPeers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testProviderPeersExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ProviderPeer{}
	if err = randomize.Struct(seed, o, providerPeerDBTypes, true, providerPeerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ProviderPeer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := ProviderPeerExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if ProviderPeer exists: %s", err)
	}
	if !e {
		t.Errorf("Expected ProviderPeerExists to return true, but got false.")
	}
}

func testProviderPeersFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ProviderPeer{}
	if err = randomize.Struct(seed, o, providerPeerDBTypes, true, providerPeerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ProviderPeer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	providerPeerFound, err := FindProviderPeer(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if providerPeerFound == nil {
		t.Error("want a record, got nil")
	}
}

func testProviderPeersBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ProviderPeer{}
	if err = randomize.Struct(seed, o, providerPeerDBTypes, true, providerPeerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ProviderPeer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = ProviderPeers().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testProviderPeersOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ProviderPeer{}
	if err = randomize.Struct(seed, o, providerPeerDBTypes, true, providerPeerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ProviderPeer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := ProviderPeers().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testProviderPeersAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	providerPeerOne := &ProviderPeer{}
	providerPeerTwo := &ProviderPeer{}
	if err = randomize.Struct(seed, providerPeerOne, providerPeerDBTypes, false, providerPeerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ProviderPeer struct: %s", err)
	}
	if err = randomize.Struct(seed, providerPeerTwo, providerPeerDBTypes, false, providerPeerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ProviderPeer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = providerPeerOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = providerPeerTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := ProviderPeers().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testProviderPeersCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	providerPeerOne := &ProviderPeer{}
	providerPeerTwo := &ProviderPeer{}
	if err = randomize.Struct(seed, providerPeerOne, providerPeerDBTypes, false, providerPeerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ProviderPeer struct: %s", err)
	}
	if err = randomize.Struct(seed, providerPeerTwo, providerPeerDBTypes, false, providerPeerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ProviderPeer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = providerPeerOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = providerPeerTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ProviderPeers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func providerPeerBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *ProviderPeer) error {
	*o = ProviderPeer{}
	return nil
}

func providerPeerAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *ProviderPeer) error {
	*o = ProviderPeer{}
	return nil
}

func providerPeerAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *ProviderPeer) error {
	*o = ProviderPeer{}
	return nil
}

func providerPeerBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *ProviderPeer) error {
	*o = ProviderPeer{}
	return nil
}

func providerPeerAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *ProviderPeer) error {
	*o = ProviderPeer{}
	return nil
}

func providerPeerBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *ProviderPeer) error {
	*o = ProviderPeer{}
	return nil
}

func providerPeerAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *ProviderPeer) error {
	*o = ProviderPeer{}
	return nil
}

func providerPeerBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *ProviderPeer) error {
	*o = ProviderPeer{}
	return nil
}

func providerPeerAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *ProviderPeer) error {
	*o = ProviderPeer{}
	return nil
}

func testProviderPeersHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &ProviderPeer{}
	o := &ProviderPeer{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, providerPeerDBTypes, false); err != nil {
		t.Errorf("Unable to randomize ProviderPeer object: %s", err)
	}

	AddProviderPeerHook(boil.BeforeInsertHook, providerPeerBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	providerPeerBeforeInsertHooks = []ProviderPeerHook{}

	AddProviderPeerHook(boil.AfterInsertHook, providerPeerAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	providerPeerAfterInsertHooks = []ProviderPeerHook{}

	AddProviderPeerHook(boil.AfterSelectHook, providerPeerAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	providerPeerAfterSelectHooks = []ProviderPeerHook{}

	AddProviderPeerHook(boil.BeforeUpdateHook, providerPeerBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	providerPeerBeforeUpdateHooks = []ProviderPeerHook{}

	AddProviderPeerHook(boil.AfterUpdateHook, providerPeerAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	providerPeerAfterUpdateHooks = []ProviderPeerHook{}

	AddProviderPeerHook(boil.BeforeDeleteHook, providerPeerBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	providerPeerBeforeDeleteHooks = []ProviderPeerHook{}

	AddProviderPeerHook(boil.AfterDeleteHook, providerPeerAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	providerPeerAfterDeleteHooks = []ProviderPeerHook{}

	AddProviderPeerHook(boil.BeforeUpsertHook, providerPeerBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	providerPeerBeforeUpsertHooks = []ProviderPeerHook{}

	AddProviderPeerHook(boil.AfterUpsertHook, providerPeerAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	providerPeerAfterUpsertHooks = []ProviderPeerHook{}
}

func testProviderPeersInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ProviderPeer{}
	if err = randomize.Struct(seed, o, providerPeerDBTypes, true, providerPeerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ProviderPeer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ProviderPeers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testProviderPeersInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ProviderPeer{}
	if err = randomize.Struct(seed, o, providerPeerDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ProviderPeer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(providerPeerColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := ProviderPeers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testProviderPeerToOneGetProviderUsingGetProvider(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local ProviderPeer
	var foreign GetProvider

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, providerPeerDBTypes, false, providerPeerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ProviderPeer struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, getProviderDBTypes, false, getProviderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GetProvider struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.GetProvidersID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.GetProvider().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := ProviderPeerSlice{&local}
	if err = local.L.LoadGetProvider(ctx, tx, false, (*[]*ProviderPeer)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.GetProvider == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.GetProvider = nil
	if err = local.L.LoadGetProvider(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.GetProvider == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testProviderPeerToOnePeerUsingProvider(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local ProviderPeer
	var foreign Peer

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, providerPeerDBTypes, false, providerPeerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ProviderPeer struct: %s", err)
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

	slice := ProviderPeerSlice{&local}
	if err = local.L.LoadProvider(ctx, tx, false, (*[]*ProviderPeer)(&slice), nil); err != nil {
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

func testProviderPeerToOneSetOpGetProviderUsingGetProvider(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a ProviderPeer
	var b, c GetProvider

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, providerPeerDBTypes, false, strmangle.SetComplement(providerPeerPrimaryKeyColumns, providerPeerColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, getProviderDBTypes, false, strmangle.SetComplement(getProviderPrimaryKeyColumns, getProviderColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, getProviderDBTypes, false, strmangle.SetComplement(getProviderPrimaryKeyColumns, getProviderColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*GetProvider{&b, &c} {
		err = a.SetGetProvider(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.GetProvider != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.ProviderPeers[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.GetProvidersID != x.ID {
			t.Error("foreign key was wrong value", a.GetProvidersID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.GetProvidersID))
		reflect.Indirect(reflect.ValueOf(&a.GetProvidersID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.GetProvidersID != x.ID {
			t.Error("foreign key was wrong value", a.GetProvidersID, x.ID)
		}
	}
}
func testProviderPeerToOneSetOpPeerUsingProvider(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a ProviderPeer
	var b, c Peer

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, providerPeerDBTypes, false, strmangle.SetComplement(providerPeerPrimaryKeyColumns, providerPeerColumnsWithoutDefault)...); err != nil {
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

		if x.R.ProviderProviderPeers[0] != &a {
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

func testProviderPeersReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ProviderPeer{}
	if err = randomize.Struct(seed, o, providerPeerDBTypes, true, providerPeerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ProviderPeer struct: %s", err)
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

func testProviderPeersReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ProviderPeer{}
	if err = randomize.Struct(seed, o, providerPeerDBTypes, true, providerPeerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ProviderPeer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ProviderPeerSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testProviderPeersSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ProviderPeer{}
	if err = randomize.Struct(seed, o, providerPeerDBTypes, true, providerPeerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ProviderPeer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := ProviderPeers().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	providerPeerDBTypes = map[string]string{`ID`: `integer`, `GetProvidersID`: `integer`, `ProviderID`: `integer`, `MultiAddressIds`: `ARRAYinteger`}
	_                   = bytes.MinRead
)

func testProviderPeersUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(providerPeerPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(providerPeerAllColumns) == len(providerPeerPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &ProviderPeer{}
	if err = randomize.Struct(seed, o, providerPeerDBTypes, true, providerPeerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ProviderPeer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ProviderPeers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, providerPeerDBTypes, true, providerPeerPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ProviderPeer struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testProviderPeersSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(providerPeerAllColumns) == len(providerPeerPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &ProviderPeer{}
	if err = randomize.Struct(seed, o, providerPeerDBTypes, true, providerPeerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ProviderPeer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ProviderPeers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, providerPeerDBTypes, true, providerPeerPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ProviderPeer struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(providerPeerAllColumns, providerPeerPrimaryKeyColumns) {
		fields = providerPeerAllColumns
	} else {
		fields = strmangle.SetComplement(
			providerPeerAllColumns,
			providerPeerPrimaryKeyColumns,
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

	slice := ProviderPeerSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testProviderPeersUpsert(t *testing.T) {
	t.Parallel()

	if len(providerPeerAllColumns) == len(providerPeerPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := ProviderPeer{}
	if err = randomize.Struct(seed, &o, providerPeerDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ProviderPeer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert ProviderPeer: %s", err)
	}

	count, err := ProviderPeers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, providerPeerDBTypes, false, providerPeerPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ProviderPeer struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert ProviderPeer: %s", err)
	}

	count, err = ProviderPeers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
