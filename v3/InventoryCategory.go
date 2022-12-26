// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	"time"
)

type InventoryCategory struct {
	CompanyId   string     `gorm:"not null" json:"companyId"`
	ID          string     `gorm:"primaryKey;not null" json:"id"`
	CreatedDate time.Time  `gorm:"column:createdDate;not null" json:"createdDate"`
	LocationId  string     `gorm:"not null" json:"locationId"`
	Meta        *Meta      `gorm:"embedded;column:meta;not null" json:"meta,omitempty"`         // the metadata about the most recent change to the row
	Metadata    any        `gorm:"embedded;column:metadata;not null" json:"metadata,omitempty"` // metadata reserved for customers to control
	Name        string     `gorm:"not null" json:"name"`
	ParentId    *string    `json:"parentId"`
	UpdatedDate *time.Time `gorm:"column:updatedDate" json:"updatedDate"`
}

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *InventoryCategory) TableName() string {
	return "inventory_category"
}

func (m *InventoryCategory) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}
