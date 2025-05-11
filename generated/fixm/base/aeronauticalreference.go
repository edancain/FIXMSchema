// Code generated from AeronauticalReference.xsd; DO NOT EDIT.

package base

// The primary official name of an aerodrome as designated by an appropriate authority. [AIXM 5.1]
type AerodromeNameType TextNameType

// The name of the ATC center with authority over the corresponding Flight Information Region (FIR). [FIXM]
type AtcUnitNameType CharacterStringType

// The name of the designated point.
type DesignatedPointDesignatorType CharacterStringType

// The three letter coded location identifier of an aerodrome according to the IATA Resolution 763. [adapted from AIXM 5.1]
type IataAerodromeDesignatorType CharacterStringType

// A series of coordinates limited to two instances.  Latitude followed by longitude.
type LatLongPosType LatLongPosListType

// A series of coordinates.
type LatLongPosListType []float64

// A four-letter code group formulated in accordance with rules prescribed by ICAO and assigned to an aerodrome, a region or an ATS unit. [ICAO Doc 4444]
type LocationIndicatorType CharacterStringType

// Longitude value expressed to a decimal precision.
type LongitudeType MeasureType

// A name of the Navaid.
type NavaidDesignatorType CharacterStringType

// Types of Navaid Services. [AIXM 5.1]
type NavaidServiceTypeType string

const (
	NavaidServiceTypeTypeDF NavaidServiceTypeType = "DF" // Direction Finder.
	NavaidServiceTypeTypeDME NavaidServiceTypeType = "DME" // Distance Measuring Equipment.
	NavaidServiceTypeTypeILS NavaidServiceTypeType = "ILS" // Instrument Landing System
	NavaidServiceTypeTypeILS_DME NavaidServiceTypeType = "ILS_DME" // ILS with collocated DME. Collocation can be with the antenna or the loc.
	NavaidServiceTypeTypeLOC NavaidServiceTypeType = "LOC" // Localizer
	NavaidServiceTypeTypeLOC_DME NavaidServiceTypeType = "LOC_DME" // Localizer and DME collocated
	NavaidServiceTypeTypeMKR NavaidServiceTypeType = "MKR" // Marker Beacon
	NavaidServiceTypeTypeMLS NavaidServiceTypeType = "MLS" // Microwave Landing System
	NavaidServiceTypeTypeMLS_DME NavaidServiceTypeType = "MLS_DME" // MLS and DME collocated
	NavaidServiceTypeTypeNDB NavaidServiceTypeType = "NDB" // Non-Directional Radio Beacon
	NavaidServiceTypeTypeNDB_DME NavaidServiceTypeType = "NDB_DME" // NDB and DME collocated
	NavaidServiceTypeTypeNDB_MKR NavaidServiceTypeType = "NDB_MKR" // Non-Directional Radio Beacon and Marker Beacon
	NavaidServiceTypeTypeSDF NavaidServiceTypeType = "SDF" // Simplified Directional Facility
	NavaidServiceTypeTypeTACAN NavaidServiceTypeType = "TACAN" // Tactical Air Navigation Beacon
	NavaidServiceTypeTypeTLS NavaidServiceTypeType = "TLS" // Transponder Landing System
	NavaidServiceTypeTypeVOR NavaidServiceTypeType = "VOR" // VHF Omni directional Radio Range.
	NavaidServiceTypeTypeVOR_DME NavaidServiceTypeType = "VOR_DME" // VOR and DME collocated
	NavaidServiceTypeTypeVORTAC NavaidServiceTypeType = "VORTAC" // VOR and TACAN collocated.
)

// Helper type for restrictions on AirspaceDesignator.
type RestrictedAirspaceDesignatorType CharacterStringType

// Helper type for restrictions on RouteDesignator.
type RestrictedRouteDesignatorType CharacterStringType

// Helper type for restrictions on RunwayDirectionDesignator.
type RestrictedRunwayDirectionDesignatorType CharacterStringType

// A shortened designator of the SID or STAR. [FIXM] This abbreviated designator is based on ARINC 424 chapter 7.4 rules.
type SidStarAbbreviatedDesignatorType CharacterStringType

// The textual designator of a SID or STAR. [Specialised from AIXM 5.1]
type SidStarDesignatorType CharacterStringType

// An aerodrome reference being either:
	// - the ICAO location indicator of the aerodrome, as listed in ICAO Doc 7910 (E.g. "KDFW")
	// - if no location indicator is available, the aerodrome name (E.g. "Dallas Fort Worth") or the identifier assigned to the aerodrome location in accordance with rules (resolution 767) governed by the International Air Transport Association. (E.g. "DFW")
[FIXM]
type AerodromeReferenceType struct {
	// An extension hook for attaching extension (non-core) classes.
	Extension []AerodromeReferenceExtensionType `xml:"extension"`
	// The three letter coded location identifier of an aerodrome according to the IATA Resolution 763. [adapted from AIXM 5.1]
	IataDesignator *IataAerodromeDesignatorType `xml:"iataDesignator"`
	// A four-letter code group formulated in accordance with rules prescribed by ICAO and assigned to the aerodrome. [Specialised from ICAO Annex 10] 
	// The list of ICAO location indicators is provided by ICAO Doc 7910. [FIXM]	
	LocationIndicator *LocationIndicatorType `xml:"locationIndicator"`
	// The primary official name of an aerodrome as designated by an appropriate authority. [AIXM 5.1]
	Name *AerodromeNameType `xml:"name"`
	// The designated geographical location of an aerodrome. [ICAO]
	ReferencePoint *GeographicalPositionType `xml:"referencePoint"`
	// The location of an aerodrome expressed as a relative point.
	ReferenceRelativePoint *RelativePointType `xml:"referenceRelativePoint"`
}

// A published sequence of characters allowing the identification of the airspace. [AIXM 5.1]
type AirspaceDesignatorType struct {
}

// A reference to an area control centre, approach control unit or aerodrome control tower being either:
	// - the ICAO location indicator of the atc unit, as listed in ICAO Doc 7910
	// - if no ICAO location indicator is available, the ATC unit name and optionally its geographical position. [FIXM
type AtcUnitReferenceType struct {
	// The full textual name of the ATC unit [adapted from AIXM 5.1] or an alternate identifier for the unit. [FIXM]
	AtcUnitNameOrAlternate *TextNameType `xml:"atcUnitNameOrAlternate"`
	// A published sequence of characters allowing the identification of a subdivision of a designated control area within which responsibility is assigned to one controller or to a small group of controllers.
	ControlSectorDesignator *AirspaceDesignatorType `xml:"controlSectorDesignator"`
	// An extension hook for attaching extension (non-core) classes.
	Extension []AtcUnitReferenceExtensionType `xml:"extension"`
	// A four-letter code group formulated in accordance with rules prescribed by ICAO and assigned to the atc unit. [Specialised from ICAO Annex 10] The list of ICAO location indicators is provided by ICAO Doc 7910.
	LocationIndicator *LocationIndicatorType `xml:"locationIndicator"`
	// The geographical point of the Unit. [AIXM 5.1]
	Position *GeographicalPositionType `xml:"position"`
	// TODO href missing
}

// A named geographical location used in defining an ATS route, the flight path of an aircraft or for other navigation or ATS purposes. [FIXM]
// The coded designator of a designated point is not always sufficient for unambiguously referring to that feature: 
// The 5-letter coded designator of a waypoint is supposed to be unique world-wide (according to ICAO Annex 11), but in reality it is not. 
// The combinations of fields that would make the references unique are name + position.
// This FIXM class adds an optional properties 'position' which may be used as a complement to the 'name' information in order to remove any ambiguity on the coded designator.
type DesignatedPointType struct {
	// The name of the designated point.
	Designator DesignatedPointDesignatorType `xml:"designator"`
	// An extension hook for attaching extension (non-core) classes.
	Extension []DesignatedPointExtensionType `xml:"extension"`
	// The values of latitude and longitude that define the position of the Designated Point on the surface of the Earth with respect to a reference datum. [specialised from ICAO Doc 9881]
	// This optional field may be used to achieve unambiguous reference to the designated point. 
	// The combinations of fields that would make the reference unique are name + position
	Position *GeographicalPositionType `xml:"position"`
	// TODO href
}

// An unnamed point designated only with a set of coordinates - latitude/longitude. [FIXM]
// Set of coordinates (latitude and longitude) referenced to the mathematical reference ellipsoid which define the position of a  point on the surface of the Earth. [ICAO Annex 15]
// The "srsName" attribute names the coordinate reference system (CRS) that defines the semantics of the lat/long pair according to the ISO6709 standard. FIXM uses only "urn:ogc:def:crs:EPSG::4326".
type GeographicalPositionType struct {
	// An extension hook for attaching extension (non-core) classes.
	Extension []GeographicalPositionExtensionType `xml:"extension"`
	// This list of doubles contains the latitude and longitude of the location in order of LATITUDE FIRST, THEN LONGITUDE.
	Pos LatLongPosType `xml:"pos"`
	// TODO
	// Think we are missing srsName
	// Names the coordinate reference system (CRS) that defines the semantics of the lat/long pair according to the ISO6709 standard. FIXM uses only "urn:ogc:def:crs:EPSG::4326".
}

// A radio navigation aid used in defining an ATS route, the flight path of an aircraft or for other navigation or ATS purposes. [FIXM]
// The coded designator of a navaid is not always sufficient for unambiguously referring to that feature: 
// The en-route navaids (VOR, DME, NDB) designator is supposed to be unique (according to ICAO Annex 11) within 600 NM. This means that these designators are not unique world-wide. For airport navaids, there is no limitation.
// The combinations of fields that would make the references unique are designator + position + navaid service type.
// This FIXM class adds two optional properties 'position' and 'navaidServiceType' which may be used as a complement to the 'designator' information in order to remove any ambiguity on the coded designator.
type NavaidType struct {
	// The name of the navaid.
	Designator NavaidDesignatorType `xml:"designator"`
	// An extension hook for attaching extension (non-core) classes.
	Extension []NavaidExtensionType `xml:"extension"`
	// Type of the navaid service. [AIXM 5.1]
	// This optional field may be used as a complement of the designator in order to achieve unambiguous reference to the navaid. 
	// The combinations of fields that would make the reference unique are designator + position + navaid service type.
	NavaidServiceType *NavaidServiceTypeType `xml:"navaidServiceType"`
	// The values of latitude and longitude that define the position of the Navaid on the surface of the Earth with respect to a reference datum. [specialised from ICAO Doc 9881]
	// This optional field may be used to achieve unambiguous reference to the navaid. 
	// The combinations of fields that would make the reference unique are name + position + navaid service type.
	Position *GeographicalPositionType `xml:"position"`
	// TODO href
}

// Bearing and distance from a reference point. [ICAO Doc 4444, Appendix 2, ITEM 15]
// This FIXM class adds one optional property 'position' which may be used as a complement to the bearing, distance and reference point information in order to provide the computed position of the relative point, if available.
type RelativePointType struct {
	// The radius from a chosen Navaid
	Bearing BearingType `xml:"bearing"`
	// Distance from the chosen Navaid to the Intersection.
	Distance DistanceType `xml:"distance"`
	// An extension hook for attaching extension (non-core) classes.
	Extension []RelativePointExtensionType `xml:"extension"`
	// The values of latitude and longitude that define the position of the Relative Point on the surface of the Earth with respect to a reference datum. [specialised from ICAO Doc 9881]
	// This optional field may be used to provide the actual position of the relative point if already known / computed.
	Position *GeographicalPositionType `xml:"position"`
	// The navaid used as reference for building an intersection
	ReferencePoint NavaidType `xml:"referencePoint"`
}

// The coded designator for a published ATS route or route segment.
type RouteDesignatorType struct {
	// TODO href
}

// - A number between 01 and 36 - Optionally Left (L), Center (C), or Right (R)
type RunwayDirectionDesignatorType struct {
	// TODO href
}

// A reference to a SID or a STAR. [FIXM]
type SidStarReferenceType struct {
	// A shortened designator of the SID or STAR. [FIXM] This abbreviated designator is based on ARINC 424 chapter 7.4 rules.
	AbbreviatedDesignator *SidStarAbbreviatedDesignatorType `xml:"abbreviatedDesignator"`
	// The textual designator of a SID or STAR. [Specialised from AIXM 5.1]
	Designator SidStarDesignatorType `xml:"designator"`
	// An extension hook for attaching extension (non-core) classes.
	Extension []SidStarReferenceExtensionType `xml:"extension"`
	//TODO href
}

// A location type restricted to lat/long location, named location, or reference location.
type SignificantPointChoiceType struct {
	// The designated geographical location of an aerodrome. [ICAO Annex 14].
	AerodromeReferencePoint AerodromeReferenceType `xml:"aerodromeReferencePoint"`
	// A named geographical location used in defining an ATS route, the flight path of an aircraft or for other navigation or ATS purposes. [FIXM]
	DesignatedPoint DesignatedPointType `xml:"designatedPoint"`
	// A radio navigation aid used in defining an ATS route, the flight path of an aircraft or for other navigation or ATS purposes. [FIXM]
	Navaid NavaidType `xml:"navaid"`
	// The values of latitude and longitude that define the position of a point on the surface of the Earth with respect to a reference datum. [specialised from ICAO Doc 9881]
	Position GeographicalPositionType `xml:"position"`
	// Bearing and distance from a reference point. [ICAO Doc 4444, Appendix 2, ITEM 15]
	RelativePoint RelativePointType `xml:"relativePoint"`
}

