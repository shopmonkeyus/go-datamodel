// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	"time"
)

type ServiceAuthorizationStatusEnum string

const (
	ServiceAuthorizationStatusNotAuthorized ServiceAuthorizationStatusEnum = "NotAuthorized"
	ServiceAuthorizationStatusAuthorized                                   = "Authorized"
	ServiceAuthorizationStatusDeclined                                     = "Declined"
)

type ServiceDeferredReasonEnum string

const (
	ServiceDeferredReasonArchived              ServiceDeferredReasonEnum = "Archived"
	ServiceDeferredReasonDeclined                                        = "Declined"
	ServiceDeferredReasonInvoicedNotAuthorized                           = "InvoicedNotAuthorized"
)

type ServiceDiscountValueTypeEnum string

const (
	ServiceDiscountValueTypePercent    ServiceDiscountValueTypeEnum = "Percent"
	ServiceDiscountValueTypeFixedCents                              = "FixedCents"
)

type ServiceEpaValueTypeEnum string

const (
	ServiceEpaValueTypePercent    ServiceEpaValueTypeEnum = "Percent"
	ServiceEpaValueTypeFixedCents                         = "FixedCents"
)

type ServiceGstValueTypeEnum string

const (
	ServiceGstValueTypePercent    ServiceGstValueTypeEnum = "Percent"
	ServiceGstValueTypeFixedCents                         = "FixedCents"
)

type ServiceHstValueTypeEnum string

const (
	ServiceHstValueTypePercent    ServiceHstValueTypeEnum = "Percent"
	ServiceHstValueTypeFixedCents                         = "FixedCents"
)

type ServicePstValueTypeEnum string

const (
	ServicePstValueTypePercent    ServicePstValueTypeEnum = "Percent"
	ServicePstValueTypeFixedCents                         = "FixedCents"
)

type ServiceShopSuppliesValueTypeEnum string

const (
	ServiceShopSuppliesValueTypePercent    ServiceShopSuppliesValueTypeEnum = "Percent"
	ServiceShopSuppliesValueTypeFixedCents                                  = "FixedCents"
)

type ServiceTaxValueTypeEnum string

const (
	ServiceTaxValueTypePercent    ServiceTaxValueTypeEnum = "Percent"
	ServiceTaxValueTypeFixedCents                         = "FixedCents"
)

type Service struct {
	ID                          string                           `gorm:"primaryKey;not null" json:"id"`
	AuthorizationStatus         ServiceAuthorizationStatusEnum   `gorm:"not null" json:"authorizationStatus"`
	CalculatedDiscountCents     int64                            `gorm:"not null" json:"calculatedDiscountCents"`
	CalculatedDiscountPercent   float64                          `gorm:"not null" json:"calculatedDiscountPercent"`
	CalculatedEpaCents          int64                            `gorm:"not null" json:"calculatedEpaCents"`
	CalculatedFeeCents          int64                            `gorm:"not null" json:"calculatedFeeCents"`
	CalculatedLaborCents        int64                            `gorm:"not null" json:"calculatedLaborCents"`
	CalculatedPartsCents        int64                            `gorm:"not null" json:"calculatedPartsCents"`
	CalculatedShopSuppliesCents int64                            `gorm:"not null" json:"calculatedShopSuppliesCents"`
	CalculatedSubcontractsCents int64                            `gorm:"not null" json:"calculatedSubcontractsCents"`
	CalculatedTaxCents          int64                            `gorm:"not null" json:"calculatedTaxCents"`
	CalculatedTiresCents        int64                            `gorm:"not null" json:"calculatedTiresCents"`
	CompanyId                   string                           `gorm:"not null" json:"companyId"`
	CreatedDate                 time.Time                        `gorm:"column:createdDate;not null" json:"createdDate"`
	DeferredDate                *time.Time                       `json:"deferredDate"`
	DeferredReason              *ServiceDeferredReasonEnum       `json:"deferredReason"`
	DiscountCents               int64                            `gorm:"not null" json:"discountCents"`
	DiscountPercent             float64                          `gorm:"not null" json:"discountPercent"`
	DiscountValueType           ServiceDiscountValueTypeEnum     `gorm:"not null" json:"discountValueType"`
	EpaCents                    int64                            `gorm:"not null" json:"epaCents"`
	EpaPercent                  float64                          `gorm:"not null" json:"epaPercent"`
	EpaValueType                ServiceEpaValueTypeEnum          `gorm:"not null" json:"epaValueType"`
	ExcludedFromDeferred        bool                             `gorm:"not null" json:"excludedFromDeferred"` // service that has been removed from the deferred list
	GstCents                    int64                            `gorm:"not null" json:"gstCents"`
	GstPercent                  float64                          `gorm:"not null" json:"gstPercent"`
	GstValueType                ServiceGstValueTypeEnum          `gorm:"not null" json:"gstValueType"`
	Hidden                      bool                             `gorm:"not null" json:"hidden"`
	HstCents                    int64                            `gorm:"not null" json:"hstCents"`
	HstPercent                  float64                          `gorm:"not null" json:"hstPercent"`
	HstValueType                ServiceHstValueTypeEnum          `gorm:"not null" json:"hstValueType"`
	LocationId                  string                           `gorm:"not null" json:"locationId"`
	LumpSum                     bool                             `gorm:"not null" json:"lumpSum"`
	Meta                        *Meta                            `gorm:"embedded;column:meta;not null" json:"meta,omitempty"`         // the metadata about the most recent change to the row
	Metadata                    any                              `gorm:"embedded;column:metadata;not null" json:"metadata,omitempty"` // metadata reserved for customers to control
	Name                        string                           `gorm:"not null" json:"name"`
	Note                        string                           `gorm:"not null" json:"note"`
	OrderId                     string                           `gorm:"not null" json:"orderId"`
	Ordinal                     float64                          `gorm:"not null" json:"ordinal"`
	PstCents                    int64                            `gorm:"not null" json:"pstCents"`
	PstPercent                  float64                          `gorm:"not null" json:"pstPercent"`
	PstValueType                ServicePstValueTypeEnum          `gorm:"not null" json:"pstValueType"`
	Recommended                 bool                             `gorm:"not null" json:"recommended"`
	Revived                     bool                             `gorm:"not null" json:"revived"`             // if a deferred service is re-added
	RevivedFromId               *string                          `json:"revivedFromId"`                       // service that was deferred and then re-added, used for chain of deferment tracing
	ShopSuppliesApplied         bool                             `gorm:"not null" json:"shopSuppliesApplied"` // if SS max cap percentage should apply, always true if taxConfig SS setting is NoCap
	ShopSuppliesCents           int64                            `gorm:"not null" json:"shopSuppliesCents"`
	ShopSuppliesPercent         float64                          `gorm:"not null" json:"shopSuppliesPercent"`
	ShopSuppliesValueType       ServiceShopSuppliesValueTypeEnum `gorm:"not null" json:"shopSuppliesValueType"`
	SourceServiceId             *string                          `json:"sourceServiceId"`
	TaxCents                    int64                            `gorm:"not null" json:"taxCents"`
	TaxPercent                  float64                          `gorm:"not null" json:"taxPercent"`
	TaxValueType                ServiceTaxValueTypeEnum          `gorm:"not null" json:"taxValueType"`
	TotalCents                  int64                            `gorm:"not null" json:"totalCents"`
	UpdatedDate                 *time.Time                       `gorm:"column:updatedDate" json:"updatedDate"`
}

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *Service) TableName() string {
	return "service"
}

func (m *Service) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}
