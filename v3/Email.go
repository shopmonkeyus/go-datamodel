// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	"time"
)

type Email struct {
	ID          string     `gorm:"primaryKey;not null" json:"id"`
	CompanyId   string     `gorm:"not null" json:"companyId"`
	CreatedDate time.Time  `gorm:"column:createdDate;not null" json:"createdDate"`
	CustomerId  string     `gorm:"not null" json:"customerId"`
	Email       string     `gorm:"not null" json:"email"`
	LocationId  string     `gorm:"not null" json:"locationId"`
	Meta        *Meta      `gorm:"embedded;column:meta;not null" json:"meta,omitempty"`         // the metadata about the most recent change to the row
	Metadata    any        `gorm:"embedded;column:metadata;not null" json:"metadata,omitempty"` // metadata reserved for customers to control
	Primary     bool       `gorm:"not null" json:"primary"`
	Subscribed  bool       `gorm:"not null" json:"subscribed"`
	UpdatedDate *time.Time `gorm:"column:updatedDate" json:"updatedDate"`
}

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *Email) TableName() string {
	return "email"
}

func (m *Email) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}
