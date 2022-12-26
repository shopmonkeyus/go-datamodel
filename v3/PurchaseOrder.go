// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	mapstructure "github.com/mitchellh/mapstructure"
	"time"
)

type PurchaseOrderProviderEnum string

const (
	PurchaseOrderProviderPartsTech PurchaseOrderProviderEnum = "PartsTech"
	PurchaseOrderProviderNexpart                             = "Nexpart"
	PurchaseOrderProviderEpicor                              = "Epicor"
	PurchaseOrderProviderWorldpac                            = "Worldpac"
	PurchaseOrderProviderATD                                 = "ATD"
)

type PurchaseOrderStatusEnum string

const (
	PurchaseOrderStatusDraft     PurchaseOrderStatusEnum = "Draft"
	PurchaseOrderStatusOrdered                           = "Ordered"
	PurchaseOrderStatusReceived                          = "Received"
	PurchaseOrderStatusFullfiled                         = "Fullfiled"
	PurchaseOrderStatusCancelled                         = "Cancelled"
)

type PurchaseOrder struct {
	CompanyId       string                     `gorm:"not null" json:"companyId"`
	CreatedDate     time.Time                  `gorm:"column:createdDate;not null" json:"createdDate"`
	ID              string                     `gorm:"primaryKey;not null" json:"id"`
	FullfiledDate   *time.Time                 `json:"fullfiledDate"`
	InvoiceNumber   *string                    `json:"invoiceNumber"`
	LocationId      string                     `gorm:"not null" json:"locationId"`
	Meta            *Meta                      `gorm:"type:json;embedded;column:meta;not null" json:"meta,omitempty"` // the metadata about the most recent change to the row
	Metadata        any                        `gorm:"type:json" json:"metadata,omitempty"`                           // metadata reserved for customers to control
	Note            *string                    `json:"note"`
	Number          string                     `gorm:"not null" json:"number"`
	OrderId         *string                    `json:"orderId"`
	OrderedDate     *time.Time                 `json:"orderedDate"`
	Provider        *PurchaseOrderProviderEnum `json:"provider"`
	ProviderData    any                        `gorm:"type:json" json:"providerData"`
	Session         any                        `gorm:"type:json" json:"session"`
	Status          PurchaseOrderStatusEnum    `gorm:"not null" json:"status"`
	TotalPriceCents int64                      `gorm:"not null" json:"totalPriceCents"`
	UpdatedDate     *time.Time                 `gorm:"column:updatedDate" json:"updatedDate"`
}

var _ Model = (*PurchaseOrder)(nil)

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *PurchaseOrder) TableName() string {
	return "purchase_order"
}

func (m *PurchaseOrder) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewPurchaseOrder returns a new model instance from a json key/value map
func NewPurchaseOrder(kv map[string]any) (*PurchaseOrder, error) {
	var result PurchaseOrder
	err := mapstructure.Decode(kv, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
