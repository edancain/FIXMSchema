// Code generated from EnRoute.xsd; DO NOT EDIT.

package enroute

import (
	"github.com/edancain/FIXMSchema.git/generated/fixm/base"
)

// Specifies whether the boundary crossing occurs at or above or at or below the specified level.
type BoundaryCrossingConditionType string

const (
	// Indicates that the transition altitude is at or above the specified.
	BoundaryCrossingConditionTypeAT_OR_ABOVE BoundaryCrossingConditionType = "AT_OR_ABOVE"
	// Indicates that the transition altitude is at or below the specified.
	BoundaryCrossingConditionTypeAT_OR_BELOW BoundaryCrossingConditionType = "AT_OR_BELOW"
)

// An altitude (flight level) at or above/below which (specified in Boundary Crossing Condition) an aircraft will cross the associated boundary point.
type AltitudeInTransitionType struct {
	// Specifies whether the boundary crossing occurs at or above or at or below the specified level.
	CrossingCondition *BoundaryCrossingConditionType `xml:"crossingCondition"`
	// An extension hook for attaching extension (non-core) classes.
	Extension []base.AltitudeInTransitionExtensionType `xml:"extension"`
	// An altitude (flight level) at or above/below which (specified in Boundary Crossing Condition) an aircraft will cross the associated boundary point.
	Level *base.FlightLevelOrAltitudeChoiceType `xml:"level"`
}

// Boundary Crossing contains estimate data conveyed between ATS Units during the process of Controller Coordination. [FIXM]
type BoundaryCrossingType struct {
	// Negotiated boundary crossing transition altitude.
	AltitudeInTransition *AltitudeInTransitionType `xml:"altitudeInTransition"`
	// The cleared altitude (flight level) at which the aircraft will cross the boundary crossing point if in level cruising flight or, if the aircraft is climbing or descending at the boundary crossing point, the cleared flight level to which it is proceeding.
	ClearedLevel *base.FlightLevelOrAltitudeChoiceType `xml:"clearedLevel"`
	// The point where the flight will cross an ATS facility boundary.
	CrossingPoint *base.SignificantPointChoiceType `xml:"crossingPoint"`
	// The estimated time at which a flight will cross the associated boundary crossing point.
	CrossingTime *base.DateTimeUtcType `xml:"crossingTime"`
	// An extension hook for attaching extension (non-core) classes.
	Extension []base.BoundaryCrossingExtensionType `xml:"extension"`
}

// Groups the en route information about the flight.
type EnRouteType struct {
	// An ICAO designator of the aerodrome to which a flight could be diverted while en route, if needed.
	AlternateAerodrome []base.AerodromeReferenceType `xml:"alternateAerodrome"`
	// A proposed boundary crossing information with associated details including altitude, crossing point and crossing time.
	BoundaryCrossingCoordination *BoundaryCrossingType `xml:"boundaryCrossingCoordination"`
	// Current assigned Mode A code. [FIXM]
	CurrentModeACode *base.ModeACodeType `xml:"currentModeACode"`
	// An extension hook for attaching extension (non-core) classes.
	Extension []base.EnRouteExtensionType `xml:"extension"`
}

