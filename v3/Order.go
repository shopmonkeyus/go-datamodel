// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	datatypes "github.com/shopmonkeyus/go-datamodel/datatypes"
)

// Order schema
type OrderInspectionStatusEnum string

const (
	OrderInspectionStatusNone         OrderInspectionStatusEnum = "None"
	OrderInspectionStatusCompleted    OrderInspectionStatusEnum = "Completed"
	OrderInspectionStatusNotCompleted OrderInspectionStatusEnum = "NotCompleted"
)

type Order struct {
	ID           string              `bson:"_id" gorm:"primaryKey;not null;column:id" json:"id"`
	CreatedDate  datatypes.DateTime  `gorm:"column:createdDate;not null;column:createdDate" json:"createdDate"`
	UpdatedDate  *datatypes.DateTime `gorm:"column:updatedDate;column:updatedDate" json:"updatedDate"`
	Meta         *datatypes.JSON     `gorm:"column:meta;not null;column:meta" json:"meta,omitempty"` // the metadata about the most recent change to the row
	Metadata     *datatypes.JSON     `gorm:"column:metadata;column:metadata" json:"metadata,omitempty"`
	CompanyID    string              `gorm:"not null;column:companyId" json:"companyId"`
	LocationID   string              `gorm:"not null;column:locationId" json:"locationId"`
	CustomFields datatypes.JSON      `gorm:"column:customFields" json:"customFields"`

	AppointmentDates       datatypes.StringArray     `gorm:"not null;column:appointmentDates" json:"appointmentDates"`
	Archived               bool                      `gorm:"not null;column:archived" json:"archived"`
	AssignedTechnicianIds  datatypes.StringArray     `gorm:"not null;column:assignedTechnicianIds" json:"assignedTechnicianIds"`
	Authorized             bool                      `gorm:"not null;column:authorized" json:"authorized"`
	AuthorizedDate         *datatypes.DateTime       `gorm:"column:authorizedDate" json:"authorizedDate"`
	CoalescedName          *string                   `gorm:"column:coalescedName" json:"coalescedName"`
	Complaint              *string                   `gorm:"column:complaint" json:"complaint"`
	CompletedDate          *datatypes.DateTime       `gorm:"column:completedDate" json:"completedDate"`
	ConversationID         string                    `gorm:"not null;column:conversationId" json:"conversationId"`
	CustomerID             *string                   `gorm:"column:customerId" json:"customerId"`
	DeferredServiceCount   int64                     `gorm:"not null;column:deferredServiceCount" json:"deferredServiceCount"`
	Deleted                bool                      `gorm:"not null;column:deleted" json:"deleted"`    // if the record has been deleted
	DeletedDate            *datatypes.DateTime       `gorm:"column:deletedDate" json:"deletedDate"`     // the date that the record was deleted or null if not deleted
	DeletedReason          *string                   `gorm:"column:deletedReason" json:"deletedReason"` // the reason that the record was deleted
	DeletedUserID          *string                   `gorm:"column:deletedUserId" json:"deletedUserId"` // the user that deleted the record or null if not deleted
	DiscountCents          int64                     `gorm:"not null;column:discountCents" json:"discountCents"`
	DiscountPercent        float64                   `gorm:"not null;column:discountPercent" json:"discountPercent"`
	DueDate                *datatypes.DateTime       `gorm:"column:dueDate" json:"dueDate"`
	EmailID                *string                   `gorm:"column:emailId" json:"emailId"` // id of the email to use instead of the customer's default email
	EpaCents               int64                     `gorm:"not null;column:epaCents" json:"epaCents"`
	FeesCents              int64                     `gorm:"not null;column:feesCents" json:"feesCents"`
	FullyPaidDate          *datatypes.DateTime       `gorm:"column:fullyPaidDate" json:"fullyPaidDate"`
	GeneratedCustomerName  *string                   `gorm:"column:generatedCustomerName" json:"generatedCustomerName"` // "[firstName] [lastName]" pulled from the customer, if any
	GeneratedName          *string                   `gorm:"column:generatedName" json:"generatedName"`
	GeneratedVehicleName   *string                   `gorm:"column:generatedVehicleName" json:"generatedVehicleName"` // "[year] [make] [model] [submodel]" pulled from the vehicle, if any
	GstCents               int64                     `gorm:"not null;column:gstCents" json:"gstCents"`
	HstCents               int64                     `gorm:"not null;column:hstCents" json:"hstCents"`
	InspectionCount        int64                     `gorm:"not null;column:inspectionCount" json:"inspectionCount"`
	InspectionStatus       OrderInspectionStatusEnum `gorm:"not null;column:inspectionStatus" json:"inspectionStatus"`
	Invoiced               bool                      `gorm:"not null;column:invoiced" json:"invoiced"`
	InvoicedDate           *datatypes.DateTime       `gorm:"column:invoicedDate" json:"invoicedDate"`
	LaborCents             int64                     `gorm:"not null;column:laborCents" json:"laborCents"`
	MessageCount           int64                     `gorm:"not null;column:messageCount" json:"messageCount"`
	MessagedDate           *datatypes.DateTime       `gorm:"column:messagedDate" json:"messagedDate"`
	MileageIn              *float64                  `gorm:"column:mileageIn" json:"mileageIn"`
	MileageOut             *float64                  `gorm:"column:mileageOut" json:"mileageOut"`
	Name                   *string                   `gorm:"column:name" json:"name"`
	Number                 int64                     `gorm:"not null;column:number" json:"number"`
	OrderCreatedDate       datatypes.DateTime        `gorm:"not null;column:orderCreatedDate" json:"orderCreatedDate"` // allow user to override created date
	Paid                   bool                      `gorm:"not null;column:paid" json:"paid"`
	PaidCostCents          int64                     `gorm:"not null;column:paidCostCents" json:"paidCostCents"`
	PartsCents             int64                     `gorm:"not null;column:partsCents" json:"partsCents"`
	PaymentTermID          string                    `gorm:"not null;column:paymentTermId" json:"paymentTermId"` // id of the payment term for the order
	PhoneNumberID          *string                   `gorm:"column:phoneNumberId" json:"phoneNumberId"`          // id of the phone number to use instead of the customer's default number
	Profitability          datatypes.JSON            `gorm:"column:profitability" json:"profitability"`
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
	StatementID            *string                   `gorm:"column:statementId" json:"statementId"` // a statement this order included in
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

// PrimaryKeys returns an array of the primary keys for this model
func (m *Order) PrimaryKeys() []string {
	return []string{"id"}
}

// String returns a string representation as JSON for this model
func (m *Order) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewOrder returns a new model instance from an encoded buffer
func NewOrder(buf []byte) (*Order, error) {
	var result Order
	err := json.Unmarshal(buf, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// NewOrderFromChangeEvent returns a new model instance from an encoded buffer as change event
func NewOrderFromChangeEvent(buf []byte, gzip bool) (*datatypes.ChangeEvent[Order], error) {
	var result datatypes.ChangeEvent[Order]
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
