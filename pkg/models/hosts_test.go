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

func testHosts(t *testing.T) {
	t.Parallel()

	query := Hosts()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testHostsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Host{}
	if err = randomize.Struct(seed, o, hostDBTypes, true, hostColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Host struct: %s", err)
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

	count, err := Hosts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testHostsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Host{}
	if err = randomize.Struct(seed, o, hostDBTypes, true, hostColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Host struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Hosts().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Hosts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testHostsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Host{}
	if err = randomize.Struct(seed, o, hostDBTypes, true, hostColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Host struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := HostSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Hosts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testHostsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Host{}
	if err = randomize.Struct(seed, o, hostDBTypes, true, hostColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Host struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := HostExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Host exists: %s", err)
	}
	if !e {
		t.Errorf("Expected HostExists to return true, but got false.")
	}
}

func testHostsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Host{}
	if err = randomize.Struct(seed, o, hostDBTypes, true, hostColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Host struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	hostFound, err := FindHost(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if hostFound == nil {
		t.Error("want a record, got nil")
	}
}

func testHostsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Host{}
	if err = randomize.Struct(seed, o, hostDBTypes, true, hostColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Host struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Hosts().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testHostsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Host{}
	if err = randomize.Struct(seed, o, hostDBTypes, true, hostColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Host struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Hosts().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testHostsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	hostOne := &Host{}
	hostTwo := &Host{}
	if err = randomize.Struct(seed, hostOne, hostDBTypes, false, hostColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Host struct: %s", err)
	}
	if err = randomize.Struct(seed, hostTwo, hostDBTypes, false, hostColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Host struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = hostOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = hostTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Hosts().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testHostsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	hostOne := &Host{}
	hostTwo := &Host{}
	if err = randomize.Struct(seed, hostOne, hostDBTypes, false, hostColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Host struct: %s", err)
	}
	if err = randomize.Struct(seed, hostTwo, hostDBTypes, false, hostColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Host struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = hostOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = hostTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Hosts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func hostBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Host) error {
	*o = Host{}
	return nil
}

func hostAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Host) error {
	*o = Host{}
	return nil
}

func hostAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Host) error {
	*o = Host{}
	return nil
}

func hostBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Host) error {
	*o = Host{}
	return nil
}

func hostAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Host) error {
	*o = Host{}
	return nil
}

func hostBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Host) error {
	*o = Host{}
	return nil
}

func hostAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Host) error {
	*o = Host{}
	return nil
}

func hostBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Host) error {
	*o = Host{}
	return nil
}

func hostAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Host) error {
	*o = Host{}
	return nil
}

func testHostsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Host{}
	o := &Host{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, hostDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Host object: %s", err)
	}

	AddHostHook(boil.BeforeInsertHook, hostBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	hostBeforeInsertHooks = []HostHook{}

	AddHostHook(boil.AfterInsertHook, hostAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	hostAfterInsertHooks = []HostHook{}

	AddHostHook(boil.AfterSelectHook, hostAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	hostAfterSelectHooks = []HostHook{}

	AddHostHook(boil.BeforeUpdateHook, hostBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	hostBeforeUpdateHooks = []HostHook{}

	AddHostHook(boil.AfterUpdateHook, hostAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	hostAfterUpdateHooks = []HostHook{}

	AddHostHook(boil.BeforeDeleteHook, hostBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	hostBeforeDeleteHooks = []HostHook{}

	AddHostHook(boil.AfterDeleteHook, hostAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	hostAfterDeleteHooks = []HostHook{}

	AddHostHook(boil.BeforeUpsertHook, hostBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	hostBeforeUpsertHooks = []HostHook{}

	AddHostHook(boil.AfterUpsertHook, hostAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	hostAfterUpsertHooks = []HostHook{}
}

func testHostsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Host{}
	if err = randomize.Struct(seed, o, hostDBTypes, true, hostColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Host struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Hosts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testHostsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Host{}
	if err = randomize.Struct(seed, o, hostDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Host struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(hostColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Hosts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testHostToManyNetworkSizeEstimates(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Host
	var b, c NetworkSizeEstimate

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, hostDBTypes, true, hostColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Host struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, networkSizeEstimateDBTypes, false, networkSizeEstimateColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, networkSizeEstimateDBTypes, false, networkSizeEstimateColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.HostID = a.ID
	c.HostID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.NetworkSizeEstimates().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.HostID == b.HostID {
			bFound = true
		}
		if v.HostID == c.HostID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := HostSlice{&a}
	if err = a.L.LoadNetworkSizeEstimates(ctx, tx, false, (*[]*Host)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.NetworkSizeEstimates); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.NetworkSizeEstimates = nil
	if err = a.L.LoadNetworkSizeEstimates(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.NetworkSizeEstimates); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testHostToManyAddOpNetworkSizeEstimates(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Host
	var b, c, d, e NetworkSizeEstimate

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, hostDBTypes, false, strmangle.SetComplement(hostPrimaryKeyColumns, hostColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*NetworkSizeEstimate{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, networkSizeEstimateDBTypes, false, strmangle.SetComplement(networkSizeEstimatePrimaryKeyColumns, networkSizeEstimateColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*NetworkSizeEstimate{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddNetworkSizeEstimates(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.HostID {
			t.Error("foreign key was wrong value", a.ID, first.HostID)
		}
		if a.ID != second.HostID {
			t.Error("foreign key was wrong value", a.ID, second.HostID)
		}

		if first.R.Host != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Host != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.NetworkSizeEstimates[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.NetworkSizeEstimates[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.NetworkSizeEstimates().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testHostToOnePeerUsingPeer(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Host
	var foreign Peer

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, hostDBTypes, false, hostColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Host struct: %s", err)
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

	slice := HostSlice{&local}
	if err = local.L.LoadPeer(ctx, tx, false, (*[]*Host)(&slice), nil); err != nil {
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

func testHostToOneSetOpPeerUsingPeer(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Host
	var b, c Peer

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, hostDBTypes, false, strmangle.SetComplement(hostPrimaryKeyColumns, hostColumnsWithoutDefault)...); err != nil {
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

		if x.R.Hosts[0] != &a {
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

func testHostsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Host{}
	if err = randomize.Struct(seed, o, hostDBTypes, true, hostColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Host struct: %s", err)
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

func testHostsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Host{}
	if err = randomize.Struct(seed, o, hostDBTypes, true, hostColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Host struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := HostSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testHostsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Host{}
	if err = randomize.Struct(seed, o, hostDBTypes, true, hostColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Host struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Hosts().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	hostDBTypes = map[string]string{`ID`: `integer`, `PeerID`: `integer`, `Name`: `text`, `PrivateKey`: `bytea`, `CreatedAt`: `timestamp with time zone`, `UpdatedAt`: `timestamp with time zone`, `ArchivedAt`: `timestamp with time zone`, `Network`: `enum.network_type('IPFS','FILECOIN','POLKADOT','KUSAMA')`}
	_           = bytes.MinRead
)

func testHostsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(hostPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(hostAllColumns) == len(hostPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Host{}
	if err = randomize.Struct(seed, o, hostDBTypes, true, hostColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Host struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Hosts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, hostDBTypes, true, hostPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Host struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testHostsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(hostAllColumns) == len(hostPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Host{}
	if err = randomize.Struct(seed, o, hostDBTypes, true, hostColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Host struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Hosts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, hostDBTypes, true, hostPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Host struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(hostAllColumns, hostPrimaryKeyColumns) {
		fields = hostAllColumns
	} else {
		fields = strmangle.SetComplement(
			hostAllColumns,
			hostPrimaryKeyColumns,
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

	slice := HostSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testHostsUpsert(t *testing.T) {
	t.Parallel()

	if len(hostAllColumns) == len(hostPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Host{}
	if err = randomize.Struct(seed, &o, hostDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Host struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Host: %s", err)
	}

	count, err := Hosts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, hostDBTypes, false, hostPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Host struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Host: %s", err)
	}

	count, err = Hosts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
