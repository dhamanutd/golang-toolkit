package models

import (
	"github.com/go-openapi/strfmt"
	"gorm.io/gorm"
)

type DeletedAt struct {
	gorm.DeletedAt
}

func (DeletedAt) Validate(strfmt.Registry) error {
	return nil
}
