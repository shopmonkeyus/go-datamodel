// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	mapstructure "github.com/mitchellh/mapstructure"
	"time"
)

type LabelCannedServiceLaborConnection struct {
	CreatedDate time.Time  `gorm:"column:createdDate;not null;column:createdDate" json:"createdDate"`
	UpdatedDate *time.Time `gorm:"column:updatedDate;column:updatedDate" json:"updatedDate"`
	Meta        *Meta      `gorm:"type:json;embedded;serializer:json;column:meta;not null;column:meta" json:"meta,omitempty"` // the metadata about the most recent change to the row
	Metadata    any        `gorm:"type:json;serializer:json;column:metadata" json:"metadata,omitempty"`                       // metadata reserved for customers to control
	CompanyID   string     `gorm:"not null;column:companyId" json:"companyId"`

	CannedServiceLaborID string `gorm:"not null;column:cannedServiceLaborId" json:"cannedServiceLaborId"`
	LabelID              string `gorm:"not null;column:labelId" json:"labelId"`
}

var _ Model = (*LabelCannedServiceLaborConnection)(nil)

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *LabelCannedServiceLaborConnection) TableName() string {
	return "label_canned_service_labor_connection"
}

func (m *LabelCannedServiceLaborConnection) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewLabelCannedServiceLaborConnection returns a new model instance from a json key/value map
func NewLabelCannedServiceLaborConnection(kv map[string]any) (*LabelCannedServiceLaborConnection, error) {
	var result LabelCannedServiceLaborConnection
	err := mapstructure.Decode(kv, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
