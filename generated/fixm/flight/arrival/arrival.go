// Code generated from Arrival.xsd; DO NOT EDIT.

package arrival

import (
	"github.com/edancain/FIXMSchema.git/generated/fixm/base"
)

// The type of arrival time represented. [FIXM]
type ArrivalTimeTypeType flight.TrajectoryPointPropertyTypeType

const (
	ArrivalTimeTypeTypeEND_LANDING_ROLL ArrivalTimeTypeType = "END_LANDING_ROLL" // Indicates that the associated trajectory point corresponds to the point at which the aircraft is predicted to come to a full stop on the arrival runway. (A prediction only, the flight will likely exit the runway without coming to a full stop).
	ArrivalTimeTypeTypeIN_BLOCKS ArrivalTimeTypeType = "IN_BLOCKS" // Indicates the point and time at which an arriving aircraft is/was in blocks.
	ArrivalTimeTypeTypeWHEELS_ON ArrivalTimeTypeType = "WHEELS_ON" // Indicates that the associated trajectory point corresponds to the point at which the aircraft is predicted to be wheels on the runway for arrival.
)

// The time, type of time (wheels-on, in blocks, etc.), and geographical position of the arrival.  [FIXM]
type ActualTimeOfArrivalType struct {
	// An extension hook for attaching extension (non-core) classes.
	Extension []base.ActualTimeOfArrivalExtensionType `xml:"extension"`
	// The geographical position associated with the actual time of arrival. [FIXM]
	Position *base.GeographicalPositionType `xml:"position"`
	// The actual time of arrival. [ICAO Doc 4444, Appendix 3, Field Type 17]
	Time *base.DateTimeUtcType `xml:"time"`
	// The type of arrival time represented. [FIXM]
	Type *flight.ArrivalTimeTypeType `xml:"type"`
}

// Information about the destination and actual arrival of the flight. [FIXM]
type ArrivalType struct {
	// The time, type of time (wheels-on, in blocks, etc.), and geographical position of the arrival.  [FIXM]
	ActualTimeOfArrival *flight.ActualTimeOfArrivalType `xml:"actualTimeOfArrival"`
	// The identifier of the arrival airport slot allocated to the flight. [FIXM]
	AirportSlotIdentification *base.AirportSlotIdentificationType `xml:"airportSlotIdentification"`
	// The aerodrome at which the flight has actually arrived. [FIXM]
	ArrivalAerodrome *base.AerodromeReferenceType `xml:"arrivalAerodrome"`
	// The aerodrome of intended landing. [ICAO Doc 4444 - extracted from the definition of Destination Alternate]
	// This is the aerodrome at which the flight is scheduled to arrive. [FIXM]
	DestinationAerodrome *base.AerodromeReferenceType `xml:"destinationAerodrome"`
	// An alternate aerodrome to which an aircraft may proceed should it become either impossible or inadvisable to land at the aerodrome of intended landing. [ICAO Doc 4444]
	DestinationAerodromeAlternate []base.AerodromeReferenceType `xml:"destinationAerodromeAlternate"`
	// Specifies the previous destination aerodrome value when a change is made.
	DestinationAerodromePrevious *base.AerodromeReferenceType `xml:"destinationAerodromePrevious"`
	// An extension hook for attaching extension (non-core) classes.
	Extension []base.ArrivalExtensionType `xml:"extension"`
	// The route details to the revised destination aerodrome, followed by the ICAO four-letter location indicator of
	// the aerodrome. The revised route is subject to reclearance in flight. [ICAO Doc 4444]
	ReclearanceInFlight *flight.ReclearanceInFlightType `xml:"reclearanceInFlight"`
	// The runway direction for landing. [FIXM]
	RunwayDirection *base.RunwayDirectionDesignatorType `xml:"runwayDirection"`
}

// The route details to the revised destination aerodrome, followed by the ICAO four-letter location indicator of
// the aerodrome. The revised route is subject to reclearance in flight. [ICAO Doc 4444]
type ReclearanceInFlightType struct {
	// An extension hook for attaching extension (non-core) classes.
	Extension []base.ReclearanceInFlightExtensionType `xml:"extension"`
	// Aerodrome where the flight will divert to based on the revised route. For example, an aerodrome where the flight will divert to if more fuel is used than expected.[FIXM]
	FiledRevisedDestinationAerodrome *base.AerodromeReferenceType `xml:"filedRevisedDestinationAerodrome"`
	// Route to the revised destination aerodrome. [FIXM]
	RouteToRevisedDestination *base.CharacterStringType `xml:"routeToRevisedDestination"`
}

