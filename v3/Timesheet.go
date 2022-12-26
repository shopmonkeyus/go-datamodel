// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	mapstructure "github.com/mitchellh/mapstructure"
	"time"
)

type TimesheetActivityEnum string

const (
	TimesheetActivityGeneral TimesheetActivityEnum = "General"
	TimesheetActivityOrder                         = "Order"
	TimesheetActivityService                       = "Service"
	TimesheetActivityLabor                         = "Labor"
)

type TimesheetClockInPlatformEnum string

const (
	TimesheetClockInPlatformWeb    TimesheetClockInPlatformEnum = "Web"
	TimesheetClockInPlatformMobile                              = "Mobile"
)

type TimesheetClockOutPlatformEnum string

const (
	TimesheetClockOutPlatformWeb    TimesheetClockOutPlatformEnum = "Web"
	TimesheetClockOutPlatformMobile                               = "Mobile"
)

type TimesheetTypeEnum string

const (
	TimesheetTypeTimeclock TimesheetTypeEnum = "Timeclock"
	TimesheetTypeManual                      = "Manual"
)

type Timesheet struct {
	Activity         TimesheetActivityEnum          `gorm:"not null" json:"activity"`
	ClockIn          time.Time                      `gorm:"not null" json:"clockIn"`
	ClockInPlatform  *TimesheetClockInPlatformEnum  `json:"clockInPlatform"`
	ClockOut         *time.Time                     `json:"clockOut"`
	ClockOutPlatform *TimesheetClockOutPlatformEnum `json:"clockOutPlatform"`
	CompanyId        string                         `gorm:"not null" json:"companyId"`
	CreatedDate      time.Time                      `gorm:"column:createdDate;not null" json:"createdDate"`
	Duration         *float64                       `json:"duration"`                 // the amount of time clocked (in milliseconds)
	FlatRate         bool                           `gorm:"not null" json:"flatRate"` // if the technician uses a flat rate
	ID               string                         `gorm:"primaryKey;not null" json:"id"`
	InProgress       bool                           `gorm:"not null" json:"inProgress"`
	LaborId          *string                        `json:"laborId"`
	LocationId       string                         `gorm:"not null" json:"locationId"`
	Meta             *Meta                          `gorm:"type:json;embedded;column:meta;not null" json:"meta,omitempty"` // the metadata about the most recent change to the row
	Metadata         any                            `gorm:"type:json" json:"metadata,omitempty"`                           // metadata reserved for customers to control
	Note             string                         `gorm:"not null" json:"note"`
	Number           int64                          `gorm:"not null" json:"number"`
	OrderId          *string                        `json:"orderId"`
	RateCents        *int64                         `json:"rateCents"`
	ServiceId        *string                        `json:"serviceId"`
	TechnicianId     string                         `gorm:"not null" json:"technicianId"`
	Type             TimesheetTypeEnum              `gorm:"not null" json:"type"`
	UpdatedDate      *time.Time                     `gorm:"column:updatedDate" json:"updatedDate"`
}

var _ Model = (*Timesheet)(nil)

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *Timesheet) TableName() string {
	return "timesheet"
}

func (m *Timesheet) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewTimesheet returns a new model instance from a json key/value map
func NewTimesheet(kv map[string]any) (*Timesheet, error) {
	var result Timesheet
	err := mapstructure.Decode(kv, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
