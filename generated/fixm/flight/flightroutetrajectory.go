package flight/flightroutetrajectory

import (
	"github.com/yourusername/fixm/generated/fixm/base"
)
import (
	"time"
)
// FlightRulesType is Visual Flight Rules
type FlightRulesType string

// ModifiedRouteItemIndicatorType is Indication that the route item was modified
type ModifiedRouteItemIndicatorType string

// OtherRouteDesignatorType is If the element is a delay segment, the route to next element should be labeled as "Unspecified". [FF-ICE]
type OtherRouteDesignatorType string

// PlannedDelayTypeType is The delay is an Operator (Airspace User) request to 'spend time' at the segment starting at point, thus the delay has to be understood as the total duration between the point and the next one.
type PlannedDelayTypeType string

// RouteTruncationIndicatorType is Indicates the route description is truncated at the preceding point
type RouteTruncationIndicatorType string

// TrajectoryPointPropertyTypeType is Indicates that the associated trajectory point corresponds to the point at which the aircraft is predicted to be wheels on the runway for arrival.
type TrajectoryPointPropertyTypeType string

// ElapsedTimeLocationChoiceType is Flight information boundary associated with the elapsed time.
type ElapsedTimeLocationChoiceType struct {
	Longitude	*LongitudeType	`xml:"longitude"`
	Point	*SignificantPointChoiceType	`xml:"point"`
	Region	*LocationIndicatorType	`xml:"region"`
}

// EstimatedElapsedTimeType is An extension hook for attaching extension (non-core) classes.
type EstimatedElapsedTimeType struct {
	SeqNumAttr	*CountType	`xml:"seqNum,attr,omitempty"`
	ElapsedTime	*DurationType	`xml:"elapsedTime"`
	Extension	[]*EstimatedElapsedTimeExtensionType	`xml:"extension"`
	Location	interface{}	`xml:"location"`
}

// FlightRouteInformationType ...
type FlightRouteInformationType struct {
	AiracReference	*DateUtcType	`xml:"airacReference"`
	CruisingLevel	*FlightLevelOrAltitudeOrVfrChoiceType	`xml:"cruisingLevel"`
	CruisingSpeed	*TrueAirspeedType	`xml:"cruisingSpeed"`
	EstimatedElapsedTime	[]interface{}	`xml:"estimatedElapsedTime"`
	Extension	[]*FlightRouteInformationExtensionType	`xml:"extension"`
	RouteText	*base.CharacterStringType	`xml:"routeText"`
	TotalEstimatedElapsedTime	*DurationType	`xml:"totalEstimatedElapsedTime"`
}

// MeteorologicalDataType ...
type MeteorologicalDataType struct {
	Extension	[]*MeteorologicalDataExtensionType	`xml:"extension"`
	Temperature	*TemperatureType	`xml:"temperature"`
	WindDirection	*WindDirectionType	`xml:"windDirection"`
	WindSpeed	*WindSpeedType	`xml:"windSpeed"`
}

// PerformanceProfileType ...
type PerformanceProfileType struct {
	Extension	[]*PerformanceProfileExtensionType	`xml:"extension"`
	ProfilePoint	[]interface{}	`xml:"profilePoint"`
}

// PlannedDelayType ...
type PlannedDelayType struct {
	DelayReason	*base.CharacterStringType	`xml:"delayReason"`
	DelayReference	*base.CharacterStringType	`xml:"delayReference"`
	DelayType	string	`xml:"delayType"`
	DelayValue	*DurationType	`xml:"delayValue"`
	Extension	[]*PlannedDelayExtensionType	`xml:"extension"`
}

// Point4DTimeChoiceType ...
type Point4DTimeChoiceType struct {
	AbsoluteTime	*DateTimeUtcType	`xml:"absoluteTime"`
	RelativeTimeFromInitialPredictionPoint	*DurationType	`xml:"relativeTimeFromInitialPredictionPoint"`
}

// ProfilePointType ...
type ProfilePointType struct {
	SeqNumAttr	*CountType	`xml:"seqNum,attr,omitempty"`
	Airspeed	*TrueAirspeedType	`xml:"airspeed"`
	Distance	*DistanceType	`xml:"distance"`
	Extension	[]*ProfilePointExtensionType	`xml:"extension"`
	Level	*FlightLevelOrAltitudeChoiceType	`xml:"level"`
	Time	*DurationType	`xml:"time"`
}

// RouteDesignatorToNextElementChoiceType ...
type RouteDesignatorToNextElementChoiceType struct {
	OtherRouteDesignator	string	`xml:"otherRouteDesignator"`
	RouteDesignator	*RouteDesignatorType	`xml:"routeDesignator"`
	StandardInstrumentArrival	*SidStarReferenceType	`xml:"standardInstrumentArrival"`
	StandardInstrumentDeparture	*SidStarReferenceType	`xml:"standardInstrumentDeparture"`
}

// RouteTrajectoryElementType ...
type RouteTrajectoryElementType struct {
	SeqNumAttr	*CountType	`xml:"seqNum,attr,omitempty"`
	AlongRouteDistance	*DistanceType	`xml:"alongRouteDistance"`
	Constraint	[]interface{}	`xml:"constraint"`
	ElementStartPoint	*SignificantPointChoiceType	`xml:"elementStartPoint"`
	Extension	[]*RouteTrajectoryElementExtensionType	`xml:"extension"`
	FlightRulesChange	string	`xml:"flightRulesChange"`
	Modified	string	`xml:"modified"`
	ModifiedRouteItemReference	*base.CharacterStringType	`xml:"modifiedRouteItemReference"`
	PlannedDelay	interface{}	`xml:"plannedDelay"`
	Point4D	interface{}	`xml:"point4D"`
	RouteChange	interface{}	`xml:"routeChange"`
	RouteDesignatorToNextElement	interface{}	`xml:"routeDesignatorToNextElement"`
	RouteTruncationIndicator	string	`xml:"routeTruncationIndicator"`
}

// RouteTrajectoryGroupType ...
type RouteTrajectoryGroupType struct {
	ClimbProfile	interface{}	`xml:"climbProfile"`
	ClimbSchedule	interface{}	`xml:"climbSchedule"`
	DescentProfile	interface{}	`xml:"descentProfile"`
	DescentSchedule	interface{}	`xml:"descentSchedule"`
	Element	[]interface{}	`xml:"element"`
	Extension	[]*RouteTrajectoryGroupExtensionType	`xml:"extension"`
	RouteInformation	interface{}	`xml:"routeInformation"`
	TakeoffMass	*MassType	`xml:"takeoffMass"`
}

// SpeedScheduleType ...
type SpeedScheduleType struct {
	Extension	[]*SpeedScheduleExtensionType	`xml:"extension"`
	InitialSpeed	*IndicatedAirspeedType	`xml:"initialSpeed"`
	SubsequentSpeed	*IndicatedAirspeedType	`xml:"subsequentSpeed"`
}

// TrajectoryPoint4DType ...
type TrajectoryPoint4DType struct {
	AltimeterSetting	*PressureType	`xml:"altimeterSetting"`
	Extension	[]*TrajectoryPoint4DExtensionType	`xml:"extension"`
	Level	*FlightLevelOrAltitudeChoiceType	`xml:"level"`
	MetData	interface{}	`xml:"metData"`
	PointProperty	[]interface{}	`xml:"pointProperty"`
	Position	*GeographicalPositionType	`xml:"position"`
	PredictedAirspeed	*IndicatedAirspeedType	`xml:"predictedAirspeed"`
	PredictedGroundspeed	*GroundSpeedType	`xml:"predictedGroundspeed"`
	Time	time.Time	`xml:"time"`
	VerticalRange	*VerticalRangeType	`xml:"verticalRange"`
}

// TrajectoryPointPropertyType ...
type TrajectoryPointPropertyType struct {
	Description	*base.CharacterStringType	`xml:"description"`
	Extension	[]*TrajectoryPointPropertyExtensionType	`xml:"extension"`
	PropertyType	string	`xml:"propertyType"`
	Reference	interface{}	`xml:"reference"`
}

// TrajectoryPointReferenceType ...
type TrajectoryPointReferenceType struct {
	HrefAttr	*HypertextReferenceType	`xml:"href,attr,omitempty"`
	Extension	[]*TrajectoryPointReferenceExtensionType	`xml:"extension"`
	Identifier	*base.CharacterStringType	`xml:"identifier"`
	Type	*base.CharacterStringType	`xml:"type"`
}
