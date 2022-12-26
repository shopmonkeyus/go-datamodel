// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
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
	ID                  string                  `gorm:"primaryKey;not null" json:"id"`
	AuthorizedCostCents int64                   `gorm:"not null" json:"authorizedCostCents"`
	CompanyId           string                  `gorm:"not null" json:"companyId"`
	CreatedDate         time.Time               `gorm:"column:createdDate;not null" json:"createdDate"`
	CustomerId          *string                 `json:"customerId"`
	Date                time.Time               `gorm:"not null" json:"date"`
	LocationId          string                  `gorm:"not null" json:"locationId"`
	Meta                *Meta                   `gorm:"embedded;column:meta;not null" json:"meta,omitempty"`         // the metadata about the most recent change to the row
	Metadata            any                     `gorm:"embedded;column:metadata;not null" json:"metadata,omitempty"` // metadata reserved for customers to control
	Method              AuthorizationMethodEnum `gorm:"not null" json:"method"`
	Note                string                  `gorm:"not null" json:"note"`
	OrderId             string                  `gorm:"not null" json:"orderId"`
	ServiceWriterId     *string                 `json:"serviceWriterId"`
	UpdatedDate         *time.Time              `gorm:"column:updatedDate" json:"updatedDate"`
}

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *Authorization) TableName() string {
	return "authorization"
}

func (m *Authorization) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}
