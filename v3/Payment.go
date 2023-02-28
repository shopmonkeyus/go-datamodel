// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	datatypes "github.com/shopmonkeyus/go-datamodel/datatypes"
)

// Payment schema
type PaymentPaymentModeEnum string

const (
	PaymentPaymentModeManual   PaymentPaymentModeEnum = "Manual"
	PaymentPaymentModeOnline   PaymentPaymentModeEnum = "Online"
	PaymentPaymentModeInPerson PaymentPaymentModeEnum = "InPerson"
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
)

type PaymentRefundReasonEnum string

const (
	PaymentRefundReasonDuplicate           PaymentRefundReasonEnum = "Duplicate"
	PaymentRefundReasonFraudulent          PaymentRefundReasonEnum = "Fraudulent"
	PaymentRefundReasonRequestedByCustomer PaymentRefundReasonEnum = "RequestedByCustomer"
)

type PaymentTransactionTypeEnum string

const (
	PaymentTransactionTypeCharge   PaymentTransactionTypeEnum = "Charge"
	PaymentTransactionTypeRefund   PaymentTransactionTypeEnum = "Refund"
	PaymentTransactionTypeTransfer PaymentTransactionTypeEnum = "Transfer"
)

type Payment struct {
	ID          string              `bson:"_id" gorm:"primaryKey;not null;column:id" json:"id"`
	CreatedDate datatypes.DateTime  `gorm:"column:createdDate;not null;column:createdDate" json:"createdDate"`
	UpdatedDate *datatypes.DateTime `gorm:"column:updatedDate;column:updatedDate" json:"updatedDate"`
	Meta        *datatypes.JSON     `gorm:"column:meta;not null;column:meta" json:"meta,omitempty"` // the metadata about the most recent change to the row
	Metadata    *datatypes.JSON     `gorm:"column:metadata;column:metadata" json:"metadata,omitempty"`
	CompanyID   string              `gorm:"not null;column:companyId" json:"companyId"`
	LocationID  string              `gorm:"not null;column:locationId" json:"locationId"`

	AmountCents          int64                      `gorm:"not null;column:amountCents" json:"amountCents"` // amount charged or refuned
	Bulk                 bool                       `gorm:"not null;column:bulk" json:"bulk"`
	CardConfirmation     *string                    `gorm:"column:cardConfirmation" json:"cardConfirmation"`
	CardDigits           *string                    `gorm:"column:cardDigits" json:"cardDigits"`
	CardName             *string                    `gorm:"column:cardName" json:"cardName"`
	CardType             *string                    `gorm:"column:cardType" json:"cardType"`
	ChargeFromPublicPage bool                       `gorm:"not null;column:chargeFromPublicPage" json:"chargeFromPublicPage"`
	ChargeID             *string                    `gorm:"column:chargeId" json:"chargeId"` // reference for original payment if transaction type is different from charge
	CheckNumber          *string                    `gorm:"column:checkNumber" json:"checkNumber"`
	DebitCard            bool                       `gorm:"not null;column:debitCard" json:"debitCard"`
	Deposit              bool                       `gorm:"not null;column:deposit" json:"deposit"`
	Note                 string                     `gorm:"not null;column:note" json:"note"`
	OrderID              string                     `gorm:"not null;column:orderId" json:"orderId"`
	PayerID              *string                    `gorm:"column:payerId" json:"payerId"`
	PaymentMode          PaymentPaymentModeEnum     `gorm:"not null;column:paymentMode" json:"paymentMode"`
	PaymentType          PaymentPaymentTypeEnum     `gorm:"not null;column:paymentType" json:"paymentType"`
	Provider             *PaymentProviderEnum       `gorm:"column:provider" json:"provider"`
	ProviderData         datatypes.JSON             `gorm:"column:providerData" json:"providerData"`
	ReceiptNumber        int64                      `gorm:"not null;column:receiptNumber" json:"receiptNumber"`
	RecordedDate         *datatypes.DateTime        `gorm:"column:recordedDate" json:"recordedDate"` // the date that the payment was recorded
	RefundReason         *PaymentRefundReasonEnum   `gorm:"column:refundReason" json:"refundReason"`
	Refunded             *bool                      `gorm:"column:refunded" json:"refunded"`
	RefundedAmountCents  *int64                     `gorm:"column:refundedAmountCents" json:"refundedAmountCents"` // amount refunded for charge transactions
	StatementID          *string                    `gorm:"column:statementId" json:"statementId"`
	TransactionType      PaymentTransactionTypeEnum `gorm:"not null;column:transactionType" json:"transactionType"`
	UserData             datatypes.JSON             `gorm:"column:userData" json:"userData"`
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
