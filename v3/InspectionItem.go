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
	ID          string     `gorm:"primaryKey;not null;column:id" json:"id"`
	CreatedDate time.Time  `gorm:"column:createdDate;not null;column:createdDate" json:"createdDate"`
	UpdatedDate *time.Time `gorm:"column:updatedDate;column:updatedDate" json:"updatedDate"`
	Meta        *Meta      `gorm:"type:json;embedded;serializer:json;column:meta;not null;column:meta" json:"meta,omitempty"` // the metadata about the most recent change to the row
	Metadata    any        `gorm:"type:json;serializer:json;column:metadata" json:"metadata,omitempty"`                       // metadata reserved for customers to control
	CompanyID   string     `gorm:"not null;column:companyId" json:"companyId"`
	LocationID  string     `gorm:"not null;column:locationId" json:"locationId"`

	Approved            bool                                  `gorm:"not null;column:approved" json:"approved"`
	AuthorizationStatus InspectionItemAuthorizationStatusEnum `gorm:"not null;column:authorizationStatus" json:"authorizationStatus"`
	InspectionDate      *time.Time                            `gorm:"column:inspectionDate" json:"inspectionDate"`
	InspectionID        string                                `gorm:"not null;column:inspectionId" json:"inspectionId"`
	InspectorUserID     *string                               `gorm:"column:inspectorUserId" json:"inspectorUserId"`
	Message             string                                `gorm:"not null;column:message" json:"message"`
	Name                string                                `gorm:"not null;column:name" json:"name"`
	Ordinal             float64                               `gorm:"not null;column:ordinal" json:"ordinal"`
	Status              *InspectionItemStatusEnum             `gorm:"column:status" json:"status"`
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
