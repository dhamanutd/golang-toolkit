package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-openapi/strfmt"
)

type DateTime strfmt.DateTime

// NewDateTime is a representation of zero value for DateTime type
func NewDateTime() DateTime {
	return DateTime(time.Unix(0, 0).UTC())
}

func (t DateTime) String() string {
	return time.Time(t).Format("2006-01-02 15:04:05.00")
}

func (t DateTime) IsZero() bool {
	parseTime := strfmt.DateTime(t)
	return parseTime.IsZero()
}

// New DateTime with time now
func NewDateTimeNow() DateTime {
	t := time.Now().UTC()
	return DateTime(t)
}

// MarshalText implements the text marshaller interface
func (t DateTime) MarshalText() ([]byte, error) {
	return []byte(t.String()), nil
}

// UnmarshalText implements the text unmarshaller interface
func (t *DateTime) UnmarshalText(text []byte) error {
	tt, err := ParseDateTime(string(text))
	if err != nil {
		return err
	}
	*t = tt
	return nil

}

// Scan scans a DateTime value from database driver type.
func (t *DateTime) Scan(raw interface{}) error {
	switch v := raw.(type) {
	case []byte:
		return t.UnmarshalText(v)
	case string:
		return t.UnmarshalText([]byte(v))
	case time.Time:
		*t = DateTime(v)
	case nil:
		*t = DateTime{}
	default:
		return fmt.Errorf("cannot sql.Scan() strfmt.DateTime from: %#v", v)
	}

	return nil
}

// Value converts DateTime to a primitive value ready to written to a database.
func (t DateTime) Value() (driver.Value, error) {
	return driver.Value(time.Time(t).Format("2006-01-02 15:04:05.00")), nil
}

// MarshalJSON returns the DateTime as JSON
func (t DateTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(strfmt.NormalizeTimeForMarshal(time.Time(t)).Format(strfmt.MarshalFormat))
}

// UnmarshalJSON sets the DateTime from JSON
func (t *DateTime) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}

	var tstr string
	if err := json.Unmarshal(data, &tstr); err != nil {
		return err
	}
	tt, err := ParseDateTime(tstr)
	if err != nil {
		return err
	}
	*t = tt
	return nil
}

// DeepCopyInto copies the receiver and writes its value into out.
func (t *DateTime) DeepCopyInto(out *DateTime) {
	*out = *t
}

// DeepCopy copies the receiver into a new DateTime.
func (t *DateTime) DeepCopy() *DateTime {
	if t == nil {
		return nil
	}
	out := new(DateTime)
	t.DeepCopyInto(out)
	return out
}

// Equal checks if two DateTime instances are equal using time.Time's Equal method
func (t DateTime) Equal(t2 DateTime) bool {
	return time.Time(t).Equal(time.Time(t2))
}

func (t DateTime) Before(t2 DateTime) bool {
	return time.Time(t).Before(time.Time(t2))
}

func ParseDateTime(data string) (DateTime, error) {
	if data == "" {
		return NewDateTime(), nil
	}
	var lastError error
	for _, layout := range strfmt.DateTimeFormats {
		dd, err := time.Parse(layout, data)
		if err != nil {
			lastError = err
			continue
		}
		return DateTime(dd), nil
	}
	return DateTime{}, lastError
}

func (t DateTime) Validate(formats strfmt.Registry) error {
	return nil
}
