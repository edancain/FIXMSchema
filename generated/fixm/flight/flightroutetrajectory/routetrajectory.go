// Code generated from RouteTrajectory.xsd; DO NOT EDIT.
// Modified manually to fix choice types issues.

package flightroutetrajectory

import (
	"github.com/edancain/FIXMSchema.git/generated/fixm/base"
)

// The regulation, or combination of regulations, that governs all aspects of operations under which the pilot plans to fly.
type FlightRulesType string

const (
	// Instrument Flight Rules
	FlightRulesTypeIFR FlightRulesType = "IFR"
	// Visual Flight Rules
	FlightRulesTypeVFR FlightRulesType = "VFR"
)

// Indication that the route item was modified
type ModifiedRouteItemIndicatorType string

const (
	ModifiedRouteItemIndicatorTypeMODIFIED_ROUTE_ITEM ModifiedRouteItemIndicatorType = "MODIFIED_ROUTE_ITEM"
)

// Indicates that the route to next element is either DIRECT or UNSPECIFIED.
type OtherRouteDesignatorType string

const (
	// When the element is direct, "Direct" must be specified. [FF-ICE]
	OtherRouteDesignatorTypeDIRECT OtherRouteDesignatorType = "DIRECT"
	// If the element is a delay segment, the route to next element should be labeled as "Unspecified". [FF-ICE]
	OtherRouteDesignatorTypeUNSPECIFIED OtherRouteDesignatorType = "UNSPECIFIED"
)

// Indicates if the delay is planned airborne holding or Operator (Airspace User) requested operations at a specified location.
type PlannedDelayTypeType string

const (
	// The delay is an ASP request for airborne holding. The holding pattern name is in the delayReference field. The delay value used to compute estimates as in OPERATOR_REQUEST_POINT.
	PlannedDelayTypeTypeATFM PlannedDelayTypeType = "ATFM"
	// The delay is an Operator (Airspace User) request to 'spend time' at an aerodrome after the RoutePoint, the name of the aerodrome is in the delayReference field, the points at which the Operator leaves its route and rejoin it go to the aerodrome are in the BEGIN_STAY and END_STAY info of the trajectory
	PlannedDelayTypeTypeOPERATOR_REQUEST_AERODROME PlannedDelayTypeType = "OPERATOR_REQUEST_AERODROME"
	// The delay is an Operator (Airspace User) request to 'spend time' in an airspace after the RoutePoint the name of which is in delayReference, the delayValue is the time in that airspace, the entry and exit time into that airspace is in the BEGIN_STAY and END_STAY info of the trajectory.
	PlannedDelayTypeTypeOPERATOR_REQUEST_AIRSPACE PlannedDelayTypeType = "OPERATOR_REQUEST_AIRSPACE"
	// The delay is an Operator (Airspace User) request to 'spend time' at a holding pattern the anchor point of which is the RoutePoint.  The holding pattern name if any is in delayReference.  delayValue used to compute estimates as in OPERATOR_REQUEST_POINT
	PlannedDelayTypeTypeOPERATOR_REQUEST_HOLDING PlannedDelayTypeType = "OPERATOR_REQUEST_HOLDING"
	// The delay is an Operator (Airspace User) request to 'spend time' at the point, thus the delay has to be added to the flight duration to the next point to compute the estimate to the next point.
	PlannedDelayTypeTypeOPERATOR_REQUEST_POINT PlannedDelayTypeType = "OPERATOR_REQUEST_POINT"
	// The delay is an Operator (Airspace User) request to 'spend time' at the segment starting at point, thus the delay has to be understood as the total duration between the point and the next one.
	PlannedDelayTypeTypeOPERATOR_REQUEST_SEGMENT PlannedDelayTypeType = "OPERATOR_REQUEST_SEGMENT"
)

// Indicates the route description is truncated at the preceding point
type RouteTruncationIndicatorType string

const (
	RouteTruncationIndicatorTypeROUTE_TRUNCATION RouteTruncationIndicatorType = "ROUTE_TRUNCATION"
)

// Enumerates any applicable properties of the trajectory point.
type TrajectoryPointPropertyTypeType string

const (
	// Indicates that the associated trajectory point is the airport reference location.
	TrajectoryPointPropertyTypeTypeAIRPORT_REFERENCE_LOCATION TrajectoryPointPropertyTypeType = "AIRPORT_REFERENCE_LOCATION"
	// And many more constants...
)

// The location associated with estimated elapsed time. It can be a longitude, significant point or flight information region.
type ElapsedTimeLocationChoiceType struct {
	// Longitude associated with the elapsed time.
	Longitude *base.LongitudeType `xml:"longitude,omitempty"`
	// Point associated with the estimated elapsed time.
	Point *base.SignificantPointChoiceType `xml:"point,omitempty"`
	// Flight information boundary associated with the elapsed time.
	Region *base.LocationIndicatorType `xml:"region,omitempty"`
}

// The estimated amount of time from takeoff to reach a significant point or Flight Information Region (FIR) boundary along the route of flight.
type EstimatedElapsedTimeType struct {
	// The estimated amount of elapsed time.
	ElapsedTime *base.DurationType `xml:"elapsedTime,omitempty"`
	// An extension hook for attaching extension (non-core) classes.
	Extension []base.EstimatedElapsedTimeExtensionType `xml:"extension,omitempty"`
	// The location associated with estimated elapsed time. It can be a longitude, significant point or flight information region.
	Location *ElapsedTimeLocationChoiceType `xml:"location,omitempty"`
	// Incrementing identifier used to ensure the sequence of ordered lists is retained. The identifier should begin at zero. [FIXM]
	SeqNum int `xml:"seqNum,attr,omitempty"`
}

// Contains basic information about the Flight Route that pertains to the whole flight.
type FlightRouteInformationType struct {
	// The internationally agreed effective date of the AIRAC (Aeronautical Information Regulation and Control) data set that was used in the computation of the route/trajectory.  [FIXM]
	AiracReference *base.DateUtcType `xml:"airacReference,omitempty"`
	// The filed altitude (flight level) for the first or the whole cruising portion of the flight.
	CruisingLevel *base.FlightLevelOrAltitudeOrVfrChoiceType `xml:"cruisingLevel,omitempty"`
	// The true airspeed for the first or the whole cruising portion of the flight.  This can be for a filed flight or an active flight.  This element is strategic in nature.
	CruisingSpeed *base.TrueAirspeedType `xml:"cruisingSpeed,omitempty"`
	// The estimated amount of time from takeoff to reach a significant point or Flight Information Region (FIR) boundary along the route of flight. Elements are ordered according to increasing duration from start. 
	EstimatedElapsedTime []EstimatedElapsedTimeType `xml:"estimatedElapsedTime,omitempty"`
	// An extension hook for attaching extension (non-core) classes.
	Extension []base.FlightRouteInformationExtensionType `xml:"extension,omitempty"`
	// A string of route elements complying with PANS-ATM Item 15c
	RouteText *base.CharacterStringType `xml:"routeText,omitempty"`
	// For IFR flights, the estimated time required from take-off to arrive over that designated point, defined by reference to navigation aids, from which it is intended that an instrument approach procedure will be commenced, or, if no navigation aid is associated with the destination aerodrome, to arrive over the destination aerodrome. For VFR flights, the estimated time required from take-off to arrive over the destination aerodrome. [ICAO Doc 4444]
	// For a flight plan received from an aircraft in flight, the total estimated elapsed time is the estimated time from the first point of the route to which the flight plan applies to the termination point of the flight plan. [ICAO Doc 4444, Appendix 2, ITEM 16]
	TotalEstimatedElapsedTime *base.DurationType `xml:"totalEstimatedElapsedTime,omitempty"`
}

// In a predicted trajectory, the instantaneous temperature and wind vector used at the 4D Point for creating the 4D trajectory.
type MeteorologicalDataType struct {
	// An extension hook for attaching extension (non-core) classes.
	Extension []base.MeteorologicalDataExtensionType `xml:"extension,omitempty"`
	// In a predicted trajectory, the instantaneous temperature used at the 4D Point for creating the 4D trajectory.
	Temperature *base.TemperatureType `xml:"temperature,omitempty"`
	// In a predicted trajectory, the instantaneous wind direction used at the 4D Point for creating the 4D trajectory.
	WindDirection *base.WindDirectionType `xml:"windDirection,omitempty"`
	// In a predicted trajectory, the instantaneous wind speed used at the 4D Point for creating the 4D trajectory.
	WindSpeed *base.WindSpeedType `xml:"windSpeed,omitempty"`
}

// A zero-wind, standard atmosphere, unconstrained movement profile reflective of the flight capabilities and desired parameters during climb or descent.
type PerformanceProfileType struct {
	// An extension hook for attaching extension (non-core) classes.
	Extension []base.PerformanceProfileExtensionType `xml:"extension,omitempty"`
	// A point in a performance climb or descent profile.  Order is based on time from earliest to latest.
	ProfilePoint []ProfilePointType `xml:"profilePoint,omitempty"`
}

// Delay or holding planned to occur at a significant point or along a route element.
type PlannedDelayType struct {
	// The reason for the delay. [FIXM]
	DelayReason *base.CharacterStringType `xml:"delayReason,omitempty"`
	// Indicates a named hold pattern, airspace, or aerodrome at which the delay is expected to occur.
	DelayReference *base.CharacterStringType `xml:"delayReference,omitempty"`
	// Indicates if the delay is planned airborne holding or Operator (Airspace User) requested operations at a specified location.
	DelayType *PlannedDelayTypeType `xml:"delayType,omitempty"`
	// The length of time the flight is expected to be delayed at a specific point.
	DelayValue *base.DurationType `xml:"delayValue,omitempty"`
	// An extension hook for attaching extension (non-core) classes.
	Extension []base.PlannedDelayExtensionType `xml:"extension,omitempty"`
}

// Allows to either specify a 4D point time as either an absolute time or relative time.
type Point4DTimeChoiceType struct {
	// Absolute Time of the 4D point.
	AbsoluteTime *base.DateTimeUtcType `xml:"absoluteTime,omitempty"`
	// Relative Time of the 4D point (expressed as duration from takeoff).
	RelativeTimeFromInitialPredictionPoint base.DurationType `xml:"relativeTimeFromInitialPredictionPoint,omitempty"`
}

// A point in a performance climb or descent profile.
type ProfilePointType struct {
	// The true airspeed of a point in a zero-wind, standard atmosphere, unconstrained profile.
	Airspeed *base.TrueAirspeedType `xml:"airspeed,omitempty"`
	// The distance (from the start of the profile) at a point in a zero-wind, unconstrained profile.
	Distance *base.DistanceType `xml:"distance,omitempty"`
	// An extension hook for attaching extension (non-core) classes.
	Extension []base.ProfilePointExtensionType `xml:"extension,omitempty"`
	// The altitude of a point in a zero-wind, unconstrained profile.
	Level *base.FlightLevelOrAltitudeChoiceType `xml:"level,omitempty"`
	// The time (from the start of the profile) at a point in a zero-wind, unconstrained profile.
	Time *base.DurationType `xml:"time,omitempty"`
	// Incrementing identifier used to ensure the sequence of ordered lists is retained. The identifier should begin at zero. [FIXM]
	SeqNum int `xml:"seqNum,attr,omitempty"`
}

// The route (airway) to the next element on the route. Can be either an en route airway, standard instrument departure route, or standard instrument arrival route. [FIXM]
type RouteDesignatorToNextElementChoiceType struct {
	// Indicates that the route to next element is either DIRECT or UNSPECIFIED.
	OtherRouteDesignator *OtherRouteDesignatorType `xml:"otherRouteDesignator,omitempty"`
	// The coded designator for a published ATS route or route segment.
	RouteDesignator *base.RouteDesignatorType `xml:"routeDesignator,omitempty"`
	// The textual designator of the Standard Instrument Arrival (STAR).
	StandardInstrumentArrival *base.SidStarReferenceType `xml:"standardInstrumentArrival,omitempty"`
	// This is the name of a published procedure extending from the departure runway to the start of the en route part of the aircraft's route.
	StandardInstrumentDeparture *base.SidStarReferenceType `xml:"standardInstrumentDeparture,omitempty"`
}

// A portion of the route that can be at one of three levels as described below:
// a) A Route Element defined by a Significant Point and the ATS Route (or direct route) to be followed until a change in route.
// b) An Expanded Route Element defined by a Significant Point and the ATS Route (or direct route) to be followed until the next Significant Point, which may be along the same ATS Route.
// c)A Trajectory Element defined by a geographic point and the ATS Route (or direct route) followed until the next trajectory point. [FF-ICE/1 Implementation Guidance Manual, Appendix B, Chapter B-3.2]
type RouteTrajectoryElementType struct {
	// The distance along the route.  For an eASP-provided expanded route, it is computed from the first converted point in the eASP's airspace for each route point in the expanded route. For an operator-provided expanded route, it is computed from the beginning of the route.
	AlongRouteDistance *base.DistanceType `xml:"alongRouteDistance,omitempty"`
	// A number of constraints can be associated with each Expanded Route Point and Trajectory point.
	Constraint []RouteTrajectoryConstraintType `xml:"constraint,omitempty"`
	// A specified geographical location used in defining the flight route or expanded route. [FIXM]
	ElementStartPoint *base.SignificantPointChoiceType `xml:"elementStartPoint,omitempty"`
	// An extension hook for attaching extension (non-core) classes.
	Extension []base.RouteTrajectoryElementExtensionType `xml:"extension,omitempty"`
	// Describes the planned change of flight rules (e.g., IFR/VFR) associated with a route point.
	FlightRulesChange *FlightRulesType `xml:"flightRulesChange,omitempty"`
	// Identifies if the route point or element was modified from the reference version, identified in the message, during negotiation by an ASP. 
	Modified *ModifiedRouteItemIndicatorType `xml:"modified,omitempty"`
	// A reference to an ATFM program name that modified the route.
	ModifiedRouteItemReference *base.CharacterStringType `xml:"modifiedRouteItemReference,omitempty"`
	// Delay or holding planned to occur at a significant point or along a route element.
	PlannedDelay *PlannedDelayType `xml:"plannedDelay,omitempty"`
	// Identifies the location, altitude and time of a trajectory point.
	Point4D *TrajectoryPoint4DType `xml:"point4D,omitempty"`
	// Flight Route element may contain a list of route changes.
	RouteChange *RouteChangeType `xml:"routeChange,omitempty"`
	// The route (airway) to the next element on the route. Can be either an en route airway, standard instrument departure route, or standard instrument arrival route. [FIXM]
	RouteDesignatorToNextElement *RouteDesignatorToNextElementChoiceType `xml:"routeDesignatorToNextElement,omitempty"`
	// Indicates the route description is truncated at the preceding point.
	RouteTruncationIndicator *RouteTruncationIndicatorType `xml:"routeTruncationIndicator,omitempty"`
	//Incrementing identifier used to ensure the sequence of ordered lists is retained. The identifier should begin at zero. [FIXM]</xs:documentation>
    SeqNum int `xml:"seqNum,attr,omitempty"`
}

// A data container supporting the representation of aircraft movement descriptions with variable levels of granularity.  This container enables the representation of a flight route and a 4D trajectory, as described in ICAO Doc 9965 Volume 2.  [FIXM]
type RouteTrajectoryGroupType struct {
	// A zero-wind, standard atmosphere, unconstrained climb profile reflective of the flight capabilities and desired parameters during climb.
	ClimbProfile *PerformanceProfileType `xml:"climbProfile,omitempty"`
	// Initially submitted by the airspace user, this defines the target speed in both Indicated Airspeed (IAS) and MACH so the aircraft can climb through the crossover altitude and target the MACH speed when needed.
	ClimbSchedule *SpeedScheduleType `xml:"climbSchedule,omitempty"`
	// A zero-wind, standard atmosphere, unconstrained descent profile reflective of the flight capabilities and desired parameters during descent.
	DescentProfile *PerformanceProfileType `xml:"descentProfile,omitempty"`
	// Initially submitted by the airspace user, this defines the target speed in both (Indicated Airspeed) IAS and MACH so the aircraft can descend through the crossover altitude and target the IAS speed when needed.
	DescentSchedule *SpeedScheduleType `xml:"descentSchedule,omitempty"`
	// A container for information pertinent to a single point in a trajectory. The elements are ordered according to increasing time/distance from the start point.
	Element []RouteTrajectoryElementType `xml:"element,omitempty"`
	// An extension hook for attaching extension (non-core) classes.
	Extension []base.RouteTrajectoryGroupExtensionType `xml:"extension,omitempty"`
	// Contains information about the Flight Route. The route is comprised of Route Segments, which are comprised of Route Point/Airway pairs. Route also contains an expanded route.
	RouteInformation *FlightRouteInformationType `xml:"routeInformation,omitempty"`
	// The estimated takeoff mass of the aircraft.
	TakeoffMass *base.MassType `xml:"takeoffMass,omitempty"`
}

// Defines the speed schedule for climb and descent of the aircraft through the crossover altitude.
type SpeedScheduleType struct {
	// An extension hook for attaching extension (non-core) classes.
	Extension []base.SpeedScheduleExtensionType `xml:"extension,omitempty"`
	// Initial speed of the aircraft during the climb
	InitialSpeed *base.IndicatedAirspeedType `xml:"initialSpeed,omitempty"`
	// Subsequent speed of the aircraft during the climb
	SubsequentSpeed *base.IndicatedAirspeedType `xml:"subsequentSpeed,omitempty"`
}

// Identifies relevant information about a trajectory point including location, altitude, time, etc.
type TrajectoryPoint4DType struct {
	// The barometric pressure reading used to adjust a pressure altimeter for variations in existing atmospheric pressure or to the standard altimeter setting (29.92).
	AltimeterSetting *base.PressureType `xml:"altimeterSetting,omitempty"`
	// An extension hook for attaching extension (non-core) classes.
	Extension []base.TrajectoryPoint4DExtensionType `xml:"extension,omitempty"`
	// Altitude of the 4D point.
	Level *base.FlightLevelOrAltitudeChoiceType `xml:"level,omitempty"`
	// In a predicted trajectory, the instantaneous temperature and wind vector used at the 4D Point for creating the 4D trajectory.
	MetData *MeteorologicalDataType `xml:"metData,omitempty"`
	// Describes any applicable properties of the trajectory point.  May include multiple properties per point.
	PointProperty []TrajectoryPointPropertyType `xml:"pointProperty,omitempty"`
	// The geographical location of the 4D point.
	Position *base.GeographicalPositionType `xml:"position,omitempty"`
	// The predicted indicated airspeed or Mach at the trajectory point.
	PredictedAirspeed *base.IndicatedAirspeedType `xml:"predictedAirspeed,omitempty"`
	// The predicted ground speed at the Trajectory Point in knots or kilometers per hour.
	PredictedGroundspeed *base.GroundSpeedType `xml:"predictedGroundspeed,omitempty"`
	// 4D Point time expressed as either absolute or relative time.
	Time *Point4DTimeChoiceType `xml:"time,omitempty"`
	// Vertical Range representing block altitude clearances
	VerticalRange *base.VerticalRangeType `xml:"verticalRange,omitempty"`
}

// Describes any applicable properties of the trajectory point.  May include multiple properties per point.
type TrajectoryPointPropertyType struct {
	// Each Trajectory Point Property shall have the ability to include a textual description. [FF-ICE]
	Description *base.CharacterStringType `xml:"description,omitempty"`
	// An extension hook for attaching extension (non-core) classes.
	Extension []base.TrajectoryPointPropertyExtensionType `xml:"extension,omitempty"`
	// Provides the type of trajectory point property.
	PropertyType *TrajectoryPointPropertyTypeType `xml:"propertyType,omitempty"`
	// A reference providing an identifier of airspace or program affecting the point.
	Reference *TrajectoryPointReferenceType `xml:"reference,omitempty"`
}

// A reference providing an identifier of airspace or program affecting the point.
type TrajectoryPointReferenceType struct {
	// An extension hook for attaching extension (non-core) classes.
	Extension []base.TrajectoryPointReferenceExtensionType `xml:"extension,omitempty"`
	// The name of the particular airspace or program (e.g., FCA016, CTP001, RRDCC123)
	Identifier *base.CharacterStringType `xml:"identifier,omitempty"`
	// The nature of the airspace or program (e.g., special activity airspace, GDP, MIT, etc.)
	Type *base.CharacterStringType `xml:"type,omitempty"`
	// Provides an optional mechanism enabling FIXM aeronautical fields to be supplemented with references to AIXM features.
	Href string `xml:"href,attr,omitempty"`
}

// Helper methods for Point4DTimeChoiceType
func (p *Point4DTimeChoiceType) IsAbsoluteTimeSet() bool {
	return p != nil && p.AbsoluteTime != nil && !p.AbsoluteTime.IsZero()
}

func (p *Point4DTimeChoiceType) IsRelativeTimeSet() bool {
   return p != nil && p.RelativeTimeFromInitialPredictionPoint != 0
}

// Helper methods for RouteDesignatorToNextElementChoiceType
func (r *RouteDesignatorToNextElementChoiceType) IsOtherRouteDesignatorSet() bool {
   return r != nil && r.OtherRouteDesignator != nil && string(*r.OtherRouteDesignator) != ""
}

func (r *RouteDesignatorToNextElementChoiceType) IsRouteDesignatorSet() bool {
   // Since RouteDesignator is now a struct with Value field
   return r != nil && r.RouteDesignator != nil && r.RouteDesignator.Value != ""
}

func (r *RouteDesignatorToNextElementChoiceType) IsStandardInstrumentArrivalSet() bool {
   return r != nil && r.StandardInstrumentArrival != nil && r.StandardInstrumentArrival.Designator != ""
}

func (r *RouteDesignatorToNextElementChoiceType) IsStandardInstrumentDepartureSet() bool {
   return r != nil && r.StandardInstrumentDeparture != nil && r.StandardInstrumentDeparture.Designator != ""
}

// Helper methods for ElapsedTimeLocationChoiceType
func (e *ElapsedTimeLocationChoiceType) IsLongitudeSet() bool {
   return e != nil && e.Longitude != nil
}

func (e *ElapsedTimeLocationChoiceType) IsPointSet() bool {
   return e != nil && e.Point != nil
}

func (e *ElapsedTimeLocationChoiceType) IsRegionSet() bool {
   return e != nil && e.Region != nil
}