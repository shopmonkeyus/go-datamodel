// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	codec "github.com/hashicorp/go-msgpack/v2/codec"
	datatypes "github.com/shopmonkeyus/go-datamodel/datatypes"
	"time"
)

type CoreChargeStatusEnum string

const (
	CoreChargeStatusNotReady      CoreChargeStatusEnum = "NotReady"
	CoreChargeStatusReady                              = "Ready"
	CoreChargeStatusPickedUp                           = "PickedUp"
	CoreChargeStatusRefunded                           = "Refunded"
	CoreChargeStatusNotRefundable                      = "NotRefundable"
)

type CoreCharge struct {
	ID          string          `gorm:"primaryKey;not null;column:id" json:"id"`
	CreatedDate time.Time       `gorm:"column:createdDate;not null;column:createdDate" json:"createdDate"`
	UpdatedDate *time.Time      `gorm:"column:updatedDate;column:updatedDate" json:"updatedDate"`
	Meta        datatypes.Meta  `gorm:"column:meta;not null;column:meta" json:"meta,omitempty"`    // the metadata about the most recent change to the row
	Metadata    *datatypes.JSON `gorm:"column:metadata;column:metadata" json:"metadata,omitempty"` // metadata reserved for customers to control
	CompanyID   string          `gorm:"not null;column:companyId" json:"companyId"`
	LocationID  string          `gorm:"not null;column:locationId" json:"locationId"`

	AmountCents int64                `gorm:"not null;column:amountCents" json:"amountCents"`
	PartID      string               `gorm:"not null;column:partId" json:"partId"`
	Status      CoreChargeStatusEnum `gorm:"not null;column:status" json:"status"`
}

var _ Model = (*CoreCharge)(nil)

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *CoreCharge) TableName() string {
	return "core_charge"
}

func (m *CoreCharge) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewCoreCharge returns a new model instance from an encoded buffer
func NewCoreCharge(buf []byte, enctype EncodingType) (*CoreCharge, error) {
	var result CoreCharge
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
