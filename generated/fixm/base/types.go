// Code generated from Types.xsd; DO NOT EDIT.

package base

import (
	"github.com/shopspring/decimal"
	"time"
)

// A group of letters, figures or a combination thereof which is either identical to, or the coded equivalent of, the aircraft call sign to be used in air-ground communications, and which is used to identify the aircraft in ground-ground air traffic services communications. [ICAO Doc 4444]
type AircraftIdentificationType CharacterStringType

// A group of alphanumeric characters used to identify, in an abbreviated form, a type of aircraft. [ICAO Doc 9426]
// Approved aircraft type designators are defined in ICAO Doc. 8643 - Aircraft Type Designators. [FIXM]
type AircraftTypeDesignatorType CharacterStringType

// The identifier of the scheduled time of arrival or departure available for allocation by, or as allocated by, a coordinator for an aircraft movement on a specific date at a coordinated airport. [adapted from IATA Worldwide Scheduling Guidelines, 21st Edition]
type AirportSlotIdentificationType CharacterStringType

// A sequence of characters [specialised from ISO19103, chapter 6.5.2.7]. 
// In FIXM, the maximum length of a CharacterString is set to 4096 characters. [FIXM]
type CharacterStringType string

// Represents non-negative integer counts of objects.
type CountType int

// Represents positive integer counts of objects.
type CountPositiveType CountType

// Describes instances identified by the combination of a date and a time expressed in Coordinated Universal Time (UTC) and restricted to whole second precision.  [FIXM]
type DateTimeUtcType DateTimeUtcHighPrecisionType // TODO value range

// Describes instances identified by the combination of a date and a time expressed in Coordinated Universal Time (UTC).  [FIXM]
type DateTimeUtcHighPrecisionType time.Time //dateTime // TODO value range

// Describes a date represented in Coordinated Universal Time (UTC). [FIXM]
type DateUtcType time.Time // date // TODO value range

// Generic decimal fraction expressed to tenths, used as scaling or comparison factor. [FIXM]
type DecimalIndexType decimal.Decimal

// Length or distance in the temporal dimension. [ISO 19108, chapter 5.2.3.7]
// Duration has no reference to start or stop times.
type DurationType time.Duration // TODO value range

// This class provides an optional mechanism enabling FIXM aeronautical fields to be supplemented with references to AIXM features.  Uses of this field should be considered functionally equivalent to the xlink:href field used in AIXM to reference features.  [FIXM]
type HypertextReferenceType string

// Mode A SSR code. [FIXM]
type ModeACodeType CharacterStringType

// Identifies the particular type of namespace used by the originator of a GUFI.
type NamespaceDomainType string

const (
	NamespaceDomainTypeFULLY_QUALIFIED_DOMAIN_NAME NamespaceDomainType = "FULLY_QUALIFIED_DO" // An organisation can be identified by its registered Fully Qualified Domain Name (FQDN). The domain name used should uniquely identify the organisation, and it could be from either the organisation's email or from the organisation's website. For example, for the website https://www.example.com/ and email person@example.com, the domain name used for the namespace would be "example.com". The namespace can also include subdomains as needed to ensure proper uniqueness and allow for the best implementation. This could be done by adding a numeric or regional subdomain. For example, "east.example.com" or "region1.example.com".
	NamespaceDomainTypeLOCATION_INDICATOR NamespaceDomainType = "LOCATION_INDICATOR" // An ATM unit can be identified by its corresponding four-letter Location Indicator (LOCID as determined in ICAO Doc. 7910). A LOCID is available for each FIR or ACC, as well as for airport locations that operators are tied into. For example, the ATM unit "Washington ATC Center", LOCID KZDC, might use the namespace "KZDC".
	NamespaceDomainTypeOPERATING_AGENCY_DESIGNATOR NamespaceDomainType = "OPERATING_AGENCY_DESIGNATOR" // An operating agency can be identified by its three-letter Operating Agency Designator (ICAO Doc. 8585).
)

// The namespace chosen by the GUFI originator.  Used to reduce the likelihood of GUFI collisions and provide traceability as to which entity generated a particular GUFI.
type NamespaceIdentifierType CharacterStringType

// Helper type for restrictions on UniversallyUniqueIdentifier
type RestrictedUniversallyUniqueIdentifierType CharacterStringType

// The official name of a State, an aerodrome, a unit, etc... [AIXM 5.1]
type TextNameType CharacterStringType

// An immutable identifier associated with a flight that allows all eligible members of the ATM community to unambiguously refer to information pertaining to the flight.
type GloballyUniqueFlightIdentifierType struct {
	CreationTime DateTimeUtcType `xml:"creationTime"` // The time at which the GUFI was created.  Used to reduce the likelihood of GUFI collisions and ensure GUFIs remain perpetually unique.
	NamespaceDomain NamespaceDomainType `xml:"namespaceDomain"` // Identifies the particular type of namespace used by the originator of a GUFI.
	NamespaceIdentifier NamespaceIdentifierType `xml:"namespaceIdentifier"` // The namespace chosen by the GUFI originator.  Used to reduce the likelihood of GUFI collisions and provide traceability as to which entity generated a particular GUFI.
}

// Reference to a published restriction should allow reference to a NOTAM; advis...
type RestrictionReferenceType struct {
	// An extension hook for attaching extension (non-core) classes.
	Extension []RestrictionReferenceExtensionType `xml:"extension"`
	// Identifier associated with the published restriction being referenced. [FIXM]
	RestrictionIdentifier *CharacterStringType `xml:"restrictionIdentifier"`
	// Used to identify the type of published restriction being referenced. [FIXM]
	RestrictionType *CharacterStringType `xml:"restrictionType"`
}

// A 128-bit number used, with a high probability, to uniquely identify information in computer systems. [FIXM]
type UniversallyUniqueIdentifierType struct {
	CodeSpace string `xml:"codeSpace"` // Code Space of the UUID
}

