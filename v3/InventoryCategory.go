// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	datatypes "github.com/shopmonkeyus/go-datamodel/datatypes"
)

type InventoryCategory struct {
	ID          string              `gorm:"primaryKey;not null;column:id" json:"id"`
	CreatedDate datatypes.DateTime  `gorm:"column:createdDate;not null;column:createdDate" json:"createdDate"`
	UpdatedDate *datatypes.DateTime `gorm:"column:updatedDate;column:updatedDate" json:"updatedDate"`
	Meta        datatypes.Meta      `gorm:"column:meta;not null;column:meta" json:"meta,omitempty"`    // the metadata about the most recent change to the row
	Metadata    *datatypes.JSON     `gorm:"column:metadata;column:metadata" json:"metadata,omitempty"` // metadata reserved for customers to control
	CompanyID   string              `gorm:"not null;column:companyId" json:"companyId"`
	LocationID  string              `gorm:"not null;column:locationId" json:"locationId"`

	Name     string  `gorm:"not null;column:name" json:"name"`
	ParentID *string `gorm:"column:parentId" json:"parentId"`
}

var _ Model = (*InventoryCategory)(nil)

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *InventoryCategory) TableName() string {
	return "inventory_category"
}

// String returns a string representation as JSON for this model
func (m *InventoryCategory) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewInventoryCategory returns a new model instance from an encoded buffer
func NewInventoryCategory(buf []byte) (*InventoryCategory, error) {
	var result InventoryCategory
	err := json.Unmarshal(buf, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// NewInventoryCategoryFromChangeEvent returns a new model instance from an encoded buffer as change event
func NewInventoryCategoryFromChangeEvent(buf []byte, gzip bool) (*datatypes.ChangeEvent[InventoryCategory], error) {
	var result datatypes.ChangeEvent[InventoryCategory]
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
