// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	datatypes "github.com/shopmonkeyus/go-datamodel/datatypes"
)

// InventoryLabor schema
type InventoryLaborDiscountValueTypeEnum string

const (
	InventoryLaborDiscountValueTypePercent    InventoryLaborDiscountValueTypeEnum = "Percent"
	InventoryLaborDiscountValueTypeFixedCents InventoryLaborDiscountValueTypeEnum = "FixedCents"
)

type InventoryLaborMultiplierTypeEnum string

const (
	InventoryLaborMultiplierTypeHours InventoryLaborMultiplierTypeEnum = "Hours"
	InventoryLaborMultiplierTypeRate  InventoryLaborMultiplierTypeEnum = "Rate"
)

type InventoryLaborSkillRequiredEnum string

const (
	InventoryLaborSkillRequiredGeneral     InventoryLaborSkillRequiredEnum = "General"
	InventoryLaborSkillRequiredMaintenance InventoryLaborSkillRequiredEnum = "Maintenance"
	InventoryLaborSkillRequiredPrecision   InventoryLaborSkillRequiredEnum = "Precision"
)

type InventoryLabor struct {
	ID          string              `bson:"_id" gorm:"primaryKey;not null;column:id" json:"id"`
	CreatedDate datatypes.DateTime  `gorm:"column:createdDate;not null;column:createdDate" json:"createdDate"`
	UpdatedDate *datatypes.DateTime `gorm:"column:updatedDate;column:updatedDate" json:"updatedDate"`
	Meta        *datatypes.JSON     `gorm:"column:meta;not null;column:meta" json:"meta,omitempty"` // the metadata about the most recent change to the row
	Metadata    *datatypes.JSON     `gorm:"column:metadata;column:metadata" json:"metadata,omitempty"`
	CompanyID   string              `gorm:"not null;column:companyId" json:"companyId"`
	LocationID  string              `gorm:"not null;column:locationId" json:"locationId"`

	CategoryID               *string                             `gorm:"column:categoryId" json:"categoryId"`
	CostHours                *float64                            `gorm:"column:costHours" json:"costHours"`
	CostRateCents            *int64                              `gorm:"column:costRateCents" json:"costRateCents"`
	Deleted                  bool                                `gorm:"not null;column:deleted" json:"deleted"`    // if the record has been deleted
	DeletedDate              *datatypes.DateTime                 `gorm:"column:deletedDate" json:"deletedDate"`     // the date that the record was deleted or null if not deleted
	DeletedReason            *string                             `gorm:"column:deletedReason" json:"deletedReason"` // the reason that the record was deleted
	DeletedUserID            *string                             `gorm:"column:deletedUserId" json:"deletedUserId"` // the user that deleted the record or null if not deleted
	DiscountCents            int64                               `gorm:"not null;column:discountCents" json:"discountCents"`
	DiscountPercent          float64                             `gorm:"not null;column:discountPercent" json:"discountPercent"`
	DiscountValueType        InventoryLaborDiscountValueTypeEnum `gorm:"not null;column:discountValueType" json:"discountValueType"`
	Hours                    float64                             `gorm:"not null;column:hours" json:"hours"`
	LaborMatrixDate          *datatypes.DateTime                 `gorm:"column:laborMatrixDate" json:"laborMatrixDate"` // datetime when laborMatrixId was set, for determining if matrix has been changed
	LaborMatrixID            *string                             `gorm:"column:laborMatrixId" json:"laborMatrixId"`
	Multiplier               float64                             `gorm:"not null;column:multiplier" json:"multiplier"`
	MultiplierType           InventoryLaborMultiplierTypeEnum    `gorm:"not null;column:multiplierType" json:"multiplierType"`
	Name                     string                              `gorm:"not null;column:name" json:"name"`
	Note                     string                              `gorm:"not null;column:note" json:"note"`
	RateCents                int64                               `gorm:"not null;column:rateCents" json:"rateCents"`
	RateID                   *string                             `gorm:"column:rateId" json:"rateId"`
	ShowHours                bool                                `gorm:"not null;column:showHours" json:"showHours"`
	ShowNote                 bool                                `gorm:"not null;column:showNote" json:"showNote"`
	SkillRequired            *InventoryLaborSkillRequiredEnum    `gorm:"column:skillRequired" json:"skillRequired"`
	SkillRequiredDescription *string                             `gorm:"column:skillRequiredDescription" json:"skillRequiredDescription"`
	Taxable                  bool                                `gorm:"not null;column:taxable" json:"taxable"`
	TotalCostCents           *int64                              `gorm:"column:totalCostCents" json:"totalCostCents"`
	UserID                   *string                             `gorm:"column:userId" json:"userId"`
}

var _ Model = (*InventoryLabor)(nil)

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *InventoryLabor) TableName() string {
	return "inventory_labor"
}

// PrimaryKeys returns an array of the primary keys for this model
func (m *InventoryLabor) PrimaryKeys() []string {
	return []string{"id"}
}

// String returns a string representation as JSON for this model
func (m *InventoryLabor) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewInventoryLabor returns a new model instance from an encoded buffer
func NewInventoryLabor(buf []byte) (*InventoryLabor, error) {
	var result InventoryLabor
	err := json.Unmarshal(buf, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// NewInventoryLaborFromChangeEvent returns a new model instance from an encoded buffer as change event
func NewInventoryLaborFromChangeEvent(buf []byte, gzip bool) (*datatypes.ChangeEvent[InventoryLabor], error) {
	var result datatypes.ChangeEvent[InventoryLabor]
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
