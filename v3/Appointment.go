// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	datatypes "github.com/shopmonkeyus/go-datamodel/datatypes"
)

// Appointment schema
type AppointmentColorEnum string

const (
	AppointmentColoraqua   AppointmentColorEnum = "aqua"
	AppointmentColorblack  AppointmentColorEnum = "black"
	AppointmentColorblue   AppointmentColorEnum = "blue"
	AppointmentColorbrown  AppointmentColorEnum = "brown"
	AppointmentColorgray   AppointmentColorEnum = "gray"
	AppointmentColorgreen  AppointmentColorEnum = "green"
	AppointmentColororange AppointmentColorEnum = "orange"
	AppointmentColorpurple AppointmentColorEnum = "purple"
	AppointmentColorred    AppointmentColorEnum = "red"
	AppointmentColoryellow AppointmentColorEnum = "yellow"
)

type Appointment struct {
	ID          string              `bson:"_id" gorm:"primaryKey;not null;column:id" json:"id"`
	CreatedDate datatypes.DateTime  `gorm:"column:createdDate;not null;column:createdDate" json:"createdDate"`
	UpdatedDate *datatypes.DateTime `gorm:"column:updatedDate;column:updatedDate" json:"updatedDate"`
	Meta        *datatypes.JSON     `gorm:"column:meta;not null;column:meta" json:"meta,omitempty"` // the metadata about the most recent change to the row
	Metadata    *datatypes.JSON     `gorm:"column:metadata;column:metadata" json:"metadata,omitempty"`
	CompanyID   string              `gorm:"not null;column:companyId" json:"companyId"`
	LocationID  string              `gorm:"not null;column:locationId" json:"locationId"`

	AllDay                bool                 `gorm:"not null;column:allDay" json:"allDay"`
	CancelationNote       *string              `gorm:"column:cancelationNote" json:"cancelationNote"` // if the appointment was canceled
	Color                 AppointmentColorEnum `gorm:"not null;column:color" json:"color"`
	Confirmed             bool                 `gorm:"not null;column:confirmed" json:"confirmed"`
	CustomerEmailID       *string              `gorm:"column:customerEmailId" json:"customerEmailId"`
	CustomerID            *string              `gorm:"column:customerId" json:"customerId"`
	CustomerPhoneNumberID *string              `gorm:"column:customerPhoneNumberId" json:"customerPhoneNumberId"`
	EndDate               datatypes.DateTime   `gorm:"not null;column:endDate" json:"endDate"` // end date and time of the appointment
	Name                  string               `gorm:"not null;column:name" json:"name"`       // name of the appointment like 'Oil change'
	Note                  string               `gorm:"not null;column:note" json:"note"`       // notes for the appointment
	OrderID               *string              `gorm:"column:orderId" json:"orderId"`
	PendingConfirmation   bool                 `gorm:"not null;column:pendingConfirmation" json:"pendingConfirmation"`
	RemovedFromRecurrency bool                 `gorm:"not null;column:removedFromRecurrency" json:"removedFromRecurrency"`
	Rruleset              *string              `gorm:"column:rruleset" json:"rruleset"`
	SendConfirmation      bool                 `gorm:"not null;column:sendConfirmation" json:"sendConfirmation"` // Send confirmation notification at the moment of saving the appointment
	SendReminder          bool                 `gorm:"not null;column:sendReminder" json:"sendReminder"`         // Send reminder will send a notification one day before the apponintment. This would need some sort of scheduling mecanism to fire at the right time.
	StartDate             datatypes.DateTime   `gorm:"not null;column:startDate" json:"startDate"`               // start date and time of the appointment
	UseEmail              bool                 `gorm:"not null;column:useEmail" json:"useEmail"`                 // In case we want to use email to send confirmation and/or reminder
	UseSMS                bool                 `gorm:"not null;column:useSMS" json:"useSMS"`                     // In case we want to use email to send confirmation and/or reminder
	VehicleID             *string              `gorm:"column:vehicleId" json:"vehicleId"`
}

var _ Model = (*Appointment)(nil)

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *Appointment) TableName() string {
	return "appointment"
}

// PrimaryKeys returns an array of the primary keys for this model
func (m *Appointment) PrimaryKeys() []string {
	return []string{"id"}
}

// String returns a string representation as JSON for this model
func (m *Appointment) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewAppointment returns a new model instance from an encoded buffer
func NewAppointment(buf []byte) (*Appointment, error) {
	var result Appointment
	err := json.Unmarshal(buf, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// NewAppointmentFromChangeEvent returns a new model instance from an encoded buffer as change event
func NewAppointmentFromChangeEvent(buf []byte, gzip bool) (*datatypes.ChangeEvent[Appointment], error) {
	var result datatypes.ChangeEvent[Appointment]
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
