// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	mapstructure "github.com/mitchellh/mapstructure"
	"time"
)

type LabelPartConnection struct {
	CreatedDate time.Time  `gorm:"column:createdDate;not null;column:createdDate" json:"createdDate"`
	UpdatedDate *time.Time `gorm:"column:updatedDate;column:updatedDate" json:"updatedDate"`
	Meta        *Meta      `gorm:"type:json;embedded;serializer:json;column:meta;not null;column:meta" json:"meta,omitempty"` // the metadata about the most recent change to the row
	Metadata    any        `gorm:"type:json;serializer:json;column:metadata" json:"metadata,omitempty"`                       // metadata reserved for customers to control
	CompanyID   string     `gorm:"not null;column:companyId" json:"companyId"`

	LabelID string `gorm:"not null;column:labelId" json:"labelId"`
	PartID  string `gorm:"not null;column:partId" json:"partId"`
}

var _ Model = (*LabelPartConnection)(nil)

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *LabelPartConnection) TableName() string {
	return "label_part_connection"
}

func (m *LabelPartConnection) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewLabelPartConnection returns a new model instance from a json key/value map
func NewLabelPartConnection(kv map[string]any) (*LabelPartConnection, error) {
	var result LabelPartConnection
	err := mapstructure.Decode(kv, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
