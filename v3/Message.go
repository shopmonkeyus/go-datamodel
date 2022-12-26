// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
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
	CompanyId      string                  `gorm:"not null" json:"companyId"`
	ID             string                  `gorm:"primaryKey;not null" json:"id"`
	ConversationId string                  `gorm:"not null" json:"conversationId"`
	CreatedDate    time.Time               `gorm:"column:createdDate;not null" json:"createdDate"`
	CustomerId     *string                 `json:"customerId"` // the id of the customer that authored the message
	EmailError     *string                 `json:"emailError"` // error sending email, if any
	EmailStatus    *MessageEmailStatusEnum `json:"emailStatus"`
	Internal       *bool                   `json:"internal"` // true if internal note, is null if author is a Customer
	LocationId     string                  `gorm:"not null" json:"locationId"`
	Meta           *Meta                   `gorm:"embedded;column:meta;not null" json:"meta,omitempty"`         // the metadata about the most recent change to the row
	Metadata       any                     `gorm:"embedded;column:metadata;not null" json:"metadata,omitempty"` // metadata reserved for customers to control
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

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *Message) TableName() string {
	return "message"
}

func (m *Message) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}
