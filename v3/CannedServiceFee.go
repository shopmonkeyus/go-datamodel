// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	mapstructure "github.com/mitchellh/mapstructure"
	"time"
)

type CannedServiceFeeFeeTypeEnum string

const (
	CannedServiceFeeFeeTypePercentLineItem CannedServiceFeeFeeTypeEnum = "PercentLineItem"
	CannedServiceFeeFeeTypePercentService                              = "PercentService"
	CannedServiceFeeFeeTypeFixedCents                                  = "FixedCents"
)

type CannedServiceFeeLineItemEntityEnum string

const (
	CannedServiceFeeLineItemEntityLabor       CannedServiceFeeLineItemEntityEnum = "Labor"
	CannedServiceFeeLineItemEntityPart                                           = "Part"
	CannedServiceFeeLineItemEntitySubcontract                                    = "Subcontract"
	CannedServiceFeeLineItemEntityTire                                           = "Tire"
)

type CannedServiceFee struct {
	ID          string     `gorm:"primaryKey;not null;column:id" json:"id"`
	CreatedDate time.Time  `gorm:"column:createdDate;not null;column:createdDate" json:"createdDate"`
	UpdatedDate *time.Time `gorm:"column:updatedDate;column:updatedDate" json:"updatedDate"`
	Meta        *Meta      `gorm:"type:json;embedded;serializer:json;column:meta;not null;column:meta" json:"meta,omitempty"` // the metadata about the most recent change to the row
	Metadata    any        `gorm:"type:json;serializer:json;column:metadata" json:"metadata,omitempty"`                       // metadata reserved for customers to control
	CompanyID   string     `gorm:"not null;column:companyId" json:"companyId"`
	LocationID  string     `gorm:"not null;column:locationId" json:"locationId"`

	AmountCents     int64                               `gorm:"not null;column:amountCents" json:"amountCents"`
	CannedServiceID string                              `gorm:"not null;column:cannedServiceId" json:"cannedServiceId"`
	CategoryID      *string                             `gorm:"column:categoryId" json:"categoryId"`
	FeeType         CannedServiceFeeFeeTypeEnum         `gorm:"not null;column:feeType" json:"feeType"`
	LaborID         *string                             `gorm:"column:laborId" json:"laborId"`
	LineItemEntity  *CannedServiceFeeLineItemEntityEnum `gorm:"column:lineItemEntity" json:"lineItemEntity"`
	Name            string                              `gorm:"not null;column:name" json:"name"`
	Ordinal         float64                             `gorm:"not null;column:ordinal" json:"ordinal"`
	PartID          *string                             `gorm:"column:partId" json:"partId"`
	Percent         float64                             `gorm:"not null;column:percent" json:"percent"`
	SourceItemID    *string                             `gorm:"column:sourceItemId" json:"sourceItemId"`
	SubcontractID   *string                             `gorm:"column:subcontractId" json:"subcontractId"`
	Subtotal        *int64                              `gorm:"column:subtotal" json:"subtotal"`
	TireID          *string                             `gorm:"column:tireId" json:"tireId"`
}

var _ Model = (*CannedServiceFee)(nil)

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *CannedServiceFee) TableName() string {
	return "canned_service_fee"
}

func (m *CannedServiceFee) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewCannedServiceFee returns a new model instance from a json key/value map
func NewCannedServiceFee(kv map[string]any) (*CannedServiceFee, error) {
	var result CannedServiceFee
	err := mapstructure.Decode(kv, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
