package model

// Provides information about the securities account.
type SecuritiesAccount9 struct {

	// Specifies whether the value is a debit or credit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Identification of the party that owns the account.
	AccountOwnerIdentification *PartyIdentification2Choice `xml:"AcctOwnrId,omitempty"`

	// Idenfitication of the account.
	AccountIdentification *Max35Text `xml:"AcctId"`

	// Type of balance.
	BalanceType *SecuritiesBalanceType10FormatChoice `xml:"BalTp,omitempty"`

	// Specifies the corporate action options available to the account owner.
	OptionType *CorporateActionOption1FormatChoice `xml:"OptnTp,omitempty"`

	// Number identifying the available corporate action options.
	OptionNumber *Exact3NumericText `xml:"OptnNb,omitempty"`

	// Specifies the form of the financial instrument.
	SecurityHoldingForm *FormOfSecurity1Code `xml:"SctyHldgForm,omitempty"`
}

func (s *SecuritiesAccount9) SetCreditDebitIndicator(value string) {
	s.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (s *SecuritiesAccount9) AddAccountOwnerIdentification() *PartyIdentification2Choice {
	s.AccountOwnerIdentification = new(PartyIdentification2Choice)
	return s.AccountOwnerIdentification
}

func (s *SecuritiesAccount9) SetAccountIdentification(value string) {
	s.AccountIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesAccount9) AddBalanceType() *SecuritiesBalanceType10FormatChoice {
	s.BalanceType = new(SecuritiesBalanceType10FormatChoice)
	return s.BalanceType
}

func (s *SecuritiesAccount9) AddOptionType() *CorporateActionOption1FormatChoice {
	s.OptionType = new(CorporateActionOption1FormatChoice)
	return s.OptionType
}

func (s *SecuritiesAccount9) SetOptionNumber(value string) {
	s.OptionNumber = (*Exact3NumericText)(&value)
}

func (s *SecuritiesAccount9) SetSecurityHoldingForm(value string) {
	s.SecurityHoldingForm = (*FormOfSecurity1Code)(&value)
}
