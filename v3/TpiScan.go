// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	datatypes "github.com/shopmonkeyus/go-datamodel/datatypes"
)

type TpiScan struct {
	ID          string              `gorm:"primaryKey;not null;column:id" json:"id"`
	CreatedDate datatypes.DateTime  `gorm:"column:createdDate;not null;column:createdDate" json:"createdDate"`
	UpdatedDate *datatypes.DateTime `gorm:"column:updatedDate;column:updatedDate" json:"updatedDate"`
	Meta        datatypes.Meta      `gorm:"column:meta;not null;column:meta" json:"meta,omitempty"`    // the metadata about the most recent change to the row
	Metadata    *datatypes.JSON     `gorm:"column:metadata;column:metadata" json:"metadata,omitempty"` // metadata reserved for customers to control
	CompanyID   string              `gorm:"not null;column:companyId" json:"companyId"`
	LocationID  string              `gorm:"not null;column:locationId" json:"locationId"`

	IsAlignmentRecommended bool                  `gorm:"not null;column:isAlignmentRecommended" json:"isAlignmentRecommended"`
	IsRotationRecommended  bool                  `gorm:"not null;column:isRotationRecommended" json:"isRotationRecommended"`
	PdfUrl                 string                `gorm:"not null;column:pdfUrl" json:"pdfUrl"`
	ScanDate               datatypes.DateTime    `gorm:"not null;column:scanDate" json:"scanDate"`
	ScannedBy              *string               `gorm:"column:scannedBy" json:"scannedBy"`
	VehicleConditions      datatypes.StringArray `gorm:"not null;column:vehicleConditions" json:"vehicleConditions"`
	VehicleID              *string               `gorm:"column:vehicleId" json:"vehicleId"`
}

var _ Model = (*TpiScan)(nil)

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *TpiScan) TableName() string {
	return "tpi_scan"
}

// String returns a string representation as JSON for this model
func (m *TpiScan) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewTpiScan returns a new model instance from an encoded buffer
func NewTpiScan(buf []byte) (*TpiScan, error) {
	var result TpiScan
	err := json.Unmarshal(buf, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// NewTpiScanFromChangeEvent returns a new model instance from an encoded buffer as change event
func NewTpiScanFromChangeEvent(buf []byte, gzip bool) (*datatypes.ChangeEvent[TpiScan], error) {
	var result datatypes.ChangeEvent[TpiScan]
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
