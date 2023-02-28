// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	datatypes "github.com/shopmonkeyus/go-datamodel/datatypes"
)

// WorkflowStatus schema
type WorkflowStatus struct {
	ID          string              `bson:"_id" gorm:"primaryKey;not null;column:id" json:"id"`
	CreatedDate datatypes.DateTime  `gorm:"column:createdDate;not null;column:createdDate" json:"createdDate"`
	UpdatedDate *datatypes.DateTime `gorm:"column:updatedDate;column:updatedDate" json:"updatedDate"`
	Meta        *datatypes.JSON     `gorm:"column:meta;not null;column:meta" json:"meta,omitempty"` // the metadata about the most recent change to the row
	Metadata    *datatypes.JSON     `gorm:"column:metadata;column:metadata" json:"metadata,omitempty"`
	CompanyID   string              `gorm:"not null;column:companyId" json:"companyId"`
	LocationID  string              `gorm:"not null;column:locationId" json:"locationId"`

	ArchiveWhenInactive bool    `gorm:"not null;column:archiveWhenInactive" json:"archiveWhenInactive"`
	DaysToArchive       int64   `gorm:"not null;column:daysToArchive" json:"daysToArchive"`
	InvoiceWorkflow     bool    `gorm:"not null;column:invoiceWorkflow" json:"invoiceWorkflow"`
	Name                *string `gorm:"column:name" json:"name"`
	Position            float64 `gorm:"not null;column:position" json:"position"` // NOTE: this will be changed to an Int after the upgrade
}

var _ Model = (*WorkflowStatus)(nil)

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *WorkflowStatus) TableName() string {
	return "workflow_status"
}

// PrimaryKeys returns an array of the primary keys for this model
func (m *WorkflowStatus) PrimaryKeys() []string {
	return []string{"id"}
}

// String returns a string representation as JSON for this model
func (m *WorkflowStatus) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewWorkflowStatus returns a new model instance from an encoded buffer
func NewWorkflowStatus(buf []byte) (*WorkflowStatus, error) {
	var result WorkflowStatus
	err := json.Unmarshal(buf, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// NewWorkflowStatusFromChangeEvent returns a new model instance from an encoded buffer as change event
func NewWorkflowStatusFromChangeEvent(buf []byte, gzip bool) (*datatypes.ChangeEvent[WorkflowStatus], error) {
	var result datatypes.ChangeEvent[WorkflowStatus]
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
