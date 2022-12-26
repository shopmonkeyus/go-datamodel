// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	mapstructure "github.com/mitchellh/mapstructure"
	"time"
)

type LabelPartConnection struct {
	CompanyId   string     `gorm:"not null" json:"companyId"`
	CreatedDate time.Time  `gorm:"column:createdDate;not null" json:"createdDate"`
	LabelId     string     `gorm:"not null" json:"labelId"`
	Meta        *Meta      `gorm:"type:json;embedded;column:meta;not null" json:"meta,omitempty"` // the metadata about the most recent change to the row
	Metadata    any        `gorm:"type:json" json:"metadata,omitempty"`                           // metadata reserved for customers to control
	PartId      string     `gorm:"not null" json:"partId"`
	UpdatedDate *time.Time `gorm:"column:updatedDate" json:"updatedDate"`
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
