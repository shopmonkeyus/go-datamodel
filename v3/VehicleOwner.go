// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	"time"
)

type VehicleOwner struct {
	CreatedDate time.Time  `gorm:"column:createdDate;not null;column:createdDate" json:"createdDate"`
	UpdatedDate *time.Time `gorm:"column:updatedDate;column:updatedDate" json:"updatedDate"`
	Meta        *Meta      `gorm:"type:json;embedded;serializer:json;column:meta;not null;column:meta" json:"meta,omitempty"` // the metadata about the most recent change to the row
	Metadata    any        `gorm:"type:json;serializer:json;column:metadata" json:"metadata,omitempty"`                       // metadata reserved for customers to control
	CompanyID   string     `gorm:"not null;column:companyId" json:"companyId"`

	CustomerID string `gorm:"not null;column:customerId" json:"customerId"`
	VehicleID  string `gorm:"not null;column:vehicleId" json:"vehicleId"`
}

var _ Model = (*VehicleOwner)(nil)

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *VehicleOwner) TableName() string {
	return "vehicle_owner"
}

func (m *VehicleOwner) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewVehicleOwner returns a new model instance from a json key/value map
func NewVehicleOwner(buf []byte) (*VehicleOwner, error) {
	var result VehicleOwner
	err := json.Unmarshal(buf, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
