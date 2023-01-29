// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	datatypes "github.com/shopmonkeyus/go-datamodel/datatypes"
)

// MileageLog schema
type MileageLogTypeEnum string

const (
	MileageLogTypeManualEntry MileageLogTypeEnum = "ManualEntry"
	MileageLogTypeOrderIn     MileageLogTypeEnum = "OrderIn"
	MileageLogTypeOrderOut    MileageLogTypeEnum = "OrderOut"
)

type MileageLog struct {
	ID          string              `bson:"_id" gorm:"primaryKey;not null;column:id" json:"id"`
	CreatedDate datatypes.DateTime  `gorm:"column:createdDate;not null;column:createdDate" json:"createdDate"`
	UpdatedDate *datatypes.DateTime `gorm:"column:updatedDate;column:updatedDate" json:"updatedDate"`
	Meta        *datatypes.JSON     `gorm:"column:meta;not null;column:meta" json:"meta,omitempty"`    // the metadata about the most recent change to the row
	Metadata    *datatypes.JSON     `gorm:"column:metadata;column:metadata" json:"metadata,omitempty"` // metadata reserved for customers to control
	CompanyID   string              `gorm:"not null;column:companyId" json:"companyId"`
	LocationID  string              `gorm:"not null;column:locationId" json:"locationId"`

	Mileage     float64            `gorm:"not null;column:mileage" json:"mileage"`
	MileageDate datatypes.DateTime `gorm:"not null;column:mileageDate" json:"mileageDate"`
	OrderID     *string            `gorm:"column:orderId" json:"orderId"`
	Type        MileageLogTypeEnum `gorm:"not null;column:type" json:"type"`
	VehicleID   string             `gorm:"not null;column:vehicleId" json:"vehicleId"`
}

var _ Model = (*MileageLog)(nil)

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *MileageLog) TableName() string {
	return "mileage_log"
}

// PrimaryKeys returns an array of the primary keys for this model
func (m *MileageLog) PrimaryKeys() []string {
	return []string{"id"}
}

// String returns a string representation as JSON for this model
func (m *MileageLog) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewMileageLog returns a new model instance from an encoded buffer
func NewMileageLog(buf []byte) (*MileageLog, error) {
	var result MileageLog
	err := json.Unmarshal(buf, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// NewMileageLogFromChangeEvent returns a new model instance from an encoded buffer as change event
func NewMileageLogFromChangeEvent(buf []byte, gzip bool) (*datatypes.ChangeEvent[MileageLog], error) {
	var result datatypes.ChangeEvent[MileageLog]
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
