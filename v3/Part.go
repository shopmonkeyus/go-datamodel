// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	codec "github.com/hashicorp/go-msgpack/v2/codec"
	"time"
)

type PartDiscountValueTypeEnum string

const (
	PartDiscountValueTypePercent    PartDiscountValueTypeEnum = "Percent"
	PartDiscountValueTypeFixedCents                           = "FixedCents"
)

type Part struct {
	ID          string     `gorm:"primaryKey;not null;column:id" json:"id"`
	CreatedDate time.Time  `gorm:"column:createdDate;not null;column:createdDate" json:"createdDate"`
	UpdatedDate *time.Time `gorm:"column:updatedDate;column:updatedDate" json:"updatedDate"`
	Meta        *Meta      `gorm:"type:json;embedded;serializer:json;column:meta;not null;column:meta" json:"meta,omitempty"` // the metadata about the most recent change to the row
	Metadata    any        `gorm:"type:json;serializer:json;column:metadata" json:"metadata,omitempty"`                       // metadata reserved for customers to control
	CompanyID   string     `gorm:"not null;column:companyId" json:"companyId"`
	LocationID  string     `gorm:"not null;column:locationId" json:"locationId"`

	BinLocation          string                    `gorm:"not null;column:binLocation" json:"binLocation"`
	CategoryID           *string                   `gorm:"column:categoryId" json:"categoryId"`
	DiscountCents        int64                     `gorm:"not null;column:discountCents" json:"discountCents"`
	DiscountPercent      float64                   `gorm:"not null;column:discountPercent" json:"discountPercent"`
	DiscountValueType    PartDiscountValueTypeEnum `gorm:"not null;column:discountValueType" json:"discountValueType"`
	InventoryPartID      *string                   `gorm:"column:inventoryPartId" json:"inventoryPartId"`
	Name                 string                    `gorm:"not null;column:name" json:"name"`
	Note                 string                    `gorm:"not null;column:note" json:"note"`
	OrderID              string                    `gorm:"not null;column:orderId" json:"orderId"`
	Ordinal              float64                   `gorm:"not null;column:ordinal" json:"ordinal"`
	PartNumber           string                    `gorm:"not null;column:partNumber" json:"partNumber"`
	PricingMatrixDate    *time.Time                `gorm:"column:pricingMatrixDate" json:"pricingMatrixDate"` // datetime when pricingMatrixId was set, for determining if matrix has been changed
	PricingMatrixID      *string                   `gorm:"column:pricingMatrixId" json:"pricingMatrixId"`
	Quantity             int64                     `gorm:"not null;column:quantity" json:"quantity"`
	ReduceInventoryCount bool                      `gorm:"not null;column:reduceInventoryCount" json:"reduceInventoryCount"`
	RetailCostCents      int64                     `gorm:"not null;column:retailCostCents" json:"retailCostCents"`
	ServiceID            string                    `gorm:"not null;column:serviceId" json:"serviceId"`
	ShowCostAndQuantity  bool                      `gorm:"not null;column:showCostAndQuantity" json:"showCostAndQuantity"`
	ShowNote             bool                      `gorm:"not null;column:showNote" json:"showNote"`
	ShowPartNumber       bool                      `gorm:"not null;column:showPartNumber" json:"showPartNumber"`
	SourceItemID         *string                   `gorm:"column:sourceItemId" json:"sourceItemId"`
	Taxable              bool                      `gorm:"not null;column:taxable" json:"taxable"`
	VendorID             *string                   `gorm:"column:vendorId" json:"vendorId"`
	WholesaleCostCents   *int64                    `gorm:"column:wholesaleCostCents" json:"wholesaleCostCents"`
}

var _ Model = (*Part)(nil)

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *Part) TableName() string {
	return "part"
}

func (m *Part) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewPart returns a new model instance from an encoded buffer
func NewPart(buf []byte, enctype EncodingType) (*Part, error) {
	var result Part
	var handle codec.Handle
	if enctype == JSONEncoding {
		handle = &jsonHandle
	} else {
		handle = &msgpackHandle
	}
	dec := codec.NewDecoderBytes(buf, handle)
	err := dec.Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
