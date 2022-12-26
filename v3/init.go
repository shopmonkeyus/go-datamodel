// Code generated. DO NOT EDIT.
package v3

import "errors"

type Meta struct {
	UserID    *string `json:"userId,omitempty"`
	SessionID *string `json:"sessionId,omitempty"`
	Version   *int64  `json:"version,omitempty"`
}

type Model interface {
	TableName() string
	String() string
}

// NewFromModel will return a model from a model name and keyvalue map as JSON
func NewFromModel(name string, kv map[string]any) (Model, error) {
	switch name {
	case "LabelFeeConnection", "label_fee_connection":
		{
			return NewLabelFeeConnection(kv)
		}
	case "InventoryCategory", "inventory_category":
		{
			return NewInventoryCategory(kv)
		}
	case "TpiScan", "tpi_scan":
		{
			return NewTpiScan(kv)
		}
	case "Vehicle", "vehicle":
		{
			return NewVehicle(kv)
		}
	case "WorkflowStatus", "workflow_status":
		{
			return NewWorkflowStatus(kv)
		}
	case "PricingMatrixRange", "pricing_matrix_range":
		{
			return NewPricingMatrixRange(kv)
		}
	case "Tire", "tire":
		{
			return NewTire(kv)
		}
	case "InspectionTemplate", "inspection_template":
		{
			return NewInspectionTemplate(kv)
		}
	case "LabelOrderConnection", "label_order_connection":
		{
			return NewLabelOrderConnection(kv)
		}
	case "LabelCannedServiceTireConnection", "label_canned_service_tire_connection":
		{
			return NewLabelCannedServiceTireConnection(kv)
		}
	case "Customer", "customer":
		{
			return NewCustomer(kv)
		}
	case "Authorization", "authorization":
		{
			return NewAuthorization(kv)
		}
	case "Fee", "fee":
		{
			return NewFee(kv)
		}
	case "VehicleOwner", "vehicle_owner":
		{
			return NewVehicleOwner(kv)
		}
	case "LaborMatrix", "labor_matrix":
		{
			return NewLaborMatrix(kv)
		}
	case "TpiScanInspection", "tpi_scan_inspection":
		{
			return NewTpiScanInspection(kv)
		}
	case "Labor", "labor":
		{
			return NewLabor(kv)
		}
	case "ReferralSource", "referral_source":
		{
			return NewReferralSource(kv)
		}
	case "PurchaseOrderItem", "purchase_order_item":
		{
			return NewPurchaseOrderItem(kv)
		}
	case "Reminder", "reminder":
		{
			return NewReminder(kv)
		}
	case "CannedService", "canned_service":
		{
			return NewCannedService(kv)
		}
	case "InventoryFee", "inventory_fee":
		{
			return NewInventoryFee(kv)
		}
	case "VehicleLocationConnection", "vehicle_location":
		{
			return NewVehicleLocationConnection(kv)
		}
	case "LaborMatrixRange", "labor_matrix_range":
		{
			return NewLaborMatrixRange(kv)
		}
	case "CannedServicePart", "canned_service_part":
		{
			return NewCannedServicePart(kv)
		}
	case "Inspection", "inspection":
		{
			return NewInspection(kv)
		}
	case "Conversation", "conversation":
		{
			return NewConversation(kv)
		}
	case "User", "user":
		{
			return NewUser(kv)
		}
	case "LabelCannedServiceSubcontractConnection", "label_canned_service_subcontract_connection":
		{
			return NewLabelCannedServiceSubcontractConnection(kv)
		}
	case "Statement", "statement":
		{
			return NewStatement(kv)
		}
	case "InventoryLabor", "inventory_labor":
		{
			return NewInventoryLabor(kv)
		}
	case "Timesheet", "timesheet":
		{
			return NewTimesheet(kv)
		}
	case "Email", "email":
		{
			return NewEmail(kv)
		}
	case "TaxConfig", "tax_config":
		{
			return NewTaxConfig(kv)
		}
	case "PurchaseOrder", "purchase_order":
		{
			return NewPurchaseOrder(kv)
		}
	case "Order", "order":
		{
			return NewOrder(kv)
		}
	case "InspectionTemplateItem", "inspection_template_item":
		{
			return NewInspectionTemplateItem(kv)
		}
	case "LabelCustomerConnection", "label_customer_connection":
		{
			return NewLabelCustomerConnection(kv)
		}
	case "LabelCannedServiceConnection", "label_canned_service":
		{
			return NewLabelCannedServiceConnection(kv)
		}
	case "Service", "service":
		{
			return NewService(kv)
		}
	case "Label", "label":
		{
			return NewLabel(kv)
		}
	case "Payment", "payment":
		{
			return NewPayment(kv)
		}
	case "CustomerLocationConnection", "customer_location_connection":
		{
			return NewCustomerLocationConnection(kv)
		}
	case "Part", "part":
		{
			return NewPart(kv)
		}
	case "LabelLaborConnection", "label_labor_connection":
		{
			return NewLabelLaborConnection(kv)
		}
	case "LabelVehicleConnection", "label_vehicle_connection":
		{
			return NewLabelVehicleConnection(kv)
		}
	case "PaymentTerm", "payment_term":
		{
			return NewPaymentTerm(kv)
		}
	case "Vendor", "vendor":
		{
			return NewVendor(kv)
		}
	case "Brand", "brand":
		{
			return NewBrand(kv)
		}
	case "LaborRate", "labor_rate":
		{
			return NewLaborRate(kv)
		}
	case "CannedServiceSubcontract", "canned_service_subcontract":
		{
			return NewCannedServiceSubcontract(kv)
		}
	case "Subcontract", "subcontract":
		{
			return NewSubcontract(kv)
		}
	case "LabelCannedServiceFeeConnection", "label_canned_service_fee_connection":
		{
			return NewLabelCannedServiceFeeConnection(kv)
		}
	case "InspectionItemQuickNote", "inspection_item_quick_note":
		{
			return NewInspectionItemQuickNote(kv)
		}
	case "PhoneNumber", "phone_number":
		{
			return NewPhoneNumber(kv)
		}
	case "CoreCharge", "core_charge":
		{
			return NewCoreCharge(kv)
		}
	case "PricingMatrix", "pricing_matrix":
		{
			return NewPricingMatrix(kv)
		}
	case "AuthorizationService", "authorization_service":
		{
			return NewAuthorizationService(kv)
		}
	case "LabelCannedServiceLaborConnection", "label_canned_service_labor_connection":
		{
			return NewLabelCannedServiceLaborConnection(kv)
		}
	case "LabelCannedServicePartConnection", "label_canned_service_part_connection":
		{
			return NewLabelCannedServicePartConnection(kv)
		}
	case "LabelSubcontractConnection", "label_subcontract_connection":
		{
			return NewLabelSubcontractConnection(kv)
		}
	case "Location", "location":
		{
			return NewLocation(kv)
		}
	case "InventoryPart", "inventory_part":
		{
			return NewInventoryPart(kv)
		}
	case "CannedServiceLabor", "canned_service_labor":
		{
			return NewCannedServiceLabor(kv)
		}
	case "CannedServiceFee", "canned_service_fee":
		{
			return NewCannedServiceFee(kv)
		}
	case "InspectionItem", "inspection_item":
		{
			return NewInspectionItem(kv)
		}
	case "LabelTireConnection", "label_tire_connection":
		{
			return NewLabelTireConnection(kv)
		}
	case "TirePressureLog", "tire_pressure_log":
		{
			return NewTirePressureLog(kv)
		}
	case "CannedServiceTire", "canned_service_tire":
		{
			return NewCannedServiceTire(kv)
		}
	case "LabelPartConnection", "label_part_connection":
		{
			return NewLabelPartConnection(kv)
		}
	case "Appointment", "appointment":
		{
			return NewAppointment(kv)
		}
	case "Message", "message":
		{
			return NewMessage(kv)
		}
	}

	return nil, errors.New("invalid model: " + name)
}
