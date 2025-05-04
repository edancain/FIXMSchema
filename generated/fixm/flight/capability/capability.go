// Code generated from Capability.xsd; DO NOT EDIT.

package capability

import (
	"github.com/edancain/FIXMSchema.git/generated/fixm/base"
)

// A code indicating the capability of the aircraft and associated flight crew q...
type CommunicationCapabilityCodeType string

const (
	CommunicationCapabilityCodeTypeE1 CommunicationCapabilityCodeType = "E1"
	CommunicationCapabilityCodeTypeE2 CommunicationCapabilityCodeType = "E2"
	CommunicationCapabilityCodeTypeE3 CommunicationCapabilityCodeType = "E3"
	CommunicationCapabilityCodeTypeH CommunicationCapabilityCodeType = "H"
	CommunicationCapabilityCodeTypeM1 CommunicationCapabilityCodeType = "M1"
	CommunicationCapabilityCodeTypeM2 CommunicationCapabilityCodeType = "M2"
	CommunicationCapabilityCodeTypeM3 CommunicationCapabilityCodeType = "M3"
	CommunicationCapabilityCodeTypeP1 CommunicationCapabilityCodeType = "P1"
	CommunicationCapabilityCodeTypeP2 CommunicationCapabilityCodeType = "P2"
	CommunicationCapabilityCodeTypeP3 CommunicationCapabilityCodeType = "P3"
	CommunicationCapabilityCodeTypeP4 CommunicationCapabilityCodeType = "P4"
	CommunicationCapabilityCodeTypeP5 CommunicationCapabilityCodeType = "P5"
	CommunicationCapabilityCodeTypeP6 CommunicationCapabilityCodeType = "P6"
	CommunicationCapabilityCodeTypeP7 CommunicationCapabilityCodeType = "P7"
	CommunicationCapabilityCodeTypeP8 CommunicationCapabilityCodeType = "P8"
	CommunicationCapabilityCodeTypeP9 CommunicationCapabilityCodeType = "P9"
	CommunicationCapabilityCodeTypeU CommunicationCapabilityCodeType = "U"
	CommunicationCapabilityCodeTypeV CommunicationCapabilityCodeType = "V"
	CommunicationCapabilityCodeTypeY CommunicationCapabilityCodeType = "Y"
)

// Helper simpleType to allow for enumerated lists in FIXM.
type CommunicationCapabilityCodeListType []flight.CommunicationCapabilityCodeType

// A code indicating the capability of the aircraft to communicate or receive da...
type DatalinkCommunicationCapabilityCodeType string

const (
	DatalinkCommunicationCapabilityCodeTypeJ1 DatalinkCommunicationCapabilityCodeType = "J1"
	DatalinkCommunicationCapabilityCodeTypeJ2 DatalinkCommunicationCapabilityCodeType = "J2"
	DatalinkCommunicationCapabilityCodeTypeJ3 DatalinkCommunicationCapabilityCodeType = "J3"
	DatalinkCommunicationCapabilityCodeTypeJ4 DatalinkCommunicationCapabilityCodeType = "J4"
	DatalinkCommunicationCapabilityCodeTypeJ5 DatalinkCommunicationCapabilityCodeType = "J5"
	DatalinkCommunicationCapabilityCodeTypeJ6 DatalinkCommunicationCapabilityCodeType = "J6"
	DatalinkCommunicationCapabilityCodeTypeJ7 DatalinkCommunicationCapabilityCodeType = "J7"
)

// Helper simpleType to allow for enumerated lists in FIXM.
type DatalinkCommunicationCapabilityCodeListType []flight.DatalinkCommunicationCapabilityCodeType

// Indication of the covered/uncovered nature of the dinghies carried by the air...
type DinghyCoverIndicatorType string

const (
	DinghyCoverIndicatorTypeCOVERED DinghyCoverIndicatorType = "COVERED"
)

// The identifier of an emergency locator transmitter. [FIXM]
type EltHexIdentifierType base.CharacterStringType

// Helper simpleType that allows representation of emergency locator transmitter...
type EltHexIdentifierListType []flight.EltHexIdentifierType

// The type of serviceable communication devices available on the aircraft that ...
type EmergencyRadioCapabilityTypeType string

const (
	EmergencyRadioCapabilityTypeTypeEMERGENCY_LOCATOR_TRANSMITTER EmergencyRadioCapabilityTypeType = "EMERGENCY_LOCATOR_TRANSMITTER"
	EmergencyRadioCapabilityTypeTypeULTRA_HIGH_FREQUENCY EmergencyRadioCapabilityTypeType = "ULTRA_HIGH_FREQUENCY"
	EmergencyRadioCapabilityTypeTypeVERY_HIGH_FREQUENCY EmergencyRadioCapabilityTypeType = "VERY_HIGH_FREQUENCY"
)

// Helper simpleType to allow for enumerated lists in FIXM.
type EmergencyRadioCapabilityTypeListType []flight.EmergencyRadioCapabilityTypeType

// The type of life jackets available on board the aircraft. [FIXM]
type LifeJacketTypeType string

const (
	LifeJacketTypeTypeFLUORESCENCE LifeJacketTypeType = "FLUORESCENCE"
	LifeJacketTypeTypeLIGHTS LifeJacketTypeType = "LIGHTS"
	LifeJacketTypeTypeULTRA_HIGH_FREQUENCY LifeJacketTypeType = "ULTRA_HIGH_FREQUENCY"
	LifeJacketTypeTypeVERY_HIGH_FREQUENCY LifeJacketTypeType = "VERY_HIGH_FREQUENCY"
)

// Helper simpleType to allow for enumerated lists in FIXM.
type LifeJacketTypeListType []flight.LifeJacketTypeType

// A code indicating the navigation capability of the aircraft and associated fl...
type NavigationCapabilityCodeType string

const (
	NavigationCapabilityCodeTypeA NavigationCapabilityCodeType = "A"
	NavigationCapabilityCodeTypeB NavigationCapabilityCodeType = "B"
	NavigationCapabilityCodeTypeC NavigationCapabilityCodeType = "C"
	NavigationCapabilityCodeTypeD NavigationCapabilityCodeType = "D"
	NavigationCapabilityCodeTypeF NavigationCapabilityCodeType = "F"
	NavigationCapabilityCodeTypeG NavigationCapabilityCodeType = "G"
	NavigationCapabilityCodeTypeI NavigationCapabilityCodeType = "I"
	NavigationCapabilityCodeTypeK NavigationCapabilityCodeType = "K"
	NavigationCapabilityCodeTypeL NavigationCapabilityCodeType = "L"
	NavigationCapabilityCodeTypeO NavigationCapabilityCodeType = "O"
	NavigationCapabilityCodeTypeT NavigationCapabilityCodeType = "T"
	NavigationCapabilityCodeTypeW NavigationCapabilityCodeType = "W"
	NavigationCapabilityCodeTypeX NavigationCapabilityCodeType = "X"
)

// Helper simpleType to allow for enumerated lists in FIXM.
type NavigationCapabilityCodeListType []flight.NavigationCapabilityCodeType

// A coded category denoting which Required Navigation Performance (RNP) and Are...
type PerformanceBasedNavigationCapabilityCodeType string

const (
	PerformanceBasedNavigationCapabilityCodeTypeA1 PerformanceBasedNavigationCapabilityCodeType = "A1"
	PerformanceBasedNavigationCapabilityCodeTypeB1 PerformanceBasedNavigationCapabilityCodeType = "B1"
	PerformanceBasedNavigationCapabilityCodeTypeB2 PerformanceBasedNavigationCapabilityCodeType = "B2"
	PerformanceBasedNavigationCapabilityCodeTypeB3 PerformanceBasedNavigationCapabilityCodeType = "B3"
	PerformanceBasedNavigationCapabilityCodeTypeB4 PerformanceBasedNavigationCapabilityCodeType = "B4"
	PerformanceBasedNavigationCapabilityCodeTypeB5 PerformanceBasedNavigationCapabilityCodeType = "B5"
	PerformanceBasedNavigationCapabilityCodeTypeB6 PerformanceBasedNavigationCapabilityCodeType = "B6"
	PerformanceBasedNavigationCapabilityCodeTypeC1 PerformanceBasedNavigationCapabilityCodeType = "C1"
	PerformanceBasedNavigationCapabilityCodeTypeC2 PerformanceBasedNavigationCapabilityCodeType = "C2"
	PerformanceBasedNavigationCapabilityCodeTypeC3 PerformanceBasedNavigationCapabilityCodeType = "C3"
	PerformanceBasedNavigationCapabilityCodeTypeC4 PerformanceBasedNavigationCapabilityCodeType = "C4"
	PerformanceBasedNavigationCapabilityCodeTypeD1 PerformanceBasedNavigationCapabilityCodeType = "D1"
	PerformanceBasedNavigationCapabilityCodeTypeD2 PerformanceBasedNavigationCapabilityCodeType = "D2"
	PerformanceBasedNavigationCapabilityCodeTypeD3 PerformanceBasedNavigationCapabilityCodeType = "D3"
	PerformanceBasedNavigationCapabilityCodeTypeD4 PerformanceBasedNavigationCapabilityCodeType = "D4"
	PerformanceBasedNavigationCapabilityCodeTypeL1 PerformanceBasedNavigationCapabilityCodeType = "L1"
	PerformanceBasedNavigationCapabilityCodeTypeO1 PerformanceBasedNavigationCapabilityCodeType = "O1"
	PerformanceBasedNavigationCapabilityCodeTypeO2 PerformanceBasedNavigationCapabilityCodeType = "O2"
	PerformanceBasedNavigationCapabilityCodeTypeO3 PerformanceBasedNavigationCapabilityCodeType = "O3"
	PerformanceBasedNavigationCapabilityCodeTypeO4 PerformanceBasedNavigationCapabilityCodeType = "O4"
	PerformanceBasedNavigationCapabilityCodeTypeS1 PerformanceBasedNavigationCapabilityCodeType = "S1"
	PerformanceBasedNavigationCapabilityCodeTypeS2 PerformanceBasedNavigationCapabilityCodeType = "S2"
	PerformanceBasedNavigationCapabilityCodeTypeT1 PerformanceBasedNavigationCapabilityCodeType = "T1"
	PerformanceBasedNavigationCapabilityCodeTypeT2 PerformanceBasedNavigationCapabilityCodeType = "T2"
)

// Helper simpleType to allow for enumerated lists in FIXM.
type PerformanceBasedNavigationCapabilityCodeListType []flight.PerformanceBasedNavigationCapabilityCodeType

// A code that consists of two 2-letter pairs and acts as a paging system for an...
type SelectiveCallingCodeType base.CharacterStringType

// This element indicates the aircraft carries the set of capabilities considere...
type StandardCapabilitiesIndicatorType string

const (
	StandardCapabilitiesIndicatorTypeSTANDARD StandardCapabilitiesIndicatorType = "STANDARD"
)

// A code indicating the SSR and ADS capability of the aircraft. [FIXM]
type SurveillanceCapabilityCodeType string

const (
	SurveillanceCapabilityCodeTypeA SurveillanceCapabilityCodeType = "A"
	SurveillanceCapabilityCodeTypeB1 SurveillanceCapabilityCodeType = "B1"
	SurveillanceCapabilityCodeTypeB2 SurveillanceCapabilityCodeType = "B2"
	SurveillanceCapabilityCodeTypeC SurveillanceCapabilityCodeType = "C"
	SurveillanceCapabilityCodeTypeD1 SurveillanceCapabilityCodeType = "D1"
	SurveillanceCapabilityCodeTypeE SurveillanceCapabilityCodeType = "E"
	SurveillanceCapabilityCodeTypeG1 SurveillanceCapabilityCodeType = "G1"
	SurveillanceCapabilityCodeTypeH SurveillanceCapabilityCodeType = "H"
	SurveillanceCapabilityCodeTypeI SurveillanceCapabilityCodeType = "I"
	SurveillanceCapabilityCodeTypeL SurveillanceCapabilityCodeType = "L"
	SurveillanceCapabilityCodeTypeP SurveillanceCapabilityCodeType = "P"
	SurveillanceCapabilityCodeTypeS SurveillanceCapabilityCodeType = "S"
	SurveillanceCapabilityCodeTypeU1 SurveillanceCapabilityCodeType = "U1"
	SurveillanceCapabilityCodeTypeU2 SurveillanceCapabilityCodeType = "U2"
	SurveillanceCapabilityCodeTypeV1 SurveillanceCapabilityCodeType = "V1"
	SurveillanceCapabilityCodeTypeV2 SurveillanceCapabilityCodeType = "V2"
	SurveillanceCapabilityCodeTypeX SurveillanceCapabilityCodeType = "X"
)

// Helper simpleType to allow for enumerated lists in FIXM.
type SurveillanceCapabilityCodeListType []flight.SurveillanceCapabilityCodeType

// The type of equipment carried on board the aircraft that can be used by the c...
type SurvivalEquipmentTypeType string

const (
	SurvivalEquipmentTypeTypeDESERT SurvivalEquipmentTypeType = "DESERT"
	SurvivalEquipmentTypeTypeJUNGLE SurvivalEquipmentTypeType = "JUNGLE"
	SurvivalEquipmentTypeTypeMARITIME SurvivalEquipmentTypeType = "MARITIME"
	SurvivalEquipmentTypeTypePOLAR SurvivalEquipmentTypeType = "POLAR"
)

// Helper simpleType to allow for enumerated lists in FIXM.
type SurvivalEquipmentTypeListType []flight.SurvivalEquipmentTypeType

// The serviceable communications equipment, available on the aircraft at the ti...
type CommunicationCapabilitiesType struct {
	// A code indicating the capability of the aircraft and associated flight crew q...
	CommunicationCapabilityCode *flight.CommunicationCapabilityCodeListType `xml:"communicationCapabilityCode"`
	// A code indicating the capability of the aircraft to communicate or receive da...
	DatalinkCommunicationCapabilityCode *flight.DatalinkCommunicationCapabilityCodeListType `xml:"datalinkCommunicationCapabilityCode"`
	// An extension hook for attaching extension (non-core) classes.
	Extension []base.CommunicationCapabilitiesExtensionType `xml:"extension"`
	// Additional Communication capabilities available on the aircraft.
	OtherCommunicationCapabilities *base.CharacterStringType `xml:"otherCommunicationCapabilities"`
	// Additional data link capabilities available on the aircraft.
	OtherDatalinkCapabilities *base.CharacterStringType `xml:"otherDatalinkCapabilities"`
	// A code that consists of two 2-letter pairs and acts as a paging system for an...
	SelectiveCallingCode *flight.SelectiveCallingCodeType `xml:"selectiveCallingCode"`
}

// Dinghy information including total capacity in persons of all dinghies carrie...
type DinghiesType struct {
	// The colour of the dinghies carried by the aircraft. [FIXM]
	Colour *base.CharacterStringType `xml:"colour"`
	// Indication of the covered/uncovered nature of the dinghies carried by the air...
	Covered *flight.DinghyCoverIndicatorType `xml:"covered"`
	// An extension hook for attaching extension (non-core) classes.
	Extension []base.DinghiesExtensionType `xml:"extension"`
	// The number of dinghies carried by the aircraft. [adapted from ICAO Doc 4444, ...
	Number *base.CountType `xml:"number"`
	// Total capacity, in persons, of all dinghies carried. [ICAO Doc 4444, Appendix...
	TotalCapacity *base.CountPositiveType `xml:"totalCapacity"`
}

// The capabilities of the flight comprising of the: a) presence of relevant ser...
type FlightCapabilitiesType struct {
	// The serviceable communications equipment, available on the aircraft at the ti...
	Communication *flight.CommunicationCapabilitiesType `xml:"communication"`
	// An extension hook for attaching extension (non-core) classes.
	Extension []base.FlightCapabilitiesExtensionType `xml:"extension"`
	// The navigation capability of the aircraft and associated flight crew qualific...
	Navigation *flight.NavigationCapabilitiesType `xml:"navigation"`
	// If present, indicates that aircraft has the "standard" capabilities for the f...
	StandardCapabilities *flight.StandardCapabilitiesIndicatorType `xml:"standardCapabilities"`
	// The serviceable Secondary Surveillance Radar (SSR) and/or Automatic Dependent...
	Surveillance *flight.SurveillanceCapabilitiesType `xml:"surveillance"`
	// The emergency and survival equipments available on board the aircraft. [FIXM]
	Survival *flight.SurvivalCapabilitiesType `xml:"survival"`
}

// The navigation capability of the aircraft and associated flight crew qualific...
type NavigationCapabilitiesType struct {
	// An extension hook for attaching extension (non-core) classes.
	Extension []base.NavigationCapabilitiesExtensionType `xml:"extension"`
	// A code indicating the navigation capability of the aircraft and associated fl...
	NavigationCapabilityCode *flight.NavigationCapabilityCodeListType `xml:"navigationCapabilityCode"`
	// Additional navigation capabilities available on board the aircraft.
	OtherNavigationCapabilities *base.CharacterStringType `xml:"otherNavigationCapabilities"`
	// A coded category denoting which Required Navigation Performance (RNP) and Are...
	PerformanceBasedCode *flight.PerformanceBasedNavigationCapabilityCodeListType `xml:"performanceBasedCode"`
	// The minimum RVR value required by a flight in order to execute an approach to...
	RequiredRunwayVisualRange *base.DistanceType `xml:"requiredRunwayVisualRange"`
}

// The serviceable Secondary Surveillance Radar (SSR) and/or Automatic Dependent...
type SurveillanceCapabilitiesType struct {
	// An extension hook for attaching extension (non-core) classes.
	Extension []base.SurveillanceCapabilitiesExtensionType `xml:"extension"`
	// Additional surveillance capabilities available on board the aircraft.
	OtherSurveillanceCapabilities *base.CharacterStringType `xml:"otherSurveillanceCapabilities"`
	// A code indicating the SSR and ADS capability of the aircraft. [FIXM]
	SurveillanceCapabilityCode *flight.SurveillanceCapabilityCodeListType `xml:"surveillanceCapabilityCode"`
}

// The emergency and survival equipments available on board the aircraft. [FIXM]
type SurvivalCapabilitiesType struct {
	// The identifier of an emergency locator transmitter carried by aircraft. [FIXM]
	CarriedEltHexIdentifier *flight.EltHexIdentifierListType `xml:"carriedEltHexIdentifier"`
	// Dinghy information including total capacity in persons of all dinghies carrie...
	DinghyInformation *flight.DinghiesType `xml:"dinghyInformation"`
	// The type of serviceable communication devices available on the aircraft that ...
	EmergencyRadioCapabilityType *flight.EmergencyRadioCapabilityTypeListType `xml:"emergencyRadioCapabilityType"`
	// An extension hook for attaching extension (non-core) classes.
	Extension []base.SurvivalCapabilitiesExtensionType `xml:"extension"`
	// The type of life jackets available on board the aircraft. [FIXM]
	LifeJacketType *flight.LifeJacketTypeListType `xml:"lifeJacketType"`
	// A description of survival equipment carried on the aircraft and any other use...
	SurvivalEquipmentRemarks *base.CharacterStringType `xml:"survivalEquipmentRemarks"`
	// The type of equipment carried on board the aircraft that can be used by the c...
	SurvivalEquipmentType *flight.SurvivalEquipmentTypeListType `xml:"survivalEquipmentType"`
}

