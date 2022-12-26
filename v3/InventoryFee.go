// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	"time"
)

type InventoryFeeFeeTypeEnum string

const (
	InventoryFeeFeeTypePercentLineItem InventoryFeeFeeTypeEnum = "PercentLineItem"
	InventoryFeeFeeTypePercentService                          = "PercentService"
	InventoryFeeFeeTypeFixedCents                              = "FixedCents"
)

type InventoryFee struct {
	AmountCents   int64                   `gorm:"not null" json:"amountCents"`
	CategoryId    *string                 `json:"categoryId"`
	CompanyId     string                  `gorm:"not null" json:"companyId"`
	CreatedDate   time.Time               `gorm:"column:createdDate;not null" json:"createdDate"`
	Deleted       bool                    `gorm:"not null" json:"deleted"` // if the record has been deleted
	DeletedDate   *time.Time              `json:"deletedDate"`             // the date that the record was deleted or null if not deleted
	ID            string                  `gorm:"primaryKey;not null" json:"id"`
	DeletedReason *string                 `json:"deletedReason"` // the reason that the record was deleted
	DeletedUserId *string                 `json:"deletedUserId"` // the user that deleted the record or null if not deleted
	FeeType       InventoryFeeFeeTypeEnum `gorm:"not null" json:"feeType"`
	LocationId    string                  `gorm:"not null" json:"locationId"`
	Meta          *Meta                   `gorm:"embedded;column:meta;not null" json:"meta,omitempty"`         // the metadata about the most recent change to the row
	Metadata      any                     `gorm:"embedded;column:metadata;not null" json:"metadata,omitempty"` // metadata reserved for customers to control
	Name          string                  `gorm:"not null" json:"name"`
	Percent       float64                 `gorm:"not null" json:"percent"`
	UpdatedDate   *time.Time              `gorm:"column:updatedDate" json:"updatedDate"`
	UserId        *string                 `json:"userId"`
}

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *InventoryFee) TableName() string {
	return "inventory_fee"
}

func (m *InventoryFee) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}
