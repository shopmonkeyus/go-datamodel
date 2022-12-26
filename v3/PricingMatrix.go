// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	"time"
)

type PricingMatrix struct {
	CompanyId   string     `gorm:"not null" json:"companyId"`
	CreatedDate time.Time  `gorm:"column:createdDate;not null" json:"createdDate"`
	Default     bool       `gorm:"not null" json:"default"`
	ID          string     `gorm:"primaryKey;not null" json:"id"`
	LocationId  string     `gorm:"not null" json:"locationId"`
	Meta        *Meta      `gorm:"embedded;column:meta;not null" json:"meta,omitempty"`         // the metadata about the most recent change to the row
	Metadata    any        `gorm:"embedded;column:metadata;not null" json:"metadata,omitempty"` // metadata reserved for customers to control
	Name        string     `gorm:"not null" json:"name"`
	UpdatedDate *time.Time `gorm:"column:updatedDate" json:"updatedDate"`
}

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *PricingMatrix) TableName() string {
	return "pricing_matrix"
}

func (m *PricingMatrix) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}
