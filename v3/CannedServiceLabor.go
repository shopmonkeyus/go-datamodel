// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	datatypes "github.com/shopmonkeyus/go-datamodel/datatypes"
)

type CannedServiceLaborDiscountValueTypeEnum string

const (
	CannedServiceLaborDiscountValueTypePercent    CannedServiceLaborDiscountValueTypeEnum = "Percent"
	CannedServiceLaborDiscountValueTypeFixedCents CannedServiceLaborDiscountValueTypeEnum = "FixedCents"
)

type CannedServiceLaborMultiplierTypeEnum string

const (
	CannedServiceLaborMultiplierTypeHours CannedServiceLaborMultiplierTypeEnum = "Hours"
	CannedServiceLaborMultiplierTypeRate  CannedServiceLaborMultiplierTypeEnum = "Rate"
)

type CannedServiceLaborSkillRequiredEnum string

const (
	CannedServiceLaborSkillRequiredGeneral     CannedServiceLaborSkillRequiredEnum = "General"
	CannedServiceLaborSkillRequiredMaintenance CannedServiceLaborSkillRequiredEnum = "Maintenance"
	CannedServiceLaborSkillRequiredPrecision   CannedServiceLaborSkillRequiredEnum = "Precision"
)

type CannedServiceLabor struct {
	ID          string              `bson:"_id" gorm:"primaryKey;not null;column:id" json:"id"`
	CreatedDate datatypes.DateTime  `gorm:"column:createdDate;not null;column:createdDate" json:"createdDate"`
	UpdatedDate *datatypes.DateTime `gorm:"column:updatedDate;column:updatedDate" json:"updatedDate"`
	Meta        *datatypes.JSON     `gorm:"column:meta;not null;column:meta" json:"meta,omitempty"`    // the metadata about the most recent change to the row
	Metadata    *datatypes.JSON     `gorm:"column:metadata;column:metadata" json:"metadata,omitempty"` // metadata reserved for customers to control
	CompanyID   string              `gorm:"not null;column:companyId" json:"companyId"`
	LocationID  string              `gorm:"not null;column:locationId" json:"locationId"`

	ApplicationID            *int64                                  `gorm:"column:applicationId" json:"applicationId"`
	CannedServiceID          string                                  `gorm:"not null;column:cannedServiceId" json:"cannedServiceId"`
	CategoryID               *string                                 `gorm:"column:categoryId" json:"categoryId"`
	CostHours                *float64                                `gorm:"column:costHours" json:"costHours"`
	CostRateCents            *int64                                  `gorm:"column:costRateCents" json:"costRateCents"`
	CostTotalCents           *int64                                  `gorm:"column:costTotalCents" json:"costTotalCents"`
	DiscountCents            int64                                   `gorm:"not null;column:discountCents" json:"discountCents"`
	DiscountPercent          float64                                 `gorm:"not null;column:discountPercent" json:"discountPercent"`
	DiscountValueType        CannedServiceLaborDiscountValueTypeEnum `gorm:"not null;column:discountValueType" json:"discountValueType"`
	Hours                    float64                                 `gorm:"not null;column:hours" json:"hours"`
	LaborMatrixDate          *datatypes.DateTime                     `gorm:"column:laborMatrixDate" json:"laborMatrixDate"` // datetime when laborMatrixId was set, for determining if matrix has been changed
	LaborMatrixID            *string                                 `gorm:"column:laborMatrixId" json:"laborMatrixId"`
	Multiplier               float64                                 `gorm:"not null;column:multiplier" json:"multiplier"`
	MultiplierType           CannedServiceLaborMultiplierTypeEnum    `gorm:"not null;column:multiplierType" json:"multiplierType"`
	Name                     *string                                 `gorm:"column:name" json:"name"`
	Note                     string                                  `gorm:"not null;column:note" json:"note"`
	Ordinal                  float64                                 `gorm:"not null;column:ordinal" json:"ordinal"`
	RateCents                int64                                   `gorm:"not null;column:rateCents" json:"rateCents"`
	RateID                   *string                                 `gorm:"column:rateId" json:"rateId"`
	ShowHours                bool                                    `gorm:"not null;column:showHours" json:"showHours"`
	ShowNote                 bool                                    `gorm:"not null;column:showNote" json:"showNote"`
	SkillRequired            *CannedServiceLaborSkillRequiredEnum    `gorm:"column:skillRequired" json:"skillRequired"`
	SkillRequiredDescription *string                                 `gorm:"column:skillRequiredDescription" json:"skillRequiredDescription"`
	SourceItemID             *string                                 `gorm:"column:sourceItemId" json:"sourceItemId"`
	Taxable                  bool                                    `gorm:"not null;column:taxable" json:"taxable"`
}

var _ Model = (*CannedServiceLabor)(nil)

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *CannedServiceLabor) TableName() string {
	return "canned_service_labor"
}

// PrimaryKeys returns an array of the primary keys for this model
func (m *CannedServiceLabor) PrimaryKeys() []string {
	return []string{"id"}
}

// String returns a string representation as JSON for this model
func (m *CannedServiceLabor) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewCannedServiceLabor returns a new model instance from an encoded buffer
func NewCannedServiceLabor(buf []byte) (*CannedServiceLabor, error) {
	var result CannedServiceLabor
	err := json.Unmarshal(buf, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// NewCannedServiceLaborFromChangeEvent returns a new model instance from an encoded buffer as change event
func NewCannedServiceLaborFromChangeEvent(buf []byte, gzip bool) (*datatypes.ChangeEvent[CannedServiceLabor], error) {
	var result datatypes.ChangeEvent[CannedServiceLabor]
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
