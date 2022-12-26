// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	mapstructure "github.com/mitchellh/mapstructure"
	"time"
)

type PricingMatrixRange struct {
	CompanyId       string     `gorm:"not null" json:"companyId"`
	CreatedDate     time.Time  `gorm:"column:createdDate;not null" json:"createdDate"`
	ID              string     `gorm:"primaryKey;not null" json:"id"`
	EndCents        *int64     `json:"endCents"`
	MarginPercent   float64    `gorm:"not null" json:"marginPercent"`
	MarkupPercent   float64    `gorm:"not null" json:"markupPercent"`
	Meta            *Meta      `gorm:"type:json;embedded;column:meta;not null" json:"meta,omitempty"` // the metadata about the most recent change to the row
	Metadata        any        `gorm:"type:json" json:"metadata,omitempty"`                           // metadata reserved for customers to control
	PricingMatrixId string     `gorm:"not null" json:"pricingMatrixId"`
	StartCents      int64      `gorm:"not null" json:"startCents"`
	UpdatedDate     *time.Time `gorm:"column:updatedDate" json:"updatedDate"`
}

var _ Model = (*PricingMatrixRange)(nil)

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *PricingMatrixRange) TableName() string {
	return "pricing_matrix_range"
}

func (m *PricingMatrixRange) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewPricingMatrixRange returns a new model instance from a json key/value map
func NewPricingMatrixRange(kv map[string]any) (*PricingMatrixRange, error) {
	var result PricingMatrixRange
	err := mapstructure.Decode(kv, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
