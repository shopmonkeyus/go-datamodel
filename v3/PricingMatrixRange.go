// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	"time"
)

type PricingMatrixRange struct {
	ID              string     `gorm:"primaryKey;not null" json:"id"`
	CompanyId       string     `gorm:"not null" json:"companyId"`
	CreatedDate     time.Time  `gorm:"column:createdDate;not null" json:"createdDate"`
	EndCents        *int64     `json:"endCents"`
	MarginPercent   float64    `gorm:"not null" json:"marginPercent"`
	MarkupPercent   float64    `gorm:"not null" json:"markupPercent"`
	Meta            *Meta      `gorm:"embedded;column:meta;not null" json:"meta,omitempty"`         // the metadata about the most recent change to the row
	Metadata        any        `gorm:"embedded;column:metadata;not null" json:"metadata,omitempty"` // metadata reserved for customers to control
	PricingMatrixId string     `gorm:"not null" json:"pricingMatrixId"`
	StartCents      int64      `gorm:"not null" json:"startCents"`
	UpdatedDate     *time.Time `gorm:"column:updatedDate" json:"updatedDate"`
}

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *PricingMatrixRange) TableName() string {
	return "pricing_matrix_range"
}

func (m *PricingMatrixRange) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}
