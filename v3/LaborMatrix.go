// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	mapstructure "github.com/mitchellh/mapstructure"
	"time"
)

type LaborMatrixTypeEnum string

const (
	LaborMatrixTypeHours LaborMatrixTypeEnum = "Hours"
	LaborMatrixTypeRate                      = "Rate"
)

type LaborMatrix struct {
	CompanyId   string              `gorm:"not null" json:"companyId"`
	CreatedDate time.Time           `gorm:"column:createdDate;not null" json:"createdDate"`
	ID          string              `gorm:"primaryKey;not null" json:"id"`
	Default     bool                `gorm:"not null" json:"default"`
	LocationId  string              `gorm:"not null" json:"locationId"`
	Meta        *Meta               `gorm:"type:json;embedded;column:meta;not null" json:"meta,omitempty"` // the metadata about the most recent change to the row
	Metadata    any                 `gorm:"type:json" json:"metadata,omitempty"`                           // metadata reserved for customers to control
	Name        string              `gorm:"not null" json:"name"`
	Type        LaborMatrixTypeEnum `gorm:"not null" json:"type"`
	UpdatedDate *time.Time          `gorm:"column:updatedDate" json:"updatedDate"`
}

var _ Model = (*LaborMatrix)(nil)

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *LaborMatrix) TableName() string {
	return "labor_matrix"
}

func (m *LaborMatrix) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewLaborMatrix returns a new model instance from a json key/value map
func NewLaborMatrix(kv map[string]any) (*LaborMatrix, error) {
	var result LaborMatrix
	err := mapstructure.Decode(kv, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
