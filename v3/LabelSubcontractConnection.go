// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	mapstructure "github.com/mitchellh/mapstructure"
	"time"
)

type LabelSubcontractConnection struct {
	CompanyId     string     `gorm:"not null" json:"companyId"`
	CreatedDate   time.Time  `gorm:"column:createdDate;not null" json:"createdDate"`
	LabelId       string     `gorm:"not null" json:"labelId"`
	Meta          *Meta      `gorm:"type:json;embedded;column:meta;not null" json:"meta,omitempty"` // the metadata about the most recent change to the row
	Metadata      any        `gorm:"type:json" json:"metadata,omitempty"`                           // metadata reserved for customers to control
	SubcontractId string     `gorm:"not null" json:"subcontractId"`
	UpdatedDate   *time.Time `gorm:"column:updatedDate" json:"updatedDate"`
}

var _ Model = (*LabelSubcontractConnection)(nil)

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *LabelSubcontractConnection) TableName() string {
	return "label_subcontract_connection"
}

func (m *LabelSubcontractConnection) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewLabelSubcontractConnection returns a new model instance from a json key/value map
func NewLabelSubcontractConnection(kv map[string]any) (*LabelSubcontractConnection, error) {
	var result LabelSubcontractConnection
	err := mapstructure.Decode(kv, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
