// Code generated from Aircraft.xsd; DO NOT EDIT.

package aircraft

import (
	"github.com/edancain/FIXMSchema/generated/fixm/base"
	"github.com/edancain/FIXMSchema/generated/fixm/flight/capability"
)

// A unique combination of twenty-four bits available for assignment to an aircraft for the purpose of air-ground communications, navigation and surveillance. [ICAO Doc 4444]
type AircraftAddressType base.CharacterStringType

// Classification of aircraft based on 1.3 times stall speed in landing configuration at maximum certified landing mass. [AIXM 5.1]
type AircraftApproachCategoryType string

const (
	AircraftApproachCategoryTypeA AircraftApproachCategoryType = "A" // less than 169 km/h (91 kt) indicated airspeed (IAS) [ICAO Doc 8168, Vol. I, chapter 1.3.5]
	AircraftApproachCategoryTypeB AircraftApproachCategoryType = "B" // 169 km/h (91 kt) or more but less than 224 km/h (121 kt) IAS [ICAO Doc 8168, Vol. I, chapter 1.3.5]
	AircraftApproachCategoryTypeC AircraftApproachCategoryType = "C" // 224 km/h (121 kt) or more but less than 261 km/h (141 kt) IAS [ICAO Doc 8168, Vol. I, chapter 1.3.5]
	AircraftApproachCategoryTypeD AircraftApproachCategoryType = "D" // 261 km/h (141 kt) or more but less than 307 km/h (166 kt) IAS [ICAO Doc 8168, Vol. I, chapter 1.3.5]
	AircraftApproachCategoryTypeE AircraftApproachCategoryType = "E" // 307 km/h (166 kt) or more but less than 391 km/h (211 kt) IAS [ICAO Doc 8168, Vol. I, chapter 1.3.5]
	AircraftApproachCategoryTypeH AircraftApproachCategoryType = "H" // Helicopters [ICAO Doc 8168, Vol. I, chapter 1.3.5]
)

// A unique, alphanumeric string that identifies a civil aircraft and consists of the Aircraft Nationality or Common Mark and an additional alphanumeric string, the registration mark, assigned by the state of registry or common mark registering authority. [FIXM]
type AircraftRegistrationType base.CharacterStringType

// Helper simpleType that allows representation of aircraft registrations as a list.
type AircraftRegistrationListType []AircraftRegistrationType

// ICAO classification of the aircraft wake turbulence, based on the maximum certified take off mass. [FIXM]
type WakeTurbulenceCategoryType string

const (
	WakeTurbulenceCategoryTypeH WakeTurbulenceCategoryType = "H" // An aircraft type with a maximum certified take-off mass of 136000 kg or more. [ICAO Doc 4444, Appendix 2, ITEM M]
	WakeTurbulenceCategoryTypeJ WakeTurbulenceCategoryType = "J" // A super heavy aircraft type e.g. the Airbus A380-800. (In the order of 560,000kg). [FIXM]
	WakeTurbulenceCategoryTypeL WakeTurbulenceCategoryType = "L" // An aircraft type with a maximum certified take-off mass of 7000 kg or less. [ICAO Doc 4444, Appendix 2, ITEM M]
	WakeTurbulenceCategoryTypeM WakeTurbulenceCategoryType = "M" // An aircraft type with a maximum certified take-off mass of less than 136000 kg. but more than 7000 kg. [ICAO Doc 4444, Appendix 2, ITEM M]
)

// Aircraft enabling the flight. [FIXM]
type AircraftType struct {
	// A unique combination of twenty-four bits available for assignment to an aircraft for the purpose of air-ground communications, navigation and surveillance. [ICAO Doc 4444]
	AircraftAddress *AircraftAddressType `xml:"aircraftAddress"`
	// Classification of aircraft based on 1.3 times stall speed in landing configuration at maximum certified landing mass. [AIXM 5.1]
	AircraftApproachCategory *AircraftApproachCategoryType `xml:"aircraftApproachCategory"`
	// The type of aircraft enabling the flight. [FIXM]
	AircraftType []AircraftTypeType `xml:"aircraftType"`
	// The capabilities of the flight comprising of the:
			// a) presence of relevant serviceable equipment on board the aircraft;
			// b) equipment and capabilities commensurate with flight crew qualifications; and
			// c) where applicable, authorization from the appropriate authority.
	Capabilities *capability.FlightCapabilitiesType `xml:"capabilities"`
	// The colours and markings of the aircraft. [ICAO Doc 4444, Appendix 3]
	ColoursAndMarkings *base.CharacterStringType `xml:"coloursAndMarkings"`
	// An extension hook for attaching extension (non-core) classes.
	Extension []base.AircraftExtensionType `xml:"extension"`
	// A unique, alphanumeric string that identifies a civil aircraft and consists of the Aircraft Nationality or Common Mark and an additional alphanumeric string, the registration mark, assigned by the state of registry or common mark registering authority. [FIXM]
	Registration *AircraftRegistrationListType `xml:"registration"`
	// ICAO classification of the aircraft wake turbulence, based on the maximum certified take off mass. [FIXM]
	WakeTurbulence *WakeTurbulenceCategoryType `xml:"wakeTurbulence"`
}

// The type of aircraft enabling the flight. [FIXM]
type AircraftTypeType struct {
	// In the case of a formation flight, number of aircraft of the same aircraft type within the formation. [FIXM]
	AircraftCount *base.CountPositiveType `xml:"aircraftCount"`
	// An extension hook for attaching extension (non-core) classes.
	Extension []base.AircraftTypeExtensionType `xml:"extension"`
	// The aircraft type designator of the aircraft as specified in ICAO Doc 8643. [FIXM]
	IcaoAircraftTypeDesignator base.AircraftTypeDesignatorType `xml:"icaoAircraftTypeDesignator"`
	// Other, non-ICAO identification of the aircraft type. [FIXM]
	OtherAircraftType base.CharacterStringType `xml:"otherAircraftType"`
}

