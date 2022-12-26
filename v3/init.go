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

// NewFromModel will return a model from a model name and JSON as byte buffer
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
