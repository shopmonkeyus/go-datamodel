// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	datatypes "github.com/shopmonkeyus/go-datamodel/datatypes"
)

// LabelVehicleConnection schema
type LabelVehicleConnection struct {
	LabelID   string `gorm:"primaryKey;not null;column:labelId" json:"labelId"`
	VehicleID string `gorm:"primaryKey;not null;column:vehicleId" json:"vehicleId"`

	CompanyID   string              `gorm:"not null;column:companyId" json:"companyId"`
	CreatedDate *datatypes.DateTime `gorm:"column:createdDate;column:createdDate" json:"createdDate"`
	Meta        *datatypes.JSON     `gorm:"column:meta;not null;column:meta" json:"meta,omitempty"`    // the metadata about the most recent change to the row
	Metadata    *datatypes.JSON     `gorm:"column:metadata;column:metadata" json:"metadata,omitempty"` // metadata reserved for customers to control
	UpdatedDate *datatypes.DateTime `gorm:"column:updatedDate;column:updatedDate" json:"updatedDate"`
}

var _ Model = (*LabelVehicleConnection)(nil)

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *LabelVehicleConnection) TableName() string {
	return "label_vehicle_connection"
}

// PrimaryKeys returns an array of the primary keys for this model
func (m *LabelVehicleConnection) PrimaryKeys() []string {
	return []string{"labelId", "vehicleId"}
}

// String returns a string representation as JSON for this model
func (m *LabelVehicleConnection) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewLabelVehicleConnection returns a new model instance from an encoded buffer
func NewLabelVehicleConnection(buf []byte) (*LabelVehicleConnection, error) {
	var result LabelVehicleConnection
	err := json.Unmarshal(buf, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// NewLabelVehicleConnectionFromChangeEvent returns a new model instance from an encoded buffer as change event
func NewLabelVehicleConnectionFromChangeEvent(buf []byte, gzip bool) (*datatypes.ChangeEvent[LabelVehicleConnection], error) {
	var result datatypes.ChangeEvent[LabelVehicleConnection]
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
