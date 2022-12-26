// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	mapstructure "github.com/mitchellh/mapstructure"
	"time"
)

type InventoryCategory struct {
	CompanyId   string     `gorm:"not null" json:"companyId"`
	CreatedDate time.Time  `gorm:"column:createdDate;not null" json:"createdDate"`
	ID          string     `gorm:"primaryKey;not null" json:"id"`
	LocationId  string     `gorm:"not null" json:"locationId"`
	Meta        *Meta      `gorm:"type:json;embedded;column:meta;not null" json:"meta,omitempty"` // the metadata about the most recent change to the row
	Metadata    any        `gorm:"type:json" json:"metadata,omitempty"`                           // metadata reserved for customers to control
	Name        string     `gorm:"not null" json:"name"`
	ParentId    *string    `json:"parentId"`
	UpdatedDate *time.Time `gorm:"column:updatedDate" json:"updatedDate"`
}

var _ Model = (*InventoryCategory)(nil)

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *InventoryCategory) TableName() string {
	return "inventory_category"
}

func (m *InventoryCategory) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewInventoryCategory returns a new model instance from a json key/value map
func NewInventoryCategory(kv map[string]any) (*InventoryCategory, error) {
	var result InventoryCategory
	err := mapstructure.Decode(kv, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
