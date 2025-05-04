// Code generated from Emergency.xsd; DO NOT EDIT.

package emergency

import (
	"github.com/edancain/FIXMSchema.git/generated/fixm/base"
)

// A generic term meaning, as the case may be, uncertainty phase, alert phase or...
type EmergencyPhaseType string

const (
	EmergencyPhaseTypeALERFA EmergencyPhaseType = "ALERFA"
	EmergencyPhaseTypeDETRESFA EmergencyPhaseType = "DETRESFA"
	EmergencyPhaseTypeINCERFA EmergencyPhaseType = "INCERFA"
)

// Groups emergency information (description, phase, position, etc) for the flight.
type FlightEmergencyType struct {
	// A description of the actions taken by the reporting Air Traffic Service (ATS)...
	ActionTaken *base.CharacterStringType `xml:"actionTaken"`
	// A short, plain-language description of the nature of the emergency.
	EmergencyDescription *base.CharacterStringType `xml:"emergencyDescription"`
	// An extension hook for attaching extension (non-core) classes.
	Extension []base.FlightEmergencyExtensionType `xml:"extension"`
	// The last two-way contact between an ATS unit and the aircraft. [FIXM]
	LastContact *flight.LastContactType `xml:"lastContact"`
	// The ICAO identifier of the ATS unit originating the emergency message.
	Originator *base.AtcUnitReferenceType `xml:"originator"`
	// Other pertinent information not captured elsewhere needed to notify appropria...
	OtherInformation *base.CharacterStringType `xml:"otherInformation"`
	// A generic term meaning, as the case may be, uncertainty phase, alert phase or...
	Phase *flight.EmergencyPhaseType `xml:"phase"`
}

// The last two-way contact between an ATS unit and the aircraft. [FIXM]
type LastContactType struct {
	// An extension hook for attaching extension (non-core) classes.
	Extension []base.LastContactExtensionType `xml:"extension"`
	// The transmitting/receiving frequency of the last two-way contact between the ...
	LastContactFrequency *base.FrequencyType `xml:"lastContactFrequency"`
	// The time of the last two-way contact between the aircraft and an ATS unit. Th...
	LastContactTime *base.DateTimeUtcType `xml:"lastContactTime"`
	// The last ATS unit which had two-way contact with the aircraft.
	LastContactUnit *base.AtcUnitNameType `xml:"lastContactUnit"`
	// The position of the aircraft last known to ATS and a corresponding timestamp.
	Position *flight.LastPositionReportType `xml:"position"`
}

// The position of the aircraft last known to ATS and a corresponding timestamp.
type LastPositionReportType struct {
	// The vertical distance, at the last known position, of the aircraft considered...
	Altitude *base.AltitudeWithSourceType `xml:"altitude"`
	// A plain-language description of the method used to determine the last known p...
	DeterminationMethod *base.CharacterStringType `xml:"determinationMethod"`
	// An extension hook for attaching extension (non-core) classes.
	Extension []base.LastPositionReportExtensionType `xml:"extension"`
	// The speed of the aircraft relative to the surface of the earth at the last kn...
	GroundSpeed *base.GroundSpeedType `xml:"groundSpeed"`
	// The direction, at the last known position, in which the longitudinal axis of ...
	Heading *base.BearingType `xml:"heading"`
	// The difference between the measured horizontal coordinates of the aircraft an...
	HorizontalAccuracy *base.DistanceType `xml:"horizontalAccuracy"`
	// The position of the aircraft last known to ATS.
	Position *base.SignificantPointChoiceType `xml:"position"`
	// Timestamp corresponding to the last known position.
	TimeAtPosition *base.DateTimeUtcType `xml:"timeAtPosition"`
}

// Groups information regarding loss of radio communication capabilities.
type RadioCommunicationFailureType struct {
	// The last ATS unit which had two-way contact with the aircraft.
	Contact *flight.LastContactType `xml:"contact"`
	// An extension hook for attaching extension (non-core) classes.
	Extension []base.RadioCommunicationFailureExtensionType `xml:"extension"`
	// Pertinent information needed to notify appropriate organizations regarding lo...
	RadioFailureRemarks *base.CharacterStringType `xml:"radioFailureRemarks"`
	// The remaining communication capability of the aircraft following radio failure.
	RemainingComCapability *base.CharacterStringType `xml:"remainingComCapability"`
}

