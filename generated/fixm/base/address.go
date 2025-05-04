// Code generated from Address.xsd; DO NOT EDIT.

package base

// The Telecom Networks that can be used to address an organisation. [adapted fr...
type TelecomNetworkTypeType string

const (
	TelecomNetworkTypeTypeAFTN TelecomNetworkTypeType = "AFTN"
	TelecomNetworkTypeTypeINTERNET TelecomNetworkTypeType = "INTERNET"
)

// Basis for e-mail and street addresses. [FIXM]
type TextAddressType CharacterStringType

// Address city. [FIXM]
type TextCityType TextNameType

// The country code as specified by ISO 3166. [FIXM]
type TextCountryCodeType CharacterStringType

// The country of the physical address for the location or organisation. Full na...
type TextCountryNameType TextNameType

// A phone or facsimile number. [AIXM 5.1]
type TextPhoneType CharacterStringType

// Information required to enable contact with the responsible person and/or org...
type ContactInformationType struct {
	// Optional postal address of the contact.
	Address *PostalAddressType `xml:"address"`
	// An extension hook for attaching extension (non-core) classes.
	Extension []ContactInformationExtensionType `xml:"extension"`
	// The official name of the contact. In case of the organization use, it is the ...
	Name *TextNameType `xml:"name"`
	// Optional e-mail address of the contact.[FIXM]
	OnlineContact []OnlineContactType `xml:"onlineContact"`
	// Optional phone or facsimile number of the contact.
	PhoneFax *TelephoneContactType `xml:"phoneFax"`
	// The official title of the contact.
	Title *TextNameType `xml:"title"`
}

// The Telecom Networks that can be used to address an organisation. [adapted fr...
type NetworkChoiceType struct {
	// Used to specify network types not included in the TelecomNetworkType list.
	Other CharacterStringType `xml:"other"`
	// Type of telecom network used.
	Type TelecomNetworkTypeType `xml:"type"`
}

// On-line or Network information that can be used to contact the individual or ...
type OnlineContactType struct {
	// The address of the electronic mailbox of the responsible organisation or indi...
	Email *TextAddressType `xml:"email"`
	// An extension hook for attaching extension (non-core) classes.
	Extension []OnlineContactExtensionType `xml:"extension"`
	// Location (address) for on-line access using a Uniform Resource Locator addres...
	Linkage *TextAddressType `xml:"linkage"`
	// The Telecom Networks that can be used to address an organisation. [adapted fr...
	Network *NetworkChoiceType `xml:"network"`
}

// Physical address at which the organization or individual may be contacted. De...
type PostalAddressType struct {
	// The state or province of the location or organisation.
	AdministrativeArea *TextNameType `xml:"administrativeArea"`
	// The city of the location or organisation.
	City *TextCityType `xml:"city"`
	// The country of the physical address for the location or organisation. ISO 316...
	CountryCode *TextCountryCodeType `xml:"countryCode"`
	// The country of the physical address for the location or organisation. Full na...
	CountryName *TextCountryNameType `xml:"countryName"`
	// The street address line for the location. More than one address line may be u...
	DeliveryPoint *TextAddressType `xml:"deliveryPoint"`
	// An extension hook for attaching extension (non-core) classes.
	Extension []PostalAddressExtensionType `xml:"extension"`
	// The ZIP or other postal code for the location or organisation.
	PostalCode *TextNameType `xml:"postalCode"`
}

// Telephone numbers at which the organisation or individual may be contacted. F...
type TelephoneContactType struct {
	// An extension hook for attaching extension (non-core) classes.
	Extension []TelephoneContactExtensionType `xml:"extension"`
	// The telephone number of a facsimile machine for the responsible organisation ...
	Facsimile *TextPhoneType `xml:"facsimile"`
	// The telephone number by which individuals can speak to the responsible organi...
	Voice *TextPhoneType `xml:"voice"`
}

