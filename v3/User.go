// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	"time"
)

// user is a model for users that have access to the system
type User struct {
	CompanyId    string     `gorm:"not null" json:"companyId"`
	CreatedDate  time.Time  `gorm:"column:createdDate;not null" json:"createdDate"`
	CustomFields any        `gorm:"not null" json:"customFields"` // custom field values
	FirstName    *string    `json:"firstName"`
	ID           string     `gorm:"primaryKey;not null" json:"id"`
	LastName     *string    `json:"lastName"`
	LocationId   *string    `json:"locationId"`
	Meta         *Meta      `gorm:"embedded;column:meta;not null" json:"meta,omitempty"`         // the metadata about the most recent change to the row
	Metadata     any        `gorm:"embedded;column:metadata;not null" json:"metadata,omitempty"` // metadata reserved for customers to control
	UpdatedDate  *time.Time `gorm:"column:updatedDate" json:"updatedDate"`
}

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *User) TableName() string {
	return "user"
}

func (m *User) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}
