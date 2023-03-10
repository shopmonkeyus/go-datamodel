// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	datatypes "github.com/shopmonkeyus/go-datamodel/datatypes"
)

// LaborMatrix schema
type LaborMatrixTypeEnum string

const (
	LaborMatrixTypeHours LaborMatrixTypeEnum = "Hours"
	LaborMatrixTypeRate  LaborMatrixTypeEnum = "Rate"
)

type LaborMatrix struct {
	ID          string              `bson:"_id" gorm:"primaryKey;not null;column:id" json:"id"`
	CreatedDate datatypes.DateTime  `gorm:"column:createdDate;not null;column:createdDate" json:"createdDate"`
	UpdatedDate *datatypes.DateTime `gorm:"column:updatedDate;column:updatedDate" json:"updatedDate"`
	Meta        *datatypes.JSON     `gorm:"column:meta;not null;column:meta" json:"meta,omitempty"` // the metadata about the most recent change to the row
	Metadata    *datatypes.JSON     `gorm:"column:metadata;column:metadata" json:"metadata,omitempty"`
	CompanyID   string              `gorm:"not null;column:companyId" json:"companyId"`
	LocationID  string              `gorm:"not null;column:locationId" json:"locationId"`

	Default bool                `gorm:"not null;column:default" json:"default"`
	Name    string              `gorm:"not null;column:name" json:"name"`
	Type    LaborMatrixTypeEnum `gorm:"not null;column:type" json:"type"`
}

var _ Model = (*LaborMatrix)(nil)

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *LaborMatrix) TableName() string {
	return "labor_matrix"
}

// PrimaryKeys returns an array of the primary keys for this model
func (m *LaborMatrix) PrimaryKeys() []string {
	return []string{"id"}
}

// String returns a string representation as JSON for this model
func (m *LaborMatrix) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewLaborMatrix returns a new model instance from an encoded buffer
func NewLaborMatrix(buf []byte) (*LaborMatrix, error) {
	var result LaborMatrix
	err := json.Unmarshal(buf, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// NewLaborMatrixFromChangeEvent returns a new model instance from an encoded buffer as change event
func NewLaborMatrixFromChangeEvent(buf []byte, gzip bool) (*datatypes.ChangeEvent[LaborMatrix], error) {
	var result datatypes.ChangeEvent[LaborMatrix]
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
