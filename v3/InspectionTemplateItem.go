// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	codec "github.com/hashicorp/go-msgpack/v2/codec"
	datatypes "github.com/shopmonkeyus/go-datamodel/datatypes"
	"time"
)

type InspectionTemplateItemStatusEnum string

const (
	InspectionTemplateItemStatusGreen         InspectionTemplateItemStatusEnum = "Green"
	InspectionTemplateItemStatusYellow                                         = "Yellow"
	InspectionTemplateItemStatusRed                                            = "Red"
	InspectionTemplateItemStatusNotApplicable                                  = "NotApplicable"
)

type InspectionTemplateItem struct {
	ID          string          `gorm:"primaryKey;not null;column:id" json:"id"`
	CreatedDate time.Time       `gorm:"column:createdDate;not null;column:createdDate" json:"createdDate"`
	UpdatedDate *time.Time      `gorm:"column:updatedDate;column:updatedDate" json:"updatedDate"`
	Meta        datatypes.Meta  `gorm:"column:meta;not null;column:meta" json:"meta,omitempty"`    // the metadata about the most recent change to the row
	Metadata    *datatypes.JSON `gorm:"column:metadata;column:metadata" json:"metadata,omitempty"` // metadata reserved for customers to control
	CompanyID   string          `gorm:"not null;column:companyId" json:"companyId"`
	LocationID  string          `gorm:"not null;column:locationId" json:"locationId"`

	InspectionTemplateID string                            `gorm:"not null;column:inspectionTemplateId" json:"inspectionTemplateId"`
	Message              string                            `gorm:"not null;column:message" json:"message"`
	Name                 string                            `gorm:"not null;column:name" json:"name"`
	Ordinal              float64                           `gorm:"not null;column:ordinal" json:"ordinal"`
	Status               *InspectionTemplateItemStatusEnum `gorm:"column:status" json:"status"`
}

var _ Model = (*InspectionTemplateItem)(nil)

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *InspectionTemplateItem) TableName() string {
	return "inspection_template_item"
}

func (m *InspectionTemplateItem) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewInspectionTemplateItem returns a new model instance from an encoded buffer
func NewInspectionTemplateItem(buf []byte, enctype EncodingType) (*InspectionTemplateItem, error) {
	var result InspectionTemplateItem
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
