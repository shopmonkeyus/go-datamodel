// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	codec "github.com/hashicorp/go-msgpack/v2/codec"
	datatypes "github.com/shopmonkeyus/go-datamodel/datatypes"
	"time"
)

type MessageAuthorTypeEnum string

const (
	MessageAuthorTypeCustomer MessageAuthorTypeEnum = "Customer"
	MessageAuthorTypeUser                           = "User"
)

type MessageEmailStatusEnum string

const (
	MessageEmailStatusPending MessageEmailStatusEnum = "Pending"
	MessageEmailStatusSent                           = "Sent"
	MessageEmailStatusRead                           = "Read"
	MessageEmailStatusError                          = "Error"
)

type MessageOriginEnum string

const (
	MessageOriginWeb               MessageOriginEnum = "Web"
	MessageOriginMobile                              = "Mobile"
	MessageOriginCustomerOrderPage                   = "CustomerOrderPage"
	MessageOriginSMS                                 = "SMS"
	MessageOriginEmail                               = "Email"
)

type MessageSmsStatusEnum string

const (
	MessageSmsStatusPending MessageSmsStatusEnum = "Pending"
	MessageSmsStatusSent                         = "Sent"
	MessageSmsStatusRead                         = "Read"
	MessageSmsStatusError                        = "Error"
)

type Message struct {
	ID          string          `gorm:"primaryKey;not null;column:id" json:"id"`
	CreatedDate time.Time       `gorm:"column:createdDate;not null;column:createdDate" json:"createdDate"`
	UpdatedDate *time.Time      `gorm:"column:updatedDate;column:updatedDate" json:"updatedDate"`
	Meta        datatypes.Meta  `gorm:"column:meta;not null;column:meta" json:"meta,omitempty"`    // the metadata about the most recent change to the row
	Metadata    *datatypes.JSON `gorm:"column:metadata;column:metadata" json:"metadata,omitempty"` // metadata reserved for customers to control
	CompanyID   string          `gorm:"not null;column:companyId" json:"companyId"`
	LocationID  string          `gorm:"not null;column:locationId" json:"locationId"`

	AuthorType     MessageAuthorTypeEnum   `gorm:"not null;column:authorType" json:"authorType"`
	ConversationID string                  `gorm:"not null;column:conversationId" json:"conversationId"`
	CustomerID     *string                 `gorm:"column:customerId" json:"customerId"` // the id of the customer that authored the message
	EmailError     *string                 `gorm:"column:emailError" json:"emailError"` // error sending email, if any
	EmailStatus    *MessageEmailStatusEnum `gorm:"column:emailStatus" json:"emailStatus"`
	Internal       *bool                   `gorm:"column:internal" json:"internal"` // true if internal note, is null if author is a Customer
	Origin         MessageOriginEnum       `gorm:"not null;column:origin" json:"origin"`
	SendEmail      *bool                   `gorm:"column:sendEmail" json:"sendEmail"` // if an email should be sent, is null if author is a Customer
	SendSms        *bool                   `gorm:"column:sendSms" json:"sendSms"`     // if an sms should be sent, is null if author is a Customer
	ShopRead       *bool                   `gorm:"column:shopRead" json:"shopRead"`   // true if someone at the shop has read a Customer message, is null if author is a User
	SmsError       *string                 `gorm:"column:smsError" json:"smsError"`   // error sending sms, if any
	SmsStatus      *MessageSmsStatusEnum   `gorm:"column:smsStatus" json:"smsStatus"`
	Text           string                  `gorm:"not null;column:text" json:"text"`
	UserID         *string                 `gorm:"column:userId" json:"userId"` // the user who authored the message
}

var _ Model = (*Message)(nil)

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *Message) TableName() string {
	return "message"
}

func (m *Message) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewMessage returns a new model instance from an encoded buffer
func NewMessage(buf []byte, enctype EncodingType) (*Message, error) {
	var result Message
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
