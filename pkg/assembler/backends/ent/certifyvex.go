// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/artifact"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/certifyvex"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/packageversion"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/vulnerabilitytype"
)

// CertifyVex is the model entity for the CertifyVex schema.
type CertifyVex struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// PackageID holds the value of the "package_id" field.
	PackageID *int `json:"package_id,omitempty"`
	// ArtifactID holds the value of the "artifact_id" field.
	ArtifactID *int `json:"artifact_id,omitempty"`
	// Vulnerability is one of OSV, GHSA, or CVE, or nil if not vulnerable
	VulnerabilityID int `json:"vulnerability_id,omitempty"`
	// KnownSince holds the value of the "knownSince" field.
	KnownSince time.Time `json:"knownSince,omitempty"`
	// Status holds the value of the "status" field.
	Status string `json:"status,omitempty"`
	// Statement holds the value of the "statement" field.
	Statement string `json:"statement,omitempty"`
	// StatusNotes holds the value of the "statusNotes" field.
	StatusNotes string `json:"statusNotes,omitempty"`
	// Justification holds the value of the "justification" field.
	Justification string `json:"justification,omitempty"`
	// Origin holds the value of the "origin" field.
	Origin string `json:"origin,omitempty"`
	// Collector holds the value of the "collector" field.
	Collector string `json:"collector,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CertifyVexQuery when eager-loading is set.
	Edges        CertifyVexEdges `json:"edges"`
	selectValues sql.SelectValues
}

// CertifyVexEdges holds the relations/edges for other nodes in the graph.
type CertifyVexEdges struct {
	// Package holds the value of the package edge.
	Package *PackageVersion `json:"package,omitempty"`
	// Artifact holds the value of the artifact edge.
	Artifact *Artifact `json:"artifact,omitempty"`
	// Vulnerability is one of OSV, GHSA, or CVE
	Vulnerability *VulnerabilityType `json:"vulnerability,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
	// totalCount holds the count of the edges above.
	totalCount [3]map[string]int
}

// PackageOrErr returns the Package value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CertifyVexEdges) PackageOrErr() (*PackageVersion, error) {
	if e.loadedTypes[0] {
		if e.Package == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: packageversion.Label}
		}
		return e.Package, nil
	}
	return nil, &NotLoadedError{edge: "package"}
}

// ArtifactOrErr returns the Artifact value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CertifyVexEdges) ArtifactOrErr() (*Artifact, error) {
	if e.loadedTypes[1] {
		if e.Artifact == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: artifact.Label}
		}
		return e.Artifact, nil
	}
	return nil, &NotLoadedError{edge: "artifact"}
}

// VulnerabilityOrErr returns the Vulnerability value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CertifyVexEdges) VulnerabilityOrErr() (*VulnerabilityType, error) {
	if e.loadedTypes[2] {
		if e.Vulnerability == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: vulnerabilitytype.Label}
		}
		return e.Vulnerability, nil
	}
	return nil, &NotLoadedError{edge: "vulnerability"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*CertifyVex) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case certifyvex.FieldID, certifyvex.FieldPackageID, certifyvex.FieldArtifactID, certifyvex.FieldVulnerabilityID:
			values[i] = new(sql.NullInt64)
		case certifyvex.FieldStatus, certifyvex.FieldStatement, certifyvex.FieldStatusNotes, certifyvex.FieldJustification, certifyvex.FieldOrigin, certifyvex.FieldCollector:
			values[i] = new(sql.NullString)
		case certifyvex.FieldKnownSince:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the CertifyVex fields.
func (cv *CertifyVex) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case certifyvex.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			cv.ID = int(value.Int64)
		case certifyvex.FieldPackageID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field package_id", values[i])
			} else if value.Valid {
				cv.PackageID = new(int)
				*cv.PackageID = int(value.Int64)
			}
		case certifyvex.FieldArtifactID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field artifact_id", values[i])
			} else if value.Valid {
				cv.ArtifactID = new(int)
				*cv.ArtifactID = int(value.Int64)
			}
		case certifyvex.FieldVulnerabilityID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field vulnerability_id", values[i])
			} else if value.Valid {
				cv.VulnerabilityID = int(value.Int64)
			}
		case certifyvex.FieldKnownSince:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field knownSince", values[i])
			} else if value.Valid {
				cv.KnownSince = value.Time
			}
		case certifyvex.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				cv.Status = value.String
			}
		case certifyvex.FieldStatement:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field statement", values[i])
			} else if value.Valid {
				cv.Statement = value.String
			}
		case certifyvex.FieldStatusNotes:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field statusNotes", values[i])
			} else if value.Valid {
				cv.StatusNotes = value.String
			}
		case certifyvex.FieldJustification:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field justification", values[i])
			} else if value.Valid {
				cv.Justification = value.String
			}
		case certifyvex.FieldOrigin:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field origin", values[i])
			} else if value.Valid {
				cv.Origin = value.String
			}
		case certifyvex.FieldCollector:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field collector", values[i])
			} else if value.Valid {
				cv.Collector = value.String
			}
		default:
			cv.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the CertifyVex.
// This includes values selected through modifiers, order, etc.
func (cv *CertifyVex) Value(name string) (ent.Value, error) {
	return cv.selectValues.Get(name)
}

// QueryPackage queries the "package" edge of the CertifyVex entity.
func (cv *CertifyVex) QueryPackage() *PackageVersionQuery {
	return NewCertifyVexClient(cv.config).QueryPackage(cv)
}

// QueryArtifact queries the "artifact" edge of the CertifyVex entity.
func (cv *CertifyVex) QueryArtifact() *ArtifactQuery {
	return NewCertifyVexClient(cv.config).QueryArtifact(cv)
}

// QueryVulnerability queries the "vulnerability" edge of the CertifyVex entity.
func (cv *CertifyVex) QueryVulnerability() *VulnerabilityTypeQuery {
	return NewCertifyVexClient(cv.config).QueryVulnerability(cv)
}

// Update returns a builder for updating this CertifyVex.
// Note that you need to call CertifyVex.Unwrap() before calling this method if this CertifyVex
// was returned from a transaction, and the transaction was committed or rolled back.
func (cv *CertifyVex) Update() *CertifyVexUpdateOne {
	return NewCertifyVexClient(cv.config).UpdateOne(cv)
}

// Unwrap unwraps the CertifyVex entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (cv *CertifyVex) Unwrap() *CertifyVex {
	_tx, ok := cv.config.driver.(*txDriver)
	if !ok {
		panic("ent: CertifyVex is not a transactional entity")
	}
	cv.config.driver = _tx.drv
	return cv
}

// String implements the fmt.Stringer.
func (cv *CertifyVex) String() string {
	var builder strings.Builder
	builder.WriteString("CertifyVex(")
	builder.WriteString(fmt.Sprintf("id=%v, ", cv.ID))
	if v := cv.PackageID; v != nil {
		builder.WriteString("package_id=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := cv.ArtifactID; v != nil {
		builder.WriteString("artifact_id=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	builder.WriteString("vulnerability_id=")
	builder.WriteString(fmt.Sprintf("%v", cv.VulnerabilityID))
	builder.WriteString(", ")
	builder.WriteString("knownSince=")
	builder.WriteString(cv.KnownSince.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(cv.Status)
	builder.WriteString(", ")
	builder.WriteString("statement=")
	builder.WriteString(cv.Statement)
	builder.WriteString(", ")
	builder.WriteString("statusNotes=")
	builder.WriteString(cv.StatusNotes)
	builder.WriteString(", ")
	builder.WriteString("justification=")
	builder.WriteString(cv.Justification)
	builder.WriteString(", ")
	builder.WriteString("origin=")
	builder.WriteString(cv.Origin)
	builder.WriteString(", ")
	builder.WriteString("collector=")
	builder.WriteString(cv.Collector)
	builder.WriteByte(')')
	return builder.String()
}

// CertifyVexes is a parsable slice of CertifyVex.
type CertifyVexes []*CertifyVex
