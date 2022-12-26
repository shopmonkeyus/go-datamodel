// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	codec "github.com/hashicorp/go-msgpack/v2/codec"
	"time"
)

type LabelCannedServiceFeeConnection struct {
	CreatedDate time.Time  `gorm:"column:createdDate;not null;column:createdDate" json:"createdDate"`
	UpdatedDate *time.Time `gorm:"column:updatedDate;column:updatedDate" json:"updatedDate"`
	Meta        *Meta      `gorm:"type:json;embedded;serializer:json;column:meta;not null;column:meta" json:"meta,omitempty"` // the metadata about the most recent change to the row
	Metadata    any        `gorm:"type:json;serializer:json;column:metadata" json:"metadata,omitempty"`                       // metadata reserved for customers to control
	CompanyID   string     `gorm:"not null;column:companyId" json:"companyId"`

	CannedServiceFeeID string `gorm:"not null;column:cannedServiceFeeId" json:"cannedServiceFeeId"`
	LabelID            string `gorm:"not null;column:labelId" json:"labelId"`
}

var _ Model = (*LabelCannedServiceFeeConnection)(nil)

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *LabelCannedServiceFeeConnection) TableName() string {
	return "label_canned_service_fee_connection"
}

func (m *LabelCannedServiceFeeConnection) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewLabelCannedServiceFeeConnection returns a new model instance from an encoded buffer
func NewLabelCannedServiceFeeConnection(buf []byte, enctype EncodingType) (*LabelCannedServiceFeeConnection, error) {
	var result LabelCannedServiceFeeConnection
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
