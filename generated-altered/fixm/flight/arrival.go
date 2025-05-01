package flight/arrival

import (
	"github.com/yourusername/fixm/generated/fixm/base"
)

// ArrivalTimeTypeType is Indicates that the associated trajectory point corresponds to the point at which the aircraft is predicted to be wheels on the runway for arrival.
type ArrivalTimeTypeType string

// ActualTimeOfArrivalType is The type of arrival time represented. [FIXM]
type ActualTimeOfArrivalType struct {
	Extension	[]*ActualTimeOfArrivalExtensionType	`xml:"extension"`
	Position	*GeographicalPositionType	`xml:"position"`
	Time	*DateTimeUtcType	`xml:"time"`
	Type	string	`xml:"type"`
}

// ArrivalType is Information about the destination and actual arrival of the flight. [FIXM]
type ArrivalType struct {
	ActualTimeOfArrival	interface{}	`xml:"actualTimeOfArrival"`
	AirportSlotIdentification	*AirportSlotIdentificationType	`xml:"airportSlotIdentification"`
	ArrivalAerodrome	*AerodromeReferenceType	`xml:"arrivalAerodrome"`
	DestinationAerodrome	*AerodromeReferenceType	`xml:"destinationAerodrome"`
	DestinationAerodromeAlternate	[]*AerodromeReferenceType	`xml:"destinationAerodromeAlternate"`
	DestinationAerodromePrevious	*AerodromeReferenceType	`xml:"destinationAerodromePrevious"`
	Extension	[]*ArrivalExtensionType	`xml:"extension"`
	ReclearanceInFlight	interface{}	`xml:"reclearanceInFlight"`
	RunwayDirection	*RunwayDirectionDesignatorType	`xml:"runwayDirection"`
}

// ReclearanceInFlightType ...
type ReclearanceInFlightType struct {
	Extension	[]*ReclearanceInFlightExtensionType	`xml:"extension"`
	FiledRevisedDestinationAerodrome	*AerodromeReferenceType	`xml:"filedRevisedDestinationAerodrome"`
	RouteToRevisedDestination	*base.CharacterStringType	`xml:"routeToRevisedDestination"`
}
