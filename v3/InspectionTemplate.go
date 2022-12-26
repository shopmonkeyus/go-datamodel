// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	mapstructure "github.com/mitchellh/mapstructure"
	"time"
)

type InspectionTemplate struct {
	CompanyId     string     `gorm:"not null" json:"companyId"`
	ID            string     `gorm:"primaryKey;not null" json:"id"`
	CreatedById   *string    `json:"createdById"`
	CreatedDate   time.Time  `gorm:"column:createdDate;not null" json:"createdDate"`
	Deleted       bool       `gorm:"not null" json:"deleted"` // if the record has been deleted
	DeletedDate   *time.Time `json:"deletedDate"`             // the date that the record was deleted or null if not deleted
	DeletedReason *string    `json:"deletedReason"`           // the reason that the record was deleted
	DeletedUserId *string    `json:"deletedUserId"`           // the user that deleted the record or null if not deleted
	LocationId    string     `gorm:"not null" json:"locationId"`
	Meta          *Meta      `gorm:"type:json;embedded;column:meta;not null" json:"meta,omitempty"` // the metadata about the most recent change to the row
	Metadata      any        `gorm:"type:json" json:"metadata,omitempty"`                           // metadata reserved for customers to control
	Name          string     `gorm:"not null" json:"name"`
	UpdatedDate   *time.Time `gorm:"column:updatedDate" json:"updatedDate"`
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

// NewInspectionTemplate returns a new model instance from a json key/value map
func NewInspectionTemplate(kv map[string]any) (*InspectionTemplate, error) {
	var result InspectionTemplate
	err := mapstructure.Decode(kv, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
