// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	codec "github.com/hashicorp/go-msgpack/v2/codec"
	datatypes "github.com/shopmonkeyus/go-datamodel/datatypes"
	"time"
)

type CustomerLocationConnection struct {
	CreatedDate time.Time       `gorm:"column:createdDate;not null;column:createdDate" json:"createdDate"`
	UpdatedDate *time.Time      `gorm:"column:updatedDate;column:updatedDate" json:"updatedDate"`
	Meta        datatypes.Meta  `gorm:"column:meta;not null;column:meta" json:"meta,omitempty"`    // the metadata about the most recent change to the row
	Metadata    *datatypes.JSON `gorm:"column:metadata;column:metadata" json:"metadata,omitempty"` // metadata reserved for customers to control
	CompanyID   string          `gorm:"not null;column:companyId" json:"companyId"`
	LocationID  string          `gorm:"not null;column:locationId" json:"locationId"`

	CustomerID string `gorm:"not null;column:customerId" json:"customerId"`
}

var _ Model = (*CustomerLocationConnection)(nil)

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *CustomerLocationConnection) TableName() string {
	return "customer_location_connection"
}

func (m *CustomerLocationConnection) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewCustomerLocationConnection returns a new model instance from an encoded buffer
func NewCustomerLocationConnection(buf []byte, enctype EncodingType) (*CustomerLocationConnection, error) {
	var result CustomerLocationConnection
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
