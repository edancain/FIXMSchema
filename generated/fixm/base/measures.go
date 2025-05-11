// Code generated from Measures.xsd; DO NOT EDIT.

package base

// Code indicating the source of the altitude.
type AltitudeSourceType string

const (
	AltitudeSourceTypeBAROMETRIC AltitudeSourceType = "BAROMETRIC" // This value indicates that the source is Barometric.
	AltitudeSourceTypeGNSS AltitudeSourceType = "GNSS" // This value indicates that the source is the Global Navigation Satellite System.
)

// The result from performing the act or process of ascertaining the value of a characteristic of some entity. [ISO 19103, chapter 6.5.7.2]
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
	VerticalReferenceTypeSFC VerticalReferenceType = "SFC" // The distance measured from the surface of the Earth (equivalent to AGL - Above Ground Level). [AIXM 5.1]
	VerticalReferenceTypeW84 VerticalReferenceType = "W84" // The distance measured from the WGS84 ellipsoid. [AIXM 5.1]
)

// A code indicating the direction of the zero bearing.
type ZeroBearingTypeType string

const (
	ZeroBearingTypeTypeMAGNETIC_NORTH ZeroBearingTypeType = "MAGNETIC_NORTH" // This value indicates that the direction is relative to Magnetic North.
	ZeroBearingTypeTypeTRUE_NORTH ZeroBearingTypeType = "TRUE_NORTH" // This value indicates that the direction is relative to True North.
)

// The vertical distance of a level, a point or an object considered as a point, measured from mean sea level (MSL). [ICAO Doc 4444]
type AltitudeType struct {
	//TODO check this as I add it manually
	uom *UomAltitudeType `xml:"uom"`// Unit of measure
}

// An altitude with a specified source. [FIXM]
type AltitudeWithSourceType struct {
	//TODO check this as I add it manually
	source AltitudeSourceType `xml:"source"`// The source of the altitude. [FIXM]
}

// The amount of rotation needed to bring one line or plane into coincidence with another, generally measured in radians or degrees. [ISO 19103, chapter 6.5.7.9]
type AngleType struct {
	//TODO check this as I add it manually
	uom UomAngleType `xml:"uom"`// Unit of measure
}

// A data type used to represent direction in the coordinate reference system. [ISO 19107, chapter 6.3.12.1]
// The value of a bearing indication (at a given point) is measured as the angle between the bearing and either True North or Magnetic North. The angle is measured clockwise from 0 degrees up to and including 360 degrees. The value can also be a VOR radial.  [AIXM 5.1]
type BearingType struct {
	//TODO check this as I add it manually
	zeroBearingType ZeroBearingTypeType `xml:"zeroBearingType"` // A code indicating the direction of the zero bearing.
}

// A type for returning the separation between two points. [ISO 19103, chapter 6.5.7.7]
type DistanceType struct {
}

// A surface of constant atmospheric pressure which is related to a specific pressure datum, 1 013.2 hectopascals (hPa), and is separated from other such surfaces by specific pressure intervals. [ICAO Doc 4444]
type FlightLevelType struct {
	//TODO check this as I add it manually
	uom UomFlightLevelType `xml:"uom"`// Unit of measure
}

// The frequency value of a navigation system. [AIXM 5.1]
type FrequencyType struct {
	//TODO check this as I add it manually
	uom UomFrequencyType `xml:"uom"`// Unit of measure
}

// The speed of an aircraft relative to the surface of the earth. [ICAO Doc 9426]
type GroundSpeedType struct {
	//TODO check this as I add it manually
	uom UomGroundSpeedType `xml:"uom"`// Unit of measure
}

// The vertical distance of a level, a point or an object considered as a point, measured from a specified datum. [ICAO Doc 4444]
type HeightType struct {
	//TODO check this as I add it manually
	ref VerticalReferenceType `xml:"ref"` // Reference for the vertical measure
	uom UomHeightType `xml:"uom"`// Unit of measure
}

// The uncorrected reading on the airspeed indicator. [ICAO Doc 9426]
type IndicatedAirspeedType struct {
	//TODO check this as I add it manually
	uom UomAirspeedType `xml:"uom"`// Unit of measure
}

// The measure of distance as an integral, for example the length of curve, the perimeter of a polygon as the length of the boundary. [ISO 19103, chapter 6.5.7.6]
type LengthType struct {
	//TODO check this as I add it manually
	uom UomLengthType `xml:"uom"`// Unit of measure
}

// The value of a mass. [AIXM 5.1]
type MassType struct {
	//TODO check this as I add it manually
	uom UomMassType `xml:"uom"`// Unit of measure
}

// The value of a pressure. [AIXM 5.1]
type PressureType struct {
	//TODO check this as I add it manually
	uom UomPressureType `xml:"uom"`// Unit of measure
}

// The value of a temperature. [AIXM 5.1]
type TemperatureType struct {
	//TODO check this as I add it manually
	uom UomTemperatureType `xml:"uom"`// Unit of measure
}

// The speed of the aeroplane relative to undisturbed air. [ICAO Doc 9713]
type TrueAirspeedType struct {
	//TODO check this as I add it manually
	uom UomAirspeedType `xml:"uom"`// Unit of measure
}

// An expression of an aircraft's vertical rate of change. [FIXM]
type VerticalRateType struct {
	//TODO check this as I add it manually
	uom UomVerticalRateType `xml:"uom"`// Unit of measure
}

// The measure of the physical space of any 3-D geometric object. [ISO 19103, ch...
type VolumeType struct {
	//TODO check this as I add it manually
	uom UomVolumeType `xml:"uom"`// Unit of measure
}

// The value of a weight. [AIXM 5.1]
type WeightType struct {
	//TODO check this as I add it manually
	uom UomWeightType `xml:"uom"`// Unit of measure
}

// Direction from which the wind blows. [International Meteorological Vocabulary, WMO]
type WindDirectionType struct {
}

// The speed of wind. [FIXM]
type WindSpeedType struct {
	//TODO check this as I add it manually
	uom UomWindSpeedType `xml:"uom"`// Unit of measure
}

