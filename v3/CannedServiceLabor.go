// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	"time"
)

type CannedServiceLaborDiscountValueTypeEnum string

const (
	CannedServiceLaborDiscountValueTypePercent    CannedServiceLaborDiscountValueTypeEnum = "Percent"
	CannedServiceLaborDiscountValueTypeFixedCents                                         = "FixedCents"
)

type CannedServiceLaborMultiplierTypeEnum string

const (
	CannedServiceLaborMultiplierTypeHours CannedServiceLaborMultiplierTypeEnum = "Hours"
	CannedServiceLaborMultiplierTypeRate                                       = "Rate"
)

type CannedServiceLaborSkillRequiredEnum string

const (
	CannedServiceLaborSkillRequiredGeneral     CannedServiceLaborSkillRequiredEnum = "General"
	CannedServiceLaborSkillRequiredMaintenance                                     = "Maintenance"
	CannedServiceLaborSkillRequiredPrecision                                       = "Precision"
)

type CannedServiceLabor struct {
	ApplicationId            *int64                                  `json:"applicationId"`
	CannedServiceId          string                                  `gorm:"not null" json:"cannedServiceId"`
	CategoryId               *string                                 `json:"categoryId"`
	CompanyId                string                                  `gorm:"not null" json:"companyId"`
	CostHours                *float64                                `json:"costHours"`
	CostRateCents            *int64                                  `json:"costRateCents"`
	CostTotalCents           *int64                                  `json:"costTotalCents"`
	CreatedDate              time.Time                               `gorm:"column:createdDate;not null" json:"createdDate"`
	DiscountCents            int64                                   `gorm:"not null" json:"discountCents"`
	DiscountPercent          float64                                 `gorm:"not null" json:"discountPercent"`
	DiscountValueType        CannedServiceLaborDiscountValueTypeEnum `gorm:"not null" json:"discountValueType"`
	Hours                    float64                                 `gorm:"not null" json:"hours"`
	ID                       string                                  `gorm:"primaryKey;not null" json:"id"`
	LaborMatrixDate          *time.Time                              `json:"laborMatrixDate"` // datetime when laborMatrixId was set, for determining if matrix has been changed
	LaborMatrixId            *string                                 `json:"laborMatrixId"`
	LocationId               string                                  `gorm:"not null" json:"locationId"`
	Meta                     *Meta                                   `gorm:"embedded;column:meta;not null" json:"meta,omitempty"`         // the metadata about the most recent change to the row
	Metadata                 any                                     `gorm:"embedded;column:metadata;not null" json:"metadata,omitempty"` // metadata reserved for customers to control
	Multiplier               float64                                 `gorm:"not null" json:"multiplier"`
	MultiplierType           CannedServiceLaborMultiplierTypeEnum    `gorm:"not null" json:"multiplierType"`
	Name                     *string                                 `json:"name"`
	Note                     string                                  `gorm:"not null" json:"note"`
	Ordinal                  float64                                 `gorm:"not null" json:"ordinal"`
	RateCents                int64                                   `gorm:"not null" json:"rateCents"`
	RateId                   *string                                 `json:"rateId"`
	ShowHours                bool                                    `gorm:"not null" json:"showHours"`
	ShowNote                 bool                                    `gorm:"not null" json:"showNote"`
	SkillRequired            *CannedServiceLaborSkillRequiredEnum    `json:"skillRequired"`
	SkillRequiredDescription *string                                 `json:"skillRequiredDescription"`
	SourceItemId             *string                                 `json:"sourceItemId"`
	Taxable                  bool                                    `gorm:"not null" json:"taxable"`
	UpdatedDate              *time.Time                              `gorm:"column:updatedDate" json:"updatedDate"`
}

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *CannedServiceLabor) TableName() string {
	return "canned_service_labor"
}

func (m *CannedServiceLabor) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}
