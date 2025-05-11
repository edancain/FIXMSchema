// Code generated from Departure.xsd; DO NOT EDIT.

package departure

import (
	"github.com/edancain/FIXMSchema.git/generated/fixm/base"
	"github.com/edancain/FIXMSchema.git/generated/fixm/flight/flightroutetrajectory"
)

// Identifies a flight that has filed its flight plan while in the air, beginning its route description from a specified point en-route, and therefore may not have provided a departure aerodrome.
type AirfileIndicatorType string

const (
	// No attribution
	AirfileIndicatorTypeAIRFILE AirfileIndicatorType = "AIRFILE"
)

// The type of departure time represented. [FIXM]
type DepartureTimeTypeType flightroutetrajectory.TrajectoryPointPropertyTypeType

const (
	DepartureTimeTypeTypeOFF_BLOCKS DepartureTimeTypeType = "OFF_BLOCKS" // This is the point at which the aircraft pushes back and begins to taxi for departure.
	DepartureTimeTypeTypeSTART_TAKEOFF_ROLL DepartureTimeTypeType = "START_TAKEOFF_ROLL" // Indicates that the associated trajectory point corresponds to the point at the start of take-off roll (used for departures only).
	DepartureTimeTypeTypeWHEELS_OFF DepartureTimeTypeType = "WHEELS_OFF" // Indicates that the associated trajectory point corresponds to the point at which the aircraft is predicted to be wheels off the runway on departure.
)

// The time, type of time (wheels-off, off blocks, etc.), and geographical position of the departure.  [FIXM]
type ActualTimeOfDepartureType struct {
	Extension []base.ActualTimeOfDepartureExtensionType `xml:"extension"`
	// The geographical position associated with the actual time of departure. [FIXM]
	Position *base.GeographicalPositionType `xml:"position"`
	// The actual time of departure from the aerodrome. [ICAO Doc 4444, Appendix 3, Field Type 13]
	Time *base.DateTimeUtcType `xml:"time"`
	// The type of departure time represented. [FIXM]
	Type *DepartureTimeTypeType `xml:"type"`
}

// Groups information pertaining to the flight's departure.
type DepartureType struct {
	// The time, type of time (wheels-off, off blocks, etc.), and geographical position of the departure.  [FIXM]
	ActualTimeOfDeparture *ActualTimeOfDepartureType `xml:"actualTimeOfDeparture"`
	// Identifies a flight that has filed its flight plan while in the air, beginning its route description from a specified point en-route, and therefore may not have provided a departure aerodrome.
	AirfileIndicator *AirfileIndicatorType `xml:"airfileIndicator"`
	// The identifier of the departure airport slot allocated to the flight. [FIXM]
	AirportSlotIdentification *base.AirportSlotIdentificationType `xml:"airportSlotIdentification"`
	// An extension hook for attaching extension (non-core) classes.
	Extension []base.DepartureExtensionType `xml:"extension"`
	// The runway direction for taking off. [FIXM]
	RunwayDirection *base.RunwayDirectionDesignatorType `xml:"runwayDirection"`
	// An alternate aerodrome at which an aircraft can land should this become necessary shortly after take-off and it is not possible to use the aerodrome of departure. [ICAO Doc 4444]
	TakeoffAlternateAerodrome []base.AerodromeReferenceType `xml:"takeoffAlternateAerodrome"`
	// The aerodrome from which the flight departs. [FIXM]
	DepartureAerodrome *base.AerodromeReferenceType `xml:"departureAerodrome"`
	// The first point of the route, for use when a departure aerodrome does not exist (such as when the flight plan is filed from the air) or is otherwise unknown.
	DeparturePoint *DeparturePointChoiceType `xml:"departurePoint"`
	// Specifies the previous departure aerodrome value when a change is made.
	DepartureAerodromePrevious *base.AerodromeReferenceType `xml:"departureAerodromePrevious"`
	// Specifies the previous departure point value when a change is made.
	DeparturePointPrevious *DeparturePointChoiceType `xml:"departurePointPrevious"`
	// The estimated time at which the aircraft will commence movement associated with departure.
	EstimatedOffBlockTime *base.DateTimeUtcType `xml:"estimatedOffBlockTime"`
	// The estimated time of departure from the first point on the route, for use when a departure aerodrome does not exist (such as when the flight plan is filed from the air) or is otherwise unknown.
	EstimatedRouteStartTime *base.DateTimeUtcType `xml:"estimatedRouteStartTime"`
	// Specifies the previous estimated off block time value when a change is made.
	EstimatedOffBlockTimePrevious *base.DateTimeUtcType `xml:"estimatedOffBlockTimePrevious"`
	// Specifies the previous estimated route start time value when a change is made.
	EstimatedRouteStartTimePrevious *base.DateTimeUtcType `xml:"estimatedRouteStartTimePrevious"`
}

// Used for representing the point of departure when a departure aerodrome does not exist or is otherwise unknown.
type DeparturePointChoiceType struct {
	// A named geographical location used in defining an ATS route, the flight path of an aircraft or for other navigation or ATS purposes. [FIXM]
	DesignatedPoint base.DesignatedPointType `xml:"designatedPoint"`
	// A radio navigation aid used in defining an ATS route, the flight path of an aircraft or for other navigation or ATS purposes. [FIXM]
	Navaid base.NavaidType `xml:"navaid"`
	// The values of latitude and longitude that define the position of a point on the surface of the Earth with respect to a reference datum. [specialised from ICAO Doc 9881]
	Position base.GeographicalPositionType `xml:"position"`
	// Bearing and distance from a reference point. [ICAO Doc 4444, Appendix 2, ITEM 15]
	RelativePoint base.RelativePointType `xml:"relativePoint"`
}

