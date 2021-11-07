package nulltype

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

// Time represents a bool that may be null.
type Time struct {
	Time  time.Time
	Valid bool
}

// NewTime creates an instance of Time.
func NewTime(t time.Time, valid bool) Time {
	return Time{Time: t, Valid: valid}
}

// Scan implements the Scanner interface.
func (t *Time) Scan(value interface{}) error {
	if value == nil {
		t.Time, t.Valid = time.Time{}, false
		return nil
	}

	t.Valid = true
	switch data := value.(type) {
	case time.Time:
		t.Time = data
		return nil
	default:
		return fmt.Errorf("unsupported type: %T", value)
	}
}

// Value implements the driver Valuer interface.
func (t Time) Value() (driver.Value, error) {
	if !t.Valid {
		return nil, nil
	}
	return t.Time, nil
}

// MarshalJSON encode the value to JSON.
func (t Time) MarshalJSON() ([]byte, error) {
	if !t.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(t.Time)
}

// UnmarshalJSON decode data to the value.
func (t *Time) UnmarshalJSON(data []byte) error {
	var tt *time.Time
	if err := json.Unmarshal(data, &tt); err != nil {
		return err
	}
	t.Valid = tt != nil
	if tt == nil {
		t.Time = time.Time{}
	} else {
		t.Time = *tt
	}
	return nil
}

// String return string indicated the value.
func (t Time) String() string {
	if !t.Valid {
		return "<null>"
	}
	return t.Time.String()
}
