// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	datatypes "github.com/shopmonkeyus/go-datamodel/datatypes"
)

// InventoryTire schema
type InventoryTireCrossUnitEnum string

const (
	InventoryTireCrossUnitmm InventoryTireCrossUnitEnum = "mm"
	InventoryTireCrossUnitin InventoryTireCrossUnitEnum = "in"
)

type InventoryTireDiscountValueTypeEnum string

const (
	InventoryTireDiscountValueTypePercent    InventoryTireDiscountValueTypeEnum = "Percent"
	InventoryTireDiscountValueTypeFixedCents InventoryTireDiscountValueTypeEnum = "FixedCents"
)

type InventoryTireOuterDiameterUnitEnum string

const (
	InventoryTireOuterDiameterUnitmm InventoryTireOuterDiameterUnitEnum = "mm"
	InventoryTireOuterDiameterUnitin InventoryTireOuterDiameterUnitEnum = "in"
)

type InventoryTireRimUnitEnum string

const (
	InventoryTireRimUnitmm InventoryTireRimUnitEnum = "mm"
	InventoryTireRimUnitin InventoryTireRimUnitEnum = "in"
)

type InventoryTireSeasonalityEnum string

const (
	InventoryTireSeasonalitySummer     InventoryTireSeasonalityEnum = "Summer"
	InventoryTireSeasonalityWinter     InventoryTireSeasonalityEnum = "Winter"
	InventoryTireSeasonalityAllSeasons InventoryTireSeasonalityEnum = "AllSeasons"
)

type InventoryTireSizeFormatEnum string

const (
	InventoryTireSizeFormatConventional InventoryTireSizeFormatEnum = "Conventional"
	InventoryTireSizeFormatOffRoadSAE   InventoryTireSizeFormatEnum = "OffRoadSAE"
)

type InventoryTireTreadDepthUnitEnum string

const (
	InventoryTireTreadDepthUnitmm InventoryTireTreadDepthUnitEnum = "mm"
	InventoryTireTreadDepthUnitin InventoryTireTreadDepthUnitEnum = "in"
)

type InventoryTire struct {
	ID          string              `bson:"_id" gorm:"primaryKey;not null;column:id" json:"id"`
	CreatedDate datatypes.DateTime  `gorm:"column:createdDate;not null;column:createdDate" json:"createdDate"`
	UpdatedDate *datatypes.DateTime `gorm:"column:updatedDate;column:updatedDate" json:"updatedDate"`
	Meta        *datatypes.JSON     `gorm:"column:meta;not null;column:meta" json:"meta,omitempty"`    // the metadata about the most recent change to the row
	Metadata    *datatypes.JSON     `gorm:"column:metadata;column:metadata" json:"metadata,omitempty"` // metadata reserved for customers to control
	CompanyID   string              `gorm:"not null;column:companyId" json:"companyId"`
	LocationID  string              `gorm:"not null;column:locationId" json:"locationId"`

	AspectValue           *float64                           `gorm:"column:aspectValue" json:"aspectValue"`
	BinLocation           *string                            `gorm:"column:binLocation" json:"binLocation"`
	BrandID               *string                            `gorm:"column:brandId" json:"brandId"`
	CategoryID            *string                            `gorm:"column:categoryId" json:"categoryId"`
	CategoryType          *string                            `gorm:"column:categoryType" json:"categoryType"`
	Construction          *string                            `gorm:"column:construction" json:"construction"`
	CrossUnit             InventoryTireCrossUnitEnum         `gorm:"not null;column:crossUnit" json:"crossUnit"`
	CrossValue            *float64                           `gorm:"column:crossValue" json:"crossValue"`
	Deleted               bool                               `gorm:"not null;column:deleted" json:"deleted"`    // if the record has been deleted
	DeletedDate           *datatypes.DateTime                `gorm:"column:deletedDate" json:"deletedDate"`     // the date that the record was deleted or null if not deleted
	DeletedReason         *string                            `gorm:"column:deletedReason" json:"deletedReason"` // the reason that the record was deleted
	DeletedUserID         *string                            `gorm:"column:deletedUserId" json:"deletedUserId"` // the user that deleted the record or null if not deleted
	DiscountCents         int64                              `gorm:"not null;column:discountCents" json:"discountCents"`
	DiscountPercent       float64                            `gorm:"not null;column:discountPercent" json:"discountPercent"`
	DiscountValueType     InventoryTireDiscountValueTypeEnum `gorm:"not null;column:discountValueType" json:"discountValueType"`
	FederalExciseTaxCents int64                              `gorm:"not null;column:federalExciseTaxCents" json:"federalExciseTaxCents"`
	InventoryNumber       *string                            `gorm:"column:inventoryNumber" json:"inventoryNumber"`
	IsRunFlat             bool                               `gorm:"not null;column:isRunFlat" json:"isRunFlat"`
	IsStudded             bool                               `gorm:"not null;column:isStudded" json:"isStudded"`
	LoadIndex             *int64                             `gorm:"column:loadIndex" json:"loadIndex"`
	LoadRange             *string                            `gorm:"column:loadRange" json:"loadRange"`
	MaxTirePressurePSI    *float64                           `gorm:"column:maxTirePressurePSI" json:"maxTirePressurePSI"`
	Model                 *string                            `gorm:"column:model" json:"model"`
	Name                  string                             `gorm:"not null;column:name" json:"name"`
	Note                  *string                            `gorm:"column:note" json:"note"`
	OuterDiameterUnit     InventoryTireOuterDiameterUnitEnum `gorm:"not null;column:outerDiameterUnit" json:"outerDiameterUnit"`
	OuterDiameterValue    *float64                           `gorm:"column:outerDiameterValue" json:"outerDiameterValue"`
	PartNumber            *string                            `gorm:"column:partNumber" json:"partNumber"`
	PricingMatrixDate     *datatypes.DateTime                `gorm:"column:pricingMatrixDate" json:"pricingMatrixDate"` // datetime when pricingMatrixId was set, for determining if matrix has been changed api_schema(calculated)
	PricingMatrixID       *string                            `gorm:"column:pricingMatrixId" json:"pricingMatrixId"`
	Quantity              int64                              `gorm:"not null;column:quantity" json:"quantity"`
	RetailCostCents       int64                              `gorm:"not null;column:retailCostCents" json:"retailCostCents"`
	RimUnit               InventoryTireRimUnitEnum           `gorm:"not null;column:rimUnit" json:"rimUnit"`
	RimValue              *float64                           `gorm:"column:rimValue" json:"rimValue"`
	Seasonality           *InventoryTireSeasonalityEnum      `gorm:"column:seasonality" json:"seasonality"`
	ServiceType           *string                            `gorm:"column:serviceType" json:"serviceType"`
	ShowNote              bool                               `gorm:"not null;column:showNote" json:"showNote"`
	ShowPartNumber        bool                               `gorm:"not null;column:showPartNumber" json:"showPartNumber"`
	ShowPriceAndQuantity  bool                               `gorm:"not null;column:showPriceAndQuantity" json:"showPriceAndQuantity"`
	Size                  *string                            `gorm:"column:size" json:"size"`
	SizeFormat            InventoryTireSizeFormatEnum        `gorm:"not null;column:sizeFormat" json:"sizeFormat"`
	SpeedRating           *string                            `gorm:"column:speedRating" json:"speedRating"`
	Taxable               bool                               `gorm:"not null;column:taxable" json:"taxable"`
	TemperatureUTQG       *string                            `gorm:"column:temperature_UTQG" json:"temperature_UTQG"`
	TractionUTQG          *string                            `gorm:"column:traction_UTQG" json:"traction_UTQG"`
	TreadDepthUnit        InventoryTireTreadDepthUnitEnum    `gorm:"not null;column:treadDepthUnit" json:"treadDepthUnit"`
	TreadDepthValue       *float64                           `gorm:"column:treadDepthValue" json:"treadDepthValue"`
	TreadWearUTQG         *float64                           `gorm:"column:treadWear_UTQG" json:"treadWear_UTQG"`
	UserID                *string                            `gorm:"column:userId" json:"userId"`
	VendorID              *string                            `gorm:"column:vendorId" json:"vendorId"`
	WasteTireFee          int64                              `gorm:"not null;column:wasteTireFee" json:"wasteTireFee"`
	WholesaleCostCents    *int64                             `gorm:"column:wholesaleCostCents" json:"wholesaleCostCents"`
}

var _ Model = (*InventoryTire)(nil)

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *InventoryTire) TableName() string {
	return "inventory_tire"
}

// PrimaryKeys returns an array of the primary keys for this model
func (m *InventoryTire) PrimaryKeys() []string {
	return []string{"id"}
}

// String returns a string representation as JSON for this model
func (m *InventoryTire) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewInventoryTire returns a new model instance from an encoded buffer
func NewInventoryTire(buf []byte) (*InventoryTire, error) {
	var result InventoryTire
	err := json.Unmarshal(buf, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// NewInventoryTireFromChangeEvent returns a new model instance from an encoded buffer as change event
func NewInventoryTireFromChangeEvent(buf []byte, gzip bool) (*datatypes.ChangeEvent[InventoryTire], error) {
	var result datatypes.ChangeEvent[InventoryTire]
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
