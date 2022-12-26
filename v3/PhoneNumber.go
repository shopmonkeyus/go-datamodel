// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	codec "github.com/hashicorp/go-msgpack/v2/codec"
	"time"
)

type PhoneNumberCountryEnum string

const (
	PhoneNumberCountryUS PhoneNumberCountryEnum = "US"
	PhoneNumberCountryCA                        = "CA"
	PhoneNumberCountryMX                        = "MX"
	PhoneNumberCountryAF                        = "AF"
	PhoneNumberCountryAX                        = "AX"
	PhoneNumberCountryAL                        = "AL"
	PhoneNumberCountryDZ                        = "DZ"
	PhoneNumberCountryAS                        = "AS"
	PhoneNumberCountryAD                        = "AD"
	PhoneNumberCountryAO                        = "AO"
	PhoneNumberCountryAI                        = "AI"
	PhoneNumberCountryAQ                        = "AQ"
	PhoneNumberCountryAG                        = "AG"
	PhoneNumberCountryAR                        = "AR"
	PhoneNumberCountryAM                        = "AM"
	PhoneNumberCountryAW                        = "AW"
	PhoneNumberCountryAU                        = "AU"
	PhoneNumberCountryAT                        = "AT"
	PhoneNumberCountryAZ                        = "AZ"
	PhoneNumberCountryBS                        = "BS"
	PhoneNumberCountryBH                        = "BH"
	PhoneNumberCountryBD                        = "BD"
	PhoneNumberCountryBB                        = "BB"
	PhoneNumberCountryBY                        = "BY"
	PhoneNumberCountryBE                        = "BE"
	PhoneNumberCountryBZ                        = "BZ"
	PhoneNumberCountryBJ                        = "BJ"
	PhoneNumberCountryBM                        = "BM"
	PhoneNumberCountryBT                        = "BT"
	PhoneNumberCountryBO                        = "BO"
	PhoneNumberCountryBQ                        = "BQ"
	PhoneNumberCountryBA                        = "BA"
	PhoneNumberCountryBW                        = "BW"
	PhoneNumberCountryBV                        = "BV"
	PhoneNumberCountryBR                        = "BR"
	PhoneNumberCountryIO                        = "IO"
	PhoneNumberCountryVG                        = "VG"
	PhoneNumberCountryBN                        = "BN"
	PhoneNumberCountryBG                        = "BG"
	PhoneNumberCountryBF                        = "BF"
	PhoneNumberCountryBI                        = "BI"
	PhoneNumberCountryCV                        = "CV"
	PhoneNumberCountryKH                        = "KH"
	PhoneNumberCountryCM                        = "CM"
	PhoneNumberCountryKY                        = "KY"
	PhoneNumberCountryCF                        = "CF"
	PhoneNumberCountryTD                        = "TD"
	PhoneNumberCountryCL                        = "CL"
	PhoneNumberCountryCN                        = "CN"
	PhoneNumberCountryHK                        = "HK"
	PhoneNumberCountryMO                        = "MO"
	PhoneNumberCountryCX                        = "CX"
	PhoneNumberCountryCC                        = "CC"
	PhoneNumberCountryCO                        = "CO"
	PhoneNumberCountryKM                        = "KM"
	PhoneNumberCountryCG                        = "CG"
	PhoneNumberCountryCK                        = "CK"
	PhoneNumberCountryCR                        = "CR"
	PhoneNumberCountryHR                        = "HR"
	PhoneNumberCountryCU                        = "CU"
	PhoneNumberCountryCW                        = "CW"
	PhoneNumberCountryCY                        = "CY"
	PhoneNumberCountryCZ                        = "CZ"
	PhoneNumberCountryCI                        = "CI"
	PhoneNumberCountryKP                        = "KP"
	PhoneNumberCountryCD                        = "CD"
	PhoneNumberCountryDK                        = "DK"
	PhoneNumberCountryDJ                        = "DJ"
	PhoneNumberCountryDM                        = "DM"
	PhoneNumberCountryDO                        = "DO"
	PhoneNumberCountryEC                        = "EC"
	PhoneNumberCountryEG                        = "EG"
	PhoneNumberCountrySV                        = "SV"
	PhoneNumberCountryGQ                        = "GQ"
	PhoneNumberCountryER                        = "ER"
	PhoneNumberCountryEE                        = "EE"
	PhoneNumberCountrySZ                        = "SZ"
	PhoneNumberCountryET                        = "ET"
	PhoneNumberCountryFK                        = "FK"
	PhoneNumberCountryFO                        = "FO"
	PhoneNumberCountryFJ                        = "FJ"
	PhoneNumberCountryFI                        = "FI"
	PhoneNumberCountryFR                        = "FR"
	PhoneNumberCountryGF                        = "GF"
	PhoneNumberCountryPF                        = "PF"
	PhoneNumberCountryTF                        = "TF"
	PhoneNumberCountryGA                        = "GA"
	PhoneNumberCountryGM                        = "GM"
	PhoneNumberCountryGE                        = "GE"
	PhoneNumberCountryDE                        = "DE"
	PhoneNumberCountryGH                        = "GH"
	PhoneNumberCountryGI                        = "GI"
	PhoneNumberCountryGR                        = "GR"
	PhoneNumberCountryGL                        = "GL"
	PhoneNumberCountryGD                        = "GD"
	PhoneNumberCountryGP                        = "GP"
	PhoneNumberCountryGU                        = "GU"
	PhoneNumberCountryGT                        = "GT"
	PhoneNumberCountryGG                        = "GG"
	PhoneNumberCountryGN                        = "GN"
	PhoneNumberCountryGW                        = "GW"
	PhoneNumberCountryGY                        = "GY"
	PhoneNumberCountryHT                        = "HT"
	PhoneNumberCountryHM                        = "HM"
	PhoneNumberCountryVA                        = "VA"
	PhoneNumberCountryHN                        = "HN"
	PhoneNumberCountryHU                        = "HU"
	PhoneNumberCountryIS                        = "IS"
	PhoneNumberCountryIN                        = "IN"
	PhoneNumberCountryID                        = "ID"
	PhoneNumberCountryIR                        = "IR"
	PhoneNumberCountryIQ                        = "IQ"
	PhoneNumberCountryIE                        = "IE"
	PhoneNumberCountryIM                        = "IM"
	PhoneNumberCountryIL                        = "IL"
	PhoneNumberCountryIT                        = "IT"
	PhoneNumberCountryJM                        = "JM"
	PhoneNumberCountryJP                        = "JP"
	PhoneNumberCountryJE                        = "JE"
	PhoneNumberCountryJO                        = "JO"
	PhoneNumberCountryKZ                        = "KZ"
	PhoneNumberCountryKE                        = "KE"
	PhoneNumberCountryKI                        = "KI"
	PhoneNumberCountryKW                        = "KW"
	PhoneNumberCountryKG                        = "KG"
	PhoneNumberCountryLA                        = "LA"
	PhoneNumberCountryLV                        = "LV"
	PhoneNumberCountryLB                        = "LB"
	PhoneNumberCountryLS                        = "LS"
	PhoneNumberCountryLR                        = "LR"
	PhoneNumberCountryLY                        = "LY"
	PhoneNumberCountryLI                        = "LI"
	PhoneNumberCountryLT                        = "LT"
	PhoneNumberCountryLU                        = "LU"
	PhoneNumberCountryMG                        = "MG"
	PhoneNumberCountryMW                        = "MW"
	PhoneNumberCountryMY                        = "MY"
	PhoneNumberCountryMV                        = "MV"
	PhoneNumberCountryML                        = "ML"
	PhoneNumberCountryMT                        = "MT"
	PhoneNumberCountryMH                        = "MH"
	PhoneNumberCountryMQ                        = "MQ"
	PhoneNumberCountryMR                        = "MR"
	PhoneNumberCountryMU                        = "MU"
	PhoneNumberCountryYT                        = "YT"
	PhoneNumberCountryFM                        = "FM"
	PhoneNumberCountryMC                        = "MC"
	PhoneNumberCountryMN                        = "MN"
	PhoneNumberCountryME                        = "ME"
	PhoneNumberCountryMS                        = "MS"
	PhoneNumberCountryMA                        = "MA"
	PhoneNumberCountryMZ                        = "MZ"
	PhoneNumberCountryMM                        = "MM"
	PhoneNumberCountryNA                        = "NA"
	PhoneNumberCountryNR                        = "NR"
	PhoneNumberCountryNP                        = "NP"
	PhoneNumberCountryNL                        = "NL"
	PhoneNumberCountryNC                        = "NC"
	PhoneNumberCountryNZ                        = "NZ"
	PhoneNumberCountryNI                        = "NI"
	PhoneNumberCountryNE                        = "NE"
	PhoneNumberCountryNG                        = "NG"
	PhoneNumberCountryNU                        = "NU"
	PhoneNumberCountryNF                        = "NF"
	PhoneNumberCountryMP                        = "MP"
	PhoneNumberCountryNO                        = "NO"
	PhoneNumberCountryOM                        = "OM"
	PhoneNumberCountryPK                        = "PK"
	PhoneNumberCountryPW                        = "PW"
	PhoneNumberCountryPA                        = "PA"
	PhoneNumberCountryPG                        = "PG"
	PhoneNumberCountryPY                        = "PY"
	PhoneNumberCountryPE                        = "PE"
	PhoneNumberCountryPH                        = "PH"
	PhoneNumberCountryPN                        = "PN"
	PhoneNumberCountryPL                        = "PL"
	PhoneNumberCountryPT                        = "PT"
	PhoneNumberCountryPR                        = "PR"
	PhoneNumberCountryQA                        = "QA"
	PhoneNumberCountryKR                        = "KR"
	PhoneNumberCountryMD                        = "MD"
	PhoneNumberCountryRO                        = "RO"
	PhoneNumberCountryRU                        = "RU"
	PhoneNumberCountryRW                        = "RW"
	PhoneNumberCountryRE                        = "RE"
	PhoneNumberCountryBL                        = "BL"
	PhoneNumberCountrySH                        = "SH"
	PhoneNumberCountryKN                        = "KN"
	PhoneNumberCountryLC                        = "LC"
	PhoneNumberCountryMF                        = "MF"
	PhoneNumberCountryPM                        = "PM"
	PhoneNumberCountryVC                        = "VC"
	PhoneNumberCountryWS                        = "WS"
	PhoneNumberCountrySM                        = "SM"
	PhoneNumberCountryST                        = "ST"
	PhoneNumberCountrySA                        = "SA"
	PhoneNumberCountrySN                        = "SN"
	PhoneNumberCountryRS                        = "RS"
	PhoneNumberCountrySC                        = "SC"
	PhoneNumberCountrySL                        = "SL"
	PhoneNumberCountrySG                        = "SG"
	PhoneNumberCountrySX                        = "SX"
	PhoneNumberCountrySK                        = "SK"
	PhoneNumberCountrySI                        = "SI"
	PhoneNumberCountrySB                        = "SB"
	PhoneNumberCountrySO                        = "SO"
	PhoneNumberCountryZA                        = "ZA"
	PhoneNumberCountryGS                        = "GS"
	PhoneNumberCountrySS                        = "SS"
	PhoneNumberCountryES                        = "ES"
	PhoneNumberCountryLK                        = "LK"
	PhoneNumberCountryPS                        = "PS"
	PhoneNumberCountrySD                        = "SD"
	PhoneNumberCountrySR                        = "SR"
	PhoneNumberCountrySJ                        = "SJ"
	PhoneNumberCountrySE                        = "SE"
	PhoneNumberCountryCH                        = "CH"
	PhoneNumberCountrySY                        = "SY"
	PhoneNumberCountryTW                        = "TW"
	PhoneNumberCountryTJ                        = "TJ"
	PhoneNumberCountryTH                        = "TH"
	PhoneNumberCountryMK                        = "MK"
	PhoneNumberCountryTL                        = "TL"
	PhoneNumberCountryTG                        = "TG"
	PhoneNumberCountryTK                        = "TK"
	PhoneNumberCountryTO                        = "TO"
	PhoneNumberCountryTT                        = "TT"
	PhoneNumberCountryTN                        = "TN"
	PhoneNumberCountryTR                        = "TR"
	PhoneNumberCountryTM                        = "TM"
	PhoneNumberCountryTC                        = "TC"
	PhoneNumberCountryTV                        = "TV"
	PhoneNumberCountryUG                        = "UG"
	PhoneNumberCountryUA                        = "UA"
	PhoneNumberCountryAE                        = "AE"
	PhoneNumberCountryGB                        = "GB"
	PhoneNumberCountryTZ                        = "TZ"
	PhoneNumberCountryUM                        = "UM"
	PhoneNumberCountryVI                        = "VI"
	PhoneNumberCountryUY                        = "UY"
	PhoneNumberCountryUZ                        = "UZ"
	PhoneNumberCountryVU                        = "VU"
	PhoneNumberCountryVE                        = "VE"
	PhoneNumberCountryVN                        = "VN"
	PhoneNumberCountryWF                        = "WF"
	PhoneNumberCountryEH                        = "EH"
	PhoneNumberCountryYE                        = "YE"
	PhoneNumberCountryZM                        = "ZM"
	PhoneNumberCountryZW                        = "ZW"
)

type PhoneNumberTypeEnum string

const (
	PhoneNumberTypeMobile   PhoneNumberTypeEnum = "Mobile"
	PhoneNumberTypeLandline                     = "Landline"
	PhoneNumberTypeVoIP                         = "VoIP"
	PhoneNumberTypeOther                        = "Other"
)

type PhoneNumberUserDefinedTypeEnum string

const (
	PhoneNumberUserDefinedTypeMobile   PhoneNumberUserDefinedTypeEnum = "Mobile"
	PhoneNumberUserDefinedTypeWork                                    = "Work"
	PhoneNumberUserDefinedTypeHome                                    = "Home"
	PhoneNumberUserDefinedTypeOffice                                  = "Office"
	PhoneNumberUserDefinedTypeFax                                     = "Fax"
	PhoneNumberUserDefinedTypeVoIP                                    = "VoIP"
	PhoneNumberUserDefinedTypeLandline                                = "Landline"
	PhoneNumberUserDefinedTypeOther                                   = "Other"
)

type PhoneNumber struct {
	ID          string     `gorm:"primaryKey;not null;column:id" json:"id"`
	CreatedDate time.Time  `gorm:"column:createdDate;not null;column:createdDate" json:"createdDate"`
	UpdatedDate *time.Time `gorm:"column:updatedDate;column:updatedDate" json:"updatedDate"`
	Meta        *Meta      `gorm:"type:json;embedded;serializer:json;column:meta;not null;column:meta" json:"meta,omitempty"` // the metadata about the most recent change to the row
	Metadata    any        `gorm:"type:json;serializer:json;column:metadata" json:"metadata,omitempty"`                       // metadata reserved for customers to control
	CompanyID   string     `gorm:"not null;column:companyId" json:"companyId"`
	LocationID  string     `gorm:"not null;column:locationId" json:"locationId"`

	CarrierName       *string                         `gorm:"column:carrierName" json:"carrierName"`
	Country           *PhoneNumberCountryEnum         `gorm:"column:country" json:"country"`
	CustomerID        string                          `gorm:"not null;column:customerId" json:"customerId"`
	Extension         *string                         `gorm:"column:extension" json:"extension"`
	LastVerifiedDate  *time.Time                      `gorm:"column:lastVerifiedDate" json:"lastVerifiedDate"`
	MobileCountryCode *string                         `gorm:"column:mobileCountryCode" json:"mobileCountryCode"` // if a mobile number, the mobile country code
	MobileNetworkCode *string                         `gorm:"column:mobileNetworkCode" json:"mobileNetworkCode"` // if a mobile number, the mobile network code
	Number            string                          `gorm:"not null;column:number" json:"number"`
	OptIn             bool                            `gorm:"not null;column:optIn" json:"optIn"`
	OptInVerifiedDate *time.Time                      `gorm:"column:optInVerifiedDate" json:"optInVerifiedDate"`
	Primary           bool                            `gorm:"not null;column:primary" json:"primary"`
	Type              *PhoneNumberTypeEnum            `gorm:"column:type" json:"type"`
	UserDefinedType   *PhoneNumberUserDefinedTypeEnum `gorm:"column:userDefinedType" json:"userDefinedType"`
}

var _ Model = (*PhoneNumber)(nil)

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *PhoneNumber) TableName() string {
	return "phone_number"
}

func (m *PhoneNumber) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewPhoneNumber returns a new model instance from an encoded buffer
func NewPhoneNumber(buf []byte, enctype EncodingType) (*PhoneNumber, error) {
	var result PhoneNumber
	var handle codec.Handle
	if enctype == JSONEncoding {
		handle = &jsonHandle
	} else {
		handle = &msgpackHandle
	}
	dec := codec.NewDecoderBytes(buf, handle)
	err := dec.Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
