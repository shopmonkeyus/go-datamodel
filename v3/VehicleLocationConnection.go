// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	mapstructure "github.com/mitchellh/mapstructure"
	"time"
)

type VehicleLocationConnection struct {
	CreatedDate time.Time  `gorm:"column:createdDate;not null;column:createdDate" json:"createdDate"`
	UpdatedDate *time.Time `gorm:"column:updatedDate;column:updatedDate" json:"updatedDate"`
	Meta        *Meta      `gorm:"type:json;embedded;serializer:json;column:meta;not null;column:meta" json:"meta,omitempty"` // the metadata about the most recent change to the row
	Metadata    any        `gorm:"type:json;serializer:json;column:metadata" json:"metadata,omitempty"`                       // metadata reserved for customers to control
	CompanyID   string     `gorm:"not null;column:companyId" json:"companyId"`
	LocationID  string     `gorm:"not null;column:locationId" json:"locationId"`

	VehicleID string `gorm:"not null;column:vehicleId" json:"vehicleId"`
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
