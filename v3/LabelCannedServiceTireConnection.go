// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	datatypes "github.com/shopmonkeyus/go-datamodel/datatypes"
)

type LabelCannedServiceTireConnection struct {
	CompanyID string `gorm:"not null;column:companyId" json:"companyId"`

	LabelID             string              `gorm:"primaryKey;not null;column:labelId" json:"labelId"`
	CannedServiceTireID string              `gorm:"primaryKey;not null;column:cannedServiceTireId" json:"cannedServiceTireId"`
	Meta                *datatypes.JSON     `gorm:"column:meta;not null;column:meta" json:"meta,omitempty"`    // the metadata about the most recent change to the row
	Metadata            *datatypes.JSON     `gorm:"column:metadata;column:metadata" json:"metadata,omitempty"` // metadata reserved for customers to control
	CreatedDate         datatypes.DateTime  `gorm:"column:createdDate;not null;column:createdDate" json:"createdDate"`
	UpdatedDate         *datatypes.DateTime `gorm:"column:updatedDate;column:updatedDate" json:"updatedDate"`
}

var _ Model = (*LabelCannedServiceTireConnection)(nil)

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *LabelCannedServiceTireConnection) TableName() string {
	return "label_canned_service_tire_connection"
}

// PrimaryKeys returns an array of the primary keys for this model
func (m *LabelCannedServiceTireConnection) PrimaryKeys() []string {
	return []string{"labelId", "cannedServiceTireId"}
}

// String returns a string representation as JSON for this model
func (m *LabelCannedServiceTireConnection) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewLabelCannedServiceTireConnection returns a new model instance from an encoded buffer
func NewLabelCannedServiceTireConnection(buf []byte) (*LabelCannedServiceTireConnection, error) {
	var result LabelCannedServiceTireConnection
	err := json.Unmarshal(buf, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// NewLabelCannedServiceTireConnectionFromChangeEvent returns a new model instance from an encoded buffer as change event
func NewLabelCannedServiceTireConnectionFromChangeEvent(buf []byte, gzip bool) (*datatypes.ChangeEvent[LabelCannedServiceTireConnection], error) {
	var result datatypes.ChangeEvent[LabelCannedServiceTireConnection]
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
