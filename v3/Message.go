// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	mapstructure "github.com/mitchellh/mapstructure"
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
	AuthorType     MessageAuthorTypeEnum   `gorm:"not null" json:"authorType"`
	ID             string                  `gorm:"primaryKey;not null" json:"id"`
	CompanyId      string                  `gorm:"not null" json:"companyId"`
	ConversationId string                  `gorm:"not null" json:"conversationId"`
	CreatedDate    time.Time               `gorm:"column:createdDate;not null" json:"createdDate"`
	CustomerId     *string                 `json:"customerId"` // the id of the customer that authored the message
	EmailError     *string                 `json:"emailError"` // error sending email, if any
	EmailStatus    *MessageEmailStatusEnum `json:"emailStatus"`
	Internal       *bool                   `json:"internal"` // true if internal note, is null if author is a Customer
	LocationId     string                  `gorm:"not null" json:"locationId"`
	Meta           *Meta                   `gorm:"type:json;embedded;column:meta;not null" json:"meta,omitempty"` // the metadata about the most recent change to the row
	Metadata       any                     `gorm:"type:json" json:"metadata,omitempty"`                           // metadata reserved for customers to control
	Origin         MessageOriginEnum       `gorm:"not null" json:"origin"`
	SendEmail      *bool                   `json:"sendEmail"` // if an email should be sent, is null if author is a Customer
	SendSms        *bool                   `json:"sendSms"`   // if an sms should be sent, is null if author is a Customer
	ShopRead       *bool                   `json:"shopRead"`  // true if someone at the shop has read a Customer message, is null if author is a User
	SmsError       *string                 `json:"smsError"`  // error sending sms, if any
	SmsStatus      *MessageSmsStatusEnum   `json:"smsStatus"`
	Text           string                  `gorm:"not null" json:"text"`
	UpdatedDate    *time.Time              `gorm:"column:updatedDate" json:"updatedDate"`
	UserId         *string                 `json:"userId"` // the user who authored the message
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

// NewMessage returns a new model instance from a json key/value map
func NewMessage(kv map[string]any) (*Message, error) {
	var result Message
	err := mapstructure.Decode(kv, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
