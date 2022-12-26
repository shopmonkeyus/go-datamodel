// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	mapstructure "github.com/mitchellh/mapstructure"
	"time"
)

type InspectionItemQuickNote struct {
	CompanyId   string     `gorm:"not null" json:"companyId"`
	CreatedById *string    `json:"createdById"`
	CreatedDate time.Time  `gorm:"column:createdDate;not null" json:"createdDate"`
	ID          string     `gorm:"primaryKey;not null" json:"id"`
	Label       string     `gorm:"not null" json:"label"`
	LocationId  string     `gorm:"not null" json:"locationId"`
	Meta        *Meta      `gorm:"type:json;embedded;column:meta;not null" json:"meta,omitempty"` // the metadata about the most recent change to the row
	Metadata    any        `gorm:"type:json" json:"metadata,omitempty"`                           // metadata reserved for customers to control
	Note        string     `gorm:"not null" json:"note"`
	UpdatedDate *time.Time `gorm:"column:updatedDate" json:"updatedDate"`
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

// NewInspectionItemQuickNote returns a new model instance from a json key/value map
func NewInspectionItemQuickNote(kv map[string]any) (*InspectionItemQuickNote, error) {
	var result InspectionItemQuickNote
	err := mapstructure.Decode(kv, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
