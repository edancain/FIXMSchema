package capability

import (
	"github.com/edancain/FIXMSchema/generated/fixm/base/"
)

// CommunicationCapabilityCodeType is VHF with 8.33 kHz channel spacing capability
type CommunicationCapabilityCodeType string

// CommunicationCapabilityCodeListType is Helper simpleType to allow for enumerated lists in FIXM.
type CommunicationCapabilityCodeListType []string

// DatalinkCommunicationCapabilityCodeType is CPDLC FANS 1/A SATCOM (Iridium)
type DatalinkCommunicationCapabilityCodeType string

// DatalinkCommunicationCapabilityCodeListType is Helper simpleType to allow for enumerated lists in FIXM.
type DatalinkCommunicationCapabilityCodeListType []string

// DinghyCoverIndicatorType is Indication of the covered/uncovered nature of the dinghies carried by the aircraft. [FIXM]
type DinghyCoverIndicatorType string

// EltHexIdentifierType is The identifier of an emergency locator transmitter. [FIXM]
type EltHexIdentifierType interface{}

// EltHexIdentifierListType is Helper simpleType that allows representation of emergency locator transmitter hexadecimal identifiers as a list.
type EltHexIdentifierListType []interface{}

// EmergencyRadioCapabilityTypeType is VHF on frequency 121.5 MHz is available. [adapted from ICAO Doc 4444, Appendix 2, ITEM 19 R/]
type EmergencyRadioCapabilityTypeType string

// EmergencyRadioCapabilityTypeListType is Helper simpleType to allow for enumerated lists in FIXM.
type EmergencyRadioCapabilityTypeListType []string

// LifeJacketTypeType is VHF on frequency 121.5 MHz is available. [adapted from ICAO Doc 4444, Appendix 2, ITEM 19 R/ and J/]
type LifeJacketTypeType string

// LifeJacketTypeListType is Helper simpleType to allow for enumerated lists in FIXM.
type LifeJacketTypeListType []string

// NavigationCapabilityCodeType is MNPS
type NavigationCapabilityCodeType string

// NavigationCapabilityCodeListType is Helper simpleType to allow for enumerated lists in FIXM.
type NavigationCapabilityCodeListType []string

// PerformanceBasedNavigationCapabilityCodeType is RNP AR APCH without RF (Authorization Required)
type PerformanceBasedNavigationCapabilityCodeType string

// PerformanceBasedNavigationCapabilityCodeListType is Helper simpleType to allow for enumerated lists in FIXM.
type PerformanceBasedNavigationCapabilityCodeListType []string

// SelectiveCallingCodeType is A code that consists of two 2-letter pairs and acts as a paging system for an ATS unit to establish voice communications with the pilot of an aircraft.
type SelectiveCallingCodeType interface{}

// StandardCapabilitiesIndicatorType is This element indicates the aircraft carries the set of capabilities considered standard by the appropriate authority.
type StandardCapabilitiesIndicatorType string

// SurveillanceCapabilityCodeType is Transponder Mode S with neither aircraft identification nor pressure-altitude capability
type SurveillanceCapabilityCodeType string

// SurveillanceCapabilityCodeListType is Helper simpleType to allow for enumerated lists in FIXM.
type SurveillanceCapabilityCodeListType []string

// SurvivalEquipmentTypeType is Polar survival equipment is carried. [ICAO Doc 4444, Appendix 3]
type SurvivalEquipmentTypeType string

// SurvivalEquipmentTypeListType is Helper simpleType to allow for enumerated lists in FIXM.
type SurvivalEquipmentTypeListType []string

// CommunicationCapabilitiesType is The serviceable communications equipment, available on the aircraft at the time of flight, and associated flight crew qualifications that may be used to communicate with ATS units. [FIXM]
type CommunicationCapabilitiesType struct {
	CommunicationCapabilityCode         interface{}                                    `xml:"communicationCapabilityCode"`
	DatalinkCommunicationCapabilityCode interface{}                                    `xml:"datalinkCommunicationCapabilityCode"`
	Extension                           []*base.CommunicationCapabilitiesExtensionType `xml:"extension"`
	OtherCommunicationCapabilities      *base.CharacterStringType                      `xml:"otherCommunicationCapabilities"`
	OtherDatalinkCapabilities           *base.CharacterStringType                      `xml:"otherDatalinkCapabilities"`
	SelectiveCallingCode                interface{}                                    `xml:"selectiveCallingCode"`
}

// DinghiesType ...
type DinghiesType struct {
	Colour        *base.CharacterStringType     `xml:"colour"`
	Covered       string                        `xml:"covered"`
	Extension     []*base.DinghiesExtensionType `xml:"extension"`
	Number        *base.CountType               `xml:"number"`
	TotalCapacity *base.CountPositiveType       `xml:"totalCapacity"`
}

// FlightCapabilitiesType ...
type FlightCapabilitiesType struct {
	Communication        interface{}                             `xml:"communication"`
	Extension            []*base.FlightCapabilitiesExtensionType `xml:"extension"`
	Navigation           interface{}                             `xml:"navigation"`
	StandardCapabilities string                                  `xml:"standardCapabilities"`
	Surveillance         interface{}                             `xml:"surveillance"`
	Survival             interface{}                             `xml:"survival"`
}

// NavigationCapabilitiesType ...
type NavigationCapabilitiesType struct {
	Extension                   []*base.NavigationCapabilitiesExtensionType `xml:"extension"`
	NavigationCapabilityCode    interface{}                                 `xml:"navigationCapabilityCode"`
	OtherNavigationCapabilities *base.CharacterStringType                   `xml:"otherNavigationCapabilities"`
	PerformanceBasedCode        interface{}                                 `xml:"performanceBasedCode"`
	RequiredRunwayVisualRange   *base.DistanceType                          `xml:"requiredRunwayVisualRange"`
}

// SurveillanceCapabilitiesType ...
type SurveillanceCapabilitiesType struct {
	Extension                     []*base.SurveillanceCapabilitiesExtensionType `xml:"extension"`
	OtherSurveillanceCapabilities *base.CharacterStringType                     `xml:"otherSurveillanceCapabilities"`
	SurveillanceCapabilityCode    interface{}                                   `xml:"surveillanceCapabilityCode"`
}

// SurvivalCapabilitiesType ...
type SurvivalCapabilitiesType struct {
	CarriedEltHexIdentifier      interface{}                               `xml:"carriedEltHexIdentifier"`
	DinghyInformation            interface{}                               `xml:"dinghyInformation"`
	EmergencyRadioCapabilityType interface{}                               `xml:"emergencyRadioCapabilityType"`
	Extension                    []*base.SurvivalCapabilitiesExtensionType `xml:"extension"`
	LifeJacketType               interface{}                               `xml:"lifeJacketType"`
	SurvivalEquipmentRemarks     *base.CharacterStringType                 `xml:"survivalEquipmentRemarks"`
	SurvivalEquipmentType        interface{}                               `xml:"survivalEquipmentType"`
}
