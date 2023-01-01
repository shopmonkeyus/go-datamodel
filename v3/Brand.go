// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	datatypes "github.com/shopmonkeyus/go-datamodel/datatypes"
)

type Brand struct {
	ID          string              `gorm:"primaryKey;not null;column:id" json:"id"`
	CreatedDate datatypes.DateTime  `gorm:"column:createdDate;not null;column:createdDate" json:"createdDate"`
	UpdatedDate *datatypes.DateTime `gorm:"column:updatedDate;column:updatedDate" json:"updatedDate"`
	Meta        datatypes.Meta      `gorm:"column:meta;not null;column:meta" json:"meta,omitempty"`    // the metadata about the most recent change to the row
	Metadata    *datatypes.JSON     `gorm:"column:metadata;column:metadata" json:"metadata,omitempty"` // metadata reserved for customers to control
	CompanyID   string              `gorm:"not null;column:companyId" json:"companyId"`
	LocationID  string              `gorm:"not null;column:locationId" json:"locationId"`

	Name string `gorm:"not null;column:name" json:"name"`
}

var _ Model = (*Brand)(nil)

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *Brand) TableName() string {
	return "brand"
}

// String returns a string representation as JSON for this model
func (m *Brand) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewBrand returns a new model instance from an encoded buffer
func NewBrand(buf []byte) (*Brand, error) {
	var result Brand
	err := json.Unmarshal(buf, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// NewBrandFromChangeEvent returns a new model instance from an encoded buffer as change event
func NewBrandFromChangeEvent(buf []byte, gzip bool) (*datatypes.ChangeEvent[Brand], error) {
	var result datatypes.ChangeEvent[Brand]
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
