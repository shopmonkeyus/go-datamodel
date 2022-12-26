// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	"time"
)

type LaborMatrixRange struct {
	ID          string     `gorm:"primaryKey;not null;column:id" json:"id"`
	CreatedDate time.Time  `gorm:"column:createdDate;not null;column:createdDate" json:"createdDate"`
	UpdatedDate *time.Time `gorm:"column:updatedDate;column:updatedDate" json:"updatedDate"`
	Meta        *Meta      `gorm:"type:json;embedded;serializer:json;column:meta;not null;column:meta" json:"meta,omitempty"` // the metadata about the most recent change to the row
	Metadata    any        `gorm:"type:json;serializer:json;column:metadata" json:"metadata,omitempty"`                       // metadata reserved for customers to control
	CompanyID   string     `gorm:"not null;column:companyId" json:"companyId"`

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

func (m *LaborMatrixRange) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewLaborMatrixRange returns a new model instance from a json key/value map
func NewLaborMatrixRange(buf []byte) (*LaborMatrixRange, error) {
	var result LaborMatrixRange
	err := json.Unmarshal(buf, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
