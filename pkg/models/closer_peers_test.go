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

func testCloserPeers(t *testing.T) {
	t.Parallel()

	query := CloserPeers()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testCloserPeersDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CloserPeer{}
	if err = randomize.Struct(seed, o, closerPeerDBTypes, true, closerPeerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CloserPeer struct: %s", err)
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

	count, err := CloserPeers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testCloserPeersQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CloserPeer{}
	if err = randomize.Struct(seed, o, closerPeerDBTypes, true, closerPeerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CloserPeer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := CloserPeers().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := CloserPeers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testCloserPeersSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CloserPeer{}
	if err = randomize.Struct(seed, o, closerPeerDBTypes, true, closerPeerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CloserPeer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := CloserPeerSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := CloserPeers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testCloserPeersExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CloserPeer{}
	if err = randomize.Struct(seed, o, closerPeerDBTypes, true, closerPeerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CloserPeer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := CloserPeerExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if CloserPeer exists: %s", err)
	}
	if !e {
		t.Errorf("Expected CloserPeerExists to return true, but got false.")
	}
}

func testCloserPeersFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CloserPeer{}
	if err = randomize.Struct(seed, o, closerPeerDBTypes, true, closerPeerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CloserPeer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	closerPeerFound, err := FindCloserPeer(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if closerPeerFound == nil {
		t.Error("want a record, got nil")
	}
}

func testCloserPeersBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CloserPeer{}
	if err = randomize.Struct(seed, o, closerPeerDBTypes, true, closerPeerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CloserPeer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = CloserPeers().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testCloserPeersOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CloserPeer{}
	if err = randomize.Struct(seed, o, closerPeerDBTypes, true, closerPeerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CloserPeer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := CloserPeers().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testCloserPeersAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	closerPeerOne := &CloserPeer{}
	closerPeerTwo := &CloserPeer{}
	if err = randomize.Struct(seed, closerPeerOne, closerPeerDBTypes, false, closerPeerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CloserPeer struct: %s", err)
	}
	if err = randomize.Struct(seed, closerPeerTwo, closerPeerDBTypes, false, closerPeerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CloserPeer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = closerPeerOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = closerPeerTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := CloserPeers().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testCloserPeersCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	closerPeerOne := &CloserPeer{}
	closerPeerTwo := &CloserPeer{}
	if err = randomize.Struct(seed, closerPeerOne, closerPeerDBTypes, false, closerPeerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CloserPeer struct: %s", err)
	}
	if err = randomize.Struct(seed, closerPeerTwo, closerPeerDBTypes, false, closerPeerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CloserPeer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = closerPeerOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = closerPeerTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := CloserPeers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func closerPeerBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *CloserPeer) error {
	*o = CloserPeer{}
	return nil
}

func closerPeerAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *CloserPeer) error {
	*o = CloserPeer{}
	return nil
}

func closerPeerAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *CloserPeer) error {
	*o = CloserPeer{}
	return nil
}

func closerPeerBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *CloserPeer) error {
	*o = CloserPeer{}
	return nil
}

func closerPeerAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *CloserPeer) error {
	*o = CloserPeer{}
	return nil
}

func closerPeerBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *CloserPeer) error {
	*o = CloserPeer{}
	return nil
}

func closerPeerAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *CloserPeer) error {
	*o = CloserPeer{}
	return nil
}

func closerPeerBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *CloserPeer) error {
	*o = CloserPeer{}
	return nil
}

func closerPeerAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *CloserPeer) error {
	*o = CloserPeer{}
	return nil
}

func testCloserPeersHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &CloserPeer{}
	o := &CloserPeer{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, closerPeerDBTypes, false); err != nil {
		t.Errorf("Unable to randomize CloserPeer object: %s", err)
	}

	AddCloserPeerHook(boil.BeforeInsertHook, closerPeerBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	closerPeerBeforeInsertHooks = []CloserPeerHook{}

	AddCloserPeerHook(boil.AfterInsertHook, closerPeerAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	closerPeerAfterInsertHooks = []CloserPeerHook{}

	AddCloserPeerHook(boil.AfterSelectHook, closerPeerAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	closerPeerAfterSelectHooks = []CloserPeerHook{}

	AddCloserPeerHook(boil.BeforeUpdateHook, closerPeerBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	closerPeerBeforeUpdateHooks = []CloserPeerHook{}

	AddCloserPeerHook(boil.AfterUpdateHook, closerPeerAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	closerPeerAfterUpdateHooks = []CloserPeerHook{}

	AddCloserPeerHook(boil.BeforeDeleteHook, closerPeerBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	closerPeerBeforeDeleteHooks = []CloserPeerHook{}

	AddCloserPeerHook(boil.AfterDeleteHook, closerPeerAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	closerPeerAfterDeleteHooks = []CloserPeerHook{}

	AddCloserPeerHook(boil.BeforeUpsertHook, closerPeerBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	closerPeerBeforeUpsertHooks = []CloserPeerHook{}

	AddCloserPeerHook(boil.AfterUpsertHook, closerPeerAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	closerPeerAfterUpsertHooks = []CloserPeerHook{}
}

func testCloserPeersInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CloserPeer{}
	if err = randomize.Struct(seed, o, closerPeerDBTypes, true, closerPeerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CloserPeer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := CloserPeers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testCloserPeersInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CloserPeer{}
	if err = randomize.Struct(seed, o, closerPeerDBTypes, true); err != nil {
		t.Errorf("Unable to randomize CloserPeer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(closerPeerColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := CloserPeers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testCloserPeerToOneFindNodesRPCUsingFindNodeRPC(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local CloserPeer
	var foreign FindNodesRPC

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, closerPeerDBTypes, true, closerPeerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CloserPeer struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, findNodesRPCDBTypes, false, findNodesRPCColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FindNodesRPC struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&local.FindNodeRPCID, foreign.ID)
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.FindNodeRPC().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if !queries.Equal(check.ID, foreign.ID) {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := CloserPeerSlice{&local}
	if err = local.L.LoadFindNodeRPC(ctx, tx, false, (*[]*CloserPeer)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.FindNodeRPC == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.FindNodeRPC = nil
	if err = local.L.LoadFindNodeRPC(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.FindNodeRPC == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testCloserPeerToOneGetProvidersRPCUsingGetProvidersRPC(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local CloserPeer
	var foreign GetProvidersRPC

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, closerPeerDBTypes, true, closerPeerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CloserPeer struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, getProvidersRPCDBTypes, false, getProvidersRPCColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GetProvidersRPC struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&local.GetProvidersRPCID, foreign.ID)
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.GetProvidersRPC().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if !queries.Equal(check.ID, foreign.ID) {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := CloserPeerSlice{&local}
	if err = local.L.LoadGetProvidersRPC(ctx, tx, false, (*[]*CloserPeer)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.GetProvidersRPC == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.GetProvidersRPC = nil
	if err = local.L.LoadGetProvidersRPC(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.GetProvidersRPC == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testCloserPeerToOnePeerUsingPeer(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local CloserPeer
	var foreign Peer

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, closerPeerDBTypes, false, closerPeerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CloserPeer struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, peerDBTypes, false, peerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Peer struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.PeerID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Peer().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := CloserPeerSlice{&local}
	if err = local.L.LoadPeer(ctx, tx, false, (*[]*CloserPeer)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Peer == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Peer = nil
	if err = local.L.LoadPeer(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Peer == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testCloserPeerToOneSetOpFindNodesRPCUsingFindNodeRPC(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a CloserPeer
	var b, c FindNodesRPC

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, closerPeerDBTypes, false, strmangle.SetComplement(closerPeerPrimaryKeyColumns, closerPeerColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, findNodesRPCDBTypes, false, strmangle.SetComplement(findNodesRPCPrimaryKeyColumns, findNodesRPCColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, findNodesRPCDBTypes, false, strmangle.SetComplement(findNodesRPCPrimaryKeyColumns, findNodesRPCColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*FindNodesRPC{&b, &c} {
		err = a.SetFindNodeRPC(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.FindNodeRPC != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.FindNodeRPCCloserPeers[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if !queries.Equal(a.FindNodeRPCID, x.ID) {
			t.Error("foreign key was wrong value", a.FindNodeRPCID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.FindNodeRPCID))
		reflect.Indirect(reflect.ValueOf(&a.FindNodeRPCID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if !queries.Equal(a.FindNodeRPCID, x.ID) {
			t.Error("foreign key was wrong value", a.FindNodeRPCID, x.ID)
		}
	}
}

func testCloserPeerToOneRemoveOpFindNodesRPCUsingFindNodeRPC(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a CloserPeer
	var b FindNodesRPC

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, closerPeerDBTypes, false, strmangle.SetComplement(closerPeerPrimaryKeyColumns, closerPeerColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, findNodesRPCDBTypes, false, strmangle.SetComplement(findNodesRPCPrimaryKeyColumns, findNodesRPCColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = a.SetFindNodeRPC(ctx, tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemoveFindNodeRPC(ctx, tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.FindNodeRPC().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.FindNodeRPC != nil {
		t.Error("R struct entry should be nil")
	}

	if !queries.IsValuerNil(a.FindNodeRPCID) {
		t.Error("foreign key value should be nil")
	}

	if len(b.R.FindNodeRPCCloserPeers) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testCloserPeerToOneSetOpGetProvidersRPCUsingGetProvidersRPC(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a CloserPeer
	var b, c GetProvidersRPC

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, closerPeerDBTypes, false, strmangle.SetComplement(closerPeerPrimaryKeyColumns, closerPeerColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, getProvidersRPCDBTypes, false, strmangle.SetComplement(getProvidersRPCPrimaryKeyColumns, getProvidersRPCColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, getProvidersRPCDBTypes, false, strmangle.SetComplement(getProvidersRPCPrimaryKeyColumns, getProvidersRPCColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*GetProvidersRPC{&b, &c} {
		err = a.SetGetProvidersRPC(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.GetProvidersRPC != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.CloserPeers[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if !queries.Equal(a.GetProvidersRPCID, x.ID) {
			t.Error("foreign key was wrong value", a.GetProvidersRPCID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.GetProvidersRPCID))
		reflect.Indirect(reflect.ValueOf(&a.GetProvidersRPCID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if !queries.Equal(a.GetProvidersRPCID, x.ID) {
			t.Error("foreign key was wrong value", a.GetProvidersRPCID, x.ID)
		}
	}
}

func testCloserPeerToOneRemoveOpGetProvidersRPCUsingGetProvidersRPC(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a CloserPeer
	var b GetProvidersRPC

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, closerPeerDBTypes, false, strmangle.SetComplement(closerPeerPrimaryKeyColumns, closerPeerColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, getProvidersRPCDBTypes, false, strmangle.SetComplement(getProvidersRPCPrimaryKeyColumns, getProvidersRPCColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = a.SetGetProvidersRPC(ctx, tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemoveGetProvidersRPC(ctx, tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.GetProvidersRPC().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.GetProvidersRPC != nil {
		t.Error("R struct entry should be nil")
	}

	if !queries.IsValuerNil(a.GetProvidersRPCID) {
		t.Error("foreign key value should be nil")
	}

	if len(b.R.CloserPeers) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testCloserPeerToOneSetOpPeerUsingPeer(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a CloserPeer
	var b, c Peer

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, closerPeerDBTypes, false, strmangle.SetComplement(closerPeerPrimaryKeyColumns, closerPeerColumnsWithoutDefault)...); err != nil {
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
		err = a.SetPeer(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Peer != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.CloserPeers[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.PeerID != x.ID {
			t.Error("foreign key was wrong value", a.PeerID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.PeerID))
		reflect.Indirect(reflect.ValueOf(&a.PeerID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.PeerID != x.ID {
			t.Error("foreign key was wrong value", a.PeerID, x.ID)
		}
	}
}

func testCloserPeersReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CloserPeer{}
	if err = randomize.Struct(seed, o, closerPeerDBTypes, true, closerPeerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CloserPeer struct: %s", err)
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

func testCloserPeersReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CloserPeer{}
	if err = randomize.Struct(seed, o, closerPeerDBTypes, true, closerPeerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CloserPeer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := CloserPeerSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testCloserPeersSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &CloserPeer{}
	if err = randomize.Struct(seed, o, closerPeerDBTypes, true, closerPeerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CloserPeer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := CloserPeers().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	closerPeerDBTypes = map[string]string{`ID`: `integer`, `FindNodeRPCID`: `integer`, `GetProvidersRPCID`: `integer`, `PeerID`: `integer`, `MultiAddressIds`: `ARRAYinteger`}
	_                 = bytes.MinRead
)

func testCloserPeersUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(closerPeerPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(closerPeerAllColumns) == len(closerPeerPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &CloserPeer{}
	if err = randomize.Struct(seed, o, closerPeerDBTypes, true, closerPeerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CloserPeer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := CloserPeers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, closerPeerDBTypes, true, closerPeerPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize CloserPeer struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testCloserPeersSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(closerPeerAllColumns) == len(closerPeerPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &CloserPeer{}
	if err = randomize.Struct(seed, o, closerPeerDBTypes, true, closerPeerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize CloserPeer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := CloserPeers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, closerPeerDBTypes, true, closerPeerPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize CloserPeer struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(closerPeerAllColumns, closerPeerPrimaryKeyColumns) {
		fields = closerPeerAllColumns
	} else {
		fields = strmangle.SetComplement(
			closerPeerAllColumns,
			closerPeerPrimaryKeyColumns,
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

	slice := CloserPeerSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testCloserPeersUpsert(t *testing.T) {
	t.Parallel()

	if len(closerPeerAllColumns) == len(closerPeerPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := CloserPeer{}
	if err = randomize.Struct(seed, &o, closerPeerDBTypes, true); err != nil {
		t.Errorf("Unable to randomize CloserPeer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert CloserPeer: %s", err)
	}

	count, err := CloserPeers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, closerPeerDBTypes, false, closerPeerPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize CloserPeer struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert CloserPeer: %s", err)
	}

	count, err = CloserPeers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
