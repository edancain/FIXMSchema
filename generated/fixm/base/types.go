// Code generated from Types.xsd; DO NOT EDIT.

package base

// A group of letters, figures or a combination thereof which is either identica...
type AircraftIdentificationType CharacterStringType

// A group of alphanumeric characters used to identify, in an abbreviated form, ...
type AircraftTypeDesignatorType CharacterStringType

// The identifier of the scheduled time of arrival or departure available for al...
type AirportSlotIdentificationType CharacterStringType

// A sequence of characters [specialised from ISO19103, chapter 6.5.2.7]. In FIX...
type CharacterStringType string

// Represents non-negative integer counts of objects.
type CountType int

// Represents positive integer counts of objects.
type CountPositiveType CountType

// Describes instances identified by the combination of a date and a time expres...
type DateTimeUtcType DateTimeUtcHighPrecisionType

// Describes instances identified by the combination of a date and a time expres...
type DateTimeUtcHighPrecisionType dateTime

// Describes a date represented in Coordinated Universal Time (UTC). [FIXM]
type DateUtcType date

// Generic decimal fraction expressed to tenths, used as scaling or comparison f...
type DecimalIndexType decimal

// Length or distance in the temporal dimension. [ISO 19108, chapter 5.2.3.7] Du...
type DurationType duration

// This class provides an optional mechanism enabling FIXM aeronautical fields t...
type HypertextReferenceType anyURI

// Mode A SSR code. [FIXM]
type ModeACodeType CharacterStringType

// Identifies the particular type of namespace used by the originator of a GUFI.
type NamespaceDomainType string

const (
	NamespaceDomainTypeFULLY_QUALIFIED_DOMAIN_NAME NamespaceDomainType = "FULLY_QUALIFIED_DOMAIN_NAME"
	NamespaceDomainTypeLOCATION_INDICATOR NamespaceDomainType = "LOCATION_INDICATOR"
	NamespaceDomainTypeOPERATING_AGENCY_DESIGNATOR NamespaceDomainType = "OPERATING_AGENCY_DESIGNATOR"
)

// The namespace chosen by the GUFI originator. Used to reduce the likelihood of...
type NamespaceIdentifierType CharacterStringType

// Helper type for restrictions on UniversallyUniqueIdentifier
type RestrictedUniversallyUniqueIdentifierType CharacterStringType

// The official name of a State, an aerodrome, a unit, etc... [AIXM 5.1]
type TextNameType CharacterStringType

// An immutable identifier associated with a flight that allows all eligible mem...
type GloballyUniqueFlightIdentifierType struct {
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

// A 128-bit number used, with a high probability, to uniquely identify informat...
type UniversallyUniqueIdentifierType struct {
}

