package models

import (
	"database/sql/driver"
	"strings"

	"github.com/go-openapi/strfmt"
)

type StringArray []string

func (m *StringArray) Validate(formats strfmt.Registry) error {
	return nil
}

func (m *StringArray) Scan(src interface{}) error {
	switch src := src.(type) {
	case []byte:
		*m = strings.Split(string(src), ",")
	case nil:
		*m = nil
		return nil
	}

	return nil
}

func (m *StringArray) Value() (driver.Value, error) {
	if m == nil {
		return nil, nil
	}

	return strings.Join(*m, ","), nil
}
