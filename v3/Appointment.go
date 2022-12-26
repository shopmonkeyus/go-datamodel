// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	mapstructure "github.com/mitchellh/mapstructure"
	"time"
)

type AppointmentRecurringParamsFrequencyEnum string

const (
	AppointmentRecurringParamsFrequencyDay   AppointmentRecurringParamsFrequencyEnum = "Day"
	AppointmentRecurringParamsFrequencyWeek                                          = "Week"
	AppointmentRecurringParamsFrequencyMonth                                         = "Month"
	AppointmentRecurringParamsFrequencyYear                                          = "Year"
)

type Appointment struct {
	ID                       string                                   `gorm:"primaryKey;not null" json:"id"`
	AllDay                   bool                                     `gorm:"not null" json:"allDay"`
	CancelationNote          *string                                  `json:"cancelationNote"` // if the appointment was canceled
	Color                    *string                                  `json:"color"`           // Color for the appointment
	CompanyId                string                                   `gorm:"not null" json:"companyId"`
	Confirmed                bool                                     `gorm:"not null" json:"confirmed"`
	CreatedDate              time.Time                                `gorm:"column:createdDate;not null" json:"createdDate"`
	CustomerId               *string                                  `json:"customerId"`
	EndDate                  *time.Time                               `json:"endDate"`                 // end date and time of the appointment
	EndRecurringParamsCount  *int64                                   `json:"endRecurringParamsCount"` // in case selected "After" in the drop down, in End Repeat section.
	EndRecurringParamsUntil  *time.Time                               `json:"endRecurringParamsUntil"` // in case user selected "On" in the drop down, in End Repeat section.
	LocationId               string                                   `gorm:"not null" json:"locationId"`
	Meta                     *Meta                                    `gorm:"type:json;embedded;column:meta;not null" json:"meta,omitempty"` // the metadata about the most recent change to the row
	Metadata                 any                                      `gorm:"type:json" json:"metadata,omitempty"`                           // metadata reserved for customers to control
	Name                     string                                   `gorm:"not null" json:"name"`                                          // name of the appointment like 'Oil change'
	Note                     string                                   `gorm:"not null" json:"note"`                                          // notes for the appointment
	OrderId                  *string                                  `json:"orderId"`
	PendingConfirmation      bool                                     `gorm:"not null" json:"pendingConfirmation"`
	Recurring                bool                                     `gorm:"not null" json:"recurring"` // In case this appointment need a repeat pattern. We probably want to create new appointment records based on the repeat pattern
	RecurringParamsFrequency *AppointmentRecurringParamsFrequencyEnum `json:"recurringParamsFrequency"`
	RecurringParamsInterval  *int64                                   `json:"recurringParamsInterval"` // the frequency of the appointment
	RemovedFromRecurrency    bool                                     `gorm:"not null" json:"removedFromRecurrency"`
	SendConfirmation         bool                                     `gorm:"not null" json:"sendConfirmation"` // Send confirmation notification at the moment of saving the appointment
	SendReminder             bool                                     `gorm:"not null" json:"sendReminder"`     // Send reminder will send a notification one day before the apponintment. This would need some sort of scheduling mecanism to fire at the right time.
	StartDate                *time.Time                               `json:"startDate"`                        // start date and time of the appointment
	TechnicianIds            []string                                 `gorm:"not null" json:"technicianIds"`
	UpdatedDate              *time.Time                               `gorm:"column:updatedDate" json:"updatedDate"`
	UseEmail                 bool                                     `gorm:"not null" json:"useEmail"` // In case we want to use email to send confirmation and/or reminder
	UseSMS                   bool                                     `gorm:"not null" json:"useSMS"`   // In case we want to use email to send confirmation and/or reminder
	VehicleId                *string                                  `json:"vehicleId"`
}

var _ Model = (*Appointment)(nil)

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *Appointment) TableName() string {
	return "appointment"
}

func (m *Appointment) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewAppointment returns a new model instance from a json key/value map
func NewAppointment(kv map[string]any) (*Appointment, error) {
	var result Appointment
	err := mapstructure.Decode(kv, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
