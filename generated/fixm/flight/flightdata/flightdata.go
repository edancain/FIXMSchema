// Code generated from FlightData.xsd; DO NOT EDIT.

package flightdata

import (
	"github.com/edancain/FIXMSchema.git/generated/fixm/base"
)

// Up to four-digit commercial flight number. [FIXM]
type FlightNumberType base.CountPositiveType

// The category of flight rules with which the pilot intends to comply. [ICAO Doc 4444, Appendix 2, Item 8]
type FlightRulesCategoryType string

const (
	FlightRulesCategoryTypeI FlightRulesCategoryType = "I" // The entire flight is intended to be operated under the IFR.
	FlightRulesCategoryTypeV FlightRulesCategoryType = "V" // The entire flight is intended to be operated under the VFR.
	FlightRulesCategoryTypeY FlightRulesCategoryType = "Y" // The flight will initially be operated under the IFR, followed by one or more subsequent changes of flight rules.
	FlightRulesCategoryTypeZ FlightRulesCategoryType = "Z" // The flight will initially be operated under the VFR, followed by one or more subsequent changes of flight rules.
)

// IATA identifier for the operator of the flight. [FIXM]
type IataOperatorCodeType base.CharacterStringType

// One character suffix used to further identify a flight. [FIXM]
type OperationalSuffixType base.CharacterStringType

// The reason for special handling of a flight by ATS. [adapted from ICAO Doc 4444, Appendix 2, ITEM 18 STS/]
type SpecialHandlingReasonCodeType string

const (
	// for a flight operated in accordance with an altitude reservation [ICAO Doc 4444, Appendix 2, ITEM 18 STS/]
	SpecialHandlingReasonCodeTypeALTRV SpecialHandlingReasonCodeType = "ALTRV"
	// for a flight approved for exemption from ATFM measures by the appropriate ATS authority; [ICAO Doc 4444, Appendix 2, ITEM 18 STS/]
	SpecialHandlingReasonCodeTypeATFMX SpecialHandlingReasonCodeType = "ATFMX"
	// fire-fighting; [ICAO Doc 4444, Appendix 2, ITEM 18 STS/]
	SpecialHandlingReasonCodeTypeFFR SpecialHandlingReasonCodeType = "FFR"
	// flight check for calibration of navaids; [ICAO Doc 4444, Appendix 2, ITEM 18 STS/]
	SpecialHandlingReasonCodeTypeFLTCK SpecialHandlingReasonCodeType = "FLTCK"
	// for a flight carrying hazardous material; [ICAO Doc 4444, Appendix 2, ITEM 18 STS/]
	SpecialHandlingReasonCodeTypeHAZMAT SpecialHandlingReasonCodeType = "HAZMAT"
	// a flight with Head of State status; [ICAO Doc 4444, Appendix 2, ITEM 18 STS/]
	SpecialHandlingReasonCodeTypeHEAD SpecialHandlingReasonCodeType = "HEAD"
	// for a medical flight declared by medical authorities; [ICAO Doc 4444, Appendix 2, ITEM 18 STS/]
	SpecialHandlingReasonCodeTypeHOSP SpecialHandlingReasonCodeType = "HOSP"
	// for a flight operating on a humanitarian mission; [ICAO Doc 4444, Appendix 2, ITEM 18 STS/]
	SpecialHandlingReasonCodeTypeHUM SpecialHandlingReasonCodeType = "HUM"
	// for a flight for which a military entity assumes responsibility for separation of military aircraft; [ICAO Doc 4444, Appendix 2, ITEM 18 STS/]
	SpecialHandlingReasonCodeTypeMARSA SpecialHandlingReasonCodeType = "MARSA"
	// for a life critical medical emergency evacuation; [ICAO Doc 4444, Appendix 2, ITEM 18 STS/]
	SpecialHandlingReasonCodeTypeMEDEVAC SpecialHandlingReasonCodeType = "MEDEVAC"
	// for a non-RVSM capable flight intending to operate in RVSM airspace; [ICAO Doc 4444, Appendix 2, ITEM 18 STS/]
	SpecialHandlingReasonCodeTypeNONRVSM SpecialHandlingReasonCodeType = "NONRVSM"
	// for a flight engaged in a search and rescue mission; [ICAO Doc 4444, Appendix 2, ITEM 18 STS/]
	SpecialHandlingReasonCodeTypeSAR SpecialHandlingReasonCodeType = "SAR"
	// for a flight engaged in military, customs or police services. [ICAO Doc 4444, Appendix 2, ITEM 18 STS/]
	SpecialHandlingReasonCodeTypeSTATE SpecialHandlingReasonCodeType = "STATE"
)

// Helper simpleType to allow for enumerated lists in FIXM.
type SpecialHandlingReasonCodeListType []flight.SpecialHandlingReasonCodeType

// Indication of the rule under which an air traffic controller provides categorical handling of a flight. [FIXM]
type TypeOfFlightType string

const (
	// General aviation [ICAO Doc 4444, Item 8]
	TypeOfFlightTypeG TypeOfFlightType = "G"
	// Military [ICAO Doc 4444, Item 8]
	TypeOfFlightTypeM TypeOfFlightType = "M"
	// Non-scheduled air transport operation [ICAO Doc 4444, Item 8]
	TypeOfFlightTypeN TypeOfFlightType = "N"
	// Scheduled air service [ICAO Doc 4444, Item 8]
	TypeOfFlightTypeS TypeOfFlightType = "S"
	// Other than any of the defined categories above. [ICAO Doc 4444, Item 8]
	TypeOfFlightTypeX TypeOfFlightType = "X"
)

// This is the central object of the FIXM Logical Model. It groups all information about the flight. [FIXM]
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
	// A general flight constraint is intended to express a constraint on the flight that cannot be associated with a specific trajectory point, either because it is not relevant to only a specific point or because it is not possible to identify the point. [FF-ICE]
	FlightConstraint []flight.FlightConstraintType `xml:"flightConstraint"`
	// A container of flight identifying data. [FIXM]
	FlightIdentification *flight.FlightIdentificationType `xml:"flightIdentification"`
	// Originator of the Flight Plan, can be FF-ICE Participant or a non-upgraded ASP. [adapted from ICAO Draft FF-ICE Implementation Guidance] Equivalent to ATS 18 ORGN/.
	FlightPlanOriginator *base.PersonOrOrganizationType `xml:"flightPlanOriginator"`
	// The FF-ICE Participant that submitted the flight plan. [ICAO Draft FF-ICE Implementation Guidance]
	FlightPlanSubmitter *base.PersonOrOrganizationType `xml:"flightPlanSubmitter"`
	// The category of flight rules with which the pilot intends to comply. [ICAO Doc 4444, Appendix 2, Item 8]
	FlightRulesCategory *flight.FlightRulesCategoryType `xml:"flightRulesCategory"`
	// Indication of the rule under which an air traffic controller provides categorical handling of a flight. [FIXM]
	FlightType *flight.TypeOfFlightType `xml:"flightType"`
	// A person, organization or enterprise engaged in or offering to engage in an aircraft operation. [ICAO Annex 9]
	Operator *base.AircraftOperatorType `xml:"operator"`
	// Contains information about radio communication failure
	RadioCommunicationFailure *flight.RadioCommunicationFailureType `xml:"radioCommunicationFailure"`
	// Any other plain-language remarks when required by the appropriate ATS authority or deemed necessary, by the pilot-in-command for the provision of air traffic services.
	Remarks *base.CharacterStringType `xml:"remarks"`
	// A logical grouping for all Route Trajectory Groups associated with the flight. [FIXM]
	RouteTrajectoryGroup *flight.RouteTrajectoryGroupContainerType `xml:"routeTrajectoryGroup"`
	// The reason for special handling of a flight by ATS. [adapted from ICAO Doc 4444, Appendix 2, ITEM 18 STS/]
	SpecialHandling *flight.SpecialHandlingReasonCodeListType `xml:"specialHandling"`
	// A container for supplementary information about the flight. This container does not capture the complete set of elements described in ICAO Doc 4444 ITEM 19 "Supplementary Information"; it only contains those elements that could not be modeled more logically in other FIXM structures. [FIXM]
	SupplementaryInformation *flight.SupplementaryInformationType `xml:"supplementaryInformation"`
}

// A general flight constraint is intended to express a constraint on the flight that cannot be associated with a specific trajectory point, either because it is not relevant to only a specific point or because it is not possible to identify the point. [FF-ICE]
type FlightConstraintType struct {
	// Expression of a general flight constraint shall allow for a description of applicability.  Description of the applicability shall accommodate a free-text description. [FF-ICE]
	Applicability *base.CharacterStringType `xml:"applicability"`
	// An extension hook for attaching extension (non-core) classes.
	Extension []base.FlightConstraintExtensionType `xml:"extension"`
	// Expression of a general flight constraint shall allow for a description of the impact of the restriction on the flight. Description of the impact of the constraint shall accommodate a free-text description.  [FF-ICE]
	Impact *base.CharacterStringType `xml:"impact"`
	// A general flight constraint should allow reference to a NOTAM; advisory; or AIXM identifier as necessary. [FF-ICE]
	RestrictionReference *base.RestrictionReferenceType `xml:"restrictionReference"`
}

// A container of flight identifying data. [FIXM]
type FlightIdentificationType struct {
	// A group of letters, figures or a combination thereof which is either identical to, or the coded equivalent of, the aircraft call sign to be used in air-ground communications, and which is used to identify the aircraft in ground-ground air traffic services communications. [ICAO Doc 4444]
	AircraftIdentification *base.AircraftIdentificationType `xml:"aircraftIdentification"`
	// Specifies the previous aircraft identification value when a change is made.
	AircraftIdentificationPrevious *base.AircraftIdentificationType `xml:"aircraftIdentificationPrevious"`
	// An extension hook for attaching extension (non-core) classes.
	Extension []base.FlightIdentificationExtensionType `xml:"extension"`
	// International Air Transport Association (IATA) flight designator. [FIXM]
	IataFlightDesignator *flight.IataFlightDesignatorType `xml:"iataFlightDesignator"`
	// A single reference for FF-ICE information pertinent to a flight that is unique globally. [ICAO Doc 9965]
	// This is a reference that uniquely identifies a specific flight and is independent of any particular system. [FIXM]
	Gufi *base.GloballyUniqueFlightIdentifierType `xml:"gufi"`
	// The legacy version of the GUFI included in the Core 4.3.0 release solely for backwards compatibility reasons.  This field should only be used when the GUFI assigned to a flight is in Core 4.2.0 format but there is a need to publish information about the flight in Core 4.3.0 format.  In all other cases, the new version of the GUFI should be used.
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

// A logical grouping for all Route Trajectory Groups associated with the flight [FIXM]
type RouteTrajectoryGroupContainerType struct {
	// The 4D Trajectory agreed to by a FF-ICE enabled Air Traffic Management Service Providers (eASP) after collaboration between the FF-ICE enabled airspace user (eAU) and the eASP.
	// The route of flight agreed to by a FF-ICE enabled Air Traffic Management Service Providers (eASP) after collaboration between the FF-ICE enabled airspace user (eAU) and the eASP. 
	Agreed *flight.RouteTrajectoryGroupType `xml:"agreed"`
	// Represents the current flight plan concept described by ICAO PANS-ATM Doc 4444.
	Current *flight.RouteTrajectoryGroupType `xml:"current"`
	// The preferred route of flight submitted by the FF-ICE enabled airspace user (eAU) to the FF-ICE enabled Air Traffic Management Service Providers (eASP) subject to required constraints.
	// This trajectory indicates the preferred 4D trajectory submitted by the FF-ICE enabled airspace user (eAU) to the FF-ICE enabled Air Traffic Management Service Providers (eASP) subject to required constraints.
	Desired *flight.RouteTrajectoryGroupType `xml:"desired"`
	// An extension hook for attaching extension (non-core) classes.
	Extension []base.RouteTrajectoryGroupContainerExtensionType `xml:"extension"`
	// The 4D Trajectory used during the collaboration between the FF-ICE enabled airspace user (eAU) and the FF-ICE enabled Air Traffic Management Service Providers (eASP) in order to agree on a 4D trajectory. This trajectory is intended to be transitory. 
	// This Route is used during collaboration between the FF-ICE enabled airspace user (eAU) and the FF-ICE enabled Air Traffic Management Service Providers (eASP) in order to agree on a route. This route field is intended to be transitory.
	Negotiating *flight.RouteTrajectoryGroupType `xml:"negotiating"`
}

// A container for supplementary information about the flight. This container does not capture the complete set of elements described in ICAO Doc 4444 ITEM 19 "Supplementary Information"; it only contains those elements that could not be modeled more logically in other FIXM structures. [FIXM]
type SupplementaryInformationType struct {
	// An extension hook for attaching extension (non-core) classes.
	Extension []base.SupplementaryInformationExtensionType `xml:"extension"`
	// The estimated maximum length of time the aircraft can spend in the cruise phase of flight, determined by the amount of fuel at takeoff. [FIXM]
	FuelEndurance *base.DurationType `xml:"fuelEndurance"`
	// The total number of persons (passengers and crew) on board the aircraft. [ICAO Doc 4444, Appendix 2, Item 19]
	PersonsOnBoard *base.CountType `xml:"personsOnBoard"`
	// The pilot designated by the operator, or in the case of general aviation, the owner, as being in command and charged with the safe conduct of a flight. [ICAO Doc 4444]
	PilotInCommand *base.PersonOrOrganizationType `xml:"pilotInCommand"`
	// Describes the source of the supplementary information. Can be either a Person/Organization or ATC Unit.
	// A Supplementary Information Source can be identified using a location identifier (from Doc. 7910 + Doc. 8585) for an ATS unit; an AFTN address; FF-ICE Participant identification per B-2.28; or appropriate contact information per B-2.12.
	SupplementaryInformationSource *flight.SupplementaryInformationSourceChoiceType `xml:"supplementaryInformationSource"`
}

// Describes the source of the supplementary information. Can be either a Person/Organization or ATC Unit.
// A Supplementary Information Source can be identified using a location identifier (from Doc. 7910 + Doc. 8585) for an ATS unit; an AFTN address; FF-ICE Participant identification per B-2.28; or appropriate contact information per B-2.12.
type SupplementaryInformationSourceChoiceType struct {
	// A Person/Organization source of the supplementary information.
	PersonOrOrganization base.PersonOrOrganizationType `xml:"personOrOrganization"`
	// An ATC unit source of the supplementary information.
	Unit base.AtcUnitReferenceType `xml:"unit"`
}

// The Flight element is an entry point to the FIXM model.
var Flight flight.FlightType

