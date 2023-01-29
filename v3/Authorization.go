// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	datatypes "github.com/shopmonkeyus/go-datamodel/datatypes"
)

// Authorization schema
type AuthorizationMethodEnum string

const (
	AuthorizationMethodInPerson AuthorizationMethodEnum = "InPerson"
	AuthorizationMethodPhone    AuthorizationMethodEnum = "Phone"
	AuthorizationMethodText     AuthorizationMethodEnum = "Text"
	AuthorizationMethodEmail    AuthorizationMethodEnum = "Email"
	AuthorizationMethodInApp    AuthorizationMethodEnum = "InApp"
)

type Authorization struct {
	ID          string              `bson:"_id" gorm:"primaryKey;not null;column:id" json:"id"`
	CreatedDate datatypes.DateTime  `gorm:"column:createdDate;not null;column:createdDate" json:"createdDate"`
	UpdatedDate *datatypes.DateTime `gorm:"column:updatedDate;column:updatedDate" json:"updatedDate"`
	Meta        *datatypes.JSON     `gorm:"column:meta;not null;column:meta" json:"meta,omitempty"`    // the metadata about the most recent change to the row
	Metadata    *datatypes.JSON     `gorm:"column:metadata;column:metadata" json:"metadata,omitempty"` // metadata reserved for customers to control
	CompanyID   string              `gorm:"not null;column:companyId" json:"companyId"`
	LocationID  string              `gorm:"not null;column:locationId" json:"locationId"`

	AuthorizedCostCents int64                   `gorm:"not null;column:authorizedCostCents" json:"authorizedCostCents"`
	CustomerID          *string                 `gorm:"column:customerId" json:"customerId"`
	Date                datatypes.DateTime      `gorm:"not null;column:date" json:"date"`
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

// PrimaryKeys returns an array of the primary keys for this model
func (m *Authorization) PrimaryKeys() []string {
	return []string{"id"}
}

// String returns a string representation as JSON for this model
func (m *Authorization) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewAuthorization returns a new model instance from an encoded buffer
func NewAuthorization(buf []byte) (*Authorization, error) {
	var result Authorization
	err := json.Unmarshal(buf, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// NewAuthorizationFromChangeEvent returns a new model instance from an encoded buffer as change event
func NewAuthorizationFromChangeEvent(buf []byte, gzip bool) (*datatypes.ChangeEvent[Authorization], error) {
	var result datatypes.ChangeEvent[Authorization]
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
