// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	mapstructure "github.com/mitchellh/mapstructure"
	"time"
)

type TpiScan struct {
	ID          string     `gorm:"primaryKey;not null;column:id" json:"id"`
	CreatedDate time.Time  `gorm:"column:createdDate;not null;column:createdDate" json:"createdDate"`
	UpdatedDate *time.Time `gorm:"column:updatedDate;column:updatedDate" json:"updatedDate"`
	Meta        *Meta      `gorm:"type:json;embedded;serializer:json;column:meta;not null;column:meta" json:"meta,omitempty"` // the metadata about the most recent change to the row
	Metadata    any        `gorm:"type:json;serializer:json;column:metadata" json:"metadata,omitempty"`                       // metadata reserved for customers to control
	CompanyID   string     `gorm:"not null;column:companyId" json:"companyId"`
	LocationID  string     `gorm:"not null;column:locationId" json:"locationId"`

	IsAlignmentRecommended bool      `gorm:"not null;column:isAlignmentRecommended" json:"isAlignmentRecommended"`
	IsRotationRecommended  bool      `gorm:"not null;column:isRotationRecommended" json:"isRotationRecommended"`
	PdfUrl                 string    `gorm:"not null;column:pdfUrl" json:"pdfUrl"`
	ScanDate               time.Time `gorm:"not null;column:scanDate" json:"scanDate"`
	ScannedBy              *string   `gorm:"column:scannedBy" json:"scannedBy"`
	VehicleConditions      []string  `gorm:"not null;column:vehicleConditions" json:"vehicleConditions"`
	VehicleID              *string   `gorm:"column:vehicleId" json:"vehicleId"`
}

var _ Model = (*TpiScan)(nil)

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *TpiScan) TableName() string {
	return "tpi_scan"
}

func (m *TpiScan) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewTpiScan returns a new model instance from a json key/value map
func NewTpiScan(kv map[string]any) (*TpiScan, error) {
	var result TpiScan
	err := mapstructure.Decode(kv, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
