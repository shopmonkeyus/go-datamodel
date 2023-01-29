// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	datatypes "github.com/shopmonkeyus/go-datamodel/datatypes"
)

// Message schema
type MessageAuthorTypeEnum string

const (
	MessageAuthorTypeCustomer MessageAuthorTypeEnum = "Customer"
	MessageAuthorTypeUser     MessageAuthorTypeEnum = "User"
)

type MessageEmailStatusEnum string

const (
	MessageEmailStatusPending MessageEmailStatusEnum = "Pending"
	MessageEmailStatusSent    MessageEmailStatusEnum = "Sent"
	MessageEmailStatusRead    MessageEmailStatusEnum = "Read"
	MessageEmailStatusError   MessageEmailStatusEnum = "Error"
)

type MessageOriginEnum string

const (
	MessageOriginWeb               MessageOriginEnum = "Web"
	MessageOriginMobile            MessageOriginEnum = "Mobile"
	MessageOriginCustomerOrderPage MessageOriginEnum = "CustomerOrderPage"
	MessageOriginSMS               MessageOriginEnum = "SMS"
	MessageOriginEmail             MessageOriginEnum = "Email"
)

type MessageSmsStatusEnum string

const (
	MessageSmsStatusPending MessageSmsStatusEnum = "Pending"
	MessageSmsStatusSent    MessageSmsStatusEnum = "Sent"
	MessageSmsStatusRead    MessageSmsStatusEnum = "Read"
	MessageSmsStatusError   MessageSmsStatusEnum = "Error"
)

type Message struct {
	ID          string              `bson:"_id" gorm:"primaryKey;not null;column:id" json:"id"`
	CreatedDate datatypes.DateTime  `gorm:"column:createdDate;not null;column:createdDate" json:"createdDate"`
	UpdatedDate *datatypes.DateTime `gorm:"column:updatedDate;column:updatedDate" json:"updatedDate"`
	Meta        *datatypes.JSON     `gorm:"column:meta;not null;column:meta" json:"meta,omitempty"`    // the metadata about the most recent change to the row
	Metadata    *datatypes.JSON     `gorm:"column:metadata;column:metadata" json:"metadata,omitempty"` // metadata reserved for customers to control
	CompanyID   string              `gorm:"not null;column:companyId" json:"companyId"`
	LocationID  string              `gorm:"not null;column:locationId" json:"locationId"`

	AuthorType     MessageAuthorTypeEnum   `gorm:"not null;column:authorType" json:"authorType"`
	ConversationID string                  `gorm:"not null;column:conversationId" json:"conversationId"`
	CustomerID     *string                 `gorm:"column:customerId" json:"customerId"`       // the id of the customer that authored the message
	Deleted        bool                    `gorm:"not null;column:deleted" json:"deleted"`    // if the record has been deleted
	DeletedDate    *datatypes.DateTime     `gorm:"column:deletedDate" json:"deletedDate"`     // the date that the record was deleted or null if not deleted
	DeletedReason  *string                 `gorm:"column:deletedReason" json:"deletedReason"` // the reason that the record was deleted
	DeletedUserID  *string                 `gorm:"column:deletedUserId" json:"deletedUserId"` // the user that deleted the record or null if not deleted
	EmailError     *string                 `gorm:"column:emailError" json:"emailError"`       // error sending email, if any
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

// PrimaryKeys returns an array of the primary keys for this model
func (m *Message) PrimaryKeys() []string {
	return []string{"id"}
}

// String returns a string representation as JSON for this model
func (m *Message) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewMessage returns a new model instance from an encoded buffer
func NewMessage(buf []byte) (*Message, error) {
	var result Message
	err := json.Unmarshal(buf, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// NewMessageFromChangeEvent returns a new model instance from an encoded buffer as change event
func NewMessageFromChangeEvent(buf []byte, gzip bool) (*datatypes.ChangeEvent[Message], error) {
	var result datatypes.ChangeEvent[Message]
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
