package flight/flightdata

import (
	"github.com/yourusername/fixm/generated/fixm/base"
)

// Flight is The Flight element is an entry point to the FIXM model.
type Flight	*FlightType

// FlightNumberType is Up to four-digit commercial flight number.  [FIXM]
type FlightNumberType interface{}

// FlightRulesCategoryType is The flight will initially be operated under the VFR, followed by one or more subsequent changes of flight rules.
type FlightRulesCategoryType string

// IataOperatorCodeType is IATA identifier for the operator of the flight.  [FIXM]
type IataOperatorCodeType interface{}

// OperationalSuffixType is One character suffix used to further identify a flight. [FIXM]
type OperationalSuffixType interface{}

// SpecialHandlingReasonCodeType is for a flight engaged in military, customs or police services. [ICAO Doc 4444, Appendix 2, ITEM 18 STS/]
type SpecialHandlingReasonCodeType string

// SpecialHandlingReasonCodeListType is Helper simpleType to allow for enumerated lists in FIXM.
type SpecialHandlingReasonCodeListType []string

// TypeOfFlightType is Other than any of the defined categories above. [ICAO Doc 4444, Item 8]
type TypeOfFlightType string

// FlightType is This is the central object of the FIXM Logical Model. It groups all information about the flight. [FIXM]
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
	FlightRulesCategory	string	`xml:"flightRulesCategory"`
	FlightType	string	`xml:"flightType"`
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
