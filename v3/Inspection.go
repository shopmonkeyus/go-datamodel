// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	datatypes "github.com/shopmonkeyus/go-datamodel/datatypes"
)

// Inspection schema
type Inspection struct {
	ID          string              `bson:"_id" gorm:"primaryKey;not null;column:id" json:"id"`
	CreatedDate datatypes.DateTime  `gorm:"column:createdDate;not null;column:createdDate" json:"createdDate"`
	UpdatedDate *datatypes.DateTime `gorm:"column:updatedDate;column:updatedDate" json:"updatedDate"`
	Meta        *datatypes.JSON     `gorm:"column:meta;not null;column:meta" json:"meta,omitempty"` // the metadata about the most recent change to the row
	Metadata    *datatypes.JSON     `gorm:"column:metadata;column:metadata" json:"metadata,omitempty"`
	CompanyID   string              `gorm:"not null;column:companyId" json:"companyId"`
	LocationID  string              `gorm:"not null;column:locationId" json:"locationId"`

	Completed     bool                `gorm:"not null;column:completed" json:"completed"`
	CompletedByID *string             `gorm:"column:completedById" json:"completedById"`
	CompletedDate *datatypes.DateTime `gorm:"column:completedDate" json:"completedDate"`
	CreatedByID   *string             `gorm:"column:createdById" json:"createdById"`
	Name          string              `gorm:"not null;column:name" json:"name"`
	OrderID       string              `gorm:"not null;column:orderId" json:"orderId"`
	Ordinal       float64             `gorm:"not null;column:ordinal" json:"ordinal"`
	TemplateID    *string             `gorm:"column:templateId" json:"templateId"`
}

var _ Model = (*Inspection)(nil)

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *Inspection) TableName() string {
	return "inspection"
}

// PrimaryKeys returns an array of the primary keys for this model
func (m *Inspection) PrimaryKeys() []string {
	return []string{"id"}
}

// String returns a string representation as JSON for this model
func (m *Inspection) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewInspection returns a new model instance from an encoded buffer
func NewInspection(buf []byte) (*Inspection, error) {
	var result Inspection
	err := json.Unmarshal(buf, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// NewInspectionFromChangeEvent returns a new model instance from an encoded buffer as change event
func NewInspectionFromChangeEvent(buf []byte, gzip bool) (*datatypes.ChangeEvent[Inspection], error) {
	var result datatypes.ChangeEvent[Inspection]
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
