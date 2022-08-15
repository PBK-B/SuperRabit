// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"
	"yayar/internal/data/ent/history"
	"yayar/internal/data/ent/version"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// HistoryCreate is the builder for creating a History entity.
type HistoryCreate struct {
	config
	mutation *HistoryMutation
	hooks    []Hook
}

// SetDevice sets the "device" field.
func (hc *HistoryCreate) SetDevice(s string) *HistoryCreate {
	hc.mutation.SetDevice(s)
	return hc
}

// SetIP sets the "ip" field.
func (hc *HistoryCreate) SetIP(s string) *HistoryCreate {
	hc.mutation.SetIP(s)
	return hc
}

// SetCreatedAt sets the "createdAt" field.
func (hc *HistoryCreate) SetCreatedAt(t time.Time) *HistoryCreate {
	hc.mutation.SetCreatedAt(t)
	return hc
}

// SetNillableCreatedAt sets the "createdAt" field if the given value is not nil.
func (hc *HistoryCreate) SetNillableCreatedAt(t *time.Time) *HistoryCreate {
	if t != nil {
		hc.SetCreatedAt(*t)
	}
	return hc
}

// AddVersionIDs adds the "version" edge to the Version entity by IDs.
func (hc *HistoryCreate) AddVersionIDs(ids ...int) *HistoryCreate {
	hc.mutation.AddVersionIDs(ids...)
	return hc
}

// AddVersion adds the "version" edges to the Version entity.
func (hc *HistoryCreate) AddVersion(v ...*Version) *HistoryCreate {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return hc.AddVersionIDs(ids...)
}

// Mutation returns the HistoryMutation object of the builder.
func (hc *HistoryCreate) Mutation() *HistoryMutation {
	return hc.mutation
}

// Save creates the History in the database.
func (hc *HistoryCreate) Save(ctx context.Context) (*History, error) {
	var (
		err  error
		node *History
	)
	hc.defaults()
	if len(hc.hooks) == 0 {
		if err = hc.check(); err != nil {
			return nil, err
		}
		node, err = hc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*HistoryMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = hc.check(); err != nil {
				return nil, err
			}
			hc.mutation = mutation
			if node, err = hc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(hc.hooks) - 1; i >= 0; i-- {
			if hc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = hc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, hc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (hc *HistoryCreate) SaveX(ctx context.Context) *History {
	v, err := hc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (hc *HistoryCreate) Exec(ctx context.Context) error {
	_, err := hc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hc *HistoryCreate) ExecX(ctx context.Context) {
	if err := hc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (hc *HistoryCreate) defaults() {
	if _, ok := hc.mutation.CreatedAt(); !ok {
		v := history.DefaultCreatedAt()
		hc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (hc *HistoryCreate) check() error {
	if _, ok := hc.mutation.Device(); !ok {
		return &ValidationError{Name: "device", err: errors.New(`ent: missing required field "History.device"`)}
	}
	if _, ok := hc.mutation.IP(); !ok {
		return &ValidationError{Name: "ip", err: errors.New(`ent: missing required field "History.ip"`)}
	}
	if _, ok := hc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "createdAt", err: errors.New(`ent: missing required field "History.createdAt"`)}
	}
	return nil
}

func (hc *HistoryCreate) sqlSave(ctx context.Context) (*History, error) {
	_node, _spec := hc.createSpec()
	if err := sqlgraph.CreateNode(ctx, hc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (hc *HistoryCreate) createSpec() (*History, *sqlgraph.CreateSpec) {
	var (
		_node = &History{config: hc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: history.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: history.FieldID,
			},
		}
	)
	if value, ok := hc.mutation.Device(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: history.FieldDevice,
		})
		_node.Device = value
	}
	if value, ok := hc.mutation.IP(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: history.FieldIP,
		})
		_node.IP = value
	}
	if value, ok := hc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: history.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if nodes := hc.mutation.VersionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   history.VersionTable,
			Columns: history.VersionPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: version.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// HistoryCreateBulk is the builder for creating many History entities in bulk.
type HistoryCreateBulk struct {
	config
	builders []*HistoryCreate
}

// Save creates the History entities in the database.
func (hcb *HistoryCreateBulk) Save(ctx context.Context) ([]*History, error) {
	specs := make([]*sqlgraph.CreateSpec, len(hcb.builders))
	nodes := make([]*History, len(hcb.builders))
	mutators := make([]Mutator, len(hcb.builders))
	for i := range hcb.builders {
		func(i int, root context.Context) {
			builder := hcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*HistoryMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, hcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, hcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, hcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (hcb *HistoryCreateBulk) SaveX(ctx context.Context) []*History {
	v, err := hcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (hcb *HistoryCreateBulk) Exec(ctx context.Context) error {
	_, err := hcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hcb *HistoryCreateBulk) ExecX(ctx context.Context) {
	if err := hcb.Exec(ctx); err != nil {
		panic(err)
	}
}
