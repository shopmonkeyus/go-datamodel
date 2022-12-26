// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	"time"
)

type PaymentTerm struct {
	ID              string     `gorm:"primaryKey;not null" json:"id"`
	CompanyId       string     `gorm:"not null" json:"companyId"`
	CreatedDate     time.Time  `gorm:"column:createdDate;not null" json:"createdDate"`
	CustomerDefault bool       `gorm:"not null" json:"customerDefault"` // if it is the default for new customers
	DueInDays       float64    `gorm:"not null" json:"dueInDays"`
	Editable        bool       `gorm:"not null" json:"editable"`     // if it is not a built-in payment term (eg. On Receipt)
	FleetDefault    bool       `gorm:"not null" json:"fleetDefault"` // if it is the default for new fleets
	LocationId      string     `gorm:"not null" json:"locationId"`
	Meta            *Meta      `gorm:"embedded;column:meta;not null" json:"meta,omitempty"`         // the metadata about the most recent change to the row
	Metadata        any        `gorm:"embedded;column:metadata;not null" json:"metadata,omitempty"` // metadata reserved for customers to control
	Name            string     `gorm:"not null" json:"name"`
	Note            string     `gorm:"not null" json:"note"`
	UpdatedDate     *time.Time `gorm:"column:updatedDate" json:"updatedDate"`
}

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *PaymentTerm) TableName() string {
	return "payment_term"
}

func (m *PaymentTerm) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}
