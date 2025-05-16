// Code generated from RangesAndChoices.xsd; DO NOT EDIT.
// Modified manually to fix choice types issues.

package base

import (
	"encoding/xml"
)

// Indicates that no specific altitude has been provided for a flight operated under visual flight rules.
type VisualFlightRulesLevelType string

const (
	VisualFlightRulesLevelTypeVFR VisualFlightRulesLevelType = "VFR"
)

// The Choice between flight level or altitude specification.
type FlightLevelOrAltitudeChoiceType struct {
	// Altitude specification.
	Altitude *AltitudeType `xml:"altitude,omitempty"`
	// Flight Level specification.
	FlightLevel *FlightLevelType `xml:"flightLevel,omitempty"`
}

// Represent either a specific level or a level range with a lower and/or upper bound.
type FlightLevelOrAltitudeOrRangeChoiceType struct {
	// Altitude specification.
	Altitude *AltitudeType `xml:"altitude,omitempty"`
	// Flight Level specification.
	FlightLevel *FlightLevelType `xml:"flightLevel,omitempty"`
	// A vertical range with a lower and/or upper bound.
	Range *VerticalRangeType `xml:"range,omitempty"`
}

// A choice between flight level or altitude or VFR specification.
type FlightLevelOrAltitudeOrVfrChoiceType struct {
	// Altitude specification.
	Altitude *AltitudeType `xml:"altitude,omitempty"`
	// Flight Level specification.
	FlightLevel *FlightLevelType `xml:"flightLevel,omitempty"`
	// Visual Flight Rules specification.
	VisualFlightRules VisualFlightRulesLevelType `xml:"visualFlightRules,omitempty"`
}

// Represent either a specific time or a time range with a lower and/or upper bound.
type TimeChoiceType struct {
	// Represents a time range with a lower and/or upper bound.
	Range *TimeRangeType `xml:"range,omitempty"`
	// A specific time value.
	Value *DateTimeUtcType `xml:"value,omitempty"`
}

// Represents a time range with a lower and/or upper bound.
type TimeRangeType struct {
	// Lower bound of the time range.
	Earliest *DateTimeUtcType `xml:"earliest,omitempty"`
	// An extension hook for attaching extension (non-core) classes.
	Extension []TimeRangeExtensionType `xml:"extension,omitempty"`
	// Upper bound of the time range.
	Latest *DateTimeUtcType `xml:"latest,omitempty"`
}

// Represent either a specific true airspeed or a true airspeed range with a lower and/or upper bound.
type TrueAirspeedChoiceType struct {
	// Represents a true airspeed range with a lower and/or upper bound.
	Range *TrueAirspeedRangeType `xml:"range,omitempty"`
	// A specific true airspeed value.
	Value *TrueAirspeedType `xml:"value,omitempty"`
}

// Represents a true airspeed range with a lower and/or upper bound.
type TrueAirspeedRangeType struct {
	// An extension hook for attaching extension (non-core) classes.
	Extension []TrueAirspeedRangeExtensionType `xml:"extension,omitempty"`
	// Lower bound of the true airspeed range.
	LowerSpeed *TrueAirspeedType `xml:"lowerSpeed,omitempty"`
	// Upper bound of the true airspeed range.
	UpperSpeed *TrueAirspeedType `xml:"upperSpeed,omitempty"`
}

// Represents a vertical range with a lower and/or upper bound.
type VerticalRangeType struct {
	// An extension hook for attaching extension (non-core) classes.
	Extension []VerticalRangeExtensionType `xml:"extension,omitempty"`
	// Lower bound of the vertical range.
	LowerBound *FlightLevelOrAltitudeChoiceType `xml:"lowerBound,omitempty"`
	// Upper bound of the vertical range.
	UpperBound *FlightLevelOrAltitudeChoiceType `xml:"upperBound,omitempty"`
}

// Helper methods to check which choice is set in FlightLevelOrAltitudeChoiceType
func (c *FlightLevelOrAltitudeChoiceType) IsAltitudeSet() bool {
	return c != nil && c.Altitude != nil
}

func (c *FlightLevelOrAltitudeChoiceType) IsFlightLevelSet() bool {
	return c != nil && c.FlightLevel != nil
}

// Helper methods for FlightLevelOrAltitudeOrVfrChoiceType
func (c *FlightLevelOrAltitudeOrVfrChoiceType) IsAltitudeSet() bool {
	return c != nil && c.Altitude != nil
}

func (c *FlightLevelOrAltitudeOrVfrChoiceType) IsFlightLevelSet() bool {
	return c != nil && c.FlightLevel != nil
}

func (c *FlightLevelOrAltitudeOrVfrChoiceType) IsVFRSet() bool {
	return c != nil && string(c.VisualFlightRules) != ""
}

// Helper methods for FlightLevelOrAltitudeOrRangeChoiceType
func (c *FlightLevelOrAltitudeOrRangeChoiceType) IsAltitudeSet() bool {
	return c != nil && c.Altitude != nil
}

func (c *FlightLevelOrAltitudeOrRangeChoiceType) IsFlightLevelSet() bool {
	return c != nil && c.FlightLevel != nil
}

func (c *FlightLevelOrAltitudeOrRangeChoiceType) IsRangeSet() bool {
	return c != nil && c.Range != nil
}

// Helper methods for TimeChoiceType
func (t *TimeChoiceType) IsRangeSet() bool {
	return t != nil && t.Range != nil
}

func (t *TimeChoiceType) IsValueSet() bool {
	return t != nil && t.Value != nil
}

// Additional helper method for TimeChoiceType
func (t *TimeChoiceType) IsZero() bool {
	if t == nil {
		return true
	}
	
	if t.IsValueSet() {
		return t.Value.IsZero()
	}
	
	if t.IsRangeSet() {
		// If both earliest and latest are zero or nil
		if (t.Range.Earliest == nil || t.Range.Earliest.IsZero()) &&
		   (t.Range.Latest == nil || t.Range.Latest.IsZero()) {
			return true
		}
		return false
	}
	
	return true
}

// Format method for TimeChoiceType
func (t *TimeChoiceType) Format() string {
	if t == nil || t.IsZero() {
		return "undefined"
	}
	
	if t.IsValueSet() {
		return t.Value.Format()
	}
	
	if t.IsRangeSet() {
		result := "Range: "
		if t.Range.Earliest != nil && !t.Range.Earliest.IsZero() {
			result += "From " + t.Range.Earliest.Format()
		} else {
			result += "From undefined"
		}
		
		if t.Range.Latest != nil && !t.Range.Latest.IsZero() {
			result += " To " + t.Range.Latest.Format()
		} else {
			result += " To undefined"
		}
		
		return result
	}
	
	return "undefined"
}

// Helper methods for TrueAirspeedChoiceType
func (t *TrueAirspeedChoiceType) IsRangeSet() bool {
	return t != nil && t.Range != nil
}

func (t *TrueAirspeedChoiceType) IsValueSet() bool {
	return t != nil && t.Value != nil
}

// Custom XML marshaling for choice types
func (c *FlightLevelOrAltitudeChoiceType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// Handle choice between altitude and flightLevel
	for {
		token, err := d.Token()
		if err != nil {
			return err
		}

		switch se := token.(type) {
		case xml.StartElement:
			switch se.Name.Local {
			case "altitude":
				var alt AltitudeType
				if err := d.DecodeElement(&alt, &se); err != nil {
					return err
				}
				c.Altitude = &alt
				c.FlightLevel = nil
			case "flightLevel":
				var fl FlightLevelType
				if err := d.DecodeElement(&fl, &se); err != nil {
					return err
				}
				c.FlightLevel = &fl
				c.Altitude = nil
			default:
				// Skip other elements
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			if se.Name == start.Name {
				return nil
			}
		}
	}
}

func (t *TimeChoiceType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// Handle choice between range and value
	for {
		token, err := d.Token()
		if err != nil {
			return err
		}

		switch se := token.(type) {
		case xml.StartElement:
			switch se.Name.Local {
			case "range":
				var rng TimeRangeType
				if err := d.DecodeElement(&rng, &se); err != nil {
					return err
				}
				t.Range = &rng
				t.Value = nil
			case "value":
				var val DateTimeUtcType
				if err := d.DecodeElement(&val, &se); err != nil {
					return err
				}
				t.Value = &val
				t.Range = nil
			default:
				// Skip other elements
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			if se.Name == start.Name {
				return nil
			}
		}
	}
}