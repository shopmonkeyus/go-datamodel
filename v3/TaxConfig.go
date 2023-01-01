// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	datatypes "github.com/shopmonkeyus/go-datamodel/datatypes"
)

type TaxConfigShopSuppliesConfigEnum string

const (
	TaxConfigShopSuppliesConfigNoCap    TaxConfigShopSuppliesConfigEnum = "NoCap"
	TaxConfigShopSuppliesConfigOrderCap TaxConfigShopSuppliesConfigEnum = "OrderCap"
)

type TaxConfigShopSuppliesServiceTypeEnum string

const (
	TaxConfigShopSuppliesServiceTypePercent    TaxConfigShopSuppliesServiceTypeEnum = "Percent"
	TaxConfigShopSuppliesServiceTypeFixedCents TaxConfigShopSuppliesServiceTypeEnum = "FixedCents"
)

type TaxConfigTaxSystemEnum string

const (
	TaxConfigTaxSystemUS TaxConfigTaxSystemEnum = "US"
	TaxConfigTaxSystemCA TaxConfigTaxSystemEnum = "CA"
	TaxConfigTaxSystemMX TaxConfigTaxSystemEnum = "MX"
)

type TaxConfig struct {
	ID          string              `gorm:"primaryKey;not null;column:id" json:"id"`
	CreatedDate datatypes.DateTime  `gorm:"column:createdDate;not null;column:createdDate" json:"createdDate"`
	UpdatedDate *datatypes.DateTime `gorm:"column:updatedDate;column:updatedDate" json:"updatedDate"`
	Meta        datatypes.Meta      `gorm:"column:meta;not null;column:meta" json:"meta,omitempty"`    // the metadata about the most recent change to the row
	Metadata    *datatypes.JSON     `gorm:"column:metadata;column:metadata" json:"metadata,omitempty"` // metadata reserved for customers to control
	CompanyID   string              `gorm:"not null;column:companyId" json:"companyId"`
	LocationID  string              `gorm:"not null;column:locationId" json:"locationId"`

	EpaPercent                       float64                              `gorm:"not null;column:epaPercent" json:"epaPercent"`
	EpaTaxable                       bool                                 `gorm:"not null;column:epaTaxable" json:"epaTaxable"`
	GstPercent                       float64                              `gorm:"not null;column:gstPercent" json:"gstPercent"`
	Hash                             string                               `gorm:"not null;column:hash" json:"hash"`
	HstPercent                       float64                              `gorm:"not null;column:hstPercent" json:"hstPercent"`
	IncludeEpaOnLabor                bool                                 `gorm:"not null;column:includeEpaOnLabor" json:"includeEpaOnLabor"`
	IncludeEpaOnParts                bool                                 `gorm:"not null;column:includeEpaOnParts" json:"includeEpaOnParts"`
	LaborShopSupplies                bool                                 `gorm:"not null;column:laborShopSupplies" json:"laborShopSupplies"`
	LaborTaxable                     bool                                 `gorm:"not null;column:laborTaxable" json:"laborTaxable"`
	PartShopSupplies                 bool                                 `gorm:"not null;column:partShopSupplies" json:"partShopSupplies"`
	PartTaxable                      bool                                 `gorm:"not null;column:partTaxable" json:"partTaxable"`
	PriceCalculatorVersion           float64                              `gorm:"not null;column:priceCalculatorVersion" json:"priceCalculatorVersion"`
	PstPercent                       float64                              `gorm:"not null;column:pstPercent" json:"pstPercent"`
	RateID                           string                               `gorm:"not null;column:rateId" json:"rateId"`
	ShopSuppliesConfig               TaxConfigShopSuppliesConfigEnum      `gorm:"not null;column:shopSuppliesConfig" json:"shopSuppliesConfig"`
	ShopSuppliesMaxCapCents          int64                                `gorm:"not null;column:shopSuppliesMaxCapCents" json:"shopSuppliesMaxCapCents"`
	ShopSuppliesMaxCapServicePercent float64                              `gorm:"not null;column:shopSuppliesMaxCapServicePercent" json:"shopSuppliesMaxCapServicePercent"`
	ShopSuppliesServiceCents         int64                                `gorm:"not null;column:shopSuppliesServiceCents" json:"shopSuppliesServiceCents"`
	ShopSuppliesServicePercent       float64                              `gorm:"not null;column:shopSuppliesServicePercent" json:"shopSuppliesServicePercent"`
	ShopSuppliesServiceType          TaxConfigShopSuppliesServiceTypeEnum `gorm:"not null;column:shopSuppliesServiceType" json:"shopSuppliesServiceType"`
	ShopSuppliesTaxable              bool                                 `gorm:"not null;column:shopSuppliesTaxable" json:"shopSuppliesTaxable"`
	SubcontractTaxable               bool                                 `gorm:"not null;column:subcontractTaxable" json:"subcontractTaxable"`
	TaxPercent                       float64                              `gorm:"not null;column:taxPercent" json:"taxPercent"`
	TaxSystem                        TaxConfigTaxSystemEnum               `gorm:"not null;column:taxSystem" json:"taxSystem"`
	Version                          int64                                `gorm:"not null;column:version" json:"version"`
}

var _ Model = (*TaxConfig)(nil)

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *TaxConfig) TableName() string {
	return "tax_config"
}

// String returns a string representation as JSON for this model
func (m *TaxConfig) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewTaxConfig returns a new model instance from an encoded buffer
func NewTaxConfig(buf []byte) (*TaxConfig, error) {
	var result TaxConfig
	err := json.Unmarshal(buf, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// NewTaxConfigFromChangeEvent returns a new model instance from an encoded buffer as change event
func NewTaxConfigFromChangeEvent(buf []byte, gzip bool) (*datatypes.ChangeEvent[TaxConfig], error) {
	var result datatypes.ChangeEvent[TaxConfig]
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
