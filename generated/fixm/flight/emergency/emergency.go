// Code generated from Emergency.xsd; DO NOT EDIT.

package emergency

import (
	"github.com/edancain/FIXMSchema.git/generated/fixm/base"
)

// A generic term meaning, as the case may be, uncertainty phase, alert phase or distress phase. [ICAO Annex 11]
// This is the stage of emergency the flight is currently under, as designated by an ATS unit. [FIXM]
type EmergencyPhaseType string

const (
	// The code word used to designate an alert phase. [ICAO Annex 11]
	EmergencyPhaseTypeALERFA EmergencyPhaseType = "ALERFA"
	// The code word used to designate a distress phase. [ICAO Annex 11]
	EmergencyPhaseTypeDETRESFA EmergencyPhaseType = "DETRESFA"
	// The code word used to designate an uncertainty phase. [ICAO Annex 11]
	EmergencyPhaseTypeINCERFA EmergencyPhaseType = "INCERFA"
)

// Groups emergency information (description, phase, position, etc) for the flight.
type FlightEmergencyType struct {
	// A description of the actions taken by the reporting Air Traffic Service (ATS) unit, in the event of search and rescue.
	ActionTaken *base.CharacterStringType `xml:"actionTaken"`
	// A short, plain-language description of the nature of the emergency.
	EmergencyDescription *base.CharacterStringType `xml:"emergencyDescription"`
	// An extension hook for attaching extension (non-core) classes.
	Extension []base.FlightEmergencyExtensionType `xml:"extension"`
	// The last two-way contact between an ATS unit and the aircraft. [FIXM]
	LastContact *flight.LastContactType `xml:"lastContact"`
	// The ICAO identifier of the ATS unit originating the emergency message.
	Originator *base.AtcUnitReferenceType `xml:"originator"`
	// Other pertinent information not captured elsewhere needed to notify appropriate organizations regarding aircraft in need of search and rescue.
	OtherInformation *base.CharacterStringType `xml:"otherInformation"`
	// A generic term meaning, as the case may be, uncertainty phase, alert phase or distress phase. [ICAO Annex 11]
	// This is the stage of emergency the flight is currently under or an indication the emergency has been cancelled, as designated by an ATS unit. [FIXM]
	Phase *flight.EmergencyPhaseType `xml:"phase"`
}

// The last two-way contact between an ATS unit and the aircraft. [FIXM]
type LastContactType struct {
	// An extension hook for attaching extension (non-core) classes.
	Extension []base.LastContactExtensionType `xml:"extension"`
	// The transmitting/receiving frequency of the last two-way contact between the aircraft and an ATS unit.
	LastContactFrequency *base.FrequencyType `xml:"lastContactFrequency"`
	// The time of the last two-way contact between the aircraft and an ATS unit. The time is given in UTC.
	LastContactTime *base.DateTimeUtcType `xml:"lastContactTime"`
	// The last ATS unit which had two-way contact with the aircraft.
	LastContactUnit *base.AtcUnitNameType `xml:"lastContactUnit"`
	// The position of the aircraft last known to ATS and a corresponding timestamp.
	Position *flight.LastPositionReportType `xml:"position"`
}

// The position of the aircraft last known to ATS and a corresponding timestamp.
type LastPositionReportType struct {
	// The vertical distance, at the last known position, of the aircraft considered as a point, measured from mean Sea level (MSL). [FIXM]
	Altitude *base.AltitudeWithSourceType `xml:"altitude"`
	// A plain-language description of the method used to determine the last known position of an aircraft.
	DeterminationMethod *base.CharacterStringType `xml:"determinationMethod"`
	// An extension hook for attaching extension (non-core) classes.
	Extension []base.LastPositionReportExtensionType `xml:"extension"`
	// The speed of the aircraft relative to the surface of the earth at the last known position. [FIXM]
	GroundSpeed *base.GroundSpeedType `xml:"groundSpeed"`
	// The direction, at the last known position, in which the longitudinal axis of the aircraft was pointed. [FIXM]
	Heading *base.BearingType `xml:"heading"`
	// The difference between the measured horizontal coordinates of the aircraft and its true position referenced to the same geodetic datum expressed as a circular error at 95 percent probability. [FIXM]
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
	// Pertinent information needed to notify appropriate organizations regarding loss of radio communication capabilities.
	RadioFailureRemarks *base.CharacterStringType `xml:"radioFailureRemarks"`
	// The remaining communication capability of the aircraft following radio failure.
	RemainingComCapability *base.CharacterStringType `xml:"remainingComCapability"`
}

