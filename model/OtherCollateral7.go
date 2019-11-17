package model

// Provides details about the letter of credit or bank guarantee, or other collateral, posted as collateral.
type OtherCollateral7 struct {

	// Provides the identification of the proposed collateral.
	CollateralIdentification *Max35Text `xml:"CollId,omitempty"`

	// Identifies the register number of the collateral deposit assigned by the central counterparty.
	AssetNumber *Max35Text `xml:"AsstNb,omitempty"`

	// Provides the unique identification of the letter of credit.
	LetterOfCreditIdentification *Max35Text `xml:"LttrOfCdtId,omitempty"`

	// Amount of the letter/documentary credit.
	LetterOfCreditAmount *ActiveCurrencyAndAmount `xml:"LttrOfCdtAmt,omitempty"`

	// Amount of the bank guarantee.
	GuaranteeAmount *ActiveCurrencyAndAmount `xml:"GrntAmt,omitempty"`

	// Provides a description and an amount of another type of collateral.
	OtherTypeOfCollateral *OtherTypeOfCollateral2 `xml:"OthrTpOfColl,omitempty"`

	// Date on which the other collateral was issued.
	IssueDate *DateFormat14Choice `xml:"IsseDt,omitempty"`

	// Date on which the other collateral expires.
	ExpiryDate *DateFormat14Choice `xml:"XpryDt,omitempty"`

	// Indicates that the collateral deposited in the clearing house covers the margin until a specific timeframe.
	LimitedCoverageIndicator *YesNoIndicator `xml:"LtdCvrgInd,omitempty"`

	// Party that issues the bank guarantee or letter of / documentary credit.
	Issuer *PartyIdentification100Choice `xml:"Issr,omitempty"`

	// Valuation date of the other collateral when it was taken as collateral.
	ValueDate *ISODate `xml:"ValDt,omitempty"`

	// Exchange rate.
	ExchangeRate *BaseOneRate `xml:"XchgRate,omitempty"`

	// Value of the collateral based on current market prices.
	MarketValue *ActiveCurrencyAndAmount `xml:"MktVal,omitempty"`

	// Haircut or valuation factor on the collateral expressed as a percentage.
	Haircut *PercentageRate `xml:"Hrcut,omitempty"`

	// Value of the collateral after taking into account the haircut, if any.
	CollateralValue *ActiveCurrencyAndAmount `xml:"CollVal"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormat10Choice `xml:"SfkpgPlc,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount19 `xml:"SfkpgAcct,omitempty"`
}

func (o *OtherCollateral7) SetCollateralIdentification(value string) {
	o.CollateralIdentification = (*Max35Text)(&value)
}

func (o *OtherCollateral7) SetAssetNumber(value string) {
	o.AssetNumber = (*Max35Text)(&value)
}

func (o *OtherCollateral7) SetLetterOfCreditIdentification(value string) {
	o.LetterOfCreditIdentification = (*Max35Text)(&value)
}

func (o *OtherCollateral7) SetLetterOfCreditAmount(value, currency string) {
	o.LetterOfCreditAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (o *OtherCollateral7) SetGuaranteeAmount(value, currency string) {
	o.GuaranteeAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (o *OtherCollateral7) AddOtherTypeOfCollateral() *OtherTypeOfCollateral2 {
	o.OtherTypeOfCollateral = new(OtherTypeOfCollateral2)
	return o.OtherTypeOfCollateral
}

func (o *OtherCollateral7) AddIssueDate() *DateFormat14Choice {
	o.IssueDate = new(DateFormat14Choice)
	return o.IssueDate
}

func (o *OtherCollateral7) AddExpiryDate() *DateFormat14Choice {
	o.ExpiryDate = new(DateFormat14Choice)
	return o.ExpiryDate
}

func (o *OtherCollateral7) SetLimitedCoverageIndicator(value string) {
	o.LimitedCoverageIndicator = (*YesNoIndicator)(&value)
}

func (o *OtherCollateral7) AddIssuer() *PartyIdentification100Choice {
	o.Issuer = new(PartyIdentification100Choice)
	return o.Issuer
}

func (o *OtherCollateral7) SetValueDate(value string) {
	o.ValueDate = (*ISODate)(&value)
}

func (o *OtherCollateral7) SetExchangeRate(value string) {
	o.ExchangeRate = (*BaseOneRate)(&value)
}

func (o *OtherCollateral7) SetMarketValue(value, currency string) {
	o.MarketValue = NewActiveCurrencyAndAmount(value, currency)
}

func (o *OtherCollateral7) SetHaircut(value string) {
	o.Haircut = (*PercentageRate)(&value)
}

func (o *OtherCollateral7) SetCollateralValue(value, currency string) {
	o.CollateralValue = NewActiveCurrencyAndAmount(value, currency)
}

func (o *OtherCollateral7) AddSafekeepingPlace() *SafekeepingPlaceFormat10Choice {
	o.SafekeepingPlace = new(SafekeepingPlaceFormat10Choice)
	return o.SafekeepingPlace
}

func (o *OtherCollateral7) AddSafekeepingAccount() *SecuritiesAccount19 {
	o.SafekeepingAccount = new(SecuritiesAccount19)
	return o.SafekeepingAccount
}
