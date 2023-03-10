// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	datatypes "github.com/shopmonkeyus/go-datamodel/datatypes"
)

// LabelOrderConnection schema
type LabelOrderConnection struct {
	LabelID string `gorm:"primaryKey;not null;column:labelId" json:"labelId"`
	OrderID string `gorm:"primaryKey;not null;column:orderId" json:"orderId"`

	CompanyID   string              `gorm:"not null;column:companyId" json:"companyId"`
	CreatedDate datatypes.DateTime  `gorm:"column:createdDate;not null;column:createdDate" json:"createdDate"`
	Meta        *datatypes.JSON     `gorm:"column:meta;not null;column:meta" json:"meta,omitempty"` // the metadata about the most recent change to the row
	Metadata    *datatypes.JSON     `gorm:"column:metadata;column:metadata" json:"metadata,omitempty"`
	UpdatedDate *datatypes.DateTime `gorm:"column:updatedDate;column:updatedDate" json:"updatedDate"`
}

var _ Model = (*LabelOrderConnection)(nil)

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *LabelOrderConnection) TableName() string {
	return "label_order_connection"
}

// PrimaryKeys returns an array of the primary keys for this model
func (m *LabelOrderConnection) PrimaryKeys() []string {
	return []string{"labelId", "orderId"}
}

// String returns a string representation as JSON for this model
func (m *LabelOrderConnection) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewLabelOrderConnection returns a new model instance from an encoded buffer
func NewLabelOrderConnection(buf []byte) (*LabelOrderConnection, error) {
	var result LabelOrderConnection
	err := json.Unmarshal(buf, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// NewLabelOrderConnectionFromChangeEvent returns a new model instance from an encoded buffer as change event
func NewLabelOrderConnectionFromChangeEvent(buf []byte, gzip bool) (*datatypes.ChangeEvent[LabelOrderConnection], error) {
	var result datatypes.ChangeEvent[LabelOrderConnection]
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
