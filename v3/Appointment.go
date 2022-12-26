// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
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
	ID          string     `gorm:"primaryKey;not null;column:id" json:"id"`
	CreatedDate time.Time  `gorm:"column:createdDate;not null;column:createdDate" json:"createdDate"`
	UpdatedDate *time.Time `gorm:"column:updatedDate;column:updatedDate" json:"updatedDate"`
	Meta        *Meta      `gorm:"type:json;embedded;serializer:json;column:meta;not null;column:meta" json:"meta,omitempty"` // the metadata about the most recent change to the row
	Metadata    any        `gorm:"type:json;serializer:json;column:metadata" json:"metadata,omitempty"`                       // metadata reserved for customers to control
	CompanyID   string     `gorm:"not null;column:companyId" json:"companyId"`
	LocationID  string     `gorm:"not null;column:locationId" json:"locationId"`

	AllDay                   bool                                     `gorm:"not null;column:allDay" json:"allDay"`
	CancelationNote          *string                                  `gorm:"column:cancelationNote" json:"cancelationNote"` // if the appointment was canceled
	Color                    *string                                  `gorm:"column:color" json:"color"`                     // Color for the appointment
	Confirmed                bool                                     `gorm:"not null;column:confirmed" json:"confirmed"`
	CustomerID               *string                                  `gorm:"column:customerId" json:"customerId"`
	EndDate                  *time.Time                               `gorm:"column:endDate" json:"endDate"`                                 // end date and time of the appointment
	EndRecurringParamsCount  *int64                                   `gorm:"column:endRecurringParamsCount" json:"endRecurringParamsCount"` // in case selected "After" in the drop down, in End Repeat section.
	EndRecurringParamsUntil  *time.Time                               `gorm:"column:endRecurringParamsUntil" json:"endRecurringParamsUntil"` // in case user selected "On" in the drop down, in End Repeat section.
	Name                     string                                   `gorm:"not null;column:name" json:"name"`                              // name of the appointment like 'Oil change'
	Note                     string                                   `gorm:"not null;column:note" json:"note"`                              // notes for the appointment
	OrderID                  *string                                  `gorm:"column:orderId" json:"orderId"`
	PendingConfirmation      bool                                     `gorm:"not null;column:pendingConfirmation" json:"pendingConfirmation"`
	Recurring                bool                                     `gorm:"not null;column:recurring" json:"recurring"` // In case this appointment need a repeat pattern. We probably want to create new appointment records based on the repeat pattern
	RecurringParamsFrequency *AppointmentRecurringParamsFrequencyEnum `gorm:"column:recurringParamsFrequency" json:"recurringParamsFrequency"`
	RecurringParamsInterval  *int64                                   `gorm:"column:recurringParamsInterval" json:"recurringParamsInterval"` // the frequency of the appointment
	RemovedFromRecurrency    bool                                     `gorm:"not null;column:removedFromRecurrency" json:"removedFromRecurrency"`
	SendConfirmation         bool                                     `gorm:"not null;column:sendConfirmation" json:"sendConfirmation"` // Send confirmation notification at the moment of saving the appointment
	SendReminder             bool                                     `gorm:"not null;column:sendReminder" json:"sendReminder"`         // Send reminder will send a notification one day before the apponintment. This would need some sort of scheduling mecanism to fire at the right time.
	StartDate                *time.Time                               `gorm:"column:startDate" json:"startDate"`                        // start date and time of the appointment
	TechnicianIds            []string                                 `gorm:"not null;column:technicianIds" json:"technicianIds"`
	UseEmail                 bool                                     `gorm:"not null;column:useEmail" json:"useEmail"` // In case we want to use email to send confirmation and/or reminder
	UseSMS                   bool                                     `gorm:"not null;column:useSMS" json:"useSMS"`     // In case we want to use email to send confirmation and/or reminder
	VehicleID                *string                                  `gorm:"column:vehicleId" json:"vehicleId"`
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
func NewAppointment(buf []byte) (*Appointment, error) {
	var result Appointment
	err := json.Unmarshal(buf, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
