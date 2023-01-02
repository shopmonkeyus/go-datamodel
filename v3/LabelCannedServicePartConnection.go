// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	datatypes "github.com/shopmonkeyus/go-datamodel/datatypes"
)

type LabelCannedServicePartConnection struct {
	LabelID             string `gorm:"primaryKey;not null;column:labelId" json:"labelId"`
	CannedServicePartID string `gorm:"primaryKey;not null;column:cannedServicePartId" json:"cannedServicePartId"`

	CompanyID   string              `gorm:"not null;column:companyId" json:"companyId"`
	CreatedDate datatypes.DateTime  `gorm:"column:createdDate;not null;column:createdDate" json:"createdDate"`
	Meta        *datatypes.JSON     `gorm:"column:meta;not null;column:meta" json:"meta,omitempty"`    // the metadata about the most recent change to the row
	Metadata    *datatypes.JSON     `gorm:"column:metadata;column:metadata" json:"metadata,omitempty"` // metadata reserved for customers to control
	UpdatedDate *datatypes.DateTime `gorm:"column:updatedDate;column:updatedDate" json:"updatedDate"`
}

var _ Model = (*LabelCannedServicePartConnection)(nil)

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *LabelCannedServicePartConnection) TableName() string {
	return "label_canned_service_part_connection"
}

// PrimaryKeys returns an array of the primary keys for this model
func (m *LabelCannedServicePartConnection) PrimaryKeys() []string {
	return []string{"labelId", "cannedServicePartId"}
}

// String returns a string representation as JSON for this model
func (m *LabelCannedServicePartConnection) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewLabelCannedServicePartConnection returns a new model instance from an encoded buffer
func NewLabelCannedServicePartConnection(buf []byte) (*LabelCannedServicePartConnection, error) {
	var result LabelCannedServicePartConnection
	err := json.Unmarshal(buf, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// NewLabelCannedServicePartConnectionFromChangeEvent returns a new model instance from an encoded buffer as change event
func NewLabelCannedServicePartConnectionFromChangeEvent(buf []byte, gzip bool) (*datatypes.ChangeEvent[LabelCannedServicePartConnection], error) {
	var result datatypes.ChangeEvent[LabelCannedServicePartConnection]
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
