// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	"time"
)

type TpiScanInspection struct {
	ID          string     `gorm:"primaryKey;not null;column:id" json:"id"`
	CreatedDate time.Time  `gorm:"column:createdDate;not null;column:createdDate" json:"createdDate"`
	UpdatedDate *time.Time `gorm:"column:updatedDate;column:updatedDate" json:"updatedDate"`
	Meta        *Meta      `gorm:"type:json;embedded;serializer:json;column:meta;not null;column:meta" json:"meta,omitempty"` // the metadata about the most recent change to the row
	Metadata    any        `gorm:"type:json;serializer:json;column:metadata" json:"metadata,omitempty"`                       // metadata reserved for customers to control
	CompanyID   string     `gorm:"not null;column:companyId" json:"companyId"`
	LocationID  string     `gorm:"not null;column:locationId" json:"locationId"`

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

func (m *TpiScanInspection) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewTpiScanInspection returns a new model instance from a json key/value map
func NewTpiScanInspection(buf []byte) (*TpiScanInspection, error) {
	var result TpiScanInspection
	err := json.Unmarshal(buf, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
