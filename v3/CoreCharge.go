// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	mapstructure "github.com/mitchellh/mapstructure"
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
	ID          string     `gorm:"primaryKey;not null;column:id" json:"id"`
	CreatedDate time.Time  `gorm:"column:createdDate;not null;column:createdDate" json:"createdDate"`
	UpdatedDate *time.Time `gorm:"column:updatedDate;column:updatedDate" json:"updatedDate"`
	Meta        *Meta      `gorm:"type:json;embedded;serializer:json;column:meta;not null;column:meta" json:"meta,omitempty"` // the metadata about the most recent change to the row
	Metadata    any        `gorm:"type:json;serializer:json;column:metadata" json:"metadata,omitempty"`                       // metadata reserved for customers to control
	CompanyID   string     `gorm:"not null;column:companyId" json:"companyId"`
	LocationID  string     `gorm:"not null;column:locationId" json:"locationId"`

	AmountCents int64                `gorm:"not null;column:amountCents" json:"amountCents"`
	PartID      string               `gorm:"not null;column:partId" json:"partId"`
	Status      CoreChargeStatusEnum `gorm:"not null;column:status" json:"status"`
}

var _ Model = (*CoreCharge)(nil)

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *CoreCharge) TableName() string {
	return "core_charge"
}

func (m *CoreCharge) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewCoreCharge returns a new model instance from a json key/value map
func NewCoreCharge(kv map[string]any) (*CoreCharge, error) {
	var result CoreCharge
	err := mapstructure.Decode(kv, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
