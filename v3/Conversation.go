// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	datatypes "github.com/shopmonkeyus/go-datamodel/datatypes"
)

// Conversation schema
type Conversation struct {
	ID          string              `bson:"_id" gorm:"primaryKey;not null;column:id" json:"id"`
	CreatedDate datatypes.DateTime  `gorm:"column:createdDate;not null;column:createdDate" json:"createdDate"`
	UpdatedDate *datatypes.DateTime `gorm:"column:updatedDate;column:updatedDate" json:"updatedDate"`
	Meta        *datatypes.JSON     `gorm:"column:meta;not null;column:meta" json:"meta,omitempty"`    // the metadata about the most recent change to the row
	Metadata    *datatypes.JSON     `gorm:"column:metadata;column:metadata" json:"metadata,omitempty"` // metadata reserved for customers to control
	CompanyID   string              `gorm:"not null;column:companyId" json:"companyId"`
	LocationID  string              `gorm:"not null;column:locationId" json:"locationId"`

	UPGRADEOrderID *string             `gorm:"column:UPGRADE_orderId" json:"UPGRADE_orderId"` // Boolean /// if the record has been deleted
	Archived       bool                `gorm:"not null;column:archived" json:"archived"`
	CustomerID     *string             `gorm:"column:customerId" json:"customerId"`
	Deleted        bool                `gorm:"not null;column:deleted" json:"deleted"`    // if the record has been deleted
	DeletedDate    *datatypes.DateTime `gorm:"column:deletedDate" json:"deletedDate"`     // the date that the record was deleted or null if not deleted
	DeletedReason  *string             `gorm:"column:deletedReason" json:"deletedReason"` // the reason that the record was deleted
	DeletedUserID  *string             `gorm:"column:deletedUserId" json:"deletedUserId"` // the user that deleted the record or null if not deleted
}

var _ Model = (*Conversation)(nil)

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *Conversation) TableName() string {
	return "conversation"
}

// PrimaryKeys returns an array of the primary keys for this model
func (m *Conversation) PrimaryKeys() []string {
	return []string{"id"}
}

// String returns a string representation as JSON for this model
func (m *Conversation) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewConversation returns a new model instance from an encoded buffer
func NewConversation(buf []byte) (*Conversation, error) {
	var result Conversation
	err := json.Unmarshal(buf, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// NewConversationFromChangeEvent returns a new model instance from an encoded buffer as change event
func NewConversationFromChangeEvent(buf []byte, gzip bool) (*datatypes.ChangeEvent[Conversation], error) {
	var result datatypes.ChangeEvent[Conversation]
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
