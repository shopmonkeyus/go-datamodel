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
	Address1               *string                             `json:"address1"`
	AppointmentCount       int64                               `gorm:"not null" json:"appointmentCount"`
	ID                     string                              `gorm:"primaryKey;not null" json:"id"`
	Address2               *string                             `json:"address2"`
	City                   *string                             `json:"city"`
	CompanyId              string                              `gorm:"not null" json:"companyId"`
	CompanyName            *string                             `json:"companyName"`
	Country                *CustomerCountryEnum                `json:"country"`
	CreatedDate            time.Time                           `gorm:"column:createdDate;not null" json:"createdDate"`
	CustomFields           any                                 `gorm:"type:json" json:"customFields"` // custom field values
	CustomerType           CustomerCustomerTypeEnum            `gorm:"not null" json:"customerType"`
	DeferredServiceCount   int64                               `gorm:"not null" json:"deferredServiceCount"`
	DiscountStatus         bool                                `gorm:"not null" json:"discountStatus"`
	DotNumber              *string                             `json:"dotNumber"`
	FirstName              *string                             `json:"firstName"`
	FleetId                *string                             `json:"fleetId"`
	LaborMatrixId          *string                             `json:"laborMatrixId"`
	LaborRateOverride      bool                                `gorm:"not null" json:"laborRateOverride"`
	LastName               *string                             `json:"lastName"`
	LastTimeOrderWorked    *time.Time                          `json:"lastTimeOrderWorked"`
	MarketingOptIn         bool                                `gorm:"not null" json:"marketingOptIn"`
	MessageCount           int64                               `gorm:"not null" json:"messageCount"`
	Meta                   *Meta                               `gorm:"type:json;embedded;column:meta;not null" json:"meta,omitempty"` // the metadata about the most recent change to the row
	Metadata               any                                 `gorm:"type:json" json:"metadata,omitempty"`                           // metadata reserved for customers to control
	NormalizedFirstName    *string                             `json:"normalizedFirstName"`
	NormalizedLastName     *string                             `json:"normalizedLastName"`
	Note                   string                              `gorm:"not null" json:"note"`
	PaymentTermId          string                              `gorm:"not null" json:"paymentTermId"`
	PostalCode             *string                             `json:"postalCode"`
	PreferredContactMethod *CustomerPreferredContactMethodEnum `json:"preferredContactMethod"`
	PricingMatrixId        *string                             `json:"pricingMatrixId"`
	ReferralSourceId       *string                             `json:"referralSourceId"`
	State                  *string                             `json:"state"`
	StatementCount         int64                               `gorm:"not null" json:"statementCount"`
	TaxExempt              bool                                `gorm:"not null" json:"taxExempt"`
	TransactionCount       int64                               `gorm:"not null" json:"transactionCount"`
	UpdatedDate            *time.Time                          `gorm:"column:updatedDate" json:"updatedDate"`
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
