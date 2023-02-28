// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	datatypes "github.com/shopmonkeyus/go-datamodel/datatypes"
)

// CannedService schema
type CannedServiceDiscountValueTypeEnum string

const (
	CannedServiceDiscountValueTypePercent    CannedServiceDiscountValueTypeEnum = "Percent"
	CannedServiceDiscountValueTypeFixedCents CannedServiceDiscountValueTypeEnum = "FixedCents"
)

type CannedServiceEpaValueTypeEnum string

const (
	CannedServiceEpaValueTypePercent    CannedServiceEpaValueTypeEnum = "Percent"
	CannedServiceEpaValueTypeFixedCents CannedServiceEpaValueTypeEnum = "FixedCents"
)

type CannedServiceGstValueTypeEnum string

const (
	CannedServiceGstValueTypePercent    CannedServiceGstValueTypeEnum = "Percent"
	CannedServiceGstValueTypeFixedCents CannedServiceGstValueTypeEnum = "FixedCents"
)

type CannedServiceHstValueTypeEnum string

const (
	CannedServiceHstValueTypePercent    CannedServiceHstValueTypeEnum = "Percent"
	CannedServiceHstValueTypeFixedCents CannedServiceHstValueTypeEnum = "FixedCents"
)

type CannedServicePstValueTypeEnum string

const (
	CannedServicePstValueTypePercent    CannedServicePstValueTypeEnum = "Percent"
	CannedServicePstValueTypeFixedCents CannedServicePstValueTypeEnum = "FixedCents"
)

type CannedServiceShopSuppliesValueTypeEnum string

const (
	CannedServiceShopSuppliesValueTypePercent    CannedServiceShopSuppliesValueTypeEnum = "Percent"
	CannedServiceShopSuppliesValueTypeFixedCents CannedServiceShopSuppliesValueTypeEnum = "FixedCents"
)

type CannedServiceTaxValueTypeEnum string

const (
	CannedServiceTaxValueTypePercent    CannedServiceTaxValueTypeEnum = "Percent"
	CannedServiceTaxValueTypeFixedCents CannedServiceTaxValueTypeEnum = "FixedCents"
)

type CannedService struct {
	ID          string              `bson:"_id" gorm:"primaryKey;not null;column:id" json:"id"`
	CreatedDate datatypes.DateTime  `gorm:"column:createdDate;not null;column:createdDate" json:"createdDate"`
	UpdatedDate *datatypes.DateTime `gorm:"column:updatedDate;column:updatedDate" json:"updatedDate"`
	Meta        *datatypes.JSON     `gorm:"column:meta;not null;column:meta" json:"meta,omitempty"` // the metadata about the most recent change to the row
	Metadata    *datatypes.JSON     `gorm:"column:metadata;column:metadata" json:"metadata,omitempty"`
	CompanyID   string              `gorm:"not null;column:companyId" json:"companyId"`
	LocationID  string              `gorm:"not null;column:locationId" json:"locationId"`

	CalculatedDiscountCents     int64                                  `gorm:"not null;column:calculatedDiscountCents" json:"calculatedDiscountCents"`
	CalculatedDiscountPercent   float64                                `gorm:"not null;column:calculatedDiscountPercent" json:"calculatedDiscountPercent"`
	CalculatedEpaCents          int64                                  `gorm:"not null;column:calculatedEpaCents" json:"calculatedEpaCents"`
	CalculatedFeeCents          int64                                  `gorm:"not null;column:calculatedFeeCents" json:"calculatedFeeCents"`
	CalculatedLaborCents        int64                                  `gorm:"not null;column:calculatedLaborCents" json:"calculatedLaborCents"`
	CalculatedPartsCents        int64                                  `gorm:"not null;column:calculatedPartsCents" json:"calculatedPartsCents"`
	CalculatedShopSuppliesCents int64                                  `gorm:"not null;column:calculatedShopSuppliesCents" json:"calculatedShopSuppliesCents"`
	CalculatedSubcontractsCents int64                                  `gorm:"not null;column:calculatedSubcontractsCents" json:"calculatedSubcontractsCents"`
	CalculatedTaxCents          int64                                  `gorm:"not null;column:calculatedTaxCents" json:"calculatedTaxCents"`
	CalculatedTiresCents        int64                                  `gorm:"not null;column:calculatedTiresCents" json:"calculatedTiresCents"`
	Deleted                     bool                                   `gorm:"not null;column:deleted" json:"deleted"`    // if the record has been deleted
	DeletedDate                 *datatypes.DateTime                    `gorm:"column:deletedDate" json:"deletedDate"`     // the date that the record was deleted or null if not deleted
	DeletedReason               *string                                `gorm:"column:deletedReason" json:"deletedReason"` // the reason that the record was deleted
	DeletedUserID               *string                                `gorm:"column:deletedUserId" json:"deletedUserId"` // the user that deleted the record or null if not deleted
	DiscountCents               int64                                  `gorm:"not null;column:discountCents" json:"discountCents"`
	DiscountPercent             float64                                `gorm:"not null;column:discountPercent" json:"discountPercent"`
	DiscountValueType           CannedServiceDiscountValueTypeEnum     `gorm:"not null;column:discountValueType" json:"discountValueType"`
	EpaCents                    int64                                  `gorm:"not null;column:epaCents" json:"epaCents"`
	EpaPercent                  float64                                `gorm:"not null;column:epaPercent" json:"epaPercent"`
	EpaValueType                CannedServiceEpaValueTypeEnum          `gorm:"not null;column:epaValueType" json:"epaValueType"`
	GstCents                    int64                                  `gorm:"not null;column:gstCents" json:"gstCents"`
	GstPercent                  float64                                `gorm:"not null;column:gstPercent" json:"gstPercent"`
	GstValueType                CannedServiceGstValueTypeEnum          `gorm:"not null;column:gstValueType" json:"gstValueType"`
	HstCents                    int64                                  `gorm:"not null;column:hstCents" json:"hstCents"`
	HstPercent                  float64                                `gorm:"not null;column:hstPercent" json:"hstPercent"`
	HstValueType                CannedServiceHstValueTypeEnum          `gorm:"not null;column:hstValueType" json:"hstValueType"`
	LumpSum                     bool                                   `gorm:"not null;column:lumpSum" json:"lumpSum"`
	Name                        string                                 `gorm:"not null;column:name" json:"name"`
	Note                        string                                 `gorm:"not null;column:note" json:"note"`
	PstCents                    int64                                  `gorm:"not null;column:pstCents" json:"pstCents"`
	PstPercent                  float64                                `gorm:"not null;column:pstPercent" json:"pstPercent"`
	PstValueType                CannedServicePstValueTypeEnum          `gorm:"not null;column:pstValueType" json:"pstValueType"`
	Recommended                 bool                                   `gorm:"not null;column:recommended" json:"recommended"`
	ShopSuppliesApplied         bool                                   `gorm:"not null;column:shopSuppliesApplied" json:"shopSuppliesApplied"`
	ShopSuppliesCents           int64                                  `gorm:"not null;column:shopSuppliesCents" json:"shopSuppliesCents"`
	ShopSuppliesPercent         float64                                `gorm:"not null;column:shopSuppliesPercent" json:"shopSuppliesPercent"`
	ShopSuppliesValueType       CannedServiceShopSuppliesValueTypeEnum `gorm:"not null;column:shopSuppliesValueType" json:"shopSuppliesValueType"`
	TaxCents                    int64                                  `gorm:"not null;column:taxCents" json:"taxCents"`
	TaxPercent                  float64                                `gorm:"not null;column:taxPercent" json:"taxPercent"`
	TaxValueType                CannedServiceTaxValueTypeEnum          `gorm:"not null;column:taxValueType" json:"taxValueType"`
	TotalCents                  int64                                  `gorm:"not null;column:totalCents" json:"totalCents"`
}

var _ Model = (*CannedService)(nil)

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *CannedService) TableName() string {
	return "canned_service"
}

// PrimaryKeys returns an array of the primary keys for this model
func (m *CannedService) PrimaryKeys() []string {
	return []string{"id"}
}

// String returns a string representation as JSON for this model
func (m *CannedService) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewCannedService returns a new model instance from an encoded buffer
func NewCannedService(buf []byte) (*CannedService, error) {
	var result CannedService
	err := json.Unmarshal(buf, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// NewCannedServiceFromChangeEvent returns a new model instance from an encoded buffer as change event
func NewCannedServiceFromChangeEvent(buf []byte, gzip bool) (*datatypes.ChangeEvent[CannedService], error) {
	var result datatypes.ChangeEvent[CannedService]
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
