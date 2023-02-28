// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	datatypes "github.com/shopmonkeyus/go-datamodel/datatypes"
)

// CannedServiceSubcontract schema
type CannedServiceSubcontractDiscountValueTypeEnum string

const (
	CannedServiceSubcontractDiscountValueTypePercent    CannedServiceSubcontractDiscountValueTypeEnum = "Percent"
	CannedServiceSubcontractDiscountValueTypeFixedCents CannedServiceSubcontractDiscountValueTypeEnum = "FixedCents"
)

type CannedServiceSubcontract struct {
	ID          string              `bson:"_id" gorm:"primaryKey;not null;column:id" json:"id"`
	CreatedDate datatypes.DateTime  `gorm:"column:createdDate;not null;column:createdDate" json:"createdDate"`
	UpdatedDate *datatypes.DateTime `gorm:"column:updatedDate;column:updatedDate" json:"updatedDate"`
	Meta        *datatypes.JSON     `gorm:"column:meta;not null;column:meta" json:"meta,omitempty"` // the metadata about the most recent change to the row
	Metadata    *datatypes.JSON     `gorm:"column:metadata;column:metadata" json:"metadata,omitempty"`
	CompanyID   string              `gorm:"not null;column:companyId" json:"companyId"`
	LocationID  string              `gorm:"not null;column:locationId" json:"locationId"`

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

// PrimaryKeys returns an array of the primary keys for this model
func (m *CannedServiceSubcontract) PrimaryKeys() []string {
	return []string{"id"}
}

// String returns a string representation as JSON for this model
func (m *CannedServiceSubcontract) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewCannedServiceSubcontract returns a new model instance from an encoded buffer
func NewCannedServiceSubcontract(buf []byte) (*CannedServiceSubcontract, error) {
	var result CannedServiceSubcontract
	err := json.Unmarshal(buf, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// NewCannedServiceSubcontractFromChangeEvent returns a new model instance from an encoded buffer as change event
func NewCannedServiceSubcontractFromChangeEvent(buf []byte, gzip bool) (*datatypes.ChangeEvent[CannedServiceSubcontract], error) {
	var result datatypes.ChangeEvent[CannedServiceSubcontract]
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
