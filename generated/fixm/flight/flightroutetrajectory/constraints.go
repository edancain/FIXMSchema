// Code generated from Constraints.xsd; DO NOT EDIT.

package flightroutetrajectory

import (
	"github.com/edancain/FIXMSchema.git/generated/fixm/base"
)

// Identifies whether the constraint is applicable on climb or descent. This pro...
type DepartureOrArrivalIndicatorType string

const (
	DepartureOrArrivalIndicatorTypeARRIVAL DepartureOrArrivalIndicatorType = "ARRIVAL"
	DepartureOrArrivalIndicatorTypeDEPARTURE DepartureOrArrivalIndicatorType = "DEPARTURE"
)

// The altitude constraint applicable to a specific point or segment on the route.
type LevelConstraintType struct {
	// Provides an indication of whether the profile constraint is met or initiated ...
	Activation *flight.ActivationType `xml:"activation"`
	// An extension hook for attaching extension (non-core) classes.
	Extension []base.LevelConstraintExtensionType `xml:"extension"`
	// Level or altitude specification of the constraint.
	Level *base.FlightLevelOrAltitudeOrRangeChoiceType `xml:"level"`
}

// A class that defines a constraint to a route point or trajectory.
type RouteTrajectoryConstraintType struct {
	// Identifies whether the constraint is applicable on climb or descent. This pro...
	DepartureOrArrivalIndicator *flight.DepartureOrArrivalIndicatorType `xml:"departureOrArrivalIndicator"`
	// Textual description of the constraint.
	Description *base.CharacterStringType `xml:"description"`
	// An extension hook for attaching extension (non-core) classes.
	Extension []base.RouteTrajectoryConstraintExtensionType `xml:"extension"`
	// The level constraint applicable to a specific point or segment on the route.
	Level *flight.LevelConstraintType `xml:"level"`
	// Reference to a named restriction, if applicable.
	RestrictionReference *base.RestrictionReferenceType `xml:"restrictionReference"`
	// The speed constraint applicable to a specific point or segment on the route.
	Speed *flight.SpeedConstraintType `xml:"speed"`
	// The time constraint applicable to a specific point on the route.
	Time *flight.TimeConstraintType `xml:"time"`
}

// The speed constraint applicable to a specific point or segment on the route.
type SpeedConstraintType struct {
	// Provides an indication of whether the profile constraint is met or initiated ...
	Activation *flight.ActivationType `xml:"activation"`
	// An extension hook for attaching extension (non-core) classes.
	Extension []base.SpeedConstraintExtensionType `xml:"extension"`
	// Speed specification of the speed constraint. Can be either an absolute speed ...
	Speed *base.TrueAirspeedChoiceType `xml:"speed"`
}

// The time constraint applicable to a specific point on the route.
type TimeConstraintType struct {
	// An extension hook for attaching extension (non-core) classes.
	Extension []base.TimeConstraintExtensionType `xml:"extension"`
	// Date and time associated with the constraint.
	TimeSpecification *base.TimeChoiceType `xml:"timeSpecification"`
}

