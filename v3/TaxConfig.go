// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	mapstructure "github.com/mitchellh/mapstructure"
	"time"
)

type TaxConfigShopSuppliesConfigEnum string

const (
	TaxConfigShopSuppliesConfigNoCap    TaxConfigShopSuppliesConfigEnum = "NoCap"
	TaxConfigShopSuppliesConfigOrderCap                                 = "OrderCap"
)

type TaxConfigShopSuppliesServiceTypeEnum string

const (
	TaxConfigShopSuppliesServiceTypePercent    TaxConfigShopSuppliesServiceTypeEnum = "Percent"
	TaxConfigShopSuppliesServiceTypeFixedCents                                      = "FixedCents"
)

type TaxConfigTaxSystemEnum string

const (
	TaxConfigTaxSystemUS TaxConfigTaxSystemEnum = "US"
	TaxConfigTaxSystemCA                        = "CA"
	TaxConfigTaxSystemMX                        = "MX"
)

type TaxConfig struct {
	CompanyId                        string                               `gorm:"not null" json:"companyId"`
	CreatedDate                      time.Time                            `gorm:"column:createdDate;not null" json:"createdDate"`
	EpaPercent                       float64                              `gorm:"not null" json:"epaPercent"`
	EpaTaxable                       bool                                 `gorm:"not null" json:"epaTaxable"`
	GstPercent                       float64                              `gorm:"not null" json:"gstPercent"`
	Hash                             string                               `gorm:"not null" json:"hash"`
	HstPercent                       float64                              `gorm:"not null" json:"hstPercent"`
	ID                               string                               `gorm:"primaryKey;not null" json:"id"`
	IncludeEpaOnLabor                bool                                 `gorm:"not null" json:"includeEpaOnLabor"`
	IncludeEpaOnParts                bool                                 `gorm:"not null" json:"includeEpaOnParts"`
	LaborShopSupplies                bool                                 `gorm:"not null" json:"laborShopSupplies"`
	LaborTaxable                     bool                                 `gorm:"not null" json:"laborTaxable"`
	LocationId                       string                               `gorm:"not null" json:"locationId"`
	Meta                             *Meta                                `gorm:"type:json;embedded;column:meta;not null" json:"meta,omitempty"` // the metadata about the most recent change to the row
	Metadata                         any                                  `gorm:"type:json" json:"metadata,omitempty"`                           // metadata reserved for customers to control
	PartShopSupplies                 bool                                 `gorm:"not null" json:"partShopSupplies"`
	PartTaxable                      bool                                 `gorm:"not null" json:"partTaxable"`
	PriceCalculatorVersion           float64                              `gorm:"not null" json:"priceCalculatorVersion"`
	PstPercent                       float64                              `gorm:"not null" json:"pstPercent"`
	RateId                           string                               `gorm:"not null" json:"rateId"`
	ShopSuppliesConfig               TaxConfigShopSuppliesConfigEnum      `gorm:"not null" json:"shopSuppliesConfig"`
	ShopSuppliesMaxCapCents          int64                                `gorm:"not null" json:"shopSuppliesMaxCapCents"`
	ShopSuppliesMaxCapServicePercent float64                              `gorm:"not null" json:"shopSuppliesMaxCapServicePercent"`
	ShopSuppliesServiceCents         int64                                `gorm:"not null" json:"shopSuppliesServiceCents"`
	ShopSuppliesServicePercent       float64                              `gorm:"not null" json:"shopSuppliesServicePercent"`
	ShopSuppliesServiceType          TaxConfigShopSuppliesServiceTypeEnum `gorm:"not null" json:"shopSuppliesServiceType"`
	ShopSuppliesTaxable              bool                                 `gorm:"not null" json:"shopSuppliesTaxable"`
	SubcontractTaxable               bool                                 `gorm:"not null" json:"subcontractTaxable"`
	TaxPercent                       float64                              `gorm:"not null" json:"taxPercent"`
	TaxSystem                        TaxConfigTaxSystemEnum               `gorm:"not null" json:"taxSystem"`
	UpdatedDate                      *time.Time                           `gorm:"column:updatedDate" json:"updatedDate"`
	Version                          int64                                `gorm:"not null" json:"version"`
}

var _ Model = (*TaxConfig)(nil)

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *TaxConfig) TableName() string {
	return "tax_config"
}

func (m *TaxConfig) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewTaxConfig returns a new model instance from a json key/value map
func NewTaxConfig(kv map[string]any) (*TaxConfig, error) {
	var result TaxConfig
	err := mapstructure.Decode(kv, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
