// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"
	"yayar/internal/data/ent/history"

	"entgo.io/ent/dialect/sql"
)

// History is the model entity for the History schema.
type History struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Device holds the value of the "device" field.
	Device string `json:"device,omitempty"`
	// IP holds the value of the "ip" field.
	IP string `json:"ip,omitempty"`
	// CreatedAt holds the value of the "createdAt" field.
	CreatedAt time.Time `json:"createdAt,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the HistoryQuery when eager-loading is set.
	Edges HistoryEdges `json:"edges"`
}

// HistoryEdges holds the relations/edges for other nodes in the graph.
type HistoryEdges struct {
	// Version holds the value of the version edge.
	Version []*Version `json:"version,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// VersionOrErr returns the Version value or an error if the edge
// was not loaded in eager-loading.
func (e HistoryEdges) VersionOrErr() ([]*Version, error) {
	if e.loadedTypes[0] {
		return e.Version, nil
	}
	return nil, &NotLoadedError{edge: "version"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*History) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case history.FieldID:
			values[i] = new(sql.NullInt64)
		case history.FieldDevice, history.FieldIP:
			values[i] = new(sql.NullString)
		case history.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type History", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the History fields.
func (h *History) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case history.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			h.ID = int(value.Int64)
		case history.FieldDevice:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field device", values[i])
			} else if value.Valid {
				h.Device = value.String
			}
		case history.FieldIP:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ip", values[i])
			} else if value.Valid {
				h.IP = value.String
			}
		case history.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field createdAt", values[i])
			} else if value.Valid {
				h.CreatedAt = value.Time
			}
		}
	}
	return nil
}

// QueryVersion queries the "version" edge of the History entity.
func (h *History) QueryVersion() *VersionQuery {
	return (&HistoryClient{config: h.config}).QueryVersion(h)
}

// Update returns a builder for updating this History.
// Note that you need to call History.Unwrap() before calling this method if this History
// was returned from a transaction, and the transaction was committed or rolled back.
func (h *History) Update() *HistoryUpdateOne {
	return (&HistoryClient{config: h.config}).UpdateOne(h)
}

// Unwrap unwraps the History entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (h *History) Unwrap() *History {
	tx, ok := h.config.driver.(*txDriver)
	if !ok {
		panic("ent: History is not a transactional entity")
	}
	h.config.driver = tx.drv
	return h
}

// String implements the fmt.Stringer.
func (h *History) String() string {
	var builder strings.Builder
	builder.WriteString("History(")
	builder.WriteString(fmt.Sprintf("id=%v", h.ID))
	builder.WriteString(", device=")
	builder.WriteString(h.Device)
	builder.WriteString(", ip=")
	builder.WriteString(h.IP)
	builder.WriteString(", createdAt=")
	builder.WriteString(h.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Histories is a parsable slice of History.
type Histories []*History

func (h Histories) config(cfg config) {
	for _i := range h {
		h[_i].config = cfg
	}
}
