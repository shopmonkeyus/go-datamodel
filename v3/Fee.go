// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	codec "github.com/hashicorp/go-msgpack/v2/codec"
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
	ID          string     `gorm:"primaryKey;not null;column:id" json:"id"`
	CreatedDate time.Time  `gorm:"column:createdDate;not null;column:createdDate" json:"createdDate"`
	UpdatedDate *time.Time `gorm:"column:updatedDate;column:updatedDate" json:"updatedDate"`
	Meta        *Meta      `gorm:"type:json;embedded;serializer:json;column:meta;not null;column:meta" json:"meta,omitempty"` // the metadata about the most recent change to the row
	Metadata    any        `gorm:"type:json;serializer:json;column:metadata" json:"metadata,omitempty"`                       // metadata reserved for customers to control
	CompanyID   string     `gorm:"not null;column:companyId" json:"companyId"`
	LocationID  string     `gorm:"not null;column:locationId" json:"locationId"`

	AmountCents    int64                  `gorm:"not null;column:amountCents" json:"amountCents"`
	CategoryID     *string                `gorm:"column:categoryId" json:"categoryId"`
	FeeType        FeeFeeTypeEnum         `gorm:"not null;column:feeType" json:"feeType"`
	LaborID        *string                `gorm:"column:laborId" json:"laborId"`
	LineItemEntity *FeeLineItemEntityEnum `gorm:"column:lineItemEntity" json:"lineItemEntity"`
	Name           string                 `gorm:"not null;column:name" json:"name"`
	OrderID        string                 `gorm:"not null;column:orderId" json:"orderId"`
	Ordinal        float64                `gorm:"not null;column:ordinal" json:"ordinal"`
	PartID         *string                `gorm:"column:partId" json:"partId"`
	Percent        float64                `gorm:"not null;column:percent" json:"percent"`
	ServiceID      string                 `gorm:"not null;column:serviceId" json:"serviceId"`
	SourceItemID   *string                `gorm:"column:sourceItemId" json:"sourceItemId"`
	SubcontractID  *string                `gorm:"column:subcontractId" json:"subcontractId"`
	Subtotal       *int64                 `gorm:"column:subtotal" json:"subtotal"`
	TireID         *string                `gorm:"column:tireId" json:"tireId"`
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

// NewFee returns a new model instance from an encoded buffer
func NewFee(buf []byte, enctype EncodingType) (*Fee, error) {
	var result Fee
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
