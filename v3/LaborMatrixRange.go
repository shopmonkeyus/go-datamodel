// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	"time"
)

type LaborMatrixRange struct {
	ID            string     `gorm:"primaryKey;not null" json:"id"`
	CompanyId     string     `gorm:"not null" json:"companyId"`
	CreatedDate   time.Time  `gorm:"column:createdDate;not null" json:"createdDate"`
	End           *float64   `json:"end"`
	LaborMatrixId string     `gorm:"not null" json:"laborMatrixId"`
	Meta          *Meta      `gorm:"embedded;column:meta;not null" json:"meta,omitempty"`         // the metadata about the most recent change to the row
	Metadata      any        `gorm:"embedded;column:metadata;not null" json:"metadata,omitempty"` // metadata reserved for customers to control
	Multiplier    float64    `gorm:"not null" json:"multiplier"`
	Start         float64    `gorm:"not null" json:"start"`
	UpdatedDate   *time.Time `gorm:"column:updatedDate" json:"updatedDate"`
}

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *LaborMatrixRange) TableName() string {
	return "labor_matrix_range"
}

func (m *LaborMatrixRange) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}
