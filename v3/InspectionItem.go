// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	mapstructure "github.com/mitchellh/mapstructure"
	"time"
)

type InspectionItemAuthorizationStatusEnum string

const (
	InspectionItemAuthorizationStatusUndecided  InspectionItemAuthorizationStatusEnum = "Undecided"
	InspectionItemAuthorizationStatusAuthorized                                       = "Authorized"
	InspectionItemAuthorizationStatusDeclined                                         = "Declined"
)

type InspectionItemStatusEnum string

const (
	InspectionItemStatusGreen         InspectionItemStatusEnum = "Green"
	InspectionItemStatusYellow                                 = "Yellow"
	InspectionItemStatusRed                                    = "Red"
	InspectionItemStatusNotApplicable                          = "NotApplicable"
)

type InspectionItem struct {
	Approved            bool                                  `gorm:"not null" json:"approved"`
	AuthorizationStatus InspectionItemAuthorizationStatusEnum `gorm:"not null" json:"authorizationStatus"`
	CompanyId           string                                `gorm:"not null" json:"companyId"`
	CreatedDate         time.Time                             `gorm:"column:createdDate;not null" json:"createdDate"`
	ID                  string                                `gorm:"primaryKey;not null" json:"id"`
	InspectionDate      *time.Time                            `json:"inspectionDate"`
	InspectionId        string                                `gorm:"not null" json:"inspectionId"`
	InspectorUserId     *string                               `json:"inspectorUserId"`
	LocationId          string                                `gorm:"not null" json:"locationId"`
	Message             string                                `gorm:"not null" json:"message"`
	Meta                *Meta                                 `gorm:"type:json;embedded;column:meta;not null" json:"meta,omitempty"` // the metadata about the most recent change to the row
	Metadata            any                                   `gorm:"type:json" json:"metadata,omitempty"`                           // metadata reserved for customers to control
	Name                string                                `gorm:"not null" json:"name"`
	Ordinal             float64                               `gorm:"not null" json:"ordinal"`
	Status              *InspectionItemStatusEnum             `json:"status"`
	UpdatedDate         *time.Time                            `gorm:"column:updatedDate" json:"updatedDate"`
}

var _ Model = (*InspectionItem)(nil)

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *InspectionItem) TableName() string {
	return "inspection_item"
}

func (m *InspectionItem) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewInspectionItem returns a new model instance from a json key/value map
func NewInspectionItem(kv map[string]any) (*InspectionItem, error) {
	var result InspectionItem
	err := mapstructure.Decode(kv, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
