// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	datatypes "github.com/shopmonkeyus/go-datamodel/datatypes"
)

// LabelLaborConnection schema
type LabelLaborConnection struct {
	LabelID string `gorm:"primaryKey;not null;column:labelId" json:"labelId"`
	LaborID string `gorm:"primaryKey;not null;column:laborId" json:"laborId"`

	CompanyID   string              `gorm:"not null;column:companyId" json:"companyId"`
	CreatedDate datatypes.DateTime  `gorm:"column:createdDate;not null;column:createdDate" json:"createdDate"`
	Meta        *datatypes.JSON     `gorm:"column:meta;not null;column:meta" json:"meta,omitempty"`    // the metadata about the most recent change to the row
	Metadata    *datatypes.JSON     `gorm:"column:metadata;column:metadata" json:"metadata,omitempty"` // metadata reserved for customers to control
	UpdatedDate *datatypes.DateTime `gorm:"column:updatedDate;column:updatedDate" json:"updatedDate"`
}

var _ Model = (*LabelLaborConnection)(nil)

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *LabelLaborConnection) TableName() string {
	return "label_labor_connection"
}

// PrimaryKeys returns an array of the primary keys for this model
func (m *LabelLaborConnection) PrimaryKeys() []string {
	return []string{"labelId", "laborId"}
}

// String returns a string representation as JSON for this model
func (m *LabelLaborConnection) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewLabelLaborConnection returns a new model instance from an encoded buffer
func NewLabelLaborConnection(buf []byte) (*LabelLaborConnection, error) {
	var result LabelLaborConnection
	err := json.Unmarshal(buf, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// NewLabelLaborConnectionFromChangeEvent returns a new model instance from an encoded buffer as change event
func NewLabelLaborConnectionFromChangeEvent(buf []byte, gzip bool) (*datatypes.ChangeEvent[LabelLaborConnection], error) {
	var result datatypes.ChangeEvent[LabelLaborConnection]
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
