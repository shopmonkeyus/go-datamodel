// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	mapstructure "github.com/mitchellh/mapstructure"
	"time"
)

type LabelCustomerConnection struct {
	CompanyId   string     `gorm:"not null" json:"companyId"`
	CreatedDate *time.Time `gorm:"column:createdDate" json:"createdDate"`
	CustomerId  string     `gorm:"not null" json:"customerId"`
	LabelId     string     `gorm:"not null" json:"labelId"`
	Meta        *Meta      `gorm:"type:json;embedded;column:meta;not null" json:"meta,omitempty"` // the metadata about the most recent change to the row
	Metadata    any        `gorm:"type:json" json:"metadata,omitempty"`                           // metadata reserved for customers to control
	UpdatedDate *time.Time `gorm:"column:updatedDate" json:"updatedDate"`
}

var _ Model = (*LabelCustomerConnection)(nil)

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *LabelCustomerConnection) TableName() string {
	return "label_customer_connection"
}

func (m *LabelCustomerConnection) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewLabelCustomerConnection returns a new model instance from a json key/value map
func NewLabelCustomerConnection(kv map[string]any) (*LabelCustomerConnection, error) {
	var result LabelCustomerConnection
	err := mapstructure.Decode(kv, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
