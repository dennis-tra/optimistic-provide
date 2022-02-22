// Code generated by SQLBoiler 4.6.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// PeerState is an object representing the database table.
type PeerState struct {
	ProvideID  int    `boil:"provide_id" json:"provide_id" toml:"provide_id" yaml:"provide_id"`
	PeerID     int    `boil:"peer_id" json:"peer_id" toml:"peer_id" yaml:"peer_id"`
	ReferrerID int    `boil:"referrer_id" json:"referrer_id" toml:"referrer_id" yaml:"referrer_id"`
	State      string `boil:"state" json:"state" toml:"state" yaml:"state"`
	Distance   []byte `boil:"distance" json:"distance" toml:"distance" yaml:"distance"`

	R *peerStateR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L peerStateL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var PeerStateColumns = struct {
	ProvideID  string
	PeerID     string
	ReferrerID string
	State      string
	Distance   string
}{
	ProvideID:  "provide_id",
	PeerID:     "peer_id",
	ReferrerID: "referrer_id",
	State:      "state",
	Distance:   "distance",
}

var PeerStateTableColumns = struct {
	ProvideID  string
	PeerID     string
	ReferrerID string
	State      string
	Distance   string
}{
	ProvideID:  "peer_states.provide_id",
	PeerID:     "peer_states.peer_id",
	ReferrerID: "peer_states.referrer_id",
	State:      "peer_states.state",
	Distance:   "peer_states.distance",
}

// Generated where

var PeerStateWhere = struct {
	ProvideID  whereHelperint
	PeerID     whereHelperint
	ReferrerID whereHelperint
	State      whereHelperstring
	Distance   whereHelper__byte
}{
	ProvideID:  whereHelperint{field: "\"peer_states\".\"provide_id\""},
	PeerID:     whereHelperint{field: "\"peer_states\".\"peer_id\""},
	ReferrerID: whereHelperint{field: "\"peer_states\".\"referrer_id\""},
	State:      whereHelperstring{field: "\"peer_states\".\"state\""},
	Distance:   whereHelper__byte{field: "\"peer_states\".\"distance\""},
}

// PeerStateRels is where relationship names are stored.
var PeerStateRels = struct {
	Peer     string
	Provide  string
	Referrer string
}{
	Peer:     "Peer",
	Provide:  "Provide",
	Referrer: "Referrer",
}

// peerStateR is where relationships are stored.
type peerStateR struct {
	Peer     *Peer    `boil:"Peer" json:"Peer" toml:"Peer" yaml:"Peer"`
	Provide  *Provide `boil:"Provide" json:"Provide" toml:"Provide" yaml:"Provide"`
	Referrer *Peer    `boil:"Referrer" json:"Referrer" toml:"Referrer" yaml:"Referrer"`
}

// NewStruct creates a new relationship struct
func (*peerStateR) NewStruct() *peerStateR {
	return &peerStateR{}
}

// peerStateL is where Load methods for each relationship are stored.
type peerStateL struct{}

var (
	peerStateAllColumns            = []string{"provide_id", "peer_id", "referrer_id", "state", "distance"}
	peerStateColumnsWithoutDefault = []string{"provide_id", "peer_id", "referrer_id", "state", "distance"}
	peerStateColumnsWithDefault    = []string{}
	peerStatePrimaryKeyColumns     = []string{"provide_id", "peer_id"}
)

type (
	// PeerStateSlice is an alias for a slice of pointers to PeerState.
	// This should almost always be used instead of []PeerState.
	PeerStateSlice []*PeerState
	// PeerStateHook is the signature for custom PeerState hook methods
	PeerStateHook func(context.Context, boil.ContextExecutor, *PeerState) error

	peerStateQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	peerStateType                 = reflect.TypeOf(&PeerState{})
	peerStateMapping              = queries.MakeStructMapping(peerStateType)
	peerStatePrimaryKeyMapping, _ = queries.BindMapping(peerStateType, peerStateMapping, peerStatePrimaryKeyColumns)
	peerStateInsertCacheMut       sync.RWMutex
	peerStateInsertCache          = make(map[string]insertCache)
	peerStateUpdateCacheMut       sync.RWMutex
	peerStateUpdateCache          = make(map[string]updateCache)
	peerStateUpsertCacheMut       sync.RWMutex
	peerStateUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var peerStateBeforeInsertHooks []PeerStateHook
var peerStateBeforeUpdateHooks []PeerStateHook
var peerStateBeforeDeleteHooks []PeerStateHook
var peerStateBeforeUpsertHooks []PeerStateHook

var peerStateAfterInsertHooks []PeerStateHook
var peerStateAfterSelectHooks []PeerStateHook
var peerStateAfterUpdateHooks []PeerStateHook
var peerStateAfterDeleteHooks []PeerStateHook
var peerStateAfterUpsertHooks []PeerStateHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *PeerState) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range peerStateBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *PeerState) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range peerStateBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *PeerState) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range peerStateBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *PeerState) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range peerStateBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *PeerState) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range peerStateAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *PeerState) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range peerStateAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *PeerState) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range peerStateAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *PeerState) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range peerStateAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *PeerState) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range peerStateAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddPeerStateHook registers your hook function for all future operations.
func AddPeerStateHook(hookPoint boil.HookPoint, peerStateHook PeerStateHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		peerStateBeforeInsertHooks = append(peerStateBeforeInsertHooks, peerStateHook)
	case boil.BeforeUpdateHook:
		peerStateBeforeUpdateHooks = append(peerStateBeforeUpdateHooks, peerStateHook)
	case boil.BeforeDeleteHook:
		peerStateBeforeDeleteHooks = append(peerStateBeforeDeleteHooks, peerStateHook)
	case boil.BeforeUpsertHook:
		peerStateBeforeUpsertHooks = append(peerStateBeforeUpsertHooks, peerStateHook)
	case boil.AfterInsertHook:
		peerStateAfterInsertHooks = append(peerStateAfterInsertHooks, peerStateHook)
	case boil.AfterSelectHook:
		peerStateAfterSelectHooks = append(peerStateAfterSelectHooks, peerStateHook)
	case boil.AfterUpdateHook:
		peerStateAfterUpdateHooks = append(peerStateAfterUpdateHooks, peerStateHook)
	case boil.AfterDeleteHook:
		peerStateAfterDeleteHooks = append(peerStateAfterDeleteHooks, peerStateHook)
	case boil.AfterUpsertHook:
		peerStateAfterUpsertHooks = append(peerStateAfterUpsertHooks, peerStateHook)
	}
}

// One returns a single peerState record from the query.
func (q peerStateQuery) One(ctx context.Context, exec boil.ContextExecutor) (*PeerState, error) {
	o := &PeerState{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for peer_states")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all PeerState records from the query.
func (q peerStateQuery) All(ctx context.Context, exec boil.ContextExecutor) (PeerStateSlice, error) {
	var o []*PeerState

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to PeerState slice")
	}

	if len(peerStateAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all PeerState records in the query.
func (q peerStateQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count peer_states rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q peerStateQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if peer_states exists")
	}

	return count > 0, nil
}

// Peer pointed to by the foreign key.
func (o *PeerState) Peer(mods ...qm.QueryMod) peerQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.PeerID),
	}

	queryMods = append(queryMods, mods...)

	query := Peers(queryMods...)
	queries.SetFrom(query.Query, "\"peers\"")

	return query
}

// Provide pointed to by the foreign key.
func (o *PeerState) Provide(mods ...qm.QueryMod) provideQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.ProvideID),
	}

	queryMods = append(queryMods, mods...)

	query := Provides(queryMods...)
	queries.SetFrom(query.Query, "\"provides\"")

	return query
}

// Referrer pointed to by the foreign key.
func (o *PeerState) Referrer(mods ...qm.QueryMod) peerQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.ReferrerID),
	}

	queryMods = append(queryMods, mods...)

	query := Peers(queryMods...)
	queries.SetFrom(query.Query, "\"peers\"")

	return query
}

// LoadPeer allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (peerStateL) LoadPeer(ctx context.Context, e boil.ContextExecutor, singular bool, maybePeerState interface{}, mods queries.Applicator) error {
	var slice []*PeerState
	var object *PeerState

	if singular {
		object = maybePeerState.(*PeerState)
	} else {
		slice = *maybePeerState.(*[]*PeerState)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &peerStateR{}
		}
		args = append(args, object.PeerID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &peerStateR{}
			}

			for _, a := range args {
				if a == obj.PeerID {
					continue Outer
				}
			}

			args = append(args, obj.PeerID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`peers`),
		qm.WhereIn(`peers.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Peer")
	}

	var resultSlice []*Peer
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Peer")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for peers")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for peers")
	}

	if len(peerStateAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Peer = foreign
		if foreign.R == nil {
			foreign.R = &peerR{}
		}
		foreign.R.PeerStates = append(foreign.R.PeerStates, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.PeerID == foreign.ID {
				local.R.Peer = foreign
				if foreign.R == nil {
					foreign.R = &peerR{}
				}
				foreign.R.PeerStates = append(foreign.R.PeerStates, local)
				break
			}
		}
	}

	return nil
}

// LoadProvide allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (peerStateL) LoadProvide(ctx context.Context, e boil.ContextExecutor, singular bool, maybePeerState interface{}, mods queries.Applicator) error {
	var slice []*PeerState
	var object *PeerState

	if singular {
		object = maybePeerState.(*PeerState)
	} else {
		slice = *maybePeerState.(*[]*PeerState)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &peerStateR{}
		}
		args = append(args, object.ProvideID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &peerStateR{}
			}

			for _, a := range args {
				if a == obj.ProvideID {
					continue Outer
				}
			}

			args = append(args, obj.ProvideID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`provides`),
		qm.WhereIn(`provides.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Provide")
	}

	var resultSlice []*Provide
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Provide")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for provides")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for provides")
	}

	if len(peerStateAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Provide = foreign
		if foreign.R == nil {
			foreign.R = &provideR{}
		}
		foreign.R.PeerStates = append(foreign.R.PeerStates, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.ProvideID == foreign.ID {
				local.R.Provide = foreign
				if foreign.R == nil {
					foreign.R = &provideR{}
				}
				foreign.R.PeerStates = append(foreign.R.PeerStates, local)
				break
			}
		}
	}

	return nil
}

// LoadReferrer allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (peerStateL) LoadReferrer(ctx context.Context, e boil.ContextExecutor, singular bool, maybePeerState interface{}, mods queries.Applicator) error {
	var slice []*PeerState
	var object *PeerState

	if singular {
		object = maybePeerState.(*PeerState)
	} else {
		slice = *maybePeerState.(*[]*PeerState)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &peerStateR{}
		}
		args = append(args, object.ReferrerID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &peerStateR{}
			}

			for _, a := range args {
				if a == obj.ReferrerID {
					continue Outer
				}
			}

			args = append(args, obj.ReferrerID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`peers`),
		qm.WhereIn(`peers.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Peer")
	}

	var resultSlice []*Peer
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Peer")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for peers")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for peers")
	}

	if len(peerStateAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Referrer = foreign
		if foreign.R == nil {
			foreign.R = &peerR{}
		}
		foreign.R.ReferrerPeerStates = append(foreign.R.ReferrerPeerStates, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.ReferrerID == foreign.ID {
				local.R.Referrer = foreign
				if foreign.R == nil {
					foreign.R = &peerR{}
				}
				foreign.R.ReferrerPeerStates = append(foreign.R.ReferrerPeerStates, local)
				break
			}
		}
	}

	return nil
}

// SetPeer of the peerState to the related item.
// Sets o.R.Peer to related.
// Adds o to related.R.PeerStates.
func (o *PeerState) SetPeer(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Peer) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"peer_states\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"peer_id"}),
		strmangle.WhereClause("\"", "\"", 2, peerStatePrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ProvideID, o.PeerID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.PeerID = related.ID
	if o.R == nil {
		o.R = &peerStateR{
			Peer: related,
		}
	} else {
		o.R.Peer = related
	}

	if related.R == nil {
		related.R = &peerR{
			PeerStates: PeerStateSlice{o},
		}
	} else {
		related.R.PeerStates = append(related.R.PeerStates, o)
	}

	return nil
}

// SetProvide of the peerState to the related item.
// Sets o.R.Provide to related.
// Adds o to related.R.PeerStates.
func (o *PeerState) SetProvide(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Provide) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"peer_states\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"provide_id"}),
		strmangle.WhereClause("\"", "\"", 2, peerStatePrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ProvideID, o.PeerID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.ProvideID = related.ID
	if o.R == nil {
		o.R = &peerStateR{
			Provide: related,
		}
	} else {
		o.R.Provide = related
	}

	if related.R == nil {
		related.R = &provideR{
			PeerStates: PeerStateSlice{o},
		}
	} else {
		related.R.PeerStates = append(related.R.PeerStates, o)
	}

	return nil
}

// SetReferrer of the peerState to the related item.
// Sets o.R.Referrer to related.
// Adds o to related.R.ReferrerPeerStates.
func (o *PeerState) SetReferrer(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Peer) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"peer_states\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"referrer_id"}),
		strmangle.WhereClause("\"", "\"", 2, peerStatePrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ProvideID, o.PeerID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.ReferrerID = related.ID
	if o.R == nil {
		o.R = &peerStateR{
			Referrer: related,
		}
	} else {
		o.R.Referrer = related
	}

	if related.R == nil {
		related.R = &peerR{
			ReferrerPeerStates: PeerStateSlice{o},
		}
	} else {
		related.R.ReferrerPeerStates = append(related.R.ReferrerPeerStates, o)
	}

	return nil
}

// PeerStates retrieves all the records using an executor.
func PeerStates(mods ...qm.QueryMod) peerStateQuery {
	mods = append(mods, qm.From("\"peer_states\""))
	return peerStateQuery{NewQuery(mods...)}
}

// FindPeerState retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindPeerState(ctx context.Context, exec boil.ContextExecutor, provideID int, peerID int, selectCols ...string) (*PeerState, error) {
	peerStateObj := &PeerState{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"peer_states\" where \"provide_id\"=$1 AND \"peer_id\"=$2", sel,
	)

	q := queries.Raw(query, provideID, peerID)

	err := q.Bind(ctx, exec, peerStateObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from peer_states")
	}

	if err = peerStateObj.doAfterSelectHooks(ctx, exec); err != nil {
		return peerStateObj, err
	}

	return peerStateObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *PeerState) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no peer_states provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(peerStateColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	peerStateInsertCacheMut.RLock()
	cache, cached := peerStateInsertCache[key]
	peerStateInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			peerStateAllColumns,
			peerStateColumnsWithDefault,
			peerStateColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(peerStateType, peerStateMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(peerStateType, peerStateMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"peer_states\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"peer_states\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into peer_states")
	}

	if !cached {
		peerStateInsertCacheMut.Lock()
		peerStateInsertCache[key] = cache
		peerStateInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the PeerState.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *PeerState) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	peerStateUpdateCacheMut.RLock()
	cache, cached := peerStateUpdateCache[key]
	peerStateUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			peerStateAllColumns,
			peerStatePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update peer_states, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"peer_states\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, peerStatePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(peerStateType, peerStateMapping, append(wl, peerStatePrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update peer_states row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for peer_states")
	}

	if !cached {
		peerStateUpdateCacheMut.Lock()
		peerStateUpdateCache[key] = cache
		peerStateUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q peerStateQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for peer_states")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for peer_states")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o PeerStateSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), peerStatePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"peer_states\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, peerStatePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in peerState slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all peerState")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *PeerState) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no peer_states provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(peerStateColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	peerStateUpsertCacheMut.RLock()
	cache, cached := peerStateUpsertCache[key]
	peerStateUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			peerStateAllColumns,
			peerStateColumnsWithDefault,
			peerStateColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			peerStateAllColumns,
			peerStatePrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert peer_states, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(peerStatePrimaryKeyColumns))
			copy(conflict, peerStatePrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"peer_states\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(peerStateType, peerStateMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(peerStateType, peerStateMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert peer_states")
	}

	if !cached {
		peerStateUpsertCacheMut.Lock()
		peerStateUpsertCache[key] = cache
		peerStateUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single PeerState record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *PeerState) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no PeerState provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), peerStatePrimaryKeyMapping)
	sql := "DELETE FROM \"peer_states\" WHERE \"provide_id\"=$1 AND \"peer_id\"=$2"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from peer_states")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for peer_states")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q peerStateQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no peerStateQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from peer_states")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for peer_states")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o PeerStateSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(peerStateBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), peerStatePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"peer_states\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, peerStatePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from peerState slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for peer_states")
	}

	if len(peerStateAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *PeerState) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindPeerState(ctx, exec, o.ProvideID, o.PeerID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *PeerStateSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := PeerStateSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), peerStatePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"peer_states\".* FROM \"peer_states\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, peerStatePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in PeerStateSlice")
	}

	*o = slice

	return nil
}

// PeerStateExists checks if the PeerState row exists.
func PeerStateExists(ctx context.Context, exec boil.ContextExecutor, provideID int, peerID int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"peer_states\" where \"provide_id\"=$1 AND \"peer_id\"=$2 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, provideID, peerID)
	}
	row := exec.QueryRowContext(ctx, sql, provideID, peerID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if peer_states exists")
	}

	return exists, nil
}
