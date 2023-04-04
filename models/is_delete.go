package models

import (
	"github.com/go-openapi/strfmt"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/soft_delete"
)

type IsDelete soft_delete.DeletedAt

func (t IsDelete) Validate(formats strfmt.Registry) error {
	return nil
}

func (IsDelete) QueryClauses(f *schema.Field) []clause.Interface {
	return []clause.Interface{soft_delete.SoftDeleteQueryClause{Field: f}}
}

func (IsDelete) DeleteClauses(f *schema.Field) []clause.Interface {
	settings := schema.ParseTagSetting(f.TagSettings["SOFTDELETE"], ",")
	softDeleteClause := soft_delete.SoftDeleteDeleteClause{
		Field:    f,
		Flag:     settings["FLAG"] != "",
		TimeType: schema.UnixSecond,
	}

	// flag is much more priority
	if !softDeleteClause.Flag {
		if settings["NANO"] != "" {
			softDeleteClause.TimeType = schema.UnixNanosecond
		} else if settings["MILLI"] != "" {
			softDeleteClause.TimeType = schema.UnixMillisecond
		} else {
			softDeleteClause.TimeType = schema.UnixSecond
		}
	}

	if v := settings["DELETEDATFIELD"]; v != "" { // DeletedAtField
		softDeleteClause.DeleteAtField = f.Schema.LookUpField(v)
	}

	return []clause.Interface{softDeleteClause}
}

func (IsDelete) UpdateClauses(f *schema.Field) []clause.Interface {
	return []clause.Interface{soft_delete.SoftDeleteUpdateClause{Field: f}}
}
