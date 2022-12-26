// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	"time"
)

type CannedServiceSubcontractDiscountValueTypeEnum string

const (
	CannedServiceSubcontractDiscountValueTypePercent    CannedServiceSubcontractDiscountValueTypeEnum = "Percent"
	CannedServiceSubcontractDiscountValueTypeFixedCents                                               = "FixedCents"
)

type CannedServiceSubcontract struct {
	ID          string     `gorm:"primaryKey;not null;column:id" json:"id"`
	CreatedDate time.Time  `gorm:"column:createdDate;not null;column:createdDate" json:"createdDate"`
	UpdatedDate *time.Time `gorm:"column:updatedDate;column:updatedDate" json:"updatedDate"`
	Meta        *Meta      `gorm:"type:json;embedded;serializer:json;column:meta;not null;column:meta" json:"meta,omitempty"` // the metadata about the most recent change to the row
	Metadata    any        `gorm:"type:json;serializer:json;column:metadata" json:"metadata,omitempty"`                       // metadata reserved for customers to control
	CompanyID   string     `gorm:"not null;column:companyId" json:"companyId"`
	LocationID  string     `gorm:"not null;column:locationId" json:"locationId"`

	CannedServiceID   string                                        `gorm:"not null;column:cannedServiceId" json:"cannedServiceId"`
	CategoryID        *string                                       `gorm:"column:categoryId" json:"categoryId"`
	CostCents         int64                                         `gorm:"not null;column:costCents" json:"costCents"`
	DiscountCents     int64                                         `gorm:"not null;column:discountCents" json:"discountCents"`
	DiscountPercent   float64                                       `gorm:"not null;column:discountPercent" json:"discountPercent"`
	DiscountValueType CannedServiceSubcontractDiscountValueTypeEnum `gorm:"not null;column:discountValueType" json:"discountValueType"`
	Name              string                                        `gorm:"not null;column:name" json:"name"`
	Note              *string                                       `gorm:"column:note" json:"note"`
	Ordinal           float64                                       `gorm:"not null;column:ordinal" json:"ordinal"`
	RetailCostCents   int64                                         `gorm:"not null;column:retailCostCents" json:"retailCostCents"`
	ShowNote          bool                                          `gorm:"not null;column:showNote" json:"showNote"`
	SourceItemID      *string                                       `gorm:"column:sourceItemId" json:"sourceItemId"`
	Taxable           bool                                          `gorm:"not null;column:taxable" json:"taxable"`
	VendorID          *string                                       `gorm:"column:vendorId" json:"vendorId"`
}

var _ Model = (*CannedServiceSubcontract)(nil)

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *CannedServiceSubcontract) TableName() string {
	return "canned_service_subcontract"
}

func (m *CannedServiceSubcontract) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewCannedServiceSubcontract returns a new model instance from a json key/value map
func NewCannedServiceSubcontract(buf []byte) (*CannedServiceSubcontract, error) {
	var result CannedServiceSubcontract
	err := json.Unmarshal(buf, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
