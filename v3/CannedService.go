// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	"time"
)

type CannedServiceDiscountValueTypeEnum string

const (
	CannedServiceDiscountValueTypePercent    CannedServiceDiscountValueTypeEnum = "Percent"
	CannedServiceDiscountValueTypeFixedCents                                    = "FixedCents"
)

type CannedServiceEpaValueTypeEnum string

const (
	CannedServiceEpaValueTypePercent    CannedServiceEpaValueTypeEnum = "Percent"
	CannedServiceEpaValueTypeFixedCents                               = "FixedCents"
)

type CannedServiceGstValueTypeEnum string

const (
	CannedServiceGstValueTypePercent    CannedServiceGstValueTypeEnum = "Percent"
	CannedServiceGstValueTypeFixedCents                               = "FixedCents"
)

type CannedServiceHstValueTypeEnum string

const (
	CannedServiceHstValueTypePercent    CannedServiceHstValueTypeEnum = "Percent"
	CannedServiceHstValueTypeFixedCents                               = "FixedCents"
)

type CannedServicePstValueTypeEnum string

const (
	CannedServicePstValueTypePercent    CannedServicePstValueTypeEnum = "Percent"
	CannedServicePstValueTypeFixedCents                               = "FixedCents"
)

type CannedServiceShopSuppliesValueTypeEnum string

const (
	CannedServiceShopSuppliesValueTypePercent    CannedServiceShopSuppliesValueTypeEnum = "Percent"
	CannedServiceShopSuppliesValueTypeFixedCents                                        = "FixedCents"
)

type CannedServiceTaxValueTypeEnum string

const (
	CannedServiceTaxValueTypePercent    CannedServiceTaxValueTypeEnum = "Percent"
	CannedServiceTaxValueTypeFixedCents                               = "FixedCents"
)

type CannedService struct {
	CalculatedDiscountCents     int64                                  `gorm:"not null" json:"calculatedDiscountCents"`
	ID                          string                                 `gorm:"primaryKey;not null" json:"id"`
	CalculatedDiscountPercent   float64                                `gorm:"not null" json:"calculatedDiscountPercent"`
	CalculatedEpaCents          int64                                  `gorm:"not null" json:"calculatedEpaCents"`
	CalculatedFeeCents          int64                                  `gorm:"not null" json:"calculatedFeeCents"`
	CalculatedLaborCents        int64                                  `gorm:"not null" json:"calculatedLaborCents"`
	CalculatedPartsCents        int64                                  `gorm:"not null" json:"calculatedPartsCents"`
	CalculatedShopSuppliesCents int64                                  `gorm:"not null" json:"calculatedShopSuppliesCents"`
	CalculatedSubcontractsCents int64                                  `gorm:"not null" json:"calculatedSubcontractsCents"`
	CalculatedTaxCents          int64                                  `gorm:"not null" json:"calculatedTaxCents"`
	CalculatedTiresCents        int64                                  `gorm:"not null" json:"calculatedTiresCents"`
	CompanyId                   string                                 `gorm:"not null" json:"companyId"`
	CreatedDate                 time.Time                              `gorm:"column:createdDate;not null" json:"createdDate"`
	Deleted                     bool                                   `gorm:"not null" json:"deleted"` // if the record has been deleted
	DeletedDate                 *time.Time                             `json:"deletedDate"`             // the date that the record was deleted or null if not deleted
	DeletedReason               *string                                `json:"deletedReason"`           // the reason that the record was deleted
	DeletedUserId               *string                                `json:"deletedUserId"`           // the user that deleted the record or null if not deleted
	DiscountCents               int64                                  `gorm:"not null" json:"discountCents"`
	DiscountPercent             float64                                `gorm:"not null" json:"discountPercent"`
	DiscountValueType           CannedServiceDiscountValueTypeEnum     `gorm:"not null" json:"discountValueType"`
	EpaCents                    int64                                  `gorm:"not null" json:"epaCents"`
	EpaPercent                  float64                                `gorm:"not null" json:"epaPercent"`
	EpaValueType                CannedServiceEpaValueTypeEnum          `gorm:"not null" json:"epaValueType"`
	GstCents                    int64                                  `gorm:"not null" json:"gstCents"`
	GstPercent                  float64                                `gorm:"not null" json:"gstPercent"`
	GstValueType                CannedServiceGstValueTypeEnum          `gorm:"not null" json:"gstValueType"`
	HstCents                    int64                                  `gorm:"not null" json:"hstCents"`
	HstPercent                  float64                                `gorm:"not null" json:"hstPercent"`
	HstValueType                CannedServiceHstValueTypeEnum          `gorm:"not null" json:"hstValueType"`
	LocationId                  string                                 `gorm:"not null" json:"locationId"`
	LumpSum                     bool                                   `gorm:"not null" json:"lumpSum"`
	Meta                        *Meta                                  `gorm:"embedded;column:meta;not null" json:"meta,omitempty"`         // the metadata about the most recent change to the row
	Metadata                    any                                    `gorm:"embedded;column:metadata;not null" json:"metadata,omitempty"` // metadata reserved for customers to control
	Name                        string                                 `gorm:"not null" json:"name"`
	Note                        string                                 `gorm:"not null" json:"note"`
	PstCents                    int64                                  `gorm:"not null" json:"pstCents"`
	PstPercent                  float64                                `gorm:"not null" json:"pstPercent"`
	PstValueType                CannedServicePstValueTypeEnum          `gorm:"not null" json:"pstValueType"`
	Recommended                 bool                                   `gorm:"not null" json:"recommended"`
	ShopSuppliesApplied         bool                                   `gorm:"not null" json:"shopSuppliesApplied"`
	ShopSuppliesCents           int64                                  `gorm:"not null" json:"shopSuppliesCents"`
	ShopSuppliesPercent         float64                                `gorm:"not null" json:"shopSuppliesPercent"`
	ShopSuppliesValueType       CannedServiceShopSuppliesValueTypeEnum `gorm:"not null" json:"shopSuppliesValueType"`
	TaxCents                    int64                                  `gorm:"not null" json:"taxCents"`
	TaxPercent                  float64                                `gorm:"not null" json:"taxPercent"`
	TaxValueType                CannedServiceTaxValueTypeEnum          `gorm:"not null" json:"taxValueType"`
	TotalCents                  int64                                  `gorm:"not null" json:"totalCents"`
	UpdatedDate                 *time.Time                             `gorm:"column:updatedDate" json:"updatedDate"`
}

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *CannedService) TableName() string {
	return "canned_service"
}

func (m *CannedService) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}
