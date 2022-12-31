// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	codec "github.com/hashicorp/go-msgpack/v2/codec"
	datatypes "github.com/shopmonkeyus/go-datamodel/datatypes"
	"time"
)

type LabelCannedServiceTireConnection struct {
	CreatedDate time.Time       `gorm:"column:createdDate;not null;column:createdDate" json:"createdDate"`
	UpdatedDate *time.Time      `gorm:"column:updatedDate;column:updatedDate" json:"updatedDate"`
	Meta        datatypes.Meta  `gorm:"column:meta;not null;column:meta" json:"meta,omitempty"`    // the metadata about the most recent change to the row
	Metadata    *datatypes.JSON `gorm:"column:metadata;column:metadata" json:"metadata,omitempty"` // metadata reserved for customers to control
	CompanyID   string          `gorm:"not null;column:companyId" json:"companyId"`

	CannedServiceTireID string `gorm:"not null;column:cannedServiceTireId" json:"cannedServiceTireId"`
	LabelID             string `gorm:"not null;column:labelId" json:"labelId"`
}

var _ Model = (*LabelCannedServiceTireConnection)(nil)

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *LabelCannedServiceTireConnection) TableName() string {
	return "label_canned_service_tire_connection"
}

func (m *LabelCannedServiceTireConnection) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewLabelCannedServiceTireConnection returns a new model instance from an encoded buffer
func NewLabelCannedServiceTireConnection(buf []byte, enctype EncodingType) (*LabelCannedServiceTireConnection, error) {
	var result LabelCannedServiceTireConnection
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
