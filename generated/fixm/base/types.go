// Code generated from Types.xsd; DO NOT EDIT.
// Modified manually to fix simple types with attributes.

package base

import (
	"encoding/xml"
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
type DateTimeUtcType time.Time

// Describes instances identified by the combination of a date and a time expressed in Coordinated Universal Time (UTC).  [FIXM]
type DateTimeUtcHighPrecisionType time.Time

// Describes a date represented in Coordinated Universal Time (UTC). [FIXM]
type DateUtcType time.Time

// Generic decimal fraction expressed to tenths, used as scaling or comparison factor. [FIXM]
type DecimalIndexType decimal.Decimal

// Length or distance in the temporal dimension. [ISO 19108, chapter 5.2.3.7]
// Duration has no reference to start or stop times.
type DurationType time.Duration

// This class provides an optional mechanism enabling FIXM aeronautical fields to be supplemented with references to AIXM features.  Uses of this field should be considered functionally equivalent to the xlink:href field used in AIXM to reference features.  [FIXM]
type HypertextReferenceType string

// Mode A SSR code. [FIXM]
type ModeACodeType CharacterStringType

// Identifies the particular type of namespace used by the originator of a GUFI.
type NamespaceDomainType string

const (
	NamespaceDomainTypeFULLY_QUALIFIED_DOMAIN_NAME NamespaceDomainType = "FULLY_QUALIFIED_DOMAIN_NAME" // An organisation can be identified by its registered Fully Qualified Domain Name (FQDN).
	NamespaceDomainTypeLOCATION_INDICATOR NamespaceDomainType = "LOCATION_INDICATOR" // An ATM unit can be identified by its corresponding four-letter Location Indicator (LOCID as determined in ICAO Doc. 7910).
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
	Value string `xml:",chardata"`
	CreationTime DateTimeUtcType `xml:"creationTime,attr"` // The time at which the GUFI was created.
	NamespaceDomain NamespaceDomainType `xml:"namespaceDomain,attr"` // Identifies the particular type of namespace used by the originator of a GUFI.
	NamespaceIdentifier NamespaceIdentifierType `xml:"namespaceIdentifier,attr"` // The namespace chosen by the GUFI originator.
	CodeSpace string `xml:"codeSpace,attr"` // Code Space of the UUID
}

// Reference to a published restriction should allow reference to a NOTAM; advis...
type RestrictionReferenceType struct {
	// An extension hook for attaching extension (non-core) classes.
	Extension []RestrictionReferenceExtensionType `xml:"extension,omitempty"`
	// Identifier associated with the published restriction being referenced. [FIXM]
	RestrictionIdentifier *CharacterStringType `xml:"restrictionIdentifier,omitempty"`
	// Used to identify the type of published restriction being referenced. [FIXM]
	RestrictionType *CharacterStringType `xml:"restrictionType,omitempty"`
	// Hypertextreference
	Href HypertextReferenceType `xml:"href,attr,omitempty"`
}

// A 128-bit number used, with a high probability, to uniquely identify information in computer systems. [FIXM]
type UniversallyUniqueIdentifierType struct {
	Value string `xml:",chardata"`
	CodeSpace string `xml:"codeSpace,attr"` // Code Space of the UUID
}

// Helper methods for DateTimeUtcType
func (dt *DateTimeUtcType) IsZero() bool {
	if dt == nil {
		return true
	}
	
	// Since DateTimeUtcType is essentially a time.Time
	t := time.Time(*dt)
	return t.IsZero()
}

func (dt *DateTimeUtcType) Format() string {
	if dt == nil || dt.IsZero() {
		return "undefined"
	}
	
	// Convert to time.Time and format
	t := time.Time(*dt)
	return t.Format("2006-01-02T15:04:05Z")
}

// Helper methods for DateTimeUtcHighPrecisionType
func (dt *DateTimeUtcHighPrecisionType) IsZero() bool {
	if dt == nil {
		return true
	}
	
	// Since DateTimeUtcHighPrecisionType is essentially a time.Time
	t := time.Time(*dt)
	return t.IsZero()
}

func (dt *DateTimeUtcHighPrecisionType) Format() string {
	if dt == nil || dt.IsZero() {
		return "undefined"
	}
	
	// Convert to time.Time and format with higher precision
	t := time.Time(*dt)
	return t.Format("2006-01-02T15:04:05.000Z")
}

// Helper method for DateUtcType
func (dt *DateUtcType) IsZero() bool {
	if dt == nil {
		return true
	}
	
	// Since DateUtcType is essentially a time.Time
	t := time.Time(*dt)
	return t.IsZero()
}

func (dt *DateUtcType) Format() string {
	if dt == nil || dt.IsZero() {
		return "undefined"
	}
	
	// Convert to time.Time and format as date only
	t := time.Time(*dt)
	return t.Format("2006-01-02")
}

// Custom XML marshaling/unmarshaling for GloballyUniqueFlightIdentifierType
func (g *GloballyUniqueFlightIdentifierType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var s struct {
		Value string `xml:",chardata"`
		CreationTime DateTimeUtcType `xml:"creationTime,attr"`
		NamespaceDomain NamespaceDomainType `xml:"namespaceDomain,attr"`
		NamespaceIdentifier NamespaceIdentifierType `xml:"namespaceIdentifier,attr"`
		CodeSpace string `xml:"codeSpace,attr"`
	}
	
	if err := d.DecodeElement(&s, &start); err != nil {
		return err
	}
	
	g.Value = s.Value
	g.CreationTime = s.CreationTime
	g.NamespaceDomain = s.NamespaceDomain
	g.NamespaceIdentifier = s.NamespaceIdentifier
	g.CodeSpace = s.CodeSpace
	return nil
}

func (g GloballyUniqueFlightIdentifierType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	// Add attributes
	start.Attr = append(start.Attr, 
		xml.Attr{Name: xml.Name{Local: "creationTime"}, Value: g.CreationTime.Format()},
		xml.Attr{Name: xml.Name{Local: "namespaceDomain"}, Value: string(g.NamespaceDomain)},
		xml.Attr{Name: xml.Name{Local: "namespaceIdentifier"}, Value: string(g.NamespaceIdentifier)},
		xml.Attr{Name: xml.Name{Local: "codeSpace"}, Value: g.CodeSpace},
	)
	
	// Start element
	if err := e.EncodeToken(start); err != nil {
		return err
	}
	
	// Write value
	if err := e.EncodeToken(xml.CharData(g.Value)); err != nil {
		return err
	}
	
	// End element
	return e.EncodeToken(start.End())
}

// Custom XML marshaling/unmarshaling for UniversallyUniqueIdentifierType
func (u *UniversallyUniqueIdentifierType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var s struct {
		Value string `xml:",chardata"`
		CodeSpace string `xml:"codeSpace,attr"`
	}
	
	if err := d.DecodeElement(&s, &start); err != nil {
		return err
	}
	
	u.Value = s.Value
	u.CodeSpace = s.CodeSpace
	return nil
}

func (u UniversallyUniqueIdentifierType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	// Add codeSpace attribute
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "codeSpace"}, Value: u.CodeSpace})
	
	// Start element
	if err := e.EncodeToken(start); err != nil {
		return err
	}
	
	// Write value
	if err := e.EncodeToken(xml.CharData(u.Value)); err != nil {
		return err
	}
	
	// End element
	return e.EncodeToken(start.End())
}