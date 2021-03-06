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

func testRoutingTableEntries(t *testing.T) {
	t.Parallel()

	query := RoutingTableEntries()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testRoutingTableEntriesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RoutingTableEntry{}
	if err = randomize.Struct(seed, o, routingTableEntryDBTypes, true, routingTableEntryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RoutingTableEntry struct: %s", err)
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

	count, err := RoutingTableEntries().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRoutingTableEntriesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RoutingTableEntry{}
	if err = randomize.Struct(seed, o, routingTableEntryDBTypes, true, routingTableEntryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RoutingTableEntry struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := RoutingTableEntries().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := RoutingTableEntries().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRoutingTableEntriesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RoutingTableEntry{}
	if err = randomize.Struct(seed, o, routingTableEntryDBTypes, true, routingTableEntryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RoutingTableEntry struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := RoutingTableEntrySlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := RoutingTableEntries().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRoutingTableEntriesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RoutingTableEntry{}
	if err = randomize.Struct(seed, o, routingTableEntryDBTypes, true, routingTableEntryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RoutingTableEntry struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := RoutingTableEntryExists(ctx, tx, o.RoutingTableSnapshotID, o.PeerID)
	if err != nil {
		t.Errorf("Unable to check if RoutingTableEntry exists: %s", err)
	}
	if !e {
		t.Errorf("Expected RoutingTableEntryExists to return true, but got false.")
	}
}

func testRoutingTableEntriesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RoutingTableEntry{}
	if err = randomize.Struct(seed, o, routingTableEntryDBTypes, true, routingTableEntryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RoutingTableEntry struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	routingTableEntryFound, err := FindRoutingTableEntry(ctx, tx, o.RoutingTableSnapshotID, o.PeerID)
	if err != nil {
		t.Error(err)
	}

	if routingTableEntryFound == nil {
		t.Error("want a record, got nil")
	}
}

func testRoutingTableEntriesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RoutingTableEntry{}
	if err = randomize.Struct(seed, o, routingTableEntryDBTypes, true, routingTableEntryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RoutingTableEntry struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = RoutingTableEntries().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testRoutingTableEntriesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RoutingTableEntry{}
	if err = randomize.Struct(seed, o, routingTableEntryDBTypes, true, routingTableEntryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RoutingTableEntry struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := RoutingTableEntries().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testRoutingTableEntriesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	routingTableEntryOne := &RoutingTableEntry{}
	routingTableEntryTwo := &RoutingTableEntry{}
	if err = randomize.Struct(seed, routingTableEntryOne, routingTableEntryDBTypes, false, routingTableEntryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RoutingTableEntry struct: %s", err)
	}
	if err = randomize.Struct(seed, routingTableEntryTwo, routingTableEntryDBTypes, false, routingTableEntryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RoutingTableEntry struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = routingTableEntryOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = routingTableEntryTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := RoutingTableEntries().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testRoutingTableEntriesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	routingTableEntryOne := &RoutingTableEntry{}
	routingTableEntryTwo := &RoutingTableEntry{}
	if err = randomize.Struct(seed, routingTableEntryOne, routingTableEntryDBTypes, false, routingTableEntryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RoutingTableEntry struct: %s", err)
	}
	if err = randomize.Struct(seed, routingTableEntryTwo, routingTableEntryDBTypes, false, routingTableEntryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RoutingTableEntry struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = routingTableEntryOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = routingTableEntryTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RoutingTableEntries().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func routingTableEntryBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *RoutingTableEntry) error {
	*o = RoutingTableEntry{}
	return nil
}

func routingTableEntryAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *RoutingTableEntry) error {
	*o = RoutingTableEntry{}
	return nil
}

func routingTableEntryAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *RoutingTableEntry) error {
	*o = RoutingTableEntry{}
	return nil
}

func routingTableEntryBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *RoutingTableEntry) error {
	*o = RoutingTableEntry{}
	return nil
}

func routingTableEntryAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *RoutingTableEntry) error {
	*o = RoutingTableEntry{}
	return nil
}

func routingTableEntryBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *RoutingTableEntry) error {
	*o = RoutingTableEntry{}
	return nil
}

func routingTableEntryAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *RoutingTableEntry) error {
	*o = RoutingTableEntry{}
	return nil
}

func routingTableEntryBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *RoutingTableEntry) error {
	*o = RoutingTableEntry{}
	return nil
}

func routingTableEntryAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *RoutingTableEntry) error {
	*o = RoutingTableEntry{}
	return nil
}

func testRoutingTableEntriesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &RoutingTableEntry{}
	o := &RoutingTableEntry{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, routingTableEntryDBTypes, false); err != nil {
		t.Errorf("Unable to randomize RoutingTableEntry object: %s", err)
	}

	AddRoutingTableEntryHook(boil.BeforeInsertHook, routingTableEntryBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	routingTableEntryBeforeInsertHooks = []RoutingTableEntryHook{}

	AddRoutingTableEntryHook(boil.AfterInsertHook, routingTableEntryAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	routingTableEntryAfterInsertHooks = []RoutingTableEntryHook{}

	AddRoutingTableEntryHook(boil.AfterSelectHook, routingTableEntryAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	routingTableEntryAfterSelectHooks = []RoutingTableEntryHook{}

	AddRoutingTableEntryHook(boil.BeforeUpdateHook, routingTableEntryBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	routingTableEntryBeforeUpdateHooks = []RoutingTableEntryHook{}

	AddRoutingTableEntryHook(boil.AfterUpdateHook, routingTableEntryAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	routingTableEntryAfterUpdateHooks = []RoutingTableEntryHook{}

	AddRoutingTableEntryHook(boil.BeforeDeleteHook, routingTableEntryBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	routingTableEntryBeforeDeleteHooks = []RoutingTableEntryHook{}

	AddRoutingTableEntryHook(boil.AfterDeleteHook, routingTableEntryAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	routingTableEntryAfterDeleteHooks = []RoutingTableEntryHook{}

	AddRoutingTableEntryHook(boil.BeforeUpsertHook, routingTableEntryBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	routingTableEntryBeforeUpsertHooks = []RoutingTableEntryHook{}

	AddRoutingTableEntryHook(boil.AfterUpsertHook, routingTableEntryAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	routingTableEntryAfterUpsertHooks = []RoutingTableEntryHook{}
}

func testRoutingTableEntriesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RoutingTableEntry{}
	if err = randomize.Struct(seed, o, routingTableEntryDBTypes, true, routingTableEntryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RoutingTableEntry struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RoutingTableEntries().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testRoutingTableEntriesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RoutingTableEntry{}
	if err = randomize.Struct(seed, o, routingTableEntryDBTypes, true); err != nil {
		t.Errorf("Unable to randomize RoutingTableEntry struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(routingTableEntryColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := RoutingTableEntries().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testRoutingTableEntryToOnePeerUsingPeer(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local RoutingTableEntry
	var foreign Peer

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, routingTableEntryDBTypes, false, routingTableEntryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RoutingTableEntry struct: %s", err)
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

	slice := RoutingTableEntrySlice{&local}
	if err = local.L.LoadPeer(ctx, tx, false, (*[]*RoutingTableEntry)(&slice), nil); err != nil {
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

func testRoutingTableEntryToOneRoutingTableSnapshotUsingRoutingTableSnapshot(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local RoutingTableEntry
	var foreign RoutingTableSnapshot

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, routingTableEntryDBTypes, false, routingTableEntryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RoutingTableEntry struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, routingTableSnapshotDBTypes, false, routingTableSnapshotColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RoutingTableSnapshot struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.RoutingTableSnapshotID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.RoutingTableSnapshot().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := RoutingTableEntrySlice{&local}
	if err = local.L.LoadRoutingTableSnapshot(ctx, tx, false, (*[]*RoutingTableEntry)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.RoutingTableSnapshot == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.RoutingTableSnapshot = nil
	if err = local.L.LoadRoutingTableSnapshot(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.RoutingTableSnapshot == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testRoutingTableEntryToOneSetOpPeerUsingPeer(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a RoutingTableEntry
	var b, c Peer

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, routingTableEntryDBTypes, false, strmangle.SetComplement(routingTableEntryPrimaryKeyColumns, routingTableEntryColumnsWithoutDefault)...); err != nil {
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

		if x.R.RoutingTableEntries[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.PeerID != x.ID {
			t.Error("foreign key was wrong value", a.PeerID)
		}

		if exists, err := RoutingTableEntryExists(ctx, tx, a.RoutingTableSnapshotID, a.PeerID); err != nil {
			t.Fatal(err)
		} else if !exists {
			t.Error("want 'a' to exist")
		}

	}
}
func testRoutingTableEntryToOneSetOpRoutingTableSnapshotUsingRoutingTableSnapshot(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a RoutingTableEntry
	var b, c RoutingTableSnapshot

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, routingTableEntryDBTypes, false, strmangle.SetComplement(routingTableEntryPrimaryKeyColumns, routingTableEntryColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, routingTableSnapshotDBTypes, false, strmangle.SetComplement(routingTableSnapshotPrimaryKeyColumns, routingTableSnapshotColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, routingTableSnapshotDBTypes, false, strmangle.SetComplement(routingTableSnapshotPrimaryKeyColumns, routingTableSnapshotColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*RoutingTableSnapshot{&b, &c} {
		err = a.SetRoutingTableSnapshot(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.RoutingTableSnapshot != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.RoutingTableEntries[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.RoutingTableSnapshotID != x.ID {
			t.Error("foreign key was wrong value", a.RoutingTableSnapshotID)
		}

		if exists, err := RoutingTableEntryExists(ctx, tx, a.RoutingTableSnapshotID, a.PeerID); err != nil {
			t.Fatal(err)
		} else if !exists {
			t.Error("want 'a' to exist")
		}

	}
}

func testRoutingTableEntriesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RoutingTableEntry{}
	if err = randomize.Struct(seed, o, routingTableEntryDBTypes, true, routingTableEntryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RoutingTableEntry struct: %s", err)
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

func testRoutingTableEntriesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RoutingTableEntry{}
	if err = randomize.Struct(seed, o, routingTableEntryDBTypes, true, routingTableEntryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RoutingTableEntry struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := RoutingTableEntrySlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testRoutingTableEntriesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RoutingTableEntry{}
	if err = randomize.Struct(seed, o, routingTableEntryDBTypes, true, routingTableEntryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RoutingTableEntry struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := RoutingTableEntries().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	routingTableEntryDBTypes = map[string]string{`RoutingTableSnapshotID`: `integer`, `PeerID`: `integer`, `Bucket`: `smallint`, `LastUsefulAt`: `timestamp with time zone`, `LastSuccessfulOutboundQueryAt`: `timestamp with time zone`, `AddedAt`: `timestamp with time zone`, `ConnectedSince`: `timestamp with time zone`}
	_                        = bytes.MinRead
)

func testRoutingTableEntriesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(routingTableEntryPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(routingTableEntryAllColumns) == len(routingTableEntryPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &RoutingTableEntry{}
	if err = randomize.Struct(seed, o, routingTableEntryDBTypes, true, routingTableEntryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RoutingTableEntry struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RoutingTableEntries().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, routingTableEntryDBTypes, true, routingTableEntryPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize RoutingTableEntry struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testRoutingTableEntriesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(routingTableEntryAllColumns) == len(routingTableEntryPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &RoutingTableEntry{}
	if err = randomize.Struct(seed, o, routingTableEntryDBTypes, true, routingTableEntryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RoutingTableEntry struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RoutingTableEntries().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, routingTableEntryDBTypes, true, routingTableEntryPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize RoutingTableEntry struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(routingTableEntryAllColumns, routingTableEntryPrimaryKeyColumns) {
		fields = routingTableEntryAllColumns
	} else {
		fields = strmangle.SetComplement(
			routingTableEntryAllColumns,
			routingTableEntryPrimaryKeyColumns,
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

	slice := RoutingTableEntrySlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testRoutingTableEntriesUpsert(t *testing.T) {
	t.Parallel()

	if len(routingTableEntryAllColumns) == len(routingTableEntryPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := RoutingTableEntry{}
	if err = randomize.Struct(seed, &o, routingTableEntryDBTypes, true); err != nil {
		t.Errorf("Unable to randomize RoutingTableEntry struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert RoutingTableEntry: %s", err)
	}

	count, err := RoutingTableEntries().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, routingTableEntryDBTypes, false, routingTableEntryPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize RoutingTableEntry struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert RoutingTableEntry: %s", err)
	}

	count, err = RoutingTableEntries().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
