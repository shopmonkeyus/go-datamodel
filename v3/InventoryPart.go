// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	datatypes "github.com/shopmonkeyus/go-datamodel/datatypes"
)

type InventoryPart struct {
	ID          string              `gorm:"primaryKey;not null;column:id" json:"id"`
	CreatedDate datatypes.DateTime  `gorm:"column:createdDate;not null;column:createdDate" json:"createdDate"`
	UpdatedDate *datatypes.DateTime `gorm:"column:updatedDate;column:updatedDate" json:"updatedDate"`
	Meta        datatypes.Meta      `gorm:"column:meta;not null;column:meta" json:"meta,omitempty"`    // the metadata about the most recent change to the row
	Metadata    *datatypes.JSON     `gorm:"column:metadata;column:metadata" json:"metadata,omitempty"` // metadata reserved for customers to control
	CompanyID   string              `gorm:"not null;column:companyId" json:"companyId"`
	LocationID  string              `gorm:"not null;column:locationId" json:"locationId"`

	BinLocation        *string             `gorm:"column:binLocation" json:"binLocation"`
	CategoryID         *string             `gorm:"column:categoryId" json:"categoryId"`
	Deleted            bool                `gorm:"not null;column:deleted" json:"deleted"`    // if the record has been deleted
	DeletedDate        *datatypes.DateTime `gorm:"column:deletedDate" json:"deletedDate"`     // the date that the record was deleted or null if not deleted
	DeletedReason      *string             `gorm:"column:deletedReason" json:"deletedReason"` // the reason that the record was deleted
	DeletedUserID      *string             `gorm:"column:deletedUserId" json:"deletedUserId"` // the user that deleted the record or null if not deleted
	Name               string              `gorm:"not null;column:name" json:"name"`
	Note               *string             `gorm:"column:note" json:"note"`
	Number             string              `gorm:"not null;column:number" json:"number"`
	Quantity           int64               `gorm:"not null;column:quantity" json:"quantity"`
	RetailCostCents    int64               `gorm:"not null;column:retailCostCents" json:"retailCostCents"`
	Taxable            bool                `gorm:"not null;column:taxable" json:"taxable"`
	UserID             *string             `gorm:"column:userId" json:"userId"`
	VendorID           *string             `gorm:"column:vendorId" json:"vendorId"`
	WholesaleCostCents int64               `gorm:"not null;column:wholesaleCostCents" json:"wholesaleCostCents"`
}

var _ Model = (*InventoryPart)(nil)

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *InventoryPart) TableName() string {
	return "inventory_part"
}

// String returns a string representation as JSON for this model
func (m *InventoryPart) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewInventoryPart returns a new model instance from an encoded buffer
func NewInventoryPart(buf []byte) (*InventoryPart, error) {
	var result InventoryPart
	err := json.Unmarshal(buf, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// NewInventoryPartFromChangeEvent returns a new model instance from an encoded buffer as change event
func NewInventoryPartFromChangeEvent(buf []byte, gzip bool) (*datatypes.ChangeEvent[InventoryPart], error) {
	var result datatypes.ChangeEvent[InventoryPart]
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
