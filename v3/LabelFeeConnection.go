// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	mapstructure "github.com/mitchellh/mapstructure"
	"time"
)

type LabelFeeConnection struct {
	CompanyId   string     `gorm:"not null" json:"companyId"`
	CreatedDate time.Time  `gorm:"column:createdDate;not null" json:"createdDate"`
	FeeId       string     `gorm:"not null" json:"feeId"`
	LabelId     string     `gorm:"not null" json:"labelId"`
	Meta        *Meta      `gorm:"type:json;embedded;column:meta;not null" json:"meta,omitempty"` // the metadata about the most recent change to the row
	Metadata    any        `gorm:"type:json" json:"metadata,omitempty"`                           // metadata reserved for customers to control
	UpdatedDate *time.Time `gorm:"column:updatedDate" json:"updatedDate"`
}

var _ Model = (*LabelFeeConnection)(nil)

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *LabelFeeConnection) TableName() string {
	return "label_fee_connection"
}

func (m *LabelFeeConnection) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewLabelFeeConnection returns a new model instance from a json key/value map
func NewLabelFeeConnection(kv map[string]any) (*LabelFeeConnection, error) {
	var result LabelFeeConnection
	err := mapstructure.Decode(kv, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
