// Code generated from RangesAndChoices.xsd; DO NOT EDIT.

package base

// Indicates that no specific altitude has been provided for a flight operated u...
type VisualFlightRulesLevelType string

const (
	VisualFlightRulesLevelTypeVFR VisualFlightRulesLevelType = "VFR"
)

// The Choice between flight level or altitude specification.
type FlightLevelOrAltitudeChoiceType struct {
	// Altitude specification.
	Altitude AltitudeType `xml:"altitude"`
	// Flight Level specification.
	FlightLevel FlightLevelType `xml:"flightLevel"`
}

// Represent either a specific level or a level range with a lower and/or upper ...
type FlightLevelOrAltitudeOrRangeChoiceType struct {
	// Altitude specification.
	Altitude AltitudeType `xml:"altitude"`
	// Flight Level specification.
	FlightLevel FlightLevelType `xml:"flightLevel"`
	// A vertical range with a lower and/or upper bound.
	Range VerticalRangeType `xml:"range"`
}

// A choice between flight level or altitude or VFR specification.
type FlightLevelOrAltitudeOrVfrChoiceType struct {
	// Altitude specification.
	Altitude AltitudeType `xml:"altitude"`
	// Flight Level specification.
	FlightLevel FlightLevelType `xml:"flightLevel"`
	// Visual Flight Rules specification.
	VisualFlightRules VisualFlightRulesLevelType `xml:"visualFlightRules"`
}

// Represent either a specific time or a time range with a lower and/or upper bo...
type TimeChoiceType struct {
	// Represents a time range with a lower and/or upper bound.
	Range TimeRangeType `xml:"range"`
	// A specific time value.
	Value DateTimeUtcType `xml:"value"`
}

// Represents a time range with a lower and/or upper bound.
type TimeRangeType struct {
	// Lower bound of the time range.
	Earliest *DateTimeUtcType `xml:"earliest"`
	// An extension hook for attaching extension (non-core) classes.
	Extension []TimeRangeExtensionType `xml:"extension"`
	// Upper bound of the time range.
	Latest *DateTimeUtcType `xml:"latest"`
}

// Represent either a specific true airspeed or a true airspeed range with a low...
type TrueAirspeedChoiceType struct {
	// Represents a true airspeed range with a lower and/or upper bound.
	Range TrueAirspeedRangeType `xml:"range"`
	// A specific true airspeed value.
	Value TrueAirspeedType `xml:"value"`
}

// Represents a true airspeed range with a lower and/or upper bound.
type TrueAirspeedRangeType struct {
	// An extension hook for attaching extension (non-core) classes.
	Extension []TrueAirspeedRangeExtensionType `xml:"extension"`
	// Lower bound of the true airspeed range.
	LowerSpeed *TrueAirspeedType `xml:"lowerSpeed"`
	// Upper bound of the true airspeed range.
	UpperSpeed *TrueAirspeedType `xml:"upperSpeed"`
}

// Represents a vertical range with a lower and/or upper bound.
type VerticalRangeType struct {
	// An extension hook for attaching extension (non-core) classes.
	Extension []VerticalRangeExtensionType `xml:"extension"`
	// Lower bound of the vertical range.
	LowerBound *FlightLevelOrAltitudeChoiceType `xml:"lowerBound"`
	// Upper bound of the vertical range.
	UpperBound *FlightLevelOrAltitudeChoiceType `xml:"upperBound"`
}

