package model

// Physical representation of a security.
type SecuritiesCertificate4 struct {

	// Unique and unambiguous identifier of a certificate assigned by the issuer.
	Number *Max35Text `xml:"Nb"`

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	Issuer *Max35Text `xml:"Issr,omitempty"`

	// Short textual description of the scheme.
	SchemeName *Max35Text `xml:"SchmeNm,omitempty"`
}

func (s *SecuritiesCertificate4) SetNumber(value string) {
	s.Number = (*Max35Text)(&value)
}

func (s *SecuritiesCertificate4) SetIssuer(value string) {
	s.Issuer = (*Max35Text)(&value)
}

func (s *SecuritiesCertificate4) SetSchemeName(value string) {
	s.SchemeName = (*Max35Text)(&value)
}
