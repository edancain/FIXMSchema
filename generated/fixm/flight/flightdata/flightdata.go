// Code generated from FlightData.xsd; DO NOT EDIT.

package flightdata

import (
	"github.com/edancain/FIXMSchema.git/generated/fixm/base"
)

// Up to four-digit commercial flight number. [FIXM]
type FlightNumberType base.CountPositiveType

// The category of flight rules with which the pilot intends to comply. [ICAO Do...
type FlightRulesCategoryType string

const (
	FlightRulesCategoryTypeI FlightRulesCategoryType = "I"
	FlightRulesCategoryTypeV FlightRulesCategoryType = "V"
	FlightRulesCategoryTypeY FlightRulesCategoryType = "Y"
	FlightRulesCategoryTypeZ FlightRulesCategoryType = "Z"
)

// IATA identifier for the operator of the flight. [FIXM]
type IataOperatorCodeType base.CharacterStringType

// One character suffix used to further identify a flight. [FIXM]
type OperationalSuffixType base.CharacterStringType

// The reason for special handling of a flight by ATS. [adapted from ICAO Doc 44...
type SpecialHandlingReasonCodeType string

const (
	SpecialHandlingReasonCodeTypeALTRV SpecialHandlingReasonCodeType = "ALTRV"
	SpecialHandlingReasonCodeTypeATFMX SpecialHandlingReasonCodeType = "ATFMX"
	SpecialHandlingReasonCodeTypeFFR SpecialHandlingReasonCodeType = "FFR"
	SpecialHandlingReasonCodeTypeFLTCK SpecialHandlingReasonCodeType = "FLTCK"
	SpecialHandlingReasonCodeTypeHAZMAT SpecialHandlingReasonCodeType = "HAZMAT"
	SpecialHandlingReasonCodeTypeHEAD SpecialHandlingReasonCodeType = "HEAD"
	SpecialHandlingReasonCodeTypeHOSP SpecialHandlingReasonCodeType = "HOSP"
	SpecialHandlingReasonCodeTypeHUM SpecialHandlingReasonCodeType = "HUM"
	SpecialHandlingReasonCodeTypeMARSA SpecialHandlingReasonCodeType = "MARSA"
	SpecialHandlingReasonCodeTypeMEDEVAC SpecialHandlingReasonCodeType = "MEDEVAC"
	SpecialHandlingReasonCodeTypeNONRVSM SpecialHandlingReasonCodeType = "NONRVSM"
	SpecialHandlingReasonCodeTypeSAR SpecialHandlingReasonCodeType = "SAR"
	SpecialHandlingReasonCodeTypeSTATE SpecialHandlingReasonCodeType = "STATE"
)

// Helper simpleType to allow for enumerated lists in FIXM.
type SpecialHandlingReasonCodeListType []flight.SpecialHandlingReasonCodeType

// Indication of the rule under which an air traffic controller provides categor...
type TypeOfFlightType string

const (
	TypeOfFlightTypeG TypeOfFlightType = "G"
	TypeOfFlightTypeM TypeOfFlightType = "M"
	TypeOfFlightTypeN TypeOfFlightType = "N"
	TypeOfFlightTypeS TypeOfFlightType = "S"
	TypeOfFlightTypeX TypeOfFlightType = "X"
)

// This is the central object of the FIXM Logical Model. It groups all informati...
type FlightType struct {
	// Aircraft enabling the flight. [FIXM]
	Aircraft *flight.AircraftType `xml:"aircraft"`
	// The actual arrival of the flight. [FIXM]
	Arrival *flight.ArrivalType `xml:"arrival"`
	// Contains information about any onboard dangerous goods.
	DangerousGoods []flight.DangerousGoodsType `xml:"dangerousGoods"`
	// Contains flight departure information
	Departure *flight.DepartureType `xml:"departure"`
	// Groups emergency information (description, phase, position, etc) for the flight.
	Emergency *flight.FlightEmergencyType `xml:"emergency"`
	// Groups the en route information about the flight.
	EnRoute *flight.EnRouteType `xml:"enRoute"`
	// An extension hook for attaching extension (non-core) classes.
	Extension []base.FlightExtensionType `xml:"extension"`
	// A general flight constraint is intended to express a constraint on the flight...
	FlightConstraint []flight.FlightConstraintType `xml:"flightConstraint"`
	// A container of flight identifying data. [FIXM]
	FlightIdentification *flight.FlightIdentificationType `xml:"flightIdentification"`
	// Originator of the Flight Plan, can be FF-ICE Participant or a non-upgraded AS...
	FlightPlanOriginator *base.PersonOrOrganizationType `xml:"flightPlanOriginator"`
	// The FF-ICE Participant that submitted the flight plan. [ICAO Draft FF-ICE Imp...
	FlightPlanSubmitter *base.PersonOrOrganizationType `xml:"flightPlanSubmitter"`
	// The category of flight rules with which the pilot intends to comply. [ICAO Do...
	FlightRulesCategory *flight.FlightRulesCategoryType `xml:"flightRulesCategory"`
	// Indication of the rule under which an air traffic controller provides categor...
	FlightType *flight.TypeOfFlightType `xml:"flightType"`
	// A person, organization or enterprise engaged in or offering to engage in an a...
	Operator *base.AircraftOperatorType `xml:"operator"`
	// Contains information about radio communication failure
	RadioCommunicationFailure *flight.RadioCommunicationFailureType `xml:"radioCommunicationFailure"`
	// Any other plain-language remarks when required by the appropriate ATS authori...
	Remarks *base.CharacterStringType `xml:"remarks"`
	// A logical grouping for all Route Trajectory Groups associated with the flight...
	RouteTrajectoryGroup *flight.RouteTrajectoryGroupContainerType `xml:"routeTrajectoryGroup"`
	// The reason for special handling of a flight by ATS. [adapted from ICAO Doc 44...
	SpecialHandling *flight.SpecialHandlingReasonCodeListType `xml:"specialHandling"`
	// A container for supplementary information about the flight. This container do...
	SupplementaryInformation *flight.SupplementaryInformationType `xml:"supplementaryInformation"`
}

// A general flight constraint is intended to express a constraint on the flight...
type FlightConstraintType struct {
	// Expression of a general flight constraint shall allow for a description of ap...
	Applicability *base.CharacterStringType `xml:"applicability"`
	// An extension hook for attaching extension (non-core) classes.
	Extension []base.FlightConstraintExtensionType `xml:"extension"`
	// Expression of a general flight constraint shall allow for a description of th...
	Impact *base.CharacterStringType `xml:"impact"`
	// A general flight constraint should allow reference to a NOTAM; advisory; or A...
	RestrictionReference *base.RestrictionReferenceType `xml:"restrictionReference"`
}

// A container of flight identifying data. [FIXM]
type FlightIdentificationType struct {
	// A group of letters, figures or a combination thereof which is either identica...
	AircraftIdentification *base.AircraftIdentificationType `xml:"aircraftIdentification"`
	// Specifies the previous aircraft identification value when a change is made.
	AircraftIdentificationPrevious *base.AircraftIdentificationType `xml:"aircraftIdentificationPrevious"`
	// An extension hook for attaching extension (non-core) classes.
	Extension []base.FlightIdentificationExtensionType `xml:"extension"`
	// International Air Transport Association (IATA) flight designator. [FIXM]
	IataFlightDesignator *flight.IataFlightDesignatorType `xml:"iataFlightDesignator"`
	// A single reference for FF-ICE information pertinent to a flight that is uniqu...
	Gufi *base.GloballyUniqueFlightIdentifierType `xml:"gufi"`
	// The legacy version of the GUFI included in the Core 4.3.0 release solely for ...
	GufiLegacy *base.UniversallyUniqueIdentifierType `xml:"gufiLegacy"`
}

// International Air Transport Association (IATA) flight designator. [FIXM]
type IataFlightDesignatorType struct {
	// An extension hook for attaching extension (non-core) classes.
	Extension []base.IataFlightDesignatorExtensionType `xml:"extension"`
	// Up to four-digit commercial flight number. [FIXM]
	FlightNumber *flight.FlightNumberType `xml:"flightNumber"`
	// IATA identifier for the operator of the flight. [FIXM]
	IataOperatorCode *flight.IataOperatorCodeType `xml:"iataOperatorCode"`
	// One character suffix used to further identify a flight. [FIXM]
	OperationalSuffix *flight.OperationalSuffixType `xml:"operationalSuffix"`
}

// A logical grouping for all Route Trajectory Groups associated with the flight...
type RouteTrajectoryGroupContainerType struct {
	// The 4D Trajectory agreed to by a FF-ICE enabled Air Traffic Management Servic...
	Agreed *flight.RouteTrajectoryGroupType `xml:"agreed"`
	// Represents the current flight plan concept described by ICAO PANS-ATM Doc 4444.
	Current *flight.RouteTrajectoryGroupType `xml:"current"`
	// The preferred route of flight submitted by the FF-ICE enabled airspace user (...
	Desired *flight.RouteTrajectoryGroupType `xml:"desired"`
	// An extension hook for attaching extension (non-core) classes.
	Extension []base.RouteTrajectoryGroupContainerExtensionType `xml:"extension"`
	// The 4D Trajectory used during the collaboration between the FF-ICE enabled ai...
	Negotiating *flight.RouteTrajectoryGroupType `xml:"negotiating"`
}

// A container for supplementary information about the flight. This container do...
type SupplementaryInformationType struct {
	// An extension hook for attaching extension (non-core) classes.
	Extension []base.SupplementaryInformationExtensionType `xml:"extension"`
	// The estimated maximum length of time the aircraft can spend in the cruise pha...
	FuelEndurance *base.DurationType `xml:"fuelEndurance"`
	// The total number of persons (passengers and crew) on board the aircraft. [ICA...
	PersonsOnBoard *base.CountType `xml:"personsOnBoard"`
	// The pilot designated by the operator, or in the case of general aviation, the...
	PilotInCommand *base.PersonOrOrganizationType `xml:"pilotInCommand"`
	// Describes the source of the supplementary information. Can be either a Person...
	SupplementaryInformationSource *flight.SupplementaryInformationSourceChoiceType `xml:"supplementaryInformationSource"`
}

// Describes the source of the supplementary information. Can be either a Person...
type SupplementaryInformationSourceChoiceType struct {
	// A Person/Organization source of the supplementary information.
	PersonOrOrganization base.PersonOrOrganizationType `xml:"personOrOrganization"`
	// An ATC unit source of the supplementary information.
	Unit base.AtcUnitReferenceType `xml:"unit"`
}

// The Flight element is an entry point to the FIXM model.
var Flight flight.FlightType

