// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	datatypes "github.com/shopmonkeyus/go-datamodel/datatypes"
)

type PurchaseOrderProviderEnum string

const (
	PurchaseOrderProviderPartsTech PurchaseOrderProviderEnum = "PartsTech"
	PurchaseOrderProviderNexpart   PurchaseOrderProviderEnum = "Nexpart"
	PurchaseOrderProviderEpicor    PurchaseOrderProviderEnum = "Epicor"
	PurchaseOrderProviderWorldpac  PurchaseOrderProviderEnum = "Worldpac"
	PurchaseOrderProviderATD       PurchaseOrderProviderEnum = "ATD"
)

type PurchaseOrderStatusEnum string

const (
	PurchaseOrderStatusDraft     PurchaseOrderStatusEnum = "Draft"
	PurchaseOrderStatusOrdered   PurchaseOrderStatusEnum = "Ordered"
	PurchaseOrderStatusReceived  PurchaseOrderStatusEnum = "Received"
	PurchaseOrderStatusFullfiled PurchaseOrderStatusEnum = "Fullfiled"
	PurchaseOrderStatusCancelled PurchaseOrderStatusEnum = "Cancelled"
)

type PurchaseOrder struct {
	ID          string              `bson:"_id" gorm:"primaryKey;not null;column:id" json:"id"`
	CreatedDate datatypes.DateTime  `gorm:"column:createdDate;not null;column:createdDate" json:"createdDate"`
	UpdatedDate *datatypes.DateTime `gorm:"column:updatedDate;column:updatedDate" json:"updatedDate"`
	Meta        *datatypes.JSON     `gorm:"column:meta;not null;column:meta" json:"meta,omitempty"`    // the metadata about the most recent change to the row
	Metadata    *datatypes.JSON     `gorm:"column:metadata;column:metadata" json:"metadata,omitempty"` // metadata reserved for customers to control
	CompanyID   string              `gorm:"not null;column:companyId" json:"companyId"`
	LocationID  string              `gorm:"not null;column:locationId" json:"locationId"`

	FullfiledDate   *datatypes.DateTime        `gorm:"column:fullfiledDate" json:"fullfiledDate"`
	InvoiceNumber   *string                    `gorm:"column:invoiceNumber" json:"invoiceNumber"`
	Note            *string                    `gorm:"column:note" json:"note"`
	Number          string                     `gorm:"not null;column:number" json:"number"`
	OrderID         *string                    `gorm:"column:orderId" json:"orderId"`
	OrderedDate     *datatypes.DateTime        `gorm:"column:orderedDate" json:"orderedDate"`
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

// PrimaryKeys returns an array of the primary keys for this model
func (m *PurchaseOrder) PrimaryKeys() []string {
	return []string{"id"}
}

// String returns a string representation as JSON for this model
func (m *PurchaseOrder) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewPurchaseOrder returns a new model instance from an encoded buffer
func NewPurchaseOrder(buf []byte) (*PurchaseOrder, error) {
	var result PurchaseOrder
	err := json.Unmarshal(buf, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// NewPurchaseOrderFromChangeEvent returns a new model instance from an encoded buffer as change event
func NewPurchaseOrderFromChangeEvent(buf []byte, gzip bool) (*datatypes.ChangeEvent[PurchaseOrder], error) {
	var result datatypes.ChangeEvent[PurchaseOrder]
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
