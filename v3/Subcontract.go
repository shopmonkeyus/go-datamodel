// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	mapstructure "github.com/mitchellh/mapstructure"
	"time"
)

type SubcontractDiscountValueTypeEnum string

const (
	SubcontractDiscountValueTypePercent    SubcontractDiscountValueTypeEnum = "Percent"
	SubcontractDiscountValueTypeFixedCents                                  = "FixedCents"
)

type Subcontract struct {
	ID                string                           `gorm:"primaryKey;not null" json:"id"`
	CategoryId        *string                          `json:"categoryId"`
	CompanyId         string                           `gorm:"not null" json:"companyId"`
	CostCents         int64                            `gorm:"not null" json:"costCents"`
	CreatedDate       time.Time                        `gorm:"column:createdDate;not null" json:"createdDate"`
	DiscountCents     int64                            `gorm:"not null" json:"discountCents"`
	DiscountPercent   float64                          `gorm:"not null" json:"discountPercent"`
	DiscountValueType SubcontractDiscountValueTypeEnum `gorm:"not null" json:"discountValueType"`
	LocationId        string                           `gorm:"not null" json:"locationId"`
	Meta              *Meta                            `gorm:"type:json;embedded;column:meta;not null" json:"meta,omitempty"` // the metadata about the most recent change to the row
	Metadata          any                              `gorm:"type:json" json:"metadata,omitempty"`                           // metadata reserved for customers to control
	Name              string                           `gorm:"not null" json:"name"`
	Note              *string                          `json:"note"`
	OrderId           string                           `gorm:"not null" json:"orderId"`
	Ordinal           float64                          `gorm:"not null" json:"ordinal"`
	RetailCostCents   int64                            `gorm:"not null" json:"retailCostCents"`
	ServiceId         string                           `gorm:"not null" json:"serviceId"`
	ShowNote          bool                             `gorm:"not null" json:"showNote"`
	SourceItemId      *string                          `json:"sourceItemId"`
	Taxable           bool                             `gorm:"not null" json:"taxable"`
	UpdatedDate       *time.Time                       `gorm:"column:updatedDate" json:"updatedDate"`
	VendorId          *string                          `json:"vendorId"`
}

var _ Model = (*Subcontract)(nil)

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *Subcontract) TableName() string {
	return "subcontract"
}

func (m *Subcontract) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewSubcontract returns a new model instance from a json key/value map
func NewSubcontract(kv map[string]any) (*Subcontract, error) {
	var result Subcontract
	err := mapstructure.Decode(kv, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
