// Code generated from Aircraft.xsd; DO NOT EDIT.

package aircraft

import (
	"github.com/edancain/FIXMSchema.git/generated/fixm/base"
)

// A unique combination of twenty-four bits available for assignment to an aircr...
type AircraftAddressType base.CharacterStringType

// Classification of aircraft based on 1.3 times stall speed in landing configur...
type AircraftApproachCategoryType string

const (
	AircraftApproachCategoryTypeA AircraftApproachCategoryType = "A"
	AircraftApproachCategoryTypeB AircraftApproachCategoryType = "B"
	AircraftApproachCategoryTypeC AircraftApproachCategoryType = "C"
	AircraftApproachCategoryTypeD AircraftApproachCategoryType = "D"
	AircraftApproachCategoryTypeE AircraftApproachCategoryType = "E"
	AircraftApproachCategoryTypeH AircraftApproachCategoryType = "H"
)

// A unique, alphanumeric string that identifies a civil aircraft and consists o...
type AircraftRegistrationType base.CharacterStringType

// Helper simpleType that allows representation of aircraft registrations as a l...
type AircraftRegistrationListType []flight.AircraftRegistrationType

// ICAO classification of the aircraft wake turbulence, based on the maximum cer...
type WakeTurbulenceCategoryType string

const (
	WakeTurbulenceCategoryTypeH WakeTurbulenceCategoryType = "H"
	WakeTurbulenceCategoryTypeJ WakeTurbulenceCategoryType = "J"
	WakeTurbulenceCategoryTypeL WakeTurbulenceCategoryType = "L"
	WakeTurbulenceCategoryTypeM WakeTurbulenceCategoryType = "M"
)

// Aircraft enabling the flight. [FIXM]
type AircraftType struct {
	// A unique combination of twenty-four bits available for assignment to an aircr...
	AircraftAddress *flight.AircraftAddressType `xml:"aircraftAddress"`
	// Classification of aircraft based on 1.3 times stall speed in landing configur...
	AircraftApproachCategory *flight.AircraftApproachCategoryType `xml:"aircraftApproachCategory"`
	// The type of aircraft enabling the flight. [FIXM]
	AircraftType []flight.AircraftTypeType `xml:"aircraftType"`
	// The capabilities of the flight comprising of the: a) presence of relevant ser...
	Capabilities *flight.FlightCapabilitiesType `xml:"capabilities"`
	// The colours and markings of the aircraft. [ICAO Doc 4444, Appendix 3]
	ColoursAndMarkings *base.CharacterStringType `xml:"coloursAndMarkings"`
	// An extension hook for attaching extension (non-core) classes.
	Extension []base.AircraftExtensionType `xml:"extension"`
	// A unique, alphanumeric string that identifies a civil aircraft and consists o...
	Registration *flight.AircraftRegistrationListType `xml:"registration"`
	// ICAO classification of the aircraft wake turbulence, based on the maximum cer...
	WakeTurbulence *flight.WakeTurbulenceCategoryType `xml:"wakeTurbulence"`
}

// The type of aircraft enabling the flight. [FIXM]
type AircraftTypeType struct {
	// In the case of a formation flight, number of aircraft of the same aircraft ty...
	AircraftCount *base.CountPositiveType `xml:"aircraftCount"`
	// An extension hook for attaching extension (non-core) classes.
	Extension []base.AircraftTypeExtensionType `xml:"extension"`
	// The aircraft type designator of the aircraft as specified in ICAO Doc 8643. [...
	IcaoAircraftTypeDesignator base.AircraftTypeDesignatorType `xml:"icaoAircraftTypeDesignator"`
	// Other, non-ICAO identification of the aircraft type. [FIXM]
	OtherAircraftType base.CharacterStringType `xml:"otherAircraftType"`
}

