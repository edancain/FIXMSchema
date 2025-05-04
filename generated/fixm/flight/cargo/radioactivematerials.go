// Code generated from RadioactiveMaterials.xsd; DO NOT EDIT.

package cargo

import (
	"github.com/edancain/FIXMSchema.git/generated/fixm/base"
)

// The dimensionless number (rounded up to the next tenth) assigned to and place...
type CriticalSafetyIndexType base.DecimalIndexType

// A category used for radioactive materials in a package, overpack or freight c...
type RadioactiveMaterialCategoryType string

const (
	RadioactiveMaterialCategoryTypeI_WHITE RadioactiveMaterialCategoryType = "I_WHITE"
	RadioactiveMaterialCategoryTypeII_YELLOW RadioactiveMaterialCategoryType = "II_YELLOW"
	RadioactiveMaterialCategoryTypeIII_YELLOW RadioactiveMaterialCategoryType = "III_YELLOW"
)

// A figure representing the radiation level measured at one meter from the pack...
type TransportIndexType base.DecimalIndexType

// The grouping element for goods that contain radioactive materials.
type RadioactiveMaterialType struct {
	// A category used for radioactive materials in a package, overpack or freight c...
	Category *flight.RadioactiveMaterialCategoryType `xml:"category"`
	// The dimensionless number (rounded up to the next tenth) assigned to and place...
	CriticalSafetyIndex *flight.CriticalSafetyIndexType `xml:"criticalSafetyIndex"`
	// An extension hook for attaching extension (non-core) classes.
	Extension []base.RadioactiveMaterialExtensionType `xml:"extension"`
	// A figure representing the radiation level measured at one meter from the pack...
	TransportIndex *flight.TransportIndexType `xml:"transportIndex"`
}

