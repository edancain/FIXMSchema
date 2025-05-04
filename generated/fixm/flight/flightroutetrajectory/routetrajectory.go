// Code generated from RouteTrajectory.xsd; DO NOT EDIT.

package flightroutetrajectory

import (
	"github.com/edancain/FIXMSchema.git/generated/fixm/base"
)

// The regulation, or combination of regulations, that governs all aspects of op...
type FlightRulesType string

const (
	FlightRulesTypeIFR FlightRulesType = "IFR"
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
	OtherRouteDesignatorTypeDIRECT OtherRouteDesignatorType = "DIRECT"
	OtherRouteDesignatorTypeUNSPECIFIED OtherRouteDesignatorType = "UNSPECIFIED"
)

// Indicates if the delay is planned airborne holding or Operator (Airspace User...
type PlannedDelayTypeType string

const (
	PlannedDelayTypeTypeATFM PlannedDelayTypeType = "ATFM"
	PlannedDelayTypeTypeOPERATOR_REQUEST_AERODROME PlannedDelayTypeType = "OPERATOR_REQUEST_AERODROME"
	PlannedDelayTypeTypeOPERATOR_REQUEST_AIRSPACE PlannedDelayTypeType = "OPERATOR_REQUEST_AIRSPACE"
	PlannedDelayTypeTypeOPERATOR_REQUEST_HOLDING PlannedDelayTypeType = "OPERATOR_REQUEST_HOLDING"
	PlannedDelayTypeTypeOPERATOR_REQUEST_POINT PlannedDelayTypeType = "OPERATOR_REQUEST_POINT"
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
	TrajectoryPointPropertyTypeTypeAIRPORT_REFERENCE_LOCATION TrajectoryPointPropertyTypeType = "AIRPORT_REFERENCE_LOCATION"
	TrajectoryPointPropertyTypeTypeBEGIN_STAY TrajectoryPointPropertyTypeType = "BEGIN_STAY"
	TrajectoryPointPropertyTypeTypeCONSTRAINT_POINT TrajectoryPointPropertyTypeType = "CONSTRAINT_POINT"
	TrajectoryPointPropertyTypeTypeCROSSING_CONSTRAINED_AIRSPACE TrajectoryPointPropertyTypeType = "CROSSING_CONSTRAINED_AIRSPACE"
	TrajectoryPointPropertyTypeTypeCROSSOVER_ALTITUDE TrajectoryPointPropertyTypeType = "CROSSOVER_ALTITUDE"
	TrajectoryPointPropertyTypeTypeDEPARTURE_RUNWAY_END TrajectoryPointPropertyTypeType = "DEPARTURE_RUNWAY_END"
	TrajectoryPointPropertyTypeTypeEND_EXPECT_VECTORS TrajectoryPointPropertyTypeType = "END_EXPECT_VECTORS"
	TrajectoryPointPropertyTypeTypeEND_LANDING_ROLL TrajectoryPointPropertyTypeType = "END_LANDING_ROLL"
	TrajectoryPointPropertyTypeTypeEND_PREDICTION_POINT TrajectoryPointPropertyTypeType = "END_PREDICTION_POINT"
	TrajectoryPointPropertyTypeTypeEND_STAY TrajectoryPointPropertyTypeType = "END_STAY"
	TrajectoryPointPropertyTypeTypeENTRY_CONSTRAINED_AIRSPACE TrajectoryPointPropertyTypeType = "ENTRY_CONSTRAINED_AIRSPACE"
	TrajectoryPointPropertyTypeTypeENTRY_RESTRICTED_OR_RESERVED_AIRSPACE TrajectoryPointPropertyTypeType = "ENTRY_RESTRICTED_OR_RESERVED_AIRSPACE"
	TrajectoryPointPropertyTypeTypeEXIT_CONSTRAINED_AIRSPACE TrajectoryPointPropertyTypeType = "EXIT_CONSTRAINED_AIRSPACE"
	TrajectoryPointPropertyTypeTypeEXIT_RESTRICTED_OR_RESERVED_AIRSPACE TrajectoryPointPropertyTypeType = "EXIT_RESTRICTED_OR_RESERVED_AIRSPACE"
	TrajectoryPointPropertyTypeTypeFIR_BOUNDARY_CROSSING_POINT TrajectoryPointPropertyTypeType = "FIR_BOUNDARY_CROSSING_POINT"
	TrajectoryPointPropertyTypeTypeHOLD_ENTRY TrajectoryPointPropertyTypeType = "HOLD_ENTRY"
	TrajectoryPointPropertyTypeTypeHOLD_EXIT TrajectoryPointPropertyTypeType = "HOLD_EXIT"
	TrajectoryPointPropertyTypeTypeINITIAL_PREDICTION_POINT TrajectoryPointPropertyTypeType = "INITIAL_PREDICTION_POINT"
	TrajectoryPointPropertyTypeTypeIN_BLOCKS TrajectoryPointPropertyTypeType = "IN_BLOCKS"
	TrajectoryPointPropertyTypeTypeOFF_BLOCKS TrajectoryPointPropertyTypeType = "OFF_BLOCKS"
	TrajectoryPointPropertyTypeTypePRESCRIBED_EET_POINT TrajectoryPointPropertyTypeType = "PRESCRIBED_EET_POINT"
	TrajectoryPointPropertyTypeTypeRUNWAY_THRESHOLD TrajectoryPointPropertyTypeType = "RUNWAY_THRESHOLD"
	TrajectoryPointPropertyTypeTypeSTART_EXPECT_VECTORS TrajectoryPointPropertyTypeType = "START_EXPECT_VECTORS"
	TrajectoryPointPropertyTypeTypeSTART_TAKEOFF_ROLL TrajectoryPointPropertyTypeType = "START_TAKEOFF_ROLL"
	TrajectoryPointPropertyTypeTypeTCP_LATERAL TrajectoryPointPropertyTypeType = "TCP_LATERAL"
	TrajectoryPointPropertyTypeTypeTCP_SPEED TrajectoryPointPropertyTypeType = "TCP_SPEED"
	TrajectoryPointPropertyTypeTypeTCP_VERTICAL TrajectoryPointPropertyTypeType = "TCP_VERTICAL"
	TrajectoryPointPropertyTypeTypeTOP_OF_CLIMB TrajectoryPointPropertyTypeType = "TOP_OF_CLIMB"
	TrajectoryPointPropertyTypeTypeTOP_OF_DESCENT TrajectoryPointPropertyTypeType = "TOP_OF_DESCENT"
	TrajectoryPointPropertyTypeTypeTRANSITION_ALTITUDE_OR_LEVEL TrajectoryPointPropertyTypeType = "TRANSITION_ALTITUDE_OR_LEVEL"
	TrajectoryPointPropertyTypeTypeWHEELS_OFF TrajectoryPointPropertyTypeType = "WHEELS_OFF"
	TrajectoryPointPropertyTypeTypeWHEELS_ON TrajectoryPointPropertyTypeType = "WHEELS_ON"
)

// The location associated with estimated elapsed time. It can be a longitude, s...
type ElapsedTimeLocationChoiceType struct {
	// Longitude associated with the elapsed time.
	Longitude base.LongitudeType `xml:"longitude"`
	// Point associated with the estimated elapsed time.
	Point base.SignificantPointChoiceType `xml:"point"`
	// Flight information boundary associated with the elapsed time.
	Region base.LocationIndicatorType `xml:"region"`
}

// The estimated amount of time from takeoff to reach a significant point or Fli...
type EstimatedElapsedTimeType struct {
	// The estimated amount of elapsed time.
	ElapsedTime *base.DurationType `xml:"elapsedTime"`
	// An extension hook for attaching extension (non-core) classes.
	Extension []base.EstimatedElapsedTimeExtensionType `xml:"extension"`
	// The location associated with estimated elapsed time. It can be a longitude, s...
	Location *flight.ElapsedTimeLocationChoiceType `xml:"location"`
}

// Contains basic information about the Flight Route that pertains to the whole ...
type FlightRouteInformationType struct {
	// The internationally agreed effective date of the AIRAC (Aeronautical Informat...
	AiracReference *base.DateUtcType `xml:"airacReference"`
	// The filed altitude (flight level) for the first or the whole cruising portion...
	CruisingLevel *base.FlightLevelOrAltitudeOrVfrChoiceType `xml:"cruisingLevel"`
	// The true airspeed for the first or the whole cruising portion of the flight. ...
	CruisingSpeed *base.TrueAirspeedType `xml:"cruisingSpeed"`
	// The estimated amount of time from takeoff to reach a significant point or Fli...
	EstimatedElapsedTime []flight.EstimatedElapsedTimeType `xml:"estimatedElapsedTime"`
	// An extension hook for attaching extension (non-core) classes.
	Extension []base.FlightRouteInformationExtensionType `xml:"extension"`
	// A string of route elements complying with PANS-ATM Item 15c
	RouteText *base.CharacterStringType `xml:"routeText"`
	// For IFR flights, the estimated time required from take-off to arrive over tha...
	TotalEstimatedElapsedTime *base.DurationType `xml:"totalEstimatedElapsedTime"`
}

// In a predicted trajectory, the instantaneous temperature and wind vector used...
type MeteorologicalDataType struct {
	// An extension hook for attaching extension (non-core) classes.
	Extension []base.MeteorologicalDataExtensionType `xml:"extension"`
	// In a predicted trajectory, the instantaneous temperature used at the 4D Point...
	Temperature *base.TemperatureType `xml:"temperature"`
	// In a predicted trajectory, the instantaneous wind direction used at the 4D Po...
	WindDirection *base.WindDirectionType `xml:"windDirection"`
	// In a predicted trajectory, the instantaneous wind speed used at the 4D Point ...
	WindSpeed *base.WindSpeedType `xml:"windSpeed"`
}

// A zero-wind, standard atmosphere, unconstrained movement profile reflective o...
type PerformanceProfileType struct {
	// An extension hook for attaching extension (non-core) classes.
	Extension []base.PerformanceProfileExtensionType `xml:"extension"`
	// A point in a performance climb or descent profile. Order is based on time fro...
	ProfilePoint []flight.ProfilePointType `xml:"profilePoint"`
}

// Delay or holding planned to occur at a significant point or along a route ele...
type PlannedDelayType struct {
	// The reason for the delay. [FIXM]
	DelayReason *base.CharacterStringType `xml:"delayReason"`
	// Indicates a named hold pattern, airspace, or aerodrome at which the delay is ...
	DelayReference *base.CharacterStringType `xml:"delayReference"`
	// Indicates if the delay is planned airborne holding or Operator (Airspace User...
	DelayType *flight.PlannedDelayTypeType `xml:"delayType"`
	// The length of time the flight is expected to be delayed at a specific point.
	DelayValue *base.DurationType `xml:"delayValue"`
	// An extension hook for attaching extension (non-core) classes.
	Extension []base.PlannedDelayExtensionType `xml:"extension"`
}

// Allows to either specify a 4D point time as either an absolute time or relati...
type Point4DTimeChoiceType struct {
	// Absolute Time of the 4D point.
	AbsoluteTime base.DateTimeUtcType `xml:"absoluteTime"`
	// Relative Time of the 4D point (expressed as duration from takeoff).
	RelativeTimeFromInitialPredictionPoint base.DurationType `xml:"relativeTimeFromInitialPredictionPoint"`
}

// A point in a performance climb or descent profile.
type ProfilePointType struct {
	// The true airspeed of a point in a zero-wind, standard atmosphere, unconstrain...
	Airspeed *base.TrueAirspeedType `xml:"airspeed"`
	// The distance (from the start of the profile) at a point in a zero-wind, uncon...
	Distance *base.DistanceType `xml:"distance"`
	// An extension hook for attaching extension (non-core) classes.
	Extension []base.ProfilePointExtensionType `xml:"extension"`
	// The altitude of a point in a zero-wind, unconstrained profile.
	Level *base.FlightLevelOrAltitudeChoiceType `xml:"level"`
	// The time (from the start of the profile) at a point in a zero-wind, unconstra...
	Time *base.DurationType `xml:"time"`
}

// The route (airway) to the next element on the route. Can be either an en rout...
type RouteDesignatorToNextElementChoiceType struct {
	// Indicates that the route to next element is either DIRECT or UNSPECIFIED.
	OtherRouteDesignator flight.OtherRouteDesignatorType `xml:"otherRouteDesignator"`
	// The coded designator for a published ATS route or route segment.
	RouteDesignator base.RouteDesignatorType `xml:"routeDesignator"`
	// The textual designator of the Standard Instrument Arrival (STAR).
	StandardInstrumentArrival base.SidStarReferenceType `xml:"standardInstrumentArrival"`
	// This is the name of a published procedure extending from the departure runway...
	StandardInstrumentDeparture base.SidStarReferenceType `xml:"standardInstrumentDeparture"`
}

// A portion of the route that can be at one of three levels as described below:...
type RouteTrajectoryElementType struct {
	// The distance along the route. For an eASP-provided expanded route, it is comp...
	AlongRouteDistance *base.DistanceType `xml:"alongRouteDistance"`
	// A number of constraints can be associated with each Expanded Route Point and ...
	Constraint []flight.RouteTrajectoryConstraintType `xml:"constraint"`
	// A specified geographical location used in defining the flight route or expand...
	ElementStartPoint *base.SignificantPointChoiceType `xml:"elementStartPoint"`
	// An extension hook for attaching extension (non-core) classes.
	Extension []base.RouteTrajectoryElementExtensionType `xml:"extension"`
	// Describes the planned change of flight rules (e.g., IFR/VFR) associated with ...
	FlightRulesChange *flight.FlightRulesType `xml:"flightRulesChange"`
	// Identifies if the route point or element was modified from the reference vers...
	Modified *flight.ModifiedRouteItemIndicatorType `xml:"modified"`
	// A reference to an ATFM program name that modified the route.
	ModifiedRouteItemReference *base.CharacterStringType `xml:"modifiedRouteItemReference"`
	// Delay or holding planned to occur at a significant point or along a route ele...
	PlannedDelay *flight.PlannedDelayType `xml:"plannedDelay"`
	// Identifies the location, altitude and time of a trajectory point.
	Point4D *flight.TrajectoryPoint4DType `xml:"point4D"`
	// Flight Route element may contain a list of route changes.
	RouteChange *flight.RouteChangeType `xml:"routeChange"`
	// The route (airway) to the next element on the route. Can be either an en rout...
	RouteDesignatorToNextElement *flight.RouteDesignatorToNextElementChoiceType `xml:"routeDesignatorToNextElement"`
	// Indicates the route description is truncated at the preceding point.
	RouteTruncationIndicator *flight.RouteTruncationIndicatorType `xml:"routeTruncationIndicator"`
}

// A data container supporting the representation of aircraft movement descripti...
type RouteTrajectoryGroupType struct {
	// A zero-wind, standard atmosphere, unconstrained climb profile reflective of t...
	ClimbProfile *flight.PerformanceProfileType `xml:"climbProfile"`
	// Initially submitted by the airspace user, this defines the target speed in bo...
	ClimbSchedule *flight.SpeedScheduleType `xml:"climbSchedule"`
	// A zero-wind, standard atmosphere, unconstrained descent profile reflective of...
	DescentProfile *flight.PerformanceProfileType `xml:"descentProfile"`
	// Initially submitted by the airspace user, this defines the target speed in bo...
	DescentSchedule *flight.SpeedScheduleType `xml:"descentSchedule"`
	// A container for information pertinent to a single point in a trajectory. The ...
	Element []flight.RouteTrajectoryElementType `xml:"element"`
	// An extension hook for attaching extension (non-core) classes.
	Extension []base.RouteTrajectoryGroupExtensionType `xml:"extension"`
	// Contains information about the Flight Route. The route is comprised of Route ...
	RouteInformation *flight.FlightRouteInformationType `xml:"routeInformation"`
	// The estimated takeoff mass of the aircraft.
	TakeoffMass *base.MassType `xml:"takeoffMass"`
}

// Defines the speed schedule for climb and descent of the aircraft through the ...
type SpeedScheduleType struct {
	// An extension hook for attaching extension (non-core) classes.
	Extension []base.SpeedScheduleExtensionType `xml:"extension"`
	// Initial speed of the aircraft during the climb
	InitialSpeed *base.IndicatedAirspeedType `xml:"initialSpeed"`
	// Subsequent speed of the aircraft during the climb
	SubsequentSpeed *base.IndicatedAirspeedType `xml:"subsequentSpeed"`
}

// Identifies relevant information about a trajectory point including location, ...
type TrajectoryPoint4DType struct {
	// The barometric pressure reading used to adjust a pressure altimeter for varia...
	AltimeterSetting *base.PressureType `xml:"altimeterSetting"`
	// An extension hook for attaching extension (non-core) classes.
	Extension []base.TrajectoryPoint4DExtensionType `xml:"extension"`
	// Altitude of the 4D point.
	Level *base.FlightLevelOrAltitudeChoiceType `xml:"level"`
	// In a predicted trajectory, the instantaneous temperature and wind vector used...
	MetData *flight.MeteorologicalDataType `xml:"metData"`
	// Describes any applicable properties of the trajectory point. May include mult...
	PointProperty []flight.TrajectoryPointPropertyType `xml:"pointProperty"`
	// The geographical location of the 4D point.
	Position *base.GeographicalPositionType `xml:"position"`
	// The predicted indicated airspeed or Mach at the trajectory point.
	PredictedAirspeed *base.IndicatedAirspeedType `xml:"predictedAirspeed"`
	// The predicted ground speed at the Trajectory Point in knots or kilometers per...
	PredictedGroundspeed *base.GroundSpeedType `xml:"predictedGroundspeed"`
	// 4D Point time expressed as either absolute or relative time.
	Time *flight.Point4DTimeChoiceType `xml:"time"`
	// Vertical Range representing block altitude clearances
	VerticalRange *base.VerticalRangeType `xml:"verticalRange"`
}

// Describes any applicable properties of the trajectory point. May include mult...
type TrajectoryPointPropertyType struct {
	// Each Trajectory Point Property shall have the ability to include a textual de...
	Description *base.CharacterStringType `xml:"description"`
	// An extension hook for attaching extension (non-core) classes.
	Extension []base.TrajectoryPointPropertyExtensionType `xml:"extension"`
	// Provides the type of trajectory point property.
	PropertyType *flight.TrajectoryPointPropertyTypeType `xml:"propertyType"`
	// A reference providing an identifier of airspace or program affecting the point.
	Reference *flight.TrajectoryPointReferenceType `xml:"reference"`
}

// A reference providing an identifier of airspace or program affecting the point.
type TrajectoryPointReferenceType struct {
	// An extension hook for attaching extension (non-core) classes.
	Extension []base.TrajectoryPointReferenceExtensionType `xml:"extension"`
	// The name of the particular airspace or program (e.g., FCA016, CTP001, RRDCC123)
	Identifier *base.CharacterStringType `xml:"identifier"`
	// The nature of the airspace or program (e.g., special activity airspace, GDP, ...
	Type *base.CharacterStringType `xml:"type"`
}

