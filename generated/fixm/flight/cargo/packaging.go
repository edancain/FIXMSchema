// Code generated from Packaging.xsd; DO NOT EDIT.

package cargo

import (
	"github.com/edancain/FIXMSchema/generated/fixm/base"
)

// Describes whether the shipment is packed to comply with the limitations prescribed for passenger and cargo aircraft or the limitations for cargo aircraft only.
type AircraftDangerousGoodsLimitationType string

const (
	AircraftDangerousGoodsLimitationTypeCARGO_AIRCRAFT_ONLY AircraftDangerousGoodsLimitationType = "CARGO_AIRCRAFT_ONLY" // Specified that the shipment is packed to comply with the limitations prescribed for the limitations for cargo aircraft only.
	AircraftDangerousGoodsLimitationTypePASSENGER_AND_CARGO_AIRCRAFT AircraftDangerousGoodsLimitationType = "PASSENGER_AND_CARGO_AIRCRAFT" // Specified that the shipment is packed to comply with the limitations prescribed for passenger and cargo aircraft.
)

// When shipping dangerous goods, the reference to the group which identifies the kind of substances and articles deemed to be compatible. 
type CompatibilityGroupType base.CharacterStringType

// A number representing a subdivision (Division) within the Class.
type HazardDivisionType base.CharacterStringType

// A code that indicates the relative degree of danger presented by various articles and substances within a Class or Division.
type PackingGroupType string

const (
	PackingGroupTypeI PackingGroupType = "I" // Represents high danger.
	PackingGroupTypeII PackingGroupType = "II" // Represents medium danger.
	PackingGroupTypeIII PackingGroupType = "III" // Represents low danger.
)

// Helper type for restrictions on HazardClass.
type RestrictedHazardClassType base.CharacterStringType

// A four-digit identification number assigned by the United Nations Committee of Experts on the Transport of Dangerous Goods to identify a substance or a particular group of substances considered dangerous goods.
type UnNumberType base.CharacterStringType

// A statement identifying the dangerous goods listed are all contained within the same outer packaging.
type AllPackedInOneType struct {
	// An extension hook for attaching extension (non-core) classes.
	Extension []base.AllPackedInOneExtensionType `xml:"extension"`
	// The number of packages in the same outer packaging.
	NumberOfPackages *base.CountPositiveType `xml:"numberOfPackages"`
}

// Weight and volume (not height, width, and depth):
type DangerousGoodsDimensionsType struct {
	// An extension hook for attaching extension (non-core) classes.
	Extension []base.DangerousGoodsDimensionsExtensionType `xml:"extension"`
	// The total gross weight of dangerous goods transported for each unique UN number.
	GrossWeight *base.WeightType `xml:"grossWeight"`
	// The total net weight of dangerous goods transported for each unique UN number.
	NetWeight *base.WeightType `xml:"netWeight"`
	// The total displacement of dangerous goods transported for each unique UN number.
	Volume *base.VolumeType `xml:"volume"`
}

// The part of the IATA Shipper's Declaration For Dangerous Goods that contains the Line Item information for the shipment.
// The part of the IATA Shipper's Declaration For Dangerous Goods that contains the Overpack Detail for the shipment.
// The part of the IATA Shipper's Declaration For Dangerous Goods that contains the Package Details for the shipment.
type DangerousGoodsPackageType struct {
	// A statement identifying the dangerous goods listed are all contained within the same outer packaging.
	AllPackedInOne *AllPackedInOneType `xml:"allPackedInOne"`
	// When shipping dangerous goods, the reference to the group which identifies the kind of substances and articles deemed to be compatible. 
	CompatibilityGroup *CompatibilityGroupType `xml:"compatibilityGroup"`
	// Describes whether the shipment is packed to comply with the limitations prescribed for passenger and cargo aircraft or the limitations for cargo aircraft only.
	DangerousGoodsLimitation *AircraftDangerousGoodsLimitationType `xml:"dangerousGoodsLimitation"`
	// The total number of dangerous good packages of the same type and content.
	DangerousGoodsQuantity *base.CountPositiveType `xml:"dangerousGoodsQuantity"`
	// An extension hook for attaching extension (non-core) classes.
	Extension []base.DangerousGoodsPackageExtensionType `xml:"extension"`
	// A number assigned to a dangerous good that represents a classification (Class) according to the most dominant hazard it represents, potentially followed by a number representing a subdivision (Division) within the Class.
	HazardClass *HazardClassType `xml:"hazardClass"`
	// A code that indicates the relative degree of danger presented by various articles and substances within a Class or Division.
	PackingGroup *PackingGroupType `xml:"packingGroup"`
	// The name used to describe a particular article or substance in all shipping documents and notifications and, where appropriate, on packaging, as shown in the UN Model Regulations Dangerous Goods List.
	ProperShippingName *base.CharacterStringType `xml:"properShippingName"`
	// The grouping element for goods that contain radioactive materials.
	RadioactiveMaterials *RadioactiveMaterialType `xml:"radioactiveMaterials"`
	// Weight and volume (not height, width, and depth)
	ShipmentDimensions *DangerousGoodsDimensionsType `xml:"shipmentDimensions"`
	// An identifier of any subsidiary hazard class(es)/division(s) in addition to the primary hazard class and division.
	SubsidiaryHazardClass []HazardClassType `xml:"subsidiaryHazardClass"`
	// A four-digit identification number assigned by the United Nations Committee of Experts on the Transport of Dangerous Goods to identify a substance or a particular group of substances considered dangerous goods.
	UnNumber *UnNumberType `xml:"unNumber"`
}

// The part of the IATA Shipper's Declaration For Dangerous Goods that contains the Package Details for the shipment.
type DangerousGoodsPackageGroupType struct {
	// A collection of at least one DangerousGoodsPackage.
	DangerousGoodsPackage []DangerousGoodsPackageType `xml:"dangerousGoodsPackage"`
	// An extension hook for attaching extension (non-core) classes.
	Extension []base.DangerousGoodsPackageGroupExtensionType `xml:"extension"`
	// Weight and volume (not height, width, and depth)
	ShipmentDimensions *DangerousGoodsDimensionsType `xml:"shipmentDimensions"`
}

// A number assigned to a dangerous good that represents a classification (Class) according to the most dominant hazard it represents, potentially followed by a number representing a subdivision (Division) within the Class.
// An identifier of any subsidiary hazard class(es)/division(s) in addition to the primary hazard class and division.
type HazardClassType struct {
	// A number assigned to a dangerous good that represents a classification (Class) according to the most dominant hazard it represents.
	Class *RestrictedHazardClassType `xml:"class"`
	// A number representing a subdivision (Division) within the Class.
	Division *HazardDivisionType `xml:"division"`
	// An extension hook for attaching extension (non-core) classes.
	Extension []base.HazardClassExtensionType `xml:"extension"`
}

