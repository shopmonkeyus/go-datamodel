// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	datatypes "github.com/shopmonkeyus/go-datamodel/datatypes"
)

type AuthorizationServiceAuthorizationStatusEnum string

const (
	AuthorizationServiceAuthorizationStatusNotAuthorized AuthorizationServiceAuthorizationStatusEnum = "NotAuthorized"
	AuthorizationServiceAuthorizationStatusAuthorized    AuthorizationServiceAuthorizationStatusEnum = "Authorized"
	AuthorizationServiceAuthorizationStatusDeclined      AuthorizationServiceAuthorizationStatusEnum = "Declined"
)

type AuthorizationService struct {
	ID          string              `bson:"_id" gorm:"primaryKey;not null;column:id" json:"id"`
	CreatedDate datatypes.DateTime  `gorm:"column:createdDate;not null;column:createdDate" json:"createdDate"`
	UpdatedDate *datatypes.DateTime `gorm:"column:updatedDate;column:updatedDate" json:"updatedDate"`
	Meta        *datatypes.JSON     `gorm:"column:meta;not null;column:meta" json:"meta,omitempty"`    // the metadata about the most recent change to the row
	Metadata    *datatypes.JSON     `gorm:"column:metadata;column:metadata" json:"metadata,omitempty"` // metadata reserved for customers to control
	CompanyID   string              `gorm:"not null;column:companyId" json:"companyId"`
	LocationID  string              `gorm:"not null;column:locationId" json:"locationId"`

	AuthorizationID     string                                      `gorm:"not null;column:authorizationId" json:"authorizationId"`
	AuthorizationStatus AuthorizationServiceAuthorizationStatusEnum `gorm:"not null;column:authorizationStatus" json:"authorizationStatus"`
	AuthorizedCostCents int64                                       `gorm:"not null;column:authorizedCostCents" json:"authorizedCostCents"`
	Name                string                                      `gorm:"not null;column:name" json:"name"`
	ServiceID           string                                      `gorm:"not null;column:serviceId" json:"serviceId"`
}

var _ Model = (*AuthorizationService)(nil)

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *AuthorizationService) TableName() string {
	return "authorization_service"
}

// PrimaryKeys returns an array of the primary keys for this model
func (m *AuthorizationService) PrimaryKeys() []string {
	return []string{"id"}
}

// String returns a string representation as JSON for this model
func (m *AuthorizationService) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewAuthorizationService returns a new model instance from an encoded buffer
func NewAuthorizationService(buf []byte) (*AuthorizationService, error) {
	var result AuthorizationService
	err := json.Unmarshal(buf, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// NewAuthorizationServiceFromChangeEvent returns a new model instance from an encoded buffer as change event
func NewAuthorizationServiceFromChangeEvent(buf []byte, gzip bool) (*datatypes.ChangeEvent[AuthorizationService], error) {
	var result datatypes.ChangeEvent[AuthorizationService]
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
