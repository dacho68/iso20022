package model

// Information that locates and identifies a specific address, as defined by postal services or in free format text.
type PostalAddress17 struct {

	// Information that locates and identifies a specific address, as defined by postal services, that is presented in free format text.
	AddressLine []*Max70Text `xml:"AdrLine,omitempty"`

	// Name of a street or thoroughfare.
	StreetName *Max70Text `xml:"StrtNm,omitempty"`

	// Number that identifies the position of a building on a street.
	BuildingNumber *Max16Text `xml:"BldgNb,omitempty"`

	// Identifier consisting of a group of letters and/or numbers that is added to a postal address to assist the sorting of mail.
	PostCode *Max16Text `xml:"PstCd,omitempty"`

	// Name of a built-up area, with defined boundaries, and a local government.
	TownName *Max35Text `xml:"TwnNm"`

	// Identifies a subdivision of a country, for instance state, region, county.
	CountrySubDivision []*Max35Text `xml:"CtrySubDvsn,omitempty"`

	// Nation with its own government, occupying a particular territory.
	Country *CountryCode `xml:"Ctry"`

	// Geographic location of the ATM specified by geographic coordinates or UTM coordinates.
	GeoLocation *GeographicLocation1Choice `xml:"GLctn,omitempty"`
}

func (p *PostalAddress17) AddAddressLine(value string) {
	p.AddressLine = append(p.AddressLine, (*Max70Text)(&value))
}

func (p *PostalAddress17) SetStreetName(value string) {
	p.StreetName = (*Max70Text)(&value)
}

func (p *PostalAddress17) SetBuildingNumber(value string) {
	p.BuildingNumber = (*Max16Text)(&value)
}

func (p *PostalAddress17) SetPostCode(value string) {
	p.PostCode = (*Max16Text)(&value)
}

func (p *PostalAddress17) SetTownName(value string) {
	p.TownName = (*Max35Text)(&value)
}

func (p *PostalAddress17) AddCountrySubDivision(value string) {
	p.CountrySubDivision = append(p.CountrySubDivision, (*Max35Text)(&value))
}

func (p *PostalAddress17) SetCountry(value string) {
	p.Country = (*CountryCode)(&value)
}

func (p *PostalAddress17) AddGeoLocation() *GeographicLocation1Choice {
	p.GeoLocation = new(GeographicLocation1Choice)
	return p.GeoLocation
}
