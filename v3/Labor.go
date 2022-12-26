// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	"time"
)

type LaborDiscountValueTypeEnum string

const (
	LaborDiscountValueTypePercent    LaborDiscountValueTypeEnum = "Percent"
	LaborDiscountValueTypeFixedCents                            = "FixedCents"
)

type LaborMultiplierTypeEnum string

const (
	LaborMultiplierTypeHours LaborMultiplierTypeEnum = "Hours"
	LaborMultiplierTypeRate                          = "Rate"
)

type LaborSkillRequiredEnum string

const (
	LaborSkillRequiredGeneral     LaborSkillRequiredEnum = "General"
	LaborSkillRequiredMaintenance                        = "Maintenance"
	LaborSkillRequiredPrecision                          = "Precision"
)

type Labor struct {
	ApplicationId            *int64                     `json:"applicationId"`
	CategoryId               *string                    `json:"categoryId"`
	CompanyId                string                     `gorm:"not null" json:"companyId"`
	CostRateCents            *int64                     `json:"costRateCents"`
	ID                       string                     `gorm:"primaryKey;not null" json:"id"`
	Completed                bool                       `gorm:"not null" json:"completed"`
	CompletedDate            *time.Time                 `json:"completedDate"`
	CostHours                *float64                   `json:"costHours"`
	CostTotalCents           *int64                     `json:"costTotalCents"`
	CreatedDate              time.Time                  `gorm:"column:createdDate;not null" json:"createdDate"`
	DiscountCents            int64                      `gorm:"not null" json:"discountCents"`
	DiscountPercent          float64                    `gorm:"not null" json:"discountPercent"`
	DiscountValueType        LaborDiscountValueTypeEnum `gorm:"not null" json:"discountValueType"`
	Hours                    float64                    `gorm:"not null" json:"hours"`
	LaborMatrixDate          *time.Time                 `json:"laborMatrixDate"` // datetime when laborMatrixId was set, for determining if matrix has been changed
	LaborMatrixId            *string                    `json:"laborMatrixId"`
	LocationId               string                     `gorm:"not null" json:"locationId"`
	Meta                     *Meta                      `gorm:"embedded;column:meta;not null" json:"meta,omitempty"`         // the metadata about the most recent change to the row
	Metadata                 any                        `gorm:"embedded;column:metadata;not null" json:"metadata,omitempty"` // metadata reserved for customers to control
	Multiplier               float64                    `gorm:"not null" json:"multiplier"`
	MultiplierType           LaborMultiplierTypeEnum    `gorm:"not null" json:"multiplierType"`
	Name                     *string                    `json:"name"`
	Note                     string                     `gorm:"not null" json:"note"`
	OrderId                  string                     `gorm:"not null" json:"orderId"`
	Ordinal                  float64                    `gorm:"not null" json:"ordinal"`
	RateCents                int64                      `gorm:"not null" json:"rateCents"`
	RateId                   *string                    `json:"rateId"`
	ServiceId                string                     `gorm:"not null" json:"serviceId"`
	ShowHours                bool                       `gorm:"not null" json:"showHours"`
	ShowNote                 bool                       `gorm:"not null" json:"showNote"`
	SkillRequired            *LaborSkillRequiredEnum    `json:"skillRequired"`
	SkillRequiredDescription *string                    `json:"skillRequiredDescription"`
	SourceItemId             *string                    `json:"sourceItemId"`
	Taxable                  bool                       `gorm:"not null" json:"taxable"`
	TechnicianId             *string                    `json:"technicianId"`
	UpdatedDate              *time.Time                 `gorm:"column:updatedDate" json:"updatedDate"`
}

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *Labor) TableName() string {
	return "labor"
}

func (m *Labor) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}
