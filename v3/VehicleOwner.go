// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	"time"
)

type VehicleOwner struct {
	CompanyId   string     `gorm:"not null" json:"companyId"`
	CreatedDate time.Time  `gorm:"column:createdDate;not null" json:"createdDate"`
	CustomerId  string     `gorm:"not null" json:"customerId"`
	Meta        *Meta      `gorm:"embedded;column:meta;not null" json:"meta,omitempty"`         // the metadata about the most recent change to the row
	Metadata    any        `gorm:"embedded;column:metadata;not null" json:"metadata,omitempty"` // metadata reserved for customers to control
	UpdatedDate *time.Time `gorm:"column:updatedDate" json:"updatedDate"`
	VehicleId   string     `gorm:"not null" json:"vehicleId"`
}

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *VehicleOwner) TableName() string {
	return "vehicle_owner"
}

func (m *VehicleOwner) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}
