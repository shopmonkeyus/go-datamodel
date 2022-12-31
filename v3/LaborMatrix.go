// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	codec "github.com/hashicorp/go-msgpack/v2/codec"
	datatypes "github.com/shopmonkeyus/go-datamodel/datatypes"
	"time"
)

type LaborMatrixTypeEnum string

const (
	LaborMatrixTypeHours LaborMatrixTypeEnum = "Hours"
	LaborMatrixTypeRate                      = "Rate"
)

type LaborMatrix struct {
	ID          string          `gorm:"primaryKey;not null;column:id" json:"id"`
	CreatedDate time.Time       `gorm:"column:createdDate;not null;column:createdDate" json:"createdDate"`
	UpdatedDate *time.Time      `gorm:"column:updatedDate;column:updatedDate" json:"updatedDate"`
	Meta        datatypes.Meta  `gorm:"column:meta;not null;column:meta" json:"meta,omitempty"`    // the metadata about the most recent change to the row
	Metadata    *datatypes.JSON `gorm:"column:metadata;column:metadata" json:"metadata,omitempty"` // metadata reserved for customers to control
	CompanyID   string          `gorm:"not null;column:companyId" json:"companyId"`
	LocationID  string          `gorm:"not null;column:locationId" json:"locationId"`

	Default bool                `gorm:"not null;column:default" json:"default"`
	Name    string              `gorm:"not null;column:name" json:"name"`
	Type    LaborMatrixTypeEnum `gorm:"not null;column:type" json:"type"`
}

var _ Model = (*LaborMatrix)(nil)

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *LaborMatrix) TableName() string {
	return "labor_matrix"
}

func (m *LaborMatrix) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewLaborMatrix returns a new model instance from an encoded buffer
func NewLaborMatrix(buf []byte, enctype EncodingType) (*LaborMatrix, error) {
	var result LaborMatrix
	var handle codec.Handle
	if enctype == JSONEncoding {
		handle = &jsonHandle
	} else {
		handle = &msgpackHandle
	}
	dec := codec.NewDecoderBytes(buf, handle)
	err := dec.Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
