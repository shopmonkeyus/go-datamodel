// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	mapstructure "github.com/mitchellh/mapstructure"
	"time"
)

// user is a model for users that have access to the system
type User struct {
	ID           string     `gorm:"primaryKey;not null" json:"id"`
	CompanyId    string     `gorm:"not null" json:"companyId"`
	CreatedDate  time.Time  `gorm:"column:createdDate;not null" json:"createdDate"`
	CustomFields any        `gorm:"type:json" json:"customFields"` // custom field values
	FirstName    *string    `json:"firstName"`
	LastName     *string    `json:"lastName"`
	LocationId   *string    `json:"locationId"`
	Meta         *Meta      `gorm:"type:json;embedded;column:meta;not null" json:"meta,omitempty"` // the metadata about the most recent change to the row
	Metadata     any        `gorm:"type:json" json:"metadata,omitempty"`                           // metadata reserved for customers to control
	UpdatedDate  *time.Time `gorm:"column:updatedDate" json:"updatedDate"`
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
func NewUser(kv map[string]any) (*User, error) {
	var result User
	err := mapstructure.Decode(kv, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
