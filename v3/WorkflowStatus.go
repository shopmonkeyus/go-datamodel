// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	codec "github.com/hashicorp/go-msgpack/v2/codec"
	datatypes "github.com/shopmonkeyus/go-datamodel/datatypes"
	"time"
)

type WorkflowStatus struct {
	ID          string          `gorm:"primaryKey;not null;column:id" json:"id"`
	CreatedDate time.Time       `gorm:"column:createdDate;not null;column:createdDate" json:"createdDate"`
	UpdatedDate *time.Time      `gorm:"column:updatedDate;column:updatedDate" json:"updatedDate"`
	Meta        datatypes.Meta  `gorm:"column:meta;not null;column:meta" json:"meta,omitempty"`    // the metadata about the most recent change to the row
	Metadata    *datatypes.JSON `gorm:"column:metadata;column:metadata" json:"metadata,omitempty"` // metadata reserved for customers to control
	CompanyID   string          `gorm:"not null;column:companyId" json:"companyId"`
	LocationID  string          `gorm:"not null;column:locationId" json:"locationId"`

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

func (m *WorkflowStatus) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewWorkflowStatus returns a new model instance from an encoded buffer
func NewWorkflowStatus(buf []byte, enctype EncodingType) (*WorkflowStatus, error) {
	var result WorkflowStatus
	var handle codec.Handle
	if enctype == JSONEncoding {
		handle = &jsonHandle
	} else {
		handle = &msgpackHandle
	}
	dec := codec.NewDecoderBytes(buf, handle)
	err := dec.Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
