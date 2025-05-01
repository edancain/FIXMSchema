package flightdata

import (
	"github.com/edancain/FIXMSchema.git/generated/fixm/base"
)

// Flight is This is the central object of the FIXM Logical Model. It groups all information about the flight. [FIXM]
type Flight string

// SpecialHandlingReasonCodeListType ...
type SpecialHandlingReasonCodeListType []*interface{}

// FlightType ...
type FlightType struct {
	Aircraft                  interface{}                    `xml:"aircraft"`
	Arrival                   interface{}                    `xml:"arrival"`
	DangerousGoods            []interface{}                  `xml:"dangerousGoods"`
	Departure                 interface{}                    `xml:"departure"`
	Emergency                 interface{}                    `xml:"emergency"`
	EnRoute                   interface{}                    `xml:"enRoute"`
	Extension                 []*base.FlightExtensionType    `xml:"extension"`
	FlightConstraint          []interface{}                  `xml:"flightConstraint"`
	FlightIdentification      interface{}                    `xml:"flightIdentification"`
	FlightPlanOriginator      *base.PersonOrOrganizationType `xml:"flightPlanOriginator"`
	FlightPlanSubmitter       *base.PersonOrOrganizationType `xml:"flightPlanSubmitter"`
	FlightRulesCategory       interface{}                    `xml:"flightRulesCategory"`
	FlightType                interface{}                    `xml:"flightType"`
	Operator                  *base.AircraftOperatorType     `xml:"operator"`
	RadioCommunicationFailure interface{}                    `xml:"radioCommunicationFailure"`
	Remarks                   *base.CharacterStringType      `xml:"remarks"`
	RouteTrajectoryGroup      interface{}                    `xml:"routeTrajectoryGroup"`
	SpecialHandling           interface{}                    `xml:"specialHandling"`
	SupplementaryInformation  interface{}                    `xml:"supplementaryInformation"`
}

// FlightConstraintType ...
type FlightConstraintType struct {
	Applicability        *base.CharacterStringType             `xml:"applicability"`
	Extension            []*base.FlightConstraintExtensionType `xml:"extension"`
	Impact               *base.CharacterStringType             `xml:"impact"`
	RestrictionReference *base.RestrictionReferenceType        `xml:"restrictionReference"`
}

// FlightIdentificationType ...
type FlightIdentificationType struct {
	AircraftIdentification         *base.AircraftIdentificationType          `xml:"aircraftIdentification"`
	AircraftIdentificationPrevious *base.AircraftIdentificationType          `xml:"aircraftIdentificationPrevious"`
	Extension                      []*base.FlightIdentificationExtensionType `xml:"extension"`
	Gufi                           *base.GloballyUniqueFlightIdentifierType  `xml:"gufi"`
	GufiLegacy                     *base.UniversallyUniqueIdentifierType     `xml:"gufiLegacy"`
	IataFlightDesignator           interface{}                               `xml:"iataFlightDesignator"`
}

// IataFlightDesignatorType ...
type IataFlightDesignatorType struct {
	Extension         []*base.IataFlightDesignatorExtensionType `xml:"extension"`
	FlightNumber      interface{}                               `xml:"flightNumber"`
	IataOperatorCode  interface{}                               `xml:"iataOperatorCode"`
	OperationalSuffix interface{}                               `xml:"operationalSuffix"`
}

// RouteTrajectoryGroupContainerType ...
type RouteTrajectoryGroupContainerType struct {
	Agreed      interface{}                                        `xml:"agreed"`
	Current     interface{}                                        `xml:"current"`
	Desired     interface{}                                        `xml:"desired"`
	Extension   []*base.RouteTrajectoryGroupContainerExtensionType `xml:"extension"`
	Negotiating interface{}                                        `xml:"negotiating"`
}

// SupplementaryInformationType ...
type SupplementaryInformationType struct {
	Extension                      []*base.SupplementaryInformationExtensionType `xml:"extension"`
	FuelEndurance                  *base.DurationType                            `xml:"fuelEndurance"`
	PersonsOnBoard                 *base.CountType                               `xml:"personsOnBoard"`
	PilotInCommand                 *base.PersonOrOrganizationType                `xml:"pilotInCommand"`
	SupplementaryInformationSource interface{}                                   `xml:"supplementaryInformationSource"`
}

// SupplementaryInformationSourceChoiceType ...
type SupplementaryInformationSourceChoiceType struct {
	PersonOrOrganization *base.PersonOrOrganizationType `xml:"personOrOrganization"`
	Unit                 *base.AtcUnitReferenceType     `xml:"unit"`
}
