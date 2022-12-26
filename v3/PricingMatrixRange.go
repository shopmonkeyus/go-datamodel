// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	mapstructure "github.com/mitchellh/mapstructure"
	"time"
)

type PricingMatrixRange struct {
	ID          string     `gorm:"primaryKey;not null;column:id" json:"id"`
	CreatedDate time.Time  `gorm:"column:createdDate;not null;column:createdDate" json:"createdDate"`
	UpdatedDate *time.Time `gorm:"column:updatedDate;column:updatedDate" json:"updatedDate"`
	Meta        *Meta      `gorm:"type:json;embedded;serializer:json;column:meta;not null;column:meta" json:"meta,omitempty"` // the metadata about the most recent change to the row
	Metadata    any        `gorm:"type:json;serializer:json;column:metadata" json:"metadata,omitempty"`                       // metadata reserved for customers to control
	CompanyID   string     `gorm:"not null;column:companyId" json:"companyId"`

	EndCents        *int64  `gorm:"column:endCents" json:"endCents"`
	MarginPercent   float64 `gorm:"not null;column:marginPercent" json:"marginPercent"`
	MarkupPercent   float64 `gorm:"not null;column:markupPercent" json:"markupPercent"`
	PricingMatrixID string  `gorm:"not null;column:pricingMatrixId" json:"pricingMatrixId"`
	StartCents      int64   `gorm:"not null;column:startCents" json:"startCents"`
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
