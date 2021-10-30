// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"long2ice/longurl/ent/url"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// Url is the model entity for the Url schema.
type Url struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// URL holds the value of the "url" field.
	URL string `json:"url,omitempty"`
	// Path holds the value of the "path" field.
	Path string `json:"path,omitempty"`
	// ExpireAt holds the value of the "expire_at" field.
	ExpireAt *time.Time `json:"expire_at,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Url) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case url.FieldID:
			values[i] = new(sql.NullInt64)
		case url.FieldURL, url.FieldPath:
			values[i] = new(sql.NullString)
		case url.FieldExpireAt, url.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Url", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Url fields.
func (u *Url) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case url.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			u.ID = int(value.Int64)
		case url.FieldURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field url", values[i])
			} else if value.Valid {
				u.URL = value.String
			}
		case url.FieldPath:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field path", values[i])
			} else if value.Valid {
				u.Path = value.String
			}
		case url.FieldExpireAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field expire_at", values[i])
			} else if value.Valid {
				u.ExpireAt = new(time.Time)
				*u.ExpireAt = value.Time
			}
		case url.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				u.CreatedAt = value.Time
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Url.
// Note that you need to call Url.Unwrap() before calling this method if this Url
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *Url) Update() *URLUpdateOne {
	return (&UrlClient{config: u.config}).UpdateOne(u)
}

// Unwrap unwraps the Url entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *Url) Unwrap() *Url {
	tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: Url is not a transactional entity")
	}
	u.config.driver = tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *Url) String() string {
	var builder strings.Builder
	builder.WriteString("Url(")
	builder.WriteString(fmt.Sprintf("id=%v", u.ID))
	builder.WriteString(", url=")
	builder.WriteString(u.URL)
	builder.WriteString(", path=")
	builder.WriteString(u.Path)
	if v := u.ExpireAt; v != nil {
		builder.WriteString(", expire_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", created_at=")
	builder.WriteString(u.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Urls is a parsable slice of Url.
type Urls []*Url

func (u Urls) config(cfg config) {
	for _i := range u {
		u[_i].config = cfg
	}
}
