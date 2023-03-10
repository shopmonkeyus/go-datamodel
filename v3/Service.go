// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	datatypes "github.com/shopmonkeyus/go-datamodel/datatypes"
)

// Service schema
type ServiceAuthorizationStatusEnum string

const (
	ServiceAuthorizationStatusNotAuthorized ServiceAuthorizationStatusEnum = "NotAuthorized"
	ServiceAuthorizationStatusAuthorized    ServiceAuthorizationStatusEnum = "Authorized"
	ServiceAuthorizationStatusDeclined      ServiceAuthorizationStatusEnum = "Declined"
)

type ServiceDeferredReasonEnum string

const (
	ServiceDeferredReasonArchived              ServiceDeferredReasonEnum = "Archived"
	ServiceDeferredReasonDeclined              ServiceDeferredReasonEnum = "Declined"
	ServiceDeferredReasonInvoicedNotAuthorized ServiceDeferredReasonEnum = "InvoicedNotAuthorized"
)

type ServiceDiscountValueTypeEnum string

const (
	ServiceDiscountValueTypePercent    ServiceDiscountValueTypeEnum = "Percent"
	ServiceDiscountValueTypeFixedCents ServiceDiscountValueTypeEnum = "FixedCents"
)

type ServiceEpaValueTypeEnum string

const (
	ServiceEpaValueTypePercent    ServiceEpaValueTypeEnum = "Percent"
	ServiceEpaValueTypeFixedCents ServiceEpaValueTypeEnum = "FixedCents"
)

type ServiceGstValueTypeEnum string

const (
	ServiceGstValueTypePercent    ServiceGstValueTypeEnum = "Percent"
	ServiceGstValueTypeFixedCents ServiceGstValueTypeEnum = "FixedCents"
)

type ServiceHstValueTypeEnum string

const (
	ServiceHstValueTypePercent    ServiceHstValueTypeEnum = "Percent"
	ServiceHstValueTypeFixedCents ServiceHstValueTypeEnum = "FixedCents"
)

type ServicePstValueTypeEnum string

const (
	ServicePstValueTypePercent    ServicePstValueTypeEnum = "Percent"
	ServicePstValueTypeFixedCents ServicePstValueTypeEnum = "FixedCents"
)

type ServiceShopSuppliesValueTypeEnum string

const (
	ServiceShopSuppliesValueTypePercent    ServiceShopSuppliesValueTypeEnum = "Percent"
	ServiceShopSuppliesValueTypeFixedCents ServiceShopSuppliesValueTypeEnum = "FixedCents"
)

type ServiceTaxValueTypeEnum string

const (
	ServiceTaxValueTypePercent    ServiceTaxValueTypeEnum = "Percent"
	ServiceTaxValueTypeFixedCents ServiceTaxValueTypeEnum = "FixedCents"
)

type Service struct {
	ID          string              `bson:"_id" gorm:"primaryKey;not null;column:id" json:"id"`
	CreatedDate datatypes.DateTime  `gorm:"column:createdDate;not null;column:createdDate" json:"createdDate"`
	UpdatedDate *datatypes.DateTime `gorm:"column:updatedDate;column:updatedDate" json:"updatedDate"`
	Meta        *datatypes.JSON     `gorm:"column:meta;not null;column:meta" json:"meta,omitempty"` // the metadata about the most recent change to the row
	Metadata    *datatypes.JSON     `gorm:"column:metadata;column:metadata" json:"metadata,omitempty"`
	CompanyID   string              `gorm:"not null;column:companyId" json:"companyId"`
	LocationID  string              `gorm:"not null;column:locationId" json:"locationId"`

	AuthorizationStatus         ServiceAuthorizationStatusEnum   `gorm:"not null;column:authorizationStatus" json:"authorizationStatus"`
	CalculatedDiscountCents     int64                            `gorm:"not null;column:calculatedDiscountCents" json:"calculatedDiscountCents"`
	CalculatedDiscountPercent   float64                          `gorm:"not null;column:calculatedDiscountPercent" json:"calculatedDiscountPercent"`
	CalculatedEpaCents          int64                            `gorm:"not null;column:calculatedEpaCents" json:"calculatedEpaCents"`
	CalculatedFeeCents          int64                            `gorm:"not null;column:calculatedFeeCents" json:"calculatedFeeCents"`
	CalculatedLaborCents        int64                            `gorm:"not null;column:calculatedLaborCents" json:"calculatedLaborCents"`
	CalculatedPartsCents        int64                            `gorm:"not null;column:calculatedPartsCents" json:"calculatedPartsCents"`
	CalculatedShopSuppliesCents int64                            `gorm:"not null;column:calculatedShopSuppliesCents" json:"calculatedShopSuppliesCents"`
	CalculatedSubcontractsCents int64                            `gorm:"not null;column:calculatedSubcontractsCents" json:"calculatedSubcontractsCents"`
	CalculatedTaxCents          int64                            `gorm:"not null;column:calculatedTaxCents" json:"calculatedTaxCents"`
	CalculatedTiresCents        int64                            `gorm:"not null;column:calculatedTiresCents" json:"calculatedTiresCents"`
	DeferredDate                *datatypes.DateTime              `gorm:"column:deferredDate" json:"deferredDate"`
	DeferredReason              *ServiceDeferredReasonEnum       `gorm:"column:deferredReason" json:"deferredReason"`
	DiscountCents               int64                            `gorm:"not null;column:discountCents" json:"discountCents"`
	DiscountPercent             float64                          `gorm:"not null;column:discountPercent" json:"discountPercent"`
	DiscountValueType           ServiceDiscountValueTypeEnum     `gorm:"not null;column:discountValueType" json:"discountValueType"`
	EpaCents                    int64                            `gorm:"not null;column:epaCents" json:"epaCents"`
	EpaPercent                  float64                          `gorm:"not null;column:epaPercent" json:"epaPercent"`
	EpaValueType                ServiceEpaValueTypeEnum          `gorm:"not null;column:epaValueType" json:"epaValueType"`
	ExcludedFromDeferred        bool                             `gorm:"not null;column:excludedFromDeferred" json:"excludedFromDeferred"` // service that has been removed from the deferred list
	GstCents                    int64                            `gorm:"not null;column:gstCents" json:"gstCents"`
	GstPercent                  float64                          `gorm:"not null;column:gstPercent" json:"gstPercent"`
	GstValueType                ServiceGstValueTypeEnum          `gorm:"not null;column:gstValueType" json:"gstValueType"`
	Hidden                      bool                             `gorm:"not null;column:hidden" json:"hidden"`
	HstCents                    int64                            `gorm:"not null;column:hstCents" json:"hstCents"`
	HstPercent                  float64                          `gorm:"not null;column:hstPercent" json:"hstPercent"`
	HstValueType                ServiceHstValueTypeEnum          `gorm:"not null;column:hstValueType" json:"hstValueType"`
	LumpSum                     bool                             `gorm:"not null;column:lumpSum" json:"lumpSum"`
	Name                        string                           `gorm:"not null;column:name" json:"name"`
	Note                        string                           `gorm:"not null;column:note" json:"note"`
	OrderID                     string                           `gorm:"not null;column:orderId" json:"orderId"`
	Ordinal                     float64                          `gorm:"not null;column:ordinal" json:"ordinal"`
	PstCents                    int64                            `gorm:"not null;column:pstCents" json:"pstCents"`
	PstPercent                  float64                          `gorm:"not null;column:pstPercent" json:"pstPercent"`
	PstValueType                ServicePstValueTypeEnum          `gorm:"not null;column:pstValueType" json:"pstValueType"`
	Recommended                 bool                             `gorm:"not null;column:recommended" json:"recommended"`
	Revived                     bool                             `gorm:"not null;column:revived" json:"revived"`                         // if a deferred service is re-added
	RevivedFromID               *string                          `gorm:"column:revivedFromId" json:"revivedFromId"`                      // service that was deferred and then re-added, used for chain of deferment tracing
	ShopSuppliesApplied         bool                             `gorm:"not null;column:shopSuppliesApplied" json:"shopSuppliesApplied"` // if SS max cap percentage should apply, always true if taxConfig SS setting is NoCap
	ShopSuppliesCents           int64                            `gorm:"not null;column:shopSuppliesCents" json:"shopSuppliesCents"`
	ShopSuppliesPercent         float64                          `gorm:"not null;column:shopSuppliesPercent" json:"shopSuppliesPercent"`
	ShopSuppliesValueType       ServiceShopSuppliesValueTypeEnum `gorm:"not null;column:shopSuppliesValueType" json:"shopSuppliesValueType"`
	SourceServiceID             *string                          `gorm:"column:sourceServiceId" json:"sourceServiceId"`
	TaxCents                    int64                            `gorm:"not null;column:taxCents" json:"taxCents"`
	TaxPercent                  float64                          `gorm:"not null;column:taxPercent" json:"taxPercent"`
	TaxValueType                ServiceTaxValueTypeEnum          `gorm:"not null;column:taxValueType" json:"taxValueType"`
	TotalCents                  int64                            `gorm:"not null;column:totalCents" json:"totalCents"`
}

var _ Model = (*Service)(nil)

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *Service) TableName() string {
	return "service"
}

// PrimaryKeys returns an array of the primary keys for this model
func (m *Service) PrimaryKeys() []string {
	return []string{"id"}
}

// String returns a string representation as JSON for this model
func (m *Service) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewService returns a new model instance from an encoded buffer
func NewService(buf []byte) (*Service, error) {
	var result Service
	err := json.Unmarshal(buf, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// NewServiceFromChangeEvent returns a new model instance from an encoded buffer as change event
func NewServiceFromChangeEvent(buf []byte, gzip bool) (*datatypes.ChangeEvent[Service], error) {
	var result datatypes.ChangeEvent[Service]
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
