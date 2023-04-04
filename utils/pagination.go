package utils

import (
	"math"

	"gorm.io/gorm"
)

type Paginate struct {
	Limit int64
	Page  int64
}

type Metadata struct {
	Page      int64 `json:"page"`
	PerPage   int64 `json:"per_page"`
	TotalPage int64 `json:"total_page"`
	TotalRow  int64 `json:"total_row"`
}

func NewPaginate(limit int64, page int64) *Paginate {
	return &Paginate{Limit: limit, Page: page}
}

func (p *Paginate) Scope(db *gorm.DB) *gorm.DB {
	offset := (p.Page - 1) * p.Limit

	return db.Offset(int(offset)).
		Limit(int(p.Limit))
}

func (p *Paginate) WithCount(db *gorm.DB) (result int64) {
	db.Count(&result)
	return result
}

func (p *Paginate) WithMetadata(db *gorm.DB) Metadata {
	count := p.WithCount(db)

	return Metadata{
		Page:      p.Page,
		PerPage:   p.Limit,
		TotalRow:  count,
		TotalPage: int64(math.Ceil(float64(count) / float64(p.Limit))),
	}

}
