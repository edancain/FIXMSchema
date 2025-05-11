// Code generated from Organization.xsd; DO NOT EDIT.

package base

// The identifier of the Aircraft Operator as assigned by ICAO. [FIXM] 
// This is the ICAO Operator Code as standardised by ICAO Doc 8585 Manual on Designators for Aircraft Operating Agencies, Aeronautical Authorities and Services.
type AircraftOperatorDesignatorType CharacterStringType

// A person, organization or enterprise engaged in or offering to engage in an aircraft operation. [ICAO Annex 9]
type AircraftOperatorType struct {
	// The identifier of the Aircraft Operator as assigned by ICAO. [FIXM]
	// This is the ICAO Operator Code as standardised by ICAO Doc 8585 Manual on Designators for Aircraft Operating Agencies, Aeronautical Authorities and Services.
	DesignatorIcao *AircraftOperatorDesignatorType `xml:"designatorIcao"`
	// An extension hook for attaching extension (non-core) classes.
	Extension []AircraftOperatorExtensionType `xml:"extension"`
	// Aircraft Operator Identity: Identity of a person, organization or enterprise engaged in or offering to engage in aircraft operation.
	OperatingOrganization *PersonOrOrganizationType `xml:"operatingOrganization"`
}

// An identifiable, responsible entity that can be either a natural person or an organization.
type PersonOrOrganizationType struct {
	// Optional ContactInformation reference.[FIXM]
	Contact *ContactInformationType `xml:"contact"`
	// An extension hook for attaching extension (non-core) classes.
	Extension []PersonOrOrganizationExtensionType `xml:"extension"`
	// A designator used for identifying the Person, State, Organization, Authority, aircraft operating agency, handling agency etc.  [FIXM]
	Identifier *CharacterStringType `xml:"identifier"`
	// The relevant domain in which the person or organization's identifier is defined or used.  [FIXM]
	IdentifierDomain *CharacterStringType `xml:"identifierDomain"`
	// The full official name of the Person, State, Organisation, Authority, aircraft operating agency, handling agency etc. [FIXM]
	Name *TextNameType `xml:"name"`
}

