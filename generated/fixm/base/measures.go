// Code generated from Measures.xsd; DO NOT EDIT.

package base

// Code indicating the source of the altitude.
type AltitudeSourceType string

const (
	AltitudeSourceTypeBAROMETRIC AltitudeSourceType = "BAROMETRIC"
	AltitudeSourceTypeGNSS AltitudeSourceType = "GNSS"
)

// The result from performing the act or process of ascertaining the value of a ...
type MeasureType decimal

// Helper class for restrictions on Angle
type RestrictedAngleType MeasureType

// Helper class for restrictions on measures. [FIXM]
type RestrictedMeasureType MeasureType

type RestrictedSpeedType SpeedType

// Helper class for restrictions on fractionDigits
type RestrictedVerticalDistanceType VerticalDistanceType

// The value of a speed. [AIXM 5.1]
type SpeedType MeasureType

// A vertical distance value. [AIXM 5.1]
type VerticalDistanceType MeasureType

// A code indicating the reference for a vertical distance. [AIXM 5.1]
type VerticalReferenceType string

const (
	VerticalReferenceTypeSFC VerticalReferenceType = "SFC"
	VerticalReferenceTypeW84 VerticalReferenceType = "W84"
)

// A code indicating the direction of the zero bearing.
type ZeroBearingTypeType string

const (
	ZeroBearingTypeTypeMAGNETIC_NORTH ZeroBearingTypeType = "MAGNETIC_NORTH"
	ZeroBearingTypeTypeTRUE_NORTH ZeroBearingTypeType = "TRUE_NORTH"
)

// The vertical distance of a level, a point or an object considered as a point,...
type AltitudeType struct {
}

// An altitude with a specified source. [FIXM]
type AltitudeWithSourceType struct {
}

// The amount of rotation needed to bring one line or plane into coincidence wit...
type AngleType struct {
}

// A data type used to represent direction in the coordinate reference system. [...
type BearingType struct {
}

// A type for returning the separation between two points. [ISO 19103, chapter 6...
type DistanceType struct {
}

// A surface of constant atmospheric pressure which is related to a specific pre...
type FlightLevelType struct {
}

// The frequency value of a navigation system. [AIXM 5.1]
type FrequencyType struct {
}

// The speed of an aircraft relative to the surface of the earth. [ICAO Doc 9426]
type GroundSpeedType struct {
}

// The vertical distance of a level, a point or an object considered as a point,...
type HeightType struct {
}

// The uncorrected reading on the airspeed indicator. [ICAO Doc 9426]
type IndicatedAirspeedType struct {
}

// The measure of distance as an integral, for example the length of curve, the ...
type LengthType struct {
}

// The value of a mass. [AIXM 5.1]
type MassType struct {
}

// The value of a pressure. [AIXM 5.1]
type PressureType struct {
}

// The value of a temperature. [AIXM 5.1]
type TemperatureType struct {
}

// The speed of the aeroplane relative to undisturbed air. [ICAO Doc 9713]
type TrueAirspeedType struct {
}

// An expression of an aircraft's vertical rate of change. [FIXM]
type VerticalRateType struct {
}

// The measure of the physical space of any 3-D geometric object. [ISO 19103, ch...
type VolumeType struct {
}

// The value of a weight. [AIXM 5.1]
type WeightType struct {
}

// Direction from which the wind blows. [International Meteorological Vocabulary...
type WindDirectionType struct {
}

// The speed of wind. [FIXM]
type WindSpeedType struct {
}

