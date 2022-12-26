// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	codec "github.com/hashicorp/go-msgpack/v2/codec"
	"time"
)

type PurchaseOrderItemStatusEnum string

const (
	PurchaseOrderItemStatusAvailable    PurchaseOrderItemStatusEnum = "Available"
	PurchaseOrderItemStatusNotAvailable                             = "NotAvailable"
	PurchaseOrderItemStatusOrdered                                  = "Ordered"
)

type PurchaseOrderItemTypeEnum string

const (
	PurchaseOrderItemTypePart PurchaseOrderItemTypeEnum = "Part"
	PurchaseOrderItemTypeTire                           = "Tire"
)

type PurchaseOrderItem struct {
	ID          string     `gorm:"primaryKey;not null;column:id" json:"id"`
	CreatedDate time.Time  `gorm:"column:createdDate;not null;column:createdDate" json:"createdDate"`
	UpdatedDate *time.Time `gorm:"column:updatedDate;column:updatedDate" json:"updatedDate"`
	Meta        *Meta      `gorm:"type:json;embedded;serializer:json;column:meta;not null;column:meta" json:"meta,omitempty"` // the metadata about the most recent change to the row
	Metadata    any        `gorm:"type:json;serializer:json;column:metadata" json:"metadata,omitempty"`                       // metadata reserved for customers to control
	CompanyID   string     `gorm:"not null;column:companyId" json:"companyId"`
	LocationID  string     `gorm:"not null;column:locationId" json:"locationId"`

	CoreChargesCents    int64                        `gorm:"not null;column:coreChargesCents" json:"coreChargesCents"`
	CostCents           int64                        `gorm:"not null;column:costCents" json:"costCents"`
	ExciseTaxCents      int64                        `gorm:"not null;column:exciseTaxCents" json:"exciseTaxCents"`
	ImageUrl            *string                      `gorm:"column:imageUrl" json:"imageUrl"`
	InventoryItemID     *string                      `gorm:"column:inventoryItemId" json:"inventoryItemId"`
	LineItemID          *string                      `gorm:"column:lineItemId" json:"lineItemId"`
	Name                string                       `gorm:"not null;column:name" json:"name"`
	Note                *string                      `gorm:"column:note" json:"note"`
	Number              *string                      `gorm:"column:number" json:"number"`
	ProviderData        any                          `gorm:"type:json;serializer:json;column:providerData" json:"providerData"`
	PurchaseOrderID     string                       `gorm:"not null;column:purchaseOrderId" json:"purchaseOrderId"`
	Quantity            int64                        `gorm:"not null;column:quantity" json:"quantity"`
	ShippingChargeCents int64                        `gorm:"not null;column:shippingChargeCents" json:"shippingChargeCents"`
	Status              *PurchaseOrderItemStatusEnum `gorm:"column:status" json:"status"`
	Type                PurchaseOrderItemTypeEnum    `gorm:"not null;column:type" json:"type"`
	VendorID            string                       `gorm:"not null;column:vendorId" json:"vendorId"`
}

var _ Model = (*PurchaseOrderItem)(nil)

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *PurchaseOrderItem) TableName() string {
	return "purchase_order_item"
}

func (m *PurchaseOrderItem) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewPurchaseOrderItem returns a new model instance from an encoded buffer
func NewPurchaseOrderItem(buf []byte, enctype EncodingType) (*PurchaseOrderItem, error) {
	var result PurchaseOrderItem
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
