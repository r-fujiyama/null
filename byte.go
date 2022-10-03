package null

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

// Byte represents a byte that may be null.
type Byte struct {
	Byte  byte
	Valid bool
}

// NewByte creates a new Byte
func NewByte(b byte, valid bool) Byte {
	return Byte{Byte: b, Valid: valid}
}

// Scan implements the Scanner interface.
func (b *Byte) Scan(value interface{}) error {
	if value == nil {
		b.Byte, b.Valid = byte(0), false
		return nil
	}

	b.Valid = true
	switch data := value.(type) {
	case byte:
		b.Byte = data
		return nil
	default:
		return fmt.Errorf("unsupported type: %T", value)
	}
}

// Value implements the driver Valuer interface.
func (b Byte) Value() (driver.Value, error) {
	if !b.Valid {
		return nil, nil
	}
	return int64(b.Byte), nil
}

// MarshalJSON encode the value to JSON.
func (b Byte) MarshalJSON() ([]byte, error) {
	if !b.Valid {
		return []byte("null"), nil
	}
	return jsonMarshal(b.Byte)
}

// UnmarshalJSON decode data to the value.
func (b *Byte) UnmarshalJSON(data []byte) error {
	var bb *byte
	if err := json.Unmarshal(data, &bb); err != nil {
		return err
	}
	b.Valid = bb != nil
	if b.Valid {
		b.Byte = *bb
	} else {
		b.Byte = byte(0)
	}
	return nil
}

// IsNull returns true if Valid is false.
func (b *Byte) IsNull() bool {
	return !b.Valid
}
