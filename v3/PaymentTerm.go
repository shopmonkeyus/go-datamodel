// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	mapstructure "github.com/mitchellh/mapstructure"
	"time"
)

type PaymentTerm struct {
	ID          string     `gorm:"primaryKey;not null;column:id" json:"id"`
	CreatedDate time.Time  `gorm:"column:createdDate;not null;column:createdDate" json:"createdDate"`
	UpdatedDate *time.Time `gorm:"column:updatedDate;column:updatedDate" json:"updatedDate"`
	Meta        *Meta      `gorm:"type:json;embedded;serializer:json;column:meta;not null;column:meta" json:"meta,omitempty"` // the metadata about the most recent change to the row
	Metadata    any        `gorm:"type:json;serializer:json;column:metadata" json:"metadata,omitempty"`                       // metadata reserved for customers to control
	CompanyID   string     `gorm:"not null;column:companyId" json:"companyId"`
	LocationID  string     `gorm:"not null;column:locationId" json:"locationId"`

	CustomerDefault bool    `gorm:"not null;column:customerDefault" json:"customerDefault"` // if it is the default for new customers
	DueInDays       float64 `gorm:"not null;column:dueInDays" json:"dueInDays"`
	Editable        bool    `gorm:"not null;column:editable" json:"editable"`         // if it is not a built-in payment term (eg. On Receipt)
	FleetDefault    bool    `gorm:"not null;column:fleetDefault" json:"fleetDefault"` // if it is the default for new fleets
	Name            string  `gorm:"not null;column:name" json:"name"`
	Note            string  `gorm:"not null;column:note" json:"note"`
}

var _ Model = (*PaymentTerm)(nil)

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *PaymentTerm) TableName() string {
	return "payment_term"
}

func (m *PaymentTerm) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewPaymentTerm returns a new model instance from a json key/value map
func NewPaymentTerm(kv map[string]any) (*PaymentTerm, error) {
	var result PaymentTerm
	err := mapstructure.Decode(kv, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
