// Code generated by xgen. DO NOT EDIT.

package emergency
import (
	"github.com/edancain/FIXMSchema.git/generated/fixm/base"
)

// EmergencyPhaseType is The code word used to designate an uncertainty phase. [ICAO Annex 11]
type EmergencyPhaseType string

// FlightEmergencyType is An extension hook for attaching extension (non-core) classes.
type FlightEmergencyType struct {
	ActionTaken          *base.CharacterStringType            `xml:"actionTaken"`
	EmergencyDescription *base.CharacterStringType            `xml:"emergencyDescription"`
	Extension            []*base.FlightEmergencyExtensionType `xml:"extension"`
	LastContact          interface{}                     `xml:"lastContact"`
	Originator           *base.AtcUnitReferenceType           `xml:"originator"`
	OtherInformation     *base.CharacterStringType            `xml:"otherInformation"`
	Phase                string                          `xml:"phase"`
}

// LastContactType ...
type LastContactType struct {
	Extension            []*base.LastContactExtensionType `xml:"extension"`
	LastContactFrequency *base.FrequencyType              `xml:"lastContactFrequency"`
	LastContactTime      *base.DateTimeUtcType            `xml:"lastContactTime"`
	LastContactUnit      *base.AtcUnitNameType            `xml:"lastContactUnit"`
	Position             interface{}                 `xml:"position"`
}

// LastPositionReportType ...
type LastPositionReportType struct {
	Altitude            *base.AltitudeWithSourceType            `xml:"altitude"`
	DeterminationMethod *base.CharacterStringType               `xml:"determinationMethod"`
	Extension           []*base.LastPositionReportExtensionType `xml:"extension"`
	GroundSpeed         *base.GroundSpeedType                   `xml:"groundSpeed"`
	Heading             *base.BearingType                       `xml:"heading"`
	HorizontalAccuracy  *base.DistanceType                      `xml:"horizontalAccuracy"`
	Position            *base.SignificantPointChoiceType        `xml:"position"`
	TimeAtPosition      *base.DateTimeUtcType                   `xml:"timeAtPosition"`
}

// RadioCommunicationFailureType ...
type RadioCommunicationFailureType struct {
	Contact                interface{}                               `xml:"contact"`
	Extension              []*base.RadioCommunicationFailureExtensionType `xml:"extension"`
	RadioFailureRemarks    *base.CharacterStringType                      `xml:"radioFailureRemarks"`
	RemainingComCapability *base.CharacterStringType                      `xml:"remainingComCapability"`
}
