// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	mapstructure "github.com/mitchellh/mapstructure"
	"time"
)

type ReminderTypeEnum string

const (
	ReminderTypeDays   ReminderTypeEnum = "Days"
	ReminderTypeWeeks                   = "Weeks"
	ReminderTypeMonths                  = "Months"
	ReminderTypeYears                   = "Years"
)

type Reminder struct {
	ID            string           `gorm:"primaryKey;not null" json:"id"`
	AppointmentId string           `gorm:"not null" json:"appointmentId"`
	CompanyId     string           `gorm:"not null" json:"companyId"`
	CreatedDate   time.Time        `gorm:"column:createdDate;not null" json:"createdDate"`
	Meta          *Meta            `gorm:"type:json;embedded;column:meta;not null" json:"meta,omitempty"` // the metadata about the most recent change to the row
	Metadata      any              `gorm:"type:json" json:"metadata,omitempty"`                           // metadata reserved for customers to control
	Type          ReminderTypeEnum `gorm:"not null" json:"type"`
	UpdatedDate   *time.Time       `gorm:"column:updatedDate" json:"updatedDate"`
	Value         int64            `gorm:"not null" json:"value"`
}

var _ Model = (*Reminder)(nil)

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *Reminder) TableName() string {
	return "reminder"
}

func (m *Reminder) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewReminder returns a new model instance from a json key/value map
func NewReminder(kv map[string]any) (*Reminder, error) {
	var result Reminder
	err := mapstructure.Decode(kv, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
