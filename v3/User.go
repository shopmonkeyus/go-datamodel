// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	codec "github.com/hashicorp/go-msgpack/v2/codec"
	datatypes "github.com/shopmonkeyus/go-datamodel/datatypes"
	"time"
)

// user is a model for users that have access to the system
type User struct {
	ID           string          `gorm:"primaryKey;not null;column:id" json:"id"`
	CreatedDate  time.Time       `gorm:"column:createdDate;not null;column:createdDate" json:"createdDate"`
	UpdatedDate  *time.Time      `gorm:"column:updatedDate;column:updatedDate" json:"updatedDate"`
	Meta         datatypes.Meta  `gorm:"column:meta;not null;column:meta" json:"meta,omitempty"`    // the metadata about the most recent change to the row
	Metadata     *datatypes.JSON `gorm:"column:metadata;column:metadata" json:"metadata,omitempty"` // metadata reserved for customers to control
	CompanyID    string          `gorm:"not null;column:companyId" json:"companyId"`
	LocationID   *string         `gorm:"column:locationId" json:"locationId"`
	CustomFields datatypes.JSON  `gorm:"column:customFields" json:"customFields"` // custom field values

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

// NewUser returns a new model instance from an encoded buffer
func NewUser(buf []byte, enctype EncodingType) (*User, error) {
	var result User
	var handle codec.Handle
	if enctype == JSONEncoding {
		handle = &jsonHandle
	} else {
		handle = &msgpackHandle
	}
	dec := codec.NewDecoderBytes(buf, handle)
	err := dec.Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
