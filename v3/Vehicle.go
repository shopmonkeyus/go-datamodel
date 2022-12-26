// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	mapstructure "github.com/mitchellh/mapstructure"
	"time"
)

type VehicleColorEnum string

const (
	VehicleColorWhite    VehicleColorEnum = "White"
	VehicleColorSilver                    = "Silver"
	VehicleColorGray                      = "Gray"
	VehicleColorBlack                     = "Black"
	VehicleColorBlue                      = "Blue"
	VehicleColorRed                       = "Red"
	VehicleColorBrown                     = "Brown"
	VehicleColorBurgundy                  = "Burgundy"
	VehicleColorTan                       = "Tan"
	VehicleColorGold                      = "Gold"
	VehicleColorGreen                     = "Green"
	VehicleColorYellow                    = "Yellow"
	VehicleColorOrange                    = "Orange"
	VehicleColorPink                      = "Pink"
	VehicleColorPurple                    = "Purple"
	VehicleColorOther                     = "Other"
)

type VehicleDrivetrainEnum string

const (
	VehicleDrivetrainRWD    VehicleDrivetrainEnum = "RWD"
	VehicleDrivetrainFWD                          = "FWD"
	VehicleDrivetrainFourWD                       = "FourWD"
	VehicleDrivetrainAWD                          = "AWD"
)

type VehicleLicensePlateCountryEnum string

const (
	VehicleLicensePlateCountryUS VehicleLicensePlateCountryEnum = "US"
	VehicleLicensePlateCountryCA                                = "CA"
	VehicleLicensePlateCountryMX                                = "MX"
	VehicleLicensePlateCountryAF                                = "AF"
	VehicleLicensePlateCountryAX                                = "AX"
	VehicleLicensePlateCountryAL                                = "AL"
	VehicleLicensePlateCountryDZ                                = "DZ"
	VehicleLicensePlateCountryAS                                = "AS"
	VehicleLicensePlateCountryAD                                = "AD"
	VehicleLicensePlateCountryAO                                = "AO"
	VehicleLicensePlateCountryAI                                = "AI"
	VehicleLicensePlateCountryAQ                                = "AQ"
	VehicleLicensePlateCountryAG                                = "AG"
	VehicleLicensePlateCountryAR                                = "AR"
	VehicleLicensePlateCountryAM                                = "AM"
	VehicleLicensePlateCountryAW                                = "AW"
	VehicleLicensePlateCountryAU                                = "AU"
	VehicleLicensePlateCountryAT                                = "AT"
	VehicleLicensePlateCountryAZ                                = "AZ"
	VehicleLicensePlateCountryBS                                = "BS"
	VehicleLicensePlateCountryBH                                = "BH"
	VehicleLicensePlateCountryBD                                = "BD"
	VehicleLicensePlateCountryBB                                = "BB"
	VehicleLicensePlateCountryBY                                = "BY"
	VehicleLicensePlateCountryBE                                = "BE"
	VehicleLicensePlateCountryBZ                                = "BZ"
	VehicleLicensePlateCountryBJ                                = "BJ"
	VehicleLicensePlateCountryBM                                = "BM"
	VehicleLicensePlateCountryBT                                = "BT"
	VehicleLicensePlateCountryBO                                = "BO"
	VehicleLicensePlateCountryBQ                                = "BQ"
	VehicleLicensePlateCountryBA                                = "BA"
	VehicleLicensePlateCountryBW                                = "BW"
	VehicleLicensePlateCountryBV                                = "BV"
	VehicleLicensePlateCountryBR                                = "BR"
	VehicleLicensePlateCountryIO                                = "IO"
	VehicleLicensePlateCountryVG                                = "VG"
	VehicleLicensePlateCountryBN                                = "BN"
	VehicleLicensePlateCountryBG                                = "BG"
	VehicleLicensePlateCountryBF                                = "BF"
	VehicleLicensePlateCountryBI                                = "BI"
	VehicleLicensePlateCountryCV                                = "CV"
	VehicleLicensePlateCountryKH                                = "KH"
	VehicleLicensePlateCountryCM                                = "CM"
	VehicleLicensePlateCountryKY                                = "KY"
	VehicleLicensePlateCountryCF                                = "CF"
	VehicleLicensePlateCountryTD                                = "TD"
	VehicleLicensePlateCountryCL                                = "CL"
	VehicleLicensePlateCountryCN                                = "CN"
	VehicleLicensePlateCountryHK                                = "HK"
	VehicleLicensePlateCountryMO                                = "MO"
	VehicleLicensePlateCountryCX                                = "CX"
	VehicleLicensePlateCountryCC                                = "CC"
	VehicleLicensePlateCountryCO                                = "CO"
	VehicleLicensePlateCountryKM                                = "KM"
	VehicleLicensePlateCountryCG                                = "CG"
	VehicleLicensePlateCountryCK                                = "CK"
	VehicleLicensePlateCountryCR                                = "CR"
	VehicleLicensePlateCountryHR                                = "HR"
	VehicleLicensePlateCountryCU                                = "CU"
	VehicleLicensePlateCountryCW                                = "CW"
	VehicleLicensePlateCountryCY                                = "CY"
	VehicleLicensePlateCountryCZ                                = "CZ"
	VehicleLicensePlateCountryCI                                = "CI"
	VehicleLicensePlateCountryKP                                = "KP"
	VehicleLicensePlateCountryCD                                = "CD"
	VehicleLicensePlateCountryDK                                = "DK"
	VehicleLicensePlateCountryDJ                                = "DJ"
	VehicleLicensePlateCountryDM                                = "DM"
	VehicleLicensePlateCountryDO                                = "DO"
	VehicleLicensePlateCountryEC                                = "EC"
	VehicleLicensePlateCountryEG                                = "EG"
	VehicleLicensePlateCountrySV                                = "SV"
	VehicleLicensePlateCountryGQ                                = "GQ"
	VehicleLicensePlateCountryER                                = "ER"
	VehicleLicensePlateCountryEE                                = "EE"
	VehicleLicensePlateCountrySZ                                = "SZ"
	VehicleLicensePlateCountryET                                = "ET"
	VehicleLicensePlateCountryFK                                = "FK"
	VehicleLicensePlateCountryFO                                = "FO"
	VehicleLicensePlateCountryFJ                                = "FJ"
	VehicleLicensePlateCountryFI                                = "FI"
	VehicleLicensePlateCountryFR                                = "FR"
	VehicleLicensePlateCountryGF                                = "GF"
	VehicleLicensePlateCountryPF                                = "PF"
	VehicleLicensePlateCountryTF                                = "TF"
	VehicleLicensePlateCountryGA                                = "GA"
	VehicleLicensePlateCountryGM                                = "GM"
	VehicleLicensePlateCountryGE                                = "GE"
	VehicleLicensePlateCountryDE                                = "DE"
	VehicleLicensePlateCountryGH                                = "GH"
	VehicleLicensePlateCountryGI                                = "GI"
	VehicleLicensePlateCountryGR                                = "GR"
	VehicleLicensePlateCountryGL                                = "GL"
	VehicleLicensePlateCountryGD                                = "GD"
	VehicleLicensePlateCountryGP                                = "GP"
	VehicleLicensePlateCountryGU                                = "GU"
	VehicleLicensePlateCountryGT                                = "GT"
	VehicleLicensePlateCountryGG                                = "GG"
	VehicleLicensePlateCountryGN                                = "GN"
	VehicleLicensePlateCountryGW                                = "GW"
	VehicleLicensePlateCountryGY                                = "GY"
	VehicleLicensePlateCountryHT                                = "HT"
	VehicleLicensePlateCountryHM                                = "HM"
	VehicleLicensePlateCountryVA                                = "VA"
	VehicleLicensePlateCountryHN                                = "HN"
	VehicleLicensePlateCountryHU                                = "HU"
	VehicleLicensePlateCountryIS                                = "IS"
	VehicleLicensePlateCountryIN                                = "IN"
	VehicleLicensePlateCountryID                                = "ID"
	VehicleLicensePlateCountryIR                                = "IR"
	VehicleLicensePlateCountryIQ                                = "IQ"
	VehicleLicensePlateCountryIE                                = "IE"
	VehicleLicensePlateCountryIM                                = "IM"
	VehicleLicensePlateCountryIL                                = "IL"
	VehicleLicensePlateCountryIT                                = "IT"
	VehicleLicensePlateCountryJM                                = "JM"
	VehicleLicensePlateCountryJP                                = "JP"
	VehicleLicensePlateCountryJE                                = "JE"
	VehicleLicensePlateCountryJO                                = "JO"
	VehicleLicensePlateCountryKZ                                = "KZ"
	VehicleLicensePlateCountryKE                                = "KE"
	VehicleLicensePlateCountryKI                                = "KI"
	VehicleLicensePlateCountryKW                                = "KW"
	VehicleLicensePlateCountryKG                                = "KG"
	VehicleLicensePlateCountryLA                                = "LA"
	VehicleLicensePlateCountryLV                                = "LV"
	VehicleLicensePlateCountryLB                                = "LB"
	VehicleLicensePlateCountryLS                                = "LS"
	VehicleLicensePlateCountryLR                                = "LR"
	VehicleLicensePlateCountryLY                                = "LY"
	VehicleLicensePlateCountryLI                                = "LI"
	VehicleLicensePlateCountryLT                                = "LT"
	VehicleLicensePlateCountryLU                                = "LU"
	VehicleLicensePlateCountryMG                                = "MG"
	VehicleLicensePlateCountryMW                                = "MW"
	VehicleLicensePlateCountryMY                                = "MY"
	VehicleLicensePlateCountryMV                                = "MV"
	VehicleLicensePlateCountryML                                = "ML"
	VehicleLicensePlateCountryMT                                = "MT"
	VehicleLicensePlateCountryMH                                = "MH"
	VehicleLicensePlateCountryMQ                                = "MQ"
	VehicleLicensePlateCountryMR                                = "MR"
	VehicleLicensePlateCountryMU                                = "MU"
	VehicleLicensePlateCountryYT                                = "YT"
	VehicleLicensePlateCountryFM                                = "FM"
	VehicleLicensePlateCountryMC                                = "MC"
	VehicleLicensePlateCountryMN                                = "MN"
	VehicleLicensePlateCountryME                                = "ME"
	VehicleLicensePlateCountryMS                                = "MS"
	VehicleLicensePlateCountryMA                                = "MA"
	VehicleLicensePlateCountryMZ                                = "MZ"
	VehicleLicensePlateCountryMM                                = "MM"
	VehicleLicensePlateCountryNA                                = "NA"
	VehicleLicensePlateCountryNR                                = "NR"
	VehicleLicensePlateCountryNP                                = "NP"
	VehicleLicensePlateCountryNL                                = "NL"
	VehicleLicensePlateCountryNC                                = "NC"
	VehicleLicensePlateCountryNZ                                = "NZ"
	VehicleLicensePlateCountryNI                                = "NI"
	VehicleLicensePlateCountryNE                                = "NE"
	VehicleLicensePlateCountryNG                                = "NG"
	VehicleLicensePlateCountryNU                                = "NU"
	VehicleLicensePlateCountryNF                                = "NF"
	VehicleLicensePlateCountryMP                                = "MP"
	VehicleLicensePlateCountryNO                                = "NO"
	VehicleLicensePlateCountryOM                                = "OM"
	VehicleLicensePlateCountryPK                                = "PK"
	VehicleLicensePlateCountryPW                                = "PW"
	VehicleLicensePlateCountryPA                                = "PA"
	VehicleLicensePlateCountryPG                                = "PG"
	VehicleLicensePlateCountryPY                                = "PY"
	VehicleLicensePlateCountryPE                                = "PE"
	VehicleLicensePlateCountryPH                                = "PH"
	VehicleLicensePlateCountryPN                                = "PN"
	VehicleLicensePlateCountryPL                                = "PL"
	VehicleLicensePlateCountryPT                                = "PT"
	VehicleLicensePlateCountryPR                                = "PR"
	VehicleLicensePlateCountryQA                                = "QA"
	VehicleLicensePlateCountryKR                                = "KR"
	VehicleLicensePlateCountryMD                                = "MD"
	VehicleLicensePlateCountryRO                                = "RO"
	VehicleLicensePlateCountryRU                                = "RU"
	VehicleLicensePlateCountryRW                                = "RW"
	VehicleLicensePlateCountryRE                                = "RE"
	VehicleLicensePlateCountryBL                                = "BL"
	VehicleLicensePlateCountrySH                                = "SH"
	VehicleLicensePlateCountryKN                                = "KN"
	VehicleLicensePlateCountryLC                                = "LC"
	VehicleLicensePlateCountryMF                                = "MF"
	VehicleLicensePlateCountryPM                                = "PM"
	VehicleLicensePlateCountryVC                                = "VC"
	VehicleLicensePlateCountryWS                                = "WS"
	VehicleLicensePlateCountrySM                                = "SM"
	VehicleLicensePlateCountryST                                = "ST"
	VehicleLicensePlateCountrySA                                = "SA"
	VehicleLicensePlateCountrySN                                = "SN"
	VehicleLicensePlateCountryRS                                = "RS"
	VehicleLicensePlateCountrySC                                = "SC"
	VehicleLicensePlateCountrySL                                = "SL"
	VehicleLicensePlateCountrySG                                = "SG"
	VehicleLicensePlateCountrySX                                = "SX"
	VehicleLicensePlateCountrySK                                = "SK"
	VehicleLicensePlateCountrySI                                = "SI"
	VehicleLicensePlateCountrySB                                = "SB"
	VehicleLicensePlateCountrySO                                = "SO"
	VehicleLicensePlateCountryZA                                = "ZA"
	VehicleLicensePlateCountryGS                                = "GS"
	VehicleLicensePlateCountrySS                                = "SS"
	VehicleLicensePlateCountryES                                = "ES"
	VehicleLicensePlateCountryLK                                = "LK"
	VehicleLicensePlateCountryPS                                = "PS"
	VehicleLicensePlateCountrySD                                = "SD"
	VehicleLicensePlateCountrySR                                = "SR"
	VehicleLicensePlateCountrySJ                                = "SJ"
	VehicleLicensePlateCountrySE                                = "SE"
	VehicleLicensePlateCountryCH                                = "CH"
	VehicleLicensePlateCountrySY                                = "SY"
	VehicleLicensePlateCountryTW                                = "TW"
	VehicleLicensePlateCountryTJ                                = "TJ"
	VehicleLicensePlateCountryTH                                = "TH"
	VehicleLicensePlateCountryMK                                = "MK"
	VehicleLicensePlateCountryTL                                = "TL"
	VehicleLicensePlateCountryTG                                = "TG"
	VehicleLicensePlateCountryTK                                = "TK"
	VehicleLicensePlateCountryTO                                = "TO"
	VehicleLicensePlateCountryTT                                = "TT"
	VehicleLicensePlateCountryTN                                = "TN"
	VehicleLicensePlateCountryTR                                = "TR"
	VehicleLicensePlateCountryTM                                = "TM"
	VehicleLicensePlateCountryTC                                = "TC"
	VehicleLicensePlateCountryTV                                = "TV"
	VehicleLicensePlateCountryUG                                = "UG"
	VehicleLicensePlateCountryUA                                = "UA"
	VehicleLicensePlateCountryAE                                = "AE"
	VehicleLicensePlateCountryGB                                = "GB"
	VehicleLicensePlateCountryTZ                                = "TZ"
	VehicleLicensePlateCountryUM                                = "UM"
	VehicleLicensePlateCountryVI                                = "VI"
	VehicleLicensePlateCountryUY                                = "UY"
	VehicleLicensePlateCountryUZ                                = "UZ"
	VehicleLicensePlateCountryVU                                = "VU"
	VehicleLicensePlateCountryVE                                = "VE"
	VehicleLicensePlateCountryVN                                = "VN"
	VehicleLicensePlateCountryWF                                = "WF"
	VehicleLicensePlateCountryEH                                = "EH"
	VehicleLicensePlateCountryYE                                = "YE"
	VehicleLicensePlateCountryZM                                = "ZM"
	VehicleLicensePlateCountryZW                                = "ZW"
)

type VehicleMileageUnitEnum string

const (
	VehicleMileageUnitMile      VehicleMileageUnitEnum = "Mile"
	VehicleMileageUnitKilometer                        = "Kilometer"
)

type VehicleSizeEnum string

const (
	VehicleSizeHeavyDuty VehicleSizeEnum = "HeavyDuty"
	VehicleSizeLightDuty                 = "LightDuty"
	VehicleSizeOther                     = "Other"
)

type VehicleTransmissionEnum string

const (
	VehicleTransmissionAutomatic           VehicleTransmissionEnum = "Automatic"
	VehicleTransmissionManual                                      = "Manual"
	VehicleTransmissionAutomaticCVT                                = "AutomaticCVT"
	VehicleTransmissionAutomaticDualClutch                         = "AutomaticDualClutch"
)

type VehicleTypeEnum string

const (
	VehicleTypeConvertible VehicleTypeEnum = "Convertible"
	VehicleTypeCoupe                       = "Coupe"
	VehicleTypeHatchback                   = "Hatchback"
	VehicleTypeSUV                         = "SUV"
	VehicleTypeSedan                       = "Sedan"
	VehicleTypeTruck                       = "Truck"
	VehicleTypeVan                         = "Van"
	VehicleTypeWagon                       = "Wagon"
	VehicleTypeBicycle                     = "Bicycle"
	VehicleTypeBigRig                      = "BigRig"
	VehicleTypeBoat                        = "Boat"
	VehicleTypeMotorcycle                  = "Motorcycle"
	VehicleTypeTrailer                     = "Trailer"
	VehicleTypeOther                       = "Other"
)

type Vehicle struct {
	AppointmentCount     int64                          `gorm:"not null" json:"appointmentCount"`
	Color                *VehicleColorEnum              `json:"color"`
	CompanyId            string                         `gorm:"not null" json:"companyId"`
	CreatedDate          time.Time                      `gorm:"column:createdDate;not null" json:"createdDate"`
	CustomFields         any                            `gorm:"type:json" json:"customFields"` // custom field values
	DeferredServiceCount int64                          `gorm:"not null" json:"deferredServiceCount"`
	Drivetrain           *VehicleDrivetrainEnum         `json:"drivetrain"`
	Engine               *string                        `json:"engine"`
	ID                   string                         `gorm:"primaryKey;not null" json:"id"`
	LicensePlate         *string                        `json:"licensePlate"`
	LicensePlateCountry  VehicleLicensePlateCountryEnum `gorm:"not null" json:"licensePlateCountry"`
	LicensePlateState    *string                        `json:"licensePlateState"`
	Make                 *string                        `json:"make"`
	MakeId               *int64                         `json:"makeId"` // vcdb make id
	MessageCount         int64                          `gorm:"not null" json:"messageCount"`
	Meta                 *Meta                          `gorm:"type:json;embedded;column:meta;not null" json:"meta,omitempty"` // the metadata about the most recent change to the row
	Metadata             any                            `gorm:"type:json" json:"metadata,omitempty"`                           // metadata reserved for customers to control
	Mileage              *int64                         `json:"mileage"`
	MileageLogCount      int64                          `gorm:"not null" json:"mileageLogCount"`
	MileageUnit          VehicleMileageUnitEnum         `gorm:"not null" json:"mileageUnit"`
	Model                *string                        `json:"model"`
	ModelId              *int64                         `json:"modelId"` // vcdb model id
	Note                 string                         `gorm:"not null" json:"note"`
	Odometer             bool                           `gorm:"not null" json:"odometer"`
	OrderCount           int64                          `gorm:"not null" json:"orderCount"`
	OwnerCount           int64                          `gorm:"not null" json:"ownerCount"`
	ProductionDate       *string                        `json:"productionDate"`
	Size                 VehicleSizeEnum                `gorm:"not null" json:"size"`
	Submodel             *string                        `json:"submodel"`
	SubmodelId           *int64                         `json:"submodelId"` // vcdb submodel id
	TirePressureLogCount int64                          `gorm:"not null" json:"tirePressureLogCount"`
	Transmission         *VehicleTransmissionEnum       `json:"transmission"`
	Type                 *VehicleTypeEnum               `json:"type"`
	Unit                 *string                        `json:"unit"`
	UpdatedDate          *time.Time                     `gorm:"column:updatedDate" json:"updatedDate"`
	VcdbVehicleId        *string                        `json:"vcdbVehicleId"`
	Vin                  *string                        `json:"vin"`
	Year                 *int64                         `json:"year"`
}

var _ Model = (*Vehicle)(nil)

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *Vehicle) TableName() string {
	return "vehicle"
}

func (m *Vehicle) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewVehicle returns a new model instance from a json key/value map
func NewVehicle(kv map[string]any) (*Vehicle, error) {
	var result Vehicle
	err := mapstructure.Decode(kv, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
