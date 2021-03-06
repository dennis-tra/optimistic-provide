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
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Host is an object representing the database table.
type Host struct {
	ID         int       `boil:"id" json:"id" toml:"id" yaml:"id"`
	PeerID     int       `boil:"peer_id" json:"peer_id" toml:"peer_id" yaml:"peer_id"`
	Name       string    `boil:"name" json:"name" toml:"name" yaml:"name"`
	PrivateKey []byte    `boil:"private_key" json:"private_key" toml:"private_key" yaml:"private_key"`
	CreatedAt  time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt  time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`
	ArchivedAt null.Time `boil:"archived_at" json:"archived_at,omitempty" toml:"archived_at" yaml:"archived_at,omitempty"`
	Network    string    `boil:"network" json:"network" toml:"network" yaml:"network"`

	R *hostR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L hostL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var HostColumns = struct {
	ID         string
	PeerID     string
	Name       string
	PrivateKey string
	CreatedAt  string
	UpdatedAt  string
	ArchivedAt string
	Network    string
}{
	ID:         "id",
	PeerID:     "peer_id",
	Name:       "name",
	PrivateKey: "private_key",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
	ArchivedAt: "archived_at",
	Network:    "network",
}

var HostTableColumns = struct {
	ID         string
	PeerID     string
	Name       string
	PrivateKey string
	CreatedAt  string
	UpdatedAt  string
	ArchivedAt string
	Network    string
}{
	ID:         "hosts.id",
	PeerID:     "hosts.peer_id",
	Name:       "hosts.name",
	PrivateKey: "hosts.private_key",
	CreatedAt:  "hosts.created_at",
	UpdatedAt:  "hosts.updated_at",
	ArchivedAt: "hosts.archived_at",
	Network:    "hosts.network",
}

// Generated where

type whereHelpernull_Time struct{ field string }

func (w whereHelpernull_Time) EQ(x null.Time) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Time) NEQ(x null.Time) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Time) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Time) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }
func (w whereHelpernull_Time) LT(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Time) LTE(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Time) GT(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Time) GTE(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var HostWhere = struct {
	ID         whereHelperint
	PeerID     whereHelperint
	Name       whereHelperstring
	PrivateKey whereHelper__byte
	CreatedAt  whereHelpertime_Time
	UpdatedAt  whereHelpertime_Time
	ArchivedAt whereHelpernull_Time
	Network    whereHelperstring
}{
	ID:         whereHelperint{field: "\"hosts\".\"id\""},
	PeerID:     whereHelperint{field: "\"hosts\".\"peer_id\""},
	Name:       whereHelperstring{field: "\"hosts\".\"name\""},
	PrivateKey: whereHelper__byte{field: "\"hosts\".\"private_key\""},
	CreatedAt:  whereHelpertime_Time{field: "\"hosts\".\"created_at\""},
	UpdatedAt:  whereHelpertime_Time{field: "\"hosts\".\"updated_at\""},
	ArchivedAt: whereHelpernull_Time{field: "\"hosts\".\"archived_at\""},
	Network:    whereHelperstring{field: "\"hosts\".\"network\""},
}

// HostRels is where relationship names are stored.
var HostRels = struct {
	Peer                 string
	NetworkSizeEstimates string
}{
	Peer:                 "Peer",
	NetworkSizeEstimates: "NetworkSizeEstimates",
}

// hostR is where relationships are stored.
type hostR struct {
	Peer                 *Peer                    `boil:"Peer" json:"Peer" toml:"Peer" yaml:"Peer"`
	NetworkSizeEstimates NetworkSizeEstimateSlice `boil:"NetworkSizeEstimates" json:"NetworkSizeEstimates" toml:"NetworkSizeEstimates" yaml:"NetworkSizeEstimates"`
}

// NewStruct creates a new relationship struct
func (*hostR) NewStruct() *hostR {
	return &hostR{}
}

// hostL is where Load methods for each relationship are stored.
type hostL struct{}

var (
	hostAllColumns            = []string{"id", "peer_id", "name", "private_key", "created_at", "updated_at", "archived_at", "network"}
	hostColumnsWithoutDefault = []string{"peer_id", "name", "private_key", "created_at", "updated_at", "archived_at"}
	hostColumnsWithDefault    = []string{"id", "network"}
	hostPrimaryKeyColumns     = []string{"id"}
)

type (
	// HostSlice is an alias for a slice of pointers to Host.
	// This should almost always be used instead of []Host.
	HostSlice []*Host
	// HostHook is the signature for custom Host hook methods
	HostHook func(context.Context, boil.ContextExecutor, *Host) error

	hostQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	hostType                 = reflect.TypeOf(&Host{})
	hostMapping              = queries.MakeStructMapping(hostType)
	hostPrimaryKeyMapping, _ = queries.BindMapping(hostType, hostMapping, hostPrimaryKeyColumns)
	hostInsertCacheMut       sync.RWMutex
	hostInsertCache          = make(map[string]insertCache)
	hostUpdateCacheMut       sync.RWMutex
	hostUpdateCache          = make(map[string]updateCache)
	hostUpsertCacheMut       sync.RWMutex
	hostUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var hostBeforeInsertHooks []HostHook
var hostBeforeUpdateHooks []HostHook
var hostBeforeDeleteHooks []HostHook
var hostBeforeUpsertHooks []HostHook

var hostAfterInsertHooks []HostHook
var hostAfterSelectHooks []HostHook
var hostAfterUpdateHooks []HostHook
var hostAfterDeleteHooks []HostHook
var hostAfterUpsertHooks []HostHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Host) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range hostBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Host) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range hostBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Host) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range hostBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Host) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range hostBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Host) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range hostAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Host) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range hostAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Host) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range hostAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Host) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range hostAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Host) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range hostAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddHostHook registers your hook function for all future operations.
func AddHostHook(hookPoint boil.HookPoint, hostHook HostHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		hostBeforeInsertHooks = append(hostBeforeInsertHooks, hostHook)
	case boil.BeforeUpdateHook:
		hostBeforeUpdateHooks = append(hostBeforeUpdateHooks, hostHook)
	case boil.BeforeDeleteHook:
		hostBeforeDeleteHooks = append(hostBeforeDeleteHooks, hostHook)
	case boil.BeforeUpsertHook:
		hostBeforeUpsertHooks = append(hostBeforeUpsertHooks, hostHook)
	case boil.AfterInsertHook:
		hostAfterInsertHooks = append(hostAfterInsertHooks, hostHook)
	case boil.AfterSelectHook:
		hostAfterSelectHooks = append(hostAfterSelectHooks, hostHook)
	case boil.AfterUpdateHook:
		hostAfterUpdateHooks = append(hostAfterUpdateHooks, hostHook)
	case boil.AfterDeleteHook:
		hostAfterDeleteHooks = append(hostAfterDeleteHooks, hostHook)
	case boil.AfterUpsertHook:
		hostAfterUpsertHooks = append(hostAfterUpsertHooks, hostHook)
	}
}

// One returns a single host record from the query.
func (q hostQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Host, error) {
	o := &Host{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for hosts")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Host records from the query.
func (q hostQuery) All(ctx context.Context, exec boil.ContextExecutor) (HostSlice, error) {
	var o []*Host

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Host slice")
	}

	if len(hostAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Host records in the query.
func (q hostQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count hosts rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q hostQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if hosts exists")
	}

	return count > 0, nil
}

// Peer pointed to by the foreign key.
func (o *Host) Peer(mods ...qm.QueryMod) peerQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.PeerID),
	}

	queryMods = append(queryMods, mods...)

	query := Peers(queryMods...)
	queries.SetFrom(query.Query, "\"peers\"")

	return query
}

// NetworkSizeEstimates retrieves all the network_size_estimate's NetworkSizeEstimates with an executor.
func (o *Host) NetworkSizeEstimates(mods ...qm.QueryMod) networkSizeEstimateQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"network_size_estimates\".\"host_id\"=?", o.ID),
	)

	query := NetworkSizeEstimates(queryMods...)
	queries.SetFrom(query.Query, "\"network_size_estimates\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"network_size_estimates\".*"})
	}

	return query
}

// LoadPeer allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (hostL) LoadPeer(ctx context.Context, e boil.ContextExecutor, singular bool, maybeHost interface{}, mods queries.Applicator) error {
	var slice []*Host
	var object *Host

	if singular {
		object = maybeHost.(*Host)
	} else {
		slice = *maybeHost.(*[]*Host)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &hostR{}
		}
		args = append(args, object.PeerID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &hostR{}
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

	if len(hostAfterSelectHooks) != 0 {
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
		foreign.R.Hosts = append(foreign.R.Hosts, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.PeerID == foreign.ID {
				local.R.Peer = foreign
				if foreign.R == nil {
					foreign.R = &peerR{}
				}
				foreign.R.Hosts = append(foreign.R.Hosts, local)
				break
			}
		}
	}

	return nil
}

// LoadNetworkSizeEstimates allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (hostL) LoadNetworkSizeEstimates(ctx context.Context, e boil.ContextExecutor, singular bool, maybeHost interface{}, mods queries.Applicator) error {
	var slice []*Host
	var object *Host

	if singular {
		object = maybeHost.(*Host)
	} else {
		slice = *maybeHost.(*[]*Host)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &hostR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &hostR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`network_size_estimates`),
		qm.WhereIn(`network_size_estimates.host_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load network_size_estimates")
	}

	var resultSlice []*NetworkSizeEstimate
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice network_size_estimates")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on network_size_estimates")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for network_size_estimates")
	}

	if len(networkSizeEstimateAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.NetworkSizeEstimates = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &networkSizeEstimateR{}
			}
			foreign.R.Host = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.HostID {
				local.R.NetworkSizeEstimates = append(local.R.NetworkSizeEstimates, foreign)
				if foreign.R == nil {
					foreign.R = &networkSizeEstimateR{}
				}
				foreign.R.Host = local
				break
			}
		}
	}

	return nil
}

// SetPeer of the host to the related item.
// Sets o.R.Peer to related.
// Adds o to related.R.Hosts.
func (o *Host) SetPeer(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Peer) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"hosts\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"peer_id"}),
		strmangle.WhereClause("\"", "\"", 2, hostPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

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
		o.R = &hostR{
			Peer: related,
		}
	} else {
		o.R.Peer = related
	}

	if related.R == nil {
		related.R = &peerR{
			Hosts: HostSlice{o},
		}
	} else {
		related.R.Hosts = append(related.R.Hosts, o)
	}

	return nil
}

// AddNetworkSizeEstimates adds the given related objects to the existing relationships
// of the host, optionally inserting them as new records.
// Appends related to o.R.NetworkSizeEstimates.
// Sets related.R.Host appropriately.
func (o *Host) AddNetworkSizeEstimates(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*NetworkSizeEstimate) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.HostID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"network_size_estimates\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"host_id"}),
				strmangle.WhereClause("\"", "\"", 2, networkSizeEstimatePrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.HostID = o.ID
		}
	}

	if o.R == nil {
		o.R = &hostR{
			NetworkSizeEstimates: related,
		}
	} else {
		o.R.NetworkSizeEstimates = append(o.R.NetworkSizeEstimates, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &networkSizeEstimateR{
				Host: o,
			}
		} else {
			rel.R.Host = o
		}
	}
	return nil
}

// Hosts retrieves all the records using an executor.
func Hosts(mods ...qm.QueryMod) hostQuery {
	mods = append(mods, qm.From("\"hosts\""))
	return hostQuery{NewQuery(mods...)}
}

// FindHost retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindHost(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*Host, error) {
	hostObj := &Host{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"hosts\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, hostObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from hosts")
	}

	if err = hostObj.doAfterSelectHooks(ctx, exec); err != nil {
		return hostObj, err
	}

	return hostObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Host) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no hosts provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		if o.UpdatedAt.IsZero() {
			o.UpdatedAt = currTime
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(hostColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	hostInsertCacheMut.RLock()
	cache, cached := hostInsertCache[key]
	hostInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			hostAllColumns,
			hostColumnsWithDefault,
			hostColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(hostType, hostMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(hostType, hostMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"hosts\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"hosts\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into hosts")
	}

	if !cached {
		hostInsertCacheMut.Lock()
		hostInsertCache[key] = cache
		hostInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Host.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Host) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	hostUpdateCacheMut.RLock()
	cache, cached := hostUpdateCache[key]
	hostUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			hostAllColumns,
			hostPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update hosts, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"hosts\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, hostPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(hostType, hostMapping, append(wl, hostPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update hosts row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for hosts")
	}

	if !cached {
		hostUpdateCacheMut.Lock()
		hostUpdateCache[key] = cache
		hostUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q hostQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for hosts")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for hosts")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o HostSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), hostPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"hosts\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, hostPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in host slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all host")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Host) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no hosts provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		o.UpdatedAt = currTime
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(hostColumnsWithDefault, o)

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

	hostUpsertCacheMut.RLock()
	cache, cached := hostUpsertCache[key]
	hostUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			hostAllColumns,
			hostColumnsWithDefault,
			hostColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			hostAllColumns,
			hostPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert hosts, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(hostPrimaryKeyColumns))
			copy(conflict, hostPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"hosts\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(hostType, hostMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(hostType, hostMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert hosts")
	}

	if !cached {
		hostUpsertCacheMut.Lock()
		hostUpsertCache[key] = cache
		hostUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Host record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Host) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Host provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), hostPrimaryKeyMapping)
	sql := "DELETE FROM \"hosts\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from hosts")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for hosts")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q hostQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no hostQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from hosts")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for hosts")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o HostSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(hostBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), hostPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"hosts\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, hostPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from host slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for hosts")
	}

	if len(hostAfterDeleteHooks) != 0 {
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
func (o *Host) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindHost(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *HostSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := HostSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), hostPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"hosts\".* FROM \"hosts\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, hostPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in HostSlice")
	}

	*o = slice

	return nil
}

// HostExists checks if the Host row exists.
func HostExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"hosts\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if hosts exists")
	}

	return exists, nil
}
