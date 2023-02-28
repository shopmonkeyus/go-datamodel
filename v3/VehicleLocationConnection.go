// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	datatypes "github.com/shopmonkeyus/go-datamodel/datatypes"
)

// VehicleLocationConnection schema
type VehicleLocationConnection struct {
	VehicleID  string `gorm:"primaryKey;not null;column:vehicleId" json:"vehicleId"`
	LocationID string `gorm:"primaryKey;not null;column:locationId" json:"locationId"`

	CompanyID   string              `gorm:"not null;column:companyId" json:"companyId"`
	CreatedDate datatypes.DateTime  `gorm:"column:createdDate;not null;column:createdDate" json:"createdDate"`
	Meta        *datatypes.JSON     `gorm:"column:meta;not null;column:meta" json:"meta,omitempty"` // the metadata about the most recent change to the row
	Metadata    *datatypes.JSON     `gorm:"column:metadata;column:metadata" json:"metadata,omitempty"`
	UpdatedDate *datatypes.DateTime `gorm:"column:updatedDate;column:updatedDate" json:"updatedDate"`
}

var _ Model = (*VehicleLocationConnection)(nil)

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *VehicleLocationConnection) TableName() string {
	return "vehicle_location"
}

// PrimaryKeys returns an array of the primary keys for this model
func (m *VehicleLocationConnection) PrimaryKeys() []string {
	return []string{"vehicleId", "locationId"}
}

// String returns a string representation as JSON for this model
func (m *VehicleLocationConnection) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewVehicleLocationConnection returns a new model instance from an encoded buffer
func NewVehicleLocationConnection(buf []byte) (*VehicleLocationConnection, error) {
	var result VehicleLocationConnection
	err := json.Unmarshal(buf, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// NewVehicleLocationConnectionFromChangeEvent returns a new model instance from an encoded buffer as change event
func NewVehicleLocationConnectionFromChangeEvent(buf []byte, gzip bool) (*datatypes.ChangeEvent[VehicleLocationConnection], error) {
	var result datatypes.ChangeEvent[VehicleLocationConnection]
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
