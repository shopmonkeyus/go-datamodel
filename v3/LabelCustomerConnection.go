// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	datatypes "github.com/shopmonkeyus/go-datamodel/datatypes"
)

// LabelCustomerConnection schema
type LabelCustomerConnection struct {
	LabelID    string `gorm:"primaryKey;not null;column:labelId" json:"labelId"`
	CustomerID string `gorm:"primaryKey;not null;column:customerId" json:"customerId"`

	CompanyID   string              `gorm:"not null;column:companyId" json:"companyId"`
	CreatedDate *datatypes.DateTime `gorm:"column:createdDate;column:createdDate" json:"createdDate"`
	Meta        *datatypes.JSON     `gorm:"column:meta;not null;column:meta" json:"meta,omitempty"` // the metadata about the most recent change to the row
	Metadata    *datatypes.JSON     `gorm:"column:metadata;column:metadata" json:"metadata,omitempty"`
	UpdatedDate *datatypes.DateTime `gorm:"column:updatedDate;column:updatedDate" json:"updatedDate"`
}

var _ Model = (*LabelCustomerConnection)(nil)

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *LabelCustomerConnection) TableName() string {
	return "label_customer_connection"
}

// PrimaryKeys returns an array of the primary keys for this model
func (m *LabelCustomerConnection) PrimaryKeys() []string {
	return []string{"labelId", "customerId"}
}

// String returns a string representation as JSON for this model
func (m *LabelCustomerConnection) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewLabelCustomerConnection returns a new model instance from an encoded buffer
func NewLabelCustomerConnection(buf []byte) (*LabelCustomerConnection, error) {
	var result LabelCustomerConnection
	err := json.Unmarshal(buf, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// NewLabelCustomerConnectionFromChangeEvent returns a new model instance from an encoded buffer as change event
func NewLabelCustomerConnectionFromChangeEvent(buf []byte, gzip bool) (*datatypes.ChangeEvent[LabelCustomerConnection], error) {
	var result datatypes.ChangeEvent[LabelCustomerConnection]
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
