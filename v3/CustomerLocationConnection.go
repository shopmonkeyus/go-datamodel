// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	datatypes "github.com/shopmonkeyus/go-datamodel/datatypes"
)

type CustomerLocationConnection struct {
	UpdatedDate *datatypes.DateTime `gorm:"column:updatedDate;column:updatedDate" json:"updatedDate"`

	CompanyID   string             `gorm:"not null;column:companyId" json:"companyId"`
	CustomerID  string             `gorm:"primaryKey;not null;column:customerId" json:"customerId"`
	LocationID  string             `gorm:"primaryKey;not null;column:locationId" json:"locationId"`
	Meta        *datatypes.JSON    `gorm:"column:meta;not null;column:meta" json:"meta,omitempty"`    // the metadata about the most recent change to the row
	Metadata    *datatypes.JSON    `gorm:"column:metadata;column:metadata" json:"metadata,omitempty"` // metadata reserved for customers to control
	CreatedDate datatypes.DateTime `gorm:"column:createdDate;not null;column:createdDate" json:"createdDate"`
}

var _ Model = (*CustomerLocationConnection)(nil)

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *CustomerLocationConnection) TableName() string {
	return "customer_location_connection"
}

// PrimaryKeys returns an array of the primary keys for this model
func (m *CustomerLocationConnection) PrimaryKeys() []string {
	return []string{"customerId", "locationId"}
}

// String returns a string representation as JSON for this model
func (m *CustomerLocationConnection) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewCustomerLocationConnection returns a new model instance from an encoded buffer
func NewCustomerLocationConnection(buf []byte) (*CustomerLocationConnection, error) {
	var result CustomerLocationConnection
	err := json.Unmarshal(buf, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// NewCustomerLocationConnectionFromChangeEvent returns a new model instance from an encoded buffer as change event
func NewCustomerLocationConnectionFromChangeEvent(buf []byte, gzip bool) (*datatypes.ChangeEvent[CustomerLocationConnection], error) {
	var result datatypes.ChangeEvent[CustomerLocationConnection]
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
