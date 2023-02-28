// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	datatypes "github.com/shopmonkeyus/go-datamodel/datatypes"
)

// LaborMatrixRange schema
type LaborMatrixRange struct {
	ID          string              `bson:"_id" gorm:"primaryKey;not null;column:id" json:"id"`
	CreatedDate datatypes.DateTime  `gorm:"column:createdDate;not null;column:createdDate" json:"createdDate"`
	UpdatedDate *datatypes.DateTime `gorm:"column:updatedDate;column:updatedDate" json:"updatedDate"`
	Meta        *datatypes.JSON     `gorm:"column:meta;not null;column:meta" json:"meta,omitempty"` // the metadata about the most recent change to the row
	Metadata    *datatypes.JSON     `gorm:"column:metadata;column:metadata" json:"metadata,omitempty"`
	CompanyID   string              `gorm:"not null;column:companyId" json:"companyId"`

	End           *float64 `gorm:"column:end" json:"end"`
	LaborMatrixID string   `gorm:"not null;column:laborMatrixId" json:"laborMatrixId"`
	Multiplier    float64  `gorm:"not null;column:multiplier" json:"multiplier"`
	Start         float64  `gorm:"not null;column:start" json:"start"`
}

var _ Model = (*LaborMatrixRange)(nil)

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *LaborMatrixRange) TableName() string {
	return "labor_matrix_range"
}

// PrimaryKeys returns an array of the primary keys for this model
func (m *LaborMatrixRange) PrimaryKeys() []string {
	return []string{"id"}
}

// String returns a string representation as JSON for this model
func (m *LaborMatrixRange) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewLaborMatrixRange returns a new model instance from an encoded buffer
func NewLaborMatrixRange(buf []byte) (*LaborMatrixRange, error) {
	var result LaborMatrixRange
	err := json.Unmarshal(buf, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// NewLaborMatrixRangeFromChangeEvent returns a new model instance from an encoded buffer as change event
func NewLaborMatrixRangeFromChangeEvent(buf []byte, gzip bool) (*datatypes.ChangeEvent[LaborMatrixRange], error) {
	var result datatypes.ChangeEvent[LaborMatrixRange]
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
