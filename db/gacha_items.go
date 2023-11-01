// Code generated by SQLBoiler 4.15.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package db

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

// GachaItem is an object representing the database table.
type GachaItem struct {
	ID      int64 `boil:"id" json:"id" toml:"id" yaml:"id"`
	GachaID int   `boil:"gacha_id" json:"gacha_id" toml:"gacha_id" yaml:"gacha_id"`
	ItemID  int64 `boil:"item_id" json:"item_id" toml:"item_id" yaml:"item_id"`
	Weight  int   `boil:"weight" json:"weight" toml:"weight" yaml:"weight"`

	R *gachaItemR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L gachaItemL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var GachaItemColumns = struct {
	ID      string
	GachaID string
	ItemID  string
	Weight  string
}{
	ID:      "id",
	GachaID: "gacha_id",
	ItemID:  "item_id",
	Weight:  "weight",
}

var GachaItemTableColumns = struct {
	ID      string
	GachaID string
	ItemID  string
	Weight  string
}{
	ID:      "gacha_items.id",
	GachaID: "gacha_items.gacha_id",
	ItemID:  "gacha_items.item_id",
	Weight:  "gacha_items.weight",
}

// Generated where

type whereHelperint64 struct{ field string }

func (w whereHelperint64) EQ(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint64) NEQ(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint64) LT(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint64) LTE(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint64) GT(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint64) GTE(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperint64) IN(slice []int64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperint64) NIN(slice []int64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelperint struct{ field string }

func (w whereHelperint) EQ(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint) NEQ(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint) LT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint) LTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint) GT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint) GTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperint) IN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperint) NIN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

var GachaItemWhere = struct {
	ID      whereHelperint64
	GachaID whereHelperint
	ItemID  whereHelperint64
	Weight  whereHelperint
}{
	ID:      whereHelperint64{field: "`gacha_items`.`id`"},
	GachaID: whereHelperint{field: "`gacha_items`.`gacha_id`"},
	ItemID:  whereHelperint64{field: "`gacha_items`.`item_id`"},
	Weight:  whereHelperint{field: "`gacha_items`.`weight`"},
}

// GachaItemRels is where relationship names are stored.
var GachaItemRels = struct {
	Gacha string
	Item  string
}{
	Gacha: "Gacha",
	Item:  "Item",
}

// gachaItemR is where relationships are stored.
type gachaItemR struct {
	Gacha *Gacha `boil:"Gacha" json:"Gacha" toml:"Gacha" yaml:"Gacha"`
	Item  *Item  `boil:"Item" json:"Item" toml:"Item" yaml:"Item"`
}

// NewStruct creates a new relationship struct
func (*gachaItemR) NewStruct() *gachaItemR {
	return &gachaItemR{}
}

func (r *gachaItemR) GetGacha() *Gacha {
	if r == nil {
		return nil
	}
	return r.Gacha
}

func (r *gachaItemR) GetItem() *Item {
	if r == nil {
		return nil
	}
	return r.Item
}

// gachaItemL is where Load methods for each relationship are stored.
type gachaItemL struct{}

var (
	gachaItemAllColumns            = []string{"id", "gacha_id", "item_id", "weight"}
	gachaItemColumnsWithoutDefault = []string{"gacha_id", "item_id", "weight"}
	gachaItemColumnsWithDefault    = []string{"id"}
	gachaItemPrimaryKeyColumns     = []string{"id"}
	gachaItemGeneratedColumns      = []string{}
)

type (
	// GachaItemSlice is an alias for a slice of pointers to GachaItem.
	// This should almost always be used instead of []GachaItem.
	GachaItemSlice []*GachaItem
	// GachaItemHook is the signature for custom GachaItem hook methods
	GachaItemHook func(context.Context, boil.ContextExecutor, *GachaItem) error

	gachaItemQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	gachaItemType                 = reflect.TypeOf(&GachaItem{})
	gachaItemMapping              = queries.MakeStructMapping(gachaItemType)
	gachaItemPrimaryKeyMapping, _ = queries.BindMapping(gachaItemType, gachaItemMapping, gachaItemPrimaryKeyColumns)
	gachaItemInsertCacheMut       sync.RWMutex
	gachaItemInsertCache          = make(map[string]insertCache)
	gachaItemUpdateCacheMut       sync.RWMutex
	gachaItemUpdateCache          = make(map[string]updateCache)
	gachaItemUpsertCacheMut       sync.RWMutex
	gachaItemUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var gachaItemAfterSelectHooks []GachaItemHook

var gachaItemBeforeInsertHooks []GachaItemHook
var gachaItemAfterInsertHooks []GachaItemHook

var gachaItemBeforeUpdateHooks []GachaItemHook
var gachaItemAfterUpdateHooks []GachaItemHook

var gachaItemBeforeDeleteHooks []GachaItemHook
var gachaItemAfterDeleteHooks []GachaItemHook

var gachaItemBeforeUpsertHooks []GachaItemHook
var gachaItemAfterUpsertHooks []GachaItemHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *GachaItem) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range gachaItemAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *GachaItem) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range gachaItemBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *GachaItem) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range gachaItemAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *GachaItem) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range gachaItemBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *GachaItem) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range gachaItemAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *GachaItem) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range gachaItemBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *GachaItem) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range gachaItemAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *GachaItem) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range gachaItemBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *GachaItem) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range gachaItemAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddGachaItemHook registers your hook function for all future operations.
func AddGachaItemHook(hookPoint boil.HookPoint, gachaItemHook GachaItemHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		gachaItemAfterSelectHooks = append(gachaItemAfterSelectHooks, gachaItemHook)
	case boil.BeforeInsertHook:
		gachaItemBeforeInsertHooks = append(gachaItemBeforeInsertHooks, gachaItemHook)
	case boil.AfterInsertHook:
		gachaItemAfterInsertHooks = append(gachaItemAfterInsertHooks, gachaItemHook)
	case boil.BeforeUpdateHook:
		gachaItemBeforeUpdateHooks = append(gachaItemBeforeUpdateHooks, gachaItemHook)
	case boil.AfterUpdateHook:
		gachaItemAfterUpdateHooks = append(gachaItemAfterUpdateHooks, gachaItemHook)
	case boil.BeforeDeleteHook:
		gachaItemBeforeDeleteHooks = append(gachaItemBeforeDeleteHooks, gachaItemHook)
	case boil.AfterDeleteHook:
		gachaItemAfterDeleteHooks = append(gachaItemAfterDeleteHooks, gachaItemHook)
	case boil.BeforeUpsertHook:
		gachaItemBeforeUpsertHooks = append(gachaItemBeforeUpsertHooks, gachaItemHook)
	case boil.AfterUpsertHook:
		gachaItemAfterUpsertHooks = append(gachaItemAfterUpsertHooks, gachaItemHook)
	}
}

// One returns a single gachaItem record from the query.
func (q gachaItemQuery) One(ctx context.Context, exec boil.ContextExecutor) (*GachaItem, error) {
	o := &GachaItem{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "db: failed to execute a one query for gacha_items")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all GachaItem records from the query.
func (q gachaItemQuery) All(ctx context.Context, exec boil.ContextExecutor) (GachaItemSlice, error) {
	var o []*GachaItem

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "db: failed to assign all query results to GachaItem slice")
	}

	if len(gachaItemAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all GachaItem records in the query.
func (q gachaItemQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "db: failed to count gacha_items rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q gachaItemQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "db: failed to check if gacha_items exists")
	}

	return count > 0, nil
}

// Gacha pointed to by the foreign key.
func (o *GachaItem) Gacha(mods ...qm.QueryMod) gachaQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`id` = ?", o.GachaID),
	}

	queryMods = append(queryMods, mods...)

	return Gachas(queryMods...)
}

// Item pointed to by the foreign key.
func (o *GachaItem) Item(mods ...qm.QueryMod) itemQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`id` = ?", o.ItemID),
	}

	queryMods = append(queryMods, mods...)

	return Items(queryMods...)
}

// LoadGacha allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (gachaItemL) LoadGacha(ctx context.Context, e boil.ContextExecutor, singular bool, maybeGachaItem interface{}, mods queries.Applicator) error {
	var slice []*GachaItem
	var object *GachaItem

	if singular {
		var ok bool
		object, ok = maybeGachaItem.(*GachaItem)
		if !ok {
			object = new(GachaItem)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeGachaItem)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeGachaItem))
			}
		}
	} else {
		s, ok := maybeGachaItem.(*[]*GachaItem)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeGachaItem)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeGachaItem))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &gachaItemR{}
		}
		args = append(args, object.GachaID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &gachaItemR{}
			}

			for _, a := range args {
				if a == obj.GachaID {
					continue Outer
				}
			}

			args = append(args, obj.GachaID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`gachas`),
		qm.WhereIn(`gachas.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Gacha")
	}

	var resultSlice []*Gacha
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Gacha")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for gachas")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for gachas")
	}

	if len(gachaAfterSelectHooks) != 0 {
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
		object.R.Gacha = foreign
		if foreign.R == nil {
			foreign.R = &gachaR{}
		}
		foreign.R.GachaGachaItems = append(foreign.R.GachaGachaItems, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.GachaID == foreign.ID {
				local.R.Gacha = foreign
				if foreign.R == nil {
					foreign.R = &gachaR{}
				}
				foreign.R.GachaGachaItems = append(foreign.R.GachaGachaItems, local)
				break
			}
		}
	}

	return nil
}

// LoadItem allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (gachaItemL) LoadItem(ctx context.Context, e boil.ContextExecutor, singular bool, maybeGachaItem interface{}, mods queries.Applicator) error {
	var slice []*GachaItem
	var object *GachaItem

	if singular {
		var ok bool
		object, ok = maybeGachaItem.(*GachaItem)
		if !ok {
			object = new(GachaItem)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeGachaItem)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeGachaItem))
			}
		}
	} else {
		s, ok := maybeGachaItem.(*[]*GachaItem)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeGachaItem)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeGachaItem))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &gachaItemR{}
		}
		args = append(args, object.ItemID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &gachaItemR{}
			}

			for _, a := range args {
				if a == obj.ItemID {
					continue Outer
				}
			}

			args = append(args, obj.ItemID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`items`),
		qm.WhereIn(`items.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Item")
	}

	var resultSlice []*Item
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Item")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for items")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for items")
	}

	if len(itemAfterSelectHooks) != 0 {
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
		object.R.Item = foreign
		if foreign.R == nil {
			foreign.R = &itemR{}
		}
		foreign.R.GachaItems = append(foreign.R.GachaItems, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.ItemID == foreign.ID {
				local.R.Item = foreign
				if foreign.R == nil {
					foreign.R = &itemR{}
				}
				foreign.R.GachaItems = append(foreign.R.GachaItems, local)
				break
			}
		}
	}

	return nil
}

// SetGacha of the gachaItem to the related item.
// Sets o.R.Gacha to related.
// Adds o to related.R.GachaGachaItems.
func (o *GachaItem) SetGacha(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Gacha) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `gacha_items` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"gacha_id"}),
		strmangle.WhereClause("`", "`", 0, gachaItemPrimaryKeyColumns),
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

	o.GachaID = related.ID
	if o.R == nil {
		o.R = &gachaItemR{
			Gacha: related,
		}
	} else {
		o.R.Gacha = related
	}

	if related.R == nil {
		related.R = &gachaR{
			GachaGachaItems: GachaItemSlice{o},
		}
	} else {
		related.R.GachaGachaItems = append(related.R.GachaGachaItems, o)
	}

	return nil
}

// SetItem of the gachaItem to the related item.
// Sets o.R.Item to related.
// Adds o to related.R.GachaItems.
func (o *GachaItem) SetItem(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Item) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `gacha_items` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"item_id"}),
		strmangle.WhereClause("`", "`", 0, gachaItemPrimaryKeyColumns),
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

	o.ItemID = related.ID
	if o.R == nil {
		o.R = &gachaItemR{
			Item: related,
		}
	} else {
		o.R.Item = related
	}

	if related.R == nil {
		related.R = &itemR{
			GachaItems: GachaItemSlice{o},
		}
	} else {
		related.R.GachaItems = append(related.R.GachaItems, o)
	}

	return nil
}

// GachaItems retrieves all the records using an executor.
func GachaItems(mods ...qm.QueryMod) gachaItemQuery {
	mods = append(mods, qm.From("`gacha_items`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`gacha_items`.*"})
	}

	return gachaItemQuery{q}
}

// FindGachaItem retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindGachaItem(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*GachaItem, error) {
	gachaItemObj := &GachaItem{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `gacha_items` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, gachaItemObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "db: unable to select from gacha_items")
	}

	if err = gachaItemObj.doAfterSelectHooks(ctx, exec); err != nil {
		return gachaItemObj, err
	}

	return gachaItemObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *GachaItem) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("db: no gacha_items provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(gachaItemColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	gachaItemInsertCacheMut.RLock()
	cache, cached := gachaItemInsertCache[key]
	gachaItemInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			gachaItemAllColumns,
			gachaItemColumnsWithDefault,
			gachaItemColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(gachaItemType, gachaItemMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(gachaItemType, gachaItemMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `gacha_items` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `gacha_items` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `gacha_items` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, gachaItemPrimaryKeyColumns))
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
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "db: unable to insert into gacha_items")
	}

	var lastID int64
	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = int64(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == gachaItemMapping["id"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "db: unable to populate default values for gacha_items")
	}

CacheNoHooks:
	if !cached {
		gachaItemInsertCacheMut.Lock()
		gachaItemInsertCache[key] = cache
		gachaItemInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the GachaItem.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *GachaItem) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	gachaItemUpdateCacheMut.RLock()
	cache, cached := gachaItemUpdateCache[key]
	gachaItemUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			gachaItemAllColumns,
			gachaItemPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("db: unable to update gacha_items, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `gacha_items` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, gachaItemPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(gachaItemType, gachaItemMapping, append(wl, gachaItemPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "db: unable to update gacha_items row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "db: failed to get rows affected by update for gacha_items")
	}

	if !cached {
		gachaItemUpdateCacheMut.Lock()
		gachaItemUpdateCache[key] = cache
		gachaItemUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q gachaItemQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "db: unable to update all for gacha_items")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "db: unable to retrieve rows affected for gacha_items")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o GachaItemSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("db: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), gachaItemPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `gacha_items` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, gachaItemPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "db: unable to update all in gachaItem slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "db: unable to retrieve rows affected all in update all gachaItem")
	}
	return rowsAff, nil
}

var mySQLGachaItemUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *GachaItem) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("db: no gacha_items provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(gachaItemColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLGachaItemUniqueColumns, o)

	if len(nzUniques) == 0 {
		return errors.New("cannot upsert with a table that cannot conflict on a unique column")
	}

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
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
	buf.WriteByte('.')
	for _, c := range nzUniques {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	gachaItemUpsertCacheMut.RLock()
	cache, cached := gachaItemUpsertCache[key]
	gachaItemUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			gachaItemAllColumns,
			gachaItemColumnsWithDefault,
			gachaItemColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			gachaItemAllColumns,
			gachaItemPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("db: unable to upsert gacha_items, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`gacha_items`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `gacha_items` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(gachaItemType, gachaItemMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(gachaItemType, gachaItemMapping, ret)
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
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "db: unable to upsert for gacha_items")
	}

	var lastID int64
	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = int64(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == gachaItemMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(gachaItemType, gachaItemMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "db: unable to retrieve unique values for gacha_items")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "db: unable to populate default values for gacha_items")
	}

CacheNoHooks:
	if !cached {
		gachaItemUpsertCacheMut.Lock()
		gachaItemUpsertCache[key] = cache
		gachaItemUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single GachaItem record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *GachaItem) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("db: no GachaItem provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), gachaItemPrimaryKeyMapping)
	sql := "DELETE FROM `gacha_items` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "db: unable to delete from gacha_items")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "db: failed to get rows affected by delete for gacha_items")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q gachaItemQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("db: no gachaItemQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "db: unable to delete all from gacha_items")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "db: failed to get rows affected by deleteall for gacha_items")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o GachaItemSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(gachaItemBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), gachaItemPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `gacha_items` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, gachaItemPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "db: unable to delete all from gachaItem slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "db: failed to get rows affected by deleteall for gacha_items")
	}

	if len(gachaItemAfterDeleteHooks) != 0 {
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
func (o *GachaItem) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindGachaItem(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *GachaItemSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := GachaItemSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), gachaItemPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `gacha_items`.* FROM `gacha_items` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, gachaItemPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "db: unable to reload all in GachaItemSlice")
	}

	*o = slice

	return nil
}

// GachaItemExists checks if the GachaItem row exists.
func GachaItemExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `gacha_items` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "db: unable to check if gacha_items exists")
	}

	return exists, nil
}

// Exists checks if the GachaItem row exists.
func (o *GachaItem) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return GachaItemExists(ctx, exec, o.ID)
}
