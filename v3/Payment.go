// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	mapstructure "github.com/mitchellh/mapstructure"
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
	Amount               int64                      `gorm:"not null" json:"amount"`
	ID                   string                     `gorm:"primaryKey;not null" json:"id"`
	Bulk                 bool                       `gorm:"not null" json:"bulk"`
	CardConfirmation     *string                    `json:"cardConfirmation"`
	CardDigits           *int64                     `json:"cardDigits"`
	CardName             *string                    `json:"cardName"`
	CardType             *string                    `json:"cardType"`
	ChargeFromPublicPage bool                       `gorm:"not null" json:"chargeFromPublicPage"`
	CollectionType       PaymentCollectionTypeEnum  `gorm:"not null" json:"collectionType"`
	CompanyId            string                     `gorm:"not null" json:"companyId"`
	CreatedDate          time.Time                  `gorm:"column:createdDate;not null" json:"createdDate"`
	DebitCard            bool                       `gorm:"not null" json:"debitCard"`
	Deposit              bool                       `gorm:"not null" json:"deposit"`
	LocationId           string                     `gorm:"not null" json:"locationId"`
	Meta                 *Meta                      `gorm:"type:json;embedded;column:meta;not null" json:"meta,omitempty"` // the metadata about the most recent change to the row
	Metadata             any                        `gorm:"type:json" json:"metadata,omitempty"`                           // metadata reserved for customers to control
	Note                 string                     `gorm:"not null" json:"note"`
	OrderId              string                     `gorm:"not null" json:"orderId"`
	PayerId              *string                    `json:"payerId"`
	PaymentType          PaymentPaymentTypeEnum     `gorm:"not null" json:"paymentType"`
	Provider             *PaymentProviderEnum       `json:"provider"`
	ProviderData         any                        `gorm:"type:json" json:"providerData"`
	ReceiptNumber        int64                      `gorm:"not null" json:"receiptNumber"`
	StatementId          string                     `gorm:"not null" json:"statementId"`
	TransactionType      PaymentTransactionTypeEnum `gorm:"not null" json:"transactionType"`
	UpdatedDate          *time.Time                 `gorm:"column:updatedDate" json:"updatedDate"`
	Userdata             any                        `gorm:"type:json" json:"userdata"`
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

// NewPayment returns a new model instance from a json key/value map
func NewPayment(kv map[string]any) (*Payment, error) {
	var result Payment
	err := mapstructure.Decode(kv, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
