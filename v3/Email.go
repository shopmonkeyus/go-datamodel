// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	datatypes "github.com/shopmonkeyus/go-datamodel/datatypes"
)

type Email struct {
	ID          string              `gorm:"primaryKey;not null;column:id" json:"id"`
	CreatedDate datatypes.DateTime  `gorm:"column:createdDate;not null;column:createdDate" json:"createdDate"`
	UpdatedDate *datatypes.DateTime `gorm:"column:updatedDate;column:updatedDate" json:"updatedDate"`
	Meta        *datatypes.JSON     `gorm:"column:meta;not null;column:meta" json:"meta,omitempty"`    // the metadata about the most recent change to the row
	Metadata    *datatypes.JSON     `gorm:"column:metadata;column:metadata" json:"metadata,omitempty"` // metadata reserved for customers to control
	CompanyID   string              `gorm:"not null;column:companyId" json:"companyId"`
	LocationID  string              `gorm:"not null;column:locationId" json:"locationId"`

	CustomerID string `gorm:"not null;column:customerId" json:"customerId"`
	Email      string `gorm:"not null;column:email" json:"email"`
	Primary    bool   `gorm:"not null;column:primary" json:"primary"`
	Subscribed bool   `gorm:"not null;column:subscribed" json:"subscribed"`
}

var _ Model = (*Email)(nil)

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *Email) TableName() string {
	return "email"
}

// PrimaryKeys returns an array of the primary keys for this model
func (m *Email) PrimaryKeys() []string {
	return []string{"id"}
}

// String returns a string representation as JSON for this model
func (m *Email) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewEmail returns a new model instance from an encoded buffer
func NewEmail(buf []byte) (*Email, error) {
	var result Email
	err := json.Unmarshal(buf, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// NewEmailFromChangeEvent returns a new model instance from an encoded buffer as change event
func NewEmailFromChangeEvent(buf []byte, gzip bool) (*datatypes.ChangeEvent[Email], error) {
	var result datatypes.ChangeEvent[Email]
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
