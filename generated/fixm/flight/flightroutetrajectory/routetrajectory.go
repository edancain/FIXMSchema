// Code generated from RouteTrajectory.xsd; DO NOT EDIT.

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
	// Indicates that the associated trajectory point is a point at which the flight is expected to begin an operation at which the flight will remain for some time.
	TrajectoryPointPropertyTypeTypeBEGIN_STAY TrajectoryPointPropertyTypeType = "BEGIN_STAY"
	// Indicates that the associated trajectory point is the point of application of a constraint. These can include explicit altitude, speed or time constraints or implicit MIT/MINIT, or sequencing constraints. For named constraints, a reference to the name of the constraint should be provided under trajectory point reference.
	TrajectoryPointPropertyTypeTypeCONSTRAINT_POINT TrajectoryPointPropertyTypeType = "CONSTRAINT_POINT"
	// Indicates that the associated trajectory point is the point at which the trajectory is predicted to cross constrained airspace designated as a line.
	TrajectoryPointPropertyTypeTypeCROSSING_CONSTRAINED_AIRSPACE TrajectoryPointPropertyTypeType = "CROSSING_CONSTRAINED_AIRSPACE"
	// The point in climb or descent where the airplane will transition between Mach Number and IAS control. (ARINC 702A)
	TrajectoryPointPropertyTypeTypeCROSSOVER_ALTITUDE TrajectoryPointPropertyTypeType = "CROSSOVER_ALTITUDE"
	// Indicates that the associated trajectory point corresponds to the point at the end of the runway on departure. This point is at the center of the runway at the departure end when departing.
	TrajectoryPointPropertyTypeTypeDEPARTURE_RUNWAY_END TrajectoryPointPropertyTypeType = "DEPARTURE_RUNWAY_END"
	// When procedures specify "Expect Vectors", the associated point identifies the ending point of the vectoring. The Trajectory Point data at this point includes an estimate of the impact of vectoring.
	TrajectoryPointPropertyTypeTypeEND_EXPECT_VECTORS TrajectoryPointPropertyTypeType = "END_EXPECT_VECTORS"
	// Indicates that the associated trajectory point corresponds to the point at which the aircraft is predicted to come to a full stop on the arrival runway. (A prediction only, the flight will likely exit the runway without coming to a full stop).
	TrajectoryPointPropertyTypeTypeEND_LANDING_ROLL TrajectoryPointPropertyTypeType = "END_LANDING_ROLL"
	// Indicates that the associated trajectory point is the final point at which a prediction was made. For FF-ICE Planning, an eASP may provide a trajectory which is predicted to end at an exit point from the eASP airspace.
	TrajectoryPointPropertyTypeTypeEND_PREDICTION_POINT TrajectoryPointPropertyTypeType = "END_PREDICTION_POINT"
	// Indicates that the associated trajectory point is a point at which the flight is expected to terminate an operation at which it remained for some time.
	TrajectoryPointPropertyTypeTypeEND_STAY TrajectoryPointPropertyTypeType = "END_STAY"
	// Indicates that the associated trajectory point is the point at which the trajectory is predicted to cross into designated constrained airspace.
	TrajectoryPointPropertyTypeTypeENTRY_CONSTRAINED_AIRSPACE TrajectoryPointPropertyTypeType = "ENTRY_CONSTRAINED_AIRSPACE"
	// Indicates that the associated trajectory point is the point at which the flight is predicted to enter an airspace restriction/reservation, including any additional safety buffer. An identifier to the airspace is provided in the trajectory point reference.
	TrajectoryPointPropertyTypeTypeENTRY_RESTRICTED_OR_RESERVED_AIRSPACE TrajectoryPointPropertyTypeType = "ENTRY_RESTRICTED_OR_RESERVED_AIRSPACE"
	// Indicates that the associated trajectory point is the point at which the trajectory is predicted to exit from designated constrained airspace.
	TrajectoryPointPropertyTypeTypeEXIT_CONSTRAINED_AIRSPACE TrajectoryPointPropertyTypeType = "EXIT_CONSTRAINED_AIRSPACE"
	// Indicates that the associated trajectory point is the point at which the flight is predicted to exit an airspace restriction/reservation, including any additional separation requirements. An identifier to the airspace is provided in the trajectory point reference.
	TrajectoryPointPropertyTypeTypeEXIT_RESTRICTED_OR_RESERVED_AIRSPACE TrajectoryPointPropertyTypeType = "EXIT_RESTRICTED_OR_RESERVED_AIRSPACE"
	// Indicates the point at which the trajectory crosses from one FIR into another. A named reference to the FIR being entered may also be provided.
	TrajectoryPointPropertyTypeTypeFIR_BOUNDARY_CROSSING_POINT TrajectoryPointPropertyTypeType = "FIR_BOUNDARY_CROSSING_POINT"
	// Indicates that the associated trajectory point is a point at which the flight is expected to enter into planned holding.
	TrajectoryPointPropertyTypeTypeHOLD_ENTRY TrajectoryPointPropertyTypeType = "HOLD_ENTRY"
	// Indicates that the associated trajectory point is a point at which the flight is expected to exit from planned holding.
	TrajectoryPointPropertyTypeTypeHOLD_EXIT TrajectoryPointPropertyTypeType = "HOLD_EXIT"
	//Indicates that the associated trajectory point is the initial point at which a prediction was made. For FF-ICE Planning, an eASP may provide a trajectory which is predicted to begin at an entry point into the eASP airspace. This includes a point near entry into the Area of Responsibility.
	TrajectoryPointPropertyTypeTypeINITIAL_PREDICTION_POINT TrajectoryPointPropertyTypeType = "INITIAL_PREDICTION_POINT"
	// Indicates the point and time at which an arriving aircraft is/was in blocks.
	TrajectoryPointPropertyTypeTypeIN_BLOCKS TrajectoryPointPropertyTypeType = "IN_BLOCKS"
	// This is the point at which the aircraft pushes back and begins to taxi for departure.
	TrajectoryPointPropertyTypeTypeOFF_BLOCKS TrajectoryPointPropertyTypeType = "OFF_BLOCKS"
	// >Indicates that the associated trajectory point represents a point that has been prescribed for required Estimated Elapsed Time reporting. This can include a FIR boundary crossing point or a significant point as prescribed on the basis of regional air navigation agreements, or by the appropriate ATS authority.
	TrajectoryPointPropertyTypeTypePRESCRIBED_EET_POINT TrajectoryPointPropertyTypeType = "PRESCRIBED_EET_POINT"
	// This point is the threshold (which may be displaced) at the centerline of the runway at the approach end when arriving. See ICAO Annex 14.
	TrajectoryPointPropertyTypeTypeRUNWAY_THRESHOLD TrajectoryPointPropertyTypeType = "RUNWAY_THRESHOLD"
	// When procedures specify "Expect Vectors", the associated point identifies the starting point of the vectoring.
	TrajectoryPointPropertyTypeTypeSTART_EXPECT_VECTORS TrajectoryPointPropertyTypeType = "START_EXPECT_VECTORS"
	// Indicates that the associated trajectory point corresponds to the point at the start of take-off roll (used for departures only).
	TrajectoryPointPropertyTypeTypeSTART_TAKEOFF_ROLL TrajectoryPointPropertyTypeType = "START_TAKEOFF_ROLL"
	// Indicates that the associated trajectory change point (TCP) is one at which a course, track or heading change will be initiated or terminated. It is not required to be provided for a planned transition between published ATS routes.
	TrajectoryPointPropertyTypeTypeTCP_LATERAL TrajectoryPointPropertyTypeType = "TCP_LATERAL"
	// The point where the airplane will begin accelerating or decelerating as a result of a speed constraint or limit, or reaches the target speed. (ARINC 702A)
	TrajectoryPointPropertyTypeTypeTCP_SPEED TrajectoryPointPropertyTypeType = "TCP_SPEED"
	// Indicates that the associated trajectory change point (TCP) is one at which a level segment (intermediate or cruise) will be initiated or terminated.
	TrajectoryPointPropertyTypeTypeTCP_VERTICAL TrajectoryPointPropertyTypeType = "TCP_VERTICAL"
	// The point where the trajectory arrives at the cruise flight level after a climb. There will be one top-of-climb for each cruise flight level (step climbs). (From ARINC 702A)
	TrajectoryPointPropertyTypeTypeTOP_OF_CLIMB TrajectoryPointPropertyTypeType = "TOP_OF_CLIMB"
	// The point where the trajectory begins a descent from the final cruise flight level.
	TrajectoryPointPropertyTypeTypeTOP_OF_DESCENT TrajectoryPointPropertyTypeType = "TOP_OF_DESCENT"
	// Indicates that the associated trajectory point is the point at which the trajectory reaches the transition altitude (in climb) or level (in descent).
	TrajectoryPointPropertyTypeTypeTRANSITION_ALTITUDE_OR_LEVEL TrajectoryPointPropertyTypeType = "TRANSITION_ALTITUDE_OR_LEVEL"
	// Indicates that the associated trajectory point corresponds to the point at which the aircraft is predicted to be wheels off the runway on departure.
	TrajectoryPointPropertyTypeTypeWHEELS_OFF TrajectoryPointPropertyTypeType = "WHEELS_OFF"
	// Indicates that the associated trajectory point corresponds to the point at which the aircraft is predicted to be wheels on the runway for arrival.
	TrajectoryPointPropertyTypeTypeWHEELS_ON TrajectoryPointPropertyTypeType = "WHEELS_ON"
)

// The location associated with estimated elapsed time. It can be a longitude, significant point or flight information region.
type ElapsedTimeLocationChoiceType struct {
	// Longitude associated with the elapsed time.
	Longitude base.LongitudeType `xml:"longitude"`
	// Point associated with the estimated elapsed time.
	Point base.SignificantPointChoiceType `xml:"point"`
	// Flight information boundary associated with the elapsed time.
	Region base.LocationIndicatorType `xml:"region"`
}

// The estimated amount of time from takeoff to reach a significant point or Flight Information Region (FIR) boundary along the route of flight.
type EstimatedElapsedTimeType struct {
	// The estimated amount of elapsed time.
	ElapsedTime *base.DurationType `xml:"elapsedTime"`
	// An extension hook for attaching extension (non-core) classes.
	Extension []base.EstimatedElapsedTimeExtensionType `xml:"extension"`
	// The location associated with estimated elapsed time. It can be a longitude, significant point or flight information region.
	Location *flight.ElapsedTimeLocationChoiceType `xml:"location"`
	/* TODO
	<xs:attribute name="seqNum" use="optional" type="fb:CountType">
			<xs:annotation>
				<xs:documentation>Incrementing identifier used to ensure the sequence of ordered lists is retained. The identifier should begin at zero. [FIXM]</xs:documentation>
			</xs:annotation>
		</xs:attribute>
		*/
}

// Contains basic information about the Flight Route that pertains to the whole flight.
type FlightRouteInformationType struct {
	// The internationally agreed effective date of the AIRAC (Aeronautical Information Regulation and Control) data set that was used in the computation of the route/trajectory.  [FIXM]
	AiracReference *base.DateUtcType `xml:"airacReference"`
	// The filed altitude (flight level) for the first or the whole cruising portion of the flight.
	CruisingLevel *base.FlightLevelOrAltitudeOrVfrChoiceType `xml:"cruisingLevel"`
	// The true airspeed for the first or the whole cruising portion of the flight.  This can be for a filed flight or an active flight.  This element is strategic in nature.
	CruisingSpeed *base.TrueAirspeedType `xml:"cruisingSpeed"`
	// The estimated amount of time from takeoff to reach a significant point or Flight Information Region (FIR) boundary along the route of flight. Elements are ordered according to increasing duration from start. 
	EstimatedElapsedTime []flight.EstimatedElapsedTimeType `xml:"estimatedElapsedTime"`
	// An extension hook for attaching extension (non-core) classes.
	Extension []base.FlightRouteInformationExtensionType `xml:"extension"`
	// A string of route elements complying with PANS-ATM Item 15c
	RouteText *base.CharacterStringType `xml:"routeText"`
	// For IFR flights, the estimated time required from take-off to arrive over that designated point, defined by reference to navigation aids, from which it is intended that an instrument approach procedure will be commenced, or, if no navigation aid is associated with the destination aerodrome, to arrive over the destination aerodrome. For VFR flights, the estimated time required from take-off to arrive over the destination aerodrome. [ICAO Doc 4444]
	// For a flight plan received from an aircraft in flight, the total estimated elapsed time is the estimated time from the first point of the route to which the flight plan applies to the termination point of the flight plan. [ICAO Doc 4444, Appendix 2, ITEM 16]
	TotalEstimatedElapsedTime *base.DurationType `xml:"totalEstimatedElapsedTime"`
}

// In a predicted trajectory, the instantaneous temperature and wind vector used at the 4D Point for creating the 4D trajectory.
type MeteorologicalDataType struct {
	// An extension hook for attaching extension (non-core) classes.
	Extension []base.MeteorologicalDataExtensionType `xml:"extension"`
	// In a predicted trajectory, the instantaneous temperature used at the 4D Point for creating the 4D trajectory.
	Temperature *base.TemperatureType `xml:"temperature"`
	// In a predicted trajectory, the instantaneous wind direction used at the 4D Point for creating the 4D trajectory.
	WindDirection *base.WindDirectionType `xml:"windDirection"`
	// In a predicted trajectory, the instantaneous wind speed used at the 4D Point for creating the 4D trajectory.
	WindSpeed *base.WindSpeedType `xml:"windSpeed"`
}

// A zero-wind, standard atmosphere, unconstrained movement profile reflective of the flight capabilities and desired parameters during climb or descent.
type PerformanceProfileType struct {
	// An extension hook for attaching extension (non-core) classes.
	Extension []base.PerformanceProfileExtensionType `xml:"extension"`
	// A point in a performance climb or descent profile.  Order is based on time from earliest to latest.
	ProfilePoint []flight.ProfilePointType `xml:"profilePoint"`
}

// Delay or holding planned to occur at a significant point or along a route element.
type PlannedDelayType struct {
	// The reason for the delay. [FIXM]
	DelayReason *base.CharacterStringType `xml:"delayReason"`
	// Indicates a named hold pattern, airspace, or aerodrome at which the delay is expected to occur.
	DelayReference *base.CharacterStringType `xml:"delayReference"`
	// Indicates if the delay is planned airborne holding or Operator (Airspace User) requested operations at a specified location.
	DelayType *flight.PlannedDelayTypeType `xml:"delayType"`
	// The length of time the flight is expected to be delayed at a specific point.
	DelayValue *base.DurationType `xml:"delayValue"`
	// An extension hook for attaching extension (non-core) classes.
	Extension []base.PlannedDelayExtensionType `xml:"extension"`
}

// Allows to either specify a 4D point time as either an absolute time or relative time.
type Point4DTimeChoiceType struct {
	// Absolute Time of the 4D point.
	AbsoluteTime base.DateTimeUtcType `xml:"absoluteTime"`
	// Relative Time of the 4D point (expressed as duration from takeoff).
	RelativeTimeFromInitialPredictionPoint base.DurationType `xml:"relativeTimeFromInitialPredictionPoint"`
}

// A point in a performance climb or descent profile.
type ProfilePointType struct {
	// The true airspeed of a point in a zero-wind, standard atmosphere, unconstrained profile.
	Airspeed *base.TrueAirspeedType `xml:"airspeed"`
	// The distance (from the start of the profile) at a point in a zero-wind, unconstrained profile.
	Distance *base.DistanceType `xml:"distance"`
	// An extension hook for attaching extension (non-core) classes.
	Extension []base.ProfilePointExtensionType `xml:"extension"`
	// The altitude of a point in a zero-wind, unconstrained profile.
	Level *base.FlightLevelOrAltitudeChoiceType `xml:"level"`
	// The time (from the start of the profile) at a point in a zero-wind, unconstrained profile.
	Time *base.DurationType `xml:"time"`
	/*
	<xs:attribute name="seqNum" use="optional" type="fb:CountType">
			<xs:annotation>
				<xs:documentation>Incrementing identifier used to ensure the sequence of ordered lists is retained. The identifier should begin at zero. [FIXM]</xs:documentation>
			</xs:annotation>
		</xs:attribute>
		*/
}

// The route (airway) to the next element on the route. Can be either an en route airway, standard instrument departure route, or standard instrument arrival route. [FIXM]
type RouteDesignatorToNextElementChoiceType struct {
	// Indicates that the route to next element is either DIRECT or UNSPECIFIED.
	OtherRouteDesignator flight.OtherRouteDesignatorType `xml:"otherRouteDesignator"`
	// The coded designator for a published ATS route or route segment.
	RouteDesignator base.RouteDesignatorType `xml:"routeDesignator"`
	// The textual designator of the Standard Instrument Arrival (STAR).
	StandardInstrumentArrival base.SidStarReferenceType `xml:"standardInstrumentArrival"`
	// This is the name of a published procedure extending from the departure runway to the start of the en route part of the aircraft's route.
	StandardInstrumentDeparture base.SidStarReferenceType `xml:"standardInstrumentDeparture"`
}

// A portion of the route that can be at one of three levels as described below:
		// a) A Route Element defined by a Significant Point and the ATS Route (or direct route) to be followed until a change in route.
		// b) An Expanded Route Element defined by a Significant Point and the ATS Route (or direct route) to be followed until the next Significant Point, which may be along the same ATS Route.
		// c)A Trajectory Element defined by a geographic point and the ATS Route (or direct route) followed until the next trajectory point. [FF-ICE/1 Implementation Guidance Manual, Appendix B, Chapter B-3.2]
type RouteTrajectoryElementType struct {
	// The distance along the route.  For an eASP-provided expanded route, it is computed from the first converted point in the eASP's airspace for each route point in the expanded route. For an operator-provided expanded route, it is computed from the beginning of the route.
	AlongRouteDistance *base.DistanceType `xml:"alongRouteDistance"`
	// A number of constraints can be associated with each Expanded Route Point and Trajectory point.
	Constraint []flight.RouteTrajectoryConstraintType `xml:"constraint"`
	// A specified geographical location used in defining the flight route or expanded route. [FIXM]
	ElementStartPoint *base.SignificantPointChoiceType `xml:"elementStartPoint"`
	// An extension hook for attaching extension (non-core) classes.
	Extension []base.RouteTrajectoryElementExtensionType `xml:"extension"`
	// Describes the planned change of flight rules (e.g., IFR/VFR) associated with a route point.
	FlightRulesChange *flight.FlightRulesType `xml:"flightRulesChange"`
	// Identifies if the route point or element was modified from the reference version, identified in the message, during negotiation by an ASP. 
	Modified *flight.ModifiedRouteItemIndicatorType `xml:"modified"`
	// A reference to an ATFM program name that modified the route.
	ModifiedRouteItemReference *base.CharacterStringType `xml:"modifiedRouteItemReference"`
	// Delay or holding planned to occur at a significant point or along a route element.
	PlannedDelay *flight.PlannedDelayType `xml:"plannedDelay"`
	// Identifies the location, altitude and time of a trajectory point.
	Point4D *flight.TrajectoryPoint4DType `xml:"point4D"`
	// Flight Route element may contain a list of route changes.
	RouteChange *flight.RouteChangeType `xml:"routeChange"`
	// The route (airway) to the next element on the route. Can be either an en route airway, standard instrument departure route, or standard instrument arrival route. [FIXM]
	RouteDesignatorToNextElement *flight.RouteDesignatorToNextElementChoiceType `xml:"routeDesignatorToNextElement"`
	// Indicates the route description is truncated at the preceding point.
	RouteTruncationIndicator *flight.RouteTruncationIndicatorType `xml:"routeTruncationIndicator"`
	/* 
	<xs:attribute name="seqNum" use="optional" type="fb:CountType">
			<xs:annotation>
				<xs:documentation>Incrementing identifier used to ensure the sequence of ordered lists is retained. The identifier should begin at zero. [FIXM]</xs:documentation>
			</xs:annotation>
		</xs:attribute>
		*/
}

// A data container supporting the representation of aircraft movement descriptions with variable levels of granularity.  This container enables the representation of a flight route and a 4D trajectory, as described in ICAO Doc 9965 Volume 2.  [FIXM]
type RouteTrajectoryGroupType struct {
	// A zero-wind, standard atmosphere, unconstrained climb profile reflective of the flight capabilities and desired parameters during climb.
	ClimbProfile *flight.PerformanceProfileType `xml:"climbProfile"`
	// Initially submitted by the airspace user, this defines the target speed in both Indicated Airspeed (IAS) and MACH so the aircraft can climb through the crossover altitude and target the MACH speed when needed.
	ClimbSchedule *flight.SpeedScheduleType `xml:"climbSchedule"`
	// A zero-wind, standard atmosphere, unconstrained descent profile reflective of the flight capabilities and desired parameters during descent.
	DescentProfile *flight.PerformanceProfileType `xml:"descentProfile"`
	// Initially submitted by the airspace user, this defines the target speed in both (Indicated Airspeed) IAS and MACH so the aircraft can descend through the crossover altitude and target the IAS speed when needed.
	DescentSchedule *flight.SpeedScheduleType `xml:"descentSchedule"`
	// A container for information pertinent to a single point in a trajectory. The elements are ordered according to increasing time/distance from the start point.
	Element []flight.RouteTrajectoryElementType `xml:"element"`
	// An extension hook for attaching extension (non-core) classes.
	Extension []base.RouteTrajectoryGroupExtensionType `xml:"extension"`
	// Contains information about the Flight Route. The route is comprised of Route Segments, which are comprised of Route Point/Airway pairs. Route also contains an expanded route.
	RouteInformation *flight.FlightRouteInformationType `xml:"routeInformation"`
	// The estimated takeoff mass of the aircraft.
	TakeoffMass *base.MassType `xml:"takeoffMass"`
}

// Defines the speed schedule for climb and descent of the aircraft through the crossover altitude.
type SpeedScheduleType struct {
	// An extension hook for attaching extension (non-core) classes.
	Extension []base.SpeedScheduleExtensionType `xml:"extension"`
	// Initial speed of the aircraft during the climb
	InitialSpeed *base.IndicatedAirspeedType `xml:"initialSpeed"`
	// Subsequent speed of the aircraft during the climb
	SubsequentSpeed *base.IndicatedAirspeedType `xml:"subsequentSpeed"`
}

// Identifies relevant information about a trajectory point including location, altitude, time, etc.
type TrajectoryPoint4DType struct {
	// The barometric pressure reading used to adjust a pressure altimeter for variations in existing atmospheric pressure or to the standard altimeter setting (29.92).
	AltimeterSetting *base.PressureType `xml:"altimeterSetting"`
	// An extension hook for attaching extension (non-core) classes.
	Extension []base.TrajectoryPoint4DExtensionType `xml:"extension"`
	// Altitude of the 4D point.
	Level *base.FlightLevelOrAltitudeChoiceType `xml:"level"`
	// In a predicted trajectory, the instantaneous temperature and wind vector used at the 4D Point for creating the 4D trajectory.
	MetData *flight.MeteorologicalDataType `xml:"metData"`
	// Describes any applicable properties of the trajectory point.  May include multiple properties per point.
	PointProperty []flight.TrajectoryPointPropertyType `xml:"pointProperty"`
	// The geographical location of the 4D point.
	Position *base.GeographicalPositionType `xml:"position"`
	// The predicted indicated airspeed or Mach at the trajectory point.
	PredictedAirspeed *base.IndicatedAirspeedType `xml:"predictedAirspeed"`
	// The predicted ground speed at the Trajectory Point in knots or kilometers per hour.
	PredictedGroundspeed *base.GroundSpeedType `xml:"predictedGroundspeed"`
	// 4D Point time expressed as either absolute or relative time.
	Time *flight.Point4DTimeChoiceType `xml:"time"`
	// Vertical Range representing block altitude clearances
	VerticalRange *base.VerticalRangeType `xml:"verticalRange"`
}

// Describes any applicable properties of the trajectory point.  May include multiple properties per point.
type TrajectoryPointPropertyType struct {
	// Each Trajectory Point Property shall have the ability to include a textual description. [FF-ICE]
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
	// The nature of the airspace or program (e.g., special activity airspace, GDP, MIT, etc.)
	Type *base.CharacterStringType `xml:"type"`
	/*
	<xs:attribute name="href" use="optional" type="fb:HypertextReferenceType">
			<xs:annotation>
				<xs:documentation>Provides an optional mechanism enabling FIXM aeronautical fields to be supplemented with references to AIXM features.  This field should be considered functionally equivalent to the xlink:href field used in AIXM to reference features. [FIXM]</xs:documentation>
			</xs:annotation>
		</xs:attribute>*/
}

