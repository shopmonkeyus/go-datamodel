// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	datatypes "github.com/shopmonkeyus/go-datamodel/datatypes"
)

// PricingMatrixRange schema
type PricingMatrixRange struct {
	ID          string              `bson:"_id" gorm:"primaryKey;not null;column:id" json:"id"`
	CreatedDate datatypes.DateTime  `gorm:"column:createdDate;not null;column:createdDate" json:"createdDate"`
	UpdatedDate *datatypes.DateTime `gorm:"column:updatedDate;column:updatedDate" json:"updatedDate"`
	Meta        *datatypes.JSON     `gorm:"column:meta;not null;column:meta" json:"meta,omitempty"`    // the metadata about the most recent change to the row
	Metadata    *datatypes.JSON     `gorm:"column:metadata;column:metadata" json:"metadata,omitempty"` // metadata reserved for customers to control
	CompanyID   string              `gorm:"not null;column:companyId" json:"companyId"`

	EndCents        *int64  `gorm:"column:endCents" json:"endCents"`
	MarginPercent   float64 `gorm:"not null;column:marginPercent" json:"marginPercent"`
	MarkupPercent   float64 `gorm:"not null;column:markupPercent" json:"markupPercent"`
	PricingMatrixID string  `gorm:"not null;column:pricingMatrixId" json:"pricingMatrixId"`
	StartCents      int64   `gorm:"not null;column:startCents" json:"startCents"`
}

var _ Model = (*PricingMatrixRange)(nil)

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *PricingMatrixRange) TableName() string {
	return "pricing_matrix_range"
}

// PrimaryKeys returns an array of the primary keys for this model
func (m *PricingMatrixRange) PrimaryKeys() []string {
	return []string{"id"}
}

// String returns a string representation as JSON for this model
func (m *PricingMatrixRange) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewPricingMatrixRange returns a new model instance from an encoded buffer
func NewPricingMatrixRange(buf []byte) (*PricingMatrixRange, error) {
	var result PricingMatrixRange
	err := json.Unmarshal(buf, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// NewPricingMatrixRangeFromChangeEvent returns a new model instance from an encoded buffer as change event
func NewPricingMatrixRangeFromChangeEvent(buf []byte, gzip bool) (*datatypes.ChangeEvent[PricingMatrixRange], error) {
	var result datatypes.ChangeEvent[PricingMatrixRange]
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
