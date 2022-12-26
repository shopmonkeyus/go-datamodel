// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	mapstructure "github.com/mitchellh/mapstructure"
	"time"
)

type LabelEntityEnum string

const (
	LabelEntityCannedService LabelEntityEnum = "CannedService"
	LabelEntityCustomer                      = "Customer"
	LabelEntityFee                           = "Fee"
	LabelEntityLabor                         = "Labor"
	LabelEntityOrder                         = "Order"
	LabelEntityPart                          = "Part"
	LabelEntityService                       = "Service"
	LabelEntitySubcontract                   = "Subcontract"
	LabelEntityTire                          = "Tire"
	LabelEntityVehicle                       = "Vehicle"
)

type Label struct {
	ID          string          `gorm:"primaryKey;not null" json:"id"`
	Color       string          `gorm:"not null" json:"color"`
	CompanyId   string          `gorm:"not null" json:"companyId"`
	CreatedDate time.Time       `gorm:"column:createdDate;not null" json:"createdDate"`
	Entity      LabelEntityEnum `gorm:"not null" json:"entity"`
	LocationId  string          `gorm:"not null" json:"locationId"`
	Meta        *Meta           `gorm:"type:json;embedded;column:meta;not null" json:"meta,omitempty"` // the metadata about the most recent change to the row
	Metadata    any             `gorm:"type:json" json:"metadata,omitempty"`                           // metadata reserved for customers to control
	Name        string          `gorm:"not null" json:"name"`
	Saved       bool            `gorm:"not null" json:"saved"`
	UpdatedDate *time.Time      `gorm:"column:updatedDate" json:"updatedDate"`
}

var _ Model = (*Label)(nil)

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *Label) TableName() string {
	return "label"
}

func (m *Label) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewLabel returns a new model instance from a json key/value map
func NewLabel(kv map[string]any) (*Label, error) {
	var result Label
	err := mapstructure.Decode(kv, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
