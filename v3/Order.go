// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	"time"
)

type OrderInspectionStatusEnum string

const (
	OrderInspectionStatusNone         OrderInspectionStatusEnum = "None"
	OrderInspectionStatusCompleted                              = "Completed"
	OrderInspectionStatusNotCompleted                           = "NotCompleted"
)

type Order struct {
	AppointmentDates       []string                  `gorm:"not null" json:"appointmentDates"`
	Archived               bool                      `gorm:"not null" json:"archived"`
	AssignedTechnicianIds  []string                  `gorm:"not null" json:"assignedTechnicianIds"`
	Authorized             bool                      `gorm:"not null" json:"authorized"`
	AuthorizedDate         *time.Time                `json:"authorizedDate"`
	ID                     string                    `gorm:"primaryKey;not null" json:"id"`
	CompanyId              string                    `gorm:"not null" json:"companyId"`
	Complaint              *string                   `json:"complaint"`
	CompletedDate          *time.Time                `json:"completedDate"`
	ConversationId         string                    `gorm:"not null" json:"conversationId"`
	CreatedDate            time.Time                 `gorm:"column:createdDate;not null" json:"createdDate"`
	CustomFields           any                       `gorm:"not null" json:"customFields"` // custom field values
	CustomerId             *string                   `json:"customerId"`
	DeferredServiceCount   int64                     `gorm:"not null" json:"deferredServiceCount"`
	Deleted                bool                      `gorm:"not null" json:"deleted"` // if the record has been deleted
	DeletedDate            *time.Time                `json:"deletedDate"`             // the date that the record was deleted or null if not deleted
	DeletedReason          *string                   `json:"deletedReason"`           // the reason that the record was deleted
	DeletedUserId          *string                   `json:"deletedUserId"`           // the user that deleted the record or null if not deleted
	DiscountCents          int64                     `gorm:"not null" json:"discountCents"`
	DiscountPercent        float64                   `gorm:"not null" json:"discountPercent"`
	DueDate                *time.Time                `json:"dueDate"`
	EmailId                *string                   `json:"emailId"` // id of the email to use instead of the customer's default email
	EpaCents               int64                     `gorm:"not null" json:"epaCents"`
	FeesCents              int64                     `gorm:"not null" json:"feesCents"`
	GeneratedName          *string                   `json:"generatedName"`
	GeneratedVehicleName   *string                   `json:"generatedVehicleName"` // "[year] [make] [model] [submodel]" pulled from the vehicle, if any
	GstCents               int64                     `gorm:"not null" json:"gstCents"`
	HstCents               int64                     `gorm:"not null" json:"hstCents"`
	InspectionCount        int64                     `gorm:"not null" json:"inspectionCount"`
	InspectionStatus       OrderInspectionStatusEnum `gorm:"not null" json:"inspectionStatus"`
	Invoiced               bool                      `gorm:"not null" json:"invoiced"`
	InvoicedDate           *time.Time                `json:"invoicedDate"`
	LaborCents             int64                     `gorm:"not null" json:"laborCents"`
	LocationId             string                    `gorm:"not null" json:"locationId"`
	MessageCount           int64                     `gorm:"not null" json:"messageCount"`
	Meta                   *Meta                     `gorm:"embedded;column:meta;not null" json:"meta,omitempty"`         // the metadata about the most recent change to the row
	Metadata               any                       `gorm:"embedded;column:metadata;not null" json:"metadata,omitempty"` // metadata reserved for customers to control
	MileageIn              *int64                    `json:"mileageIn"`
	MileageOut             *int64                    `json:"mileageOut"`
	Name                   *string                   `json:"name"`
	Number                 int64                     `gorm:"not null" json:"number"`
	OrderCreatedDate       time.Time                 `gorm:"not null" json:"orderCreatedDate"` // allow user to override created date
	Paid                   bool                      `gorm:"not null" json:"paid"`
	PaidCostCents          int64                     `gorm:"not null" json:"paidCostCents"`
	PartsCents             int64                     `gorm:"not null" json:"partsCents"`
	PhoneNumberId          *string                   `json:"phoneNumberId"` // id of the phone number to use instead of the customer's default number
	Profitability          any                       `gorm:"not null" json:"profitability"`
	PstCents               int64                     `gorm:"not null" json:"pstCents"`
	PublicId               string                    `gorm:"not null" json:"publicId"`
	PurchaseOrderNumber    *string                   `json:"purchaseOrderNumber"`
	ReadOnly               bool                      `gorm:"not null" json:"readOnly"` // if this order should not be editable in the UI
	ReadOnlyReason         *string                   `json:"readOnlyReason"`           // a friendly explanation of why (eg. "migrated from previous system")
	Recommendation         *string                   `json:"recommendation"`
	RemainingCostCents     *int64                    `json:"remainingCostCents"`
	SentToCarfax           bool                      `gorm:"not null" json:"sentToCarfax"`
	ServiceWriterId        *string                   `json:"serviceWriterId"`
	ShopSuppliesCents      int64                     `gorm:"not null" json:"shopSuppliesCents"`
	ShopUnreadMessageCount int64                     `gorm:"not null" json:"shopUnreadMessageCount"`
	SubcontractsCents      int64                     `gorm:"not null" json:"subcontractsCents"`
	TaxCents               int64                     `gorm:"not null" json:"taxCents"`
	TaxConfigId            string                    `gorm:"not null" json:"taxConfigId"`
	TiresCents             int64                     `gorm:"not null" json:"tiresCents"`
	TotalCostCents         int64                     `gorm:"not null" json:"totalCostCents"`
	UpdatedDate            *time.Time                `gorm:"column:updatedDate" json:"updatedDate"`
	VehicleId              *string                   `json:"vehicleId"`
	WorkflowStatusId       string                    `gorm:"not null" json:"workflowStatusId"`
	WorkflowStatusPosition float64                   `gorm:"not null" json:"workflowStatusPosition"`
}

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *Order) TableName() string {
	return "order"
}

func (m *Order) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}
