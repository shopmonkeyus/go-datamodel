// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	codec "github.com/hashicorp/go-msgpack/v2/codec"
	datatypes "github.com/shopmonkeyus/go-datamodel/datatypes"
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
	ID          string          `gorm:"primaryKey;not null;column:id" json:"id"`
	CreatedDate time.Time       `gorm:"column:createdDate;not null;column:createdDate" json:"createdDate"`
	UpdatedDate *time.Time      `gorm:"column:updatedDate;column:updatedDate" json:"updatedDate"`
	Meta        datatypes.Meta  `gorm:"column:meta;not null;column:meta" json:"meta,omitempty"`    // the metadata about the most recent change to the row
	Metadata    *datatypes.JSON `gorm:"column:metadata;column:metadata" json:"metadata,omitempty"` // metadata reserved for customers to control
	CompanyID   string          `gorm:"not null;column:companyId" json:"companyId"`
	LocationID  string          `gorm:"not null;column:locationId" json:"locationId"`

	BinLocation           *string                   `gorm:"column:binLocation" json:"binLocation"`
	BrandID               *string                   `gorm:"column:brandId" json:"brandId"`
	CategoryID            *string                   `gorm:"column:categoryId" json:"categoryId"`
	DiscountCents         int64                     `gorm:"not null;column:discountCents" json:"discountCents"`
	DiscountPercent       float64                   `gorm:"not null;column:discountPercent" json:"discountPercent"`
	DiscountValueType     TireDiscountValueTypeEnum `gorm:"not null;column:discountValueType" json:"discountValueType"`
	FederalExciseTaxCents int64                     `gorm:"not null;column:federalExciseTaxCents" json:"federalExciseTaxCents"`
	Model                 *string                   `gorm:"column:model" json:"model"`
	Name                  string                    `gorm:"not null;column:name" json:"name"`
	Note                  *string                   `gorm:"column:note" json:"note"`
	OrderID               string                    `gorm:"not null;column:orderId" json:"orderId"`
	Ordinal               float64                   `gorm:"not null;column:ordinal" json:"ordinal"`
	PartNumber            *string                   `gorm:"column:partNumber" json:"partNumber"`
	PricingMatrixDate     *time.Time                `gorm:"column:pricingMatrixDate" json:"pricingMatrixDate"` // datetime when pricingMatrixId was set, for determining if matrix has been changed api_schema(calculated)
	PricingMatrixID       *string                   `gorm:"column:pricingMatrixId" json:"pricingMatrixId"`
	Quantity              int64                     `gorm:"not null;column:quantity" json:"quantity"`
	ReduceInventoryCount  bool                      `gorm:"not null;column:reduceInventoryCount" json:"reduceInventoryCount"`
	RetailCostCents       int64                     `gorm:"not null;column:retailCostCents" json:"retailCostCents"`
	Seasonality           *TireSeasonalityEnum      `gorm:"column:seasonality" json:"seasonality"`
	ServiceID             string                    `gorm:"not null;column:serviceId" json:"serviceId"`
	ShowNote              bool                      `gorm:"not null;column:showNote" json:"showNote"`
	ShowPartNumber        bool                      `gorm:"not null;column:showPartNumber" json:"showPartNumber"`
	ShowPriceAndQuantity  bool                      `gorm:"not null;column:showPriceAndQuantity" json:"showPriceAndQuantity"`
	Size                  *string                   `gorm:"column:size" json:"size"`
	SizeFormat            TireSizeFormatEnum        `gorm:"not null;column:sizeFormat" json:"sizeFormat"`
	SourceItemID          *string                   `gorm:"column:sourceItemId" json:"sourceItemId"`
	Taxable               bool                      `gorm:"not null;column:taxable" json:"taxable"`
	VendorID              *string                   `gorm:"column:vendorId" json:"vendorId"`
	WasteTireFee          int64                     `gorm:"not null;column:wasteTireFee" json:"wasteTireFee"`
	WholesaleCostCents    *int64                    `gorm:"column:wholesaleCostCents" json:"wholesaleCostCents"`
}

var _ Model = (*Tire)(nil)

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *Tire) TableName() string {
	return "tire"
}

func (m *Tire) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewTire returns a new model instance from an encoded buffer
func NewTire(buf []byte, enctype EncodingType) (*Tire, error) {
	var result Tire
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
