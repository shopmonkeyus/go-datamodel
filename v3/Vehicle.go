// Code generated. DO NOT EDIT.
package v3

import (
	"encoding/json"
	datatypes "github.com/shopmonkeyus/go-datamodel/datatypes"
)

type VehicleColorEnum string

const (
	VehicleColorWhite    VehicleColorEnum = "White"
	VehicleColorSilver   VehicleColorEnum = "Silver"
	VehicleColorGray     VehicleColorEnum = "Gray"
	VehicleColorBlack    VehicleColorEnum = "Black"
	VehicleColorBlue     VehicleColorEnum = "Blue"
	VehicleColorRed      VehicleColorEnum = "Red"
	VehicleColorBrown    VehicleColorEnum = "Brown"
	VehicleColorBurgundy VehicleColorEnum = "Burgundy"
	VehicleColorTan      VehicleColorEnum = "Tan"
	VehicleColorGold     VehicleColorEnum = "Gold"
	VehicleColorGreen    VehicleColorEnum = "Green"
	VehicleColorYellow   VehicleColorEnum = "Yellow"
	VehicleColorOrange   VehicleColorEnum = "Orange"
	VehicleColorPink     VehicleColorEnum = "Pink"
	VehicleColorPurple   VehicleColorEnum = "Purple"
	VehicleColorOther    VehicleColorEnum = "Other"
	VehicleColor         VehicleColorEnum = ""
)

type VehicleDrivetrainEnum string

const (
	VehicleDrivetrainRWD    VehicleDrivetrainEnum = "RWD"
	VehicleDrivetrainFWD    VehicleDrivetrainEnum = "FWD"
	VehicleDrivetrainFourWD VehicleDrivetrainEnum = "FourWD"
	VehicleDrivetrainAWD    VehicleDrivetrainEnum = "AWD"
	VehicleDrivetrain       VehicleDrivetrainEnum = ""
)

type VehicleLicensePlateCountryEnum string

const (
	VehicleLicensePlateCountryUS VehicleLicensePlateCountryEnum = "US"
	VehicleLicensePlateCountryCA VehicleLicensePlateCountryEnum = "CA"
	VehicleLicensePlateCountryMX VehicleLicensePlateCountryEnum = "MX"
	VehicleLicensePlateCountryAF VehicleLicensePlateCountryEnum = "AF"
	VehicleLicensePlateCountryAX VehicleLicensePlateCountryEnum = "AX"
	VehicleLicensePlateCountryAL VehicleLicensePlateCountryEnum = "AL"
	VehicleLicensePlateCountryDZ VehicleLicensePlateCountryEnum = "DZ"
	VehicleLicensePlateCountryAS VehicleLicensePlateCountryEnum = "AS"
	VehicleLicensePlateCountryAD VehicleLicensePlateCountryEnum = "AD"
	VehicleLicensePlateCountryAO VehicleLicensePlateCountryEnum = "AO"
	VehicleLicensePlateCountryAI VehicleLicensePlateCountryEnum = "AI"
	VehicleLicensePlateCountryAQ VehicleLicensePlateCountryEnum = "AQ"
	VehicleLicensePlateCountryAG VehicleLicensePlateCountryEnum = "AG"
	VehicleLicensePlateCountryAR VehicleLicensePlateCountryEnum = "AR"
	VehicleLicensePlateCountryAM VehicleLicensePlateCountryEnum = "AM"
	VehicleLicensePlateCountryAW VehicleLicensePlateCountryEnum = "AW"
	VehicleLicensePlateCountryAU VehicleLicensePlateCountryEnum = "AU"
	VehicleLicensePlateCountryAT VehicleLicensePlateCountryEnum = "AT"
	VehicleLicensePlateCountryAZ VehicleLicensePlateCountryEnum = "AZ"
	VehicleLicensePlateCountryBS VehicleLicensePlateCountryEnum = "BS"
	VehicleLicensePlateCountryBH VehicleLicensePlateCountryEnum = "BH"
	VehicleLicensePlateCountryBD VehicleLicensePlateCountryEnum = "BD"
	VehicleLicensePlateCountryBB VehicleLicensePlateCountryEnum = "BB"
	VehicleLicensePlateCountryBY VehicleLicensePlateCountryEnum = "BY"
	VehicleLicensePlateCountryBE VehicleLicensePlateCountryEnum = "BE"
	VehicleLicensePlateCountryBZ VehicleLicensePlateCountryEnum = "BZ"
	VehicleLicensePlateCountryBJ VehicleLicensePlateCountryEnum = "BJ"
	VehicleLicensePlateCountryBM VehicleLicensePlateCountryEnum = "BM"
	VehicleLicensePlateCountryBT VehicleLicensePlateCountryEnum = "BT"
	VehicleLicensePlateCountryBO VehicleLicensePlateCountryEnum = "BO"
	VehicleLicensePlateCountryBQ VehicleLicensePlateCountryEnum = "BQ"
	VehicleLicensePlateCountryBA VehicleLicensePlateCountryEnum = "BA"
	VehicleLicensePlateCountryBW VehicleLicensePlateCountryEnum = "BW"
	VehicleLicensePlateCountryBV VehicleLicensePlateCountryEnum = "BV"
	VehicleLicensePlateCountryBR VehicleLicensePlateCountryEnum = "BR"
	VehicleLicensePlateCountryIO VehicleLicensePlateCountryEnum = "IO"
	VehicleLicensePlateCountryVG VehicleLicensePlateCountryEnum = "VG"
	VehicleLicensePlateCountryBN VehicleLicensePlateCountryEnum = "BN"
	VehicleLicensePlateCountryBG VehicleLicensePlateCountryEnum = "BG"
	VehicleLicensePlateCountryBF VehicleLicensePlateCountryEnum = "BF"
	VehicleLicensePlateCountryBI VehicleLicensePlateCountryEnum = "BI"
	VehicleLicensePlateCountryCV VehicleLicensePlateCountryEnum = "CV"
	VehicleLicensePlateCountryKH VehicleLicensePlateCountryEnum = "KH"
	VehicleLicensePlateCountryCM VehicleLicensePlateCountryEnum = "CM"
	VehicleLicensePlateCountryKY VehicleLicensePlateCountryEnum = "KY"
	VehicleLicensePlateCountryCF VehicleLicensePlateCountryEnum = "CF"
	VehicleLicensePlateCountryTD VehicleLicensePlateCountryEnum = "TD"
	VehicleLicensePlateCountryCL VehicleLicensePlateCountryEnum = "CL"
	VehicleLicensePlateCountryCN VehicleLicensePlateCountryEnum = "CN"
	VehicleLicensePlateCountryHK VehicleLicensePlateCountryEnum = "HK"
	VehicleLicensePlateCountryMO VehicleLicensePlateCountryEnum = "MO"
	VehicleLicensePlateCountryCX VehicleLicensePlateCountryEnum = "CX"
	VehicleLicensePlateCountryCC VehicleLicensePlateCountryEnum = "CC"
	VehicleLicensePlateCountryCO VehicleLicensePlateCountryEnum = "CO"
	VehicleLicensePlateCountryKM VehicleLicensePlateCountryEnum = "KM"
	VehicleLicensePlateCountryCG VehicleLicensePlateCountryEnum = "CG"
	VehicleLicensePlateCountryCK VehicleLicensePlateCountryEnum = "CK"
	VehicleLicensePlateCountryCR VehicleLicensePlateCountryEnum = "CR"
	VehicleLicensePlateCountryHR VehicleLicensePlateCountryEnum = "HR"
	VehicleLicensePlateCountryCU VehicleLicensePlateCountryEnum = "CU"
	VehicleLicensePlateCountryCW VehicleLicensePlateCountryEnum = "CW"
	VehicleLicensePlateCountryCY VehicleLicensePlateCountryEnum = "CY"
	VehicleLicensePlateCountryCZ VehicleLicensePlateCountryEnum = "CZ"
	VehicleLicensePlateCountryCI VehicleLicensePlateCountryEnum = "CI"
	VehicleLicensePlateCountryKP VehicleLicensePlateCountryEnum = "KP"
	VehicleLicensePlateCountryCD VehicleLicensePlateCountryEnum = "CD"
	VehicleLicensePlateCountryDK VehicleLicensePlateCountryEnum = "DK"
	VehicleLicensePlateCountryDJ VehicleLicensePlateCountryEnum = "DJ"
	VehicleLicensePlateCountryDM VehicleLicensePlateCountryEnum = "DM"
	VehicleLicensePlateCountryDO VehicleLicensePlateCountryEnum = "DO"
	VehicleLicensePlateCountryEC VehicleLicensePlateCountryEnum = "EC"
	VehicleLicensePlateCountryEG VehicleLicensePlateCountryEnum = "EG"
	VehicleLicensePlateCountrySV VehicleLicensePlateCountryEnum = "SV"
	VehicleLicensePlateCountryGQ VehicleLicensePlateCountryEnum = "GQ"
	VehicleLicensePlateCountryER VehicleLicensePlateCountryEnum = "ER"
	VehicleLicensePlateCountryEE VehicleLicensePlateCountryEnum = "EE"
	VehicleLicensePlateCountrySZ VehicleLicensePlateCountryEnum = "SZ"
	VehicleLicensePlateCountryET VehicleLicensePlateCountryEnum = "ET"
	VehicleLicensePlateCountryFK VehicleLicensePlateCountryEnum = "FK"
	VehicleLicensePlateCountryFO VehicleLicensePlateCountryEnum = "FO"
	VehicleLicensePlateCountryFJ VehicleLicensePlateCountryEnum = "FJ"
	VehicleLicensePlateCountryFI VehicleLicensePlateCountryEnum = "FI"
	VehicleLicensePlateCountryFR VehicleLicensePlateCountryEnum = "FR"
	VehicleLicensePlateCountryGF VehicleLicensePlateCountryEnum = "GF"
	VehicleLicensePlateCountryPF VehicleLicensePlateCountryEnum = "PF"
	VehicleLicensePlateCountryTF VehicleLicensePlateCountryEnum = "TF"
	VehicleLicensePlateCountryGA VehicleLicensePlateCountryEnum = "GA"
	VehicleLicensePlateCountryGM VehicleLicensePlateCountryEnum = "GM"
	VehicleLicensePlateCountryGE VehicleLicensePlateCountryEnum = "GE"
	VehicleLicensePlateCountryDE VehicleLicensePlateCountryEnum = "DE"
	VehicleLicensePlateCountryGH VehicleLicensePlateCountryEnum = "GH"
	VehicleLicensePlateCountryGI VehicleLicensePlateCountryEnum = "GI"
	VehicleLicensePlateCountryGR VehicleLicensePlateCountryEnum = "GR"
	VehicleLicensePlateCountryGL VehicleLicensePlateCountryEnum = "GL"
	VehicleLicensePlateCountryGD VehicleLicensePlateCountryEnum = "GD"
	VehicleLicensePlateCountryGP VehicleLicensePlateCountryEnum = "GP"
	VehicleLicensePlateCountryGU VehicleLicensePlateCountryEnum = "GU"
	VehicleLicensePlateCountryGT VehicleLicensePlateCountryEnum = "GT"
	VehicleLicensePlateCountryGG VehicleLicensePlateCountryEnum = "GG"
	VehicleLicensePlateCountryGN VehicleLicensePlateCountryEnum = "GN"
	VehicleLicensePlateCountryGW VehicleLicensePlateCountryEnum = "GW"
	VehicleLicensePlateCountryGY VehicleLicensePlateCountryEnum = "GY"
	VehicleLicensePlateCountryHT VehicleLicensePlateCountryEnum = "HT"
	VehicleLicensePlateCountryHM VehicleLicensePlateCountryEnum = "HM"
	VehicleLicensePlateCountryVA VehicleLicensePlateCountryEnum = "VA"
	VehicleLicensePlateCountryHN VehicleLicensePlateCountryEnum = "HN"
	VehicleLicensePlateCountryHU VehicleLicensePlateCountryEnum = "HU"
	VehicleLicensePlateCountryIS VehicleLicensePlateCountryEnum = "IS"
	VehicleLicensePlateCountryIN VehicleLicensePlateCountryEnum = "IN"
	VehicleLicensePlateCountryID VehicleLicensePlateCountryEnum = "ID"
	VehicleLicensePlateCountryIR VehicleLicensePlateCountryEnum = "IR"
	VehicleLicensePlateCountryIQ VehicleLicensePlateCountryEnum = "IQ"
	VehicleLicensePlateCountryIE VehicleLicensePlateCountryEnum = "IE"
	VehicleLicensePlateCountryIM VehicleLicensePlateCountryEnum = "IM"
	VehicleLicensePlateCountryIL VehicleLicensePlateCountryEnum = "IL"
	VehicleLicensePlateCountryIT VehicleLicensePlateCountryEnum = "IT"
	VehicleLicensePlateCountryJM VehicleLicensePlateCountryEnum = "JM"
	VehicleLicensePlateCountryJP VehicleLicensePlateCountryEnum = "JP"
	VehicleLicensePlateCountryJE VehicleLicensePlateCountryEnum = "JE"
	VehicleLicensePlateCountryJO VehicleLicensePlateCountryEnum = "JO"
	VehicleLicensePlateCountryKZ VehicleLicensePlateCountryEnum = "KZ"
	VehicleLicensePlateCountryKE VehicleLicensePlateCountryEnum = "KE"
	VehicleLicensePlateCountryKI VehicleLicensePlateCountryEnum = "KI"
	VehicleLicensePlateCountryKW VehicleLicensePlateCountryEnum = "KW"
	VehicleLicensePlateCountryKG VehicleLicensePlateCountryEnum = "KG"
	VehicleLicensePlateCountryLA VehicleLicensePlateCountryEnum = "LA"
	VehicleLicensePlateCountryLV VehicleLicensePlateCountryEnum = "LV"
	VehicleLicensePlateCountryLB VehicleLicensePlateCountryEnum = "LB"
	VehicleLicensePlateCountryLS VehicleLicensePlateCountryEnum = "LS"
	VehicleLicensePlateCountryLR VehicleLicensePlateCountryEnum = "LR"
	VehicleLicensePlateCountryLY VehicleLicensePlateCountryEnum = "LY"
	VehicleLicensePlateCountryLI VehicleLicensePlateCountryEnum = "LI"
	VehicleLicensePlateCountryLT VehicleLicensePlateCountryEnum = "LT"
	VehicleLicensePlateCountryLU VehicleLicensePlateCountryEnum = "LU"
	VehicleLicensePlateCountryMG VehicleLicensePlateCountryEnum = "MG"
	VehicleLicensePlateCountryMW VehicleLicensePlateCountryEnum = "MW"
	VehicleLicensePlateCountryMY VehicleLicensePlateCountryEnum = "MY"
	VehicleLicensePlateCountryMV VehicleLicensePlateCountryEnum = "MV"
	VehicleLicensePlateCountryML VehicleLicensePlateCountryEnum = "ML"
	VehicleLicensePlateCountryMT VehicleLicensePlateCountryEnum = "MT"
	VehicleLicensePlateCountryMH VehicleLicensePlateCountryEnum = "MH"
	VehicleLicensePlateCountryMQ VehicleLicensePlateCountryEnum = "MQ"
	VehicleLicensePlateCountryMR VehicleLicensePlateCountryEnum = "MR"
	VehicleLicensePlateCountryMU VehicleLicensePlateCountryEnum = "MU"
	VehicleLicensePlateCountryYT VehicleLicensePlateCountryEnum = "YT"
	VehicleLicensePlateCountryFM VehicleLicensePlateCountryEnum = "FM"
	VehicleLicensePlateCountryMC VehicleLicensePlateCountryEnum = "MC"
	VehicleLicensePlateCountryMN VehicleLicensePlateCountryEnum = "MN"
	VehicleLicensePlateCountryME VehicleLicensePlateCountryEnum = "ME"
	VehicleLicensePlateCountryMS VehicleLicensePlateCountryEnum = "MS"
	VehicleLicensePlateCountryMA VehicleLicensePlateCountryEnum = "MA"
	VehicleLicensePlateCountryMZ VehicleLicensePlateCountryEnum = "MZ"
	VehicleLicensePlateCountryMM VehicleLicensePlateCountryEnum = "MM"
	VehicleLicensePlateCountryNA VehicleLicensePlateCountryEnum = "NA"
	VehicleLicensePlateCountryNR VehicleLicensePlateCountryEnum = "NR"
	VehicleLicensePlateCountryNP VehicleLicensePlateCountryEnum = "NP"
	VehicleLicensePlateCountryNL VehicleLicensePlateCountryEnum = "NL"
	VehicleLicensePlateCountryNC VehicleLicensePlateCountryEnum = "NC"
	VehicleLicensePlateCountryNZ VehicleLicensePlateCountryEnum = "NZ"
	VehicleLicensePlateCountryNI VehicleLicensePlateCountryEnum = "NI"
	VehicleLicensePlateCountryNE VehicleLicensePlateCountryEnum = "NE"
	VehicleLicensePlateCountryNG VehicleLicensePlateCountryEnum = "NG"
	VehicleLicensePlateCountryNU VehicleLicensePlateCountryEnum = "NU"
	VehicleLicensePlateCountryNF VehicleLicensePlateCountryEnum = "NF"
	VehicleLicensePlateCountryMP VehicleLicensePlateCountryEnum = "MP"
	VehicleLicensePlateCountryNO VehicleLicensePlateCountryEnum = "NO"
	VehicleLicensePlateCountryOM VehicleLicensePlateCountryEnum = "OM"
	VehicleLicensePlateCountryPK VehicleLicensePlateCountryEnum = "PK"
	VehicleLicensePlateCountryPW VehicleLicensePlateCountryEnum = "PW"
	VehicleLicensePlateCountryPA VehicleLicensePlateCountryEnum = "PA"
	VehicleLicensePlateCountryPG VehicleLicensePlateCountryEnum = "PG"
	VehicleLicensePlateCountryPY VehicleLicensePlateCountryEnum = "PY"
	VehicleLicensePlateCountryPE VehicleLicensePlateCountryEnum = "PE"
	VehicleLicensePlateCountryPH VehicleLicensePlateCountryEnum = "PH"
	VehicleLicensePlateCountryPN VehicleLicensePlateCountryEnum = "PN"
	VehicleLicensePlateCountryPL VehicleLicensePlateCountryEnum = "PL"
	VehicleLicensePlateCountryPT VehicleLicensePlateCountryEnum = "PT"
	VehicleLicensePlateCountryPR VehicleLicensePlateCountryEnum = "PR"
	VehicleLicensePlateCountryQA VehicleLicensePlateCountryEnum = "QA"
	VehicleLicensePlateCountryKR VehicleLicensePlateCountryEnum = "KR"
	VehicleLicensePlateCountryMD VehicleLicensePlateCountryEnum = "MD"
	VehicleLicensePlateCountryRO VehicleLicensePlateCountryEnum = "RO"
	VehicleLicensePlateCountryRU VehicleLicensePlateCountryEnum = "RU"
	VehicleLicensePlateCountryRW VehicleLicensePlateCountryEnum = "RW"
	VehicleLicensePlateCountryRE VehicleLicensePlateCountryEnum = "RE"
	VehicleLicensePlateCountryBL VehicleLicensePlateCountryEnum = "BL"
	VehicleLicensePlateCountrySH VehicleLicensePlateCountryEnum = "SH"
	VehicleLicensePlateCountryKN VehicleLicensePlateCountryEnum = "KN"
	VehicleLicensePlateCountryLC VehicleLicensePlateCountryEnum = "LC"
	VehicleLicensePlateCountryMF VehicleLicensePlateCountryEnum = "MF"
	VehicleLicensePlateCountryPM VehicleLicensePlateCountryEnum = "PM"
	VehicleLicensePlateCountryVC VehicleLicensePlateCountryEnum = "VC"
	VehicleLicensePlateCountryWS VehicleLicensePlateCountryEnum = "WS"
	VehicleLicensePlateCountrySM VehicleLicensePlateCountryEnum = "SM"
	VehicleLicensePlateCountryST VehicleLicensePlateCountryEnum = "ST"
	VehicleLicensePlateCountrySA VehicleLicensePlateCountryEnum = "SA"
	VehicleLicensePlateCountrySN VehicleLicensePlateCountryEnum = "SN"
	VehicleLicensePlateCountryRS VehicleLicensePlateCountryEnum = "RS"
	VehicleLicensePlateCountrySC VehicleLicensePlateCountryEnum = "SC"
	VehicleLicensePlateCountrySL VehicleLicensePlateCountryEnum = "SL"
	VehicleLicensePlateCountrySG VehicleLicensePlateCountryEnum = "SG"
	VehicleLicensePlateCountrySX VehicleLicensePlateCountryEnum = "SX"
	VehicleLicensePlateCountrySK VehicleLicensePlateCountryEnum = "SK"
	VehicleLicensePlateCountrySI VehicleLicensePlateCountryEnum = "SI"
	VehicleLicensePlateCountrySB VehicleLicensePlateCountryEnum = "SB"
	VehicleLicensePlateCountrySO VehicleLicensePlateCountryEnum = "SO"
	VehicleLicensePlateCountryZA VehicleLicensePlateCountryEnum = "ZA"
	VehicleLicensePlateCountryGS VehicleLicensePlateCountryEnum = "GS"
	VehicleLicensePlateCountrySS VehicleLicensePlateCountryEnum = "SS"
	VehicleLicensePlateCountryES VehicleLicensePlateCountryEnum = "ES"
	VehicleLicensePlateCountryLK VehicleLicensePlateCountryEnum = "LK"
	VehicleLicensePlateCountryPS VehicleLicensePlateCountryEnum = "PS"
	VehicleLicensePlateCountrySD VehicleLicensePlateCountryEnum = "SD"
	VehicleLicensePlateCountrySR VehicleLicensePlateCountryEnum = "SR"
	VehicleLicensePlateCountrySJ VehicleLicensePlateCountryEnum = "SJ"
	VehicleLicensePlateCountrySE VehicleLicensePlateCountryEnum = "SE"
	VehicleLicensePlateCountryCH VehicleLicensePlateCountryEnum = "CH"
	VehicleLicensePlateCountrySY VehicleLicensePlateCountryEnum = "SY"
	VehicleLicensePlateCountryTW VehicleLicensePlateCountryEnum = "TW"
	VehicleLicensePlateCountryTJ VehicleLicensePlateCountryEnum = "TJ"
	VehicleLicensePlateCountryTH VehicleLicensePlateCountryEnum = "TH"
	VehicleLicensePlateCountryMK VehicleLicensePlateCountryEnum = "MK"
	VehicleLicensePlateCountryTL VehicleLicensePlateCountryEnum = "TL"
	VehicleLicensePlateCountryTG VehicleLicensePlateCountryEnum = "TG"
	VehicleLicensePlateCountryTK VehicleLicensePlateCountryEnum = "TK"
	VehicleLicensePlateCountryTO VehicleLicensePlateCountryEnum = "TO"
	VehicleLicensePlateCountryTT VehicleLicensePlateCountryEnum = "TT"
	VehicleLicensePlateCountryTN VehicleLicensePlateCountryEnum = "TN"
	VehicleLicensePlateCountryTR VehicleLicensePlateCountryEnum = "TR"
	VehicleLicensePlateCountryTM VehicleLicensePlateCountryEnum = "TM"
	VehicleLicensePlateCountryTC VehicleLicensePlateCountryEnum = "TC"
	VehicleLicensePlateCountryTV VehicleLicensePlateCountryEnum = "TV"
	VehicleLicensePlateCountryUG VehicleLicensePlateCountryEnum = "UG"
	VehicleLicensePlateCountryUA VehicleLicensePlateCountryEnum = "UA"
	VehicleLicensePlateCountryAE VehicleLicensePlateCountryEnum = "AE"
	VehicleLicensePlateCountryGB VehicleLicensePlateCountryEnum = "GB"
	VehicleLicensePlateCountryTZ VehicleLicensePlateCountryEnum = "TZ"
	VehicleLicensePlateCountryUM VehicleLicensePlateCountryEnum = "UM"
	VehicleLicensePlateCountryVI VehicleLicensePlateCountryEnum = "VI"
	VehicleLicensePlateCountryUY VehicleLicensePlateCountryEnum = "UY"
	VehicleLicensePlateCountryUZ VehicleLicensePlateCountryEnum = "UZ"
	VehicleLicensePlateCountryVU VehicleLicensePlateCountryEnum = "VU"
	VehicleLicensePlateCountryVE VehicleLicensePlateCountryEnum = "VE"
	VehicleLicensePlateCountryVN VehicleLicensePlateCountryEnum = "VN"
	VehicleLicensePlateCountryWF VehicleLicensePlateCountryEnum = "WF"
	VehicleLicensePlateCountryEH VehicleLicensePlateCountryEnum = "EH"
	VehicleLicensePlateCountryYE VehicleLicensePlateCountryEnum = "YE"
	VehicleLicensePlateCountryZM VehicleLicensePlateCountryEnum = "ZM"
	VehicleLicensePlateCountryZW VehicleLicensePlateCountryEnum = "ZW"
)

type VehicleMileageUnitEnum string

const (
	VehicleMileageUnitMile      VehicleMileageUnitEnum = "Mile"
	VehicleMileageUnitKilometer VehicleMileageUnitEnum = "Kilometer"
)

type VehicleSizeEnum string

const (
	VehicleSizeHeavyDuty VehicleSizeEnum = "HeavyDuty"
	VehicleSizeLightDuty VehicleSizeEnum = "LightDuty"
	VehicleSizeOther     VehicleSizeEnum = "Other"
)

type VehicleTransmissionEnum string

const (
	VehicleTransmissionAutomatic           VehicleTransmissionEnum = "Automatic"
	VehicleTransmissionManual              VehicleTransmissionEnum = "Manual"
	VehicleTransmissionAutomaticCVT        VehicleTransmissionEnum = "AutomaticCVT"
	VehicleTransmissionAutomaticDualClutch VehicleTransmissionEnum = "AutomaticDualClutch"
	VehicleTransmission                    VehicleTransmissionEnum = ""
)

type VehicleTypeEnum string

const (
	VehicleTypeConvertible VehicleTypeEnum = "Convertible"
	VehicleTypeCoupe       VehicleTypeEnum = "Coupe"
	VehicleTypeHatchback   VehicleTypeEnum = "Hatchback"
	VehicleTypeSUV         VehicleTypeEnum = "SUV"
	VehicleTypeSedan       VehicleTypeEnum = "Sedan"
	VehicleTypeTruck       VehicleTypeEnum = "Truck"
	VehicleTypeVan         VehicleTypeEnum = "Van"
	VehicleTypeWagon       VehicleTypeEnum = "Wagon"
	VehicleTypeBicycle     VehicleTypeEnum = "Bicycle"
	VehicleTypeBigRig      VehicleTypeEnum = "BigRig"
	VehicleTypeBoat        VehicleTypeEnum = "Boat"
	VehicleTypeMotorcycle  VehicleTypeEnum = "Motorcycle"
	VehicleTypeTrailer     VehicleTypeEnum = "Trailer"
	VehicleTypeOther       VehicleTypeEnum = "Other"
	VehicleType            VehicleTypeEnum = ""
)

type Vehicle struct {
	ID           string              `gorm:"primaryKey;not null;column:id" json:"id"`
	CreatedDate  datatypes.DateTime  `gorm:"column:createdDate;not null;column:createdDate" json:"createdDate"`
	UpdatedDate  *datatypes.DateTime `gorm:"column:updatedDate;column:updatedDate" json:"updatedDate"`
	Meta         *datatypes.JSON     `gorm:"column:meta;not null;column:meta" json:"meta,omitempty"`    // the metadata about the most recent change to the row
	Metadata     *datatypes.JSON     `gorm:"column:metadata;column:metadata" json:"metadata,omitempty"` // metadata reserved for customers to control
	CompanyID    string              `gorm:"not null;column:companyId" json:"companyId"`
	CustomFields datatypes.JSON      `gorm:"column:customFields" json:"customFields"` // custom field values

	AppointmentCount     int64                          `gorm:"not null;column:appointmentCount" json:"appointmentCount"`
	Color                *VehicleColorEnum              `gorm:"column:color" json:"color"`
	DeferredServiceCount int64                          `gorm:"not null;column:deferredServiceCount" json:"deferredServiceCount"`
	Drivetrain           *VehicleDrivetrainEnum         `gorm:"column:drivetrain" json:"drivetrain"`
	Engine               *string                        `gorm:"column:engine" json:"engine"`
	LicensePlate         *string                        `gorm:"column:licensePlate" json:"licensePlate"`
	LicensePlateCountry  VehicleLicensePlateCountryEnum `gorm:"not null;column:licensePlateCountry" json:"licensePlateCountry"`
	LicensePlateState    *string                        `gorm:"column:licensePlateState" json:"licensePlateState"`
	Make                 *string                        `gorm:"column:make" json:"make"`
	MakeID               *int64                         `gorm:"column:makeId" json:"makeId"` // vcdb make id
	MessageCount         int64                          `gorm:"not null;column:messageCount" json:"messageCount"`
	Mileage              *int64                         `gorm:"column:mileage" json:"mileage"`
	MileageLogCount      int64                          `gorm:"not null;column:mileageLogCount" json:"mileageLogCount"`
	MileageUnit          VehicleMileageUnitEnum         `gorm:"not null;column:mileageUnit" json:"mileageUnit"`
	Model                *string                        `gorm:"column:model" json:"model"`
	ModelID              *int64                         `gorm:"column:modelId" json:"modelId"` // vcdb model id
	Note                 string                         `gorm:"not null;column:note" json:"note"`
	Odometer             bool                           `gorm:"not null;column:odometer" json:"odometer"`
	OrderCount           int64                          `gorm:"not null;column:orderCount" json:"orderCount"`
	OwnerCount           int64                          `gorm:"not null;column:ownerCount" json:"ownerCount"`
	ProductionDate       *string                        `gorm:"column:productionDate" json:"productionDate"`
	Size                 VehicleSizeEnum                `gorm:"not null;column:size" json:"size"`
	Submodel             *string                        `gorm:"column:submodel" json:"submodel"`
	SubmodelID           *int64                         `gorm:"column:submodelId" json:"submodelId"` // vcdb submodel id
	TirePressureLogCount int64                          `gorm:"not null;column:tirePressureLogCount" json:"tirePressureLogCount"`
	Transmission         *VehicleTransmissionEnum       `gorm:"column:transmission" json:"transmission"`
	Type                 *VehicleTypeEnum               `gorm:"column:type" json:"type"`
	Unit                 *string                        `gorm:"column:unit" json:"unit"`
	VcdbVehicleID        *string                        `gorm:"column:vcdbVehicleId" json:"vcdbVehicleId"`
	Vin                  *string                        `gorm:"column:vin" json:"vin"`
	Year                 *int64                         `gorm:"column:year" json:"year"`
}

var _ Model = (*Vehicle)(nil)

// TableName returns the name of the table for this model which GORM will use when using this model
func (m *Vehicle) TableName() string {
	return "vehicle"
}

// PrimaryKeys returns an array of the primary keys for this model
func (m *Vehicle) PrimaryKeys() []string {
	return []string{"id"}
}

// String returns a string representation as JSON for this model
func (m *Vehicle) String() string {
	buf, _ := json.Marshal(m)
	return string(buf)
}

// NewVehicle returns a new model instance from an encoded buffer
func NewVehicle(buf []byte) (*Vehicle, error) {
	var result Vehicle
	err := json.Unmarshal(buf, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// NewVehicleFromChangeEvent returns a new model instance from an encoded buffer as change event
func NewVehicleFromChangeEvent(buf []byte, gzip bool) (*datatypes.ChangeEvent[Vehicle], error) {
	var result datatypes.ChangeEvent[Vehicle]
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
