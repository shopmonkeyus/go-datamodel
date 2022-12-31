// Code generated. DO NOT EDIT.
package v3

import (
	"errors"
	codec "github.com/hashicorp/go-msgpack/v2/codec"
)

type Model interface {
	TableName() string
	String() string
}

var jsonHandle codec.JsonHandle
var msgpackHandle codec.MsgpackHandle

type EncodingType string

const (
	JSONEncoding    EncodingType = "json"
	MsgPackEncoding              = "msgpack"
)

// ModelNames is an array of all the models defined in this package
var ModelNames = []string{
	"Appointment",
	"Authorization",
	"AuthorizationService",
	"Brand",
	"CannedService",
	"CannedServiceFee",
	"CannedServiceLabor",
	"CannedServicePart",
	"CannedServiceSubcontract",
	"CannedServiceTire",
	"Conversation",
	"CoreCharge",
	"Customer",
	"CustomerLocationConnection",
	"Email",
	"Fee",
	"Inspection",
	"InspectionItem",
	"InspectionItemQuickNote",
	"InspectionTemplate",
	"InspectionTemplateItem",
	"InventoryCategory",
	"InventoryFee",
	"InventoryLabor",
	"InventoryPart",
	"Label",
	"LabelCannedServiceConnection",
	"LabelCannedServiceFeeConnection",
	"LabelCannedServiceLaborConnection",
	"LabelCannedServicePartConnection",
	"LabelCannedServiceSubcontractConnection",
	"LabelCannedServiceTireConnection",
	"LabelCustomerConnection",
	"LabelFeeConnection",
	"LabelLaborConnection",
	"LabelOrderConnection",
	"LabelPartConnection",
	"LabelSubcontractConnection",
	"LabelTireConnection",
	"LabelVehicleConnection",
	"Labor",
	"LaborMatrix",
	"LaborMatrixRange",
	"LaborRate",
	"Location",
	"Message",
	"Order",
	"Part",
	"Payment",
	"PaymentTerm",
	"PhoneNumber",
	"PricingMatrix",
	"PricingMatrixRange",
	"PurchaseOrder",
	"PurchaseOrderItem",
	"ReferralSource",
	"Reminder",
	"Service",
	"Statement",
	"Subcontract",
	"TaxConfig",
	"Timesheet",
	"Tire",
	"TirePressureLog",
	"TpiScan",
	"TpiScanInspection",
	"User",
	"Vehicle",
	"VehicleLocationConnection",
	"VehicleOwner",
	"Vendor",
	"WorkflowStatus",
}

// ModelInstances is an array of an empty object for each model in this package
var ModelInstances = []interface{}{
	&Appointment{},
	&Authorization{},
	&AuthorizationService{},
	&Brand{},
	&CannedService{},
	&CannedServiceFee{},
	&CannedServiceLabor{},
	&CannedServicePart{},
	&CannedServiceSubcontract{},
	&CannedServiceTire{},
	&Conversation{},
	&CoreCharge{},
	&Customer{},
	&CustomerLocationConnection{},
	&Email{},
	&Fee{},
	&Inspection{},
	&InspectionItem{},
	&InspectionItemQuickNote{},
	&InspectionTemplate{},
	&InspectionTemplateItem{},
	&InventoryCategory{},
	&InventoryFee{},
	&InventoryLabor{},
	&InventoryPart{},
	&Label{},
	&LabelCannedServiceConnection{},
	&LabelCannedServiceFeeConnection{},
	&LabelCannedServiceLaborConnection{},
	&LabelCannedServicePartConnection{},
	&LabelCannedServiceSubcontractConnection{},
	&LabelCannedServiceTireConnection{},
	&LabelCustomerConnection{},
	&LabelFeeConnection{},
	&LabelLaborConnection{},
	&LabelOrderConnection{},
	&LabelPartConnection{},
	&LabelSubcontractConnection{},
	&LabelTireConnection{},
	&LabelVehicleConnection{},
	&Labor{},
	&LaborMatrix{},
	&LaborMatrixRange{},
	&LaborRate{},
	&Location{},
	&Message{},
	&Order{},
	&Part{},
	&Payment{},
	&PaymentTerm{},
	&PhoneNumber{},
	&PricingMatrix{},
	&PricingMatrixRange{},
	&PurchaseOrder{},
	&PurchaseOrderItem{},
	&ReferralSource{},
	&Reminder{},
	&Service{},
	&Statement{},
	&Subcontract{},
	&TaxConfig{},
	&Timesheet{},
	&Tire{},
	&TirePressureLog{},
	&TpiScan{},
	&TpiScanInspection{},
	&User{},
	&Vehicle{},
	&VehicleLocationConnection{},
	&VehicleOwner{},
	&Vendor{},
	&WorkflowStatus{},
}

// NewFromModel will return a model from a model name and buffer encoded as EncodingType
func NewFromModel(name string, buf []byte, enctype EncodingType) (Model, error) {
	switch name {
	case "Appointment", "appointment":
		{
			return NewAppointment(buf, enctype)
		}
	case "Authorization", "authorization":
		{
			return NewAuthorization(buf, enctype)
		}
	case "AuthorizationService", "authorization_service":
		{
			return NewAuthorizationService(buf, enctype)
		}
	case "Brand", "brand":
		{
			return NewBrand(buf, enctype)
		}
	case "CannedService", "canned_service":
		{
			return NewCannedService(buf, enctype)
		}
	case "CannedServiceFee", "canned_service_fee":
		{
			return NewCannedServiceFee(buf, enctype)
		}
	case "CannedServiceLabor", "canned_service_labor":
		{
			return NewCannedServiceLabor(buf, enctype)
		}
	case "CannedServicePart", "canned_service_part":
		{
			return NewCannedServicePart(buf, enctype)
		}
	case "CannedServiceSubcontract", "canned_service_subcontract":
		{
			return NewCannedServiceSubcontract(buf, enctype)
		}
	case "CannedServiceTire", "canned_service_tire":
		{
			return NewCannedServiceTire(buf, enctype)
		}
	case "Conversation", "conversation":
		{
			return NewConversation(buf, enctype)
		}
	case "CoreCharge", "core_charge":
		{
			return NewCoreCharge(buf, enctype)
		}
	case "Customer", "customer":
		{
			return NewCustomer(buf, enctype)
		}
	case "CustomerLocationConnection", "customer_location_connection":
		{
			return NewCustomerLocationConnection(buf, enctype)
		}
	case "Email", "email":
		{
			return NewEmail(buf, enctype)
		}
	case "Fee", "fee":
		{
			return NewFee(buf, enctype)
		}
	case "Inspection", "inspection":
		{
			return NewInspection(buf, enctype)
		}
	case "InspectionItem", "inspection_item":
		{
			return NewInspectionItem(buf, enctype)
		}
	case "InspectionItemQuickNote", "inspection_item_quick_note":
		{
			return NewInspectionItemQuickNote(buf, enctype)
		}
	case "InspectionTemplate", "inspection_template":
		{
			return NewInspectionTemplate(buf, enctype)
		}
	case "InspectionTemplateItem", "inspection_template_item":
		{
			return NewInspectionTemplateItem(buf, enctype)
		}
	case "InventoryCategory", "inventory_category":
		{
			return NewInventoryCategory(buf, enctype)
		}
	case "InventoryFee", "inventory_fee":
		{
			return NewInventoryFee(buf, enctype)
		}
	case "InventoryLabor", "inventory_labor":
		{
			return NewInventoryLabor(buf, enctype)
		}
	case "InventoryPart", "inventory_part":
		{
			return NewInventoryPart(buf, enctype)
		}
	case "Label", "label":
		{
			return NewLabel(buf, enctype)
		}
	case "LabelCannedServiceConnection", "label_canned_service":
		{
			return NewLabelCannedServiceConnection(buf, enctype)
		}
	case "LabelCannedServiceFeeConnection", "label_canned_service_fee_connection":
		{
			return NewLabelCannedServiceFeeConnection(buf, enctype)
		}
	case "LabelCannedServiceLaborConnection", "label_canned_service_labor_connection":
		{
			return NewLabelCannedServiceLaborConnection(buf, enctype)
		}
	case "LabelCannedServicePartConnection", "label_canned_service_part_connection":
		{
			return NewLabelCannedServicePartConnection(buf, enctype)
		}
	case "LabelCannedServiceSubcontractConnection", "label_canned_service_subcontract_connection":
		{
			return NewLabelCannedServiceSubcontractConnection(buf, enctype)
		}
	case "LabelCannedServiceTireConnection", "label_canned_service_tire_connection":
		{
			return NewLabelCannedServiceTireConnection(buf, enctype)
		}
	case "LabelCustomerConnection", "label_customer_connection":
		{
			return NewLabelCustomerConnection(buf, enctype)
		}
	case "LabelFeeConnection", "label_fee_connection":
		{
			return NewLabelFeeConnection(buf, enctype)
		}
	case "LabelLaborConnection", "label_labor_connection":
		{
			return NewLabelLaborConnection(buf, enctype)
		}
	case "LabelOrderConnection", "label_order_connection":
		{
			return NewLabelOrderConnection(buf, enctype)
		}
	case "LabelPartConnection", "label_part_connection":
		{
			return NewLabelPartConnection(buf, enctype)
		}
	case "LabelSubcontractConnection", "label_subcontract_connection":
		{
			return NewLabelSubcontractConnection(buf, enctype)
		}
	case "LabelTireConnection", "label_tire_connection":
		{
			return NewLabelTireConnection(buf, enctype)
		}
	case "LabelVehicleConnection", "label_vehicle_connection":
		{
			return NewLabelVehicleConnection(buf, enctype)
		}
	case "Labor", "labor":
		{
			return NewLabor(buf, enctype)
		}
	case "LaborMatrix", "labor_matrix":
		{
			return NewLaborMatrix(buf, enctype)
		}
	case "LaborMatrixRange", "labor_matrix_range":
		{
			return NewLaborMatrixRange(buf, enctype)
		}
	case "LaborRate", "labor_rate":
		{
			return NewLaborRate(buf, enctype)
		}
	case "Location", "location":
		{
			return NewLocation(buf, enctype)
		}
	case "Message", "message":
		{
			return NewMessage(buf, enctype)
		}
	case "Order", "order":
		{
			return NewOrder(buf, enctype)
		}
	case "Part", "part":
		{
			return NewPart(buf, enctype)
		}
	case "Payment", "payment":
		{
			return NewPayment(buf, enctype)
		}
	case "PaymentTerm", "payment_term":
		{
			return NewPaymentTerm(buf, enctype)
		}
	case "PhoneNumber", "phone_number":
		{
			return NewPhoneNumber(buf, enctype)
		}
	case "PricingMatrix", "pricing_matrix":
		{
			return NewPricingMatrix(buf, enctype)
		}
	case "PricingMatrixRange", "pricing_matrix_range":
		{
			return NewPricingMatrixRange(buf, enctype)
		}
	case "PurchaseOrder", "purchase_order":
		{
			return NewPurchaseOrder(buf, enctype)
		}
	case "PurchaseOrderItem", "purchase_order_item":
		{
			return NewPurchaseOrderItem(buf, enctype)
		}
	case "ReferralSource", "referral_source":
		{
			return NewReferralSource(buf, enctype)
		}
	case "Reminder", "reminder":
		{
			return NewReminder(buf, enctype)
		}
	case "Service", "service":
		{
			return NewService(buf, enctype)
		}
	case "Statement", "statement":
		{
			return NewStatement(buf, enctype)
		}
	case "Subcontract", "subcontract":
		{
			return NewSubcontract(buf, enctype)
		}
	case "TaxConfig", "tax_config":
		{
			return NewTaxConfig(buf, enctype)
		}
	case "Timesheet", "timesheet":
		{
			return NewTimesheet(buf, enctype)
		}
	case "Tire", "tire":
		{
			return NewTire(buf, enctype)
		}
	case "TirePressureLog", "tire_pressure_log":
		{
			return NewTirePressureLog(buf, enctype)
		}
	case "TpiScan", "tpi_scan":
		{
			return NewTpiScan(buf, enctype)
		}
	case "TpiScanInspection", "tpi_scan_inspection":
		{
			return NewTpiScanInspection(buf, enctype)
		}
	case "User", "user":
		{
			return NewUser(buf, enctype)
		}
	case "Vehicle", "vehicle":
		{
			return NewVehicle(buf, enctype)
		}
	case "VehicleLocationConnection", "vehicle_location":
		{
			return NewVehicleLocationConnection(buf, enctype)
		}
	case "VehicleOwner", "vehicle_owner":
		{
			return NewVehicleOwner(buf, enctype)
		}
	case "Vendor", "vendor":
		{
			return NewVendor(buf, enctype)
		}
	case "WorkflowStatus", "workflow_status":
		{
			return NewWorkflowStatus(buf, enctype)
		}
	}

	return nil, errors.New("invalid model: " + name)
}
