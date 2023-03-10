// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	datatypes "github.com/shopmonkeyus/go-datamodel/datatypes"
)

// TpiScanInspection schema
type TpiScanInspection struct {
	ID          string              `bson:"_id" gorm:"primaryKey;not null;column:id" json:"id"`
	CreatedDate datatypes.DateTime  `gorm:"column:createdDate;not null;column:createdDate" json:"createdDate"`
	UpdatedDate *datatypes.DateTime `gorm:"column:updatedDate;column:updatedDate" json:"updatedDate"`
	Meta        *datatypes.JSON     `gorm:"column:meta;not null;column:meta" json:"meta,omitempty"` // the metadata about the most recent change to the row
	Metadata    *datatypes.JSON     `gorm:"column:metadata;column:metadata" json:"metadata,omitempty"`
	CompanyID   string              `gorm:"not null;column:companyId" json:"companyId"`
	LocationID  string              `gorm:"not null;column:locationId" json:"locationId"`

	Name      string  `gorm:"not null;column:name" json:"name"`
	OrderID   string  `gorm:"not null;column:orderId" json:"orderId"`
	Ordinal   float64 `gorm:"not null;column:ordinal" json:"ordinal"`
	TpiScanID string  `gorm:"not null;column:tpiScanId" json:"tpiScanId"`
}

var _ Model = (*TpiScanInspection)(nil)

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *TpiScanInspection) TableName() string {
	return "tpi_scan_inspection"
}

// PrimaryKeys returns an array of the primary keys for this model
func (m *TpiScanInspection) PrimaryKeys() []string {
	return []string{"id"}
}

// String returns a string representation as JSON for this model
func (m *TpiScanInspection) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewTpiScanInspection returns a new model instance from an encoded buffer
func NewTpiScanInspection(buf []byte) (*TpiScanInspection, error) {
	var result TpiScanInspection
	err := json.Unmarshal(buf, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// NewTpiScanInspectionFromChangeEvent returns a new model instance from an encoded buffer as change event
func NewTpiScanInspectionFromChangeEvent(buf []byte, gzip bool) (*datatypes.ChangeEvent[TpiScanInspection], error) {
	var result datatypes.ChangeEvent[TpiScanInspection]
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
