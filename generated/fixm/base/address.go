// Code generated from Address.xsd; DO NOT EDIT.

package base

//The Telecom Networks that can be used to address an organisation. [adapted from AIXM]
type TelecomNetworkTypeType string

const (
	TelecomNetworkTypeTypeAFTN TelecomNetworkTypeType = "AFTN" // The data interchange in the AFS is performed by the Aeronautical Fixed Telecommunications Network, AFTN. This is a message handling network running according to ICAO Standards documented in Annex 10 to the ICAO Convention.[AIXM]
	TelecomNetworkTypeTypeINTERNET TelecomNetworkTypeType = "INTERNET" // The Internet is a worldwide, publicly accessible series of interconnected computer networks that transmit data by packet switching using the standard Internet Protocol (IP). 
																	   // The usage of this value indicates that a URL will be provided for the linkage property in OnlineContact.
)

// Basis for e-mail and street addresses. [FIXM]
type TextAddressType CharacterStringType

// Address city. [FIXM]
type TextCityType TextNameType

// The country code as specified by ISO 3166. [FIXM]
type TextCountryCodeType CharacterStringType

// The country of the physical address for the location or organisation.  Full name, not ISO 3166 abbreviations. [AIXM 5.1]
type TextCountryNameType TextNameType

// A phone or facsimile number. [AIXM 5.1]
type TextPhoneType CharacterStringType

// Information required to enable contact with the responsible person and/or organisation.  This model is derived from ISO19115-2003:Geographic Information- Metadata. [AIXM 5.1]
type ContactInformationType struct {
	// Optional postal address of the contact.
	Address *PostalAddressType `xml:"address"`
	// An extension hook for attaching extension (non-core) classes.
	Extension []ContactInformationExtensionType `xml:"extension"`
	// The official name of the contact. In case of the organization use, it is the name of the person within the organization who is the contact point.
	// If the usage of  ContactInformation is associated with a person, this field should not be used, instead, the Person class name should be used instead.
	Name *TextNameType `xml:"name"`
	// Optional e-mail address of the contact.[FIXM]
	OnlineContact []OnlineContactType `xml:"onlineContact"`
	// Optional phone or facsimile number of the contact.
	PhoneFax *TelephoneContactType `xml:"phoneFax"`
	// The official title of the contact.
	Title *TextNameType `xml:"title"`
}

// The Telecom Networks that can be used to address an organisation. [adapted from AIXM]
type NetworkChoiceType struct {
	// Used to specify network types not included in the TelecomNetworkType list.
	Other CharacterStringType `xml:"other"`
	// Type of telecom network used.
	Type TelecomNetworkTypeType `xml:"type"`
}

// On-line or Network information that can be used to contact the individual or organisation, including eMail address. The contact information can include an email address and either an AFTN address or URL, but usually not both. The network attribute specifies whether the address specified by the linkage attribute is an AFTN or an internet address. [specialised from AIXM 5.1]
type OnlineContactType struct {
	// The address of the electronic mailbox of the responsible organisation or individual. [FIXM]
	Email *TextAddressType `xml:"email"`
	// An extension hook for attaching extension (non-core) classes.
	Extension []OnlineContactExtensionType `xml:"extension"`
	// Location (address) for on-line access using a Uniform Resource Locator address or AFTN address. [adapted from AIXM]
	Linkage *TextAddressType `xml:"linkage"`
	// The Telecom Networks that can be used to address an organisation. [adapted from AIXM]
	Network *NetworkChoiceType `xml:"network"`
}

// Physical address at which the organization or individual may be contacted. Derived from ISO19115-2003
type PostalAddressType struct {
	// The state or province of the location or organisation.
	AdministrativeArea *TextNameType `xml:"administrativeArea"`
	// The city of the location or organisation.
	City *TextCityType `xml:"city"`
	// The country of the physical address for the location or organisation.  ISO 3166 abbreviations.
	CountryCode *TextCountryCodeType `xml:"countryCode"`
	// The country of the physical address for the location or organisation.  Full name, not ISO 3166 abbreviations.
	CountryName *TextCountryNameType `xml:"countryName"`
	// The street address line for the location.  More than one address line may be used. [FIXM]
	DeliveryPoint *TextAddressType `xml:"deliveryPoint"`
	// An extension hook for attaching extension (non-core) classes.
	Extension []PostalAddressExtensionType `xml:"extension"`
	// The ZIP or other postal code for the location or organisation.
	PostalCode *TextNameType `xml:"postalCode"`
}

// Telephone numbers at which the organisation or individual may be contacted.  From ISO19115-2003. [AIXM 5.1]
type TelephoneContactType struct {
	// An extension hook for attaching extension (non-core) classes.
	Extension []TelephoneContactExtensionType `xml:"extension"`
	// The telephone number of a facsimile machine for the responsible organisation or individual.
	Facsimile *TextPhoneType `xml:"facsimile"`
	// The telephone number by which individuals can speak to the responsible organisation or individual.
	Voice *TextPhoneType `xml:"voice"`
}

