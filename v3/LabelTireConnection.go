// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	mapstructure "github.com/mitchellh/mapstructure"
	"time"
)

type LabelTireConnection struct {
	CompanyId   string     `gorm:"not null" json:"companyId"`
	CreatedDate time.Time  `gorm:"column:createdDate;not null" json:"createdDate"`
	LabelId     string     `gorm:"not null" json:"labelId"`
	Meta        *Meta      `gorm:"type:json;embedded;column:meta;not null" json:"meta,omitempty"` // the metadata about the most recent change to the row
	Metadata    any        `gorm:"type:json" json:"metadata,omitempty"`                           // metadata reserved for customers to control
	TireId      string     `gorm:"not null" json:"tireId"`
	UpdatedDate *time.Time `gorm:"column:updatedDate" json:"updatedDate"`
}

var _ Model = (*LabelTireConnection)(nil)

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *LabelTireConnection) TableName() string {
	return "label_tire_connection"
}

func (m *LabelTireConnection) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewLabelTireConnection returns a new model instance from a json key/value map
func NewLabelTireConnection(kv map[string]any) (*LabelTireConnection, error) {
	var result LabelTireConnection
	err := mapstructure.Decode(kv, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
