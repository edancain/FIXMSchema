package flight/flightdata

import (
	"github.com/yourusername/fixm/generated/fixm/base"
)

// Flight is This is the central object of the FIXM Logical Model. It groups all information about the flight. [FIXM]
type Flight	string

// SpecialHandlingReasonCodeListType ...
type SpecialHandlingReasonCodeListType []interface{}

// FlightType ...
type FlightType struct {
	Aircraft	interface{}	`xml:"aircraft"`
	Arrival	interface{}	`xml:"arrival"`
	DangerousGoods	[]interface{}	`xml:"dangerousGoods"`
	Departure	interface{}	`xml:"departure"`
	Emergency	interface{}	`xml:"emergency"`
	EnRoute	interface{}	`xml:"enRoute"`
	Extension	[]*FlightExtensionType	`xml:"extension"`
	FlightConstraint	[]interface{}	`xml:"flightConstraint"`
	FlightIdentification	interface{}	`xml:"flightIdentification"`
	FlightPlanOriginator	*PersonOrOrganizationType	`xml:"flightPlanOriginator"`
	FlightPlanSubmitter	*PersonOrOrganizationType	`xml:"flightPlanSubmitter"`
	FlightRulesCategory	interface{}	`xml:"flightRulesCategory"`
	FlightType	interface{}	`xml:"flightType"`
	Operator	*AircraftOperatorType	`xml:"operator"`
	RadioCommunicationFailure	interface{}	`xml:"radioCommunicationFailure"`
	Remarks	*base.CharacterStringType	`xml:"remarks"`
	RouteTrajectoryGroup	interface{}	`xml:"routeTrajectoryGroup"`
	SpecialHandling	interface{}	`xml:"specialHandling"`
	SupplementaryInformation	interface{}	`xml:"supplementaryInformation"`
}

// FlightConstraintType ...
type FlightConstraintType struct {
	Applicability	*base.CharacterStringType	`xml:"applicability"`
	Extension	[]*FlightConstraintExtensionType	`xml:"extension"`
	Impact	*base.CharacterStringType	`xml:"impact"`
	RestrictionReference	*RestrictionReferenceType	`xml:"restrictionReference"`
}

// FlightIdentificationType ...
type FlightIdentificationType struct {
	AircraftIdentification	*AircraftIdentificationType	`xml:"aircraftIdentification"`
	AircraftIdentificationPrevious	*AircraftIdentificationType	`xml:"aircraftIdentificationPrevious"`
	Extension	[]*FlightIdentificationExtensionType	`xml:"extension"`
	Gufi	*GloballyUniqueFlightIdentifierType	`xml:"gufi"`
	GufiLegacy	*UniversallyUniqueIdentifierType	`xml:"gufiLegacy"`
	IataFlightDesignator	interface{}	`xml:"iataFlightDesignator"`
}

// IataFlightDesignatorType ...
type IataFlightDesignatorType struct {
	Extension	[]*IataFlightDesignatorExtensionType	`xml:"extension"`
	FlightNumber	interface{}	`xml:"flightNumber"`
	IataOperatorCode	interface{}	`xml:"iataOperatorCode"`
	OperationalSuffix	interface{}	`xml:"operationalSuffix"`
}

// RouteTrajectoryGroupContainerType ...
type RouteTrajectoryGroupContainerType struct {
	Agreed	interface{}	`xml:"agreed"`
	Current	interface{}	`xml:"current"`
	Desired	interface{}	`xml:"desired"`
	Extension	[]*RouteTrajectoryGroupContainerExtensionType	`xml:"extension"`
	Negotiating	interface{}	`xml:"negotiating"`
}

// SupplementaryInformationType ...
type SupplementaryInformationType struct {
	Extension	[]*SupplementaryInformationExtensionType	`xml:"extension"`
	FuelEndurance	*DurationType	`xml:"fuelEndurance"`
	PersonsOnBoard	*CountType	`xml:"personsOnBoard"`
	PilotInCommand	*PersonOrOrganizationType	`xml:"pilotInCommand"`
	SupplementaryInformationSource	interface{}	`xml:"supplementaryInformationSource"`
}

// SupplementaryInformationSourceChoiceType ...
type SupplementaryInformationSourceChoiceType struct {
	PersonOrOrganization	*PersonOrOrganizationType	`xml:"personOrOrganization"`
	Unit	*AtcUnitReferenceType	`xml:"unit"`
}
