// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	datatypes "github.com/shopmonkeyus/go-datamodel/datatypes"
)

// InspectionItem schema
type InspectionItemAuthorizationStatusEnum string

const (
	InspectionItemAuthorizationStatusUndecided  InspectionItemAuthorizationStatusEnum = "Undecided"
	InspectionItemAuthorizationStatusAuthorized InspectionItemAuthorizationStatusEnum = "Authorized"
	InspectionItemAuthorizationStatusDeclined   InspectionItemAuthorizationStatusEnum = "Declined"
)

type InspectionItemStatusEnum string

const (
	InspectionItemStatusGreen         InspectionItemStatusEnum = "Green"
	InspectionItemStatusYellow        InspectionItemStatusEnum = "Yellow"
	InspectionItemStatusRed           InspectionItemStatusEnum = "Red"
	InspectionItemStatusNotApplicable InspectionItemStatusEnum = "NotApplicable"
)

type InspectionItem struct {
	ID          string              `bson:"_id" gorm:"primaryKey;not null;column:id" json:"id"`
	CreatedDate datatypes.DateTime  `gorm:"column:createdDate;not null;column:createdDate" json:"createdDate"`
	UpdatedDate *datatypes.DateTime `gorm:"column:updatedDate;column:updatedDate" json:"updatedDate"`
	Meta        *datatypes.JSON     `gorm:"column:meta;not null;column:meta" json:"meta,omitempty"`    // the metadata about the most recent change to the row
	Metadata    *datatypes.JSON     `gorm:"column:metadata;column:metadata" json:"metadata,omitempty"` // metadata reserved for customers to control
	CompanyID   string              `gorm:"not null;column:companyId" json:"companyId"`
	LocationID  string              `gorm:"not null;column:locationId" json:"locationId"`

	Approved            bool                                  `gorm:"not null;column:approved" json:"approved"`
	AuthorizationStatus InspectionItemAuthorizationStatusEnum `gorm:"not null;column:authorizationStatus" json:"authorizationStatus"`
	InspectionDate      *datatypes.DateTime                   `gorm:"column:inspectionDate" json:"inspectionDate"`
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

// PrimaryKeys returns an array of the primary keys for this model
func (m *InspectionItem) PrimaryKeys() []string {
	return []string{"id"}
}

// String returns a string representation as JSON for this model
func (m *InspectionItem) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewInspectionItem returns a new model instance from an encoded buffer
func NewInspectionItem(buf []byte) (*InspectionItem, error) {
	var result InspectionItem
	err := json.Unmarshal(buf, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// NewInspectionItemFromChangeEvent returns a new model instance from an encoded buffer as change event
func NewInspectionItemFromChangeEvent(buf []byte, gzip bool) (*datatypes.ChangeEvent[InspectionItem], error) {
	var result datatypes.ChangeEvent[InspectionItem]
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
