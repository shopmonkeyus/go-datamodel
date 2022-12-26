// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	mapstructure "github.com/mitchellh/mapstructure"
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
	ID          string     `gorm:"primaryKey;not null;column:id" json:"id"`
	CreatedDate time.Time  `gorm:"column:createdDate;not null;column:createdDate" json:"createdDate"`
	UpdatedDate *time.Time `gorm:"column:updatedDate;column:updatedDate" json:"updatedDate"`
	Meta        *Meta      `gorm:"type:json;embedded;serializer:json;column:meta;not null;column:meta" json:"meta,omitempty"` // the metadata about the most recent change to the row
	Metadata    any        `gorm:"type:json;serializer:json;column:metadata" json:"metadata,omitempty"`                       // metadata reserved for customers to control
	CompanyID   string     `gorm:"not null;column:companyId" json:"companyId"`
	LocationID  string     `gorm:"not null;column:locationId" json:"locationId"`

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

// NewInspectionTemplateItem returns a new model instance from a json key/value map
func NewInspectionTemplateItem(kv map[string]any) (*InspectionTemplateItem, error) {
	var result InspectionTemplateItem
	err := mapstructure.Decode(kv, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
