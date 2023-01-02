// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	datatypes "github.com/shopmonkeyus/go-datamodel/datatypes"
)

type PricingMatrix struct {
	ID          string              `bson:"_id" gorm:"primaryKey;not null;column:id" json:"id"`
	CreatedDate datatypes.DateTime  `gorm:"column:createdDate;not null;column:createdDate" json:"createdDate"`
	UpdatedDate *datatypes.DateTime `gorm:"column:updatedDate;column:updatedDate" json:"updatedDate"`
	Meta        *datatypes.JSON     `gorm:"column:meta;not null;column:meta" json:"meta,omitempty"`    // the metadata about the most recent change to the row
	Metadata    *datatypes.JSON     `gorm:"column:metadata;column:metadata" json:"metadata,omitempty"` // metadata reserved for customers to control
	CompanyID   string              `gorm:"not null;column:companyId" json:"companyId"`
	LocationID  string              `gorm:"not null;column:locationId" json:"locationId"`

	Default bool   `gorm:"not null;column:default" json:"default"`
	Name    string `gorm:"not null;column:name" json:"name"`
}

var _ Model = (*PricingMatrix)(nil)

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *PricingMatrix) TableName() string {
	return "pricing_matrix"
}

// PrimaryKeys returns an array of the primary keys for this model
func (m *PricingMatrix) PrimaryKeys() []string {
	return []string{"id"}
}

// String returns a string representation as JSON for this model
func (m *PricingMatrix) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewPricingMatrix returns a new model instance from an encoded buffer
func NewPricingMatrix(buf []byte) (*PricingMatrix, error) {
	var result PricingMatrix
	err := json.Unmarshal(buf, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// NewPricingMatrixFromChangeEvent returns a new model instance from an encoded buffer as change event
func NewPricingMatrixFromChangeEvent(buf []byte, gzip bool) (*datatypes.ChangeEvent[PricingMatrix], error) {
	var result datatypes.ChangeEvent[PricingMatrix]
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
