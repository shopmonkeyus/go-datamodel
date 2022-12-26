// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	"time"
)

type LabelCannedServiceSubcontractConnection struct {
	CreatedDate time.Time  `gorm:"column:createdDate;not null;column:createdDate" json:"createdDate"`
	UpdatedDate *time.Time `gorm:"column:updatedDate;column:updatedDate" json:"updatedDate"`
	Meta        *Meta      `gorm:"type:json;embedded;serializer:json;column:meta;not null;column:meta" json:"meta,omitempty"` // the metadata about the most recent change to the row
	Metadata    any        `gorm:"type:json;serializer:json;column:metadata" json:"metadata,omitempty"`                       // metadata reserved for customers to control
	CompanyID   string     `gorm:"not null;column:companyId" json:"companyId"`

	CannedServiceSubcontractID string `gorm:"not null;column:cannedServiceSubcontractId" json:"cannedServiceSubcontractId"`
	LabelID                    string `gorm:"not null;column:labelId" json:"labelId"`
}

var _ Model = (*LabelCannedServiceSubcontractConnection)(nil)

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *LabelCannedServiceSubcontractConnection) TableName() string {
	return "label_canned_service_subcontract_connection"
}

func (m *LabelCannedServiceSubcontractConnection) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewLabelCannedServiceSubcontractConnection returns a new model instance from a json key/value map
func NewLabelCannedServiceSubcontractConnection(buf []byte) (*LabelCannedServiceSubcontractConnection, error) {
	var result LabelCannedServiceSubcontractConnection
	err := json.Unmarshal(buf, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
