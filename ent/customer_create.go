// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"shopular/ent/customer"
	"shopular/ent/user"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CustomerCreate is the builder for creating a Customer entity.
type CustomerCreate struct {
	config
	mutation *CustomerMutation
	hooks    []Hook
}

// SetFirstName sets the "first_name" field.
func (cc *CustomerCreate) SetFirstName(s string) *CustomerCreate {
	cc.mutation.SetFirstName(s)
	return cc
}

// SetLastName sets the "last_name" field.
func (cc *CustomerCreate) SetLastName(s string) *CustomerCreate {
	cc.mutation.SetLastName(s)
	return cc
}

// SetAddress sets the "address" field.
func (cc *CustomerCreate) SetAddress(s string) *CustomerCreate {
	cc.mutation.SetAddress(s)
	return cc
}

// SetNillableAddress sets the "address" field if the given value is not nil.
func (cc *CustomerCreate) SetNillableAddress(s *string) *CustomerCreate {
	if s != nil {
		cc.SetAddress(*s)
	}
	return cc
}

// SetPhone sets the "phone" field.
func (cc *CustomerCreate) SetPhone(s string) *CustomerCreate {
	cc.mutation.SetPhone(s)
	return cc
}

// SetNillablePhone sets the "phone" field if the given value is not nil.
func (cc *CustomerCreate) SetNillablePhone(s *string) *CustomerCreate {
	if s != nil {
		cc.SetPhone(*s)
	}
	return cc
}

// SetEmail sets the "email" field.
func (cc *CustomerCreate) SetEmail(s string) *CustomerCreate {
	cc.mutation.SetEmail(s)
	return cc
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (cc *CustomerCreate) SetNillableEmail(s *string) *CustomerCreate {
	if s != nil {
		cc.SetEmail(*s)
	}
	return cc
}

// SetOrderCount sets the "order_count" field.
func (cc *CustomerCreate) SetOrderCount(i int) *CustomerCreate {
	cc.mutation.SetOrderCount(i)
	return cc
}

// SetNillableOrderCount sets the "order_count" field if the given value is not nil.
func (cc *CustomerCreate) SetNillableOrderCount(i *int) *CustomerCreate {
	if i != nil {
		cc.SetOrderCount(*i)
	}
	return cc
}

// SetUserID sets the "user" edge to the User entity by ID.
func (cc *CustomerCreate) SetUserID(id int) *CustomerCreate {
	cc.mutation.SetUserID(id)
	return cc
}

// SetUser sets the "user" edge to the User entity.
func (cc *CustomerCreate) SetUser(u *User) *CustomerCreate {
	return cc.SetUserID(u.ID)
}

// Mutation returns the CustomerMutation object of the builder.
func (cc *CustomerCreate) Mutation() *CustomerMutation {
	return cc.mutation
}

// Save creates the Customer in the database.
func (cc *CustomerCreate) Save(ctx context.Context) (*Customer, error) {
	var (
		err  error
		node *Customer
	)
	cc.defaults()
	if len(cc.hooks) == 0 {
		if err = cc.check(); err != nil {
			return nil, err
		}
		node, err = cc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CustomerMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cc.check(); err != nil {
				return nil, err
			}
			cc.mutation = mutation
			if node, err = cc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(cc.hooks) - 1; i >= 0; i-- {
			if cc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, cc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Customer)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from CustomerMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CustomerCreate) SaveX(ctx context.Context) *Customer {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *CustomerCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *CustomerCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cc *CustomerCreate) defaults() {
	if _, ok := cc.mutation.OrderCount(); !ok {
		v := customer.DefaultOrderCount
		cc.mutation.SetOrderCount(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *CustomerCreate) check() error {
	if _, ok := cc.mutation.FirstName(); !ok {
		return &ValidationError{Name: "first_name", err: errors.New(`ent: missing required field "Customer.first_name"`)}
	}
	if _, ok := cc.mutation.LastName(); !ok {
		return &ValidationError{Name: "last_name", err: errors.New(`ent: missing required field "Customer.last_name"`)}
	}
	if _, ok := cc.mutation.OrderCount(); !ok {
		return &ValidationError{Name: "order_count", err: errors.New(`ent: missing required field "Customer.order_count"`)}
	}
	if _, ok := cc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "Customer.user"`)}
	}
	return nil
}

func (cc *CustomerCreate) sqlSave(ctx context.Context) (*Customer, error) {
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (cc *CustomerCreate) createSpec() (*Customer, *sqlgraph.CreateSpec) {
	var (
		_node = &Customer{config: cc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: customer.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: customer.FieldID,
			},
		}
	)
	if value, ok := cc.mutation.FirstName(); ok {
		_spec.SetField(customer.FieldFirstName, field.TypeString, value)
		_node.FirstName = value
	}
	if value, ok := cc.mutation.LastName(); ok {
		_spec.SetField(customer.FieldLastName, field.TypeString, value)
		_node.LastName = value
	}
	if value, ok := cc.mutation.Address(); ok {
		_spec.SetField(customer.FieldAddress, field.TypeString, value)
		_node.Address = value
	}
	if value, ok := cc.mutation.Phone(); ok {
		_spec.SetField(customer.FieldPhone, field.TypeString, value)
		_node.Phone = value
	}
	if value, ok := cc.mutation.Email(); ok {
		_spec.SetField(customer.FieldEmail, field.TypeString, value)
		_node.Email = value
	}
	if value, ok := cc.mutation.OrderCount(); ok {
		_spec.SetField(customer.FieldOrderCount, field.TypeInt, value)
		_node.OrderCount = value
	}
	if nodes := cc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   customer.UserTable,
			Columns: []string{customer.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_customer = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// CustomerCreateBulk is the builder for creating many Customer entities in bulk.
type CustomerCreateBulk struct {
	config
	builders []*CustomerCreate
}

// Save creates the Customer entities in the database.
func (ccb *CustomerCreateBulk) Save(ctx context.Context) ([]*Customer, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Customer, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CustomerMutation)
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
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *CustomerCreateBulk) SaveX(ctx context.Context) []*Customer {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *CustomerCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *CustomerCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}
