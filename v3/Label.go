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
	ID          string     `gorm:"primaryKey;not null;column:id" json:"id"`
	CreatedDate time.Time  `gorm:"column:createdDate;not null;column:createdDate" json:"createdDate"`
	UpdatedDate *time.Time `gorm:"column:updatedDate;column:updatedDate" json:"updatedDate"`
	Meta        *Meta      `gorm:"type:json;embedded;serializer:json;column:meta;not null;column:meta" json:"meta,omitempty"` // the metadata about the most recent change to the row
	Metadata    any        `gorm:"type:json;serializer:json;column:metadata" json:"metadata,omitempty"`                       // metadata reserved for customers to control
	CompanyID   string     `gorm:"not null;column:companyId" json:"companyId"`
	LocationID  string     `gorm:"not null;column:locationId" json:"locationId"`

	Color  string          `gorm:"not null;column:color" json:"color"`
	Entity LabelEntityEnum `gorm:"not null;column:entity" json:"entity"`
	Name   string          `gorm:"not null;column:name" json:"name"`
	Saved  bool            `gorm:"not null;column:saved" json:"saved"`
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
