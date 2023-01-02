// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	datatypes "github.com/shopmonkeyus/go-datamodel/datatypes"
)

type LabelCannedServiceFeeConnection struct {
	Metadata *datatypes.JSON `gorm:"column:metadata;column:metadata" json:"metadata,omitempty"` // metadata reserved for customers to control

	CreatedDate        datatypes.DateTime  `gorm:"column:createdDate;not null;column:createdDate" json:"createdDate"`
	UpdatedDate        *datatypes.DateTime `gorm:"column:updatedDate;column:updatedDate" json:"updatedDate"`
	CompanyID          string              `gorm:"not null;column:companyId" json:"companyId"`
	LabelID            string              `gorm:"primaryKey;not null;column:labelId" json:"labelId"`
	CannedServiceFeeID string              `gorm:"primaryKey;not null;column:cannedServiceFeeId" json:"cannedServiceFeeId"`
	Meta               *datatypes.JSON     `gorm:"column:meta;not null;column:meta" json:"meta,omitempty"` // the metadata about the most recent change to the row
}

var _ Model = (*LabelCannedServiceFeeConnection)(nil)

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *LabelCannedServiceFeeConnection) TableName() string {
	return "label_canned_service_fee_connection"
}

// PrimaryKeys returns an array of the primary keys for this model
func (m *LabelCannedServiceFeeConnection) PrimaryKeys() []string {
	return []string{"labelId", "cannedServiceFeeId"}
}

// String returns a string representation as JSON for this model
func (m *LabelCannedServiceFeeConnection) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewLabelCannedServiceFeeConnection returns a new model instance from an encoded buffer
func NewLabelCannedServiceFeeConnection(buf []byte) (*LabelCannedServiceFeeConnection, error) {
	var result LabelCannedServiceFeeConnection
	err := json.Unmarshal(buf, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// NewLabelCannedServiceFeeConnectionFromChangeEvent returns a new model instance from an encoded buffer as change event
func NewLabelCannedServiceFeeConnectionFromChangeEvent(buf []byte, gzip bool) (*datatypes.ChangeEvent[LabelCannedServiceFeeConnection], error) {
	var result datatypes.ChangeEvent[LabelCannedServiceFeeConnection]
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
