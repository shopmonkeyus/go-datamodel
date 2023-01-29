// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	datatypes "github.com/shopmonkeyus/go-datamodel/datatypes"
)

// LaborRate schema
type LaborRate struct {
	ID          string              `bson:"_id" gorm:"primaryKey;not null;column:id" json:"id"`
	CreatedDate datatypes.DateTime  `gorm:"column:createdDate;not null;column:createdDate" json:"createdDate"`
	UpdatedDate *datatypes.DateTime `gorm:"column:updatedDate;column:updatedDate" json:"updatedDate"`
	Meta        *datatypes.JSON     `gorm:"column:meta;not null;column:meta" json:"meta,omitempty"`    // the metadata about the most recent change to the row
	Metadata    *datatypes.JSON     `gorm:"column:metadata;column:metadata" json:"metadata,omitempty"` // metadata reserved for customers to control
	CompanyID   string              `gorm:"not null;column:companyId" json:"companyId"`
	LocationID  string              `gorm:"not null;column:locationId" json:"locationId"`

	Name       string `gorm:"not null;column:name" json:"name"`
	ValueCents int64  `gorm:"not null;column:valueCents" json:"valueCents"`
}

var _ Model = (*LaborRate)(nil)

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *LaborRate) TableName() string {
	return "labor_rate"
}

// PrimaryKeys returns an array of the primary keys for this model
func (m *LaborRate) PrimaryKeys() []string {
	return []string{"id"}
}

// String returns a string representation as JSON for this model
func (m *LaborRate) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewLaborRate returns a new model instance from an encoded buffer
func NewLaborRate(buf []byte) (*LaborRate, error) {
	var result LaborRate
	err := json.Unmarshal(buf, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// NewLaborRateFromChangeEvent returns a new model instance from an encoded buffer as change event
func NewLaborRateFromChangeEvent(buf []byte, gzip bool) (*datatypes.ChangeEvent[LaborRate], error) {
	var result datatypes.ChangeEvent[LaborRate]
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
