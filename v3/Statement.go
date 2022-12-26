// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	"time"
)

type Statement struct {
	ID          string     `gorm:"primaryKey;not null;column:id" json:"id"`
	CreatedDate time.Time  `gorm:"column:createdDate;not null;column:createdDate" json:"createdDate"`
	UpdatedDate *time.Time `gorm:"column:updatedDate;column:updatedDate" json:"updatedDate"`
	Meta        *Meta      `gorm:"type:json;embedded;serializer:json;column:meta;not null;column:meta" json:"meta,omitempty"` // the metadata about the most recent change to the row
	Metadata    any        `gorm:"type:json;serializer:json;column:metadata" json:"metadata,omitempty"`                       // metadata reserved for customers to control
	CompanyID   string     `gorm:"not null;column:companyId" json:"companyId"`

	FromDate time.Time  `gorm:"not null;column:fromDate" json:"fromDate"`
	Name     string     `gorm:"not null;column:name" json:"name"`
	SentDate *time.Time `gorm:"column:sentDate" json:"sentDate"`
	ToDate   time.Time  `gorm:"not null;column:toDate" json:"toDate"`
}

var _ Model = (*Statement)(nil)

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *Statement) TableName() string {
	return "statement"
}

func (m *Statement) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewStatement returns a new model instance from a json key/value map
func NewStatement(buf []byte) (*Statement, error) {
	var result Statement
	err := json.Unmarshal(buf, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
