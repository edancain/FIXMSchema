// Code generated from Departure.xsd; DO NOT EDIT.

package departure

import (
	"github.com/edancain/FIXMSchema.git/generated/fixm/base"
)

// Identifies a flight that has filed its flight plan while in the air, beginnin...
type AirfileIndicatorType string

const (
	AirfileIndicatorTypeAIRFILE AirfileIndicatorType = "AIRFILE"
)

// The type of departure time represented. [FIXM]
type DepartureTimeTypeType flight.TrajectoryPointPropertyTypeType

const (
	DepartureTimeTypeTypeOFF_BLOCKS DepartureTimeTypeType = "OFF_BLOCKS"
	DepartureTimeTypeTypeSTART_TAKEOFF_ROLL DepartureTimeTypeType = "START_TAKEOFF_ROLL"
	DepartureTimeTypeTypeWHEELS_OFF DepartureTimeTypeType = "WHEELS_OFF"
)

// The time, type of time (wheels-off, off blocks, etc.), and geographical posit...
type ActualTimeOfDepartureType struct {
	Extension []base.ActualTimeOfDepartureExtensionType `xml:"extension"`
	// The geographical position associated with the actual time of departure. [FIXM]
	Position *base.GeographicalPositionType `xml:"position"`
	// The actual time of departure from the aerodrome. [ICAO Doc 4444, Appendix 3, ...
	Time *base.DateTimeUtcType `xml:"time"`
	// The type of departure time represented. [FIXM]
	Type *flight.DepartureTimeTypeType `xml:"type"`
}

// Groups information pertaining to the flight's departure.
type DepartureType struct {
	// The time, type of time (wheels-off, off blocks, etc.), and geographical posit...
	ActualTimeOfDeparture *flight.ActualTimeOfDepartureType `xml:"actualTimeOfDeparture"`
	// Identifies a flight that has filed its flight plan while in the air, beginnin...
	AirfileIndicator *flight.AirfileIndicatorType `xml:"airfileIndicator"`
	// The identifier of the departure airport slot allocated to the flight. [FIXM]
	AirportSlotIdentification *base.AirportSlotIdentificationType `xml:"airportSlotIdentification"`
	// An extension hook for attaching extension (non-core) classes.
	Extension []base.DepartureExtensionType `xml:"extension"`
	// The runway direction for taking off. [FIXM]
	RunwayDirection *base.RunwayDirectionDesignatorType `xml:"runwayDirection"`
	// An alternate aerodrome at which an aircraft can land should this become neces...
	TakeoffAlternateAerodrome []base.AerodromeReferenceType `xml:"takeoffAlternateAerodrome"`
	// The aerodrome from which the flight departs. [FIXM]
	DepartureAerodrome *base.AerodromeReferenceType `xml:"departureAerodrome"`
	// The first point of the route, for use when a departure aerodrome does not exi...
	DeparturePoint *flight.DeparturePointChoiceType `xml:"departurePoint"`
	// Specifies the previous departure aerodrome value when a change is made.
	DepartureAerodromePrevious *base.AerodromeReferenceType `xml:"departureAerodromePrevious"`
	// Specifies the previous departure point value when a change is made.
	DeparturePointPrevious *flight.DeparturePointChoiceType `xml:"departurePointPrevious"`
	// The estimated time at which the aircraft will commence movement associated wi...
	EstimatedOffBlockTime *base.DateTimeUtcType `xml:"estimatedOffBlockTime"`
	// The estimated time of departure from the first point on the route, for use wh...
	EstimatedRouteStartTime *base.DateTimeUtcType `xml:"estimatedRouteStartTime"`
	// Specifies the previous estimated off block time value when a change is made.
	EstimatedOffBlockTimePrevious *base.DateTimeUtcType `xml:"estimatedOffBlockTimePrevious"`
	// Specifies the previous estimated route start time value when a change is made.
	EstimatedRouteStartTimePrevious *base.DateTimeUtcType `xml:"estimatedRouteStartTimePrevious"`
}

// Used for representing the point of departure when a departure aerodrome does ...
type DeparturePointChoiceType struct {
	// A named geographical location used in defining an ATS route, the flight path ...
	DesignatedPoint base.DesignatedPointType `xml:"designatedPoint"`
	// A radio navigation aid used in defining an ATS route, the flight path of an a...
	Navaid base.NavaidType `xml:"navaid"`
	// The values of latitude and longitude that define the position of a point on t...
	Position base.GeographicalPositionType `xml:"position"`
	// Bearing and distance from a reference point. [ICAO Doc 4444, Appendix 2, ITEM...
	RelativePoint base.RelativePointType `xml:"relativePoint"`
}

