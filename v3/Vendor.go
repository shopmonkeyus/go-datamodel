// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	datatypes "github.com/shopmonkeyus/go-datamodel/datatypes"
)

type VendorCountryEnum string

const (
	VendorCountryUS VendorCountryEnum = "US"
	VendorCountryCA VendorCountryEnum = "CA"
	VendorCountryMX VendorCountryEnum = "MX"
	VendorCountryAF VendorCountryEnum = "AF"
	VendorCountryAX VendorCountryEnum = "AX"
	VendorCountryAL VendorCountryEnum = "AL"
	VendorCountryDZ VendorCountryEnum = "DZ"
	VendorCountryAS VendorCountryEnum = "AS"
	VendorCountryAD VendorCountryEnum = "AD"
	VendorCountryAO VendorCountryEnum = "AO"
	VendorCountryAI VendorCountryEnum = "AI"
	VendorCountryAQ VendorCountryEnum = "AQ"
	VendorCountryAG VendorCountryEnum = "AG"
	VendorCountryAR VendorCountryEnum = "AR"
	VendorCountryAM VendorCountryEnum = "AM"
	VendorCountryAW VendorCountryEnum = "AW"
	VendorCountryAU VendorCountryEnum = "AU"
	VendorCountryAT VendorCountryEnum = "AT"
	VendorCountryAZ VendorCountryEnum = "AZ"
	VendorCountryBS VendorCountryEnum = "BS"
	VendorCountryBH VendorCountryEnum = "BH"
	VendorCountryBD VendorCountryEnum = "BD"
	VendorCountryBB VendorCountryEnum = "BB"
	VendorCountryBY VendorCountryEnum = "BY"
	VendorCountryBE VendorCountryEnum = "BE"
	VendorCountryBZ VendorCountryEnum = "BZ"
	VendorCountryBJ VendorCountryEnum = "BJ"
	VendorCountryBM VendorCountryEnum = "BM"
	VendorCountryBT VendorCountryEnum = "BT"
	VendorCountryBO VendorCountryEnum = "BO"
	VendorCountryBQ VendorCountryEnum = "BQ"
	VendorCountryBA VendorCountryEnum = "BA"
	VendorCountryBW VendorCountryEnum = "BW"
	VendorCountryBV VendorCountryEnum = "BV"
	VendorCountryBR VendorCountryEnum = "BR"
	VendorCountryIO VendorCountryEnum = "IO"
	VendorCountryVG VendorCountryEnum = "VG"
	VendorCountryBN VendorCountryEnum = "BN"
	VendorCountryBG VendorCountryEnum = "BG"
	VendorCountryBF VendorCountryEnum = "BF"
	VendorCountryBI VendorCountryEnum = "BI"
	VendorCountryCV VendorCountryEnum = "CV"
	VendorCountryKH VendorCountryEnum = "KH"
	VendorCountryCM VendorCountryEnum = "CM"
	VendorCountryKY VendorCountryEnum = "KY"
	VendorCountryCF VendorCountryEnum = "CF"
	VendorCountryTD VendorCountryEnum = "TD"
	VendorCountryCL VendorCountryEnum = "CL"
	VendorCountryCN VendorCountryEnum = "CN"
	VendorCountryHK VendorCountryEnum = "HK"
	VendorCountryMO VendorCountryEnum = "MO"
	VendorCountryCX VendorCountryEnum = "CX"
	VendorCountryCC VendorCountryEnum = "CC"
	VendorCountryCO VendorCountryEnum = "CO"
	VendorCountryKM VendorCountryEnum = "KM"
	VendorCountryCG VendorCountryEnum = "CG"
	VendorCountryCK VendorCountryEnum = "CK"
	VendorCountryCR VendorCountryEnum = "CR"
	VendorCountryHR VendorCountryEnum = "HR"
	VendorCountryCU VendorCountryEnum = "CU"
	VendorCountryCW VendorCountryEnum = "CW"
	VendorCountryCY VendorCountryEnum = "CY"
	VendorCountryCZ VendorCountryEnum = "CZ"
	VendorCountryCI VendorCountryEnum = "CI"
	VendorCountryKP VendorCountryEnum = "KP"
	VendorCountryCD VendorCountryEnum = "CD"
	VendorCountryDK VendorCountryEnum = "DK"
	VendorCountryDJ VendorCountryEnum = "DJ"
	VendorCountryDM VendorCountryEnum = "DM"
	VendorCountryDO VendorCountryEnum = "DO"
	VendorCountryEC VendorCountryEnum = "EC"
	VendorCountryEG VendorCountryEnum = "EG"
	VendorCountrySV VendorCountryEnum = "SV"
	VendorCountryGQ VendorCountryEnum = "GQ"
	VendorCountryER VendorCountryEnum = "ER"
	VendorCountryEE VendorCountryEnum = "EE"
	VendorCountrySZ VendorCountryEnum = "SZ"
	VendorCountryET VendorCountryEnum = "ET"
	VendorCountryFK VendorCountryEnum = "FK"
	VendorCountryFO VendorCountryEnum = "FO"
	VendorCountryFJ VendorCountryEnum = "FJ"
	VendorCountryFI VendorCountryEnum = "FI"
	VendorCountryFR VendorCountryEnum = "FR"
	VendorCountryGF VendorCountryEnum = "GF"
	VendorCountryPF VendorCountryEnum = "PF"
	VendorCountryTF VendorCountryEnum = "TF"
	VendorCountryGA VendorCountryEnum = "GA"
	VendorCountryGM VendorCountryEnum = "GM"
	VendorCountryGE VendorCountryEnum = "GE"
	VendorCountryDE VendorCountryEnum = "DE"
	VendorCountryGH VendorCountryEnum = "GH"
	VendorCountryGI VendorCountryEnum = "GI"
	VendorCountryGR VendorCountryEnum = "GR"
	VendorCountryGL VendorCountryEnum = "GL"
	VendorCountryGD VendorCountryEnum = "GD"
	VendorCountryGP VendorCountryEnum = "GP"
	VendorCountryGU VendorCountryEnum = "GU"
	VendorCountryGT VendorCountryEnum = "GT"
	VendorCountryGG VendorCountryEnum = "GG"
	VendorCountryGN VendorCountryEnum = "GN"
	VendorCountryGW VendorCountryEnum = "GW"
	VendorCountryGY VendorCountryEnum = "GY"
	VendorCountryHT VendorCountryEnum = "HT"
	VendorCountryHM VendorCountryEnum = "HM"
	VendorCountryVA VendorCountryEnum = "VA"
	VendorCountryHN VendorCountryEnum = "HN"
	VendorCountryHU VendorCountryEnum = "HU"
	VendorCountryIS VendorCountryEnum = "IS"
	VendorCountryIN VendorCountryEnum = "IN"
	VendorCountryID VendorCountryEnum = "ID"
	VendorCountryIR VendorCountryEnum = "IR"
	VendorCountryIQ VendorCountryEnum = "IQ"
	VendorCountryIE VendorCountryEnum = "IE"
	VendorCountryIM VendorCountryEnum = "IM"
	VendorCountryIL VendorCountryEnum = "IL"
	VendorCountryIT VendorCountryEnum = "IT"
	VendorCountryJM VendorCountryEnum = "JM"
	VendorCountryJP VendorCountryEnum = "JP"
	VendorCountryJE VendorCountryEnum = "JE"
	VendorCountryJO VendorCountryEnum = "JO"
	VendorCountryKZ VendorCountryEnum = "KZ"
	VendorCountryKE VendorCountryEnum = "KE"
	VendorCountryKI VendorCountryEnum = "KI"
	VendorCountryKW VendorCountryEnum = "KW"
	VendorCountryKG VendorCountryEnum = "KG"
	VendorCountryLA VendorCountryEnum = "LA"
	VendorCountryLV VendorCountryEnum = "LV"
	VendorCountryLB VendorCountryEnum = "LB"
	VendorCountryLS VendorCountryEnum = "LS"
	VendorCountryLR VendorCountryEnum = "LR"
	VendorCountryLY VendorCountryEnum = "LY"
	VendorCountryLI VendorCountryEnum = "LI"
	VendorCountryLT VendorCountryEnum = "LT"
	VendorCountryLU VendorCountryEnum = "LU"
	VendorCountryMG VendorCountryEnum = "MG"
	VendorCountryMW VendorCountryEnum = "MW"
	VendorCountryMY VendorCountryEnum = "MY"
	VendorCountryMV VendorCountryEnum = "MV"
	VendorCountryML VendorCountryEnum = "ML"
	VendorCountryMT VendorCountryEnum = "MT"
	VendorCountryMH VendorCountryEnum = "MH"
	VendorCountryMQ VendorCountryEnum = "MQ"
	VendorCountryMR VendorCountryEnum = "MR"
	VendorCountryMU VendorCountryEnum = "MU"
	VendorCountryYT VendorCountryEnum = "YT"
	VendorCountryFM VendorCountryEnum = "FM"
	VendorCountryMC VendorCountryEnum = "MC"
	VendorCountryMN VendorCountryEnum = "MN"
	VendorCountryME VendorCountryEnum = "ME"
	VendorCountryMS VendorCountryEnum = "MS"
	VendorCountryMA VendorCountryEnum = "MA"
	VendorCountryMZ VendorCountryEnum = "MZ"
	VendorCountryMM VendorCountryEnum = "MM"
	VendorCountryNA VendorCountryEnum = "NA"
	VendorCountryNR VendorCountryEnum = "NR"
	VendorCountryNP VendorCountryEnum = "NP"
	VendorCountryNL VendorCountryEnum = "NL"
	VendorCountryNC VendorCountryEnum = "NC"
	VendorCountryNZ VendorCountryEnum = "NZ"
	VendorCountryNI VendorCountryEnum = "NI"
	VendorCountryNE VendorCountryEnum = "NE"
	VendorCountryNG VendorCountryEnum = "NG"
	VendorCountryNU VendorCountryEnum = "NU"
	VendorCountryNF VendorCountryEnum = "NF"
	VendorCountryMP VendorCountryEnum = "MP"
	VendorCountryNO VendorCountryEnum = "NO"
	VendorCountryOM VendorCountryEnum = "OM"
	VendorCountryPK VendorCountryEnum = "PK"
	VendorCountryPW VendorCountryEnum = "PW"
	VendorCountryPA VendorCountryEnum = "PA"
	VendorCountryPG VendorCountryEnum = "PG"
	VendorCountryPY VendorCountryEnum = "PY"
	VendorCountryPE VendorCountryEnum = "PE"
	VendorCountryPH VendorCountryEnum = "PH"
	VendorCountryPN VendorCountryEnum = "PN"
	VendorCountryPL VendorCountryEnum = "PL"
	VendorCountryPT VendorCountryEnum = "PT"
	VendorCountryPR VendorCountryEnum = "PR"
	VendorCountryQA VendorCountryEnum = "QA"
	VendorCountryKR VendorCountryEnum = "KR"
	VendorCountryMD VendorCountryEnum = "MD"
	VendorCountryRO VendorCountryEnum = "RO"
	VendorCountryRU VendorCountryEnum = "RU"
	VendorCountryRW VendorCountryEnum = "RW"
	VendorCountryRE VendorCountryEnum = "RE"
	VendorCountryBL VendorCountryEnum = "BL"
	VendorCountrySH VendorCountryEnum = "SH"
	VendorCountryKN VendorCountryEnum = "KN"
	VendorCountryLC VendorCountryEnum = "LC"
	VendorCountryMF VendorCountryEnum = "MF"
	VendorCountryPM VendorCountryEnum = "PM"
	VendorCountryVC VendorCountryEnum = "VC"
	VendorCountryWS VendorCountryEnum = "WS"
	VendorCountrySM VendorCountryEnum = "SM"
	VendorCountryST VendorCountryEnum = "ST"
	VendorCountrySA VendorCountryEnum = "SA"
	VendorCountrySN VendorCountryEnum = "SN"
	VendorCountryRS VendorCountryEnum = "RS"
	VendorCountrySC VendorCountryEnum = "SC"
	VendorCountrySL VendorCountryEnum = "SL"
	VendorCountrySG VendorCountryEnum = "SG"
	VendorCountrySX VendorCountryEnum = "SX"
	VendorCountrySK VendorCountryEnum = "SK"
	VendorCountrySI VendorCountryEnum = "SI"
	VendorCountrySB VendorCountryEnum = "SB"
	VendorCountrySO VendorCountryEnum = "SO"
	VendorCountryZA VendorCountryEnum = "ZA"
	VendorCountryGS VendorCountryEnum = "GS"
	VendorCountrySS VendorCountryEnum = "SS"
	VendorCountryES VendorCountryEnum = "ES"
	VendorCountryLK VendorCountryEnum = "LK"
	VendorCountryPS VendorCountryEnum = "PS"
	VendorCountrySD VendorCountryEnum = "SD"
	VendorCountrySR VendorCountryEnum = "SR"
	VendorCountrySJ VendorCountryEnum = "SJ"
	VendorCountrySE VendorCountryEnum = "SE"
	VendorCountryCH VendorCountryEnum = "CH"
	VendorCountrySY VendorCountryEnum = "SY"
	VendorCountryTW VendorCountryEnum = "TW"
	VendorCountryTJ VendorCountryEnum = "TJ"
	VendorCountryTH VendorCountryEnum = "TH"
	VendorCountryMK VendorCountryEnum = "MK"
	VendorCountryTL VendorCountryEnum = "TL"
	VendorCountryTG VendorCountryEnum = "TG"
	VendorCountryTK VendorCountryEnum = "TK"
	VendorCountryTO VendorCountryEnum = "TO"
	VendorCountryTT VendorCountryEnum = "TT"
	VendorCountryTN VendorCountryEnum = "TN"
	VendorCountryTR VendorCountryEnum = "TR"
	VendorCountryTM VendorCountryEnum = "TM"
	VendorCountryTC VendorCountryEnum = "TC"
	VendorCountryTV VendorCountryEnum = "TV"
	VendorCountryUG VendorCountryEnum = "UG"
	VendorCountryUA VendorCountryEnum = "UA"
	VendorCountryAE VendorCountryEnum = "AE"
	VendorCountryGB VendorCountryEnum = "GB"
	VendorCountryTZ VendorCountryEnum = "TZ"
	VendorCountryUM VendorCountryEnum = "UM"
	VendorCountryVI VendorCountryEnum = "VI"
	VendorCountryUY VendorCountryEnum = "UY"
	VendorCountryUZ VendorCountryEnum = "UZ"
	VendorCountryVU VendorCountryEnum = "VU"
	VendorCountryVE VendorCountryEnum = "VE"
	VendorCountryVN VendorCountryEnum = "VN"
	VendorCountryWF VendorCountryEnum = "WF"
	VendorCountryEH VendorCountryEnum = "EH"
	VendorCountryYE VendorCountryEnum = "YE"
	VendorCountryZM VendorCountryEnum = "ZM"
	VendorCountryZW VendorCountryEnum = "ZW"
	VendorCountry   VendorCountryEnum = ""
)

type Vendor struct {
	ID          string              `gorm:"primaryKey;not null;column:id" json:"id"`
	CreatedDate datatypes.DateTime  `gorm:"column:createdDate;not null;column:createdDate" json:"createdDate"`
	UpdatedDate *datatypes.DateTime `gorm:"column:updatedDate;column:updatedDate" json:"updatedDate"`
	Meta        *datatypes.JSON     `gorm:"column:meta;not null;column:meta" json:"meta,omitempty"`    // the metadata about the most recent change to the row
	Metadata    *datatypes.JSON     `gorm:"column:metadata;column:metadata" json:"metadata,omitempty"` // metadata reserved for customers to control
	CompanyID   string              `gorm:"not null;column:companyId" json:"companyId"`
	LocationID  string              `gorm:"not null;column:locationId" json:"locationId"`

	AccountNumber    *string            `gorm:"column:accountNumber" json:"accountNumber"`
	Address1         *string            `gorm:"column:address1" json:"address1"`
	Address2         *string            `gorm:"column:address2" json:"address2"`
	City             *string            `gorm:"column:city" json:"city"`
	ContactEmail     *string            `gorm:"column:contactEmail" json:"contactEmail"`
	ContactFirstName *string            `gorm:"column:contactFirstName" json:"contactFirstName"`
	ContactLastName  *string            `gorm:"column:contactLastName" json:"contactLastName"`
	ContactPhone     datatypes.JSON     `gorm:"column:contactPhone" json:"contactPhone"`
	Country          *VendorCountryEnum `gorm:"column:country" json:"country"`
	Name             string             `gorm:"not null;column:name" json:"name"`
	PostalCode       *string            `gorm:"column:postalCode" json:"postalCode"`
	State            *string            `gorm:"column:state" json:"state"`
	URL              *string            `gorm:"column:url" json:"url"`
}

var _ Model = (*Vendor)(nil)

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *Vendor) TableName() string {
	return "vendor"
}

// PrimaryKeys returns an array of the primary keys for this model
func (m *Vendor) PrimaryKeys() []string {
	return []string{"id"}
}

// String returns a string representation as JSON for this model
func (m *Vendor) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewVendor returns a new model instance from an encoded buffer
func NewVendor(buf []byte) (*Vendor, error) {
	var result Vendor
	err := json.Unmarshal(buf, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// NewVendorFromChangeEvent returns a new model instance from an encoded buffer as change event
func NewVendorFromChangeEvent(buf []byte, gzip bool) (*datatypes.ChangeEvent[Vendor], error) {
	var result datatypes.ChangeEvent[Vendor]
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
