// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	mapstructure "github.com/mitchellh/mapstructure"
	"time"
)

type InventoryPart struct {
	BinLocation        *string    `json:"binLocation"`
	CategoryId         *string    `json:"categoryId"`
	CompanyId          string     `gorm:"not null" json:"companyId"`
	CreatedDate        time.Time  `gorm:"column:createdDate;not null" json:"createdDate"`
	Deleted            bool       `gorm:"not null" json:"deleted"` // if the record has been deleted
	DeletedDate        *time.Time `json:"deletedDate"`             // the date that the record was deleted or null if not deleted
	DeletedReason      *string    `json:"deletedReason"`           // the reason that the record was deleted
	DeletedUserId      *string    `json:"deletedUserId"`           // the user that deleted the record or null if not deleted
	ID                 string     `gorm:"primaryKey;not null" json:"id"`
	LocationId         string     `gorm:"not null" json:"locationId"`
	Meta               *Meta      `gorm:"type:json;embedded;column:meta;not null" json:"meta,omitempty"` // the metadata about the most recent change to the row
	Metadata           any        `gorm:"type:json" json:"metadata,omitempty"`                           // metadata reserved for customers to control
	Name               string     `gorm:"not null" json:"name"`
	Note               *string    `json:"note"`
	Number             string     `gorm:"not null" json:"number"`
	Quantity           int64      `gorm:"not null" json:"quantity"`
	RetailCostCents    int64      `gorm:"not null" json:"retailCostCents"`
	Taxable            bool       `gorm:"not null" json:"taxable"`
	UpdatedDate        *time.Time `gorm:"column:updatedDate" json:"updatedDate"`
	UserId             *string    `json:"userId"`
	VendorId           *string    `json:"vendorId"`
	WholesaleCostCents int64      `gorm:"not null" json:"wholesaleCostCents"`
}

var _ Model = (*InventoryPart)(nil)

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *InventoryPart) TableName() string {
	return "inventory_part"
}

func (m *InventoryPart) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewInventoryPart returns a new model instance from a json key/value map
func NewInventoryPart(kv map[string]any) (*InventoryPart, error) {
	var result InventoryPart
	err := mapstructure.Decode(kv, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
