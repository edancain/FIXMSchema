package flight/cargo

// CriticalSafetyIndexType is The dimensionless number (rounded up to the next tenth) assigned to and placed on the label of a fissile material package to designate the degree of control of accumulation of packages containing fissile material during transportation.
type CriticalSafetyIndexType interface{}

// RadioactiveMaterialCategoryType is Surface radiation >50 millirem/hr, 1 meter radiation >1 millirem/hr
type RadioactiveMaterialCategoryType string

// TransportIndexType is A figure representing the radiation level measured at one meter from the package.
type TransportIndexType interface{}

// RadioactiveMaterialType is A category used for radioactive materials in a package, overpack or freight container, based on their maximum radiation level.
type RadioactiveMaterialType struct {
	Category	string	`xml:"category"`
	CriticalSafetyIndex	interface{}	`xml:"criticalSafetyIndex"`
	Extension	[]*RadioactiveMaterialExtensionType	`xml:"extension"`
	TransportIndex	interface{}	`xml:"transportIndex"`
}
