// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	codec "github.com/hashicorp/go-msgpack/v2/codec"
	datatypes "github.com/shopmonkeyus/go-datamodel/datatypes"
	"time"
)

type LabelCustomerConnection struct {
	CreatedDate *time.Time      `gorm:"column:createdDate;column:createdDate" json:"createdDate"`
	UpdatedDate *time.Time      `gorm:"column:updatedDate;column:updatedDate" json:"updatedDate"`
	Meta        datatypes.Meta  `gorm:"column:meta;not null;column:meta" json:"meta,omitempty"`    // the metadata about the most recent change to the row
	Metadata    *datatypes.JSON `gorm:"column:metadata;column:metadata" json:"metadata,omitempty"` // metadata reserved for customers to control
	CompanyID   string          `gorm:"not null;column:companyId" json:"companyId"`

	CustomerID string `gorm:"not null;column:customerId" json:"customerId"`
	LabelID    string `gorm:"not null;column:labelId" json:"labelId"`
}

var _ Model = (*LabelCustomerConnection)(nil)

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *LabelCustomerConnection) TableName() string {
	return "label_customer_connection"
}

func (m *LabelCustomerConnection) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewLabelCustomerConnection returns a new model instance from an encoded buffer
func NewLabelCustomerConnection(buf []byte, enctype EncodingType) (*LabelCustomerConnection, error) {
	var result LabelCustomerConnection
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
