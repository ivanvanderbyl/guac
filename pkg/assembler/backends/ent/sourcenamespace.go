// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/sourcenamespace"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/sourcetype"
)

// SourceNamespace is the model entity for the SourceNamespace schema.
type SourceNamespace struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Namespace holds the value of the "namespace" field.
	Namespace string `json:"namespace,omitempty"`
	// SourceID holds the value of the "source_id" field.
	SourceID int `json:"source_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SourceNamespaceQuery when eager-loading is set.
	Edges        SourceNamespaceEdges `json:"edges"`
	selectValues sql.SelectValues
}

// SourceNamespaceEdges holds the relations/edges for other nodes in the graph.
type SourceNamespaceEdges struct {
	// SourceType holds the value of the source_type edge.
	SourceType *SourceType `json:"source_type,omitempty"`
	// Names holds the value of the names edge.
	Names []*SourceName `json:"names,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// SourceTypeOrErr returns the SourceType value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SourceNamespaceEdges) SourceTypeOrErr() (*SourceType, error) {
	if e.loadedTypes[0] {
		if e.SourceType == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: sourcetype.Label}
		}
		return e.SourceType, nil
	}
	return nil, &NotLoadedError{edge: "source_type"}
}

// NamesOrErr returns the Names value or an error if the edge
// was not loaded in eager-loading.
func (e SourceNamespaceEdges) NamesOrErr() ([]*SourceName, error) {
	if e.loadedTypes[1] {
		return e.Names, nil
	}
	return nil, &NotLoadedError{edge: "names"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*SourceNamespace) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case sourcenamespace.FieldID, sourcenamespace.FieldSourceID:
			values[i] = new(sql.NullInt64)
		case sourcenamespace.FieldNamespace:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the SourceNamespace fields.
func (sn *SourceNamespace) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case sourcenamespace.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			sn.ID = int(value.Int64)
		case sourcenamespace.FieldNamespace:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field namespace", values[i])
			} else if value.Valid {
				sn.Namespace = value.String
			}
		case sourcenamespace.FieldSourceID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field source_id", values[i])
			} else if value.Valid {
				sn.SourceID = int(value.Int64)
			}
		default:
			sn.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the SourceNamespace.
// This includes values selected through modifiers, order, etc.
func (sn *SourceNamespace) Value(name string) (ent.Value, error) {
	return sn.selectValues.Get(name)
}

// QuerySourceType queries the "source_type" edge of the SourceNamespace entity.
func (sn *SourceNamespace) QuerySourceType() *SourceTypeQuery {
	return NewSourceNamespaceClient(sn.config).QuerySourceType(sn)
}

// QueryNames queries the "names" edge of the SourceNamespace entity.
func (sn *SourceNamespace) QueryNames() *SourceNameQuery {
	return NewSourceNamespaceClient(sn.config).QueryNames(sn)
}

// Update returns a builder for updating this SourceNamespace.
// Note that you need to call SourceNamespace.Unwrap() before calling this method if this SourceNamespace
// was returned from a transaction, and the transaction was committed or rolled back.
func (sn *SourceNamespace) Update() *SourceNamespaceUpdateOne {
	return NewSourceNamespaceClient(sn.config).UpdateOne(sn)
}

// Unwrap unwraps the SourceNamespace entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (sn *SourceNamespace) Unwrap() *SourceNamespace {
	_tx, ok := sn.config.driver.(*txDriver)
	if !ok {
		panic("ent: SourceNamespace is not a transactional entity")
	}
	sn.config.driver = _tx.drv
	return sn
}

// String implements the fmt.Stringer.
func (sn *SourceNamespace) String() string {
	var builder strings.Builder
	builder.WriteString("SourceNamespace(")
	builder.WriteString(fmt.Sprintf("id=%v, ", sn.ID))
	builder.WriteString("namespace=")
	builder.WriteString(sn.Namespace)
	builder.WriteString(", ")
	builder.WriteString("source_id=")
	builder.WriteString(fmt.Sprintf("%v", sn.SourceID))
	builder.WriteByte(')')
	return builder.String()
}

// SourceNamespaces is a parsable slice of SourceNamespace.
type SourceNamespaces []*SourceNamespace
