// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	datatypes "github.com/shopmonkeyus/go-datamodel/datatypes"
)

type LabelEntityEnum string

const (
	LabelEntityCannedService LabelEntityEnum = "CannedService"
	LabelEntityCustomer      LabelEntityEnum = "Customer"
	LabelEntityFee           LabelEntityEnum = "Fee"
	LabelEntityLabor         LabelEntityEnum = "Labor"
	LabelEntityOrder         LabelEntityEnum = "Order"
	LabelEntityPart          LabelEntityEnum = "Part"
	LabelEntityService       LabelEntityEnum = "Service"
	LabelEntitySubcontract   LabelEntityEnum = "Subcontract"
	LabelEntityTire          LabelEntityEnum = "Tire"
	LabelEntityVehicle       LabelEntityEnum = "Vehicle"
)

type Label struct {
	ID          string              `gorm:"primaryKey;not null;column:id" json:"id"`
	CreatedDate datatypes.DateTime  `gorm:"column:createdDate;not null;column:createdDate" json:"createdDate"`
	UpdatedDate *datatypes.DateTime `gorm:"column:updatedDate;column:updatedDate" json:"updatedDate"`
	Meta        *datatypes.JSON     `gorm:"column:meta;not null;column:meta" json:"meta,omitempty"`    // the metadata about the most recent change to the row
	Metadata    *datatypes.JSON     `gorm:"column:metadata;column:metadata" json:"metadata,omitempty"` // metadata reserved for customers to control
	CompanyID   string              `gorm:"not null;column:companyId" json:"companyId"`
	LocationID  string              `gorm:"not null;column:locationId" json:"locationId"`

	Color  string          `gorm:"not null;column:color" json:"color"`
	Entity LabelEntityEnum `gorm:"not null;column:entity" json:"entity"`
	Name   string          `gorm:"not null;column:name" json:"name"`
	Saved  bool            `gorm:"not null;column:saved" json:"saved"`
}

var _ Model = (*Label)(nil)

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *Label) TableName() string {
	return "label"
}

// PrimaryKeys returns an array of the primary keys for this model
func (m *Label) PrimaryKeys() []string {
	return []string{"id"}
}

// String returns a string representation as JSON for this model
func (m *Label) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewLabel returns a new model instance from an encoded buffer
func NewLabel(buf []byte) (*Label, error) {
	var result Label
	err := json.Unmarshal(buf, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// NewLabelFromChangeEvent returns a new model instance from an encoded buffer as change event
func NewLabelFromChangeEvent(buf []byte, gzip bool) (*datatypes.ChangeEvent[Label], error) {
	var result datatypes.ChangeEvent[Label]
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
