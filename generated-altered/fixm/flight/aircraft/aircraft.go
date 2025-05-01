package aircraft

import (
	"github.com/edancain/FIXMSchema/generated/fixm/base/"
)

// AircraftAddressType is A unique combination of twenty-four bits available for assignment to an aircraft for the purpose of air-ground communications, navigation and surveillance. [ICAO Doc 4444]
type AircraftAddressType interface{}

// AircraftApproachCategoryType is Helicopters [ICAO Doc 8168, Vol. I, chapter 1.3.5]
type AircraftApproachCategoryType string

// AircraftRegistrationType is A unique, alphanumeric string that identifies a civil aircraft and consists of the Aircraft Nationality or Common Mark and an additional alphanumeric string, the registration mark, assigned by the state of registry or common mark registering authority. [FIXM]
type AircraftRegistrationType interface{}

// AircraftRegistrationListType is Helper simpleType that allows representation of aircraft registrations as a list.
type AircraftRegistrationListType []interface{}

// WakeTurbulenceCategoryType is An aircraft type with a maximum certified take-off mass of less than 136000 kg. but more than 7000 kg. [ICAO Doc 4444, Appendix 2, ITEM M]
type WakeTurbulenceCategoryType string

// AircraftType is Aircraft enabling the flight. [FIXM]
type AircraftType struct {
	AircraftAddress          interface{}                   `xml:"aircraftAddress"`
	AircraftApproachCategory string                        `xml:"aircraftApproachCategory"`
	AircraftType             []interface{}                 `xml:"aircraftType"`
	Capabilities             interface{}                   `xml:"capabilities"`
	ColoursAndMarkings       *base.CharacterStringType     `xml:"coloursAndMarkings"`
	Extension                []*base.AircraftExtensionType `xml:"extension"`
	Registration             interface{}                   `xml:"registration"`
	WakeTurbulence           string                        `xml:"wakeTurbulence"`
}

// AircraftTypeType ...
type AircraftTypeType struct {
	AircraftCount              *base.CountPositiveType           `xml:"aircraftCount"`
	Extension                  []*base.AircraftTypeExtensionType `xml:"extension"`
	IcaoAircraftTypeDesignator *base.AircraftTypeDesignatorType  `xml:"icaoAircraftTypeDesignator"`
	OtherAircraftType          *base.CharacterStringType         `xml:"otherAircraftType"`
}
