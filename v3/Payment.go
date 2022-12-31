// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	codec "github.com/hashicorp/go-msgpack/v2/codec"
	datatypes "github.com/shopmonkeyus/go-datamodel/datatypes"
	"time"
)

type PaymentCollectionTypeEnum string

const (
	PaymentCollectionTypeManual   PaymentCollectionTypeEnum = "Manual"
	PaymentCollectionTypeOnline                             = "Online"
	PaymentCollectionTypeInPerson                           = "InPerson"
)

type PaymentPaymentTypeEnum string

const (
	PaymentPaymentTypeCheck PaymentPaymentTypeEnum = "Check"
	PaymentPaymentTypeCash                         = "Cash"
	PaymentPaymentTypeCard                         = "Card"
	PaymentPaymentTypeOther                        = "Other"
)

type PaymentProviderEnum string

const (
	PaymentProviderManual    PaymentProviderEnum = "Manual"
	PaymentProviderFirstMile                     = "FirstMile"
	PaymentProviderStripe                        = "Stripe"
)

type PaymentTransactionTypeEnum string

const (
	PaymentTransactionTypeCharge   PaymentTransactionTypeEnum = "Charge"
	PaymentTransactionTypeRefund                              = "Refund"
	PaymentTransactionTypeTransfer                            = "Transfer"
)

type Payment struct {
	ID          string          `gorm:"primaryKey;not null;column:id" json:"id"`
	CreatedDate time.Time       `gorm:"column:createdDate;not null;column:createdDate" json:"createdDate"`
	UpdatedDate *time.Time      `gorm:"column:updatedDate;column:updatedDate" json:"updatedDate"`
	Meta        datatypes.Meta  `gorm:"column:meta;not null;column:meta" json:"meta,omitempty"`    // the metadata about the most recent change to the row
	Metadata    *datatypes.JSON `gorm:"column:metadata;column:metadata" json:"metadata,omitempty"` // metadata reserved for customers to control
	CompanyID   string          `gorm:"not null;column:companyId" json:"companyId"`
	LocationID  string          `gorm:"not null;column:locationId" json:"locationId"`

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

func (m *Payment) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewPayment returns a new model instance from an encoded buffer
func NewPayment(buf []byte, enctype EncodingType) (*Payment, error) {
	var result Payment
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
