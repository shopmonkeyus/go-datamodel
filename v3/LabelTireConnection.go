// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	mapstructure "github.com/mitchellh/mapstructure"
	"time"
)

type LabelTireConnection struct {
	CreatedDate time.Time  `gorm:"column:createdDate;not null;column:createdDate" json:"createdDate"`
	UpdatedDate *time.Time `gorm:"column:updatedDate;column:updatedDate" json:"updatedDate"`
	Meta        *Meta      `gorm:"type:json;embedded;serializer:json;column:meta;not null;column:meta" json:"meta,omitempty"` // the metadata about the most recent change to the row
	Metadata    any        `gorm:"type:json;serializer:json;column:metadata" json:"metadata,omitempty"`                       // metadata reserved for customers to control
	CompanyID   string     `gorm:"not null;column:companyId" json:"companyId"`

	LabelID string `gorm:"not null;column:labelId" json:"labelId"`
	TireID  string `gorm:"not null;column:tireId" json:"tireId"`
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
