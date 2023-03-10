// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	datatypes "github.com/shopmonkeyus/go-datamodel/datatypes"
)

// Reminder schema
type ReminderTypeEnum string

const (
	ReminderTypeHours ReminderTypeEnum = "Hours"
	ReminderTypeDays  ReminderTypeEnum = "Days"
	ReminderTypeWeeks ReminderTypeEnum = "Weeks"
)

type Reminder struct {
	ID          string              `bson:"_id" gorm:"primaryKey;not null;column:id" json:"id"`
	CreatedDate datatypes.DateTime  `gorm:"column:createdDate;not null;column:createdDate" json:"createdDate"`
	UpdatedDate *datatypes.DateTime `gorm:"column:updatedDate;column:updatedDate" json:"updatedDate"`
	Meta        *datatypes.JSON     `gorm:"column:meta;not null;column:meta" json:"meta,omitempty"` // the metadata about the most recent change to the row
	Metadata    *datatypes.JSON     `gorm:"column:metadata;column:metadata" json:"metadata,omitempty"`
	CompanyID   string              `gorm:"not null;column:companyId" json:"companyId"`

	AppointmentID string           `gorm:"not null;column:appointmentId" json:"appointmentId"`
	Type          ReminderTypeEnum `gorm:"not null;column:type" json:"type"`
	Value         int64            `gorm:"not null;column:value" json:"value"`
}

var _ Model = (*Reminder)(nil)

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *Reminder) TableName() string {
	return "reminder"
}

// PrimaryKeys returns an array of the primary keys for this model
func (m *Reminder) PrimaryKeys() []string {
	return []string{"id"}
}

// String returns a string representation as JSON for this model
func (m *Reminder) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewReminder returns a new model instance from an encoded buffer
func NewReminder(buf []byte) (*Reminder, error) {
	var result Reminder
	err := json.Unmarshal(buf, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// NewReminderFromChangeEvent returns a new model instance from an encoded buffer as change event
func NewReminderFromChangeEvent(buf []byte, gzip bool) (*datatypes.ChangeEvent[Reminder], error) {
	var result datatypes.ChangeEvent[Reminder]
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
