package nulltype

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

// NewByte creates an instance of Byte.
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
	return b.Byte, nil
}

// MarshalJSON encode the value to JSON.
func (b Byte) MarshalJSON() ([]byte, error) {
	if !b.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(b.Byte)
}

// UnmarshalJSON decode data to the value.
func (b *Byte) UnmarshalJSON(data []byte) error {
	var bb *byte
	if err := json.Unmarshal(data, &bb); err != nil {
		return err
	}
	b.Valid = bb != nil
	if bb == nil {
		b.Byte = byte(0)
	} else {
		b.Byte = *bb
	}
	return nil
}
