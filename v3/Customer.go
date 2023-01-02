// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	datatypes "github.com/shopmonkeyus/go-datamodel/datatypes"
)

type CustomerCountryEnum string

const (
	CustomerCountryUS CustomerCountryEnum = "US"
	CustomerCountryCA CustomerCountryEnum = "CA"
	CustomerCountryMX CustomerCountryEnum = "MX"
	CustomerCountryAF CustomerCountryEnum = "AF"
	CustomerCountryAX CustomerCountryEnum = "AX"
	CustomerCountryAL CustomerCountryEnum = "AL"
	CustomerCountryDZ CustomerCountryEnum = "DZ"
	CustomerCountryAS CustomerCountryEnum = "AS"
	CustomerCountryAD CustomerCountryEnum = "AD"
	CustomerCountryAO CustomerCountryEnum = "AO"
	CustomerCountryAI CustomerCountryEnum = "AI"
	CustomerCountryAQ CustomerCountryEnum = "AQ"
	CustomerCountryAG CustomerCountryEnum = "AG"
	CustomerCountryAR CustomerCountryEnum = "AR"
	CustomerCountryAM CustomerCountryEnum = "AM"
	CustomerCountryAW CustomerCountryEnum = "AW"
	CustomerCountryAU CustomerCountryEnum = "AU"
	CustomerCountryAT CustomerCountryEnum = "AT"
	CustomerCountryAZ CustomerCountryEnum = "AZ"
	CustomerCountryBS CustomerCountryEnum = "BS"
	CustomerCountryBH CustomerCountryEnum = "BH"
	CustomerCountryBD CustomerCountryEnum = "BD"
	CustomerCountryBB CustomerCountryEnum = "BB"
	CustomerCountryBY CustomerCountryEnum = "BY"
	CustomerCountryBE CustomerCountryEnum = "BE"
	CustomerCountryBZ CustomerCountryEnum = "BZ"
	CustomerCountryBJ CustomerCountryEnum = "BJ"
	CustomerCountryBM CustomerCountryEnum = "BM"
	CustomerCountryBT CustomerCountryEnum = "BT"
	CustomerCountryBO CustomerCountryEnum = "BO"
	CustomerCountryBQ CustomerCountryEnum = "BQ"
	CustomerCountryBA CustomerCountryEnum = "BA"
	CustomerCountryBW CustomerCountryEnum = "BW"
	CustomerCountryBV CustomerCountryEnum = "BV"
	CustomerCountryBR CustomerCountryEnum = "BR"
	CustomerCountryIO CustomerCountryEnum = "IO"
	CustomerCountryVG CustomerCountryEnum = "VG"
	CustomerCountryBN CustomerCountryEnum = "BN"
	CustomerCountryBG CustomerCountryEnum = "BG"
	CustomerCountryBF CustomerCountryEnum = "BF"
	CustomerCountryBI CustomerCountryEnum = "BI"
	CustomerCountryCV CustomerCountryEnum = "CV"
	CustomerCountryKH CustomerCountryEnum = "KH"
	CustomerCountryCM CustomerCountryEnum = "CM"
	CustomerCountryKY CustomerCountryEnum = "KY"
	CustomerCountryCF CustomerCountryEnum = "CF"
	CustomerCountryTD CustomerCountryEnum = "TD"
	CustomerCountryCL CustomerCountryEnum = "CL"
	CustomerCountryCN CustomerCountryEnum = "CN"
	CustomerCountryHK CustomerCountryEnum = "HK"
	CustomerCountryMO CustomerCountryEnum = "MO"
	CustomerCountryCX CustomerCountryEnum = "CX"
	CustomerCountryCC CustomerCountryEnum = "CC"
	CustomerCountryCO CustomerCountryEnum = "CO"
	CustomerCountryKM CustomerCountryEnum = "KM"
	CustomerCountryCG CustomerCountryEnum = "CG"
	CustomerCountryCK CustomerCountryEnum = "CK"
	CustomerCountryCR CustomerCountryEnum = "CR"
	CustomerCountryHR CustomerCountryEnum = "HR"
	CustomerCountryCU CustomerCountryEnum = "CU"
	CustomerCountryCW CustomerCountryEnum = "CW"
	CustomerCountryCY CustomerCountryEnum = "CY"
	CustomerCountryCZ CustomerCountryEnum = "CZ"
	CustomerCountryCI CustomerCountryEnum = "CI"
	CustomerCountryKP CustomerCountryEnum = "KP"
	CustomerCountryCD CustomerCountryEnum = "CD"
	CustomerCountryDK CustomerCountryEnum = "DK"
	CustomerCountryDJ CustomerCountryEnum = "DJ"
	CustomerCountryDM CustomerCountryEnum = "DM"
	CustomerCountryDO CustomerCountryEnum = "DO"
	CustomerCountryEC CustomerCountryEnum = "EC"
	CustomerCountryEG CustomerCountryEnum = "EG"
	CustomerCountrySV CustomerCountryEnum = "SV"
	CustomerCountryGQ CustomerCountryEnum = "GQ"
	CustomerCountryER CustomerCountryEnum = "ER"
	CustomerCountryEE CustomerCountryEnum = "EE"
	CustomerCountrySZ CustomerCountryEnum = "SZ"
	CustomerCountryET CustomerCountryEnum = "ET"
	CustomerCountryFK CustomerCountryEnum = "FK"
	CustomerCountryFO CustomerCountryEnum = "FO"
	CustomerCountryFJ CustomerCountryEnum = "FJ"
	CustomerCountryFI CustomerCountryEnum = "FI"
	CustomerCountryFR CustomerCountryEnum = "FR"
	CustomerCountryGF CustomerCountryEnum = "GF"
	CustomerCountryPF CustomerCountryEnum = "PF"
	CustomerCountryTF CustomerCountryEnum = "TF"
	CustomerCountryGA CustomerCountryEnum = "GA"
	CustomerCountryGM CustomerCountryEnum = "GM"
	CustomerCountryGE CustomerCountryEnum = "GE"
	CustomerCountryDE CustomerCountryEnum = "DE"
	CustomerCountryGH CustomerCountryEnum = "GH"
	CustomerCountryGI CustomerCountryEnum = "GI"
	CustomerCountryGR CustomerCountryEnum = "GR"
	CustomerCountryGL CustomerCountryEnum = "GL"
	CustomerCountryGD CustomerCountryEnum = "GD"
	CustomerCountryGP CustomerCountryEnum = "GP"
	CustomerCountryGU CustomerCountryEnum = "GU"
	CustomerCountryGT CustomerCountryEnum = "GT"
	CustomerCountryGG CustomerCountryEnum = "GG"
	CustomerCountryGN CustomerCountryEnum = "GN"
	CustomerCountryGW CustomerCountryEnum = "GW"
	CustomerCountryGY CustomerCountryEnum = "GY"
	CustomerCountryHT CustomerCountryEnum = "HT"
	CustomerCountryHM CustomerCountryEnum = "HM"
	CustomerCountryVA CustomerCountryEnum = "VA"
	CustomerCountryHN CustomerCountryEnum = "HN"
	CustomerCountryHU CustomerCountryEnum = "HU"
	CustomerCountryIS CustomerCountryEnum = "IS"
	CustomerCountryIN CustomerCountryEnum = "IN"
	CustomerCountryID CustomerCountryEnum = "ID"
	CustomerCountryIR CustomerCountryEnum = "IR"
	CustomerCountryIQ CustomerCountryEnum = "IQ"
	CustomerCountryIE CustomerCountryEnum = "IE"
	CustomerCountryIM CustomerCountryEnum = "IM"
	CustomerCountryIL CustomerCountryEnum = "IL"
	CustomerCountryIT CustomerCountryEnum = "IT"
	CustomerCountryJM CustomerCountryEnum = "JM"
	CustomerCountryJP CustomerCountryEnum = "JP"
	CustomerCountryJE CustomerCountryEnum = "JE"
	CustomerCountryJO CustomerCountryEnum = "JO"
	CustomerCountryKZ CustomerCountryEnum = "KZ"
	CustomerCountryKE CustomerCountryEnum = "KE"
	CustomerCountryKI CustomerCountryEnum = "KI"
	CustomerCountryKW CustomerCountryEnum = "KW"
	CustomerCountryKG CustomerCountryEnum = "KG"
	CustomerCountryLA CustomerCountryEnum = "LA"
	CustomerCountryLV CustomerCountryEnum = "LV"
	CustomerCountryLB CustomerCountryEnum = "LB"
	CustomerCountryLS CustomerCountryEnum = "LS"
	CustomerCountryLR CustomerCountryEnum = "LR"
	CustomerCountryLY CustomerCountryEnum = "LY"
	CustomerCountryLI CustomerCountryEnum = "LI"
	CustomerCountryLT CustomerCountryEnum = "LT"
	CustomerCountryLU CustomerCountryEnum = "LU"
	CustomerCountryMG CustomerCountryEnum = "MG"
	CustomerCountryMW CustomerCountryEnum = "MW"
	CustomerCountryMY CustomerCountryEnum = "MY"
	CustomerCountryMV CustomerCountryEnum = "MV"
	CustomerCountryML CustomerCountryEnum = "ML"
	CustomerCountryMT CustomerCountryEnum = "MT"
	CustomerCountryMH CustomerCountryEnum = "MH"
	CustomerCountryMQ CustomerCountryEnum = "MQ"
	CustomerCountryMR CustomerCountryEnum = "MR"
	CustomerCountryMU CustomerCountryEnum = "MU"
	CustomerCountryYT CustomerCountryEnum = "YT"
	CustomerCountryFM CustomerCountryEnum = "FM"
	CustomerCountryMC CustomerCountryEnum = "MC"
	CustomerCountryMN CustomerCountryEnum = "MN"
	CustomerCountryME CustomerCountryEnum = "ME"
	CustomerCountryMS CustomerCountryEnum = "MS"
	CustomerCountryMA CustomerCountryEnum = "MA"
	CustomerCountryMZ CustomerCountryEnum = "MZ"
	CustomerCountryMM CustomerCountryEnum = "MM"
	CustomerCountryNA CustomerCountryEnum = "NA"
	CustomerCountryNR CustomerCountryEnum = "NR"
	CustomerCountryNP CustomerCountryEnum = "NP"
	CustomerCountryNL CustomerCountryEnum = "NL"
	CustomerCountryNC CustomerCountryEnum = "NC"
	CustomerCountryNZ CustomerCountryEnum = "NZ"
	CustomerCountryNI CustomerCountryEnum = "NI"
	CustomerCountryNE CustomerCountryEnum = "NE"
	CustomerCountryNG CustomerCountryEnum = "NG"
	CustomerCountryNU CustomerCountryEnum = "NU"
	CustomerCountryNF CustomerCountryEnum = "NF"
	CustomerCountryMP CustomerCountryEnum = "MP"
	CustomerCountryNO CustomerCountryEnum = "NO"
	CustomerCountryOM CustomerCountryEnum = "OM"
	CustomerCountryPK CustomerCountryEnum = "PK"
	CustomerCountryPW CustomerCountryEnum = "PW"
	CustomerCountryPA CustomerCountryEnum = "PA"
	CustomerCountryPG CustomerCountryEnum = "PG"
	CustomerCountryPY CustomerCountryEnum = "PY"
	CustomerCountryPE CustomerCountryEnum = "PE"
	CustomerCountryPH CustomerCountryEnum = "PH"
	CustomerCountryPN CustomerCountryEnum = "PN"
	CustomerCountryPL CustomerCountryEnum = "PL"
	CustomerCountryPT CustomerCountryEnum = "PT"
	CustomerCountryPR CustomerCountryEnum = "PR"
	CustomerCountryQA CustomerCountryEnum = "QA"
	CustomerCountryKR CustomerCountryEnum = "KR"
	CustomerCountryMD CustomerCountryEnum = "MD"
	CustomerCountryRO CustomerCountryEnum = "RO"
	CustomerCountryRU CustomerCountryEnum = "RU"
	CustomerCountryRW CustomerCountryEnum = "RW"
	CustomerCountryRE CustomerCountryEnum = "RE"
	CustomerCountryBL CustomerCountryEnum = "BL"
	CustomerCountrySH CustomerCountryEnum = "SH"
	CustomerCountryKN CustomerCountryEnum = "KN"
	CustomerCountryLC CustomerCountryEnum = "LC"
	CustomerCountryMF CustomerCountryEnum = "MF"
	CustomerCountryPM CustomerCountryEnum = "PM"
	CustomerCountryVC CustomerCountryEnum = "VC"
	CustomerCountryWS CustomerCountryEnum = "WS"
	CustomerCountrySM CustomerCountryEnum = "SM"
	CustomerCountryST CustomerCountryEnum = "ST"
	CustomerCountrySA CustomerCountryEnum = "SA"
	CustomerCountrySN CustomerCountryEnum = "SN"
	CustomerCountryRS CustomerCountryEnum = "RS"
	CustomerCountrySC CustomerCountryEnum = "SC"
	CustomerCountrySL CustomerCountryEnum = "SL"
	CustomerCountrySG CustomerCountryEnum = "SG"
	CustomerCountrySX CustomerCountryEnum = "SX"
	CustomerCountrySK CustomerCountryEnum = "SK"
	CustomerCountrySI CustomerCountryEnum = "SI"
	CustomerCountrySB CustomerCountryEnum = "SB"
	CustomerCountrySO CustomerCountryEnum = "SO"
	CustomerCountryZA CustomerCountryEnum = "ZA"
	CustomerCountryGS CustomerCountryEnum = "GS"
	CustomerCountrySS CustomerCountryEnum = "SS"
	CustomerCountryES CustomerCountryEnum = "ES"
	CustomerCountryLK CustomerCountryEnum = "LK"
	CustomerCountryPS CustomerCountryEnum = "PS"
	CustomerCountrySD CustomerCountryEnum = "SD"
	CustomerCountrySR CustomerCountryEnum = "SR"
	CustomerCountrySJ CustomerCountryEnum = "SJ"
	CustomerCountrySE CustomerCountryEnum = "SE"
	CustomerCountryCH CustomerCountryEnum = "CH"
	CustomerCountrySY CustomerCountryEnum = "SY"
	CustomerCountryTW CustomerCountryEnum = "TW"
	CustomerCountryTJ CustomerCountryEnum = "TJ"
	CustomerCountryTH CustomerCountryEnum = "TH"
	CustomerCountryMK CustomerCountryEnum = "MK"
	CustomerCountryTL CustomerCountryEnum = "TL"
	CustomerCountryTG CustomerCountryEnum = "TG"
	CustomerCountryTK CustomerCountryEnum = "TK"
	CustomerCountryTO CustomerCountryEnum = "TO"
	CustomerCountryTT CustomerCountryEnum = "TT"
	CustomerCountryTN CustomerCountryEnum = "TN"
	CustomerCountryTR CustomerCountryEnum = "TR"
	CustomerCountryTM CustomerCountryEnum = "TM"
	CustomerCountryTC CustomerCountryEnum = "TC"
	CustomerCountryTV CustomerCountryEnum = "TV"
	CustomerCountryUG CustomerCountryEnum = "UG"
	CustomerCountryUA CustomerCountryEnum = "UA"
	CustomerCountryAE CustomerCountryEnum = "AE"
	CustomerCountryGB CustomerCountryEnum = "GB"
	CustomerCountryTZ CustomerCountryEnum = "TZ"
	CustomerCountryUM CustomerCountryEnum = "UM"
	CustomerCountryVI CustomerCountryEnum = "VI"
	CustomerCountryUY CustomerCountryEnum = "UY"
	CustomerCountryUZ CustomerCountryEnum = "UZ"
	CustomerCountryVU CustomerCountryEnum = "VU"
	CustomerCountryVE CustomerCountryEnum = "VE"
	CustomerCountryVN CustomerCountryEnum = "VN"
	CustomerCountryWF CustomerCountryEnum = "WF"
	CustomerCountryEH CustomerCountryEnum = "EH"
	CustomerCountryYE CustomerCountryEnum = "YE"
	CustomerCountryZM CustomerCountryEnum = "ZM"
	CustomerCountryZW CustomerCountryEnum = "ZW"
)

type CustomerCustomerTypeEnum string

const (
	CustomerCustomerTypeCustomer CustomerCustomerTypeEnum = "Customer"
	CustomerCustomerTypeFleet    CustomerCustomerTypeEnum = "Fleet"
)

type CustomerPreferredContactMethodEnum string

const (
	CustomerPreferredContactMethodSMS   CustomerPreferredContactMethodEnum = "SMS"
	CustomerPreferredContactMethodEmail CustomerPreferredContactMethodEnum = "Email"
	CustomerPreferredContactMethodAll   CustomerPreferredContactMethodEnum = "All"
)

type Customer struct {
	ID           string              `gorm:"primaryKey;not null;column:id" json:"id"`
	CreatedDate  datatypes.DateTime  `gorm:"column:createdDate;not null;column:createdDate" json:"createdDate"`
	UpdatedDate  *datatypes.DateTime `gorm:"column:updatedDate;column:updatedDate" json:"updatedDate"`
	Meta         *datatypes.JSON     `gorm:"column:meta;not null;column:meta" json:"meta,omitempty"`    // the metadata about the most recent change to the row
	Metadata     *datatypes.JSON     `gorm:"column:metadata;column:metadata" json:"metadata,omitempty"` // metadata reserved for customers to control
	CompanyID    string              `gorm:"not null;column:companyId" json:"companyId"`
	CustomFields datatypes.JSON      `gorm:"column:customFields" json:"customFields"` // custom field values

	Address1               *string                             `gorm:"column:address1" json:"address1"`
	Address2               *string                             `gorm:"column:address2" json:"address2"`
	AppointmentCount       int64                               `gorm:"not null;column:appointmentCount" json:"appointmentCount"`
	City                   *string                             `gorm:"column:city" json:"city"`
	CompanyName            *string                             `gorm:"column:companyName" json:"companyName"`
	Country                *CustomerCountryEnum                `gorm:"column:country" json:"country"`
	CustomerType           CustomerCustomerTypeEnum            `gorm:"not null;column:customerType" json:"customerType"`
	DeferredServiceCount   int64                               `gorm:"not null;column:deferredServiceCount" json:"deferredServiceCount"`
	DiscountStatus         bool                                `gorm:"not null;column:discountStatus" json:"discountStatus"`
	DotNumber              *string                             `gorm:"column:dotNumber" json:"dotNumber"`
	FirstName              *string                             `gorm:"column:firstName" json:"firstName"`
	FleetID                *string                             `gorm:"column:fleetId" json:"fleetId"`
	LaborMatrixID          *string                             `gorm:"column:laborMatrixId" json:"laborMatrixId"`
	LaborRateOverride      bool                                `gorm:"not null;column:laborRateOverride" json:"laborRateOverride"`
	LastName               *string                             `gorm:"column:lastName" json:"lastName"`
	LastTimeOrderWorked    *datatypes.DateTime                 `gorm:"column:lastTimeOrderWorked" json:"lastTimeOrderWorked"`
	MarketingOptIn         bool                                `gorm:"not null;column:marketingOptIn" json:"marketingOptIn"`
	MessageCount           int64                               `gorm:"not null;column:messageCount" json:"messageCount"`
	NormalizedFirstName    *string                             `gorm:"column:normalizedFirstName" json:"normalizedFirstName"`
	NormalizedLastName     *string                             `gorm:"column:normalizedLastName" json:"normalizedLastName"`
	Note                   string                              `gorm:"not null;column:note" json:"note"`
	PaymentTermID          string                              `gorm:"not null;column:paymentTermId" json:"paymentTermId"`
	PostalCode             *string                             `gorm:"column:postalCode" json:"postalCode"`
	PreferredContactMethod *CustomerPreferredContactMethodEnum `gorm:"column:preferredContactMethod" json:"preferredContactMethod"`
	PricingMatrixID        *string                             `gorm:"column:pricingMatrixId" json:"pricingMatrixId"`
	ReferralSourceID       *string                             `gorm:"column:referralSourceId" json:"referralSourceId"`
	State                  *string                             `gorm:"column:state" json:"state"`
	StatementCount         int64                               `gorm:"not null;column:statementCount" json:"statementCount"`
	TaxExempt              bool                                `gorm:"not null;column:taxExempt" json:"taxExempt"`
	TransactionCount       int64                               `gorm:"not null;column:transactionCount" json:"transactionCount"`
}

var _ Model = (*Customer)(nil)

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *Customer) TableName() string {
	return "customer"
}

// PrimaryKeys returns an array of the primary keys for this model
func (m *Customer) PrimaryKeys() []string {
	return []string{"id"}
}

// String returns a string representation as JSON for this model
func (m *Customer) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewCustomer returns a new model instance from an encoded buffer
func NewCustomer(buf []byte) (*Customer, error) {
	var result Customer
	err := json.Unmarshal(buf, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// NewCustomerFromChangeEvent returns a new model instance from an encoded buffer as change event
func NewCustomerFromChangeEvent(buf []byte, gzip bool) (*datatypes.ChangeEvent[Customer], error) {
	var result datatypes.ChangeEvent[Customer]
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
