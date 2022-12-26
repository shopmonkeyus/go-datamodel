// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	mapstructure "github.com/mitchellh/mapstructure"
	"time"
)

type VehicleLocationConnection struct {
	CompanyId   string     `gorm:"not null" json:"companyId"`
	CreatedDate time.Time  `gorm:"column:createdDate;not null" json:"createdDate"`
	LocationId  string     `gorm:"not null" json:"locationId"`
	Meta        *Meta      `gorm:"type:json;embedded;column:meta;not null" json:"meta,omitempty"` // the metadata about the most recent change to the row
	Metadata    any        `gorm:"type:json" json:"metadata,omitempty"`                           // metadata reserved for customers to control
	UpdatedDate *time.Time `gorm:"column:updatedDate" json:"updatedDate"`
	VehicleId   string     `gorm:"not null" json:"vehicleId"`
}

var _ Model = (*VehicleLocationConnection)(nil)

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *VehicleLocationConnection) TableName() string {
	return "vehicle_location"
}

func (m *VehicleLocationConnection) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewVehicleLocationConnection returns a new model instance from a json key/value map
func NewVehicleLocationConnection(kv map[string]any) (*VehicleLocationConnection, error) {
	var result VehicleLocationConnection
	err := mapstructure.Decode(kv, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
