// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	mapstructure "github.com/mitchellh/mapstructure"
	"time"
)

type Brand struct {
	ID          string     `gorm:"primaryKey;not null" json:"id"`
	CompanyId   string     `gorm:"not null" json:"companyId"`
	CreatedDate time.Time  `gorm:"column:createdDate;not null" json:"createdDate"`
	LocationId  string     `gorm:"not null" json:"locationId"`
	Meta        *Meta      `gorm:"type:json;embedded;column:meta;not null" json:"meta,omitempty"` // the metadata about the most recent change to the row
	Metadata    any        `gorm:"type:json" json:"metadata,omitempty"`                           // metadata reserved for customers to control
	Name        string     `gorm:"not null" json:"name"`
	UpdatedDate *time.Time `gorm:"column:updatedDate" json:"updatedDate"`
}

var _ Model = (*Brand)(nil)

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *Brand) TableName() string {
	return "brand"
}

func (m *Brand) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewBrand returns a new model instance from a json key/value map
func NewBrand(kv map[string]any) (*Brand, error) {
	var result Brand
	err := mapstructure.Decode(kv, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
