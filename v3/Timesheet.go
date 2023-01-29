// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	datatypes "github.com/shopmonkeyus/go-datamodel/datatypes"
)

// Timesheet schema
type TimesheetActivityEnum string

const (
	TimesheetActivityGeneral TimesheetActivityEnum = "General"
	TimesheetActivityOrder   TimesheetActivityEnum = "Order"
	TimesheetActivityService TimesheetActivityEnum = "Service"
	TimesheetActivityLabor   TimesheetActivityEnum = "Labor"
)

type TimesheetClockInPlatformEnum string

const (
	TimesheetClockInPlatformWeb    TimesheetClockInPlatformEnum = "Web"
	TimesheetClockInPlatformMobile TimesheetClockInPlatformEnum = "Mobile"
)

type TimesheetClockOutPlatformEnum string

const (
	TimesheetClockOutPlatformWeb    TimesheetClockOutPlatformEnum = "Web"
	TimesheetClockOutPlatformMobile TimesheetClockOutPlatformEnum = "Mobile"
)

type TimesheetTypeEnum string

const (
	TimesheetTypeTimeclock TimesheetTypeEnum = "Timeclock"
	TimesheetTypeManual    TimesheetTypeEnum = "Manual"
)

type Timesheet struct {
	ID          string              `bson:"_id" gorm:"primaryKey;not null;column:id" json:"id"`
	CreatedDate datatypes.DateTime  `gorm:"column:createdDate;not null;column:createdDate" json:"createdDate"`
	UpdatedDate *datatypes.DateTime `gorm:"column:updatedDate;column:updatedDate" json:"updatedDate"`
	Meta        *datatypes.JSON     `gorm:"column:meta;not null;column:meta" json:"meta,omitempty"`    // the metadata about the most recent change to the row
	Metadata    *datatypes.JSON     `gorm:"column:metadata;column:metadata" json:"metadata,omitempty"` // metadata reserved for customers to control
	CompanyID   string              `gorm:"not null;column:companyId" json:"companyId"`
	LocationID  string              `gorm:"not null;column:locationId" json:"locationId"`

	Activity         TimesheetActivityEnum          `gorm:"not null;column:activity" json:"activity"`
	ClockIn          datatypes.DateTime             `gorm:"not null;column:clockIn" json:"clockIn"`
	ClockInPlatform  *TimesheetClockInPlatformEnum  `gorm:"column:clockInPlatform" json:"clockInPlatform"`
	ClockOut         *datatypes.DateTime            `gorm:"column:clockOut" json:"clockOut"`
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

// PrimaryKeys returns an array of the primary keys for this model
func (m *Timesheet) PrimaryKeys() []string {
	return []string{"id"}
}

// String returns a string representation as JSON for this model
func (m *Timesheet) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewTimesheet returns a new model instance from an encoded buffer
func NewTimesheet(buf []byte) (*Timesheet, error) {
	var result Timesheet
	err := json.Unmarshal(buf, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// NewTimesheetFromChangeEvent returns a new model instance from an encoded buffer as change event
func NewTimesheetFromChangeEvent(buf []byte, gzip bool) (*datatypes.ChangeEvent[Timesheet], error) {
	var result datatypes.ChangeEvent[Timesheet]
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
