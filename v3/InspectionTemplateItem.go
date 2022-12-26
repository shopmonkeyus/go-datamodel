// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
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
	CreatedDate          time.Time                         `gorm:"column:createdDate;not null" json:"createdDate"`
	ID                   string                            `gorm:"primaryKey;not null" json:"id"`
	InspectionTemplateId string                            `gorm:"not null" json:"inspectionTemplateId"`
	LocationId           string                            `gorm:"not null" json:"locationId"`
	Message              string                            `gorm:"not null" json:"message"`
	Meta                 *Meta                             `gorm:"embedded;column:meta;not null" json:"meta,omitempty"`         // the metadata about the most recent change to the row
	Metadata             any                               `gorm:"embedded;column:metadata;not null" json:"metadata,omitempty"` // metadata reserved for customers to control
	Name                 string                            `gorm:"not null" json:"name"`
	Ordinal              float64                           `gorm:"not null" json:"ordinal"`
	Status               *InspectionTemplateItemStatusEnum `json:"status"`
	UpdatedDate          *time.Time                        `gorm:"column:updatedDate" json:"updatedDate"`
}

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *InspectionTemplateItem) TableName() string {
	return "inspection_template_item"
}

func (m *InspectionTemplateItem) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}
