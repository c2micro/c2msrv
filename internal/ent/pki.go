// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/c2micro/c2msrv/internal/ent/pki"
)

// Pki is the model entity for the Pki schema.
type Pki struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Time when entity was created
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Time when entity was updated
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Time when entity was soft-deleted
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	// type of certificate blob (ca, listener, operator)
	Type pki.Type `json:"type,omitempty"`
	// certificate key
	Key []byte `json:"key,omitempty"`
	// certificate chain
	Cert         []byte `json:"cert,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Pki) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case pki.FieldKey, pki.FieldCert:
			values[i] = new([]byte)
		case pki.FieldID:
			values[i] = new(sql.NullInt64)
		case pki.FieldType:
			values[i] = new(sql.NullString)
		case pki.FieldCreatedAt, pki.FieldUpdatedAt, pki.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Pki fields.
func (pk *Pki) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case pki.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pk.ID = int(value.Int64)
		case pki.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				pk.CreatedAt = value.Time
			}
		case pki.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				pk.UpdatedAt = value.Time
			}
		case pki.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				pk.DeletedAt = value.Time
			}
		case pki.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				pk.Type = pki.Type(value.String)
			}
		case pki.FieldKey:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field key", values[i])
			} else if value != nil {
				pk.Key = *value
			}
		case pki.FieldCert:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field cert", values[i])
			} else if value != nil {
				pk.Cert = *value
			}
		default:
			pk.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Pki.
// This includes values selected through modifiers, order, etc.
func (pk *Pki) Value(name string) (ent.Value, error) {
	return pk.selectValues.Get(name)
}

// Update returns a builder for updating this Pki.
// Note that you need to call Pki.Unwrap() before calling this method if this Pki
// was returned from a transaction, and the transaction was committed or rolled back.
func (pk *Pki) Update() *PkiUpdateOne {
	return NewPkiClient(pk.config).UpdateOne(pk)
}

// Unwrap unwraps the Pki entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pk *Pki) Unwrap() *Pki {
	_tx, ok := pk.config.driver.(*txDriver)
	if !ok {
		panic("ent: Pki is not a transactional entity")
	}
	pk.config.driver = _tx.drv
	return pk
}

// String implements the fmt.Stringer.
func (pk *Pki) String() string {
	var builder strings.Builder
	builder.WriteString("Pki(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pk.ID))
	builder.WriteString("created_at=")
	builder.WriteString(pk.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(pk.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(pk.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(fmt.Sprintf("%v", pk.Type))
	builder.WriteString(", ")
	builder.WriteString("key=")
	builder.WriteString(fmt.Sprintf("%v", pk.Key))
	builder.WriteString(", ")
	builder.WriteString("cert=")
	builder.WriteString(fmt.Sprintf("%v", pk.Cert))
	builder.WriteByte(')')
	return builder.String()
}

// Pkis is a parsable slice of Pki.
type Pkis []*Pki