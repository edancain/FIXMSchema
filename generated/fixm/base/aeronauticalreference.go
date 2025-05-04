// Code generated from AeronauticalReference.xsd; DO NOT EDIT.

package base

// The primary official name of an aerodrome as designated by an appropriate aut...
type AerodromeNameType TextNameType

// The name of the ATC center with authority over the corresponding Flight Infor...
type AtcUnitNameType CharacterStringType

// The name of the designated point.
type DesignatedPointDesignatorType CharacterStringType

// The three letter coded location identifier of an aerodrome according to the I...
type IataAerodromeDesignatorType CharacterStringType

// A series of coordinates limited to two instances. Latitude followed by longit...
type LatLongPosType LatLongPosListType

// A series of coordinates.
type LatLongPosListType []float64

// A four-letter code group formulated in accordance with rules prescribed by IC...
type LocationIndicatorType CharacterStringType

// Longitude value expressed to a decimal precision.
type LongitudeType MeasureType

// A name of the Navaid.
type NavaidDesignatorType CharacterStringType

// Types of Navaid Services. [AIXM 5.1]
type NavaidServiceTypeType string

const (
	NavaidServiceTypeTypeDF NavaidServiceTypeType = "DF"
	NavaidServiceTypeTypeDME NavaidServiceTypeType = "DME"
	NavaidServiceTypeTypeILS NavaidServiceTypeType = "ILS"
	NavaidServiceTypeTypeILS_DME NavaidServiceTypeType = "ILS_DME"
	NavaidServiceTypeTypeLOC NavaidServiceTypeType = "LOC"
	NavaidServiceTypeTypeLOC_DME NavaidServiceTypeType = "LOC_DME"
	NavaidServiceTypeTypeMKR NavaidServiceTypeType = "MKR"
	NavaidServiceTypeTypeMLS NavaidServiceTypeType = "MLS"
	NavaidServiceTypeTypeMLS_DME NavaidServiceTypeType = "MLS_DME"
	NavaidServiceTypeTypeNDB NavaidServiceTypeType = "NDB"
	NavaidServiceTypeTypeNDB_DME NavaidServiceTypeType = "NDB_DME"
	NavaidServiceTypeTypeNDB_MKR NavaidServiceTypeType = "NDB_MKR"
	NavaidServiceTypeTypeSDF NavaidServiceTypeType = "SDF"
	NavaidServiceTypeTypeTACAN NavaidServiceTypeType = "TACAN"
	NavaidServiceTypeTypeTLS NavaidServiceTypeType = "TLS"
	NavaidServiceTypeTypeVOR NavaidServiceTypeType = "VOR"
	NavaidServiceTypeTypeVOR_DME NavaidServiceTypeType = "VOR_DME"
	NavaidServiceTypeTypeVORTAC NavaidServiceTypeType = "VORTAC"
)

// Helper type for restrictions on AirspaceDesignator.
type RestrictedAirspaceDesignatorType CharacterStringType

// Helper type for restrictions on RouteDesignator.
type RestrictedRouteDesignatorType CharacterStringType

// Helper type for restrictions on RunwayDirectionDesignator.
type RestrictedRunwayDirectionDesignatorType CharacterStringType

// A shortened designator of the SID or STAR. [FIXM] This abbreviated designator...
type SidStarAbbreviatedDesignatorType CharacterStringType

// The textual designator of a SID or STAR. [Specialised from AIXM 5.1]
type SidStarDesignatorType CharacterStringType

// An aerodrome reference being either: - the ICAO location indicator of the aer...
type AerodromeReferenceType struct {
	// An extension hook for attaching extension (non-core) classes.
	Extension []AerodromeReferenceExtensionType `xml:"extension"`
	// The three letter coded location identifier of an aerodrome according to the I...
	IataDesignator *IataAerodromeDesignatorType `xml:"iataDesignator"`
	// A four-letter code group formulated in accordance with rules prescribed by IC...
	LocationIndicator *LocationIndicatorType `xml:"locationIndicator"`
	// The primary official name of an aerodrome as designated by an appropriate aut...
	Name *AerodromeNameType `xml:"name"`
	// The designated geographical location of an aerodrome. [ICAO]
	ReferencePoint *GeographicalPositionType `xml:"referencePoint"`
	// The location of an aerodrome expressed as a relative point.
	ReferenceRelativePoint *RelativePointType `xml:"referenceRelativePoint"`
}

// A published sequence of characters allowing the identification of the airspac...
type AirspaceDesignatorType struct {
}

// A reference to an area control centre, approach control unit or aerodrome con...
type AtcUnitReferenceType struct {
	// The full textual name of the ATC unit [adapted from AIXM 5.1] or an alternate...
	AtcUnitNameOrAlternate *TextNameType `xml:"atcUnitNameOrAlternate"`
	// A published sequence of characters allowing the identification of a subdivisi...
	ControlSectorDesignator *AirspaceDesignatorType `xml:"controlSectorDesignator"`
	// An extension hook for attaching extension (non-core) classes.
	Extension []AtcUnitReferenceExtensionType `xml:"extension"`
	// A four-letter code group formulated in accordance with rules prescribed by IC...
	LocationIndicator *LocationIndicatorType `xml:"locationIndicator"`
	// The geographical point of the Unit. [AIXM 5.1]
	Position *GeographicalPositionType `xml:"position"`
}

// A named geographical location used in defining an ATS route, the flight path ...
type DesignatedPointType struct {
	// The name of the designated point.
	Designator DesignatedPointDesignatorType `xml:"designator"`
	// An extension hook for attaching extension (non-core) classes.
	Extension []DesignatedPointExtensionType `xml:"extension"`
	// The values of latitude and longitude that define the position of the Designat...
	Position *GeographicalPositionType `xml:"position"`
}

// An unnamed point designated only with a set of coordinates - latitude/longitu...
type GeographicalPositionType struct {
	// An extension hook for attaching extension (non-core) classes.
	Extension []GeographicalPositionExtensionType `xml:"extension"`
	// This list of doubles contains the latitude and longitude of the location in o...
	Pos LatLongPosType `xml:"pos"`
}

// A radio navigation aid used in defining an ATS route, the flight path of an a...
type NavaidType struct {
	// The name of the navaid.
	Designator NavaidDesignatorType `xml:"designator"`
	// An extension hook for attaching extension (non-core) classes.
	Extension []NavaidExtensionType `xml:"extension"`
	// Type of the navaid service. [AIXM 5.1] This optional field may be used as a c...
	NavaidServiceType *NavaidServiceTypeType `xml:"navaidServiceType"`
	// The values of latitude and longitude that define the position of the Navaid o...
	Position *GeographicalPositionType `xml:"position"`
}

// Bearing and distance from a reference point. [ICAO Doc 4444, Appendix 2, ITEM...
type RelativePointType struct {
	// The radius from a chosen Navaid
	Bearing BearingType `xml:"bearing"`
	// Distance from the chosen Navaid to the Intersection.
	Distance DistanceType `xml:"distance"`
	// An extension hook for attaching extension (non-core) classes.
	Extension []RelativePointExtensionType `xml:"extension"`
	// The values of latitude and longitude that define the position of the Relative...
	Position *GeographicalPositionType `xml:"position"`
	// The navaid used as reference for building an intersection
	ReferencePoint NavaidType `xml:"referencePoint"`
}

// The coded designator for a published ATS route or route segment.
type RouteDesignatorType struct {
}

// - A number between 01 and 36 - Optionally Left (L), Center (C), or Right (R)
type RunwayDirectionDesignatorType struct {
}

// A reference to a SID or a STAR. [FIXM]
type SidStarReferenceType struct {
	// A shortened designator of the SID or STAR. [FIXM] This abbreviated designator...
	AbbreviatedDesignator *SidStarAbbreviatedDesignatorType `xml:"abbreviatedDesignator"`
	// The textual designator of a SID or STAR. [Specialised from AIXM 5.1]
	Designator SidStarDesignatorType `xml:"designator"`
	// An extension hook for attaching extension (non-core) classes.
	Extension []SidStarReferenceExtensionType `xml:"extension"`
}

// A location type restricted to lat/long location, named location, or reference...
type SignificantPointChoiceType struct {
	// The designated geographical location of an aerodrome. [ICAO Annex 14].
	AerodromeReferencePoint AerodromeReferenceType `xml:"aerodromeReferencePoint"`
	// A named geographical location used in defining an ATS route, the flight path ...
	DesignatedPoint DesignatedPointType `xml:"designatedPoint"`
	// A radio navigation aid used in defining an ATS route, the flight path of an a...
	Navaid NavaidType `xml:"navaid"`
	// The values of latitude and longitude that define the position of a point on t...
	Position GeographicalPositionType `xml:"position"`
	// Bearing and distance from a reference point. [ICAO Doc 4444, Appendix 2, ITEM...
	RelativePoint RelativePointType `xml:"relativePoint"`
}

