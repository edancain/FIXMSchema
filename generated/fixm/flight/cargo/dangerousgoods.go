// Code generated from DangerousGoods.xsd; DO NOT EDIT.

package cargo

import (
	"github.com/edancain/FIXMSchema.git/generated/fixm/base"
)

// The number referencing the air waybill.
type AirWaybillNumberType base.CharacterStringType

// Articles or substances which are capable of posing a risk to health, safety, ...
type DangerousGoodsType struct {
	// Describes whether the shipment is packed to comply with the limitations presc...
	AircraftLimitation *flight.AircraftDangerousGoodsLimitationType `xml:"aircraftLimitation"`
	// The number referencing the air waybill.
	AirWaybillNumber *flight.AirWaybillNumberType `xml:"airWaybillNumber"`
	// An extension hook for attaching extension (non-core) classes.
	Extension []base.DangerousGoodsExtensionType `xml:"extension"`
	// The location of a dangerous goods shipment inside the airframe.
	OnboardLocation *base.CharacterStringType `xml:"onboardLocation"`
	// A code that indicates the relative degree of danger presented by various arti...
	PackageGroup []flight.DangerousGoodsPackageGroupType `xml:"packageGroup"`
	// IATA Shipper's Declaration for Dangerous Goods.
	ShippingInformation *flight.ShippingInformationType `xml:"shippingInformation"`
}

// IATA Shipper's Declaration for Dangerous Goods.
type ShippingInformationType struct {
	// An extension hook for attaching extension (non-core) classes.
	Extension []base.ShippingInformationExtensionType `xml:"extension"`
	// Additional information related to an approval, permission, or other specific ...
	ShipmentAuthorizations *base.CharacterStringType `xml:"shipmentAuthorizations"`
	// An identifier of any subsidiary hazard class(es)/division(s) in addition to t...
	SubsidiaryHazardClassAndDivision *base.CharacterStringType `xml:"subsidiaryHazardClassAndDivision"`
}

