// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	datatypes "github.com/shopmonkeyus/go-datamodel/datatypes"
)

type LabelSubcontractConnection struct {
	CreatedDate datatypes.DateTime  `gorm:"column:createdDate;not null;column:createdDate" json:"createdDate"`
	UpdatedDate *datatypes.DateTime `gorm:"column:updatedDate;column:updatedDate" json:"updatedDate"`
	Meta        datatypes.Meta      `gorm:"column:meta;not null;column:meta" json:"meta,omitempty"`    // the metadata about the most recent change to the row
	Metadata    *datatypes.JSON     `gorm:"column:metadata;column:metadata" json:"metadata,omitempty"` // metadata reserved for customers to control
	CompanyID   string              `gorm:"not null;column:companyId" json:"companyId"`

	LabelID       string `gorm:"not null;column:labelId" json:"labelId"`
	SubcontractID string `gorm:"not null;column:subcontractId" json:"subcontractId"`
}

var _ Model = (*LabelSubcontractConnection)(nil)

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *LabelSubcontractConnection) TableName() string {
	return "label_subcontract_connection"
}

// String returns a string representation as JSON for this model
func (m *LabelSubcontractConnection) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewLabelSubcontractConnection returns a new model instance from an encoded buffer
func NewLabelSubcontractConnection(buf []byte) (*LabelSubcontractConnection, error) {
	var result LabelSubcontractConnection
	err := json.Unmarshal(buf, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// NewLabelSubcontractConnectionFromChangeEvent returns a new model instance from an encoded buffer as change event
func NewLabelSubcontractConnectionFromChangeEvent(buf []byte, gzip bool) (*datatypes.ChangeEvent[LabelSubcontractConnection], error) {
	var result datatypes.ChangeEvent[LabelSubcontractConnection]
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
