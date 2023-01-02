// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	datatypes "github.com/shopmonkeyus/go-datamodel/datatypes"
)

type CannedServicePartDiscountValueTypeEnum string

const (
	CannedServicePartDiscountValueTypePercent    CannedServicePartDiscountValueTypeEnum = "Percent"
	CannedServicePartDiscountValueTypeFixedCents CannedServicePartDiscountValueTypeEnum = "FixedCents"
)

type CannedServicePart struct {
	ID          string              `gorm:"primaryKey;not null;column:id" json:"id"`
	CreatedDate datatypes.DateTime  `gorm:"column:createdDate;not null;column:createdDate" json:"createdDate"`
	UpdatedDate *datatypes.DateTime `gorm:"column:updatedDate;column:updatedDate" json:"updatedDate"`
	Meta        *datatypes.JSON     `gorm:"column:meta;not null;column:meta" json:"meta,omitempty"`    // the metadata about the most recent change to the row
	Metadata    *datatypes.JSON     `gorm:"column:metadata;column:metadata" json:"metadata,omitempty"` // metadata reserved for customers to control
	CompanyID   string              `gorm:"not null;column:companyId" json:"companyId"`
	LocationID  string              `gorm:"not null;column:locationId" json:"locationId"`

	BinLocation          string                                 `gorm:"not null;column:binLocation" json:"binLocation"`
	CannedServiceID      string                                 `gorm:"not null;column:cannedServiceId" json:"cannedServiceId"`
	CategoryID           *string                                `gorm:"column:categoryId" json:"categoryId"`
	DiscountCents        int64                                  `gorm:"not null;column:discountCents" json:"discountCents"`
	DiscountPercent      float64                                `gorm:"not null;column:discountPercent" json:"discountPercent"`
	DiscountValueType    CannedServicePartDiscountValueTypeEnum `gorm:"not null;column:discountValueType" json:"discountValueType"`
	InventoryPartID      *string                                `gorm:"column:inventoryPartId" json:"inventoryPartId"`
	Name                 string                                 `gorm:"not null;column:name" json:"name"`
	Note                 string                                 `gorm:"not null;column:note" json:"note"`
	Ordinal              float64                                `gorm:"not null;column:ordinal" json:"ordinal"`
	PartNumber           string                                 `gorm:"not null;column:partNumber" json:"partNumber"`
	PricingMatrixDate    *datatypes.DateTime                    `gorm:"column:pricingMatrixDate" json:"pricingMatrixDate"` // datetime when pricingMatrixId was set, for determining if matrix has been changed
	PricingMatrixID      *string                                `gorm:"column:pricingMatrixId" json:"pricingMatrixId"`
	Quantity             int64                                  `gorm:"not null;column:quantity" json:"quantity"`
	ReduceInventoryCount bool                                   `gorm:"not null;column:reduceInventoryCount" json:"reduceInventoryCount"`
	RetailCostCents      int64                                  `gorm:"not null;column:retailCostCents" json:"retailCostCents"`
	ShowCostAndQuantity  bool                                   `gorm:"not null;column:showCostAndQuantity" json:"showCostAndQuantity"`
	ShowNote             bool                                   `gorm:"not null;column:showNote" json:"showNote"`
	ShowPartNumber       bool                                   `gorm:"not null;column:showPartNumber" json:"showPartNumber"`
	SourceItemID         *string                                `gorm:"column:sourceItemId" json:"sourceItemId"`
	Taxable              bool                                   `gorm:"not null;column:taxable" json:"taxable"`
	VendorID             *string                                `gorm:"column:vendorId" json:"vendorId"`
	WholesaleCostCents   *int64                                 `gorm:"column:wholesaleCostCents" json:"wholesaleCostCents"`
}

var _ Model = (*CannedServicePart)(nil)

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *CannedServicePart) TableName() string {
	return "canned_service_part"
}

// PrimaryKeys returns an array of the primary keys for this model
func (m *CannedServicePart) PrimaryKeys() []string {
	return []string{"id"}
}

// String returns a string representation as JSON for this model
func (m *CannedServicePart) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewCannedServicePart returns a new model instance from an encoded buffer
func NewCannedServicePart(buf []byte) (*CannedServicePart, error) {
	var result CannedServicePart
	err := json.Unmarshal(buf, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// NewCannedServicePartFromChangeEvent returns a new model instance from an encoded buffer as change event
func NewCannedServicePartFromChangeEvent(buf []byte, gzip bool) (*datatypes.ChangeEvent[CannedServicePart], error) {
	var result datatypes.ChangeEvent[CannedServicePart]
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
