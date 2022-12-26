// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	"time"
)

type LabelCannedServiceConnection struct {
	CannedServiceId string     `gorm:"not null" json:"cannedServiceId"`
	CompanyId       string     `gorm:"not null" json:"companyId"`
	CreatedDate     *time.Time `gorm:"column:createdDate" json:"createdDate"`
	LabelId         string     `gorm:"not null" json:"labelId"`
	Meta            *Meta      `gorm:"embedded;column:meta;not null" json:"meta,omitempty"`         // the metadata about the most recent change to the row
	Metadata        any        `gorm:"embedded;column:metadata;not null" json:"metadata,omitempty"` // metadata reserved for customers to control
	UpdatedDate     *time.Time `gorm:"column:updatedDate" json:"updatedDate"`
}

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *LabelCannedServiceConnection) TableName() string {
	return "label_canned_service"
}

func (m *LabelCannedServiceConnection) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}
