// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	datatypes "github.com/shopmonkeyus/go-datamodel/datatypes"
)

// Statement schema
type Statement struct {
	ID          string              `bson:"_id" gorm:"primaryKey;not null;column:id" json:"id"`
	CreatedDate datatypes.DateTime  `gorm:"column:createdDate;not null;column:createdDate" json:"createdDate"`
	UpdatedDate *datatypes.DateTime `gorm:"column:updatedDate;column:updatedDate" json:"updatedDate"`
	Meta        *datatypes.JSON     `gorm:"column:meta;not null;column:meta" json:"meta,omitempty"`    // the metadata about the most recent change to the row
	Metadata    *datatypes.JSON     `gorm:"column:metadata;column:metadata" json:"metadata,omitempty"` // metadata reserved for customers to control
	CompanyID   string              `gorm:"not null;column:companyId" json:"companyId"`
	LocationID  string              `gorm:"not null;column:locationId" json:"locationId"`

	CustomerID         string              `gorm:"not null;column:customerId" json:"customerId"`
	FromDate           datatypes.DateTime  `gorm:"not null;column:fromDate" json:"fromDate"`
	InvoicesCount      int64               `gorm:"not null;column:invoicesCount" json:"invoicesCount"`
	Name               *string             `gorm:"column:name" json:"name"`
	Paid               bool                `gorm:"not null;column:paid" json:"paid"`
	PaidCostCents      int64               `gorm:"not null;column:paidCostCents" json:"paidCostCents"`
	RemainingCostCents int64               `gorm:"not null;column:remainingCostCents" json:"remainingCostCents"`
	Sent               bool                `gorm:"not null;column:sent" json:"sent"`
	SentDate           *datatypes.DateTime `gorm:"column:sentDate" json:"sentDate"`
	ToDate             datatypes.DateTime  `gorm:"not null;column:toDate" json:"toDate"`
	TotalCostCents     int64               `gorm:"not null;column:totalCostCents" json:"totalCostCents"`
}

var _ Model = (*Statement)(nil)

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *Statement) TableName() string {
	return "statement"
}

// PrimaryKeys returns an array of the primary keys for this model
func (m *Statement) PrimaryKeys() []string {
	return []string{"id"}
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
