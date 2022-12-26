// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	"time"
)

type TireDiscountValueTypeEnum string

const (
	TireDiscountValueTypePercent    TireDiscountValueTypeEnum = "Percent"
	TireDiscountValueTypeFixedCents                           = "FixedCents"
)

type TireSeasonalityEnum string

const (
	TireSeasonalitySummer     TireSeasonalityEnum = "Summer"
	TireSeasonalityWinter                         = "Winter"
	TireSeasonalityAllSeasons                     = "AllSeasons"
)

type TireSizeFormatEnum string

const (
	TireSizeFormatConventional TireSizeFormatEnum = "Conventional"
	TireSizeFormatOffRoadSAE                      = "OffRoadSAE"
)

type Tire struct {
	BinLocation           *string                   `json:"binLocation"`
	BrandId               *string                   `json:"brandId"`
	CategoryId            *string                   `json:"categoryId"`
	CompanyId             string                    `gorm:"not null" json:"companyId"`
	CreatedDate           time.Time                 `gorm:"column:createdDate;not null" json:"createdDate"`
	DiscountCents         int64                     `gorm:"not null" json:"discountCents"`
	DiscountPercent       float64                   `gorm:"not null" json:"discountPercent"`
	DiscountValueType     TireDiscountValueTypeEnum `gorm:"not null" json:"discountValueType"`
	FederalExciseTaxCents int64                     `gorm:"not null" json:"federalExciseTaxCents"`
	ID                    string                    `gorm:"primaryKey;not null" json:"id"`
	LocationId            string                    `gorm:"not null" json:"locationId"`
	Meta                  *Meta                     `gorm:"embedded;column:meta;not null" json:"meta,omitempty"`         // the metadata about the most recent change to the row
	Metadata              any                       `gorm:"embedded;column:metadata;not null" json:"metadata,omitempty"` // metadata reserved for customers to control
	Model                 *string                   `json:"model"`
	Name                  string                    `gorm:"not null" json:"name"`
	Note                  *string                   `json:"note"`
	OrderId               string                    `gorm:"not null" json:"orderId"`
	Ordinal               float64                   `gorm:"not null" json:"ordinal"`
	PartNumber            *string                   `json:"partNumber"`
	PricingMatrixDate     *time.Time                `json:"pricingMatrixDate"` // datetime when pricingMatrixId was set, for determining if matrix has been changed api_schema(calculated)
	PricingMatrixId       *string                   `json:"pricingMatrixId"`
	Quantity              int64                     `gorm:"not null" json:"quantity"`
	ReduceInventoryCount  bool                      `gorm:"not null" json:"reduceInventoryCount"`
	RetailCostCents       int64                     `gorm:"not null" json:"retailCostCents"`
	Seasonality           *TireSeasonalityEnum      `json:"seasonality"`
	ServiceId             string                    `gorm:"not null" json:"serviceId"`
	ShowNote              bool                      `gorm:"not null" json:"showNote"`
	ShowPartNumber        bool                      `gorm:"not null" json:"showPartNumber"`
	ShowPriceAndQuantity  bool                      `gorm:"not null" json:"showPriceAndQuantity"`
	Size                  *string                   `json:"size"`
	SizeFormat            TireSizeFormatEnum        `gorm:"not null" json:"sizeFormat"`
	SourceItemId          *string                   `json:"sourceItemId"`
	Taxable               bool                      `gorm:"not null" json:"taxable"`
	UpdatedDate           *time.Time                `gorm:"column:updatedDate" json:"updatedDate"`
	VendorId              *string                   `json:"vendorId"`
	WasteTireFee          int64                     `gorm:"not null" json:"wasteTireFee"`
	WholesaleCostCents    *int64                    `json:"wholesaleCostCents"`
}

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *Tire) TableName() string {
	return "tire"
}

func (m *Tire) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}
