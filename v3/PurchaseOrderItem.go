// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
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
	ID                  string                       `gorm:"primaryKey;not null" json:"id"`
	CompanyId           string                       `gorm:"not null" json:"companyId"`
	CoreChargesCents    int64                        `gorm:"not null" json:"coreChargesCents"`
	CostCents           int64                        `gorm:"not null" json:"costCents"`
	CreatedDate         time.Time                    `gorm:"column:createdDate;not null" json:"createdDate"`
	ExciseTaxCents      int64                        `gorm:"not null" json:"exciseTaxCents"`
	ImageUrl            *string                      `json:"imageUrl"`
	InventoryItemId     *string                      `json:"inventoryItemId"`
	LineItemId          *string                      `json:"lineItemId"`
	LocationId          string                       `gorm:"not null" json:"locationId"`
	Meta                *Meta                        `gorm:"embedded;column:meta;not null" json:"meta,omitempty"`         // the metadata about the most recent change to the row
	Metadata            any                          `gorm:"embedded;column:metadata;not null" json:"metadata,omitempty"` // metadata reserved for customers to control
	Name                string                       `gorm:"not null" json:"name"`
	Note                *string                      `json:"note"`
	Number              *string                      `json:"number"`
	ProviderData        any                          `gorm:"not null" json:"providerData"`
	PurchaseOrderId     string                       `gorm:"not null" json:"purchaseOrderId"`
	Quantity            int64                        `gorm:"not null" json:"quantity"`
	ShippingChargeCents int64                        `gorm:"not null" json:"shippingChargeCents"`
	Status              *PurchaseOrderItemStatusEnum `json:"status"`
	Type                PurchaseOrderItemTypeEnum    `gorm:"not null" json:"type"`
	UpdatedDate         *time.Time                   `gorm:"column:updatedDate" json:"updatedDate"`
	VendorId            string                       `gorm:"not null" json:"vendorId"`
}

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *PurchaseOrderItem) TableName() string {
	return "purchase_order_item"
}

func (m *PurchaseOrderItem) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}
