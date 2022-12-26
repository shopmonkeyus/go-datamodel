// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	"time"
)

type Statement struct {
	CompanyId   string     `gorm:"not null" json:"companyId"`
	ID          string     `gorm:"primaryKey;not null" json:"id"`
	CreatedDate time.Time  `gorm:"column:createdDate;not null" json:"createdDate"`
	FromDate    time.Time  `gorm:"not null" json:"fromDate"`
	Meta        *Meta      `gorm:"embedded;column:meta;not null" json:"meta,omitempty"`         // the metadata about the most recent change to the row
	Metadata    any        `gorm:"embedded;column:metadata;not null" json:"metadata,omitempty"` // metadata reserved for customers to control
	Name        string     `gorm:"not null" json:"name"`
	SentDate    *time.Time `json:"sentDate"`
	ToDate      time.Time  `gorm:"not null" json:"toDate"`
	UpdatedDate *time.Time `gorm:"column:updatedDate" json:"updatedDate"`
}

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *Statement) TableName() string {
	return "statement"
}

func (m *Statement) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}
