// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	mapstructure "github.com/mitchellh/mapstructure"
	"time"
)

type CannedServicePartDiscountValueTypeEnum string

const (
	CannedServicePartDiscountValueTypePercent    CannedServicePartDiscountValueTypeEnum = "Percent"
	CannedServicePartDiscountValueTypeFixedCents                                        = "FixedCents"
)

type CannedServicePart struct {
	ID                   string                                 `gorm:"primaryKey;not null" json:"id"`
	BinLocation          string                                 `gorm:"not null" json:"binLocation"`
	CannedServiceId      string                                 `gorm:"not null" json:"cannedServiceId"`
	CategoryId           *string                                `json:"categoryId"`
	CompanyId            string                                 `gorm:"not null" json:"companyId"`
	CreatedDate          time.Time                              `gorm:"column:createdDate;not null" json:"createdDate"`
	DiscountCents        int64                                  `gorm:"not null" json:"discountCents"`
	DiscountPercent      float64                                `gorm:"not null" json:"discountPercent"`
	DiscountValueType    CannedServicePartDiscountValueTypeEnum `gorm:"not null" json:"discountValueType"`
	InventoryPartId      *string                                `json:"inventoryPartId"`
	LocationId           string                                 `gorm:"not null" json:"locationId"`
	Meta                 *Meta                                  `gorm:"type:json;embedded;column:meta;not null" json:"meta,omitempty"` // the metadata about the most recent change to the row
	Metadata             any                                    `gorm:"type:json" json:"metadata,omitempty"`                           // metadata reserved for customers to control
	Name                 string                                 `gorm:"not null" json:"name"`
	Note                 string                                 `gorm:"not null" json:"note"`
	Ordinal              float64                                `gorm:"not null" json:"ordinal"`
	PartNumber           string                                 `gorm:"not null" json:"partNumber"`
	PricingMatrixDate    *time.Time                             `json:"pricingMatrixDate"` // datetime when pricingMatrixId was set, for determining if matrix has been changed
	PricingMatrixId      *string                                `json:"pricingMatrixId"`
	Quantity             int64                                  `gorm:"not null" json:"quantity"`
	ReduceInventoryCount bool                                   `gorm:"not null" json:"reduceInventoryCount"`
	RetailCostCents      int64                                  `gorm:"not null" json:"retailCostCents"`
	ShowCostAndQuantity  bool                                   `gorm:"not null" json:"showCostAndQuantity"`
	ShowNote             bool                                   `gorm:"not null" json:"showNote"`
	ShowPartNumber       bool                                   `gorm:"not null" json:"showPartNumber"`
	SourceItemId         *string                                `json:"sourceItemId"`
	Taxable              bool                                   `gorm:"not null" json:"taxable"`
	UpdatedDate          *time.Time                             `gorm:"column:updatedDate" json:"updatedDate"`
	VendorId             *string                                `json:"vendorId"`
	WholesaleCostCents   *int64                                 `json:"wholesaleCostCents"`
}

var _ Model = (*CannedServicePart)(nil)

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *CannedServicePart) TableName() string {
	return "canned_service_part"
}

func (m *CannedServicePart) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewCannedServicePart returns a new model instance from a json key/value map
func NewCannedServicePart(kv map[string]any) (*CannedServicePart, error) {
	var result CannedServicePart
	err := mapstructure.Decode(kv, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
