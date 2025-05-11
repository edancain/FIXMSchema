// Code generated from DangerousGoods.xsd; DO NOT EDIT.

package cargo

import (
	"github.com/edancain/FIXMSchema.git/generated/fixm/base"
)

// The number referencing the air waybill.
type AirWaybillNumberType base.CharacterStringType

// Articles or substances which are capable of posing a risk to health, safety, property or the environment and which are shown in the list of dangerous goods in the Technical Instructions or which are classified according to those Instructions. [ICAO Annex 18]
type DangerousGoodsType struct {
	// Describes whether the shipment is packed to comply with the limitations prescribed for passenger and cargo aircraft or the limitations for cargo aircraft only.
	AircraftLimitation *AircraftDangerousGoodsLimitationType `xml:"aircraftLimitation"`
	// The number referencing the air waybill.
	AirWaybillNumber *AirWaybillNumberType `xml:"airWaybillNumber"`
	// An extension hook for attaching extension (non-core) classes.
	Extension []base.DangerousGoodsExtensionType `xml:"extension"`
	// The location of a dangerous goods shipment inside the airframe.
	OnboardLocation *base.CharacterStringType `xml:"onboardLocation"`
	// A code that indicates the relative degree of danger presented by various articles and substances within a Class or Division.
	PackageGroup []DangerousGoodsPackageGroupType `xml:"packageGroup"`
	// IATA Shipper's Declaration for Dangerous Goods.
	ShippingInformation *ShippingInformationType `xml:"shippingInformation"`
}

// IATA Shipper's Declaration for Dangerous Goods.
type ShippingInformationType struct {
	// An extension hook for attaching extension (non-core) classes.
	Extension []base.ShippingInformationExtensionType `xml:"extension"`
	// Additional information related to an approval, permission, or other specific detail regarding the shipment of dangerous goods.
	ShipmentAuthorizations *base.CharacterStringType `xml:"shipmentAuthorizations"`
	// An identifier of any subsidiary hazard class(es)/division(s) in addition to the primary hazard class and division.
	SubsidiaryHazardClassAndDivision *base.CharacterStringType `xml:"subsidiaryHazardClassAndDivision"`
}

