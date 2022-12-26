// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	mapstructure "github.com/mitchellh/mapstructure"
	"time"
)

type TirePressureLog struct {
	ID          string     `gorm:"primaryKey;not null" json:"id"`
	CompanyId   string     `gorm:"not null" json:"companyId"`
	CreatedDate time.Time  `gorm:"column:createdDate;not null" json:"createdDate"`
	FrontLeft   *float64   `json:"frontLeft"`
	FrontRight  *float64   `json:"frontRight"`
	LocationId  string     `gorm:"not null" json:"locationId"`
	Meta        *Meta      `gorm:"type:json;embedded;column:meta;not null" json:"meta,omitempty"` // the metadata about the most recent change to the row
	Metadata    any        `gorm:"type:json" json:"metadata,omitempty"`                           // metadata reserved for customers to control
	OrderId     string     `gorm:"not null" json:"orderId"`
	RearLeft    *float64   `json:"rearLeft"`
	RearRight   *float64   `json:"rearRight"`
	UpdatedDate *time.Time `gorm:"column:updatedDate" json:"updatedDate"`
	VehicleId   string     `gorm:"not null" json:"vehicleId"`
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
