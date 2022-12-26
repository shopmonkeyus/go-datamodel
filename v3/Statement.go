// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	mapstructure "github.com/mitchellh/mapstructure"
	"time"
)

type Statement struct {
	CompanyId   string     `gorm:"not null" json:"companyId"`
	CreatedDate time.Time  `gorm:"column:createdDate;not null" json:"createdDate"`
	FromDate    time.Time  `gorm:"not null" json:"fromDate"`
	ID          string     `gorm:"primaryKey;not null" json:"id"`
	Meta        *Meta      `gorm:"type:json;embedded;column:meta;not null" json:"meta,omitempty"` // the metadata about the most recent change to the row
	Metadata    any        `gorm:"type:json" json:"metadata,omitempty"`                           // metadata reserved for customers to control
	Name        string     `gorm:"not null" json:"name"`
	SentDate    *time.Time `json:"sentDate"`
	ToDate      time.Time  `gorm:"not null" json:"toDate"`
	UpdatedDate *time.Time `gorm:"column:updatedDate" json:"updatedDate"`
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
func NewStatement(kv map[string]any) (*Statement, error) {
	var result Statement
	err := mapstructure.Decode(kv, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
