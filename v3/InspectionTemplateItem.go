// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	datatypes "github.com/shopmonkeyus/go-datamodel/datatypes"
)

type InspectionTemplateItemStatusEnum string

const (
	InspectionTemplateItemStatusGreen         InspectionTemplateItemStatusEnum = "Green"
	InspectionTemplateItemStatusYellow        InspectionTemplateItemStatusEnum = "Yellow"
	InspectionTemplateItemStatusRed           InspectionTemplateItemStatusEnum = "Red"
	InspectionTemplateItemStatusNotApplicable InspectionTemplateItemStatusEnum = "NotApplicable"
	InspectionTemplateItemStatus              InspectionTemplateItemStatusEnum = ""
)

type InspectionTemplateItem struct {
	ID          string              `gorm:"primaryKey;not null;column:id" json:"id"`
	CreatedDate datatypes.DateTime  `gorm:"column:createdDate;not null;column:createdDate" json:"createdDate"`
	UpdatedDate *datatypes.DateTime `gorm:"column:updatedDate;column:updatedDate" json:"updatedDate"`
	Meta        datatypes.Meta      `gorm:"column:meta;not null;column:meta" json:"meta,omitempty"`    // the metadata about the most recent change to the row
	Metadata    *datatypes.JSON     `gorm:"column:metadata;column:metadata" json:"metadata,omitempty"` // metadata reserved for customers to control
	CompanyID   string              `gorm:"not null;column:companyId" json:"companyId"`
	LocationID  string              `gorm:"not null;column:locationId" json:"locationId"`

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

// String returns a string representation as JSON for this model
func (m *InspectionTemplateItem) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewInspectionTemplateItem returns a new model instance from an encoded buffer
func NewInspectionTemplateItem(buf []byte) (*InspectionTemplateItem, error) {
	var result InspectionTemplateItem
	err := json.Unmarshal(buf, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// NewInspectionTemplateItemFromChangeEvent returns a new model instance from an encoded buffer as change event
func NewInspectionTemplateItemFromChangeEvent(buf []byte, gzip bool) (*datatypes.ChangeEvent[InspectionTemplateItem], error) {
	var result datatypes.ChangeEvent[InspectionTemplateItem]
	var decompressed = buf
	if gzip {
		dec, err := datatypes.Gunzip(buf)
		if err != nil {
			return nil, err
		}
		decompressed = dec
	}
	err := json.Unmarshal(decompressed, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
