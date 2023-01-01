// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	datatypes "github.com/shopmonkeyus/go-datamodel/datatypes"
)

type SubcontractDiscountValueTypeEnum string

const (
	SubcontractDiscountValueTypePercent    SubcontractDiscountValueTypeEnum = "Percent"
	SubcontractDiscountValueTypeFixedCents SubcontractDiscountValueTypeEnum = "FixedCents"
)

type Subcontract struct {
	ID          string              `gorm:"primaryKey;not null;column:id" json:"id"`
	CreatedDate datatypes.DateTime  `gorm:"column:createdDate;not null;column:createdDate" json:"createdDate"`
	UpdatedDate *datatypes.DateTime `gorm:"column:updatedDate;column:updatedDate" json:"updatedDate"`
	Meta        datatypes.Meta      `gorm:"column:meta;not null;column:meta" json:"meta,omitempty"`    // the metadata about the most recent change to the row
	Metadata    *datatypes.JSON     `gorm:"column:metadata;column:metadata" json:"metadata,omitempty"` // metadata reserved for customers to control
	CompanyID   string              `gorm:"not null;column:companyId" json:"companyId"`
	LocationID  string              `gorm:"not null;column:locationId" json:"locationId"`

	CategoryID        *string                          `gorm:"column:categoryId" json:"categoryId"`
	CostCents         int64                            `gorm:"not null;column:costCents" json:"costCents"`
	DiscountCents     int64                            `gorm:"not null;column:discountCents" json:"discountCents"`
	DiscountPercent   float64                          `gorm:"not null;column:discountPercent" json:"discountPercent"`
	DiscountValueType SubcontractDiscountValueTypeEnum `gorm:"not null;column:discountValueType" json:"discountValueType"`
	Name              string                           `gorm:"not null;column:name" json:"name"`
	Note              *string                          `gorm:"column:note" json:"note"`
	OrderID           string                           `gorm:"not null;column:orderId" json:"orderId"`
	Ordinal           float64                          `gorm:"not null;column:ordinal" json:"ordinal"`
	RetailCostCents   int64                            `gorm:"not null;column:retailCostCents" json:"retailCostCents"`
	ServiceID         string                           `gorm:"not null;column:serviceId" json:"serviceId"`
	ShowNote          bool                             `gorm:"not null;column:showNote" json:"showNote"`
	SourceItemID      *string                          `gorm:"column:sourceItemId" json:"sourceItemId"`
	Taxable           bool                             `gorm:"not null;column:taxable" json:"taxable"`
	VendorID          *string                          `gorm:"column:vendorId" json:"vendorId"`
}

var _ Model = (*Subcontract)(nil)

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *Subcontract) TableName() string {
	return "subcontract"
}

// String returns a string representation as JSON for this model
func (m *Subcontract) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewSubcontract returns a new model instance from an encoded buffer
func NewSubcontract(buf []byte) (*Subcontract, error) {
	var result Subcontract
	err := json.Unmarshal(buf, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// NewSubcontractFromChangeEvent returns a new model instance from an encoded buffer as change event
func NewSubcontractFromChangeEvent(buf []byte, gzip bool) (*datatypes.ChangeEvent[Subcontract], error) {
	var result datatypes.ChangeEvent[Subcontract]
	var decompressed = buf
	if gzip {
		dec, err := datatypes.Gunzip(buf)
		if err != nil {
			return nil, err
		}
		decompressed = dec
	}
	err := json.Unmarshal(decompressed, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
