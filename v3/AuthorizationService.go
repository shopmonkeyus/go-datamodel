// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	codec "github.com/hashicorp/go-msgpack/v2/codec"
	"time"
)

type AuthorizationServiceAuthorizationStatusEnum string

const (
	AuthorizationServiceAuthorizationStatusNotAuthorized AuthorizationServiceAuthorizationStatusEnum = "NotAuthorized"
	AuthorizationServiceAuthorizationStatusAuthorized                                                = "Authorized"
	AuthorizationServiceAuthorizationStatusDeclined                                                  = "Declined"
)

type AuthorizationService struct {
	ID          string     `gorm:"primaryKey;not null;column:id" json:"id"`
	CreatedDate time.Time  `gorm:"column:createdDate;not null;column:createdDate" json:"createdDate"`
	UpdatedDate *time.Time `gorm:"column:updatedDate;column:updatedDate" json:"updatedDate"`
	Meta        *Meta      `gorm:"type:json;embedded;serializer:json;column:meta;not null;column:meta" json:"meta,omitempty"` // the metadata about the most recent change to the row
	Metadata    any        `gorm:"type:json;serializer:json;column:metadata" json:"metadata,omitempty"`                       // metadata reserved for customers to control
	CompanyID   string     `gorm:"not null;column:companyId" json:"companyId"`
	LocationID  string     `gorm:"not null;column:locationId" json:"locationId"`

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

func (m *AuthorizationService) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewAuthorizationService returns a new model instance from an encoded buffer
func NewAuthorizationService(buf []byte, enctype EncodingType) (*AuthorizationService, error) {
	var result AuthorizationService
	var handle codec.Handle
	if enctype == JSONEncoding {
		handle = &jsonHandle
	} else {
		handle = &msgpackHandle
	}
	dec := codec.NewDecoderBytes(buf, handle)
	err := dec.Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
