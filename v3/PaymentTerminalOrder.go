// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	datatypes "github.com/shopmonkeyus/go-datamodel/datatypes"
)

// PaymentTerminalOrder schema
type PaymentTerminalOrderProviderEnum string

const (
	PaymentTerminalOrderProviderManual    PaymentTerminalOrderProviderEnum = "Manual"
	PaymentTerminalOrderProviderFirstMile PaymentTerminalOrderProviderEnum = "FirstMile"
	PaymentTerminalOrderProviderStripe    PaymentTerminalOrderProviderEnum = "Stripe"
)

type PaymentTerminalOrder struct {
	ID          string              `bson:"_id" gorm:"primaryKey;not null;column:id" json:"id"`
	CreatedDate datatypes.DateTime  `gorm:"column:createdDate;not null;column:createdDate" json:"createdDate"`
	UpdatedDate *datatypes.DateTime `gorm:"column:updatedDate;column:updatedDate" json:"updatedDate"`
	Meta        *datatypes.JSON     `gorm:"column:meta;not null;column:meta" json:"meta,omitempty"` // the metadata about the most recent change to the row
	Metadata    *datatypes.JSON     `gorm:"column:metadata;column:metadata" json:"metadata,omitempty"`
	CompanyID   string              `gorm:"not null;column:companyId" json:"companyId"`
	LocationID  string              `gorm:"not null;column:locationId" json:"locationId"`

	Address1     string                            `gorm:"not null;column:address1" json:"address1"`
	Address2     *string                           `gorm:"column:address2" json:"address2"`
	City         string                            `gorm:"not null;column:city" json:"city"`
	Provider     *PaymentTerminalOrderProviderEnum `gorm:"column:provider" json:"provider"`
	ProviderData datatypes.JSON                    `gorm:"column:providerData" json:"providerData"`
	Quantity     int64                             `gorm:"not null;column:quantity" json:"quantity"`
	State        string                            `gorm:"not null;column:state" json:"state"`
	Zip          string                            `gorm:"not null;column:zip" json:"zip"`
}

var _ Model = (*PaymentTerminalOrder)(nil)

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *PaymentTerminalOrder) TableName() string {
	return "payment_terminal_order"
}

// PrimaryKeys returns an array of the primary keys for this model
func (m *PaymentTerminalOrder) PrimaryKeys() []string {
	return []string{"id"}
}

// String returns a string representation as JSON for this model
func (m *PaymentTerminalOrder) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewPaymentTerminalOrder returns a new model instance from an encoded buffer
func NewPaymentTerminalOrder(buf []byte) (*PaymentTerminalOrder, error) {
	var result PaymentTerminalOrder
	err := json.Unmarshal(buf, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// NewPaymentTerminalOrderFromChangeEvent returns a new model instance from an encoded buffer as change event
func NewPaymentTerminalOrderFromChangeEvent(buf []byte, gzip bool) (*datatypes.ChangeEvent[PaymentTerminalOrder], error) {
	var result datatypes.ChangeEvent[PaymentTerminalOrder]
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
