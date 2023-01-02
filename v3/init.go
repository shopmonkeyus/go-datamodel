// Code generated. DO NOT EDIT.
package v3

import "errors"

type Model interface {
	TableName() string
	PrimaryKeys() []string
	String() string
}

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

// TableNames is an array of all the tables defined in this package
var TableNames = []string{
	"appointment",
	"authorization",
	"authorization_service",
	"brand",
	"canned_service",
	"canned_service_fee",
	"canned_service_labor",
	"canned_service_part",
	"canned_service_subcontract",
	"canned_service_tire",
	"conversation",
	"core_charge",
	"customer",
	"customer_location_connection",
	"email",
	"fee",
	"inspection",
	"inspection_item",
	"inspection_item_quick_note",
	"inspection_template",
	"inspection_template_item",
	"inventory_category",
	"inventory_fee",
	"inventory_labor",
	"inventory_part",
	"label",
	"label_canned_service",
	"label_canned_service_fee_connection",
	"label_canned_service_labor_connection",
	"label_canned_service_part_connection",
	"label_canned_service_subcontract_connection",
	"label_canned_service_tire_connection",
	"label_customer_connection",
	"label_fee_connection",
	"label_labor_connection",
	"label_order_connection",
	"label_part_connection",
	"label_subcontract_connection",
	"label_tire_connection",
	"label_vehicle_connection",
	"labor",
	"labor_matrix",
	"labor_matrix_range",
	"labor_rate",
	"location",
	"message",
	"order",
	"part",
	"payment",
	"payment_term",
	"phone_number",
	"pricing_matrix",
	"pricing_matrix_range",
	"purchase_order",
	"purchase_order_item",
	"referral_source",
	"reminder",
	"service",
	"statement",
	"subcontract",
	"tax_config",
	"timesheet",
	"tire",
	"tire_pressure_log",
	"tpi_scan",
	"tpi_scan_inspection",
	"user",
	"vehicle",
	"vehicle_location",
	"vehicle_owner",
	"vendor",
	"workflow_status",
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
func NewFromModel(name string, buf []byte) (Model, error) {
	switch name {
	case "Appointment", "appointment":
		{
			return NewAppointment(buf)
		}
	case "Authorization", "authorization":
		{
			return NewAuthorization(buf)
		}
	case "AuthorizationService", "authorization_service":
		{
			return NewAuthorizationService(buf)
		}
	case "Brand", "brand":
		{
			return NewBrand(buf)
		}
	case "CannedService", "canned_service":
		{
			return NewCannedService(buf)
		}
	case "CannedServiceFee", "canned_service_fee":
		{
			return NewCannedServiceFee(buf)
		}
	case "CannedServiceLabor", "canned_service_labor":
		{
			return NewCannedServiceLabor(buf)
		}
	case "CannedServicePart", "canned_service_part":
		{
			return NewCannedServicePart(buf)
		}
	case "CannedServiceSubcontract", "canned_service_subcontract":
		{
			return NewCannedServiceSubcontract(buf)
		}
	case "CannedServiceTire", "canned_service_tire":
		{
			return NewCannedServiceTire(buf)
		}
	case "Conversation", "conversation":
		{
			return NewConversation(buf)
		}
	case "CoreCharge", "core_charge":
		{
			return NewCoreCharge(buf)
		}
	case "Customer", "customer":
		{
			return NewCustomer(buf)
		}
	case "CustomerLocationConnection", "customer_location_connection":
		{
			return NewCustomerLocationConnection(buf)
		}
	case "Email", "email":
		{
			return NewEmail(buf)
		}
	case "Fee", "fee":
		{
			return NewFee(buf)
		}
	case "Inspection", "inspection":
		{
			return NewInspection(buf)
		}
	case "InspectionItem", "inspection_item":
		{
			return NewInspectionItem(buf)
		}
	case "InspectionItemQuickNote", "inspection_item_quick_note":
		{
			return NewInspectionItemQuickNote(buf)
		}
	case "InspectionTemplate", "inspection_template":
		{
			return NewInspectionTemplate(buf)
		}
	case "InspectionTemplateItem", "inspection_template_item":
		{
			return NewInspectionTemplateItem(buf)
		}
	case "InventoryCategory", "inventory_category":
		{
			return NewInventoryCategory(buf)
		}
	case "InventoryFee", "inventory_fee":
		{
			return NewInventoryFee(buf)
		}
	case "InventoryLabor", "inventory_labor":
		{
			return NewInventoryLabor(buf)
		}
	case "InventoryPart", "inventory_part":
		{
			return NewInventoryPart(buf)
		}
	case "Label", "label":
		{
			return NewLabel(buf)
		}
	case "LabelCannedServiceConnection", "label_canned_service":
		{
			return NewLabelCannedServiceConnection(buf)
		}
	case "LabelCannedServiceFeeConnection", "label_canned_service_fee_connection":
		{
			return NewLabelCannedServiceFeeConnection(buf)
		}
	case "LabelCannedServiceLaborConnection", "label_canned_service_labor_connection":
		{
			return NewLabelCannedServiceLaborConnection(buf)
		}
	case "LabelCannedServicePartConnection", "label_canned_service_part_connection":
		{
			return NewLabelCannedServicePartConnection(buf)
		}
	case "LabelCannedServiceSubcontractConnection", "label_canned_service_subcontract_connection":
		{
			return NewLabelCannedServiceSubcontractConnection(buf)
		}
	case "LabelCannedServiceTireConnection", "label_canned_service_tire_connection":
		{
			return NewLabelCannedServiceTireConnection(buf)
		}
	case "LabelCustomerConnection", "label_customer_connection":
		{
			return NewLabelCustomerConnection(buf)
		}
	case "LabelFeeConnection", "label_fee_connection":
		{
			return NewLabelFeeConnection(buf)
		}
	case "LabelLaborConnection", "label_labor_connection":
		{
			return NewLabelLaborConnection(buf)
		}
	case "LabelOrderConnection", "label_order_connection":
		{
			return NewLabelOrderConnection(buf)
		}
	case "LabelPartConnection", "label_part_connection":
		{
			return NewLabelPartConnection(buf)
		}
	case "LabelSubcontractConnection", "label_subcontract_connection":
		{
			return NewLabelSubcontractConnection(buf)
		}
	case "LabelTireConnection", "label_tire_connection":
		{
			return NewLabelTireConnection(buf)
		}
	case "LabelVehicleConnection", "label_vehicle_connection":
		{
			return NewLabelVehicleConnection(buf)
		}
	case "Labor", "labor":
		{
			return NewLabor(buf)
		}
	case "LaborMatrix", "labor_matrix":
		{
			return NewLaborMatrix(buf)
		}
	case "LaborMatrixRange", "labor_matrix_range":
		{
			return NewLaborMatrixRange(buf)
		}
	case "LaborRate", "labor_rate":
		{
			return NewLaborRate(buf)
		}
	case "Location", "location":
		{
			return NewLocation(buf)
		}
	case "Message", "message":
		{
			return NewMessage(buf)
		}
	case "Order", "order":
		{
			return NewOrder(buf)
		}
	case "Part", "part":
		{
			return NewPart(buf)
		}
	case "Payment", "payment":
		{
			return NewPayment(buf)
		}
	case "PaymentTerm", "payment_term":
		{
			return NewPaymentTerm(buf)
		}
	case "PhoneNumber", "phone_number":
		{
			return NewPhoneNumber(buf)
		}
	case "PricingMatrix", "pricing_matrix":
		{
			return NewPricingMatrix(buf)
		}
	case "PricingMatrixRange", "pricing_matrix_range":
		{
			return NewPricingMatrixRange(buf)
		}
	case "PurchaseOrder", "purchase_order":
		{
			return NewPurchaseOrder(buf)
		}
	case "PurchaseOrderItem", "purchase_order_item":
		{
			return NewPurchaseOrderItem(buf)
		}
	case "ReferralSource", "referral_source":
		{
			return NewReferralSource(buf)
		}
	case "Reminder", "reminder":
		{
			return NewReminder(buf)
		}
	case "Service", "service":
		{
			return NewService(buf)
		}
	case "Statement", "statement":
		{
			return NewStatement(buf)
		}
	case "Subcontract", "subcontract":
		{
			return NewSubcontract(buf)
		}
	case "TaxConfig", "tax_config":
		{
			return NewTaxConfig(buf)
		}
	case "Timesheet", "timesheet":
		{
			return NewTimesheet(buf)
		}
	case "Tire", "tire":
		{
			return NewTire(buf)
		}
	case "TirePressureLog", "tire_pressure_log":
		{
			return NewTirePressureLog(buf)
		}
	case "TpiScan", "tpi_scan":
		{
			return NewTpiScan(buf)
		}
	case "TpiScanInspection", "tpi_scan_inspection":
		{
			return NewTpiScanInspection(buf)
		}
	case "User", "user":
		{
			return NewUser(buf)
		}
	case "Vehicle", "vehicle":
		{
			return NewVehicle(buf)
		}
	case "VehicleLocationConnection", "vehicle_location":
		{
			return NewVehicleLocationConnection(buf)
		}
	case "VehicleOwner", "vehicle_owner":
		{
			return NewVehicleOwner(buf)
		}
	case "Vendor", "vendor":
		{
			return NewVendor(buf)
		}
	case "WorkflowStatus", "workflow_status":
		{
			return NewWorkflowStatus(buf)
		}
	}

	return nil, errors.New("invalid model: " + name)
}

// NewFromChangeEvent will return change event for model from a buffer encoded as EncodingType
func NewFromChangeEvent(name string, buf []byte, gzip bool) (any, error) {
	switch name {
	case "Appointment", "appointment":
		{
			return NewAppointmentFromChangeEvent(buf, gzip)
		}
	case "Authorization", "authorization":
		{
			return NewAuthorizationFromChangeEvent(buf, gzip)
		}
	case "AuthorizationService", "authorization_service":
		{
			return NewAuthorizationServiceFromChangeEvent(buf, gzip)
		}
	case "Brand", "brand":
		{
			return NewBrandFromChangeEvent(buf, gzip)
		}
	case "CannedService", "canned_service":
		{
			return NewCannedServiceFromChangeEvent(buf, gzip)
		}
	case "CannedServiceFee", "canned_service_fee":
		{
			return NewCannedServiceFeeFromChangeEvent(buf, gzip)
		}
	case "CannedServiceLabor", "canned_service_labor":
		{
			return NewCannedServiceLaborFromChangeEvent(buf, gzip)
		}
	case "CannedServicePart", "canned_service_part":
		{
			return NewCannedServicePartFromChangeEvent(buf, gzip)
		}
	case "CannedServiceSubcontract", "canned_service_subcontract":
		{
			return NewCannedServiceSubcontractFromChangeEvent(buf, gzip)
		}
	case "CannedServiceTire", "canned_service_tire":
		{
			return NewCannedServiceTireFromChangeEvent(buf, gzip)
		}
	case "Conversation", "conversation":
		{
			return NewConversationFromChangeEvent(buf, gzip)
		}
	case "CoreCharge", "core_charge":
		{
			return NewCoreChargeFromChangeEvent(buf, gzip)
		}
	case "Customer", "customer":
		{
			return NewCustomerFromChangeEvent(buf, gzip)
		}
	case "CustomerLocationConnection", "customer_location_connection":
		{
			return NewCustomerLocationConnectionFromChangeEvent(buf, gzip)
		}
	case "Email", "email":
		{
			return NewEmailFromChangeEvent(buf, gzip)
		}
	case "Fee", "fee":
		{
			return NewFeeFromChangeEvent(buf, gzip)
		}
	case "Inspection", "inspection":
		{
			return NewInspectionFromChangeEvent(buf, gzip)
		}
	case "InspectionItem", "inspection_item":
		{
			return NewInspectionItemFromChangeEvent(buf, gzip)
		}
	case "InspectionItemQuickNote", "inspection_item_quick_note":
		{
			return NewInspectionItemQuickNoteFromChangeEvent(buf, gzip)
		}
	case "InspectionTemplate", "inspection_template":
		{
			return NewInspectionTemplateFromChangeEvent(buf, gzip)
		}
	case "InspectionTemplateItem", "inspection_template_item":
		{
			return NewInspectionTemplateItemFromChangeEvent(buf, gzip)
		}
	case "InventoryCategory", "inventory_category":
		{
			return NewInventoryCategoryFromChangeEvent(buf, gzip)
		}
	case "InventoryFee", "inventory_fee":
		{
			return NewInventoryFeeFromChangeEvent(buf, gzip)
		}
	case "InventoryLabor", "inventory_labor":
		{
			return NewInventoryLaborFromChangeEvent(buf, gzip)
		}
	case "InventoryPart", "inventory_part":
		{
			return NewInventoryPartFromChangeEvent(buf, gzip)
		}
	case "Label", "label":
		{
			return NewLabelFromChangeEvent(buf, gzip)
		}
	case "LabelCannedServiceConnection", "label_canned_service":
		{
			return NewLabelCannedServiceConnectionFromChangeEvent(buf, gzip)
		}
	case "LabelCannedServiceFeeConnection", "label_canned_service_fee_connection":
		{
			return NewLabelCannedServiceFeeConnectionFromChangeEvent(buf, gzip)
		}
	case "LabelCannedServiceLaborConnection", "label_canned_service_labor_connection":
		{
			return NewLabelCannedServiceLaborConnectionFromChangeEvent(buf, gzip)
		}
	case "LabelCannedServicePartConnection", "label_canned_service_part_connection":
		{
			return NewLabelCannedServicePartConnectionFromChangeEvent(buf, gzip)
		}
	case "LabelCannedServiceSubcontractConnection", "label_canned_service_subcontract_connection":
		{
			return NewLabelCannedServiceSubcontractConnectionFromChangeEvent(buf, gzip)
		}
	case "LabelCannedServiceTireConnection", "label_canned_service_tire_connection":
		{
			return NewLabelCannedServiceTireConnectionFromChangeEvent(buf, gzip)
		}
	case "LabelCustomerConnection", "label_customer_connection":
		{
			return NewLabelCustomerConnectionFromChangeEvent(buf, gzip)
		}
	case "LabelFeeConnection", "label_fee_connection":
		{
			return NewLabelFeeConnectionFromChangeEvent(buf, gzip)
		}
	case "LabelLaborConnection", "label_labor_connection":
		{
			return NewLabelLaborConnectionFromChangeEvent(buf, gzip)
		}
	case "LabelOrderConnection", "label_order_connection":
		{
			return NewLabelOrderConnectionFromChangeEvent(buf, gzip)
		}
	case "LabelPartConnection", "label_part_connection":
		{
			return NewLabelPartConnectionFromChangeEvent(buf, gzip)
		}
	case "LabelSubcontractConnection", "label_subcontract_connection":
		{
			return NewLabelSubcontractConnectionFromChangeEvent(buf, gzip)
		}
	case "LabelTireConnection", "label_tire_connection":
		{
			return NewLabelTireConnectionFromChangeEvent(buf, gzip)
		}
	case "LabelVehicleConnection", "label_vehicle_connection":
		{
			return NewLabelVehicleConnectionFromChangeEvent(buf, gzip)
		}
	case "Labor", "labor":
		{
			return NewLaborFromChangeEvent(buf, gzip)
		}
	case "LaborMatrix", "labor_matrix":
		{
			return NewLaborMatrixFromChangeEvent(buf, gzip)
		}
	case "LaborMatrixRange", "labor_matrix_range":
		{
			return NewLaborMatrixRangeFromChangeEvent(buf, gzip)
		}
	case "LaborRate", "labor_rate":
		{
			return NewLaborRateFromChangeEvent(buf, gzip)
		}
	case "Location", "location":
		{
			return NewLocationFromChangeEvent(buf, gzip)
		}
	case "Message", "message":
		{
			return NewMessageFromChangeEvent(buf, gzip)
		}
	case "Order", "order":
		{
			return NewOrderFromChangeEvent(buf, gzip)
		}
	case "Part", "part":
		{
			return NewPartFromChangeEvent(buf, gzip)
		}
	case "Payment", "payment":
		{
			return NewPaymentFromChangeEvent(buf, gzip)
		}
	case "PaymentTerm", "payment_term":
		{
			return NewPaymentTermFromChangeEvent(buf, gzip)
		}
	case "PhoneNumber", "phone_number":
		{
			return NewPhoneNumberFromChangeEvent(buf, gzip)
		}
	case "PricingMatrix", "pricing_matrix":
		{
			return NewPricingMatrixFromChangeEvent(buf, gzip)
		}
	case "PricingMatrixRange", "pricing_matrix_range":
		{
			return NewPricingMatrixRangeFromChangeEvent(buf, gzip)
		}
	case "PurchaseOrder", "purchase_order":
		{
			return NewPurchaseOrderFromChangeEvent(buf, gzip)
		}
	case "PurchaseOrderItem", "purchase_order_item":
		{
			return NewPurchaseOrderItemFromChangeEvent(buf, gzip)
		}
	case "ReferralSource", "referral_source":
		{
			return NewReferralSourceFromChangeEvent(buf, gzip)
		}
	case "Reminder", "reminder":
		{
			return NewReminderFromChangeEvent(buf, gzip)
		}
	case "Service", "service":
		{
			return NewServiceFromChangeEvent(buf, gzip)
		}
	case "Statement", "statement":
		{
			return NewStatementFromChangeEvent(buf, gzip)
		}
	case "Subcontract", "subcontract":
		{
			return NewSubcontractFromChangeEvent(buf, gzip)
		}
	case "TaxConfig", "tax_config":
		{
			return NewTaxConfigFromChangeEvent(buf, gzip)
		}
	case "Timesheet", "timesheet":
		{
			return NewTimesheetFromChangeEvent(buf, gzip)
		}
	case "Tire", "tire":
		{
			return NewTireFromChangeEvent(buf, gzip)
		}
	case "TirePressureLog", "tire_pressure_log":
		{
			return NewTirePressureLogFromChangeEvent(buf, gzip)
		}
	case "TpiScan", "tpi_scan":
		{
			return NewTpiScanFromChangeEvent(buf, gzip)
		}
	case "TpiScanInspection", "tpi_scan_inspection":
		{
			return NewTpiScanInspectionFromChangeEvent(buf, gzip)
		}
	case "User", "user":
		{
			return NewUserFromChangeEvent(buf, gzip)
		}
	case "Vehicle", "vehicle":
		{
			return NewVehicleFromChangeEvent(buf, gzip)
		}
	case "VehicleLocationConnection", "vehicle_location":
		{
			return NewVehicleLocationConnectionFromChangeEvent(buf, gzip)
		}
	case "VehicleOwner", "vehicle_owner":
		{
			return NewVehicleOwnerFromChangeEvent(buf, gzip)
		}
	case "Vendor", "vendor":
		{
			return NewVendorFromChangeEvent(buf, gzip)
		}
	case "WorkflowStatus", "workflow_status":
		{
			return NewWorkflowStatusFromChangeEvent(buf, gzip)
		}
	}

	return nil, errors.New("invalid model: " + name)
}
