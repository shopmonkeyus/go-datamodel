// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	datatypes "github.com/shopmonkeyus/go-datamodel/datatypes"
)

// VehicleOwner schema
type VehicleOwner struct {
	VehicleID  string `gorm:"primaryKey;not null;column:vehicleId" json:"vehicleId"`
	CustomerID string `gorm:"primaryKey;not null;column:customerId" json:"customerId"`

	CompanyID   string              `gorm:"not null;column:companyId" json:"companyId"`
	CreatedDate datatypes.DateTime  `gorm:"column:createdDate;not null;column:createdDate" json:"createdDate"`
	Meta        *datatypes.JSON     `gorm:"column:meta;not null;column:meta" json:"meta,omitempty"` // the metadata about the most recent change to the row
	Metadata    *datatypes.JSON     `gorm:"column:metadata;column:metadata" json:"metadata,omitempty"`
	UpdatedDate *datatypes.DateTime `gorm:"column:updatedDate;column:updatedDate" json:"updatedDate"`
}

var _ Model = (*VehicleOwner)(nil)

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *VehicleOwner) TableName() string {
	return "vehicle_owner"
}

// PrimaryKeys returns an array of the primary keys for this model
func (m *VehicleOwner) PrimaryKeys() []string {
	return []string{"vehicleId", "customerId"}
}

// String returns a string representation as JSON for this model
func (m *VehicleOwner) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewVehicleOwner returns a new model instance from an encoded buffer
func NewVehicleOwner(buf []byte) (*VehicleOwner, error) {
	var result VehicleOwner
	err := json.Unmarshal(buf, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// NewVehicleOwnerFromChangeEvent returns a new model instance from an encoded buffer as change event
func NewVehicleOwnerFromChangeEvent(buf []byte, gzip bool) (*datatypes.ChangeEvent[VehicleOwner], error) {
	var result datatypes.ChangeEvent[VehicleOwner]
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
