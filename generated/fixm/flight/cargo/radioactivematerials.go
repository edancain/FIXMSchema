// Code generated from RadioactiveMaterials.xsd; DO NOT EDIT.

package cargo

import (
	"github.com/edancain/FIXMSchema/generated/fixm/base"
)

// The dimensionless number (rounded up to the next tenth) assigned to and placed on the label of a fissile material package to designate the degree of control of accumulation of packages containing fissile material during transportation.
type CriticalSafetyIndexType base.DecimalIndexType

// A category used for radioactive materials in a package, overpack or freight container, based on their maximum radiation level.
type RadioactiveMaterialCategoryType string

const (
	// .Surface radiation &lt;0.5 millirem/hr, 1 meter radiation: N/A
	RadioactiveMaterialCategoryTypeI_WHITE RadioactiveMaterialCategoryType = "I_WHITE"
	// 
	RadioactiveMaterialCategoryTypeII_YELLOW RadioactiveMaterialCategoryType = "II_YELLOW"
	// Surface radiation &gt;50 millirem/hr, 1 meter radiation &gt;1 millirem/hr
	RadioactiveMaterialCategoryTypeIII_YELLOW RadioactiveMaterialCategoryType = "III_YELLOW"
)

// A figure representing the radiation level measured at one meter from the package.
type TransportIndexType base.DecimalIndexType

// The grouping element for goods that contain radioactive materials.
type RadioactiveMaterialType struct {
	// >A category used for radioactive materials in a package, overpack or freight container, based on their maximum radiation level.
	Category *RadioactiveMaterialCategoryType `xml:"category"`
	// The dimensionless number (rounded up to the next tenth) assigned to and placed on the label of a fissile material package to designate the degree of control of accumulation of packages containing fissile material during transportation.
	CriticalSafetyIndex *CriticalSafetyIndexType `xml:"criticalSafetyIndex"`
	// An extension hook for attaching extension (non-core) classes.
	Extension []base.RadioactiveMaterialExtensionType `xml:"extension"`
	// A figure representing the radiation level measured at one meter from the package.
	TransportIndex *TransportIndexType `xml:"transportIndex"`
}

