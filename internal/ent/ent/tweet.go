// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/codesmith-dev/twitter/internal/ent/ent/tweet"
)

// Tweet is the model entity for the Tweet schema.
type Tweet struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Content holds the value of the "content" field.
	Content string `json:"content,omitempty"`
	// Userid holds the value of the "userid" field.
	Userid       string `json:"userid,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Tweet) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case tweet.FieldID, tweet.FieldContent, tweet.FieldUserid:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Tweet fields.
func (t *Tweet) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case tweet.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				t.ID = value.String
			}
		case tweet.FieldContent:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field content", values[i])
			} else if value.Valid {
				t.Content = value.String
			}
		case tweet.FieldUserid:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field userid", values[i])
			} else if value.Valid {
				t.Userid = value.String
			}
		default:
			t.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Tweet.
// This includes values selected through modifiers, order, etc.
func (t *Tweet) Value(name string) (ent.Value, error) {
	return t.selectValues.Get(name)
}

// Update returns a builder for updating this Tweet.
// Note that you need to call Tweet.Unwrap() before calling this method if this Tweet
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Tweet) Update() *TweetUpdateOne {
	return NewTweetClient(t.config).UpdateOne(t)
}

// Unwrap unwraps the Tweet entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Tweet) Unwrap() *Tweet {
	_tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Tweet is not a transactional entity")
	}
	t.config.driver = _tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Tweet) String() string {
	var builder strings.Builder
	builder.WriteString("Tweet(")
	builder.WriteString(fmt.Sprintf("id=%v, ", t.ID))
	builder.WriteString("content=")
	builder.WriteString(t.Content)
	builder.WriteString(", ")
	builder.WriteString("userid=")
	builder.WriteString(t.Userid)
	builder.WriteByte(')')
	return builder.String()
}

// Tweets is a parsable slice of Tweet.
type Tweets []*Tweet
