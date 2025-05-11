// Code generated from RouteChanges.xsd; DO NOT EDIT.

package flightroutetrajectory

import (
	"github.com/edancain/FIXMSchema.git/generated/fixm/base"
)

// Provides an indication of whether the profile constraint is met or initiated at the Location
type ActivationType string

const (
	ActivationTypePLAN_TO_ATTAIN ActivationType = "PLAN_TO_ATTAIN"
	ActivationTypePLAN_TO_COMMENCE ActivationType = "PLAN_TO_COMMENCE"
)

// Indicates the cruise climb will be to an unspecified level at or above the lower level.
type AtOrAboveIndicatorType string

const (
	// Indicates the cruise climb will be to an unspecified level at or above the lower level.
	AtOrAboveIndicatorTypeAT_OR_ABOVE_LOWER_LEVEL AtOrAboveIndicatorType = "AT_OR_ABOVE_LOWER_LEVEL"
)

// Describes the cruise climb parameters at the point at which a cruise climb is planned to commence.  Includes the speed, the lower level and the upper level of the cruise climb.
type CruiseClimbStartType struct {
	// An extension hook for attaching extension (non-core) classes.
	Extension []base.CruiseClimbStartExtensionType `xml:"extension"`
	// Either the lower of the two levels to be occupied during cruise climb, or the level at or above which cruise climb is planned.
	LowerLevel *base.FlightLevelOrAltitudeChoiceType `xml:"lowerLevel"`
	// The speed to be maintained during cruise climb.
	Speed *base.TrueAirspeedType `xml:"speed"`
	// Either the upper of the two levels to be occupied during cruise climb, or an indicator that the cruise climb will be to an unspecified level at or above the lower level.
	UpperLevel *flight.UpperLevelChoiceType `xml:"upperLevel"`
}

// Describes a new planned altitude or level at the specified point at which the change is planned to commence.  (May also describe a point at which the altitude or level change is planned to be attained.)
type CruisingLevelChangeType struct {
	// Qualifies the position associated with the level constraint.
	Activation *flight.ActivationType `xml:"activation"`
	// An extension hook for attaching extension (non-core) classes.
	Extension []base.CruisingLevelChangeExtensionType `xml:"extension"`
	// A new planned altitude or level at the specified point at which the change is planned to commence.
	Level *base.FlightLevelOrAltitudeChoiceType `xml:"level"`
}

// Describes the new planned speed at the specified point at which the change to the speed is planned to commence.
type CruisingSpeedChangeType struct {
	// Qualifies the position associated with the speed constraint.
	Activation *flight.ActivationType `xml:"activation"`
	// An extension hook for attaching extension (non-core) classes.
	Extension []base.CruisingSpeedChangeExtensionType `xml:"extension"`
	// Describes the new planned speed at the specified point at which the change to the new speed is planned to commence
	// .Speed Change: Describes the new planned speed at the specified point at which the change to the new speed is planned to commence.  (May also describe a point at which the speed change is planned to be attained).
	Speed *base.TrueAirspeedType `xml:"speed"`
}

// A requested change that is applicable to a route or trajectory point. Can be any of the several defined requested changes.
type RouteChangeType struct {
	// The parameters of a cruise climb executed at the associated significant point. It contains the following parameters: 1. the speed to be maintained during cruise climb; 2. either the minimum and maximum levels defining the layer to be occupied during cruise climb, or the level above which cruise climb is planned.
	CruiseClimbStart *flight.CruiseClimbStartType `xml:"cruiseClimbStart"`
	// An extension hook for attaching extension (non-core) classes.
	Extension []base.RouteChangeExtensionType `xml:"extension"`
	// The planned altitude that the aircraft will change to upon reaching the associated Significant Point along its Route.
	Level *flight.CruisingLevelChangeType `xml:"level"`
	// The planned speed  that the aircraft will change to upon reaching the associated Significant Point along its Route.
	Speed *flight.CruisingSpeedChangeType `xml:"speed"`
}

// Either the upper of the two levels to be occupied during cruise climb, or an indicator that the cruise climb will be to an unspecified level at or above the lower level.
type UpperLevelChoiceType struct {
	// Altitude specification.
	Altitude base.AltitudeType `xml:"altitude"`
	// Indicates the cruise climb will be to an unspecified level at or above the lower level.
	AtOrAbove flight.AtOrAboveIndicatorType `xml:"atOrAbove"`
	// Flight Level specification.
	FlightLevel base.FlightLevelType `xml:"flightLevel"`
}

