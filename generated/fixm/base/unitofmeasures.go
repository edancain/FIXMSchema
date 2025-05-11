// Code generated from UnitOfMeasures.xsd; DO NOT EDIT.

package base

// The reference quantities used to express the value of airspeed.
type UomAirspeedType string

const (
	UomAirspeedTypeKM_H UomAirspeedType = "KM_H"
	UomAirspeedTypeKT UomAirspeedType = "KT"
	UomAirspeedTypeMACH UomAirspeedType = "MACH"
)

// The reference quantities used to express the value of altitude.
type UomAltitudeType string

const (
	UomAltitudeTypeFT UomAltitudeType = "FT"
	UomAltitudeTypeM UomAltitudeType = "M"
)

// The reference quantities used to express the value of angles. [ISO 19103, chapter 6.5.7.10]
type UomAngleType string

const (
	UomAngleTypeDEG UomAngleType = "DEG" //Degree
)

// The reference quantities used to express the value of flight level.
type UomFlightLevelType string

const (
	UomFlightLevelTypeFL UomFlightLevelType = "FL" // Flight level in hundreds of feet.
	UomFlightLevelTypeSM UomFlightLevelType = "SM" // Flight level in tens of metres.
)

// Radio frequency unit of measure. Either kHz OR MHz.
type UomFrequencyType string

const (
	UomFrequencyTypeKHZ UomFrequencyType = "KHZ" // Indicates this radio frequency is measured in kHz.
	UomFrequencyTypeMHZ UomFrequencyType = "MHZ" // Indicates this radio frequency is measured in MHz.
)

// The reference quantities used to express the value of ground speed.
type UomGroundSpeedType string

const (
	UomGroundSpeedTypeKM_H UomGroundSpeedType = "KM_H" // Kilometres per hour.
	UomGroundSpeedTypeKT UomGroundSpeedType = "KT" // Knots.
)

// The reference quantities used to express the value of height.
type UomHeightType string

const (
	UomHeightTypeFT UomHeightType = "FT" // Foot.
	UomHeightTypeM UomHeightType = "M" // Metre.
)

// The reference quantities used to express the value of the length, or distance between two entities. [ISO 19103, chapter 6.5.7.8]
type UomLengthType string

const (
	UomLengthTypeCM UomLengthType = "CM" // Centimetre
	UomLengthTypeFT UomLengthType = "FT" // Foot
	UomLengthTypeIN UomLengthType = "IN" // Inch
	UomLengthTypeKM UomLengthType = "KM" // Kilometre
	UomLengthTypeM UomLengthType = "M" // Metre
	UomLengthTypeMI UomLengthType = "MI" // Statute Mile
	UomLengthTypeMM UomLengthType = "MM" // Millimetre
	UomLengthTypeNM UomLengthType = "NM" // Nautical Mile
)

// The reference quantities used to express the value of mass.
type UomMassType string

const (
	UomMassTypeKG UomMassType = "KG" // Kilogram
	UomMassTypeLB UomMassType = "LB" // Pound
)

// The reference quantities used to express the value of pressure.
type UomPressureType string

const (
	UomPressureTypeATM UomPressureType = "ATM" // Standard Atmosphere.
	UomPressureTypeBAR UomPressureType = "BAR" // Bar. (=100000 Pascals (Pa))
	UomPressureTypeHPA UomPressureType = "HPA" // Hectopascal
	UomPressureTypeINHG UomPressureType = "INHG" // Inches of mercury
	UomPressureTypeMBAR UomPressureType = "MBAR" // Millibar
	UomPressureTypePA UomPressureType = "PA" // Pascal
	UomPressureTypePSI UomPressureType = "PSI" // Pounds per square inch.
	UomPressureTypeTORR UomPressureType = "TORR" // Torr. (= mm of mercury (mmHg).)
)

// The reference quantities used to express the value of temperature.
type UomTemperatureType string

const (
	UomTemperatureTypeC UomTemperatureType = "C" // Degrees Celsius.
	UomTemperatureTypeF UomTemperatureType = "F" // Degrees Fahrenheit.
	UomTemperatureTypeK UomTemperatureType = "K" // Degrees Kelvin.
	UomTemperatureTypeR UomTemperatureType = "R" // Degrees Rankine.
)

// The reference quantities used to express the value of vertical rate.
type UomVerticalRateType string

const (
	UomVerticalRateTypeFT_MIN UomVerticalRateType = "FT_MIN" // Feet per minute.
	UomVerticalRateTypeM_SEC UomVerticalRateType = "M_SEC" // Metres per second.
)

// The reference quantities used to express the value of volume. [ISO 19103, chapter 6.5.7.16]
type UomVolumeType string

const (
	UomVolumeTypeL UomVolumeType = "L" // Litre
	UomVolumeTypeUS_GAL UomVolumeType = "US_GAL" // US Gallon
)

// The reference quantities used to express the value of weight.
type UomWeightType string

const (
	UomWeightTypeKG UomWeightType = "KG" // Kilogram
	UomWeightTypeLB UomWeightType = "LB" // Pound
)

// The reference quantities used to express the value of wind speed.
type UomWindSpeedType string

const (
	UomWindSpeedTypeKM_H UomWindSpeedType = "KM_H" // Kilometres per hour.
	UomWindSpeedTypeKT UomWindSpeedType = "KT" // Knots.
	UomWindSpeedTypeM_SEC UomWindSpeedType = "M_SEC" // Metres per second
	UomWindSpeedTypeMPH UomWindSpeedType = "MPH" // Miles per hour.
)

