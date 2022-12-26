// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	mapstructure "github.com/mitchellh/mapstructure"
	"time"
)

type AuthorizationServiceAuthorizationStatusEnum string

const (
	AuthorizationServiceAuthorizationStatusNotAuthorized AuthorizationServiceAuthorizationStatusEnum = "NotAuthorized"
	AuthorizationServiceAuthorizationStatusAuthorized                                                = "Authorized"
	AuthorizationServiceAuthorizationStatusDeclined                                                  = "Declined"
)

type AuthorizationService struct {
	ID                  string                                      `gorm:"primaryKey;not null" json:"id"`
	AuthorizationId     string                                      `gorm:"not null" json:"authorizationId"`
	AuthorizationStatus AuthorizationServiceAuthorizationStatusEnum `gorm:"not null" json:"authorizationStatus"`
	AuthorizedCostCents int64                                       `gorm:"not null" json:"authorizedCostCents"`
	CompanyId           string                                      `gorm:"not null" json:"companyId"`
	CreatedDate         time.Time                                   `gorm:"column:createdDate;not null" json:"createdDate"`
	LocationId          string                                      `gorm:"not null" json:"locationId"`
	Meta                *Meta                                       `gorm:"type:json;embedded;column:meta;not null" json:"meta,omitempty"` // the metadata about the most recent change to the row
	Metadata            any                                         `gorm:"type:json" json:"metadata,omitempty"`                           // metadata reserved for customers to control
	Name                string                                      `gorm:"not null" json:"name"`
	ServiceId           string                                      `gorm:"not null" json:"serviceId"`
	UpdatedDate         *time.Time                                  `gorm:"column:updatedDate" json:"updatedDate"`
}

var _ Model = (*AuthorizationService)(nil)

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *AuthorizationService) TableName() string {
	return "authorization_service"
}

func (m *AuthorizationService) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewAuthorizationService returns a new model instance from a json key/value map
func NewAuthorizationService(kv map[string]any) (*AuthorizationService, error) {
	var result AuthorizationService
	err := mapstructure.Decode(kv, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
