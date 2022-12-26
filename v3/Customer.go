// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	mapstructure "github.com/mitchellh/mapstructure"
	"time"
)

type CustomerCountryEnum string

const (
	CustomerCountryUS CustomerCountryEnum = "US"
	CustomerCountryCA                     = "CA"
	CustomerCountryMX                     = "MX"
	CustomerCountryAF                     = "AF"
	CustomerCountryAX                     = "AX"
	CustomerCountryAL                     = "AL"
	CustomerCountryDZ                     = "DZ"
	CustomerCountryAS                     = "AS"
	CustomerCountryAD                     = "AD"
	CustomerCountryAO                     = "AO"
	CustomerCountryAI                     = "AI"
	CustomerCountryAQ                     = "AQ"
	CustomerCountryAG                     = "AG"
	CustomerCountryAR                     = "AR"
	CustomerCountryAM                     = "AM"
	CustomerCountryAW                     = "AW"
	CustomerCountryAU                     = "AU"
	CustomerCountryAT                     = "AT"
	CustomerCountryAZ                     = "AZ"
	CustomerCountryBS                     = "BS"
	CustomerCountryBH                     = "BH"
	CustomerCountryBD                     = "BD"
	CustomerCountryBB                     = "BB"
	CustomerCountryBY                     = "BY"
	CustomerCountryBE                     = "BE"
	CustomerCountryBZ                     = "BZ"
	CustomerCountryBJ                     = "BJ"
	CustomerCountryBM                     = "BM"
	CustomerCountryBT                     = "BT"
	CustomerCountryBO                     = "BO"
	CustomerCountryBQ                     = "BQ"
	CustomerCountryBA                     = "BA"
	CustomerCountryBW                     = "BW"
	CustomerCountryBV                     = "BV"
	CustomerCountryBR                     = "BR"
	CustomerCountryIO                     = "IO"
	CustomerCountryVG                     = "VG"
	CustomerCountryBN                     = "BN"
	CustomerCountryBG                     = "BG"
	CustomerCountryBF                     = "BF"
	CustomerCountryBI                     = "BI"
	CustomerCountryCV                     = "CV"
	CustomerCountryKH                     = "KH"
	CustomerCountryCM                     = "CM"
	CustomerCountryKY                     = "KY"
	CustomerCountryCF                     = "CF"
	CustomerCountryTD                     = "TD"
	CustomerCountryCL                     = "CL"
	CustomerCountryCN                     = "CN"
	CustomerCountryHK                     = "HK"
	CustomerCountryMO                     = "MO"
	CustomerCountryCX                     = "CX"
	CustomerCountryCC                     = "CC"
	CustomerCountryCO                     = "CO"
	CustomerCountryKM                     = "KM"
	CustomerCountryCG                     = "CG"
	CustomerCountryCK                     = "CK"
	CustomerCountryCR                     = "CR"
	CustomerCountryHR                     = "HR"
	CustomerCountryCU                     = "CU"
	CustomerCountryCW                     = "CW"
	CustomerCountryCY                     = "CY"
	CustomerCountryCZ                     = "CZ"
	CustomerCountryCI                     = "CI"
	CustomerCountryKP                     = "KP"
	CustomerCountryCD                     = "CD"
	CustomerCountryDK                     = "DK"
	CustomerCountryDJ                     = "DJ"
	CustomerCountryDM                     = "DM"
	CustomerCountryDO                     = "DO"
	CustomerCountryEC                     = "EC"
	CustomerCountryEG                     = "EG"
	CustomerCountrySV                     = "SV"
	CustomerCountryGQ                     = "GQ"
	CustomerCountryER                     = "ER"
	CustomerCountryEE                     = "EE"
	CustomerCountrySZ                     = "SZ"
	CustomerCountryET                     = "ET"
	CustomerCountryFK                     = "FK"
	CustomerCountryFO                     = "FO"
	CustomerCountryFJ                     = "FJ"
	CustomerCountryFI                     = "FI"
	CustomerCountryFR                     = "FR"
	CustomerCountryGF                     = "GF"
	CustomerCountryPF                     = "PF"
	CustomerCountryTF                     = "TF"
	CustomerCountryGA                     = "GA"
	CustomerCountryGM                     = "GM"
	CustomerCountryGE                     = "GE"
	CustomerCountryDE                     = "DE"
	CustomerCountryGH                     = "GH"
	CustomerCountryGI                     = "GI"
	CustomerCountryGR                     = "GR"
	CustomerCountryGL                     = "GL"
	CustomerCountryGD                     = "GD"
	CustomerCountryGP                     = "GP"
	CustomerCountryGU                     = "GU"
	CustomerCountryGT                     = "GT"
	CustomerCountryGG                     = "GG"
	CustomerCountryGN                     = "GN"
	CustomerCountryGW                     = "GW"
	CustomerCountryGY                     = "GY"
	CustomerCountryHT                     = "HT"
	CustomerCountryHM                     = "HM"
	CustomerCountryVA                     = "VA"
	CustomerCountryHN                     = "HN"
	CustomerCountryHU                     = "HU"
	CustomerCountryIS                     = "IS"
	CustomerCountryIN                     = "IN"
	CustomerCountryID                     = "ID"
	CustomerCountryIR                     = "IR"
	CustomerCountryIQ                     = "IQ"
	CustomerCountryIE                     = "IE"
	CustomerCountryIM                     = "IM"
	CustomerCountryIL                     = "IL"
	CustomerCountryIT                     = "IT"
	CustomerCountryJM                     = "JM"
	CustomerCountryJP                     = "JP"
	CustomerCountryJE                     = "JE"
	CustomerCountryJO                     = "JO"
	CustomerCountryKZ                     = "KZ"
	CustomerCountryKE                     = "KE"
	CustomerCountryKI                     = "KI"
	CustomerCountryKW                     = "KW"
	CustomerCountryKG                     = "KG"
	CustomerCountryLA                     = "LA"
	CustomerCountryLV                     = "LV"
	CustomerCountryLB                     = "LB"
	CustomerCountryLS                     = "LS"
	CustomerCountryLR                     = "LR"
	CustomerCountryLY                     = "LY"
	CustomerCountryLI                     = "LI"
	CustomerCountryLT                     = "LT"
	CustomerCountryLU                     = "LU"
	CustomerCountryMG                     = "MG"
	CustomerCountryMW                     = "MW"
	CustomerCountryMY                     = "MY"
	CustomerCountryMV                     = "MV"
	CustomerCountryML                     = "ML"
	CustomerCountryMT                     = "MT"
	CustomerCountryMH                     = "MH"
	CustomerCountryMQ                     = "MQ"
	CustomerCountryMR                     = "MR"
	CustomerCountryMU                     = "MU"
	CustomerCountryYT                     = "YT"
	CustomerCountryFM                     = "FM"
	CustomerCountryMC                     = "MC"
	CustomerCountryMN                     = "MN"
	CustomerCountryME                     = "ME"
	CustomerCountryMS                     = "MS"
	CustomerCountryMA                     = "MA"
	CustomerCountryMZ                     = "MZ"
	CustomerCountryMM                     = "MM"
	CustomerCountryNA                     = "NA"
	CustomerCountryNR                     = "NR"
	CustomerCountryNP                     = "NP"
	CustomerCountryNL                     = "NL"
	CustomerCountryNC                     = "NC"
	CustomerCountryNZ                     = "NZ"
	CustomerCountryNI                     = "NI"
	CustomerCountryNE                     = "NE"
	CustomerCountryNG                     = "NG"
	CustomerCountryNU                     = "NU"
	CustomerCountryNF                     = "NF"
	CustomerCountryMP                     = "MP"
	CustomerCountryNO                     = "NO"
	CustomerCountryOM                     = "OM"
	CustomerCountryPK                     = "PK"
	CustomerCountryPW                     = "PW"
	CustomerCountryPA                     = "PA"
	CustomerCountryPG                     = "PG"
	CustomerCountryPY                     = "PY"
	CustomerCountryPE                     = "PE"
	CustomerCountryPH                     = "PH"
	CustomerCountryPN                     = "PN"
	CustomerCountryPL                     = "PL"
	CustomerCountryPT                     = "PT"
	CustomerCountryPR                     = "PR"
	CustomerCountryQA                     = "QA"
	CustomerCountryKR                     = "KR"
	CustomerCountryMD                     = "MD"
	CustomerCountryRO                     = "RO"
	CustomerCountryRU                     = "RU"
	CustomerCountryRW                     = "RW"
	CustomerCountryRE                     = "RE"
	CustomerCountryBL                     = "BL"
	CustomerCountrySH                     = "SH"
	CustomerCountryKN                     = "KN"
	CustomerCountryLC                     = "LC"
	CustomerCountryMF                     = "MF"
	CustomerCountryPM                     = "PM"
	CustomerCountryVC                     = "VC"
	CustomerCountryWS                     = "WS"
	CustomerCountrySM                     = "SM"
	CustomerCountryST                     = "ST"
	CustomerCountrySA                     = "SA"
	CustomerCountrySN                     = "SN"
	CustomerCountryRS                     = "RS"
	CustomerCountrySC                     = "SC"
	CustomerCountrySL                     = "SL"
	CustomerCountrySG                     = "SG"
	CustomerCountrySX                     = "SX"
	CustomerCountrySK                     = "SK"
	CustomerCountrySI                     = "SI"
	CustomerCountrySB                     = "SB"
	CustomerCountrySO                     = "SO"
	CustomerCountryZA                     = "ZA"
	CustomerCountryGS                     = "GS"
	CustomerCountrySS                     = "SS"
	CustomerCountryES                     = "ES"
	CustomerCountryLK                     = "LK"
	CustomerCountryPS                     = "PS"
	CustomerCountrySD                     = "SD"
	CustomerCountrySR                     = "SR"
	CustomerCountrySJ                     = "SJ"
	CustomerCountrySE                     = "SE"
	CustomerCountryCH                     = "CH"
	CustomerCountrySY                     = "SY"
	CustomerCountryTW                     = "TW"
	CustomerCountryTJ                     = "TJ"
	CustomerCountryTH                     = "TH"
	CustomerCountryMK                     = "MK"
	CustomerCountryTL                     = "TL"
	CustomerCountryTG                     = "TG"
	CustomerCountryTK                     = "TK"
	CustomerCountryTO                     = "TO"
	CustomerCountryTT                     = "TT"
	CustomerCountryTN                     = "TN"
	CustomerCountryTR                     = "TR"
	CustomerCountryTM                     = "TM"
	CustomerCountryTC                     = "TC"
	CustomerCountryTV                     = "TV"
	CustomerCountryUG                     = "UG"
	CustomerCountryUA                     = "UA"
	CustomerCountryAE                     = "AE"
	CustomerCountryGB                     = "GB"
	CustomerCountryTZ                     = "TZ"
	CustomerCountryUM                     = "UM"
	CustomerCountryVI                     = "VI"
	CustomerCountryUY                     = "UY"
	CustomerCountryUZ                     = "UZ"
	CustomerCountryVU                     = "VU"
	CustomerCountryVE                     = "VE"
	CustomerCountryVN                     = "VN"
	CustomerCountryWF                     = "WF"
	CustomerCountryEH                     = "EH"
	CustomerCountryYE                     = "YE"
	CustomerCountryZM                     = "ZM"
	CustomerCountryZW                     = "ZW"
)

type CustomerCustomerTypeEnum string

const (
	CustomerCustomerTypeCustomer CustomerCustomerTypeEnum = "Customer"
	CustomerCustomerTypeFleet                             = "Fleet"
)

type CustomerPreferredContactMethodEnum string

const (
	CustomerPreferredContactMethodSMS   CustomerPreferredContactMethodEnum = "SMS"
	CustomerPreferredContactMethodEmail                                    = "Email"
	CustomerPreferredContactMethodAll                                      = "All"
)

type Customer struct {
	ID           string     `gorm:"primaryKey;not null;column:id" json:"id"`
	CreatedDate  time.Time  `gorm:"column:createdDate;not null;column:createdDate" json:"createdDate"`
	UpdatedDate  *time.Time `gorm:"column:updatedDate;column:updatedDate" json:"updatedDate"`
	Meta         *Meta      `gorm:"type:json;embedded;serializer:json;column:meta;not null;column:meta" json:"meta,omitempty"` // the metadata about the most recent change to the row
	Metadata     any        `gorm:"type:json;serializer:json;column:metadata" json:"metadata,omitempty"`                       // metadata reserved for customers to control
	CompanyID    string     `gorm:"not null;column:companyId" json:"companyId"`
	CustomFields any        `gorm:"type:json;serializer:json;column:customFields" json:"customFields"` // custom field values

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
	LastTimeOrderWorked    *time.Time                          `gorm:"column:lastTimeOrderWorked" json:"lastTimeOrderWorked"`
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

func (m *Customer) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewCustomer returns a new model instance from a json key/value map
func NewCustomer(kv map[string]any) (*Customer, error) {
	var result Customer
	err := mapstructure.Decode(kv, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
