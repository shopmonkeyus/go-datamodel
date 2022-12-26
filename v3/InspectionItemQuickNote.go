// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	codec "github.com/hashicorp/go-msgpack/v2/codec"
	"time"
)

type InspectionItemQuickNote struct {
	ID          string     `gorm:"primaryKey;not null;column:id" json:"id"`
	CreatedDate time.Time  `gorm:"column:createdDate;not null;column:createdDate" json:"createdDate"`
	UpdatedDate *time.Time `gorm:"column:updatedDate;column:updatedDate" json:"updatedDate"`
	Meta        *Meta      `gorm:"type:json;embedded;serializer:json;column:meta;not null;column:meta" json:"meta,omitempty"` // the metadata about the most recent change to the row
	Metadata    any        `gorm:"type:json;serializer:json;column:metadata" json:"metadata,omitempty"`                       // metadata reserved for customers to control
	CompanyID   string     `gorm:"not null;column:companyId" json:"companyId"`
	LocationID  string     `gorm:"not null;column:locationId" json:"locationId"`

	CreatedByID *string `gorm:"column:createdById" json:"createdById"`
	Label       string  `gorm:"not null;column:label" json:"label"`
	Note        string  `gorm:"not null;column:note" json:"note"`
}

var _ Model = (*InspectionItemQuickNote)(nil)

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *InspectionItemQuickNote) TableName() string {
	return "inspection_item_quick_note"
}

func (m *InspectionItemQuickNote) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewInspectionItemQuickNote returns a new model instance from an encoded buffer
func NewInspectionItemQuickNote(buf []byte, enctype EncodingType) (*InspectionItemQuickNote, error) {
	var result InspectionItemQuickNote
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
