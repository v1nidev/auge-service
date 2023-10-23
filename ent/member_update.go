// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/v1nidev/auge-service/ent/emergencycontact"
	"github.com/v1nidev/auge-service/ent/member"
	"github.com/v1nidev/auge-service/ent/predicate"
)

// MemberUpdate is the builder for updating Member entities.
type MemberUpdate struct {
	config
	hooks    []Hook
	mutation *MemberMutation
}

// Where appends a list predicates to the MemberUpdate builder.
func (mu *MemberUpdate) Where(ps ...predicate.Member) *MemberUpdate {
	mu.mutation.Where(ps...)
	return mu
}

// SetName sets the "name" field.
func (mu *MemberUpdate) SetName(s string) *MemberUpdate {
	mu.mutation.SetName(s)
	return mu
}

// SetSex sets the "sex" field.
func (mu *MemberUpdate) SetSex(s string) *MemberUpdate {
	mu.mutation.SetSex(s)
	return mu
}

// SetEmail sets the "email" field.
func (mu *MemberUpdate) SetEmail(s string) *MemberUpdate {
	mu.mutation.SetEmail(s)
	return mu
}

// SetEmergencycontactID sets the "emergencycontact" edge to the EmergencyContact entity by ID.
func (mu *MemberUpdate) SetEmergencycontactID(id int) *MemberUpdate {
	mu.mutation.SetEmergencycontactID(id)
	return mu
}

// SetNillableEmergencycontactID sets the "emergencycontact" edge to the EmergencyContact entity by ID if the given value is not nil.
func (mu *MemberUpdate) SetNillableEmergencycontactID(id *int) *MemberUpdate {
	if id != nil {
		mu = mu.SetEmergencycontactID(*id)
	}
	return mu
}

// SetEmergencycontact sets the "emergencycontact" edge to the EmergencyContact entity.
func (mu *MemberUpdate) SetEmergencycontact(e *EmergencyContact) *MemberUpdate {
	return mu.SetEmergencycontactID(e.ID)
}

// Mutation returns the MemberMutation object of the builder.
func (mu *MemberUpdate) Mutation() *MemberMutation {
	return mu.mutation
}

// ClearEmergencycontact clears the "emergencycontact" edge to the EmergencyContact entity.
func (mu *MemberUpdate) ClearEmergencycontact() *MemberUpdate {
	mu.mutation.ClearEmergencycontact()
	return mu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mu *MemberUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, mu.sqlSave, mu.mutation, mu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mu *MemberUpdate) SaveX(ctx context.Context) int {
	affected, err := mu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mu *MemberUpdate) Exec(ctx context.Context) error {
	_, err := mu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mu *MemberUpdate) ExecX(ctx context.Context) {
	if err := mu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mu *MemberUpdate) check() error {
	if v, ok := mu.mutation.Email(); ok {
		if err := member.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "Member.email": %w`, err)}
		}
	}
	return nil
}

func (mu *MemberUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := mu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(member.Table, member.Columns, sqlgraph.NewFieldSpec(member.FieldID, field.TypeInt))
	if ps := mu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mu.mutation.Name(); ok {
		_spec.SetField(member.FieldName, field.TypeString, value)
	}
	if value, ok := mu.mutation.Sex(); ok {
		_spec.SetField(member.FieldSex, field.TypeString, value)
	}
	if value, ok := mu.mutation.Email(); ok {
		_spec.SetField(member.FieldEmail, field.TypeString, value)
	}
	if mu.mutation.EmergencycontactCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   member.EmergencycontactTable,
			Columns: []string{member.EmergencycontactColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(emergencycontact.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.EmergencycontactIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   member.EmergencycontactTable,
			Columns: []string{member.EmergencycontactColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(emergencycontact.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{member.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	mu.mutation.done = true
	return n, nil
}

// MemberUpdateOne is the builder for updating a single Member entity.
type MemberUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MemberMutation
}

// SetName sets the "name" field.
func (muo *MemberUpdateOne) SetName(s string) *MemberUpdateOne {
	muo.mutation.SetName(s)
	return muo
}

// SetSex sets the "sex" field.
func (muo *MemberUpdateOne) SetSex(s string) *MemberUpdateOne {
	muo.mutation.SetSex(s)
	return muo
}

// SetEmail sets the "email" field.
func (muo *MemberUpdateOne) SetEmail(s string) *MemberUpdateOne {
	muo.mutation.SetEmail(s)
	return muo
}

// SetEmergencycontactID sets the "emergencycontact" edge to the EmergencyContact entity by ID.
func (muo *MemberUpdateOne) SetEmergencycontactID(id int) *MemberUpdateOne {
	muo.mutation.SetEmergencycontactID(id)
	return muo
}

// SetNillableEmergencycontactID sets the "emergencycontact" edge to the EmergencyContact entity by ID if the given value is not nil.
func (muo *MemberUpdateOne) SetNillableEmergencycontactID(id *int) *MemberUpdateOne {
	if id != nil {
		muo = muo.SetEmergencycontactID(*id)
	}
	return muo
}

// SetEmergencycontact sets the "emergencycontact" edge to the EmergencyContact entity.
func (muo *MemberUpdateOne) SetEmergencycontact(e *EmergencyContact) *MemberUpdateOne {
	return muo.SetEmergencycontactID(e.ID)
}

// Mutation returns the MemberMutation object of the builder.
func (muo *MemberUpdateOne) Mutation() *MemberMutation {
	return muo.mutation
}

// ClearEmergencycontact clears the "emergencycontact" edge to the EmergencyContact entity.
func (muo *MemberUpdateOne) ClearEmergencycontact() *MemberUpdateOne {
	muo.mutation.ClearEmergencycontact()
	return muo
}

// Where appends a list predicates to the MemberUpdate builder.
func (muo *MemberUpdateOne) Where(ps ...predicate.Member) *MemberUpdateOne {
	muo.mutation.Where(ps...)
	return muo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (muo *MemberUpdateOne) Select(field string, fields ...string) *MemberUpdateOne {
	muo.fields = append([]string{field}, fields...)
	return muo
}

// Save executes the query and returns the updated Member entity.
func (muo *MemberUpdateOne) Save(ctx context.Context) (*Member, error) {
	return withHooks(ctx, muo.sqlSave, muo.mutation, muo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (muo *MemberUpdateOne) SaveX(ctx context.Context) *Member {
	node, err := muo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (muo *MemberUpdateOne) Exec(ctx context.Context) error {
	_, err := muo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (muo *MemberUpdateOne) ExecX(ctx context.Context) {
	if err := muo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (muo *MemberUpdateOne) check() error {
	if v, ok := muo.mutation.Email(); ok {
		if err := member.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "Member.email": %w`, err)}
		}
	}
	return nil
}

func (muo *MemberUpdateOne) sqlSave(ctx context.Context) (_node *Member, err error) {
	if err := muo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(member.Table, member.Columns, sqlgraph.NewFieldSpec(member.FieldID, field.TypeInt))
	id, ok := muo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Member.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := muo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, member.FieldID)
		for _, f := range fields {
			if !member.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != member.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := muo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := muo.mutation.Name(); ok {
		_spec.SetField(member.FieldName, field.TypeString, value)
	}
	if value, ok := muo.mutation.Sex(); ok {
		_spec.SetField(member.FieldSex, field.TypeString, value)
	}
	if value, ok := muo.mutation.Email(); ok {
		_spec.SetField(member.FieldEmail, field.TypeString, value)
	}
	if muo.mutation.EmergencycontactCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   member.EmergencycontactTable,
			Columns: []string{member.EmergencycontactColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(emergencycontact.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.EmergencycontactIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   member.EmergencycontactTable,
			Columns: []string{member.EmergencycontactColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(emergencycontact.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Member{config: muo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, muo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{member.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	muo.mutation.done = true
	return _node, nil
}