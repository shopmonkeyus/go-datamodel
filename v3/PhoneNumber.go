// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	datatypes "github.com/shopmonkeyus/go-datamodel/datatypes"
)

type PhoneNumberCountryEnum string

const (
	PhoneNumberCountryUS PhoneNumberCountryEnum = "US"
	PhoneNumberCountryCA PhoneNumberCountryEnum = "CA"
	PhoneNumberCountryMX PhoneNumberCountryEnum = "MX"
	PhoneNumberCountryAF PhoneNumberCountryEnum = "AF"
	PhoneNumberCountryAX PhoneNumberCountryEnum = "AX"
	PhoneNumberCountryAL PhoneNumberCountryEnum = "AL"
	PhoneNumberCountryDZ PhoneNumberCountryEnum = "DZ"
	PhoneNumberCountryAS PhoneNumberCountryEnum = "AS"
	PhoneNumberCountryAD PhoneNumberCountryEnum = "AD"
	PhoneNumberCountryAO PhoneNumberCountryEnum = "AO"
	PhoneNumberCountryAI PhoneNumberCountryEnum = "AI"
	PhoneNumberCountryAQ PhoneNumberCountryEnum = "AQ"
	PhoneNumberCountryAG PhoneNumberCountryEnum = "AG"
	PhoneNumberCountryAR PhoneNumberCountryEnum = "AR"
	PhoneNumberCountryAM PhoneNumberCountryEnum = "AM"
	PhoneNumberCountryAW PhoneNumberCountryEnum = "AW"
	PhoneNumberCountryAU PhoneNumberCountryEnum = "AU"
	PhoneNumberCountryAT PhoneNumberCountryEnum = "AT"
	PhoneNumberCountryAZ PhoneNumberCountryEnum = "AZ"
	PhoneNumberCountryBS PhoneNumberCountryEnum = "BS"
	PhoneNumberCountryBH PhoneNumberCountryEnum = "BH"
	PhoneNumberCountryBD PhoneNumberCountryEnum = "BD"
	PhoneNumberCountryBB PhoneNumberCountryEnum = "BB"
	PhoneNumberCountryBY PhoneNumberCountryEnum = "BY"
	PhoneNumberCountryBE PhoneNumberCountryEnum = "BE"
	PhoneNumberCountryBZ PhoneNumberCountryEnum = "BZ"
	PhoneNumberCountryBJ PhoneNumberCountryEnum = "BJ"
	PhoneNumberCountryBM PhoneNumberCountryEnum = "BM"
	PhoneNumberCountryBT PhoneNumberCountryEnum = "BT"
	PhoneNumberCountryBO PhoneNumberCountryEnum = "BO"
	PhoneNumberCountryBQ PhoneNumberCountryEnum = "BQ"
	PhoneNumberCountryBA PhoneNumberCountryEnum = "BA"
	PhoneNumberCountryBW PhoneNumberCountryEnum = "BW"
	PhoneNumberCountryBV PhoneNumberCountryEnum = "BV"
	PhoneNumberCountryBR PhoneNumberCountryEnum = "BR"
	PhoneNumberCountryIO PhoneNumberCountryEnum = "IO"
	PhoneNumberCountryVG PhoneNumberCountryEnum = "VG"
	PhoneNumberCountryBN PhoneNumberCountryEnum = "BN"
	PhoneNumberCountryBG PhoneNumberCountryEnum = "BG"
	PhoneNumberCountryBF PhoneNumberCountryEnum = "BF"
	PhoneNumberCountryBI PhoneNumberCountryEnum = "BI"
	PhoneNumberCountryCV PhoneNumberCountryEnum = "CV"
	PhoneNumberCountryKH PhoneNumberCountryEnum = "KH"
	PhoneNumberCountryCM PhoneNumberCountryEnum = "CM"
	PhoneNumberCountryKY PhoneNumberCountryEnum = "KY"
	PhoneNumberCountryCF PhoneNumberCountryEnum = "CF"
	PhoneNumberCountryTD PhoneNumberCountryEnum = "TD"
	PhoneNumberCountryCL PhoneNumberCountryEnum = "CL"
	PhoneNumberCountryCN PhoneNumberCountryEnum = "CN"
	PhoneNumberCountryHK PhoneNumberCountryEnum = "HK"
	PhoneNumberCountryMO PhoneNumberCountryEnum = "MO"
	PhoneNumberCountryCX PhoneNumberCountryEnum = "CX"
	PhoneNumberCountryCC PhoneNumberCountryEnum = "CC"
	PhoneNumberCountryCO PhoneNumberCountryEnum = "CO"
	PhoneNumberCountryKM PhoneNumberCountryEnum = "KM"
	PhoneNumberCountryCG PhoneNumberCountryEnum = "CG"
	PhoneNumberCountryCK PhoneNumberCountryEnum = "CK"
	PhoneNumberCountryCR PhoneNumberCountryEnum = "CR"
	PhoneNumberCountryHR PhoneNumberCountryEnum = "HR"
	PhoneNumberCountryCU PhoneNumberCountryEnum = "CU"
	PhoneNumberCountryCW PhoneNumberCountryEnum = "CW"
	PhoneNumberCountryCY PhoneNumberCountryEnum = "CY"
	PhoneNumberCountryCZ PhoneNumberCountryEnum = "CZ"
	PhoneNumberCountryCI PhoneNumberCountryEnum = "CI"
	PhoneNumberCountryKP PhoneNumberCountryEnum = "KP"
	PhoneNumberCountryCD PhoneNumberCountryEnum = "CD"
	PhoneNumberCountryDK PhoneNumberCountryEnum = "DK"
	PhoneNumberCountryDJ PhoneNumberCountryEnum = "DJ"
	PhoneNumberCountryDM PhoneNumberCountryEnum = "DM"
	PhoneNumberCountryDO PhoneNumberCountryEnum = "DO"
	PhoneNumberCountryEC PhoneNumberCountryEnum = "EC"
	PhoneNumberCountryEG PhoneNumberCountryEnum = "EG"
	PhoneNumberCountrySV PhoneNumberCountryEnum = "SV"
	PhoneNumberCountryGQ PhoneNumberCountryEnum = "GQ"
	PhoneNumberCountryER PhoneNumberCountryEnum = "ER"
	PhoneNumberCountryEE PhoneNumberCountryEnum = "EE"
	PhoneNumberCountrySZ PhoneNumberCountryEnum = "SZ"
	PhoneNumberCountryET PhoneNumberCountryEnum = "ET"
	PhoneNumberCountryFK PhoneNumberCountryEnum = "FK"
	PhoneNumberCountryFO PhoneNumberCountryEnum = "FO"
	PhoneNumberCountryFJ PhoneNumberCountryEnum = "FJ"
	PhoneNumberCountryFI PhoneNumberCountryEnum = "FI"
	PhoneNumberCountryFR PhoneNumberCountryEnum = "FR"
	PhoneNumberCountryGF PhoneNumberCountryEnum = "GF"
	PhoneNumberCountryPF PhoneNumberCountryEnum = "PF"
	PhoneNumberCountryTF PhoneNumberCountryEnum = "TF"
	PhoneNumberCountryGA PhoneNumberCountryEnum = "GA"
	PhoneNumberCountryGM PhoneNumberCountryEnum = "GM"
	PhoneNumberCountryGE PhoneNumberCountryEnum = "GE"
	PhoneNumberCountryDE PhoneNumberCountryEnum = "DE"
	PhoneNumberCountryGH PhoneNumberCountryEnum = "GH"
	PhoneNumberCountryGI PhoneNumberCountryEnum = "GI"
	PhoneNumberCountryGR PhoneNumberCountryEnum = "GR"
	PhoneNumberCountryGL PhoneNumberCountryEnum = "GL"
	PhoneNumberCountryGD PhoneNumberCountryEnum = "GD"
	PhoneNumberCountryGP PhoneNumberCountryEnum = "GP"
	PhoneNumberCountryGU PhoneNumberCountryEnum = "GU"
	PhoneNumberCountryGT PhoneNumberCountryEnum = "GT"
	PhoneNumberCountryGG PhoneNumberCountryEnum = "GG"
	PhoneNumberCountryGN PhoneNumberCountryEnum = "GN"
	PhoneNumberCountryGW PhoneNumberCountryEnum = "GW"
	PhoneNumberCountryGY PhoneNumberCountryEnum = "GY"
	PhoneNumberCountryHT PhoneNumberCountryEnum = "HT"
	PhoneNumberCountryHM PhoneNumberCountryEnum = "HM"
	PhoneNumberCountryVA PhoneNumberCountryEnum = "VA"
	PhoneNumberCountryHN PhoneNumberCountryEnum = "HN"
	PhoneNumberCountryHU PhoneNumberCountryEnum = "HU"
	PhoneNumberCountryIS PhoneNumberCountryEnum = "IS"
	PhoneNumberCountryIN PhoneNumberCountryEnum = "IN"
	PhoneNumberCountryID PhoneNumberCountryEnum = "ID"
	PhoneNumberCountryIR PhoneNumberCountryEnum = "IR"
	PhoneNumberCountryIQ PhoneNumberCountryEnum = "IQ"
	PhoneNumberCountryIE PhoneNumberCountryEnum = "IE"
	PhoneNumberCountryIM PhoneNumberCountryEnum = "IM"
	PhoneNumberCountryIL PhoneNumberCountryEnum = "IL"
	PhoneNumberCountryIT PhoneNumberCountryEnum = "IT"
	PhoneNumberCountryJM PhoneNumberCountryEnum = "JM"
	PhoneNumberCountryJP PhoneNumberCountryEnum = "JP"
	PhoneNumberCountryJE PhoneNumberCountryEnum = "JE"
	PhoneNumberCountryJO PhoneNumberCountryEnum = "JO"
	PhoneNumberCountryKZ PhoneNumberCountryEnum = "KZ"
	PhoneNumberCountryKE PhoneNumberCountryEnum = "KE"
	PhoneNumberCountryKI PhoneNumberCountryEnum = "KI"
	PhoneNumberCountryKW PhoneNumberCountryEnum = "KW"
	PhoneNumberCountryKG PhoneNumberCountryEnum = "KG"
	PhoneNumberCountryLA PhoneNumberCountryEnum = "LA"
	PhoneNumberCountryLV PhoneNumberCountryEnum = "LV"
	PhoneNumberCountryLB PhoneNumberCountryEnum = "LB"
	PhoneNumberCountryLS PhoneNumberCountryEnum = "LS"
	PhoneNumberCountryLR PhoneNumberCountryEnum = "LR"
	PhoneNumberCountryLY PhoneNumberCountryEnum = "LY"
	PhoneNumberCountryLI PhoneNumberCountryEnum = "LI"
	PhoneNumberCountryLT PhoneNumberCountryEnum = "LT"
	PhoneNumberCountryLU PhoneNumberCountryEnum = "LU"
	PhoneNumberCountryMG PhoneNumberCountryEnum = "MG"
	PhoneNumberCountryMW PhoneNumberCountryEnum = "MW"
	PhoneNumberCountryMY PhoneNumberCountryEnum = "MY"
	PhoneNumberCountryMV PhoneNumberCountryEnum = "MV"
	PhoneNumberCountryML PhoneNumberCountryEnum = "ML"
	PhoneNumberCountryMT PhoneNumberCountryEnum = "MT"
	PhoneNumberCountryMH PhoneNumberCountryEnum = "MH"
	PhoneNumberCountryMQ PhoneNumberCountryEnum = "MQ"
	PhoneNumberCountryMR PhoneNumberCountryEnum = "MR"
	PhoneNumberCountryMU PhoneNumberCountryEnum = "MU"
	PhoneNumberCountryYT PhoneNumberCountryEnum = "YT"
	PhoneNumberCountryFM PhoneNumberCountryEnum = "FM"
	PhoneNumberCountryMC PhoneNumberCountryEnum = "MC"
	PhoneNumberCountryMN PhoneNumberCountryEnum = "MN"
	PhoneNumberCountryME PhoneNumberCountryEnum = "ME"
	PhoneNumberCountryMS PhoneNumberCountryEnum = "MS"
	PhoneNumberCountryMA PhoneNumberCountryEnum = "MA"
	PhoneNumberCountryMZ PhoneNumberCountryEnum = "MZ"
	PhoneNumberCountryMM PhoneNumberCountryEnum = "MM"
	PhoneNumberCountryNA PhoneNumberCountryEnum = "NA"
	PhoneNumberCountryNR PhoneNumberCountryEnum = "NR"
	PhoneNumberCountryNP PhoneNumberCountryEnum = "NP"
	PhoneNumberCountryNL PhoneNumberCountryEnum = "NL"
	PhoneNumberCountryNC PhoneNumberCountryEnum = "NC"
	PhoneNumberCountryNZ PhoneNumberCountryEnum = "NZ"
	PhoneNumberCountryNI PhoneNumberCountryEnum = "NI"
	PhoneNumberCountryNE PhoneNumberCountryEnum = "NE"
	PhoneNumberCountryNG PhoneNumberCountryEnum = "NG"
	PhoneNumberCountryNU PhoneNumberCountryEnum = "NU"
	PhoneNumberCountryNF PhoneNumberCountryEnum = "NF"
	PhoneNumberCountryMP PhoneNumberCountryEnum = "MP"
	PhoneNumberCountryNO PhoneNumberCountryEnum = "NO"
	PhoneNumberCountryOM PhoneNumberCountryEnum = "OM"
	PhoneNumberCountryPK PhoneNumberCountryEnum = "PK"
	PhoneNumberCountryPW PhoneNumberCountryEnum = "PW"
	PhoneNumberCountryPA PhoneNumberCountryEnum = "PA"
	PhoneNumberCountryPG PhoneNumberCountryEnum = "PG"
	PhoneNumberCountryPY PhoneNumberCountryEnum = "PY"
	PhoneNumberCountryPE PhoneNumberCountryEnum = "PE"
	PhoneNumberCountryPH PhoneNumberCountryEnum = "PH"
	PhoneNumberCountryPN PhoneNumberCountryEnum = "PN"
	PhoneNumberCountryPL PhoneNumberCountryEnum = "PL"
	PhoneNumberCountryPT PhoneNumberCountryEnum = "PT"
	PhoneNumberCountryPR PhoneNumberCountryEnum = "PR"
	PhoneNumberCountryQA PhoneNumberCountryEnum = "QA"
	PhoneNumberCountryKR PhoneNumberCountryEnum = "KR"
	PhoneNumberCountryMD PhoneNumberCountryEnum = "MD"
	PhoneNumberCountryRO PhoneNumberCountryEnum = "RO"
	PhoneNumberCountryRU PhoneNumberCountryEnum = "RU"
	PhoneNumberCountryRW PhoneNumberCountryEnum = "RW"
	PhoneNumberCountryRE PhoneNumberCountryEnum = "RE"
	PhoneNumberCountryBL PhoneNumberCountryEnum = "BL"
	PhoneNumberCountrySH PhoneNumberCountryEnum = "SH"
	PhoneNumberCountryKN PhoneNumberCountryEnum = "KN"
	PhoneNumberCountryLC PhoneNumberCountryEnum = "LC"
	PhoneNumberCountryMF PhoneNumberCountryEnum = "MF"
	PhoneNumberCountryPM PhoneNumberCountryEnum = "PM"
	PhoneNumberCountryVC PhoneNumberCountryEnum = "VC"
	PhoneNumberCountryWS PhoneNumberCountryEnum = "WS"
	PhoneNumberCountrySM PhoneNumberCountryEnum = "SM"
	PhoneNumberCountryST PhoneNumberCountryEnum = "ST"
	PhoneNumberCountrySA PhoneNumberCountryEnum = "SA"
	PhoneNumberCountrySN PhoneNumberCountryEnum = "SN"
	PhoneNumberCountryRS PhoneNumberCountryEnum = "RS"
	PhoneNumberCountrySC PhoneNumberCountryEnum = "SC"
	PhoneNumberCountrySL PhoneNumberCountryEnum = "SL"
	PhoneNumberCountrySG PhoneNumberCountryEnum = "SG"
	PhoneNumberCountrySX PhoneNumberCountryEnum = "SX"
	PhoneNumberCountrySK PhoneNumberCountryEnum = "SK"
	PhoneNumberCountrySI PhoneNumberCountryEnum = "SI"
	PhoneNumberCountrySB PhoneNumberCountryEnum = "SB"
	PhoneNumberCountrySO PhoneNumberCountryEnum = "SO"
	PhoneNumberCountryZA PhoneNumberCountryEnum = "ZA"
	PhoneNumberCountryGS PhoneNumberCountryEnum = "GS"
	PhoneNumberCountrySS PhoneNumberCountryEnum = "SS"
	PhoneNumberCountryES PhoneNumberCountryEnum = "ES"
	PhoneNumberCountryLK PhoneNumberCountryEnum = "LK"
	PhoneNumberCountryPS PhoneNumberCountryEnum = "PS"
	PhoneNumberCountrySD PhoneNumberCountryEnum = "SD"
	PhoneNumberCountrySR PhoneNumberCountryEnum = "SR"
	PhoneNumberCountrySJ PhoneNumberCountryEnum = "SJ"
	PhoneNumberCountrySE PhoneNumberCountryEnum = "SE"
	PhoneNumberCountryCH PhoneNumberCountryEnum = "CH"
	PhoneNumberCountrySY PhoneNumberCountryEnum = "SY"
	PhoneNumberCountryTW PhoneNumberCountryEnum = "TW"
	PhoneNumberCountryTJ PhoneNumberCountryEnum = "TJ"
	PhoneNumberCountryTH PhoneNumberCountryEnum = "TH"
	PhoneNumberCountryMK PhoneNumberCountryEnum = "MK"
	PhoneNumberCountryTL PhoneNumberCountryEnum = "TL"
	PhoneNumberCountryTG PhoneNumberCountryEnum = "TG"
	PhoneNumberCountryTK PhoneNumberCountryEnum = "TK"
	PhoneNumberCountryTO PhoneNumberCountryEnum = "TO"
	PhoneNumberCountryTT PhoneNumberCountryEnum = "TT"
	PhoneNumberCountryTN PhoneNumberCountryEnum = "TN"
	PhoneNumberCountryTR PhoneNumberCountryEnum = "TR"
	PhoneNumberCountryTM PhoneNumberCountryEnum = "TM"
	PhoneNumberCountryTC PhoneNumberCountryEnum = "TC"
	PhoneNumberCountryTV PhoneNumberCountryEnum = "TV"
	PhoneNumberCountryUG PhoneNumberCountryEnum = "UG"
	PhoneNumberCountryUA PhoneNumberCountryEnum = "UA"
	PhoneNumberCountryAE PhoneNumberCountryEnum = "AE"
	PhoneNumberCountryGB PhoneNumberCountryEnum = "GB"
	PhoneNumberCountryTZ PhoneNumberCountryEnum = "TZ"
	PhoneNumberCountryUM PhoneNumberCountryEnum = "UM"
	PhoneNumberCountryVI PhoneNumberCountryEnum = "VI"
	PhoneNumberCountryUY PhoneNumberCountryEnum = "UY"
	PhoneNumberCountryUZ PhoneNumberCountryEnum = "UZ"
	PhoneNumberCountryVU PhoneNumberCountryEnum = "VU"
	PhoneNumberCountryVE PhoneNumberCountryEnum = "VE"
	PhoneNumberCountryVN PhoneNumberCountryEnum = "VN"
	PhoneNumberCountryWF PhoneNumberCountryEnum = "WF"
	PhoneNumberCountryEH PhoneNumberCountryEnum = "EH"
	PhoneNumberCountryYE PhoneNumberCountryEnum = "YE"
	PhoneNumberCountryZM PhoneNumberCountryEnum = "ZM"
	PhoneNumberCountryZW PhoneNumberCountryEnum = "ZW"
)

type PhoneNumberTypeEnum string

const (
	PhoneNumberTypeMobile   PhoneNumberTypeEnum = "Mobile"
	PhoneNumberTypeLandline PhoneNumberTypeEnum = "Landline"
	PhoneNumberTypeVoIP     PhoneNumberTypeEnum = "VoIP"
	PhoneNumberTypeOther    PhoneNumberTypeEnum = "Other"
)

type PhoneNumberUserDefinedTypeEnum string

const (
	PhoneNumberUserDefinedTypeMobile   PhoneNumberUserDefinedTypeEnum = "Mobile"
	PhoneNumberUserDefinedTypeWork     PhoneNumberUserDefinedTypeEnum = "Work"
	PhoneNumberUserDefinedTypeHome     PhoneNumberUserDefinedTypeEnum = "Home"
	PhoneNumberUserDefinedTypeOffice   PhoneNumberUserDefinedTypeEnum = "Office"
	PhoneNumberUserDefinedTypeFax      PhoneNumberUserDefinedTypeEnum = "Fax"
	PhoneNumberUserDefinedTypeVoIP     PhoneNumberUserDefinedTypeEnum = "VoIP"
	PhoneNumberUserDefinedTypeLandline PhoneNumberUserDefinedTypeEnum = "Landline"
	PhoneNumberUserDefinedTypeOther    PhoneNumberUserDefinedTypeEnum = "Other"
)

type PhoneNumber struct {
	ID          string              `bson:"_id" gorm:"primaryKey;not null;column:id" json:"id"`
	CreatedDate datatypes.DateTime  `gorm:"column:createdDate;not null;column:createdDate" json:"createdDate"`
	UpdatedDate *datatypes.DateTime `gorm:"column:updatedDate;column:updatedDate" json:"updatedDate"`
	Meta        *datatypes.JSON     `gorm:"column:meta;not null;column:meta" json:"meta,omitempty"`    // the metadata about the most recent change to the row
	Metadata    *datatypes.JSON     `gorm:"column:metadata;column:metadata" json:"metadata,omitempty"` // metadata reserved for customers to control
	CompanyID   string              `gorm:"not null;column:companyId" json:"companyId"`
	LocationID  string              `gorm:"not null;column:locationId" json:"locationId"`

	CarrierName       *string                         `gorm:"column:carrierName" json:"carrierName"`
	Country           *PhoneNumberCountryEnum         `gorm:"column:country" json:"country"`
	CustomerID        string                          `gorm:"not null;column:customerId" json:"customerId"`
	Extension         *string                         `gorm:"column:extension" json:"extension"`
	LastVerifiedDate  *datatypes.DateTime             `gorm:"column:lastVerifiedDate" json:"lastVerifiedDate"`
	MobileCountryCode *string                         `gorm:"column:mobileCountryCode" json:"mobileCountryCode"` // if a mobile number, the mobile country code
	MobileNetworkCode *string                         `gorm:"column:mobileNetworkCode" json:"mobileNetworkCode"` // if a mobile number, the mobile network code
	Number            string                          `gorm:"not null;column:number" json:"number"`
	OptIn             bool                            `gorm:"not null;column:optIn" json:"optIn"`
	OptInVerifiedDate *datatypes.DateTime             `gorm:"column:optInVerifiedDate" json:"optInVerifiedDate"`
	Primary           bool                            `gorm:"not null;column:primary" json:"primary"`
	Type              *PhoneNumberTypeEnum            `gorm:"column:type" json:"type"`
	UserDefinedType   *PhoneNumberUserDefinedTypeEnum `gorm:"column:userDefinedType" json:"userDefinedType"`
}

var _ Model = (*PhoneNumber)(nil)

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *PhoneNumber) TableName() string {
	return "phone_number"
}

// PrimaryKeys returns an array of the primary keys for this model
func (m *PhoneNumber) PrimaryKeys() []string {
	return []string{"id"}
}

// String returns a string representation as JSON for this model
func (m *PhoneNumber) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewPhoneNumber returns a new model instance from an encoded buffer
func NewPhoneNumber(buf []byte) (*PhoneNumber, error) {
	var result PhoneNumber
	err := json.Unmarshal(buf, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// NewPhoneNumberFromChangeEvent returns a new model instance from an encoded buffer as change event
func NewPhoneNumberFromChangeEvent(buf []byte, gzip bool) (*datatypes.ChangeEvent[PhoneNumber], error) {
	var result datatypes.ChangeEvent[PhoneNumber]
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
