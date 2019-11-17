package model

// Human entity, as distinguished from a corporate entity (which is sometimes referred to as an 'artificial person').
type IndividualPerson11 struct {

	// Specifies the terms used to formally address a person.
	NamePrefix *NamePrefix1Code `xml:"NmPrfx,omitempty"`

	// Specifies the terms used to formally address a person.
	ExtendedNamePrefix *Extended350Code `xml:"XtndedNmPrfx,omitempty"`

	// First name of a person.
	GivenName *Max35Text `xml:"GvnNm"`

	// Second name of a person.
	MiddleName *Max35Text `xml:"MddlNm,omitempty"`

	// Name by which a party is known and which is usually used to identify that party.
	Name *Max350Text `xml:"Nm"`

	// Additional information about a person that follows a person's name, eg, qualification such as Doctor of Philosophy (PhD).
	NameSuffix *Max35Text `xml:"NmSfx,omitempty"`

	// Specifies the gender of the person.
	Gender *GenderCode `xml:"Gndr,omitempty"`

	// Language in which a person communicates.
	Language *LanguageCode `xml:"Lang,omitempty"`

	// Date on which a person is born.
	BirthDate *ISODate `xml:"BirthDt"`

	// Country where a person was born.
	CountryOfBirth *CountryCode `xml:"CtryOfBirth,omitempty"`

	// Province where a person was born.
	ProvinceOfBirth *Max35Text `xml:"PrvcOfBirth,omitempty"`

	// City where a person was born.
	CityOfBirth *Max35Text `xml:"CityOfBirth,omitempty"`

	// Name of the occupation or job of a person.
	Profession *Max35Text `xml:"Prfssn,omitempty"`

	// Country of taxation of an individual person or an organisation.
	TaxationCountry *CountryCode `xml:"TaxtnCtry,omitempty"`

	// Country and residential status of an individual, for example, non-pernament resident.
	CountryAndResidentialStatus *CountryAndResidentialStatusType1 `xml:"CtryAndResdtlSts,omitempty"`

	// Title of the function.
	BusinessFunction *Max35Text `xml:"BizFctn,omitempty"`

	// Organisation represented by a person, or for which a person works.
	EmployingCompany *Max140Text `xml:"EmplngCpny,omitempty"`

	// Address information to be either inserted, updated or deleted.
	ModifiedPostalAddress []*ModificationScope1 `xml:"ModfdPstlAdr,omitempty"`

	// Citizenship information to be inserted or deleted.
	ModifiedCitizenship []*ModificationScope3 `xml:"ModfdCtznsh,omitempty"`

	// Communication device number or electronic address used for communication.
	PrimaryCommunicationAddress *CommunicationAddress3 `xml:"PmryComAdr,omitempty"`

	// Communication device number or electronic address used for communication.
	SecondaryCommunicationAddress *CommunicationAddress3 `xml:"ScndryComAdr,omitempty"`

	// Identification information to be either inserted or deleted.
	ModifiedOtherIdentification []*ModificationScope2 `xml:"ModfdOthrId,omitempty"`
}

func (i *IndividualPerson11) SetNamePrefix(value string) {
	i.NamePrefix = (*NamePrefix1Code)(&value)
}

func (i *IndividualPerson11) SetExtendedNamePrefix(value string) {
	i.ExtendedNamePrefix = (*Extended350Code)(&value)
}

func (i *IndividualPerson11) SetGivenName(value string) {
	i.GivenName = (*Max35Text)(&value)
}

func (i *IndividualPerson11) SetMiddleName(value string) {
	i.MiddleName = (*Max35Text)(&value)
}

func (i *IndividualPerson11) SetName(value string) {
	i.Name = (*Max350Text)(&value)
}

func (i *IndividualPerson11) SetNameSuffix(value string) {
	i.NameSuffix = (*Max35Text)(&value)
}

func (i *IndividualPerson11) SetGender(value string) {
	i.Gender = (*GenderCode)(&value)
}

func (i *IndividualPerson11) SetLanguage(value string) {
	i.Language = (*LanguageCode)(&value)
}

func (i *IndividualPerson11) SetBirthDate(value string) {
	i.BirthDate = (*ISODate)(&value)
}

func (i *IndividualPerson11) SetCountryOfBirth(value string) {
	i.CountryOfBirth = (*CountryCode)(&value)
}

func (i *IndividualPerson11) SetProvinceOfBirth(value string) {
	i.ProvinceOfBirth = (*Max35Text)(&value)
}

func (i *IndividualPerson11) SetCityOfBirth(value string) {
	i.CityOfBirth = (*Max35Text)(&value)
}

func (i *IndividualPerson11) SetProfession(value string) {
	i.Profession = (*Max35Text)(&value)
}

func (i *IndividualPerson11) SetTaxationCountry(value string) {
	i.TaxationCountry = (*CountryCode)(&value)
}

func (i *IndividualPerson11) AddCountryAndResidentialStatus() *CountryAndResidentialStatusType1 {
	i.CountryAndResidentialStatus = new(CountryAndResidentialStatusType1)
	return i.CountryAndResidentialStatus
}

func (i *IndividualPerson11) SetBusinessFunction(value string) {
	i.BusinessFunction = (*Max35Text)(&value)
}

func (i *IndividualPerson11) SetEmployingCompany(value string) {
	i.EmployingCompany = (*Max140Text)(&value)
}

func (i *IndividualPerson11) AddModifiedPostalAddress() *ModificationScope1 {
	newValue := new(ModificationScope1)
	i.ModifiedPostalAddress = append(i.ModifiedPostalAddress, newValue)
	return newValue
}

func (i *IndividualPerson11) AddModifiedCitizenship() *ModificationScope3 {
	newValue := new(ModificationScope3)
	i.ModifiedCitizenship = append(i.ModifiedCitizenship, newValue)
	return newValue
}

func (i *IndividualPerson11) AddPrimaryCommunicationAddress() *CommunicationAddress3 {
	i.PrimaryCommunicationAddress = new(CommunicationAddress3)
	return i.PrimaryCommunicationAddress
}

func (i *IndividualPerson11) AddSecondaryCommunicationAddress() *CommunicationAddress3 {
	i.SecondaryCommunicationAddress = new(CommunicationAddress3)
	return i.SecondaryCommunicationAddress
}

func (i *IndividualPerson11) AddModifiedOtherIdentification() *ModificationScope2 {
	newValue := new(ModificationScope2)
	i.ModifiedOtherIdentification = append(i.ModifiedOtherIdentification, newValue)
	return newValue
}
