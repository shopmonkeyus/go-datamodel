// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	"time"
)

type Inspection struct {
	ID          string     `gorm:"primaryKey;not null;column:id" json:"id"`
	CreatedDate time.Time  `gorm:"column:createdDate;not null;column:createdDate" json:"createdDate"`
	UpdatedDate *time.Time `gorm:"column:updatedDate;column:updatedDate" json:"updatedDate"`
	Meta        *Meta      `gorm:"type:json;embedded;serializer:json;column:meta;not null;column:meta" json:"meta,omitempty"` // the metadata about the most recent change to the row
	Metadata    any        `gorm:"type:json;serializer:json;column:metadata" json:"metadata,omitempty"`                       // metadata reserved for customers to control
	CompanyID   string     `gorm:"not null;column:companyId" json:"companyId"`
	LocationID  string     `gorm:"not null;column:locationId" json:"locationId"`

	Completed     bool       `gorm:"not null;column:completed" json:"completed"`
	CompletedByID *string    `gorm:"column:completedById" json:"completedById"`
	CompletedDate *time.Time `gorm:"column:completedDate" json:"completedDate"`
	CreatedByID   *string    `gorm:"column:createdById" json:"createdById"`
	Name          string     `gorm:"not null;column:name" json:"name"`
	OrderID       string     `gorm:"not null;column:orderId" json:"orderId"`
	Ordinal       float64    `gorm:"not null;column:ordinal" json:"ordinal"`
	TemplateID    *string    `gorm:"column:templateId" json:"templateId"`
}

var _ Model = (*Inspection)(nil)

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *Inspection) TableName() string {
	return "inspection"
}

func (m *Inspection) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewInspection returns a new model instance from a json key/value map
func NewInspection(buf []byte) (*Inspection, error) {
	var result Inspection
	err := json.Unmarshal(buf, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
