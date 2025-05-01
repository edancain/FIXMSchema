package flight/departure

// AirfileIndicatorType is Identifies a flight that has filed its flight plan while in the air, beginning its route description from a specified point en-route, and therefore may not have provided a departure aerodrome.
type AirfileIndicatorType string

// DepartureTimeTypeType is Indicates that the associated trajectory point corresponds to the point at which the aircraft is predicted to be wheels off the runway on departure.
type DepartureTimeTypeType string

// ActualTimeOfDepartureType is The type of departure time represented. [FIXM]
type ActualTimeOfDepartureType struct {
	Extension	[]*ActualTimeOfDepartureExtensionType	`xml:"extension"`
	Position	*GeographicalPositionType	`xml:"position"`
	Time	*DateTimeUtcType	`xml:"time"`
	Type	string	`xml:"type"`
}

// DepartureType is Groups information pertaining to the flight's departure.
type DepartureType struct {
	ActualTimeOfDeparture	interface{}	`xml:"actualTimeOfDeparture"`
	AirfileIndicator	string	`xml:"airfileIndicator"`
	AirportSlotIdentification	*AirportSlotIdentificationType	`xml:"airportSlotIdentification"`
	DepartureAerodrome	*AerodromeReferenceType	`xml:"departureAerodrome"`
	DeparturePoint	interface{}	`xml:"departurePoint"`
	DepartureAerodromePrevious	*AerodromeReferenceType	`xml:"departureAerodromePrevious"`
	DeparturePointPrevious	interface{}	`xml:"departurePointPrevious"`
	EstimatedOffBlockTime	*DateTimeUtcType	`xml:"estimatedOffBlockTime"`
	EstimatedRouteStartTime	*DateTimeUtcType	`xml:"estimatedRouteStartTime"`
	EstimatedOffBlockTimePrevious	*DateTimeUtcType	`xml:"estimatedOffBlockTimePrevious"`
	EstimatedRouteStartTimePrevious	*DateTimeUtcType	`xml:"estimatedRouteStartTimePrevious"`
	Extension	[]*DepartureExtensionType	`xml:"extension"`
	RunwayDirection	*RunwayDirectionDesignatorType	`xml:"runwayDirection"`
	TakeoffAlternateAerodrome	[]*AerodromeReferenceType	`xml:"takeoffAlternateAerodrome"`
}

// DeparturePointChoiceType ...
type DeparturePointChoiceType struct {
	DesignatedPoint	*DesignatedPointType	`xml:"designatedPoint"`
	Navaid	*NavaidType	`xml:"navaid"`
	Position	*GeographicalPositionType	`xml:"position"`
	RelativePoint	*RelativePointType	`xml:"relativePoint"`
}
