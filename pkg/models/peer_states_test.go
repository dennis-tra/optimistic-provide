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

func testPeerStates(t *testing.T) {
	t.Parallel()

	query := PeerStates()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testPeerStatesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PeerState{}
	if err = randomize.Struct(seed, o, peerStateDBTypes, true, peerStateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PeerState struct: %s", err)
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

	count, err := PeerStates().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPeerStatesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PeerState{}
	if err = randomize.Struct(seed, o, peerStateDBTypes, true, peerStateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PeerState struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := PeerStates().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := PeerStates().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPeerStatesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PeerState{}
	if err = randomize.Struct(seed, o, peerStateDBTypes, true, peerStateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PeerState struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := PeerStateSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := PeerStates().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPeerStatesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PeerState{}
	if err = randomize.Struct(seed, o, peerStateDBTypes, true, peerStateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PeerState struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := PeerStateExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if PeerState exists: %s", err)
	}
	if !e {
		t.Errorf("Expected PeerStateExists to return true, but got false.")
	}
}

func testPeerStatesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PeerState{}
	if err = randomize.Struct(seed, o, peerStateDBTypes, true, peerStateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PeerState struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	peerStateFound, err := FindPeerState(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if peerStateFound == nil {
		t.Error("want a record, got nil")
	}
}

func testPeerStatesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PeerState{}
	if err = randomize.Struct(seed, o, peerStateDBTypes, true, peerStateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PeerState struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = PeerStates().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testPeerStatesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PeerState{}
	if err = randomize.Struct(seed, o, peerStateDBTypes, true, peerStateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PeerState struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := PeerStates().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testPeerStatesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	peerStateOne := &PeerState{}
	peerStateTwo := &PeerState{}
	if err = randomize.Struct(seed, peerStateOne, peerStateDBTypes, false, peerStateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PeerState struct: %s", err)
	}
	if err = randomize.Struct(seed, peerStateTwo, peerStateDBTypes, false, peerStateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PeerState struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = peerStateOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = peerStateTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := PeerStates().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testPeerStatesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	peerStateOne := &PeerState{}
	peerStateTwo := &PeerState{}
	if err = randomize.Struct(seed, peerStateOne, peerStateDBTypes, false, peerStateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PeerState struct: %s", err)
	}
	if err = randomize.Struct(seed, peerStateTwo, peerStateDBTypes, false, peerStateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PeerState struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = peerStateOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = peerStateTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := PeerStates().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func peerStateBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *PeerState) error {
	*o = PeerState{}
	return nil
}

func peerStateAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *PeerState) error {
	*o = PeerState{}
	return nil
}

func peerStateAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *PeerState) error {
	*o = PeerState{}
	return nil
}

func peerStateBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *PeerState) error {
	*o = PeerState{}
	return nil
}

func peerStateAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *PeerState) error {
	*o = PeerState{}
	return nil
}

func peerStateBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *PeerState) error {
	*o = PeerState{}
	return nil
}

func peerStateAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *PeerState) error {
	*o = PeerState{}
	return nil
}

func peerStateBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *PeerState) error {
	*o = PeerState{}
	return nil
}

func peerStateAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *PeerState) error {
	*o = PeerState{}
	return nil
}

func testPeerStatesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &PeerState{}
	o := &PeerState{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, peerStateDBTypes, false); err != nil {
		t.Errorf("Unable to randomize PeerState object: %s", err)
	}

	AddPeerStateHook(boil.BeforeInsertHook, peerStateBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	peerStateBeforeInsertHooks = []PeerStateHook{}

	AddPeerStateHook(boil.AfterInsertHook, peerStateAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	peerStateAfterInsertHooks = []PeerStateHook{}

	AddPeerStateHook(boil.AfterSelectHook, peerStateAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	peerStateAfterSelectHooks = []PeerStateHook{}

	AddPeerStateHook(boil.BeforeUpdateHook, peerStateBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	peerStateBeforeUpdateHooks = []PeerStateHook{}

	AddPeerStateHook(boil.AfterUpdateHook, peerStateAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	peerStateAfterUpdateHooks = []PeerStateHook{}

	AddPeerStateHook(boil.BeforeDeleteHook, peerStateBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	peerStateBeforeDeleteHooks = []PeerStateHook{}

	AddPeerStateHook(boil.AfterDeleteHook, peerStateAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	peerStateAfterDeleteHooks = []PeerStateHook{}

	AddPeerStateHook(boil.BeforeUpsertHook, peerStateBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	peerStateBeforeUpsertHooks = []PeerStateHook{}

	AddPeerStateHook(boil.AfterUpsertHook, peerStateAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	peerStateAfterUpsertHooks = []PeerStateHook{}
}

func testPeerStatesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PeerState{}
	if err = randomize.Struct(seed, o, peerStateDBTypes, true, peerStateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PeerState struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := PeerStates().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testPeerStatesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PeerState{}
	if err = randomize.Struct(seed, o, peerStateDBTypes, true); err != nil {
		t.Errorf("Unable to randomize PeerState struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(peerStateColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := PeerStates().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testPeerStateToOnePeerUsingPeer(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local PeerState
	var foreign Peer

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, peerStateDBTypes, false, peerStateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PeerState struct: %s", err)
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

	slice := PeerStateSlice{&local}
	if err = local.L.LoadPeer(ctx, tx, false, (*[]*PeerState)(&slice), nil); err != nil {
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

func testPeerStateToOneProvideUsingProvide(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local PeerState
	var foreign Provide

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, peerStateDBTypes, true, peerStateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PeerState struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, provideDBTypes, false, provideColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Provide struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&local.ProvideID, foreign.ID)
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Provide().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if !queries.Equal(check.ID, foreign.ID) {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := PeerStateSlice{&local}
	if err = local.L.LoadProvide(ctx, tx, false, (*[]*PeerState)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Provide == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Provide = nil
	if err = local.L.LoadProvide(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Provide == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testPeerStateToOnePeerUsingReferrer(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local PeerState
	var foreign Peer

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, peerStateDBTypes, false, peerStateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PeerState struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, peerDBTypes, false, peerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Peer struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.ReferrerID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Referrer().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := PeerStateSlice{&local}
	if err = local.L.LoadReferrer(ctx, tx, false, (*[]*PeerState)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Referrer == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Referrer = nil
	if err = local.L.LoadReferrer(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Referrer == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testPeerStateToOneRetrievalUsingRetrieval(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local PeerState
	var foreign Retrieval

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, peerStateDBTypes, true, peerStateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PeerState struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, retrievalDBTypes, false, retrievalColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Retrieval struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&local.RetrievalID, foreign.ID)
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Retrieval().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if !queries.Equal(check.ID, foreign.ID) {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := PeerStateSlice{&local}
	if err = local.L.LoadRetrieval(ctx, tx, false, (*[]*PeerState)(&slice), nil); err != nil {
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

func testPeerStateToOneSetOpPeerUsingPeer(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a PeerState
	var b, c Peer

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, peerStateDBTypes, false, strmangle.SetComplement(peerStatePrimaryKeyColumns, peerStateColumnsWithoutDefault)...); err != nil {
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

		if x.R.PeerStates[0] != &a {
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
func testPeerStateToOneSetOpProvideUsingProvide(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a PeerState
	var b, c Provide

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, peerStateDBTypes, false, strmangle.SetComplement(peerStatePrimaryKeyColumns, peerStateColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, provideDBTypes, false, strmangle.SetComplement(providePrimaryKeyColumns, provideColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, provideDBTypes, false, strmangle.SetComplement(providePrimaryKeyColumns, provideColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Provide{&b, &c} {
		err = a.SetProvide(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Provide != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.PeerStates[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if !queries.Equal(a.ProvideID, x.ID) {
			t.Error("foreign key was wrong value", a.ProvideID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.ProvideID))
		reflect.Indirect(reflect.ValueOf(&a.ProvideID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if !queries.Equal(a.ProvideID, x.ID) {
			t.Error("foreign key was wrong value", a.ProvideID, x.ID)
		}
	}
}

func testPeerStateToOneRemoveOpProvideUsingProvide(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a PeerState
	var b Provide

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, peerStateDBTypes, false, strmangle.SetComplement(peerStatePrimaryKeyColumns, peerStateColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, provideDBTypes, false, strmangle.SetComplement(providePrimaryKeyColumns, provideColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = a.SetProvide(ctx, tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemoveProvide(ctx, tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.Provide().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.Provide != nil {
		t.Error("R struct entry should be nil")
	}

	if !queries.IsValuerNil(a.ProvideID) {
		t.Error("foreign key value should be nil")
	}

	if len(b.R.PeerStates) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testPeerStateToOneSetOpPeerUsingReferrer(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a PeerState
	var b, c Peer

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, peerStateDBTypes, false, strmangle.SetComplement(peerStatePrimaryKeyColumns, peerStateColumnsWithoutDefault)...); err != nil {
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
		err = a.SetReferrer(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Referrer != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.ReferrerPeerStates[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.ReferrerID != x.ID {
			t.Error("foreign key was wrong value", a.ReferrerID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.ReferrerID))
		reflect.Indirect(reflect.ValueOf(&a.ReferrerID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.ReferrerID != x.ID {
			t.Error("foreign key was wrong value", a.ReferrerID, x.ID)
		}
	}
}
func testPeerStateToOneSetOpRetrievalUsingRetrieval(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a PeerState
	var b, c Retrieval

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, peerStateDBTypes, false, strmangle.SetComplement(peerStatePrimaryKeyColumns, peerStateColumnsWithoutDefault)...); err != nil {
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

		if x.R.PeerStates[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if !queries.Equal(a.RetrievalID, x.ID) {
			t.Error("foreign key was wrong value", a.RetrievalID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.RetrievalID))
		reflect.Indirect(reflect.ValueOf(&a.RetrievalID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if !queries.Equal(a.RetrievalID, x.ID) {
			t.Error("foreign key was wrong value", a.RetrievalID, x.ID)
		}
	}
}

func testPeerStateToOneRemoveOpRetrievalUsingRetrieval(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a PeerState
	var b Retrieval

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, peerStateDBTypes, false, strmangle.SetComplement(peerStatePrimaryKeyColumns, peerStateColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, retrievalDBTypes, false, strmangle.SetComplement(retrievalPrimaryKeyColumns, retrievalColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = a.SetRetrieval(ctx, tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemoveRetrieval(ctx, tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.Retrieval().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.Retrieval != nil {
		t.Error("R struct entry should be nil")
	}

	if !queries.IsValuerNil(a.RetrievalID) {
		t.Error("foreign key value should be nil")
	}

	if len(b.R.PeerStates) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testPeerStatesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PeerState{}
	if err = randomize.Struct(seed, o, peerStateDBTypes, true, peerStateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PeerState struct: %s", err)
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

func testPeerStatesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PeerState{}
	if err = randomize.Struct(seed, o, peerStateDBTypes, true, peerStateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PeerState struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := PeerStateSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testPeerStatesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PeerState{}
	if err = randomize.Struct(seed, o, peerStateDBTypes, true, peerStateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PeerState struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := PeerStates().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	peerStateDBTypes = map[string]string{`ID`: `integer`, `ProvideID`: `integer`, `RetrievalID`: `integer`, `PeerID`: `integer`, `ReferrerID`: `integer`, `State`: `enum.peer_state('HEARD','WAITING','QUERIED','UNREACHABLE')`, `Distance`: `bytea`}
	_                = bytes.MinRead
)

func testPeerStatesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(peerStatePrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(peerStateAllColumns) == len(peerStatePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &PeerState{}
	if err = randomize.Struct(seed, o, peerStateDBTypes, true, peerStateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PeerState struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := PeerStates().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, peerStateDBTypes, true, peerStatePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize PeerState struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testPeerStatesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(peerStateAllColumns) == len(peerStatePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &PeerState{}
	if err = randomize.Struct(seed, o, peerStateDBTypes, true, peerStateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PeerState struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := PeerStates().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, peerStateDBTypes, true, peerStatePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize PeerState struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(peerStateAllColumns, peerStatePrimaryKeyColumns) {
		fields = peerStateAllColumns
	} else {
		fields = strmangle.SetComplement(
			peerStateAllColumns,
			peerStatePrimaryKeyColumns,
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

	slice := PeerStateSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testPeerStatesUpsert(t *testing.T) {
	t.Parallel()

	if len(peerStateAllColumns) == len(peerStatePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := PeerState{}
	if err = randomize.Struct(seed, &o, peerStateDBTypes, true); err != nil {
		t.Errorf("Unable to randomize PeerState struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert PeerState: %s", err)
	}

	count, err := PeerStates().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, peerStateDBTypes, false, peerStatePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize PeerState struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert PeerState: %s", err)
	}

	count, err = PeerStates().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
