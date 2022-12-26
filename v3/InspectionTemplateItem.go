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
	CompanyId            string                            `gorm:"not null" json:"companyId"`
	ID                   string                            `gorm:"primaryKey;not null" json:"id"`
	CreatedDate          time.Time                         `gorm:"column:createdDate;not null" json:"createdDate"`
	InspectionTemplateId string                            `gorm:"not null" json:"inspectionTemplateId"`
	LocationId           string                            `gorm:"not null" json:"locationId"`
	Message              string                            `gorm:"not null" json:"message"`
	Meta                 *Meta                             `gorm:"type:json;embedded;column:meta;not null" json:"meta,omitempty"` // the metadata about the most recent change to the row
	Metadata             any                               `gorm:"type:json" json:"metadata,omitempty"`                           // metadata reserved for customers to control
	Name                 string                            `gorm:"not null" json:"name"`
	Ordinal              float64                           `gorm:"not null" json:"ordinal"`
	Status               *InspectionTemplateItemStatusEnum `json:"status"`
	UpdatedDate          *time.Time                        `gorm:"column:updatedDate" json:"updatedDate"`
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
