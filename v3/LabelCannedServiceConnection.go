// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	datatypes "github.com/shopmonkeyus/go-datamodel/datatypes"
)

type LabelCannedServiceConnection struct {
	LabelID         string `gorm:"primaryKey;not null;column:labelId" json:"labelId"`
	CannedServiceID string `gorm:"primaryKey;not null;column:cannedServiceId" json:"cannedServiceId"`

	CompanyID   string              `gorm:"not null;column:companyId" json:"companyId"`
	CreatedDate *datatypes.DateTime `gorm:"column:createdDate;column:createdDate" json:"createdDate"`
	Meta        *datatypes.JSON     `gorm:"column:meta;not null;column:meta" json:"meta,omitempty"`    // the metadata about the most recent change to the row
	Metadata    *datatypes.JSON     `gorm:"column:metadata;column:metadata" json:"metadata,omitempty"` // metadata reserved for customers to control
	UpdatedDate *datatypes.DateTime `gorm:"column:updatedDate;column:updatedDate" json:"updatedDate"`
}

var _ Model = (*LabelCannedServiceConnection)(nil)

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *LabelCannedServiceConnection) TableName() string {
	return "label_canned_service"
}

// PrimaryKeys returns an array of the primary keys for this model
func (m *LabelCannedServiceConnection) PrimaryKeys() []string {
	return []string{"labelId", "cannedServiceId"}
}

// String returns a string representation as JSON for this model
func (m *LabelCannedServiceConnection) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewLabelCannedServiceConnection returns a new model instance from an encoded buffer
func NewLabelCannedServiceConnection(buf []byte) (*LabelCannedServiceConnection, error) {
	var result LabelCannedServiceConnection
	err := json.Unmarshal(buf, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// NewLabelCannedServiceConnectionFromChangeEvent returns a new model instance from an encoded buffer as change event
func NewLabelCannedServiceConnectionFromChangeEvent(buf []byte, gzip bool) (*datatypes.ChangeEvent[LabelCannedServiceConnection], error) {
	var result datatypes.ChangeEvent[LabelCannedServiceConnection]
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
