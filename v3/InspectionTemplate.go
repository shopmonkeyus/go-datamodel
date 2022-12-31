// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	codec "github.com/hashicorp/go-msgpack/v2/codec"
	datatypes "github.com/shopmonkeyus/go-datamodel/datatypes"
	"time"
)

type InspectionTemplate struct {
	ID          string          `gorm:"primaryKey;not null;column:id" json:"id"`
	CreatedDate time.Time       `gorm:"column:createdDate;not null;column:createdDate" json:"createdDate"`
	UpdatedDate *time.Time      `gorm:"column:updatedDate;column:updatedDate" json:"updatedDate"`
	Meta        datatypes.Meta  `gorm:"column:meta;not null;column:meta" json:"meta,omitempty"`    // the metadata about the most recent change to the row
	Metadata    *datatypes.JSON `gorm:"column:metadata;column:metadata" json:"metadata,omitempty"` // metadata reserved for customers to control
	CompanyID   string          `gorm:"not null;column:companyId" json:"companyId"`
	LocationID  string          `gorm:"not null;column:locationId" json:"locationId"`

	CreatedByID   *string    `gorm:"column:createdById" json:"createdById"`
	Deleted       bool       `gorm:"not null;column:deleted" json:"deleted"`    // if the record has been deleted
	DeletedDate   *time.Time `gorm:"column:deletedDate" json:"deletedDate"`     // the date that the record was deleted or null if not deleted
	DeletedReason *string    `gorm:"column:deletedReason" json:"deletedReason"` // the reason that the record was deleted
	DeletedUserID *string    `gorm:"column:deletedUserId" json:"deletedUserId"` // the user that deleted the record or null if not deleted
	Name          string     `gorm:"not null;column:name" json:"name"`
}

var _ Model = (*InspectionTemplate)(nil)

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *InspectionTemplate) TableName() string {
	return "inspection_template"
}

func (m *InspectionTemplate) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewInspectionTemplate returns a new model instance from an encoded buffer
func NewInspectionTemplate(buf []byte, enctype EncodingType) (*InspectionTemplate, error) {
	var result InspectionTemplate
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
