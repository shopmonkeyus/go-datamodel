// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	mapstructure "github.com/mitchellh/mapstructure"
	"time"
)

type LaborMatrixRange struct {
	CompanyId     string     `gorm:"not null" json:"companyId"`
	ID            string     `gorm:"primaryKey;not null" json:"id"`
	CreatedDate   time.Time  `gorm:"column:createdDate;not null" json:"createdDate"`
	End           *float64   `json:"end"`
	LaborMatrixId string     `gorm:"not null" json:"laborMatrixId"`
	Meta          *Meta      `gorm:"type:json;embedded;column:meta;not null" json:"meta,omitempty"` // the metadata about the most recent change to the row
	Metadata      any        `gorm:"type:json" json:"metadata,omitempty"`                           // metadata reserved for customers to control
	Multiplier    float64    `gorm:"not null" json:"multiplier"`
	Start         float64    `gorm:"not null" json:"start"`
	UpdatedDate   *time.Time `gorm:"column:updatedDate" json:"updatedDate"`
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
func NewLaborMatrixRange(kv map[string]any) (*LaborMatrixRange, error) {
	var result LaborMatrixRange
	err := mapstructure.Decode(kv, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
