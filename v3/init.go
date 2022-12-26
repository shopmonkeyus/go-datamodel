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
func NewFromModel(name string, buf []byte) (Model, error) {
	switch name {
	case "InventoryFee", "inventory_fee":
		{
			return NewInventoryFee(buf)
		}
	case "VehicleOwner", "vehicle_owner":
		{
			return NewVehicleOwner(buf)
		}
	case "Brand", "brand":
		{
			return NewBrand(buf)
		}
	case "TirePressureLog", "tire_pressure_log":
		{
			return NewTirePressureLog(buf)
		}
	case "InspectionItem", "inspection_item":
		{
			return NewInspectionItem(buf)
		}
	case "LabelFeeConnection", "label_fee_connection":
		{
			return NewLabelFeeConnection(buf)
		}
	case "LabelCannedServiceSubcontractConnection", "label_canned_service_subcontract_connection":
		{
			return NewLabelCannedServiceSubcontractConnection(buf)
		}
	case "Labor", "labor":
		{
			return NewLabor(buf)
		}
	case "Email", "email":
		{
			return NewEmail(buf)
		}
	case "PricingMatrixRange", "pricing_matrix_range":
		{
			return NewPricingMatrixRange(buf)
		}
	case "LaborMatrix", "labor_matrix":
		{
			return NewLaborMatrix(buf)
		}
	case "CannedServiceFee", "canned_service_fee":
		{
			return NewCannedServiceFee(buf)
		}
	case "PaymentTerm", "payment_term":
		{
			return NewPaymentTerm(buf)
		}
	case "Reminder", "reminder":
		{
			return NewReminder(buf)
		}
	case "LabelTireConnection", "label_tire_connection":
		{
			return NewLabelTireConnection(buf)
		}
	case "PricingMatrix", "pricing_matrix":
		{
			return NewPricingMatrix(buf)
		}
	case "Subcontract", "subcontract":
		{
			return NewSubcontract(buf)
		}
	case "Fee", "fee":
		{
			return NewFee(buf)
		}
	case "LabelOrderConnection", "label_order_connection":
		{
			return NewLabelOrderConnection(buf)
		}
	case "AuthorizationService", "authorization_service":
		{
			return NewAuthorizationService(buf)
		}
	case "CannedServiceLabor", "canned_service_labor":
		{
			return NewCannedServiceLabor(buf)
		}
	case "VehicleLocationConnection", "vehicle_location":
		{
			return NewVehicleLocationConnection(buf)
		}
	case "Location", "location":
		{
			return NewLocation(buf)
		}
	case "InventoryCategory", "inventory_category":
		{
			return NewInventoryCategory(buf)
		}
	case "PurchaseOrder", "purchase_order":
		{
			return NewPurchaseOrder(buf)
		}
	case "Conversation", "conversation":
		{
			return NewConversation(buf)
		}
	case "LabelCannedServiceFeeConnection", "label_canned_service_fee_connection":
		{
			return NewLabelCannedServiceFeeConnection(buf)
		}
	case "LabelCannedServiceConnection", "label_canned_service":
		{
			return NewLabelCannedServiceConnection(buf)
		}
	case "Payment", "payment":
		{
			return NewPayment(buf)
		}
	case "InspectionItemQuickNote", "inspection_item_quick_note":
		{
			return NewInspectionItemQuickNote(buf)
		}
	case "LabelPartConnection", "label_part_connection":
		{
			return NewLabelPartConnection(buf)
		}
	case "Statement", "statement":
		{
			return NewStatement(buf)
		}
	case "TaxConfig", "tax_config":
		{
			return NewTaxConfig(buf)
		}
	case "Vehicle", "vehicle":
		{
			return NewVehicle(buf)
		}
	case "PurchaseOrderItem", "purchase_order_item":
		{
			return NewPurchaseOrderItem(buf)
		}
	case "Message", "message":
		{
			return NewMessage(buf)
		}
	case "CannedServicePart", "canned_service_part":
		{
			return NewCannedServicePart(buf)
		}
	case "LabelCannedServiceLaborConnection", "label_canned_service_labor_connection":
		{
			return NewLabelCannedServiceLaborConnection(buf)
		}
	case "LabelCustomerConnection", "label_customer_connection":
		{
			return NewLabelCustomerConnection(buf)
		}
	case "CannedServiceSubcontract", "canned_service_subcontract":
		{
			return NewCannedServiceSubcontract(buf)
		}
	case "TpiScanInspection", "tpi_scan_inspection":
		{
			return NewTpiScanInspection(buf)
		}
	case "Authorization", "authorization":
		{
			return NewAuthorization(buf)
		}
	case "PhoneNumber", "phone_number":
		{
			return NewPhoneNumber(buf)
		}
	case "Tire", "tire":
		{
			return NewTire(buf)
		}
	case "Order", "order":
		{
			return NewOrder(buf)
		}
	case "LaborRate", "labor_rate":
		{
			return NewLaborRate(buf)
		}
	case "LabelCannedServiceTireConnection", "label_canned_service_tire_connection":
		{
			return NewLabelCannedServiceTireConnection(buf)
		}
	case "Vendor", "vendor":
		{
			return NewVendor(buf)
		}
	case "Timesheet", "timesheet":
		{
			return NewTimesheet(buf)
		}
	case "Inspection", "inspection":
		{
			return NewInspection(buf)
		}
	case "LabelCannedServicePartConnection", "label_canned_service_part_connection":
		{
			return NewLabelCannedServicePartConnection(buf)
		}
	case "CustomerLocationConnection", "customer_location_connection":
		{
			return NewCustomerLocationConnection(buf)
		}
	case "Part", "part":
		{
			return NewPart(buf)
		}
	case "InspectionTemplateItem", "inspection_template_item":
		{
			return NewInspectionTemplateItem(buf)
		}
	case "LabelLaborConnection", "label_labor_connection":
		{
			return NewLabelLaborConnection(buf)
		}
	case "LabelSubcontractConnection", "label_subcontract_connection":
		{
			return NewLabelSubcontractConnection(buf)
		}
	case "CoreCharge", "core_charge":
		{
			return NewCoreCharge(buf)
		}
	case "User", "user":
		{
			return NewUser(buf)
		}
	case "CannedService", "canned_service":
		{
			return NewCannedService(buf)
		}
	case "CannedServiceTire", "canned_service_tire":
		{
			return NewCannedServiceTire(buf)
		}
	case "Service", "service":
		{
			return NewService(buf)
		}
	case "LaborMatrixRange", "labor_matrix_range":
		{
			return NewLaborMatrixRange(buf)
		}
	case "Customer", "customer":
		{
			return NewCustomer(buf)
		}
	case "ReferralSource", "referral_source":
		{
			return NewReferralSource(buf)
		}
	case "InventoryPart", "inventory_part":
		{
			return NewInventoryPart(buf)
		}
	case "InventoryLabor", "inventory_labor":
		{
			return NewInventoryLabor(buf)
		}
	case "InspectionTemplate", "inspection_template":
		{
			return NewInspectionTemplate(buf)
		}
	case "Label", "label":
		{
			return NewLabel(buf)
		}
	case "LabelVehicleConnection", "label_vehicle_connection":
		{
			return NewLabelVehicleConnection(buf)
		}
	case "WorkflowStatus", "workflow_status":
		{
			return NewWorkflowStatus(buf)
		}
	case "Appointment", "appointment":
		{
			return NewAppointment(buf)
		}
	case "TpiScan", "tpi_scan":
		{
			return NewTpiScan(buf)
		}
	}

	return nil, errors.New("invalid model: " + name)
}
