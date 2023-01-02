// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	datatypes "github.com/shopmonkeyus/go-datamodel/datatypes"
)

type PaymentCollectionTypeEnum string

const (
	PaymentCollectionTypeManual   PaymentCollectionTypeEnum = "Manual"
	PaymentCollectionTypeOnline   PaymentCollectionTypeEnum = "Online"
	PaymentCollectionTypeInPerson PaymentCollectionTypeEnum = "InPerson"
)

type PaymentPaymentTypeEnum string

const (
	PaymentPaymentTypeCheck PaymentPaymentTypeEnum = "Check"
	PaymentPaymentTypeCash  PaymentPaymentTypeEnum = "Cash"
	PaymentPaymentTypeCard  PaymentPaymentTypeEnum = "Card"
	PaymentPaymentTypeOther PaymentPaymentTypeEnum = "Other"
)

type PaymentProviderEnum string

const (
	PaymentProviderManual    PaymentProviderEnum = "Manual"
	PaymentProviderFirstMile PaymentProviderEnum = "FirstMile"
	PaymentProviderStripe    PaymentProviderEnum = "Stripe"
	PaymentProvider          PaymentProviderEnum = ""
)

type PaymentTransactionTypeEnum string

const (
	PaymentTransactionTypeCharge   PaymentTransactionTypeEnum = "Charge"
	PaymentTransactionTypeRefund   PaymentTransactionTypeEnum = "Refund"
	PaymentTransactionTypeTransfer PaymentTransactionTypeEnum = "Transfer"
)

type Payment struct {
	ID          string              `gorm:"primaryKey;not null;column:id" json:"id"`
	CreatedDate datatypes.DateTime  `gorm:"column:createdDate;not null;column:createdDate" json:"createdDate"`
	UpdatedDate *datatypes.DateTime `gorm:"column:updatedDate;column:updatedDate" json:"updatedDate"`
	Meta        *datatypes.JSON     `gorm:"column:meta;not null;column:meta" json:"meta,omitempty"`    // the metadata about the most recent change to the row
	Metadata    *datatypes.JSON     `gorm:"column:metadata;column:metadata" json:"metadata,omitempty"` // metadata reserved for customers to control
	CompanyID   string              `gorm:"not null;column:companyId" json:"companyId"`
	LocationID  string              `gorm:"not null;column:locationId" json:"locationId"`

	Amount               int64                      `gorm:"not null;column:amount" json:"amount"`
	Bulk                 bool                       `gorm:"not null;column:bulk" json:"bulk"`
	CardConfirmation     *string                    `gorm:"column:cardConfirmation" json:"cardConfirmation"`
	CardDigits           *int64                     `gorm:"column:cardDigits" json:"cardDigits"`
	CardName             *string                    `gorm:"column:cardName" json:"cardName"`
	CardType             *string                    `gorm:"column:cardType" json:"cardType"`
	ChargeFromPublicPage bool                       `gorm:"not null;column:chargeFromPublicPage" json:"chargeFromPublicPage"`
	CollectionType       PaymentCollectionTypeEnum  `gorm:"not null;column:collectionType" json:"collectionType"`
	DebitCard            bool                       `gorm:"not null;column:debitCard" json:"debitCard"`
	Deposit              bool                       `gorm:"not null;column:deposit" json:"deposit"`
	Note                 string                     `gorm:"not null;column:note" json:"note"`
	OrderID              string                     `gorm:"not null;column:orderId" json:"orderId"`
	PayerID              *string                    `gorm:"column:payerId" json:"payerId"`
	PaymentType          PaymentPaymentTypeEnum     `gorm:"not null;column:paymentType" json:"paymentType"`
	Provider             *PaymentProviderEnum       `gorm:"column:provider" json:"provider"`
	ProviderData         datatypes.JSON             `gorm:"column:providerData" json:"providerData"`
	ReceiptNumber        int64                      `gorm:"not null;column:receiptNumber" json:"receiptNumber"`
	StatementID          string                     `gorm:"not null;column:statementId" json:"statementId"`
	TransactionType      PaymentTransactionTypeEnum `gorm:"not null;column:transactionType" json:"transactionType"`
	Userdata             datatypes.JSON             `gorm:"column:userdata" json:"userdata"`
}

var _ Model = (*Payment)(nil)

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *Payment) TableName() string {
	return "payment"
}

// PrimaryKeys returns an array of the primary keys for this model
func (m *Payment) PrimaryKeys() []string {
	return []string{"id"}
}

// String returns a string representation as JSON for this model
func (m *Payment) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewPayment returns a new model instance from an encoded buffer
func NewPayment(buf []byte) (*Payment, error) {
	var result Payment
	err := json.Unmarshal(buf, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// NewPaymentFromChangeEvent returns a new model instance from an encoded buffer as change event
func NewPaymentFromChangeEvent(buf []byte, gzip bool) (*datatypes.ChangeEvent[Payment], error) {
	var result datatypes.ChangeEvent[Payment]
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
