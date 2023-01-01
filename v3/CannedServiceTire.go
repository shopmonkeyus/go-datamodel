// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	datatypes "github.com/shopmonkeyus/go-datamodel/datatypes"
)

type CannedServiceTireDiscountValueTypeEnum string

const (
	CannedServiceTireDiscountValueTypePercent    CannedServiceTireDiscountValueTypeEnum = "Percent"
	CannedServiceTireDiscountValueTypeFixedCents CannedServiceTireDiscountValueTypeEnum = "FixedCents"
)

type CannedServiceTireSeasonalityEnum string

const (
	CannedServiceTireSeasonalitySummer     CannedServiceTireSeasonalityEnum = "Summer"
	CannedServiceTireSeasonalityWinter     CannedServiceTireSeasonalityEnum = "Winter"
	CannedServiceTireSeasonalityAllSeasons CannedServiceTireSeasonalityEnum = "AllSeasons"
	CannedServiceTireSeasonality           CannedServiceTireSeasonalityEnum = ""
)

type CannedServiceTireSizeFormatEnum string

const (
	CannedServiceTireSizeFormatConventional CannedServiceTireSizeFormatEnum = "Conventional"
	CannedServiceTireSizeFormatOffRoadSAE   CannedServiceTireSizeFormatEnum = "OffRoadSAE"
)

type CannedServiceTire struct {
	ID          string              `gorm:"primaryKey;not null;column:id" json:"id"`
	CreatedDate datatypes.DateTime  `gorm:"column:createdDate;not null;column:createdDate" json:"createdDate"`
	UpdatedDate *datatypes.DateTime `gorm:"column:updatedDate;column:updatedDate" json:"updatedDate"`
	Meta        datatypes.Meta      `gorm:"column:meta;not null;column:meta" json:"meta,omitempty"`    // the metadata about the most recent change to the row
	Metadata    *datatypes.JSON     `gorm:"column:metadata;column:metadata" json:"metadata,omitempty"` // metadata reserved for customers to control
	CompanyID   string              `gorm:"not null;column:companyId" json:"companyId"`
	LocationID  string              `gorm:"not null;column:locationId" json:"locationId"`

	BinLocation           *string                                `gorm:"column:binLocation" json:"binLocation"`
	BrandID               *string                                `gorm:"column:brandId" json:"brandId"`
	CannedServiceID       string                                 `gorm:"not null;column:cannedServiceId" json:"cannedServiceId"`
	CategoryID            *string                                `gorm:"column:categoryId" json:"categoryId"`
	DiscountCents         int64                                  `gorm:"not null;column:discountCents" json:"discountCents"`
	DiscountPercent       float64                                `gorm:"not null;column:discountPercent" json:"discountPercent"`
	DiscountValueType     CannedServiceTireDiscountValueTypeEnum `gorm:"not null;column:discountValueType" json:"discountValueType"`
	FederalExciseTaxCents int64                                  `gorm:"not null;column:federalExciseTaxCents" json:"federalExciseTaxCents"`
	Model                 *string                                `gorm:"column:model" json:"model"`
	Name                  string                                 `gorm:"not null;column:name" json:"name"`
	Note                  *string                                `gorm:"column:note" json:"note"`
	Ordinal               float64                                `gorm:"not null;column:ordinal" json:"ordinal"`
	PartNumber            *string                                `gorm:"column:partNumber" json:"partNumber"`
	PricingMatrixDate     *datatypes.DateTime                    `gorm:"column:pricingMatrixDate" json:"pricingMatrixDate"` // datetime when pricingMatrixId was set, for determining if matrix has been changed api_schema(calculated)
	PricingMatrixID       *string                                `gorm:"column:pricingMatrixId" json:"pricingMatrixId"`
	Quantity              int64                                  `gorm:"not null;column:quantity" json:"quantity"`
	ReduceInventoryCount  bool                                   `gorm:"not null;column:reduceInventoryCount" json:"reduceInventoryCount"`
	RetailCostCents       int64                                  `gorm:"not null;column:retailCostCents" json:"retailCostCents"`
	Seasonality           *CannedServiceTireSeasonalityEnum      `gorm:"column:seasonality" json:"seasonality"`
	ShowNote              bool                                   `gorm:"not null;column:showNote" json:"showNote"`
	ShowPartNumber        bool                                   `gorm:"not null;column:showPartNumber" json:"showPartNumber"`
	ShowPriceAndQuantity  bool                                   `gorm:"not null;column:showPriceAndQuantity" json:"showPriceAndQuantity"`
	Size                  *string                                `gorm:"column:size" json:"size"`
	SizeFormat            CannedServiceTireSizeFormatEnum        `gorm:"not null;column:sizeFormat" json:"sizeFormat"`
	SourceItemID          *string                                `gorm:"column:sourceItemId" json:"sourceItemId"`
	Taxable               bool                                   `gorm:"not null;column:taxable" json:"taxable"`
	VendorID              *string                                `gorm:"column:vendorId" json:"vendorId"`
	WasteTireFee          int64                                  `gorm:"not null;column:wasteTireFee" json:"wasteTireFee"`
	WholesaleCostCents    *int64                                 `gorm:"column:wholesaleCostCents" json:"wholesaleCostCents"`
}

var _ Model = (*CannedServiceTire)(nil)

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *CannedServiceTire) TableName() string {
	return "canned_service_tire"
}

// String returns a string representation as JSON for this model
func (m *CannedServiceTire) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewCannedServiceTire returns a new model instance from an encoded buffer
func NewCannedServiceTire(buf []byte) (*CannedServiceTire, error) {
	var result CannedServiceTire
	err := json.Unmarshal(buf, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// NewCannedServiceTireFromChangeEvent returns a new model instance from an encoded buffer as change event
func NewCannedServiceTireFromChangeEvent(buf []byte, gzip bool) (*datatypes.ChangeEvent[CannedServiceTire], error) {
	var result datatypes.ChangeEvent[CannedServiceTire]
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
