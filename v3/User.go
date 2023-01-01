// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	datatypes "github.com/shopmonkeyus/go-datamodel/datatypes"
)

// user is a model for users that have access to the system
type User struct {
	ID           string              `gorm:"primaryKey;not null;column:id" json:"id"`
	CreatedDate  datatypes.DateTime  `gorm:"column:createdDate;not null;column:createdDate" json:"createdDate"`
	UpdatedDate  *datatypes.DateTime `gorm:"column:updatedDate;column:updatedDate" json:"updatedDate"`
	Meta         datatypes.Meta      `gorm:"column:meta;not null;column:meta" json:"meta,omitempty"`    // the metadata about the most recent change to the row
	Metadata     *datatypes.JSON     `gorm:"column:metadata;column:metadata" json:"metadata,omitempty"` // metadata reserved for customers to control
	CompanyID    string              `gorm:"not null;column:companyId" json:"companyId"`
	LocationID   *string             `gorm:"column:locationId" json:"locationId"`
	CustomFields datatypes.JSON      `gorm:"column:customFields" json:"customFields"` // custom field values

	FirstName *string `gorm:"column:firstName" json:"firstName"`
	LastName  *string `gorm:"column:lastName" json:"lastName"`
}

var _ Model = (*User)(nil)

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *User) TableName() string {
	return "user"
}

// String returns a string representation as JSON for this model
func (m *User) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewUser returns a new model instance from an encoded buffer
func NewUser(buf []byte) (*User, error) {
	var result User
	err := json.Unmarshal(buf, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// NewUserFromChangeEvent returns a new model instance from an encoded buffer as change event
func NewUserFromChangeEvent(buf []byte, gzip bool) (*datatypes.ChangeEvent[User], error) {
	var result datatypes.ChangeEvent[User]
	var decompressed = buf
	if gzip {
		dec, err := datatypes.Gunzip(buf)
		if err != nil {
			return nil, err
		}
		decompressed = dec
	}
	err := json.Unmarshal(decompressed, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
