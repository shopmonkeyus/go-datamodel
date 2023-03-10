// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	datatypes "github.com/shopmonkeyus/go-datamodel/datatypes"
)

// InspectionItemQuickNote schema
type InspectionItemQuickNote struct {
	ID          string              `bson:"_id" gorm:"primaryKey;not null;column:id" json:"id"`
	CreatedDate datatypes.DateTime  `gorm:"column:createdDate;not null;column:createdDate" json:"createdDate"`
	UpdatedDate *datatypes.DateTime `gorm:"column:updatedDate;column:updatedDate" json:"updatedDate"`
	Meta        *datatypes.JSON     `gorm:"column:meta;not null;column:meta" json:"meta,omitempty"` // the metadata about the most recent change to the row
	Metadata    *datatypes.JSON     `gorm:"column:metadata;column:metadata" json:"metadata,omitempty"`
	CompanyID   string              `gorm:"not null;column:companyId" json:"companyId"`
	LocationID  string              `gorm:"not null;column:locationId" json:"locationId"`

	CreatedByID *string `gorm:"column:createdById" json:"createdById"`
	Label       string  `gorm:"not null;column:label" json:"label"`
	Note        string  `gorm:"not null;column:note" json:"note"`
}

var _ Model = (*InspectionItemQuickNote)(nil)

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *InspectionItemQuickNote) TableName() string {
	return "inspection_item_quick_note"
}

// PrimaryKeys returns an array of the primary keys for this model
func (m *InspectionItemQuickNote) PrimaryKeys() []string {
	return []string{"id"}
}

// String returns a string representation as JSON for this model
func (m *InspectionItemQuickNote) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewInspectionItemQuickNote returns a new model instance from an encoded buffer
func NewInspectionItemQuickNote(buf []byte) (*InspectionItemQuickNote, error) {
	var result InspectionItemQuickNote
	err := json.Unmarshal(buf, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// NewInspectionItemQuickNoteFromChangeEvent returns a new model instance from an encoded buffer as change event
func NewInspectionItemQuickNoteFromChangeEvent(buf []byte, gzip bool) (*datatypes.ChangeEvent[InspectionItemQuickNote], error) {
	var result datatypes.ChangeEvent[InspectionItemQuickNote]
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
