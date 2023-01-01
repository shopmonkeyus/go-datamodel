// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	datatypes "github.com/shopmonkeyus/go-datamodel/datatypes"
)

type Statement struct {
	ID          string              `gorm:"primaryKey;not null;column:id" json:"id"`
	CreatedDate datatypes.DateTime  `gorm:"column:createdDate;not null;column:createdDate" json:"createdDate"`
	UpdatedDate *datatypes.DateTime `gorm:"column:updatedDate;column:updatedDate" json:"updatedDate"`
	Meta        datatypes.Meta      `gorm:"column:meta;not null;column:meta" json:"meta,omitempty"`    // the metadata about the most recent change to the row
	Metadata    *datatypes.JSON     `gorm:"column:metadata;column:metadata" json:"metadata,omitempty"` // metadata reserved for customers to control
	CompanyID   string              `gorm:"not null;column:companyId" json:"companyId"`

	FromDate datatypes.DateTime  `gorm:"not null;column:fromDate" json:"fromDate"`
	Name     string              `gorm:"not null;column:name" json:"name"`
	SentDate *datatypes.DateTime `gorm:"column:sentDate" json:"sentDate"`
	ToDate   datatypes.DateTime  `gorm:"not null;column:toDate" json:"toDate"`
}

var _ Model = (*Statement)(nil)

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *Statement) TableName() string {
	return "statement"
}

// String returns a string representation as JSON for this model
func (m *Statement) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewStatement returns a new model instance from an encoded buffer
func NewStatement(buf []byte) (*Statement, error) {
	var result Statement
	err := json.Unmarshal(buf, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// NewStatementFromChangeEvent returns a new model instance from an encoded buffer as change event
func NewStatementFromChangeEvent(buf []byte, gzip bool) (*datatypes.ChangeEvent[Statement], error) {
	var result datatypes.ChangeEvent[Statement]
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
