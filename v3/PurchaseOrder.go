// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	codec "github.com/hashicorp/go-msgpack/v2/codec"
	datatypes "github.com/shopmonkeyus/go-datamodel/datatypes"
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
	ID          string          `gorm:"primaryKey;not null;column:id" json:"id"`
	CreatedDate time.Time       `gorm:"column:createdDate;not null;column:createdDate" json:"createdDate"`
	UpdatedDate *time.Time      `gorm:"column:updatedDate;column:updatedDate" json:"updatedDate"`
	Meta        datatypes.Meta  `gorm:"column:meta;not null;column:meta" json:"meta,omitempty"`    // the metadata about the most recent change to the row
	Metadata    *datatypes.JSON `gorm:"column:metadata;column:metadata" json:"metadata,omitempty"` // metadata reserved for customers to control
	CompanyID   string          `gorm:"not null;column:companyId" json:"companyId"`
	LocationID  string          `gorm:"not null;column:locationId" json:"locationId"`

	FullfiledDate   *time.Time                 `gorm:"column:fullfiledDate" json:"fullfiledDate"`
	InvoiceNumber   *string                    `gorm:"column:invoiceNumber" json:"invoiceNumber"`
	Note            *string                    `gorm:"column:note" json:"note"`
	Number          string                     `gorm:"not null;column:number" json:"number"`
	OrderID         *string                    `gorm:"column:orderId" json:"orderId"`
	OrderedDate     *time.Time                 `gorm:"column:orderedDate" json:"orderedDate"`
	Provider        *PurchaseOrderProviderEnum `gorm:"column:provider" json:"provider"`
	ProviderData    datatypes.JSON             `gorm:"column:providerData" json:"providerData"`
	Session         datatypes.JSON             `gorm:"column:session" json:"session"`
	Status          PurchaseOrderStatusEnum    `gorm:"not null;column:status" json:"status"`
	TotalPriceCents int64                      `gorm:"not null;column:totalPriceCents" json:"totalPriceCents"`
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

// NewPurchaseOrder returns a new model instance from an encoded buffer
func NewPurchaseOrder(buf []byte, enctype EncodingType) (*PurchaseOrder, error) {
	var result PurchaseOrder
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
