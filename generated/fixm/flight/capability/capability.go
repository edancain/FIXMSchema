// Code generated from Capability.xsd; DO NOT EDIT.

package capability

import (
	"github.com/edancain/FIXMSchema.git/generated/fixm/base"
)

// A code indicating the capability of the aircraft and associated flight crew qualifications to communicate with ATS units. [FIXM]
type CommunicationCapabilityCodeType string

const (
	CommunicationCapabilityCodeTypeE1 CommunicationCapabilityCodeType = "E1" // FMC WPR ACARS
	CommunicationCapabilityCodeTypeE2 CommunicationCapabilityCodeType = "E2" // D-FIS ACARS
	CommunicationCapabilityCodeTypeE3 CommunicationCapabilityCodeType = "E3" // PDC ACARS
	CommunicationCapabilityCodeTypeH CommunicationCapabilityCodeType = "H" // HF RTF
	CommunicationCapabilityCodeTypeM1 CommunicationCapabilityCodeType = "M1" // ATC RTF SATCOM (INMARSAT)
	CommunicationCapabilityCodeTypeM2 CommunicationCapabilityCodeType = "M2" // ATC RTF (MTSAT)
	CommunicationCapabilityCodeTypeM3 CommunicationCapabilityCodeType = "M3" // ATC RTF (Iridium)
	CommunicationCapabilityCodeTypeP1 CommunicationCapabilityCodeType = "P1" // Reserved for RCP
	CommunicationCapabilityCodeTypeP2 CommunicationCapabilityCodeType = "P2" // Reserved for RCP
	CommunicationCapabilityCodeTypeP3 CommunicationCapabilityCodeType = "P3" // Reserved for RCP
	CommunicationCapabilityCodeTypeP4 CommunicationCapabilityCodeType = "P4" // Reserved for RCP
	CommunicationCapabilityCodeTypeP5 CommunicationCapabilityCodeType = "P5" // Reserved for RCP
	CommunicationCapabilityCodeTypeP6 CommunicationCapabilityCodeType = "P6" // Reserved for RCP
	CommunicationCapabilityCodeTypeP7 CommunicationCapabilityCodeType = "P7" // Reserved for RCP
	CommunicationCapabilityCodeTypeP8 CommunicationCapabilityCodeType = "P8" // Reserved for RCP
	CommunicationCapabilityCodeTypeP9 CommunicationCapabilityCodeType = "P9" // Reserved for RCP
	CommunicationCapabilityCodeTypeU CommunicationCapabilityCodeType = "U" // UHF RTF
	CommunicationCapabilityCodeTypeV CommunicationCapabilityCodeType = "V" // VHF RTF
	CommunicationCapabilityCodeTypeY CommunicationCapabilityCodeType = "Y" // VHF with 8.33 kHz channel spacing capability
)

// Helper simpleType to allow for enumerated lists in FIXM.
type CommunicationCapabilityCodeListType []CommunicationCapabilityCodeType

// A code indicating the capability of the aircraft to communicate or receive data [FIXM]
type DatalinkCommunicationCapabilityCodeType string

const (
	DatalinkCommunicationCapabilityCodeTypeJ1 DatalinkCommunicationCapabilityCodeType = "J1" // CPDLC ATN VDL Mode 2 
	DatalinkCommunicationCapabilityCodeTypeJ2 DatalinkCommunicationCapabilityCodeType = "J2" // CPDLC FANS 1/A HFDL 
	DatalinkCommunicationCapabilityCodeTypeJ3 DatalinkCommunicationCapabilityCodeType = "J3" // CPDLC FANS 1/A VDL Mode A 
	DatalinkCommunicationCapabilityCodeTypeJ4 DatalinkCommunicationCapabilityCodeType = "J4" // CPDLC FANS 1/A VDL Mode 2 
	DatalinkCommunicationCapabilityCodeTypeJ5 DatalinkCommunicationCapabilityCodeType = "J5" // CPDLC FANS 1/A SATCOM (INMARSAT)
	DatalinkCommunicationCapabilityCodeTypeJ6 DatalinkCommunicationCapabilityCodeType = "J6" // CPDLC FANS 1/A SATCOM (MTSAT) 
	DatalinkCommunicationCapabilityCodeTypeJ7 DatalinkCommunicationCapabilityCodeType = "J7" // CPDLC FANS 1/A SATCOM (Iridium) 
)

// Helper simpleType to allow for enumerated lists in FIXM.
type DatalinkCommunicationCapabilityCodeListType []DatalinkCommunicationCapabilityCodeType

// Indication of the covered/uncovered nature of the dinghies carried by the aircraft [FIXM]
type DinghyCoverIndicatorType string

const (
	// Indication of the covered/uncovered nature of the dinghies carried by the aircraft. [FIXM]
	DinghyCoverIndicatorTypeCOVERED DinghyCoverIndicatorType = "COVERED"
)

// The identifier of an emergency locator transmitter. [FIXM]
type EltHexIdentifierType base.CharacterStringType

// Helper simpleType that allows representation of emergency locator transmitter hexadecimal identifiers as a list.
type EltHexIdentifierListType []EltHexIdentifierType

// The type of serviceable communication devices available on the aircraft that are able to transmit an emergency radio signal. [FIXM]
type EmergencyRadioCapabilityTypeType string

const (
	EmergencyRadioCapabilityTypeTypeEMERGENCY_LOCATOR_TRANSMITTER EmergencyRadioCapabilityTypeType = "EMERGENCY_LOCATOR_TRANSMITTER" // Emergency locator transmitter (ELT) is available. [adapted from ICAO Doc 4444, Appendix 2, ITEM 19 R/]
	EmergencyRadioCapabilityTypeTypeULTRA_HIGH_FREQUENCY EmergencyRadioCapabilityTypeType = "ULTRA_HIGH_FREQUENCY" // UHF on frequency 243.0 MHz is available. [adapted from ICAO Doc 4444, Appendix 2, ITEM 19 R/]
	EmergencyRadioCapabilityTypeTypeVERY_HIGH_FREQUENCY EmergencyRadioCapabilityTypeType = "VERY_HIGH_FREQUENCY" // VHF on frequency 121.5 MHz is available. [adapted from ICAO Doc 4444, Appendix 2, ITEM 19 R/]
)

// Helper simpleType to allow for enumerated lists in FIXM.
type EmergencyRadioCapabilityTypeListType []EmergencyRadioCapabilityTypeType

// The type of life jackets available on board the aircraft. [FIXM]
type LifeJacketTypeType string

const (
	LifeJacketTypeTypeFLUORESCENCE LifeJacketTypeType = "FLUORESCENCE" // Life jackets are equipped with fluorescence. [adapted from ICAO Doc 4444, Appendix 2, ITEM 19 J/]
	LifeJacketTypeTypeLIGHTS LifeJacketTypeType = "LIGHTS" // Life jackets are equipped with lights. [adapted from ICAO Doc 4444, Appendix 2, ITEM 19 J/]
	LifeJacketTypeTypeULTRA_HIGH_FREQUENCY LifeJacketTypeType = "ULTRA_HIGH_FREQUENCY" // UHF on frequency 243.0 MHz is available. [adapted from ICAO Doc 4444, Appendix 2, ITEM 19 R/ and J/]
	LifeJacketTypeTypeVERY_HIGH_FREQUENCY LifeJacketTypeType = "VERY_HIGH_FREQUENCY" // VHF on frequency 121.5 MHz is available. [adapted from ICAO Doc 4444, Appendix 2, ITEM 19 R/ and J/]
)

// Helper simpleType to allow for enumerated lists in FIXM.
type LifeJacketTypeListType []LifeJacketTypeType

// A code indicating the navigation capability of the aircraft and associated flight crew qualifications. [FIXM]
type NavigationCapabilityCodeType string

const (
	NavigationCapabilityCodeTypeA NavigationCapabilityCodeType = "A" // GBAS
	NavigationCapabilityCodeTypeB NavigationCapabilityCodeType = "B" // LPV
	NavigationCapabilityCodeTypeC NavigationCapabilityCodeType = "C" // LORAN C
	NavigationCapabilityCodeTypeD NavigationCapabilityCodeType = "D" // DME
	NavigationCapabilityCodeTypeF NavigationCapabilityCodeType = "F" // ADF
	NavigationCapabilityCodeTypeG NavigationCapabilityCodeType = "G" // GNSS
	NavigationCapabilityCodeTypeI NavigationCapabilityCodeType = "I" // Inertial Navigation 
	NavigationCapabilityCodeTypeK NavigationCapabilityCodeType = "K" // MLS
	NavigationCapabilityCodeTypeL NavigationCapabilityCodeType = "L" // ILS
	NavigationCapabilityCodeTypeO NavigationCapabilityCodeType = "O" // VOR
	NavigationCapabilityCodeTypeT NavigationCapabilityCodeType = "T" // TACAN
	NavigationCapabilityCodeTypeW NavigationCapabilityCodeType = "W" // RVSM
	NavigationCapabilityCodeTypeX NavigationCapabilityCodeType = "X" // MNPS
)

// Helper simpleType to allow for enumerated lists in FIXM.
type NavigationCapabilityCodeListType []NavigationCapabilityCodeType

// A coded category denoting which Required Navigation Performance (RNP) and Area Navigation (RNAV) requirements can be met by the aircraft while operating in the context of a particular airspace when supported by the appropriate navigation infrastructure. [FIXM]
type PerformanceBasedNavigationCapabilityCodeType string

const (
	PerformanceBasedNavigationCapabilityCodeTypeA1 PerformanceBasedNavigationCapabilityCodeType = "A1" // RNAV 10 (RNP 10)
	PerformanceBasedNavigationCapabilityCodeTypeB1 PerformanceBasedNavigationCapabilityCodeType = "B1" // RNAV 5 All Permitted Sensors 
	PerformanceBasedNavigationCapabilityCodeTypeB2 PerformanceBasedNavigationCapabilityCodeType = "B2" // RNAV 5 GNSS 
	PerformanceBasedNavigationCapabilityCodeTypeB3 PerformanceBasedNavigationCapabilityCodeType = "B3" // RNAV 5 DME/DME 
	PerformanceBasedNavigationCapabilityCodeTypeB4 PerformanceBasedNavigationCapabilityCodeType = "B4" // RNAV 5 VOR/DME 
	PerformanceBasedNavigationCapabilityCodeTypeB5 PerformanceBasedNavigationCapabilityCodeType = "B5" // RNAV 5 INS or IRS 
	PerformanceBasedNavigationCapabilityCodeTypeB6 PerformanceBasedNavigationCapabilityCodeType = "B6" // RNAV 5 LORAN-C 
	PerformanceBasedNavigationCapabilityCodeTypeC1 PerformanceBasedNavigationCapabilityCodeType = "C1" // RNAV 2 All Permitted Sensors 
	PerformanceBasedNavigationCapabilityCodeTypeC2 PerformanceBasedNavigationCapabilityCodeType = "C2" // RNAV 2 GNSS 
	PerformanceBasedNavigationCapabilityCodeTypeC3 PerformanceBasedNavigationCapabilityCodeType = "C3" // RNAV 2 DME/DME 
	PerformanceBasedNavigationCapabilityCodeTypeC4 PerformanceBasedNavigationCapabilityCodeType = "C4" // RNAV 2 DME/DME/IRU 
	PerformanceBasedNavigationCapabilityCodeTypeD1 PerformanceBasedNavigationCapabilityCodeType = "D1" // RNAV 1 All Permitted Sensors 
	PerformanceBasedNavigationCapabilityCodeTypeD2 PerformanceBasedNavigationCapabilityCodeType = "D2" // RNAV 1 GNSS 
	PerformanceBasedNavigationCapabilityCodeTypeD3 PerformanceBasedNavigationCapabilityCodeType = "D3" // RNAV 1 DME/DME 
	PerformanceBasedNavigationCapabilityCodeTypeD4 PerformanceBasedNavigationCapabilityCodeType = "D4" // RNAV 1 DME/DME/IRU 
	PerformanceBasedNavigationCapabilityCodeTypeL1 PerformanceBasedNavigationCapabilityCodeType = "L1" // RNP 4 
	PerformanceBasedNavigationCapabilityCodeTypeO1 PerformanceBasedNavigationCapabilityCodeType = "O1" // Basic RNP 1 All Permitted Sensors 
	PerformanceBasedNavigationCapabilityCodeTypeO2 PerformanceBasedNavigationCapabilityCodeType = "O2" // Basic RNP 1 GNSS 
	PerformanceBasedNavigationCapabilityCodeTypeO3 PerformanceBasedNavigationCapabilityCodeType = "O3" // Basic RNP 1 DME/DME 
	PerformanceBasedNavigationCapabilityCodeTypeO4 PerformanceBasedNavigationCapabilityCodeType = "O4" // Basic RNP 1 DME/DME/IRU 
	PerformanceBasedNavigationCapabilityCodeTypeS1 PerformanceBasedNavigationCapabilityCodeType = "S1" // RNP APCH
	PerformanceBasedNavigationCapabilityCodeTypeS2 PerformanceBasedNavigationCapabilityCodeType = "S2" // RNP APCH with Barometric Vertical Navigation 
	PerformanceBasedNavigationCapabilityCodeTypeT1 PerformanceBasedNavigationCapabilityCodeType = "T1" // RNP AR APCH with RF (Authorization Required) 
	PerformanceBasedNavigationCapabilityCodeTypeT2 PerformanceBasedNavigationCapabilityCodeType = "T2" // RNP AR APCH without RF (Authorization Required)
)

// Helper simpleType to allow for enumerated lists in FIXM.
type PerformanceBasedNavigationCapabilityCodeListType []PerformanceBasedNavigationCapabilityCodeType

// A code that consists of two 2-letter pairs and acts as a paging system for an ATS unit to establish voice communications with the pilot of an aircraft.
type SelectiveCallingCodeType base.CharacterStringType

// This element indicates the aircraft carries the set of capabilities considered standard by the appropriate authority.
type StandardCapabilitiesIndicatorType string

const (
	StandardCapabilitiesIndicatorTypeSTANDARD StandardCapabilitiesIndicatorType = "STANDARD" // This element indicates the aircraft carries the set of capabilities considered standard by the appropriate authority.
)

// A code indicating the SSR and ADS capability of the aircraft. [FIXM]
type SurveillanceCapabilityCodeType string

const (
	SurveillanceCapabilityCodeTypeA SurveillanceCapabilityCodeType = "A" // Transponder-Mode A (4 digits-4,096 codes) 
	SurveillanceCapabilityCodeTypeB1 SurveillanceCapabilityCodeType = "B1" // ADS-B with dedicated 1090 MHz ADS-B out capability 
	SurveillanceCapabilityCodeTypeB2 SurveillanceCapabilityCodeType = "B2" // ADS-B with dedicated 1090 MHz ADS-B out and in capability 
	SurveillanceCapabilityCodeTypeC SurveillanceCapabilityCodeType = "C" // Transponder-Mode A (4 digits-4,096 codes) and Mode C 
	SurveillanceCapabilityCodeTypeD1 SurveillanceCapabilityCodeType = "D1" // ADS-C with FANS 1/A capabilities 
	SurveillanceCapabilityCodeTypeE SurveillanceCapabilityCodeType = "E" // Transponder Mode S including aircraft identification, pressure-altitude, and extended squitter capability (ADS-B) 
	SurveillanceCapabilityCodeTypeG1 SurveillanceCapabilityCodeType = "G1" // ADS-C with ATN capabilities 
	SurveillanceCapabilityCodeTypeH SurveillanceCapabilityCodeType = "H" // Transponder Mode S including aircraft identification, pressure-altitude, and enhanced surveillance capability 
	SurveillanceCapabilityCodeTypeI SurveillanceCapabilityCodeType = "I" // Transponder Mode S including aircraft identification, but no pressure-altitude capability 
	SurveillanceCapabilityCodeTypeL SurveillanceCapabilityCodeType = "L" // Transponder Mode S including aircraft identification, pressure-altitude, extended squitter, and enhanced surveillance capability 
	SurveillanceCapabilityCodeTypeP SurveillanceCapabilityCodeType = "P" // Transponder Mode S including pressure-altitude, but no aircraft identification capability
	SurveillanceCapabilityCodeTypeS SurveillanceCapabilityCodeType = "S" // Transponder-Mode S, including both pressure-altitude and aircraft identification transmission
	SurveillanceCapabilityCodeTypeU1 SurveillanceCapabilityCodeType = "U1" // ADS-B out capability using UAT
	SurveillanceCapabilityCodeTypeU2 SurveillanceCapabilityCodeType = "U2" // ADS-B out and in capability using UAT 
	SurveillanceCapabilityCodeTypeV1 SurveillanceCapabilityCodeType = "V1" // ADS-B out capability using VDL mode 4 
	SurveillanceCapabilityCodeTypeV2 SurveillanceCapabilityCodeType = "V2" // ADS-B in and out capability using VDL mode 4
	SurveillanceCapabilityCodeTypeX SurveillanceCapabilityCodeType = "X" // Transponder Mode S with neither aircraft identification nor pressure-altitude capability
)

// Helper simpleType to allow for enumerated lists in FIXM.
type SurveillanceCapabilityCodeListType []SurveillanceCapabilityCodeType

// The type of equipment carried on board the aircraft that can be used by the crew and passengers to assist survival in harsh environments in case of emergency. [FIXM]
type SurvivalEquipmentTypeType string

const (
	SurvivalEquipmentTypeTypeDESERT SurvivalEquipmentTypeType = "DESERT" // Desert survival equipment is carried. [ICAO Doc 4444, Appendix 3]
	SurvivalEquipmentTypeTypeJUNGLE SurvivalEquipmentTypeType = "JUNGLE" // Jungle survival equipment is carried. [ICAO Doc 4444, Appendix 3]
	SurvivalEquipmentTypeTypeMARITIME SurvivalEquipmentTypeType = "MARITIME" // Maritime survival equipment is carried. [ICAO Doc 4444, Appendix 3]
	SurvivalEquipmentTypeTypePOLAR SurvivalEquipmentTypeType = "POLAR" // Polar survival equipment is carried. [ICAO Doc 4444, Appendix 3]
)

// Helper simpleType to allow for enumerated lists in FIXM.
type SurvivalEquipmentTypeListType []SurvivalEquipmentTypeType

// The serviceable communications equipment, available on the aircraft at the time of flight, and associated flight crew qualifications that may be used to communicate with ATS units. [FIXM]
type CommunicationCapabilitiesType struct {
	// A code indicating the capability of the aircraft and associated flight crew qualifications to communicate with ATS units. [FIXM]
	CommunicationCapabilityCode *CommunicationCapabilityCodeListType `xml:"communicationCapabilityCode"`
	// A code indicating the capability of the aircraft to communicate or receive data. [FIXM]
	DatalinkCommunicationCapabilityCode *DatalinkCommunicationCapabilityCodeListType `xml:"datalinkCommunicationCapabilityCode"`
	// An extension hook for attaching extension (non-core) classes.
	Extension []base.CommunicationCapabilitiesExtensionType `xml:"extension"`
	// Additional Communication capabilities available on the aircraft.
	OtherCommunicationCapabilities *base.CharacterStringType `xml:"otherCommunicationCapabilities"`
	// Additional data link capabilities available on the aircraft.
	OtherDatalinkCapabilities *base.CharacterStringType `xml:"otherDatalinkCapabilities"`
	// A code that consists of two 2-letter pairs and acts as a paging system for an ATS unit to establish voice communications with the pilot of an aircraft.
	SelectiveCallingCode *SelectiveCallingCodeType `xml:"selectiveCallingCode"`
}

// Dinghy information including total capacity in persons of all dinghies carried, indicator if dinghies are covered, their colours and number.  [ICAO Doc 4444, Appendix 2, ITEM 19 D]
type DinghiesType struct {
	// The colour of the dinghies carried by the aircraft. [FIXM]
	Colour *base.CharacterStringType `xml:"colour"`
	// Indication of the covered/uncovered nature of the dinghies carried by the aircraft.
	Covered *DinghyCoverIndicatorType `xml:"covered"`
	// An extension hook for attaching extension (non-core) classes.
	Extension []base.DinghiesExtensionType `xml:"extension"`
	// The number of dinghies carried by the aircraft. [adapted from ICAO Doc 4444, Appendix 2, ITEM 19 D/]
	Number *base.CountType `xml:"number"`
	// Total capacity, in persons, of all dinghies carried. [ICAO Doc 4444, Appendix 2, ITEM 19 D/]
	TotalCapacity *base.CountPositiveType `xml:"totalCapacity"`
}

// The capabilities of the flight comprising of the:
		// a) presence of relevant serviceable equipment on board the aircraft;
		// b) equipment and capabilities commensurate with flight crew qualifications; and
		// c) where applicable, authorization from the appropriate authority.
		// [FIXM]
type FlightCapabilitiesType struct {
	// The serviceable communications equipment, available on the aircraft at the time of flight, and associated flight crew qualifications that may be used to communicate with ATS units. [FIXM]
	Communication *CommunicationCapabilitiesType `xml:"communication"`
	// An extension hook for attaching extension (non-core) classes.
	Extension []base.FlightCapabilitiesExtensionType `xml:"extension"`
	// The navigation capability of the aircraft and associated flight crew qualifications. [FIXM]
	Navigation *NavigationCapabilitiesType `xml:"navigation"`
	// If present, indicates that aircraft has the "standard" capabilities for the flight.
	StandardCapabilities *StandardCapabilitiesIndicatorType `xml:"standardCapabilities"`
	// The serviceable Secondary Surveillance Radar (SSR) and/or Automatic Dependent Surveillance (ADS) equipment available on the aircraft at the time of flight that may be used to identify and/or locate the aircraft. [FIXM]
	Surveillance *SurveillanceCapabilitiesType `xml:"surveillance"`
	// The emergency and survival equipments available on board the aircraft. [FIXM]
	Survival *SurvivalCapabilitiesType `xml:"survival"`
}

// The navigation capability of the aircraft and associated flight crew qualifications [FIXM]
type NavigationCapabilitiesType struct {
	// An extension hook for attaching extension (non-core) classes.
	Extension []base.NavigationCapabilitiesExtensionType `xml:"extension"`
	// A code indicating the navigation capability of the aircraft and associated flight crew qualifications. [FIXM]
	NavigationCapabilityCode *NavigationCapabilityCodeListType `xml:"navigationCapabilityCode"`
	// Additional navigation capabilities available on board the aircraft.
	OtherNavigationCapabilities *base.CharacterStringType `xml:"otherNavigationCapabilities"`
	// A coded category denoting which Required Navigation Performance (RNP) and Area Navigation (RNAV) requirements can be met by the aircraft while operating in the context of a particular airspace when supported by the appropriate navigation infrastructure.
	PerformanceBasedCode *PerformanceBasedNavigationCapabilityCodeListType `xml:"performanceBasedCode"`
	// The minimum RVR value required by a flight in order to execute an approach to land at the destination aerodrome in accordance with the applicable ATM configuration. [FF-ICE]
	RequiredRunwayVisualRange *base.DistanceType `xml:"requiredRunwayVisualRange"`
}

// The serviceable Secondary Surveillance Radar (SSR) and/or Automatic Dependent...
type SurveillanceCapabilitiesType struct {
	// An extension hook for attaching extension (non-core) classes.
	Extension []base.SurveillanceCapabilitiesExtensionType `xml:"extension"`
	// Additional surveillance capabilities available on board the aircraft.
	OtherSurveillanceCapabilities *base.CharacterStringType `xml:"otherSurveillanceCapabilities"`
	// A code indicating the SSR and ADS capability of the aircraft. [FIXM]
	SurveillanceCapabilityCode *SurveillanceCapabilityCodeListType `xml:"surveillanceCapabilityCode"`
}

// The emergency and survival equipments available on board the aircraft. [FIXM]
type SurvivalCapabilitiesType struct {
	// The identifier of an emergency locator transmitter carried by aircraft. [FIXM]
	CarriedEltHexIdentifier *EltHexIdentifierListType `xml:"carriedEltHexIdentifier"`
	// Dinghy information including total capacity in persons of all dinghies carried, indicator if dinghies are covered, their colours and number. [ICAO Doc 4444, Appendix 2, ITEM 19 D]
	DinghyInformation *DinghiesType `xml:"dinghyInformation"`
	// The type of serviceable communication devices available on the aircraft that are able to transmit an emergency radio signal. [FIXM]
	EmergencyRadioCapabilityType *EmergencyRadioCapabilityTypeListType `xml:"emergencyRadioCapabilityType"`
	// An extension hook for attaching extension (non-core) classes.
	Extension []base.SurvivalCapabilitiesExtensionType `xml:"extension"`
	// The type of life jackets available on board the aircraft. [FIXM]
	LifeJacketType *LifeJacketTypeListType `xml:"lifeJacketType"`
	// A description of survival equipment carried on the aircraft and any other useful remarks regarding survival equipment.
	SurvivalEquipmentRemarks *base.CharacterStringType `xml:"survivalEquipmentRemarks"`
	// The type of equipment carried on board the aircraft that can be used by the crew and passengers to assist survival in harsh environments in case of emergency. [FIXM]
	SurvivalEquipmentType *SurvivalEquipmentTypeListType `xml:"survivalEquipmentType"`
}

