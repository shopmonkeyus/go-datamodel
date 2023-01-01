// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	datatypes "github.com/shopmonkeyus/go-datamodel/datatypes"
)

type InventoryLabor struct {
	ID          string              `gorm:"primaryKey;not null;column:id" json:"id"`
	CreatedDate datatypes.DateTime  `gorm:"column:createdDate;not null;column:createdDate" json:"createdDate"`
	UpdatedDate *datatypes.DateTime `gorm:"column:updatedDate;column:updatedDate" json:"updatedDate"`
	Meta        datatypes.Meta      `gorm:"column:meta;not null;column:meta" json:"meta,omitempty"`    // the metadata about the most recent change to the row
	Metadata    *datatypes.JSON     `gorm:"column:metadata;column:metadata" json:"metadata,omitempty"` // metadata reserved for customers to control
	CompanyID   string              `gorm:"not null;column:companyId" json:"companyId"`
	LocationID  string              `gorm:"not null;column:locationId" json:"locationId"`

	CategoryID     *string             `gorm:"column:categoryId" json:"categoryId"`
	Deleted        bool                `gorm:"not null;column:deleted" json:"deleted"`    // if the record has been deleted
	DeletedDate    *datatypes.DateTime `gorm:"column:deletedDate" json:"deletedDate"`     // the date that the record was deleted or null if not deleted
	DeletedReason  *string             `gorm:"column:deletedReason" json:"deletedReason"` // the reason that the record was deleted
	DeletedUserID  *string             `gorm:"column:deletedUserId" json:"deletedUserId"` // the user that deleted the record or null if not deleted
	Name           string              `gorm:"not null;column:name" json:"name"`
	RateID         *string             `gorm:"column:rateId" json:"rateId"`
	Taxable        bool                `gorm:"not null;column:taxable" json:"taxable"`
	TotalCostCents *int64              `gorm:"column:totalCostCents" json:"totalCostCents"`
	UserID         *string             `gorm:"column:userId" json:"userId"`
}

var _ Model = (*InventoryLabor)(nil)

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *InventoryLabor) TableName() string {
	return "inventory_labor"
}

// String returns a string representation as JSON for this model
func (m *InventoryLabor) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewInventoryLabor returns a new model instance from an encoded buffer
func NewInventoryLabor(buf []byte) (*InventoryLabor, error) {
	var result InventoryLabor
	err := json.Unmarshal(buf, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// NewInventoryLaborFromChangeEvent returns a new model instance from an encoded buffer as change event
func NewInventoryLaborFromChangeEvent(buf []byte, gzip bool) (*datatypes.ChangeEvent[InventoryLabor], error) {
	var result datatypes.ChangeEvent[InventoryLabor]
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
