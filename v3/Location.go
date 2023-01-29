// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	datatypes "github.com/shopmonkeyus/go-datamodel/datatypes"
)

// Location schema
type LocationCountryEnum string

const (
	LocationCountryUS LocationCountryEnum = "US"
	LocationCountryCA LocationCountryEnum = "CA"
	LocationCountryMX LocationCountryEnum = "MX"
	LocationCountryAF LocationCountryEnum = "AF"
	LocationCountryAX LocationCountryEnum = "AX"
	LocationCountryAL LocationCountryEnum = "AL"
	LocationCountryDZ LocationCountryEnum = "DZ"
	LocationCountryAS LocationCountryEnum = "AS"
	LocationCountryAD LocationCountryEnum = "AD"
	LocationCountryAO LocationCountryEnum = "AO"
	LocationCountryAI LocationCountryEnum = "AI"
	LocationCountryAQ LocationCountryEnum = "AQ"
	LocationCountryAG LocationCountryEnum = "AG"
	LocationCountryAR LocationCountryEnum = "AR"
	LocationCountryAM LocationCountryEnum = "AM"
	LocationCountryAW LocationCountryEnum = "AW"
	LocationCountryAU LocationCountryEnum = "AU"
	LocationCountryAT LocationCountryEnum = "AT"
	LocationCountryAZ LocationCountryEnum = "AZ"
	LocationCountryBS LocationCountryEnum = "BS"
	LocationCountryBH LocationCountryEnum = "BH"
	LocationCountryBD LocationCountryEnum = "BD"
	LocationCountryBB LocationCountryEnum = "BB"
	LocationCountryBY LocationCountryEnum = "BY"
	LocationCountryBE LocationCountryEnum = "BE"
	LocationCountryBZ LocationCountryEnum = "BZ"
	LocationCountryBJ LocationCountryEnum = "BJ"
	LocationCountryBM LocationCountryEnum = "BM"
	LocationCountryBT LocationCountryEnum = "BT"
	LocationCountryBO LocationCountryEnum = "BO"
	LocationCountryBQ LocationCountryEnum = "BQ"
	LocationCountryBA LocationCountryEnum = "BA"
	LocationCountryBW LocationCountryEnum = "BW"
	LocationCountryBV LocationCountryEnum = "BV"
	LocationCountryBR LocationCountryEnum = "BR"
	LocationCountryIO LocationCountryEnum = "IO"
	LocationCountryVG LocationCountryEnum = "VG"
	LocationCountryBN LocationCountryEnum = "BN"
	LocationCountryBG LocationCountryEnum = "BG"
	LocationCountryBF LocationCountryEnum = "BF"
	LocationCountryBI LocationCountryEnum = "BI"
	LocationCountryCV LocationCountryEnum = "CV"
	LocationCountryKH LocationCountryEnum = "KH"
	LocationCountryCM LocationCountryEnum = "CM"
	LocationCountryKY LocationCountryEnum = "KY"
	LocationCountryCF LocationCountryEnum = "CF"
	LocationCountryTD LocationCountryEnum = "TD"
	LocationCountryCL LocationCountryEnum = "CL"
	LocationCountryCN LocationCountryEnum = "CN"
	LocationCountryHK LocationCountryEnum = "HK"
	LocationCountryMO LocationCountryEnum = "MO"
	LocationCountryCX LocationCountryEnum = "CX"
	LocationCountryCC LocationCountryEnum = "CC"
	LocationCountryCO LocationCountryEnum = "CO"
	LocationCountryKM LocationCountryEnum = "KM"
	LocationCountryCG LocationCountryEnum = "CG"
	LocationCountryCK LocationCountryEnum = "CK"
	LocationCountryCR LocationCountryEnum = "CR"
	LocationCountryHR LocationCountryEnum = "HR"
	LocationCountryCU LocationCountryEnum = "CU"
	LocationCountryCW LocationCountryEnum = "CW"
	LocationCountryCY LocationCountryEnum = "CY"
	LocationCountryCZ LocationCountryEnum = "CZ"
	LocationCountryCI LocationCountryEnum = "CI"
	LocationCountryKP LocationCountryEnum = "KP"
	LocationCountryCD LocationCountryEnum = "CD"
	LocationCountryDK LocationCountryEnum = "DK"
	LocationCountryDJ LocationCountryEnum = "DJ"
	LocationCountryDM LocationCountryEnum = "DM"
	LocationCountryDO LocationCountryEnum = "DO"
	LocationCountryEC LocationCountryEnum = "EC"
	LocationCountryEG LocationCountryEnum = "EG"
	LocationCountrySV LocationCountryEnum = "SV"
	LocationCountryGQ LocationCountryEnum = "GQ"
	LocationCountryER LocationCountryEnum = "ER"
	LocationCountryEE LocationCountryEnum = "EE"
	LocationCountrySZ LocationCountryEnum = "SZ"
	LocationCountryET LocationCountryEnum = "ET"
	LocationCountryFK LocationCountryEnum = "FK"
	LocationCountryFO LocationCountryEnum = "FO"
	LocationCountryFJ LocationCountryEnum = "FJ"
	LocationCountryFI LocationCountryEnum = "FI"
	LocationCountryFR LocationCountryEnum = "FR"
	LocationCountryGF LocationCountryEnum = "GF"
	LocationCountryPF LocationCountryEnum = "PF"
	LocationCountryTF LocationCountryEnum = "TF"
	LocationCountryGA LocationCountryEnum = "GA"
	LocationCountryGM LocationCountryEnum = "GM"
	LocationCountryGE LocationCountryEnum = "GE"
	LocationCountryDE LocationCountryEnum = "DE"
	LocationCountryGH LocationCountryEnum = "GH"
	LocationCountryGI LocationCountryEnum = "GI"
	LocationCountryGR LocationCountryEnum = "GR"
	LocationCountryGL LocationCountryEnum = "GL"
	LocationCountryGD LocationCountryEnum = "GD"
	LocationCountryGP LocationCountryEnum = "GP"
	LocationCountryGU LocationCountryEnum = "GU"
	LocationCountryGT LocationCountryEnum = "GT"
	LocationCountryGG LocationCountryEnum = "GG"
	LocationCountryGN LocationCountryEnum = "GN"
	LocationCountryGW LocationCountryEnum = "GW"
	LocationCountryGY LocationCountryEnum = "GY"
	LocationCountryHT LocationCountryEnum = "HT"
	LocationCountryHM LocationCountryEnum = "HM"
	LocationCountryVA LocationCountryEnum = "VA"
	LocationCountryHN LocationCountryEnum = "HN"
	LocationCountryHU LocationCountryEnum = "HU"
	LocationCountryIS LocationCountryEnum = "IS"
	LocationCountryIN LocationCountryEnum = "IN"
	LocationCountryID LocationCountryEnum = "ID"
	LocationCountryIR LocationCountryEnum = "IR"
	LocationCountryIQ LocationCountryEnum = "IQ"
	LocationCountryIE LocationCountryEnum = "IE"
	LocationCountryIM LocationCountryEnum = "IM"
	LocationCountryIL LocationCountryEnum = "IL"
	LocationCountryIT LocationCountryEnum = "IT"
	LocationCountryJM LocationCountryEnum = "JM"
	LocationCountryJP LocationCountryEnum = "JP"
	LocationCountryJE LocationCountryEnum = "JE"
	LocationCountryJO LocationCountryEnum = "JO"
	LocationCountryKZ LocationCountryEnum = "KZ"
	LocationCountryKE LocationCountryEnum = "KE"
	LocationCountryKI LocationCountryEnum = "KI"
	LocationCountryKW LocationCountryEnum = "KW"
	LocationCountryKG LocationCountryEnum = "KG"
	LocationCountryLA LocationCountryEnum = "LA"
	LocationCountryLV LocationCountryEnum = "LV"
	LocationCountryLB LocationCountryEnum = "LB"
	LocationCountryLS LocationCountryEnum = "LS"
	LocationCountryLR LocationCountryEnum = "LR"
	LocationCountryLY LocationCountryEnum = "LY"
	LocationCountryLI LocationCountryEnum = "LI"
	LocationCountryLT LocationCountryEnum = "LT"
	LocationCountryLU LocationCountryEnum = "LU"
	LocationCountryMG LocationCountryEnum = "MG"
	LocationCountryMW LocationCountryEnum = "MW"
	LocationCountryMY LocationCountryEnum = "MY"
	LocationCountryMV LocationCountryEnum = "MV"
	LocationCountryML LocationCountryEnum = "ML"
	LocationCountryMT LocationCountryEnum = "MT"
	LocationCountryMH LocationCountryEnum = "MH"
	LocationCountryMQ LocationCountryEnum = "MQ"
	LocationCountryMR LocationCountryEnum = "MR"
	LocationCountryMU LocationCountryEnum = "MU"
	LocationCountryYT LocationCountryEnum = "YT"
	LocationCountryFM LocationCountryEnum = "FM"
	LocationCountryMC LocationCountryEnum = "MC"
	LocationCountryMN LocationCountryEnum = "MN"
	LocationCountryME LocationCountryEnum = "ME"
	LocationCountryMS LocationCountryEnum = "MS"
	LocationCountryMA LocationCountryEnum = "MA"
	LocationCountryMZ LocationCountryEnum = "MZ"
	LocationCountryMM LocationCountryEnum = "MM"
	LocationCountryNA LocationCountryEnum = "NA"
	LocationCountryNR LocationCountryEnum = "NR"
	LocationCountryNP LocationCountryEnum = "NP"
	LocationCountryNL LocationCountryEnum = "NL"
	LocationCountryNC LocationCountryEnum = "NC"
	LocationCountryNZ LocationCountryEnum = "NZ"
	LocationCountryNI LocationCountryEnum = "NI"
	LocationCountryNE LocationCountryEnum = "NE"
	LocationCountryNG LocationCountryEnum = "NG"
	LocationCountryNU LocationCountryEnum = "NU"
	LocationCountryNF LocationCountryEnum = "NF"
	LocationCountryMP LocationCountryEnum = "MP"
	LocationCountryNO LocationCountryEnum = "NO"
	LocationCountryOM LocationCountryEnum = "OM"
	LocationCountryPK LocationCountryEnum = "PK"
	LocationCountryPW LocationCountryEnum = "PW"
	LocationCountryPA LocationCountryEnum = "PA"
	LocationCountryPG LocationCountryEnum = "PG"
	LocationCountryPY LocationCountryEnum = "PY"
	LocationCountryPE LocationCountryEnum = "PE"
	LocationCountryPH LocationCountryEnum = "PH"
	LocationCountryPN LocationCountryEnum = "PN"
	LocationCountryPL LocationCountryEnum = "PL"
	LocationCountryPT LocationCountryEnum = "PT"
	LocationCountryPR LocationCountryEnum = "PR"
	LocationCountryQA LocationCountryEnum = "QA"
	LocationCountryKR LocationCountryEnum = "KR"
	LocationCountryMD LocationCountryEnum = "MD"
	LocationCountryRO LocationCountryEnum = "RO"
	LocationCountryRU LocationCountryEnum = "RU"
	LocationCountryRW LocationCountryEnum = "RW"
	LocationCountryRE LocationCountryEnum = "RE"
	LocationCountryBL LocationCountryEnum = "BL"
	LocationCountrySH LocationCountryEnum = "SH"
	LocationCountryKN LocationCountryEnum = "KN"
	LocationCountryLC LocationCountryEnum = "LC"
	LocationCountryMF LocationCountryEnum = "MF"
	LocationCountryPM LocationCountryEnum = "PM"
	LocationCountryVC LocationCountryEnum = "VC"
	LocationCountryWS LocationCountryEnum = "WS"
	LocationCountrySM LocationCountryEnum = "SM"
	LocationCountryST LocationCountryEnum = "ST"
	LocationCountrySA LocationCountryEnum = "SA"
	LocationCountrySN LocationCountryEnum = "SN"
	LocationCountryRS LocationCountryEnum = "RS"
	LocationCountrySC LocationCountryEnum = "SC"
	LocationCountrySL LocationCountryEnum = "SL"
	LocationCountrySG LocationCountryEnum = "SG"
	LocationCountrySX LocationCountryEnum = "SX"
	LocationCountrySK LocationCountryEnum = "SK"
	LocationCountrySI LocationCountryEnum = "SI"
	LocationCountrySB LocationCountryEnum = "SB"
	LocationCountrySO LocationCountryEnum = "SO"
	LocationCountryZA LocationCountryEnum = "ZA"
	LocationCountryGS LocationCountryEnum = "GS"
	LocationCountrySS LocationCountryEnum = "SS"
	LocationCountryES LocationCountryEnum = "ES"
	LocationCountryLK LocationCountryEnum = "LK"
	LocationCountryPS LocationCountryEnum = "PS"
	LocationCountrySD LocationCountryEnum = "SD"
	LocationCountrySR LocationCountryEnum = "SR"
	LocationCountrySJ LocationCountryEnum = "SJ"
	LocationCountrySE LocationCountryEnum = "SE"
	LocationCountryCH LocationCountryEnum = "CH"
	LocationCountrySY LocationCountryEnum = "SY"
	LocationCountryTW LocationCountryEnum = "TW"
	LocationCountryTJ LocationCountryEnum = "TJ"
	LocationCountryTH LocationCountryEnum = "TH"
	LocationCountryMK LocationCountryEnum = "MK"
	LocationCountryTL LocationCountryEnum = "TL"
	LocationCountryTG LocationCountryEnum = "TG"
	LocationCountryTK LocationCountryEnum = "TK"
	LocationCountryTO LocationCountryEnum = "TO"
	LocationCountryTT LocationCountryEnum = "TT"
	LocationCountryTN LocationCountryEnum = "TN"
	LocationCountryTR LocationCountryEnum = "TR"
	LocationCountryTM LocationCountryEnum = "TM"
	LocationCountryTC LocationCountryEnum = "TC"
	LocationCountryTV LocationCountryEnum = "TV"
	LocationCountryUG LocationCountryEnum = "UG"
	LocationCountryUA LocationCountryEnum = "UA"
	LocationCountryAE LocationCountryEnum = "AE"
	LocationCountryGB LocationCountryEnum = "GB"
	LocationCountryTZ LocationCountryEnum = "TZ"
	LocationCountryUM LocationCountryEnum = "UM"
	LocationCountryVI LocationCountryEnum = "VI"
	LocationCountryUY LocationCountryEnum = "UY"
	LocationCountryUZ LocationCountryEnum = "UZ"
	LocationCountryVU LocationCountryEnum = "VU"
	LocationCountryVE LocationCountryEnum = "VE"
	LocationCountryVN LocationCountryEnum = "VN"
	LocationCountryWF LocationCountryEnum = "WF"
	LocationCountryEH LocationCountryEnum = "EH"
	LocationCountryYE LocationCountryEnum = "YE"
	LocationCountryZM LocationCountryEnum = "ZM"
	LocationCountryZW LocationCountryEnum = "ZW"
)

type Location struct {
	ID          string              `bson:"_id" gorm:"primaryKey;not null;column:id" json:"id"`
	CreatedDate datatypes.DateTime  `gorm:"column:createdDate;not null;column:createdDate" json:"createdDate"`
	UpdatedDate *datatypes.DateTime `gorm:"column:updatedDate;column:updatedDate" json:"updatedDate"`
	Meta        *datatypes.JSON     `gorm:"column:meta;not null;column:meta" json:"meta,omitempty"`    // the metadata about the most recent change to the row
	Metadata    *datatypes.JSON     `gorm:"column:metadata;column:metadata" json:"metadata,omitempty"` // metadata reserved for customers to control
	CompanyID   string              `gorm:"not null;column:companyId" json:"companyId"`

	Address1              string              `gorm:"not null;column:address1" json:"address1"`
	Address2              *string             `gorm:"column:address2" json:"address2"`
	City                  string              `gorm:"not null;column:city" json:"city"`
	ContactName           *string             `gorm:"column:contactName" json:"contactName"` // the location contact name
	Country               LocationCountryEnum `gorm:"not null;column:country" json:"country"`
	Email                 *string             `gorm:"column:email" json:"email"`                                 // the location email for generic inqueries
	LocationPhoneNumberID *string             `gorm:"column:locationPhoneNumberId" json:"locationPhoneNumberId"` // the location phone number
	Name                  string              `gorm:"not null;column:name" json:"name"`
	PostalCode            string              `gorm:"not null;column:postalCode" json:"postalCode"`
	State                 string              `gorm:"not null;column:state" json:"state"`
	Timezone              string              `gorm:"not null;column:timezone" json:"timezone"`
	Website               *string             `gorm:"column:website" json:"website"`
}

var _ Model = (*Location)(nil)

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *Location) TableName() string {
	return "location"
}

// PrimaryKeys returns an array of the primary keys for this model
func (m *Location) PrimaryKeys() []string {
	return []string{"id"}
}

// String returns a string representation as JSON for this model
func (m *Location) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewLocation returns a new model instance from an encoded buffer
func NewLocation(buf []byte) (*Location, error) {
	var result Location
	err := json.Unmarshal(buf, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// NewLocationFromChangeEvent returns a new model instance from an encoded buffer as change event
func NewLocationFromChangeEvent(buf []byte, gzip bool) (*datatypes.ChangeEvent[Location], error) {
	var result datatypes.ChangeEvent[Location]
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
