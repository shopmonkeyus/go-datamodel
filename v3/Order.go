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
	ID           string     `gorm:"primaryKey;not null;column:id" json:"id"`
	CreatedDate  time.Time  `gorm:"column:createdDate;not null;column:createdDate" json:"createdDate"`
	UpdatedDate  *time.Time `gorm:"column:updatedDate;column:updatedDate" json:"updatedDate"`
	Meta         *Meta      `gorm:"type:json;embedded;serializer:json;column:meta;not null;column:meta" json:"meta,omitempty"` // the metadata about the most recent change to the row
	Metadata     any        `gorm:"type:json;serializer:json;column:metadata" json:"metadata,omitempty"`                       // metadata reserved for customers to control
	CompanyID    string     `gorm:"not null;column:companyId" json:"companyId"`
	LocationID   string     `gorm:"not null;column:locationId" json:"locationId"`
	CustomFields any        `gorm:"type:json;serializer:json;column:customFields" json:"customFields"` // custom field values

	AppointmentDates       []string                  `gorm:"not null;column:appointmentDates" json:"appointmentDates"`
	Archived               bool                      `gorm:"not null;column:archived" json:"archived"`
	AssignedTechnicianIds  []string                  `gorm:"not null;column:assignedTechnicianIds" json:"assignedTechnicianIds"`
	Authorized             bool                      `gorm:"not null;column:authorized" json:"authorized"`
	AuthorizedDate         *time.Time                `gorm:"column:authorizedDate" json:"authorizedDate"`
	Complaint              *string                   `gorm:"column:complaint" json:"complaint"`
	CompletedDate          *time.Time                `gorm:"column:completedDate" json:"completedDate"`
	ConversationID         string                    `gorm:"not null;column:conversationId" json:"conversationId"`
	CustomerID             *string                   `gorm:"column:customerId" json:"customerId"`
	DeferredServiceCount   int64                     `gorm:"not null;column:deferredServiceCount" json:"deferredServiceCount"`
	Deleted                bool                      `gorm:"not null;column:deleted" json:"deleted"`    // if the record has been deleted
	DeletedDate            *time.Time                `gorm:"column:deletedDate" json:"deletedDate"`     // the date that the record was deleted or null if not deleted
	DeletedReason          *string                   `gorm:"column:deletedReason" json:"deletedReason"` // the reason that the record was deleted
	DeletedUserID          *string                   `gorm:"column:deletedUserId" json:"deletedUserId"` // the user that deleted the record or null if not deleted
	DiscountCents          int64                     `gorm:"not null;column:discountCents" json:"discountCents"`
	DiscountPercent        float64                   `gorm:"not null;column:discountPercent" json:"discountPercent"`
	DueDate                *time.Time                `gorm:"column:dueDate" json:"dueDate"`
	EmailID                *string                   `gorm:"column:emailId" json:"emailId"` // id of the email to use instead of the customer's default email
	EpaCents               int64                     `gorm:"not null;column:epaCents" json:"epaCents"`
	FeesCents              int64                     `gorm:"not null;column:feesCents" json:"feesCents"`
	GeneratedName          *string                   `gorm:"column:generatedName" json:"generatedName"`
	GeneratedVehicleName   *string                   `gorm:"column:generatedVehicleName" json:"generatedVehicleName"` // "[year] [make] [model] [submodel]" pulled from the vehicle, if any
	GstCents               int64                     `gorm:"not null;column:gstCents" json:"gstCents"`
	HstCents               int64                     `gorm:"not null;column:hstCents" json:"hstCents"`
	InspectionCount        int64                     `gorm:"not null;column:inspectionCount" json:"inspectionCount"`
	InspectionStatus       OrderInspectionStatusEnum `gorm:"not null;column:inspectionStatus" json:"inspectionStatus"`
	Invoiced               bool                      `gorm:"not null;column:invoiced" json:"invoiced"`
	InvoicedDate           *time.Time                `gorm:"column:invoicedDate" json:"invoicedDate"`
	LaborCents             int64                     `gorm:"not null;column:laborCents" json:"laborCents"`
	MessageCount           int64                     `gorm:"not null;column:messageCount" json:"messageCount"`
	MileageIn              *int64                    `gorm:"column:mileageIn" json:"mileageIn"`
	MileageOut             *int64                    `gorm:"column:mileageOut" json:"mileageOut"`
	Name                   *string                   `gorm:"column:name" json:"name"`
	Number                 int64                     `gorm:"not null;column:number" json:"number"`
	OrderCreatedDate       time.Time                 `gorm:"not null;column:orderCreatedDate" json:"orderCreatedDate"` // allow user to override created date
	Paid                   bool                      `gorm:"not null;column:paid" json:"paid"`
	PaidCostCents          int64                     `gorm:"not null;column:paidCostCents" json:"paidCostCents"`
	PartsCents             int64                     `gorm:"not null;column:partsCents" json:"partsCents"`
	PhoneNumberID          *string                   `gorm:"column:phoneNumberId" json:"phoneNumberId"` // id of the phone number to use instead of the customer's default number
	Profitability          any                       `gorm:"type:json;serializer:json;column:profitability" json:"profitability"`
	PstCents               int64                     `gorm:"not null;column:pstCents" json:"pstCents"`
	PublicID               string                    `gorm:"not null;column:publicId" json:"publicId"`
	PurchaseOrderNumber    *string                   `gorm:"column:purchaseOrderNumber" json:"purchaseOrderNumber"`
	ReadOnly               bool                      `gorm:"not null;column:readOnly" json:"readOnly"`    // if this order should not be editable in the UI
	ReadOnlyReason         *string                   `gorm:"column:readOnlyReason" json:"readOnlyReason"` // a friendly explanation of why (eg. "migrated from previous system")
	Recommendation         *string                   `gorm:"column:recommendation" json:"recommendation"`
	RemainingCostCents     *int64                    `gorm:"column:remainingCostCents" json:"remainingCostCents"`
	SentToCarfax           bool                      `gorm:"not null;column:sentToCarfax" json:"sentToCarfax"`
	ServiceWriterID        *string                   `gorm:"column:serviceWriterId" json:"serviceWriterId"`
	ShopSuppliesCents      int64                     `gorm:"not null;column:shopSuppliesCents" json:"shopSuppliesCents"`
	ShopUnreadMessageCount int64                     `gorm:"not null;column:shopUnreadMessageCount" json:"shopUnreadMessageCount"`
	SubcontractsCents      int64                     `gorm:"not null;column:subcontractsCents" json:"subcontractsCents"`
	TaxCents               int64                     `gorm:"not null;column:taxCents" json:"taxCents"`
	TaxConfigID            string                    `gorm:"not null;column:taxConfigId" json:"taxConfigId"`
	TiresCents             int64                     `gorm:"not null;column:tiresCents" json:"tiresCents"`
	TotalCostCents         int64                     `gorm:"not null;column:totalCostCents" json:"totalCostCents"`
	VehicleID              *string                   `gorm:"column:vehicleId" json:"vehicleId"`
	WorkflowStatusID       string                    `gorm:"not null;column:workflowStatusId" json:"workflowStatusId"`
	WorkflowStatusPosition float64                   `gorm:"not null;column:workflowStatusPosition" json:"workflowStatusPosition"`
}

var _ Model = (*Order)(nil)

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *Order) TableName() string {
	return "order"
}

func (m *Order) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewOrder returns a new model instance from a json key/value map
func NewOrder(buf []byte) (*Order, error) {
	var result Order
	err := json.Unmarshal(buf, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
