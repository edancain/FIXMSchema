// Code generated from Organization.xsd; DO NOT EDIT.

package base

// The identifier of the Aircraft Operator as assigned by ICAO. [FIXM] This is t...
type AircraftOperatorDesignatorType CharacterStringType

// A person, organization or enterprise engaged in or offering to engage in an a...
type AircraftOperatorType struct {
	// The identifier of the Aircraft Operator as assigned by ICAO. [FIXM] This is t...
	DesignatorIcao *AircraftOperatorDesignatorType `xml:"designatorIcao"`
	// An extension hook for attaching extension (non-core) classes.
	Extension []AircraftOperatorExtensionType `xml:"extension"`
	// Aircraft Operator Identity: Identity of a person, organization or enterprise ...
	OperatingOrganization *PersonOrOrganizationType `xml:"operatingOrganization"`
}

// An identifiable, responsible entity that can be either a natural person or an...
type PersonOrOrganizationType struct {
	// Optional ContactInformation reference.[FIXM]
	Contact *ContactInformationType `xml:"contact"`
	// An extension hook for attaching extension (non-core) classes.
	Extension []PersonOrOrganizationExtensionType `xml:"extension"`
	// A designator used for identifying the Person, State, Organization, Authority,...
	Identifier *CharacterStringType `xml:"identifier"`
	// The relevant domain in which the person or organization's identifier is defin...
	IdentifierDomain *CharacterStringType `xml:"identifierDomain"`
	// The full official name of the Person, State, Organisation, Authority, aircraf...
	Name *TextNameType `xml:"name"`
}

