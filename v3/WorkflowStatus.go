// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	mapstructure "github.com/mitchellh/mapstructure"
	"time"
)

type WorkflowStatus struct {
	ArchiveWhenInactive bool       `gorm:"not null" json:"archiveWhenInactive"`
	CompanyId           string     `gorm:"not null" json:"companyId"`
	ID                  string     `gorm:"primaryKey;not null" json:"id"`
	CreatedDate         time.Time  `gorm:"column:createdDate;not null" json:"createdDate"`
	DaysToArchive       int64      `gorm:"not null" json:"daysToArchive"`
	InvoiceWorkflow     bool       `gorm:"not null" json:"invoiceWorkflow"`
	LocationId          string     `gorm:"not null" json:"locationId"`
	Meta                *Meta      `gorm:"type:json;embedded;column:meta;not null" json:"meta,omitempty"` // the metadata about the most recent change to the row
	Metadata            any        `gorm:"type:json" json:"metadata,omitempty"`                           // metadata reserved for customers to control
	Name                *string    `json:"name"`
	Position            float64    `gorm:"not null" json:"position"` // NOTE: this will be changed to an Int after the upgrade
	UpdatedDate         *time.Time `gorm:"column:updatedDate" json:"updatedDate"`
}

var _ Model = (*WorkflowStatus)(nil)

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *WorkflowStatus) TableName() string {
	return "workflow_status"
}

func (m *WorkflowStatus) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewWorkflowStatus returns a new model instance from a json key/value map
func NewWorkflowStatus(kv map[string]any) (*WorkflowStatus, error) {
	var result WorkflowStatus
	err := mapstructure.Decode(kv, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
