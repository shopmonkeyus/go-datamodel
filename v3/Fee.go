// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	mapstructure "github.com/mitchellh/mapstructure"
	"time"
)

type FeeFeeTypeEnum string

const (
	FeeFeeTypePercentLineItem FeeFeeTypeEnum = "PercentLineItem"
	FeeFeeTypePercentService                 = "PercentService"
	FeeFeeTypeFixedCents                     = "FixedCents"
)

type FeeLineItemEntityEnum string

const (
	FeeLineItemEntityLabor       FeeLineItemEntityEnum = "Labor"
	FeeLineItemEntityPart                              = "Part"
	FeeLineItemEntitySubcontract                       = "Subcontract"
	FeeLineItemEntityTire                              = "Tire"
)

type Fee struct {
	AmountCents    int64                  `gorm:"not null" json:"amountCents"`
	ID             string                 `gorm:"primaryKey;not null" json:"id"`
	CategoryId     *string                `json:"categoryId"`
	CompanyId      string                 `gorm:"not null" json:"companyId"`
	CreatedDate    time.Time              `gorm:"column:createdDate;not null" json:"createdDate"`
	FeeType        FeeFeeTypeEnum         `gorm:"not null" json:"feeType"`
	LaborId        *string                `json:"laborId"`
	LineItemEntity *FeeLineItemEntityEnum `json:"lineItemEntity"`
	LocationId     string                 `gorm:"not null" json:"locationId"`
	Meta           *Meta                  `gorm:"type:json;embedded;column:meta;not null" json:"meta,omitempty"` // the metadata about the most recent change to the row
	Metadata       any                    `gorm:"type:json" json:"metadata,omitempty"`                           // metadata reserved for customers to control
	Name           string                 `gorm:"not null" json:"name"`
	OrderId        string                 `gorm:"not null" json:"orderId"`
	Ordinal        float64                `gorm:"not null" json:"ordinal"`
	PartId         *string                `json:"partId"`
	Percent        float64                `gorm:"not null" json:"percent"`
	ServiceId      string                 `gorm:"not null" json:"serviceId"`
	SourceItemId   *string                `json:"sourceItemId"`
	SubcontractId  *string                `json:"subcontractId"`
	Subtotal       *int64                 `json:"subtotal"`
	TireId         *string                `json:"tireId"`
	UpdatedDate    *time.Time             `gorm:"column:updatedDate" json:"updatedDate"`
}

var _ Model = (*Fee)(nil)

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *Fee) TableName() string {
	return "fee"
}

func (m *Fee) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewFee returns a new model instance from a json key/value map
func NewFee(kv map[string]any) (*Fee, error) {
	var result Fee
	err := mapstructure.Decode(kv, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
