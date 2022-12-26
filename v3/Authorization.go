// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	codec "github.com/hashicorp/go-msgpack/v2/codec"
	"time"
)

type AuthorizationMethodEnum string

const (
	AuthorizationMethodInPerson AuthorizationMethodEnum = "InPerson"
	AuthorizationMethodPhone                            = "Phone"
	AuthorizationMethodText                             = "Text"
	AuthorizationMethodEmail                            = "Email"
	AuthorizationMethodInApp                            = "InApp"
)

type Authorization struct {
	ID          string     `gorm:"primaryKey;not null;column:id" json:"id"`
	CreatedDate time.Time  `gorm:"column:createdDate;not null;column:createdDate" json:"createdDate"`
	UpdatedDate *time.Time `gorm:"column:updatedDate;column:updatedDate" json:"updatedDate"`
	Meta        *Meta      `gorm:"type:json;embedded;serializer:json;column:meta;not null;column:meta" json:"meta,omitempty"` // the metadata about the most recent change to the row
	Metadata    any        `gorm:"type:json;serializer:json;column:metadata" json:"metadata,omitempty"`                       // metadata reserved for customers to control
	CompanyID   string     `gorm:"not null;column:companyId" json:"companyId"`
	LocationID  string     `gorm:"not null;column:locationId" json:"locationId"`

	AuthorizedCostCents int64                   `gorm:"not null;column:authorizedCostCents" json:"authorizedCostCents"`
	CustomerID          *string                 `gorm:"column:customerId" json:"customerId"`
	Date                time.Time               `gorm:"not null;column:date" json:"date"`
	Method              AuthorizationMethodEnum `gorm:"not null;column:method" json:"method"`
	Note                string                  `gorm:"not null;column:note" json:"note"`
	OrderID             string                  `gorm:"not null;column:orderId" json:"orderId"`
	ServiceWriterID     *string                 `gorm:"column:serviceWriterId" json:"serviceWriterId"`
}

var _ Model = (*Authorization)(nil)

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *Authorization) TableName() string {
	return "authorization"
}

func (m *Authorization) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewAuthorization returns a new model instance from an encoded buffer
func NewAuthorization(buf []byte, enctype EncodingType) (*Authorization, error) {
	var result Authorization
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
