// Code generated by xgen. DO NOT EDIT.

package base

// VisualFlightRulesLevelType is No specific altitude has been provided for this VFR flight.
type VisualFlightRulesLevelType string

// FlightLevelOrAltitudeChoiceType is Flight Level specification.
type FlightLevelOrAltitudeChoiceType struct {
	Altitude    *AltitudeType    `xml:"altitude"`
	FlightLevel *FlightLevelType `xml:"flightLevel"`
}

// FlightLevelOrAltitudeOrRangeChoiceType is A vertical range with a lower and/or upper bound.
type FlightLevelOrAltitudeOrRangeChoiceType struct {
	Altitude    *AltitudeType      `xml:"altitude"`
	FlightLevel *FlightLevelType   `xml:"flightLevel"`
	Range       *VerticalRangeType `xml:"range"`
}

// FlightLevelOrAltitudeOrVfrChoiceType is Visual Flight Rules specification.
type FlightLevelOrAltitudeOrVfrChoiceType struct {
	Altitude          *AltitudeType    `xml:"altitude"`
	FlightLevel       *FlightLevelType `xml:"flightLevel"`
	VisualFlightRules string           `xml:"visualFlightRules"`
}

// TimeChoiceType is A specific time value.
type TimeChoiceType struct {
	Range *TimeRangeType                `xml:"range"`
	Value *DateTimeUtcHighPrecisionType `xml:"value"`
}

// TimeRangeType is Upper bound of the time range.
type TimeRangeType struct {
	Earliest  *DateTimeUtcHighPrecisionType `xml:"earliest"`
	Extension []*TimeRangeExtensionType     `xml:"extension"`
	Latest    *DateTimeUtcHighPrecisionType `xml:"latest"`
}

// TrueAirspeedChoiceType is A specific true airspeed value.
type TrueAirspeedChoiceType struct {
	Range *TrueAirspeedRangeType `xml:"range"`
	Value *TrueAirspeedType      `xml:"value"`
}

// TrueAirspeedRangeType is Upper bound of the true airspeed range.
type TrueAirspeedRangeType struct {
	Extension  []*TrueAirspeedRangeExtensionType `xml:"extension"`
	LowerSpeed *TrueAirspeedType                 `xml:"lowerSpeed"`
	UpperSpeed *TrueAirspeedType                 `xml:"upperSpeed"`
}

// VerticalRangeType is Upper bound of the vertical range.
type VerticalRangeType struct {
	Extension  []*VerticalRangeExtensionType    `xml:"extension"`
	LowerBound *FlightLevelOrAltitudeChoiceType `xml:"lowerBound"`
	UpperBound *FlightLevelOrAltitudeChoiceType `xml:"upperBound"`
}
