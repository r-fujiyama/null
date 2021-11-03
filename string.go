package null

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"unsafe"
)

type String struct {
	String string
	Valid  bool
}

func NewString(str string) String {
	return String{String: str, Valid: str != ""}
}

func (s *String) Scan(value interface{}) error {
	if value == nil {
		s.String, s.Valid = "", false
		return nil
	}

	if data, ok := value.([]byte); ok {
		s.String, s.Valid = *(*string)(unsafe.Pointer(&data)), true
		return nil
	}
	return fmt.Errorf("got data of type %T but wanted []uint8", value)
}

func (s *String) Value() (driver.Value, error) {
	if !s.Valid {
		return nil, nil
	}
	return s.String, nil
}

func (s *String) MarshalJSON() ([]byte, error) {
	if s.String == "" || !s.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(s.String)
}

func (s *String) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	s.String, s.Valid = str, str != ""
	return nil
}

func (s *String) IsEmpty() bool {
	return s.String == "" || !s.Valid
}
