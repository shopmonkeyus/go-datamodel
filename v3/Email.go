// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	mapstructure "github.com/mitchellh/mapstructure"
	"time"
)

type Email struct {
	ID          string     `gorm:"primaryKey;not null;column:id" json:"id"`
	CreatedDate time.Time  `gorm:"column:createdDate;not null;column:createdDate" json:"createdDate"`
	UpdatedDate *time.Time `gorm:"column:updatedDate;column:updatedDate" json:"updatedDate"`
	Meta        *Meta      `gorm:"type:json;embedded;serializer:json;column:meta;not null;column:meta" json:"meta,omitempty"` // the metadata about the most recent change to the row
	Metadata    any        `gorm:"type:json;serializer:json;column:metadata" json:"metadata,omitempty"`                       // metadata reserved for customers to control
	CompanyID   string     `gorm:"not null;column:companyId" json:"companyId"`
	LocationID  string     `gorm:"not null;column:locationId" json:"locationId"`

	CustomerID string `gorm:"not null;column:customerId" json:"customerId"`
	Email      string `gorm:"not null;column:email" json:"email"`
	Primary    bool   `gorm:"not null;column:primary" json:"primary"`
	Subscribed bool   `gorm:"not null;column:subscribed" json:"subscribed"`
}

var _ Model = (*Email)(nil)

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *Email) TableName() string {
	return "email"
}

func (m *Email) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewEmail returns a new model instance from a json key/value map
func NewEmail(kv map[string]any) (*Email, error) {
	var result Email
	err := mapstructure.Decode(kv, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
