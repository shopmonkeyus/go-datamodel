// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	"time"
)

type CannedServiceFeeFeeTypeEnum string

const (
	CannedServiceFeeFeeTypePercentLineItem CannedServiceFeeFeeTypeEnum = "PercentLineItem"
	CannedServiceFeeFeeTypePercentService                              = "PercentService"
	CannedServiceFeeFeeTypeFixedCents                                  = "FixedCents"
)

type CannedServiceFeeLineItemEntityEnum string

const (
	CannedServiceFeeLineItemEntityLabor       CannedServiceFeeLineItemEntityEnum = "Labor"
	CannedServiceFeeLineItemEntityPart                                           = "Part"
	CannedServiceFeeLineItemEntitySubcontract                                    = "Subcontract"
	CannedServiceFeeLineItemEntityTire                                           = "Tire"
)

type CannedServiceFee struct {
	ID              string                              `gorm:"primaryKey;not null" json:"id"`
	AmountCents     int64                               `gorm:"not null" json:"amountCents"`
	CannedServiceId string                              `gorm:"not null" json:"cannedServiceId"`
	CategoryId      *string                             `json:"categoryId"`
	CompanyId       string                              `gorm:"not null" json:"companyId"`
	CreatedDate     time.Time                           `gorm:"column:createdDate;not null" json:"createdDate"`
	FeeType         CannedServiceFeeFeeTypeEnum         `gorm:"not null" json:"feeType"`
	LaborId         *string                             `json:"laborId"`
	LineItemEntity  *CannedServiceFeeLineItemEntityEnum `json:"lineItemEntity"`
	LocationId      string                              `gorm:"not null" json:"locationId"`
	Meta            *Meta                               `gorm:"embedded;column:meta;not null" json:"meta,omitempty"`         // the metadata about the most recent change to the row
	Metadata        any                                 `gorm:"embedded;column:metadata;not null" json:"metadata,omitempty"` // metadata reserved for customers to control
	Name            string                              `gorm:"not null" json:"name"`
	Ordinal         float64                             `gorm:"not null" json:"ordinal"`
	PartId          *string                             `json:"partId"`
	Percent         float64                             `gorm:"not null" json:"percent"`
	SourceItemId    *string                             `json:"sourceItemId"`
	SubcontractId   *string                             `json:"subcontractId"`
	Subtotal        *int64                              `json:"subtotal"`
	TireId          *string                             `json:"tireId"`
	UpdatedDate     *time.Time                          `gorm:"column:updatedDate" json:"updatedDate"`
}

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *CannedServiceFee) TableName() string {
	return "canned_service_fee"
}

func (m *CannedServiceFee) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}
