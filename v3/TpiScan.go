// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	mapstructure "github.com/mitchellh/mapstructure"
	"time"
)

type TpiScan struct {
	CompanyId              string     `gorm:"not null" json:"companyId"`
	CreatedDate            time.Time  `gorm:"column:createdDate;not null" json:"createdDate"`
	ID                     string     `gorm:"primaryKey;not null" json:"id"`
	IsAlignmentRecommended bool       `gorm:"not null" json:"isAlignmentRecommended"`
	IsRotationRecommended  bool       `gorm:"not null" json:"isRotationRecommended"`
	LocationId             string     `gorm:"not null" json:"locationId"`
	Meta                   *Meta      `gorm:"type:json;embedded;column:meta;not null" json:"meta,omitempty"` // the metadata about the most recent change to the row
	Metadata               any        `gorm:"type:json" json:"metadata,omitempty"`                           // metadata reserved for customers to control
	PdfUrl                 string     `gorm:"not null" json:"pdfUrl"`
	ScanDate               time.Time  `gorm:"not null" json:"scanDate"`
	ScannedBy              *string    `json:"scannedBy"`
	UpdatedDate            *time.Time `gorm:"column:updatedDate" json:"updatedDate"`
	VehicleConditions      []string   `gorm:"not null" json:"vehicleConditions"`
	VehicleId              *string    `json:"vehicleId"`
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
