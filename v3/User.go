// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	"time"
)

// user is a model for users that have access to the system
type User struct {
	ID           string     `gorm:"primaryKey;not null;column:id" json:"id"`
	CreatedDate  time.Time  `gorm:"column:createdDate;not null;column:createdDate" json:"createdDate"`
	UpdatedDate  *time.Time `gorm:"column:updatedDate;column:updatedDate" json:"updatedDate"`
	Meta         *Meta      `gorm:"type:json;embedded;serializer:json;column:meta;not null;column:meta" json:"meta,omitempty"` // the metadata about the most recent change to the row
	Metadata     any        `gorm:"type:json;serializer:json;column:metadata" json:"metadata,omitempty"`                       // metadata reserved for customers to control
	CompanyID    string     `gorm:"not null;column:companyId" json:"companyId"`
	LocationID   *string    `gorm:"column:locationId" json:"locationId"`
	CustomFields any        `gorm:"type:json;serializer:json;column:customFields" json:"customFields"` // custom field values

	FirstName *string `gorm:"column:firstName" json:"firstName"`
	LastName  *string `gorm:"column:lastName" json:"lastName"`
}

var _ Model = (*User)(nil)

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *User) TableName() string {
	return "user"
}

func (m *User) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewUser returns a new model instance from a json key/value map
func NewUser(buf []byte) (*User, error) {
	var result User
	err := json.Unmarshal(buf, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
