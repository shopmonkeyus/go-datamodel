// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	"time"
)

type CoreChargeStatusEnum string

const (
	CoreChargeStatusNotReady      CoreChargeStatusEnum = "NotReady"
	CoreChargeStatusReady                              = "Ready"
	CoreChargeStatusPickedUp                           = "PickedUp"
	CoreChargeStatusRefunded                           = "Refunded"
	CoreChargeStatusNotRefundable                      = "NotRefundable"
)

type CoreCharge struct {
	ID          string               `gorm:"primaryKey;not null" json:"id"`
	AmountCents int64                `gorm:"not null" json:"amountCents"`
	CompanyId   string               `gorm:"not null" json:"companyId"`
	CreatedDate time.Time            `gorm:"column:createdDate;not null" json:"createdDate"`
	LocationId  string               `gorm:"not null" json:"locationId"`
	Meta        *Meta                `gorm:"embedded;column:meta;not null" json:"meta,omitempty"`         // the metadata about the most recent change to the row
	Metadata    any                  `gorm:"embedded;column:metadata;not null" json:"metadata,omitempty"` // metadata reserved for customers to control
	PartId      string               `gorm:"not null" json:"partId"`
	Status      CoreChargeStatusEnum `gorm:"not null" json:"status"`
	UpdatedDate *time.Time           `gorm:"column:updatedDate" json:"updatedDate"`
}

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *CoreCharge) TableName() string {
	return "core_charge"
}

func (m *CoreCharge) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}
