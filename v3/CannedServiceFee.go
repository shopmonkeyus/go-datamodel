// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	datatypes "github.com/shopmonkeyus/go-datamodel/datatypes"
)

// CannedServiceFee schema
type CannedServiceFeeFeeTypeEnum string

const (
	CannedServiceFeeFeeTypePercentLineItem CannedServiceFeeFeeTypeEnum = "PercentLineItem"
	CannedServiceFeeFeeTypePercentService  CannedServiceFeeFeeTypeEnum = "PercentService"
	CannedServiceFeeFeeTypeFixedCents      CannedServiceFeeFeeTypeEnum = "FixedCents"
)

type CannedServiceFeeLineItemEntityEnum string

const (
	CannedServiceFeeLineItemEntityLabor       CannedServiceFeeLineItemEntityEnum = "Labor"
	CannedServiceFeeLineItemEntityPart        CannedServiceFeeLineItemEntityEnum = "Part"
	CannedServiceFeeLineItemEntitySubcontract CannedServiceFeeLineItemEntityEnum = "Subcontract"
	CannedServiceFeeLineItemEntityTire        CannedServiceFeeLineItemEntityEnum = "Tire"
)

type CannedServiceFee struct {
	ID          string              `bson:"_id" gorm:"primaryKey;not null;column:id" json:"id"`
	CreatedDate datatypes.DateTime  `gorm:"column:createdDate;not null;column:createdDate" json:"createdDate"`
	UpdatedDate *datatypes.DateTime `gorm:"column:updatedDate;column:updatedDate" json:"updatedDate"`
	Meta        *datatypes.JSON     `gorm:"column:meta;not null;column:meta" json:"meta,omitempty"` // the metadata about the most recent change to the row
	Metadata    *datatypes.JSON     `gorm:"column:metadata;column:metadata" json:"metadata,omitempty"`
	CompanyID   string              `gorm:"not null;column:companyId" json:"companyId"`
	LocationID  string              `gorm:"not null;column:locationId" json:"locationId"`

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
	SubtotalCents   *int64                              `gorm:"column:subtotalCents" json:"subtotalCents"`
	TireID          *string                             `gorm:"column:tireId" json:"tireId"`
}

var _ Model = (*CannedServiceFee)(nil)

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *CannedServiceFee) TableName() string {
	return "canned_service_fee"
}

// PrimaryKeys returns an array of the primary keys for this model
func (m *CannedServiceFee) PrimaryKeys() []string {
	return []string{"id"}
}

// String returns a string representation as JSON for this model
func (m *CannedServiceFee) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewCannedServiceFee returns a new model instance from an encoded buffer
func NewCannedServiceFee(buf []byte) (*CannedServiceFee, error) {
	var result CannedServiceFee
	err := json.Unmarshal(buf, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// NewCannedServiceFeeFromChangeEvent returns a new model instance from an encoded buffer as change event
func NewCannedServiceFeeFromChangeEvent(buf []byte, gzip bool) (*datatypes.ChangeEvent[CannedServiceFee], error) {
	var result datatypes.ChangeEvent[CannedServiceFee]
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
