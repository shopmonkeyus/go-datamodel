// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	mapstructure "github.com/mitchellh/mapstructure"
	"time"
)

type LabelCannedServiceFeeConnection struct {
	CannedServiceFeeId string     `gorm:"not null" json:"cannedServiceFeeId"`
	CompanyId          string     `gorm:"not null" json:"companyId"`
	CreatedDate        time.Time  `gorm:"column:createdDate;not null" json:"createdDate"`
	LabelId            string     `gorm:"not null" json:"labelId"`
	Meta               *Meta      `gorm:"type:json;embedded;column:meta;not null" json:"meta,omitempty"` // the metadata about the most recent change to the row
	Metadata           any        `gorm:"type:json" json:"metadata,omitempty"`                           // metadata reserved for customers to control
	UpdatedDate        *time.Time `gorm:"column:updatedDate" json:"updatedDate"`
}

var _ Model = (*LabelCannedServiceFeeConnection)(nil)

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *LabelCannedServiceFeeConnection) TableName() string {
	return "label_canned_service_fee_connection"
}

func (m *LabelCannedServiceFeeConnection) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewLabelCannedServiceFeeConnection returns a new model instance from a json key/value map
func NewLabelCannedServiceFeeConnection(kv map[string]any) (*LabelCannedServiceFeeConnection, error) {
	var result LabelCannedServiceFeeConnection
	err := mapstructure.Decode(kv, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
