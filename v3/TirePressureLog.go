// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	datatypes "github.com/shopmonkeyus/go-datamodel/datatypes"
)

type TirePressureLog struct {
	ID          string              `gorm:"primaryKey;not null;column:id" json:"id"`
	CreatedDate datatypes.DateTime  `gorm:"column:createdDate;not null;column:createdDate" json:"createdDate"`
	UpdatedDate *datatypes.DateTime `gorm:"column:updatedDate;column:updatedDate" json:"updatedDate"`
	Meta        *datatypes.JSON     `gorm:"column:meta;not null;column:meta" json:"meta,omitempty"`    // the metadata about the most recent change to the row
	Metadata    *datatypes.JSON     `gorm:"column:metadata;column:metadata" json:"metadata,omitempty"` // metadata reserved for customers to control
	CompanyID   string              `gorm:"not null;column:companyId" json:"companyId"`
	LocationID  string              `gorm:"not null;column:locationId" json:"locationId"`

	FrontLeft  *float64 `gorm:"column:frontLeft" json:"frontLeft"`
	FrontRight *float64 `gorm:"column:frontRight" json:"frontRight"`
	OrderID    string   `gorm:"not null;column:orderId" json:"orderId"`
	RearLeft   *float64 `gorm:"column:rearLeft" json:"rearLeft"`
	RearRight  *float64 `gorm:"column:rearRight" json:"rearRight"`
	VehicleID  string   `gorm:"not null;column:vehicleId" json:"vehicleId"`
}

var _ Model = (*TirePressureLog)(nil)

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *TirePressureLog) TableName() string {
	return "tire_pressure_log"
}

// PrimaryKeys returns an array of the primary keys for this model
func (m *TirePressureLog) PrimaryKeys() []string {
	return []string{"id"}
}

// String returns a string representation as JSON for this model
func (m *TirePressureLog) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewTirePressureLog returns a new model instance from an encoded buffer
func NewTirePressureLog(buf []byte) (*TirePressureLog, error) {
	var result TirePressureLog
	err := json.Unmarshal(buf, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// NewTirePressureLogFromChangeEvent returns a new model instance from an encoded buffer as change event
func NewTirePressureLogFromChangeEvent(buf []byte, gzip bool) (*datatypes.ChangeEvent[TirePressureLog], error) {
	var result datatypes.ChangeEvent[TirePressureLog]
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
