package flight/enroute

// BoundaryCrossingConditionType is Indicates that the transition altitude is at or below the specified.
type BoundaryCrossingConditionType string

// AltitudeInTransitionType is An altitude (flight level) at or above/below which (specified in Boundary Crossing Condition) an aircraft will cross the associated boundary point.
type AltitudeInTransitionType struct {
	CrossingCondition	string	`xml:"crossingCondition"`
	Extension	[]*AltitudeInTransitionExtensionType	`xml:"extension"`
	Level	*FlightLevelOrAltitudeChoiceType	`xml:"level"`
}

// BoundaryCrossingType is Boundary Crossing contains estimate data conveyed between ATS Units during the process of Controller Coordination. [FIXM]
type BoundaryCrossingType struct {
	AltitudeInTransition	interface{}	`xml:"altitudeInTransition"`
	ClearedLevel	*FlightLevelOrAltitudeChoiceType	`xml:"clearedLevel"`
	CrossingPoint	*SignificantPointChoiceType	`xml:"crossingPoint"`
	CrossingTime	*DateTimeUtcType	`xml:"crossingTime"`
	Extension	[]*BoundaryCrossingExtensionType	`xml:"extension"`
}

// EnRouteType ...
type EnRouteType struct {
	AlternateAerodrome	[]*AerodromeReferenceType	`xml:"alternateAerodrome"`
	BoundaryCrossingCoordination	interface{}	`xml:"boundaryCrossingCoordination"`
	CurrentModeACode	*ModeACodeType	`xml:"currentModeACode"`
	Extension	[]*EnRouteExtensionType	`xml:"extension"`
}
