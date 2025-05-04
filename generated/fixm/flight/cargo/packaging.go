// Code generated from Packaging.xsd; DO NOT EDIT.

package cargo

import (
	"github.com/edancain/FIXMSchema.git/generated/fixm/base"
)

// Describes whether the shipment is packed to comply with the limitations presc...
type AircraftDangerousGoodsLimitationType string

const (
	AircraftDangerousGoodsLimitationTypeCARGO_AIRCRAFT_ONLY AircraftDangerousGoodsLimitationType = "CARGO_AIRCRAFT_ONLY"
	AircraftDangerousGoodsLimitationTypePASSENGER_AND_CARGO_AIRCRAFT AircraftDangerousGoodsLimitationType = "PASSENGER_AND_CARGO_AIRCRAFT"
)

// When shipping dangerous goods, the reference to the group which identifies th...
type CompatibilityGroupType base.CharacterStringType

// A number representing a subdivision (Division) within the Class.
type HazardDivisionType base.CharacterStringType

// A code that indicates the relative degree of danger presented by various arti...
type PackingGroupType string

const (
	PackingGroupTypeI PackingGroupType = "I"
	PackingGroupTypeII PackingGroupType = "II"
	PackingGroupTypeIII PackingGroupType = "III"
)

// Helper type for restrictions on HazardClass.
type RestrictedHazardClassType base.CharacterStringType

// A four-digit identification number assigned by the United Nations Committee o...
type UnNumberType base.CharacterStringType

// A statement identifying the dangerous goods listed are all contained within t...
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

// The part of the IATA Shipper's Declaration For Dangerous Goods that contains ...
type DangerousGoodsPackageType struct {
	// A statement identifying the dangerous goods listed are all contained within t...
	AllPackedInOne *flight.AllPackedInOneType `xml:"allPackedInOne"`
	// When shipping dangerous goods, the reference to the group which identifies th...
	CompatibilityGroup *flight.CompatibilityGroupType `xml:"compatibilityGroup"`
	// Describes whether the shipment is packed to comply with the limitations presc...
	DangerousGoodsLimitation *flight.AircraftDangerousGoodsLimitationType `xml:"dangerousGoodsLimitation"`
	// The total number of dangerous good packages of the same type and content.
	DangerousGoodsQuantity *base.CountPositiveType `xml:"dangerousGoodsQuantity"`
	// An extension hook for attaching extension (non-core) classes.
	Extension []base.DangerousGoodsPackageExtensionType `xml:"extension"`
	// A number assigned to a dangerous good that represents a classification (Class...
	HazardClass *flight.HazardClassType `xml:"hazardClass"`
	// A code that indicates the relative degree of danger presented by various arti...
	PackingGroup *flight.PackingGroupType `xml:"packingGroup"`
	// The name used to describe a particular article or substance in all shipping d...
	ProperShippingName *base.CharacterStringType `xml:"properShippingName"`
	// The grouping element for goods that contain radioactive materials.
	RadioactiveMaterials *flight.RadioactiveMaterialType `xml:"radioactiveMaterials"`
	// Weight and volume (not height, width, and depth)
	ShipmentDimensions *flight.DangerousGoodsDimensionsType `xml:"shipmentDimensions"`
	// An identifier of any subsidiary hazard class(es)/division(s) in addition to t...
	SubsidiaryHazardClass []flight.HazardClassType `xml:"subsidiaryHazardClass"`
	// A four-digit identification number assigned by the United Nations Committee o...
	UnNumber *flight.UnNumberType `xml:"unNumber"`
}

// The part of the IATA Shipper's Declaration For Dangerous Goods that contains ...
type DangerousGoodsPackageGroupType struct {
	// A collection of at least one DangerousGoodsPackage.
	DangerousGoodsPackage []flight.DangerousGoodsPackageType `xml:"dangerousGoodsPackage"`
	// An extension hook for attaching extension (non-core) classes.
	Extension []base.DangerousGoodsPackageGroupExtensionType `xml:"extension"`
	// Weight and volume (not height, width, and depth)
	ShipmentDimensions *flight.DangerousGoodsDimensionsType `xml:"shipmentDimensions"`
}

// A number assigned to a dangerous good that represents a classification (Class...
type HazardClassType struct {
	// A number assigned to a dangerous good that represents a classification (Class...
	Class *flight.RestrictedHazardClassType `xml:"class"`
	// A number representing a subdivision (Division) within the Class.
	Division *flight.HazardDivisionType `xml:"division"`
	// An extension hook for attaching extension (non-core) classes.
	Extension []base.HazardClassExtensionType `xml:"extension"`
}

