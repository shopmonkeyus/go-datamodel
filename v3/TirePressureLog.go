// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	mapstructure "github.com/mitchellh/mapstructure"
	"time"
)

type TirePressureLog struct {
	ID          string     `gorm:"primaryKey;not null;column:id" json:"id"`
	CreatedDate time.Time  `gorm:"column:createdDate;not null;column:createdDate" json:"createdDate"`
	UpdatedDate *time.Time `gorm:"column:updatedDate;column:updatedDate" json:"updatedDate"`
	Meta        *Meta      `gorm:"type:json;embedded;serializer:json;column:meta;not null;column:meta" json:"meta,omitempty"` // the metadata about the most recent change to the row
	Metadata    any        `gorm:"type:json;serializer:json;column:metadata" json:"metadata,omitempty"`                       // metadata reserved for customers to control
	CompanyID   string     `gorm:"not null;column:companyId" json:"companyId"`
	LocationID  string     `gorm:"not null;column:locationId" json:"locationId"`

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

func (m *TirePressureLog) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewTirePressureLog returns a new model instance from a json key/value map
func NewTirePressureLog(kv map[string]any) (*TirePressureLog, error) {
	var result TirePressureLog
	err := mapstructure.Decode(kv, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
