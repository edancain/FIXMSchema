package flight/emergency

import (
	"github.com/yourusername/fixm/generated/fixm/base"
)

// EmergencyPhaseType is The code word used to designate an uncertainty phase. [ICAO Annex 11]
type EmergencyPhaseType string

// FlightEmergencyType is An extension hook for attaching extension (non-core) classes.
type FlightEmergencyType struct {
	ActionTaken	*base.CharacterStringType	`xml:"actionTaken"`
	EmergencyDescription	*base.CharacterStringType	`xml:"emergencyDescription"`
	Extension	[]*FlightEmergencyExtensionType	`xml:"extension"`
	LastContact	interface{}	`xml:"lastContact"`
	Originator	*AtcUnitReferenceType	`xml:"originator"`
	OtherInformation	*base.CharacterStringType	`xml:"otherInformation"`
	Phase	string	`xml:"phase"`
}

// LastContactType ...
type LastContactType struct {
	Extension	[]*LastContactExtensionType	`xml:"extension"`
	LastContactFrequency	*FrequencyType	`xml:"lastContactFrequency"`
	LastContactTime	*DateTimeUtcType	`xml:"lastContactTime"`
	LastContactUnit	*AtcUnitNameType	`xml:"lastContactUnit"`
	Position	interface{}	`xml:"position"`
}

// LastPositionReportType ...
type LastPositionReportType struct {
	Altitude	*AltitudeWithSourceType	`xml:"altitude"`
	DeterminationMethod	*base.CharacterStringType	`xml:"determinationMethod"`
	Extension	[]*LastPositionReportExtensionType	`xml:"extension"`
	GroundSpeed	*GroundSpeedType	`xml:"groundSpeed"`
	Heading	*BearingType	`xml:"heading"`
	HorizontalAccuracy	*DistanceType	`xml:"horizontalAccuracy"`
	Position	*SignificantPointChoiceType	`xml:"position"`
	TimeAtPosition	*DateTimeUtcType	`xml:"timeAtPosition"`
}

// RadioCommunicationFailureType ...
type RadioCommunicationFailureType struct {
	Contact	interface{}	`xml:"contact"`
	Extension	[]*RadioCommunicationFailureExtensionType	`xml:"extension"`
	RadioFailureRemarks	*base.CharacterStringType	`xml:"radioFailureRemarks"`
	RemainingComCapability	*base.CharacterStringType	`xml:"remainingComCapability"`
}
