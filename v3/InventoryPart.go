// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	codec "github.com/hashicorp/go-msgpack/v2/codec"
	"time"
)

type InventoryPart struct {
	ID          string     `gorm:"primaryKey;not null;column:id" json:"id"`
	CreatedDate time.Time  `gorm:"column:createdDate;not null;column:createdDate" json:"createdDate"`
	UpdatedDate *time.Time `gorm:"column:updatedDate;column:updatedDate" json:"updatedDate"`
	Meta        *Meta      `gorm:"type:json;embedded;serializer:json;column:meta;not null;column:meta" json:"meta,omitempty"` // the metadata about the most recent change to the row
	Metadata    any        `gorm:"type:json;serializer:json;column:metadata" json:"metadata,omitempty"`                       // metadata reserved for customers to control
	CompanyID   string     `gorm:"not null;column:companyId" json:"companyId"`
	LocationID  string     `gorm:"not null;column:locationId" json:"locationId"`

	BinLocation        *string    `gorm:"column:binLocation" json:"binLocation"`
	CategoryID         *string    `gorm:"column:categoryId" json:"categoryId"`
	Deleted            bool       `gorm:"not null;column:deleted" json:"deleted"`    // if the record has been deleted
	DeletedDate        *time.Time `gorm:"column:deletedDate" json:"deletedDate"`     // the date that the record was deleted or null if not deleted
	DeletedReason      *string    `gorm:"column:deletedReason" json:"deletedReason"` // the reason that the record was deleted
	DeletedUserID      *string    `gorm:"column:deletedUserId" json:"deletedUserId"` // the user that deleted the record or null if not deleted
	Name               string     `gorm:"not null;column:name" json:"name"`
	Note               *string    `gorm:"column:note" json:"note"`
	Number             string     `gorm:"not null;column:number" json:"number"`
	Quantity           int64      `gorm:"not null;column:quantity" json:"quantity"`
	RetailCostCents    int64      `gorm:"not null;column:retailCostCents" json:"retailCostCents"`
	Taxable            bool       `gorm:"not null;column:taxable" json:"taxable"`
	UserID             *string    `gorm:"column:userId" json:"userId"`
	VendorID           *string    `gorm:"column:vendorId" json:"vendorId"`
	WholesaleCostCents int64      `gorm:"not null;column:wholesaleCostCents" json:"wholesaleCostCents"`
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

// NewInventoryPart returns a new model instance from an encoded buffer
func NewInventoryPart(buf []byte, enctype EncodingType) (*InventoryPart, error) {
	var result InventoryPart
	var handle codec.Handle
	if enctype == JSONEncoding {
		handle = &jsonHandle
	} else {
		handle = &msgpackHandle
	}
	dec := codec.NewDecoderBytes(buf, handle)
	err := dec.Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
