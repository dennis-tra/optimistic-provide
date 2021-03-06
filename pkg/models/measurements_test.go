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

func testMeasurements(t *testing.T) {
	t.Parallel()

	query := Measurements()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testMeasurementsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Measurement{}
	if err = randomize.Struct(seed, o, measurementDBTypes, true, measurementColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Measurement struct: %s", err)
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

	count, err := Measurements().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testMeasurementsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Measurement{}
	if err = randomize.Struct(seed, o, measurementDBTypes, true, measurementColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Measurement struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Measurements().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Measurements().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testMeasurementsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Measurement{}
	if err = randomize.Struct(seed, o, measurementDBTypes, true, measurementColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Measurement struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := MeasurementSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Measurements().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testMeasurementsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Measurement{}
	if err = randomize.Struct(seed, o, measurementDBTypes, true, measurementColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Measurement struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := MeasurementExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Measurement exists: %s", err)
	}
	if !e {
		t.Errorf("Expected MeasurementExists to return true, but got false.")
	}
}

func testMeasurementsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Measurement{}
	if err = randomize.Struct(seed, o, measurementDBTypes, true, measurementColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Measurement struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	measurementFound, err := FindMeasurement(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if measurementFound == nil {
		t.Error("want a record, got nil")
	}
}

func testMeasurementsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Measurement{}
	if err = randomize.Struct(seed, o, measurementDBTypes, true, measurementColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Measurement struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Measurements().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testMeasurementsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Measurement{}
	if err = randomize.Struct(seed, o, measurementDBTypes, true, measurementColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Measurement struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Measurements().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testMeasurementsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	measurementOne := &Measurement{}
	measurementTwo := &Measurement{}
	if err = randomize.Struct(seed, measurementOne, measurementDBTypes, false, measurementColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Measurement struct: %s", err)
	}
	if err = randomize.Struct(seed, measurementTwo, measurementDBTypes, false, measurementColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Measurement struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = measurementOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = measurementTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Measurements().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testMeasurementsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	measurementOne := &Measurement{}
	measurementTwo := &Measurement{}
	if err = randomize.Struct(seed, measurementOne, measurementDBTypes, false, measurementColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Measurement struct: %s", err)
	}
	if err = randomize.Struct(seed, measurementTwo, measurementDBTypes, false, measurementColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Measurement struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = measurementOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = measurementTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Measurements().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func measurementBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Measurement) error {
	*o = Measurement{}
	return nil
}

func measurementAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Measurement) error {
	*o = Measurement{}
	return nil
}

func measurementAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Measurement) error {
	*o = Measurement{}
	return nil
}

func measurementBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Measurement) error {
	*o = Measurement{}
	return nil
}

func measurementAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Measurement) error {
	*o = Measurement{}
	return nil
}

func measurementBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Measurement) error {
	*o = Measurement{}
	return nil
}

func measurementAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Measurement) error {
	*o = Measurement{}
	return nil
}

func measurementBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Measurement) error {
	*o = Measurement{}
	return nil
}

func measurementAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Measurement) error {
	*o = Measurement{}
	return nil
}

func testMeasurementsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Measurement{}
	o := &Measurement{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, measurementDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Measurement object: %s", err)
	}

	AddMeasurementHook(boil.BeforeInsertHook, measurementBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	measurementBeforeInsertHooks = []MeasurementHook{}

	AddMeasurementHook(boil.AfterInsertHook, measurementAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	measurementAfterInsertHooks = []MeasurementHook{}

	AddMeasurementHook(boil.AfterSelectHook, measurementAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	measurementAfterSelectHooks = []MeasurementHook{}

	AddMeasurementHook(boil.BeforeUpdateHook, measurementBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	measurementBeforeUpdateHooks = []MeasurementHook{}

	AddMeasurementHook(boil.AfterUpdateHook, measurementAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	measurementAfterUpdateHooks = []MeasurementHook{}

	AddMeasurementHook(boil.BeforeDeleteHook, measurementBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	measurementBeforeDeleteHooks = []MeasurementHook{}

	AddMeasurementHook(boil.AfterDeleteHook, measurementAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	measurementAfterDeleteHooks = []MeasurementHook{}

	AddMeasurementHook(boil.BeforeUpsertHook, measurementBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	measurementBeforeUpsertHooks = []MeasurementHook{}

	AddMeasurementHook(boil.AfterUpsertHook, measurementAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	measurementAfterUpsertHooks = []MeasurementHook{}
}

func testMeasurementsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Measurement{}
	if err = randomize.Struct(seed, o, measurementDBTypes, true, measurementColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Measurement struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Measurements().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testMeasurementsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Measurement{}
	if err = randomize.Struct(seed, o, measurementDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Measurement struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(measurementColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Measurements().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testMeasurementToManyProvides(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Measurement
	var b, c Provide

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, measurementDBTypes, true, measurementColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Measurement struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, provideDBTypes, false, provideColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, provideDBTypes, false, provideColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&b.MeasurementID, a.ID)
	queries.Assign(&c.MeasurementID, a.ID)
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.Provides().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if queries.Equal(v.MeasurementID, b.MeasurementID) {
			bFound = true
		}
		if queries.Equal(v.MeasurementID, c.MeasurementID) {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := MeasurementSlice{&a}
	if err = a.L.LoadProvides(ctx, tx, false, (*[]*Measurement)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Provides); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Provides = nil
	if err = a.L.LoadProvides(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Provides); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testMeasurementToManyAddOpProvides(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Measurement
	var b, c, d, e Provide

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, measurementDBTypes, false, strmangle.SetComplement(measurementPrimaryKeyColumns, measurementColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Provide{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, provideDBTypes, false, strmangle.SetComplement(providePrimaryKeyColumns, provideColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*Provide{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddProvides(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if !queries.Equal(a.ID, first.MeasurementID) {
			t.Error("foreign key was wrong value", a.ID, first.MeasurementID)
		}
		if !queries.Equal(a.ID, second.MeasurementID) {
			t.Error("foreign key was wrong value", a.ID, second.MeasurementID)
		}

		if first.R.Measurement != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Measurement != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.Provides[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Provides[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Provides().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testMeasurementToManySetOpProvides(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Measurement
	var b, c, d, e Provide

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, measurementDBTypes, false, strmangle.SetComplement(measurementPrimaryKeyColumns, measurementColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Provide{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, provideDBTypes, false, strmangle.SetComplement(providePrimaryKeyColumns, provideColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.SetProvides(ctx, tx, false, &b, &c)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.Provides().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	err = a.SetProvides(ctx, tx, true, &d, &e)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.Provides().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.MeasurementID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.MeasurementID) {
		t.Error("want c's foreign key value to be nil")
	}
	if !queries.Equal(a.ID, d.MeasurementID) {
		t.Error("foreign key was wrong value", a.ID, d.MeasurementID)
	}
	if !queries.Equal(a.ID, e.MeasurementID) {
		t.Error("foreign key was wrong value", a.ID, e.MeasurementID)
	}

	if b.R.Measurement != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Measurement != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Measurement != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.Measurement != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}

	if a.R.Provides[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.Provides[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testMeasurementToManyRemoveOpProvides(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Measurement
	var b, c, d, e Provide

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, measurementDBTypes, false, strmangle.SetComplement(measurementPrimaryKeyColumns, measurementColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Provide{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, provideDBTypes, false, strmangle.SetComplement(providePrimaryKeyColumns, provideColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.AddProvides(ctx, tx, true, foreigners...)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.Provides().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 4 {
		t.Error("count was wrong:", count)
	}

	err = a.RemoveProvides(ctx, tx, foreigners[:2]...)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.Provides().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.MeasurementID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.MeasurementID) {
		t.Error("want c's foreign key value to be nil")
	}

	if b.R.Measurement != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Measurement != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Measurement != &a {
		t.Error("relationship to a should have been preserved")
	}
	if e.R.Measurement != &a {
		t.Error("relationship to a should have been preserved")
	}

	if len(a.R.Provides) != 2 {
		t.Error("should have preserved two relationships")
	}

	// Removal doesn't do a stable deletion for performance so we have to flip the order
	if a.R.Provides[1] != &d {
		t.Error("relationship to d should have been preserved")
	}
	if a.R.Provides[0] != &e {
		t.Error("relationship to e should have been preserved")
	}
}

func testMeasurementToOnePeerUsingHost(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Measurement
	var foreign Peer

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, measurementDBTypes, false, measurementColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Measurement struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, peerDBTypes, false, peerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Peer struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.HostID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Host().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := MeasurementSlice{&local}
	if err = local.L.LoadHost(ctx, tx, false, (*[]*Measurement)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Host == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Host = nil
	if err = local.L.LoadHost(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Host == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testMeasurementToOneSetOpPeerUsingHost(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Measurement
	var b, c Peer

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, measurementDBTypes, false, strmangle.SetComplement(measurementPrimaryKeyColumns, measurementColumnsWithoutDefault)...); err != nil {
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
		err = a.SetHost(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Host != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.HostMeasurements[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.HostID != x.ID {
			t.Error("foreign key was wrong value", a.HostID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.HostID))
		reflect.Indirect(reflect.ValueOf(&a.HostID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.HostID != x.ID {
			t.Error("foreign key was wrong value", a.HostID, x.ID)
		}
	}
}

func testMeasurementsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Measurement{}
	if err = randomize.Struct(seed, o, measurementDBTypes, true, measurementColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Measurement struct: %s", err)
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

func testMeasurementsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Measurement{}
	if err = randomize.Struct(seed, o, measurementDBTypes, true, measurementColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Measurement struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := MeasurementSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testMeasurementsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Measurement{}
	if err = randomize.Struct(seed, o, measurementDBTypes, true, measurementColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Measurement struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Measurements().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	measurementDBTypes = map[string]string{`ID`: `integer`, `HostID`: `integer`, `StartedAt`: `timestamp with time zone`, `EndedAt`: `timestamp with time zone`, `Configuration`: `json`, `UpdatedAt`: `timestamp with time zone`, `CreatedAt`: `timestamp with time zone`}
	_                  = bytes.MinRead
)

func testMeasurementsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(measurementPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(measurementAllColumns) == len(measurementPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Measurement{}
	if err = randomize.Struct(seed, o, measurementDBTypes, true, measurementColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Measurement struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Measurements().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, measurementDBTypes, true, measurementPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Measurement struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testMeasurementsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(measurementAllColumns) == len(measurementPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Measurement{}
	if err = randomize.Struct(seed, o, measurementDBTypes, true, measurementColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Measurement struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Measurements().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, measurementDBTypes, true, measurementPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Measurement struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(measurementAllColumns, measurementPrimaryKeyColumns) {
		fields = measurementAllColumns
	} else {
		fields = strmangle.SetComplement(
			measurementAllColumns,
			measurementPrimaryKeyColumns,
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

	slice := MeasurementSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testMeasurementsUpsert(t *testing.T) {
	t.Parallel()

	if len(measurementAllColumns) == len(measurementPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Measurement{}
	if err = randomize.Struct(seed, &o, measurementDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Measurement struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Measurement: %s", err)
	}

	count, err := Measurements().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, measurementDBTypes, false, measurementPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Measurement struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Measurement: %s", err)
	}

	count, err = Measurements().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
