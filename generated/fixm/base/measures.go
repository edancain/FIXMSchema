// Code generated from Measures.xsd; DO NOT EDIT.
// Modified manually to fix simple types with attributes.

package base

import (
	"encoding/xml"
	"github.com/shopspring/decimal"
)

// Code indicating the source of the altitude.
type AltitudeSourceType string

const (
	AltitudeSourceTypeBAROMETRIC AltitudeSourceType = "BAROMETRIC" // This value indicates that the source is Barometric.
	AltitudeSourceTypeGNSS AltitudeSourceType = "GNSS" // This value indicates that the source is the Global Navigation Satellite System.
)

// The result from performing the act or process of ascertaining the value of a characteristic of some entity. [ISO 19103, chapter 6.5.7.2]
type MeasureType decimal.Decimal

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
	Value decimal.Decimal `xml:",chardata"`
	UOM UomAltitudeType `xml:"uom,attr"`
}

// An altitude with a specified source. [FIXM]
type AltitudeWithSourceType struct {
	Value decimal.Decimal `xml:",chardata"`
	UOM UomAltitudeType `xml:"uom,attr"`
	Source AltitudeSourceType `xml:"source,attr"`
}

// The amount of rotation needed to bring one line or plane into coincidence with another, generally measured in radians or degrees. [ISO 19103, chapter 6.5.7.9]
type AngleType struct {
	Value decimal.Decimal `xml:",chardata"`
	UOM UomAngleType `xml:"uom,attr"`
}

// A data type used to represent direction in the coordinate reference system. [ISO 19107, chapter 6.3.12.1]
// The value of a bearing indication (at a given point) is measured as the angle between the bearing and either True North or Magnetic North. The angle is measured clockwise from 0 degrees up to and including 360 degrees. The value can also be a VOR radial.  [AIXM 5.1]
type BearingType struct {
	Value decimal.Decimal `xml:",chardata"`
	UOM UomAngleType `xml:"uom,attr"`
	ZeroBearingType ZeroBearingTypeType `xml:"zeroBearingType,attr"`
}

// A type for returning the separation between two points. [ISO 19103, chapter 6.5.7.7]
type DistanceType struct {
	Value decimal.Decimal `xml:",chardata"`
	UOM UomLengthType `xml:"uom,attr"`
}

// A surface of constant atmospheric pressure which is related to a specific pressure datum, 1 013.2 hectopascals (hPa), and is separated from other such surfaces by specific pressure intervals. [ICAO Doc 4444]
type FlightLevelType struct {
	Value decimal.Decimal `xml:",chardata"`
	UOM UomFlightLevelType `xml:"uom,attr"`
}

// The frequency value of a navigation system. [AIXM 5.1]
type FrequencyType struct {
	Value decimal.Decimal `xml:",chardata"`
	UOM UomFrequencyType `xml:"uom,attr"`
}

// The speed of an aircraft relative to the surface of the earth. [ICAO Doc 9426]
type GroundSpeedType struct {
	Value decimal.Decimal `xml:",chardata"`
	UOM UomGroundSpeedType `xml:"uom,attr"`
}

// The vertical distance of a level, a point or an object considered as a point, measured from a specified datum. [ICAO Doc 4444]
type HeightType struct {
	Value decimal.Decimal `xml:",chardata"`
	UOM UomHeightType `xml:"uom,attr"`
	Ref VerticalReferenceType `xml:"ref,attr"`
}

// The uncorrected reading on the airspeed indicator. [ICAO Doc 9426]
type IndicatedAirspeedType struct {
	Value decimal.Decimal `xml:",chardata"`
	UOM UomAirspeedType `xml:"uom,attr"`
}

// The measure of distance as an integral, for example the length of curve, the perimeter of a polygon as the length of the boundary. [ISO 19103, chapter 6.5.7.6]
type LengthType struct {
	Value decimal.Decimal `xml:",chardata"`
	UOM UomLengthType `xml:"uom,attr"`
}

// The value of a mass. [AIXM 5.1]
type MassType struct {
	Value decimal.Decimal `xml:",chardata"`
	UOM UomMassType `xml:"uom,attr"`
}

// The value of a pressure. [AIXM 5.1]
type PressureType struct {
	Value decimal.Decimal `xml:",chardata"`
	UOM UomPressureType `xml:"uom,attr"`
}

// The value of a temperature. [AIXM 5.1]
type TemperatureType struct {
	Value decimal.Decimal `xml:",chardata"`
	UOM UomTemperatureType `xml:"uom,attr"`
}

// The speed of the aeroplane relative to undisturbed air. [ICAO Doc 9713]
type TrueAirspeedType struct {
	Value decimal.Decimal `xml:",chardata"`
	UOM UomAirspeedType `xml:"uom,attr"`
}

// An expression of an aircraft's vertical rate of change. [FIXM]
type VerticalRateType struct {
	Value decimal.Decimal `xml:",chardata"`
	UOM UomVerticalRateType `xml:"uom,attr"`
}

// The measure of the physical space of any 3-D geometric object. [ISO 19103, ch...
type VolumeType struct {
	Value decimal.Decimal `xml:",chardata"`
	UOM UomVolumeType `xml:"uom,attr"`
}

// The value of a weight. [AIXM 5.1]
type WeightType struct {
	Value decimal.Decimal `xml:",chardata"`
	UOM UomWeightType `xml:"uom,attr"`
}

// Direction from which the wind blows. [International Meteorological Vocabulary, WMO]
type WindDirectionType struct {
	Value decimal.Decimal `xml:",chardata"`
	UOM UomAngleType `xml:"uom,attr"`
}

// The speed of wind. [FIXM]
type WindSpeedType struct {
	Value decimal.Decimal `xml:",chardata"`
	UOM UomWindSpeedType `xml:"uom,attr"`
}

// Custom XML marshaling/unmarshaling for types with simple content and attributes
func (a *AltitudeType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var s struct {
		Value decimal.Decimal `xml:",chardata"`
		UOM UomAltitudeType `xml:"uom,attr"`
	}
	
	if err := d.DecodeElement(&s, &start); err != nil {
		return err
	}
	
	a.Value = s.Value
	a.UOM = s.UOM
	return nil
}

func (a AltitudeType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	// Add UOM attribute
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "uom"}, Value: string(a.UOM)})
	
	// Start element
	if err := e.EncodeToken(start); err != nil {
		return err
	}
	
	// Write value
	if err := e.EncodeToken(xml.CharData(a.Value.String())); err != nil {
		return err
	}
	
	// End element
	return e.EncodeToken(start.End())
}

func (f *FlightLevelType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var s struct {
		Value decimal.Decimal `xml:",chardata"`
		UOM UomFlightLevelType `xml:"uom,attr"`
	}
	
	if err := d.DecodeElement(&s, &start); err != nil {
		return err
	}
	
	f.Value = s.Value
	f.UOM = s.UOM
	return nil
}

func (f FlightLevelType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	// Add UOM attribute
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "uom"}, Value: string(f.UOM)})
	
	// Start element
	if err := e.EncodeToken(start); err != nil {
		return err
	}
	
	// Write value
	if err := e.EncodeToken(xml.CharData(f.Value.String())); err != nil {
		return err
	}
	
	// End element
	return e.EncodeToken(start.End())
}

func (i *IndicatedAirspeedType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var s struct {
		Value decimal.Decimal `xml:",chardata"`
		UOM UomAirspeedType `xml:"uom,attr"`
	}
	
	if err := d.DecodeElement(&s, &start); err != nil {
		return err
	}
	
	i.Value = s.Value
	i.UOM = s.UOM
	return nil
}

func (i IndicatedAirspeedType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	// Add UOM attribute
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "uom"}, Value: string(i.UOM)})
	
	// Start element
	if err := e.EncodeToken(start); err != nil {
		return err
	}
	
	// Write value
	if err := e.EncodeToken(xml.CharData(i.Value.String())); err != nil {
		return err
	}
	
	// End element
	return e.EncodeToken(start.End())
}

func (t *TrueAirspeedType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var s struct {
		Value decimal.Decimal `xml:",chardata"`
		UOM UomAirspeedType `xml:"uom,attr"`
	}
	
	if err := d.DecodeElement(&s, &start); err != nil {
		return err
	}
	
	t.Value = s.Value
	t.UOM = s.UOM
	return nil
}

func (t TrueAirspeedType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	// Add UOM attribute
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "uom"}, Value: string(t.UOM)})
	
	// Start element
	if err := e.EncodeToken(start); err != nil {
		return err
	}
	
	// Write value
	if err := e.EncodeToken(xml.CharData(t.Value.String())); err != nil {
		return err
	}
	
	// End element
	return e.EncodeToken(start.End())
}

func (m *MassType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var s struct {
		Value decimal.Decimal `xml:",chardata"`
		UOM UomMassType `xml:"uom,attr"`
	}
	
	if err := d.DecodeElement(&s, &start); err != nil {
		return err
	}
	
	m.Value = s.Value
	m.UOM = s.UOM
	return nil
}

func (m MassType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	// Add UOM attribute
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "uom"}, Value: string(m.UOM)})
	
	// Start element
	if err := e.EncodeToken(start); err != nil {
		return err
	}
	
	// Write value
	if err := e.EncodeToken(xml.CharData(m.Value.String())); err != nil {
		return err
	}
	
	// End element
	return e.EncodeToken(start.End())
}