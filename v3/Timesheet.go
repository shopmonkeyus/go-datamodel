// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	codec "github.com/hashicorp/go-msgpack/v2/codec"
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
	ID          string     `gorm:"primaryKey;not null;column:id" json:"id"`
	CreatedDate time.Time  `gorm:"column:createdDate;not null;column:createdDate" json:"createdDate"`
	UpdatedDate *time.Time `gorm:"column:updatedDate;column:updatedDate" json:"updatedDate"`
	Meta        *Meta      `gorm:"type:json;embedded;serializer:json;column:meta;not null;column:meta" json:"meta,omitempty"` // the metadata about the most recent change to the row
	Metadata    any        `gorm:"type:json;serializer:json;column:metadata" json:"metadata,omitempty"`                       // metadata reserved for customers to control
	CompanyID   string     `gorm:"not null;column:companyId" json:"companyId"`
	LocationID  string     `gorm:"not null;column:locationId" json:"locationId"`

	Activity         TimesheetActivityEnum          `gorm:"not null;column:activity" json:"activity"`
	ClockIn          time.Time                      `gorm:"not null;column:clockIn" json:"clockIn"`
	ClockInPlatform  *TimesheetClockInPlatformEnum  `gorm:"column:clockInPlatform" json:"clockInPlatform"`
	ClockOut         *time.Time                     `gorm:"column:clockOut" json:"clockOut"`
	ClockOutPlatform *TimesheetClockOutPlatformEnum `gorm:"column:clockOutPlatform" json:"clockOutPlatform"`
	Duration         *float64                       `gorm:"column:duration" json:"duration"`          // the amount of time clocked (in milliseconds)
	FlatRate         bool                           `gorm:"not null;column:flatRate" json:"flatRate"` // if the technician uses a flat rate
	InProgress       bool                           `gorm:"not null;column:inProgress" json:"inProgress"`
	LaborID          *string                        `gorm:"column:laborId" json:"laborId"`
	Note             string                         `gorm:"not null;column:note" json:"note"`
	Number           int64                          `gorm:"not null;column:number" json:"number"`
	OrderID          *string                        `gorm:"column:orderId" json:"orderId"`
	RateCents        *int64                         `gorm:"column:rateCents" json:"rateCents"`
	ServiceID        *string                        `gorm:"column:serviceId" json:"serviceId"`
	TechnicianID     string                         `gorm:"not null;column:technicianId" json:"technicianId"`
	Type             TimesheetTypeEnum              `gorm:"not null;column:type" json:"type"`
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

// NewTimesheet returns a new model instance from an encoded buffer
func NewTimesheet(buf []byte, enctype EncodingType) (*Timesheet, error) {
	var result Timesheet
	var handle codec.Handle
	if enctype == JSONEncoding {
		handle = &jsonHandle
	} else {
		handle = &msgpackHandle
	}
	dec := codec.NewDecoderBytes(buf, handle)
	err := dec.Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
