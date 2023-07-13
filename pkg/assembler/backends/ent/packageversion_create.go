// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/billofmaterials"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/occurrence"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/packagename"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/packageversion"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/pkgequal"
	"github.com/guacsec/guac/pkg/assembler/graphql/model"
)

// PackageVersionCreate is the builder for creating a PackageVersion entity.
type PackageVersionCreate struct {
	config
	mutation *PackageVersionMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetNameID sets the "name_id" field.
func (pvc *PackageVersionCreate) SetNameID(i int) *PackageVersionCreate {
	pvc.mutation.SetNameID(i)
	return pvc
}

// SetVersion sets the "version" field.
func (pvc *PackageVersionCreate) SetVersion(s string) *PackageVersionCreate {
	pvc.mutation.SetVersion(s)
	return pvc
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (pvc *PackageVersionCreate) SetNillableVersion(s *string) *PackageVersionCreate {
	if s != nil {
		pvc.SetVersion(*s)
	}
	return pvc
}

// SetSubpath sets the "subpath" field.
func (pvc *PackageVersionCreate) SetSubpath(s string) *PackageVersionCreate {
	pvc.mutation.SetSubpath(s)
	return pvc
}

// SetNillableSubpath sets the "subpath" field if the given value is not nil.
func (pvc *PackageVersionCreate) SetNillableSubpath(s *string) *PackageVersionCreate {
	if s != nil {
		pvc.SetSubpath(*s)
	}
	return pvc
}

// SetQualifiers sets the "qualifiers" field.
func (pvc *PackageVersionCreate) SetQualifiers(mq []model.PackageQualifier) *PackageVersionCreate {
	pvc.mutation.SetQualifiers(mq)
	return pvc
}

// SetHash sets the "hash" field.
func (pvc *PackageVersionCreate) SetHash(s string) *PackageVersionCreate {
	pvc.mutation.SetHash(s)
	return pvc
}

// SetName sets the "name" edge to the PackageName entity.
func (pvc *PackageVersionCreate) SetName(p *PackageName) *PackageVersionCreate {
	return pvc.SetNameID(p.ID)
}

// AddOccurrenceIDs adds the "occurrences" edge to the Occurrence entity by IDs.
func (pvc *PackageVersionCreate) AddOccurrenceIDs(ids ...int) *PackageVersionCreate {
	pvc.mutation.AddOccurrenceIDs(ids...)
	return pvc
}

// AddOccurrences adds the "occurrences" edges to the Occurrence entity.
func (pvc *PackageVersionCreate) AddOccurrences(o ...*Occurrence) *PackageVersionCreate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return pvc.AddOccurrenceIDs(ids...)
}

// AddSbomIDs adds the "sbom" edge to the BillOfMaterials entity by IDs.
func (pvc *PackageVersionCreate) AddSbomIDs(ids ...int) *PackageVersionCreate {
	pvc.mutation.AddSbomIDs(ids...)
	return pvc
}

// AddSbom adds the "sbom" edges to the BillOfMaterials entity.
func (pvc *PackageVersionCreate) AddSbom(b ...*BillOfMaterials) *PackageVersionCreate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return pvc.AddSbomIDs(ids...)
}

// AddSimilarIDs adds the "similar" edge to the PackageVersion entity by IDs.
func (pvc *PackageVersionCreate) AddSimilarIDs(ids ...int) *PackageVersionCreate {
	pvc.mutation.AddSimilarIDs(ids...)
	return pvc
}

// AddSimilar adds the "similar" edges to the PackageVersion entity.
func (pvc *PackageVersionCreate) AddSimilar(p ...*PackageVersion) *PackageVersionCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pvc.AddSimilarIDs(ids...)
}

// AddEqualIDs adds the "equal" edge to the PkgEqual entity by IDs.
func (pvc *PackageVersionCreate) AddEqualIDs(ids ...int) *PackageVersionCreate {
	pvc.mutation.AddEqualIDs(ids...)
	return pvc
}

// AddEqual adds the "equal" edges to the PkgEqual entity.
func (pvc *PackageVersionCreate) AddEqual(p ...*PkgEqual) *PackageVersionCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pvc.AddEqualIDs(ids...)
}

// Mutation returns the PackageVersionMutation object of the builder.
func (pvc *PackageVersionCreate) Mutation() *PackageVersionMutation {
	return pvc.mutation
}

// Save creates the PackageVersion in the database.
func (pvc *PackageVersionCreate) Save(ctx context.Context) (*PackageVersion, error) {
	pvc.defaults()
	return withHooks(ctx, pvc.sqlSave, pvc.mutation, pvc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pvc *PackageVersionCreate) SaveX(ctx context.Context) *PackageVersion {
	v, err := pvc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pvc *PackageVersionCreate) Exec(ctx context.Context) error {
	_, err := pvc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pvc *PackageVersionCreate) ExecX(ctx context.Context) {
	if err := pvc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pvc *PackageVersionCreate) defaults() {
	if _, ok := pvc.mutation.Version(); !ok {
		v := packageversion.DefaultVersion
		pvc.mutation.SetVersion(v)
	}
	if _, ok := pvc.mutation.Subpath(); !ok {
		v := packageversion.DefaultSubpath
		pvc.mutation.SetSubpath(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pvc *PackageVersionCreate) check() error {
	if _, ok := pvc.mutation.NameID(); !ok {
		return &ValidationError{Name: "name_id", err: errors.New(`ent: missing required field "PackageVersion.name_id"`)}
	}
	if _, ok := pvc.mutation.Version(); !ok {
		return &ValidationError{Name: "version", err: errors.New(`ent: missing required field "PackageVersion.version"`)}
	}
	if _, ok := pvc.mutation.Subpath(); !ok {
		return &ValidationError{Name: "subpath", err: errors.New(`ent: missing required field "PackageVersion.subpath"`)}
	}
	if _, ok := pvc.mutation.Hash(); !ok {
		return &ValidationError{Name: "hash", err: errors.New(`ent: missing required field "PackageVersion.hash"`)}
	}
	if _, ok := pvc.mutation.NameID(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required edge "PackageVersion.name"`)}
	}
	return nil
}

func (pvc *PackageVersionCreate) sqlSave(ctx context.Context) (*PackageVersion, error) {
	if err := pvc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pvc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pvc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	pvc.mutation.id = &_node.ID
	pvc.mutation.done = true
	return _node, nil
}

func (pvc *PackageVersionCreate) createSpec() (*PackageVersion, *sqlgraph.CreateSpec) {
	var (
		_node = &PackageVersion{config: pvc.config}
		_spec = sqlgraph.NewCreateSpec(packageversion.Table, sqlgraph.NewFieldSpec(packageversion.FieldID, field.TypeInt))
	)
	_spec.OnConflict = pvc.conflict
	if value, ok := pvc.mutation.Version(); ok {
		_spec.SetField(packageversion.FieldVersion, field.TypeString, value)
		_node.Version = value
	}
	if value, ok := pvc.mutation.Subpath(); ok {
		_spec.SetField(packageversion.FieldSubpath, field.TypeString, value)
		_node.Subpath = value
	}
	if value, ok := pvc.mutation.Qualifiers(); ok {
		_spec.SetField(packageversion.FieldQualifiers, field.TypeJSON, value)
		_node.Qualifiers = value
	}
	if value, ok := pvc.mutation.Hash(); ok {
		_spec.SetField(packageversion.FieldHash, field.TypeString, value)
		_node.Hash = value
	}
	if nodes := pvc.mutation.NameIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   packageversion.NameTable,
			Columns: []string{packageversion.NameColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(packagename.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.NameID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pvc.mutation.OccurrencesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   packageversion.OccurrencesTable,
			Columns: []string{packageversion.OccurrencesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(occurrence.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pvc.mutation.SbomIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   packageversion.SbomTable,
			Columns: []string{packageversion.SbomColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(billofmaterials.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pvc.mutation.SimilarIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   packageversion.SimilarTable,
			Columns: packageversion.SimilarPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(packageversion.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pvc.mutation.EqualIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   packageversion.EqualTable,
			Columns: []string{packageversion.EqualColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(pkgequal.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.PackageVersion.Create().
//		SetNameID(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.PackageVersionUpsert) {
//			SetNameID(v+v).
//		}).
//		Exec(ctx)
func (pvc *PackageVersionCreate) OnConflict(opts ...sql.ConflictOption) *PackageVersionUpsertOne {
	pvc.conflict = opts
	return &PackageVersionUpsertOne{
		create: pvc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.PackageVersion.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (pvc *PackageVersionCreate) OnConflictColumns(columns ...string) *PackageVersionUpsertOne {
	pvc.conflict = append(pvc.conflict, sql.ConflictColumns(columns...))
	return &PackageVersionUpsertOne{
		create: pvc,
	}
}

type (
	// PackageVersionUpsertOne is the builder for "upsert"-ing
	//  one PackageVersion node.
	PackageVersionUpsertOne struct {
		create *PackageVersionCreate
	}

	// PackageVersionUpsert is the "OnConflict" setter.
	PackageVersionUpsert struct {
		*sql.UpdateSet
	}
)

// SetNameID sets the "name_id" field.
func (u *PackageVersionUpsert) SetNameID(v int) *PackageVersionUpsert {
	u.Set(packageversion.FieldNameID, v)
	return u
}

// UpdateNameID sets the "name_id" field to the value that was provided on create.
func (u *PackageVersionUpsert) UpdateNameID() *PackageVersionUpsert {
	u.SetExcluded(packageversion.FieldNameID)
	return u
}

// SetVersion sets the "version" field.
func (u *PackageVersionUpsert) SetVersion(v string) *PackageVersionUpsert {
	u.Set(packageversion.FieldVersion, v)
	return u
}

// UpdateVersion sets the "version" field to the value that was provided on create.
func (u *PackageVersionUpsert) UpdateVersion() *PackageVersionUpsert {
	u.SetExcluded(packageversion.FieldVersion)
	return u
}

// SetSubpath sets the "subpath" field.
func (u *PackageVersionUpsert) SetSubpath(v string) *PackageVersionUpsert {
	u.Set(packageversion.FieldSubpath, v)
	return u
}

// UpdateSubpath sets the "subpath" field to the value that was provided on create.
func (u *PackageVersionUpsert) UpdateSubpath() *PackageVersionUpsert {
	u.SetExcluded(packageversion.FieldSubpath)
	return u
}

// SetQualifiers sets the "qualifiers" field.
func (u *PackageVersionUpsert) SetQualifiers(v []model.PackageQualifier) *PackageVersionUpsert {
	u.Set(packageversion.FieldQualifiers, v)
	return u
}

// UpdateQualifiers sets the "qualifiers" field to the value that was provided on create.
func (u *PackageVersionUpsert) UpdateQualifiers() *PackageVersionUpsert {
	u.SetExcluded(packageversion.FieldQualifiers)
	return u
}

// ClearQualifiers clears the value of the "qualifiers" field.
func (u *PackageVersionUpsert) ClearQualifiers() *PackageVersionUpsert {
	u.SetNull(packageversion.FieldQualifiers)
	return u
}

// SetHash sets the "hash" field.
func (u *PackageVersionUpsert) SetHash(v string) *PackageVersionUpsert {
	u.Set(packageversion.FieldHash, v)
	return u
}

// UpdateHash sets the "hash" field to the value that was provided on create.
func (u *PackageVersionUpsert) UpdateHash() *PackageVersionUpsert {
	u.SetExcluded(packageversion.FieldHash)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.PackageVersion.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *PackageVersionUpsertOne) UpdateNewValues() *PackageVersionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.PackageVersion.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *PackageVersionUpsertOne) Ignore() *PackageVersionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *PackageVersionUpsertOne) DoNothing() *PackageVersionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the PackageVersionCreate.OnConflict
// documentation for more info.
func (u *PackageVersionUpsertOne) Update(set func(*PackageVersionUpsert)) *PackageVersionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&PackageVersionUpsert{UpdateSet: update})
	}))
	return u
}

// SetNameID sets the "name_id" field.
func (u *PackageVersionUpsertOne) SetNameID(v int) *PackageVersionUpsertOne {
	return u.Update(func(s *PackageVersionUpsert) {
		s.SetNameID(v)
	})
}

// UpdateNameID sets the "name_id" field to the value that was provided on create.
func (u *PackageVersionUpsertOne) UpdateNameID() *PackageVersionUpsertOne {
	return u.Update(func(s *PackageVersionUpsert) {
		s.UpdateNameID()
	})
}

// SetVersion sets the "version" field.
func (u *PackageVersionUpsertOne) SetVersion(v string) *PackageVersionUpsertOne {
	return u.Update(func(s *PackageVersionUpsert) {
		s.SetVersion(v)
	})
}

// UpdateVersion sets the "version" field to the value that was provided on create.
func (u *PackageVersionUpsertOne) UpdateVersion() *PackageVersionUpsertOne {
	return u.Update(func(s *PackageVersionUpsert) {
		s.UpdateVersion()
	})
}

// SetSubpath sets the "subpath" field.
func (u *PackageVersionUpsertOne) SetSubpath(v string) *PackageVersionUpsertOne {
	return u.Update(func(s *PackageVersionUpsert) {
		s.SetSubpath(v)
	})
}

// UpdateSubpath sets the "subpath" field to the value that was provided on create.
func (u *PackageVersionUpsertOne) UpdateSubpath() *PackageVersionUpsertOne {
	return u.Update(func(s *PackageVersionUpsert) {
		s.UpdateSubpath()
	})
}

// SetQualifiers sets the "qualifiers" field.
func (u *PackageVersionUpsertOne) SetQualifiers(v []model.PackageQualifier) *PackageVersionUpsertOne {
	return u.Update(func(s *PackageVersionUpsert) {
		s.SetQualifiers(v)
	})
}

// UpdateQualifiers sets the "qualifiers" field to the value that was provided on create.
func (u *PackageVersionUpsertOne) UpdateQualifiers() *PackageVersionUpsertOne {
	return u.Update(func(s *PackageVersionUpsert) {
		s.UpdateQualifiers()
	})
}

// ClearQualifiers clears the value of the "qualifiers" field.
func (u *PackageVersionUpsertOne) ClearQualifiers() *PackageVersionUpsertOne {
	return u.Update(func(s *PackageVersionUpsert) {
		s.ClearQualifiers()
	})
}

// SetHash sets the "hash" field.
func (u *PackageVersionUpsertOne) SetHash(v string) *PackageVersionUpsertOne {
	return u.Update(func(s *PackageVersionUpsert) {
		s.SetHash(v)
	})
}

// UpdateHash sets the "hash" field to the value that was provided on create.
func (u *PackageVersionUpsertOne) UpdateHash() *PackageVersionUpsertOne {
	return u.Update(func(s *PackageVersionUpsert) {
		s.UpdateHash()
	})
}

// Exec executes the query.
func (u *PackageVersionUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for PackageVersionCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *PackageVersionUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *PackageVersionUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *PackageVersionUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// PackageVersionCreateBulk is the builder for creating many PackageVersion entities in bulk.
type PackageVersionCreateBulk struct {
	config
	builders []*PackageVersionCreate
	conflict []sql.ConflictOption
}

// Save creates the PackageVersion entities in the database.
func (pvcb *PackageVersionCreateBulk) Save(ctx context.Context) ([]*PackageVersion, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pvcb.builders))
	nodes := make([]*PackageVersion, len(pvcb.builders))
	mutators := make([]Mutator, len(pvcb.builders))
	for i := range pvcb.builders {
		func(i int, root context.Context) {
			builder := pvcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PackageVersionMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, pvcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = pvcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pvcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pvcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pvcb *PackageVersionCreateBulk) SaveX(ctx context.Context) []*PackageVersion {
	v, err := pvcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pvcb *PackageVersionCreateBulk) Exec(ctx context.Context) error {
	_, err := pvcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pvcb *PackageVersionCreateBulk) ExecX(ctx context.Context) {
	if err := pvcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.PackageVersion.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.PackageVersionUpsert) {
//			SetNameID(v+v).
//		}).
//		Exec(ctx)
func (pvcb *PackageVersionCreateBulk) OnConflict(opts ...sql.ConflictOption) *PackageVersionUpsertBulk {
	pvcb.conflict = opts
	return &PackageVersionUpsertBulk{
		create: pvcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.PackageVersion.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (pvcb *PackageVersionCreateBulk) OnConflictColumns(columns ...string) *PackageVersionUpsertBulk {
	pvcb.conflict = append(pvcb.conflict, sql.ConflictColumns(columns...))
	return &PackageVersionUpsertBulk{
		create: pvcb,
	}
}

// PackageVersionUpsertBulk is the builder for "upsert"-ing
// a bulk of PackageVersion nodes.
type PackageVersionUpsertBulk struct {
	create *PackageVersionCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.PackageVersion.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *PackageVersionUpsertBulk) UpdateNewValues() *PackageVersionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.PackageVersion.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *PackageVersionUpsertBulk) Ignore() *PackageVersionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *PackageVersionUpsertBulk) DoNothing() *PackageVersionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the PackageVersionCreateBulk.OnConflict
// documentation for more info.
func (u *PackageVersionUpsertBulk) Update(set func(*PackageVersionUpsert)) *PackageVersionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&PackageVersionUpsert{UpdateSet: update})
	}))
	return u
}

// SetNameID sets the "name_id" field.
func (u *PackageVersionUpsertBulk) SetNameID(v int) *PackageVersionUpsertBulk {
	return u.Update(func(s *PackageVersionUpsert) {
		s.SetNameID(v)
	})
}

// UpdateNameID sets the "name_id" field to the value that was provided on create.
func (u *PackageVersionUpsertBulk) UpdateNameID() *PackageVersionUpsertBulk {
	return u.Update(func(s *PackageVersionUpsert) {
		s.UpdateNameID()
	})
}

// SetVersion sets the "version" field.
func (u *PackageVersionUpsertBulk) SetVersion(v string) *PackageVersionUpsertBulk {
	return u.Update(func(s *PackageVersionUpsert) {
		s.SetVersion(v)
	})
}

// UpdateVersion sets the "version" field to the value that was provided on create.
func (u *PackageVersionUpsertBulk) UpdateVersion() *PackageVersionUpsertBulk {
	return u.Update(func(s *PackageVersionUpsert) {
		s.UpdateVersion()
	})
}

// SetSubpath sets the "subpath" field.
func (u *PackageVersionUpsertBulk) SetSubpath(v string) *PackageVersionUpsertBulk {
	return u.Update(func(s *PackageVersionUpsert) {
		s.SetSubpath(v)
	})
}

// UpdateSubpath sets the "subpath" field to the value that was provided on create.
func (u *PackageVersionUpsertBulk) UpdateSubpath() *PackageVersionUpsertBulk {
	return u.Update(func(s *PackageVersionUpsert) {
		s.UpdateSubpath()
	})
}

// SetQualifiers sets the "qualifiers" field.
func (u *PackageVersionUpsertBulk) SetQualifiers(v []model.PackageQualifier) *PackageVersionUpsertBulk {
	return u.Update(func(s *PackageVersionUpsert) {
		s.SetQualifiers(v)
	})
}

// UpdateQualifiers sets the "qualifiers" field to the value that was provided on create.
func (u *PackageVersionUpsertBulk) UpdateQualifiers() *PackageVersionUpsertBulk {
	return u.Update(func(s *PackageVersionUpsert) {
		s.UpdateQualifiers()
	})
}

// ClearQualifiers clears the value of the "qualifiers" field.
func (u *PackageVersionUpsertBulk) ClearQualifiers() *PackageVersionUpsertBulk {
	return u.Update(func(s *PackageVersionUpsert) {
		s.ClearQualifiers()
	})
}

// SetHash sets the "hash" field.
func (u *PackageVersionUpsertBulk) SetHash(v string) *PackageVersionUpsertBulk {
	return u.Update(func(s *PackageVersionUpsert) {
		s.SetHash(v)
	})
}

// UpdateHash sets the "hash" field to the value that was provided on create.
func (u *PackageVersionUpsertBulk) UpdateHash() *PackageVersionUpsertBulk {
	return u.Update(func(s *PackageVersionUpsert) {
		s.UpdateHash()
	})
}

// Exec executes the query.
func (u *PackageVersionUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the PackageVersionCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for PackageVersionCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *PackageVersionUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
