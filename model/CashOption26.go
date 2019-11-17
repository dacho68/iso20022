package model

// Provides information about the cash option.
type CashOption26 struct {

	// Indicates whether the value is a debit or credit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Indicates whether the cash payment occurs or will occur in advance of receipt of proceeds from the issuer and based on a contractual agreement established with the account servicer or upon receipt of proceeds from the issuer.
	ContractualPaymentIndicator *Payment1Code `xml:"CtrctlPmtInd,omitempty"`

	// Specifies the type of income.
	// The lists of income type codes to be used, are available on the SMPG website at www.smpg.info.
	IncomeType *GenericIdentification20 `xml:"IncmTp,omitempty"`

	// Choice between a cash account, a charges account or a tax account.
	Account *Account8Choice `xml:"Acct,omitempty"`

	// Provides information about cash parties.
	CashParties *CashParties10 `xml:"CshPties,omitempty"`

	// Provides information about the amounts related to a cash movement.
	AmountDetails *CorporateActionAmounts23 `xml:"AmtDtls"`

	// Provides information about the dates related to a cash movement.
	DateDetails *CorporateActionDate24 `xml:"DtDtls"`

	// Exchange rate between the amount and the resulting amount.
	ForeignExchangeDetails *ForeignExchangeTerms11 `xml:"FXDtls,omitempty"`

	// Provides information about the tax voucher related to a cash movement.
	TaxVoucherDetails *TaxVoucher2 `xml:"TaxVchrDtls,omitempty"`

	// Provides information about the corporate action option.
	RateAndAmountDetails *RateDetails11 `xml:"RateAndAmtDtls,omitempty"`

	// Provides information about the prices related to a corporate action option.
	PriceDetails *PriceDetails11 `xml:"PricDtls,omitempty"`
}

func (c *CashOption26) SetCreditDebitIndicator(value string) {
	c.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (c *CashOption26) SetContractualPaymentIndicator(value string) {
	c.ContractualPaymentIndicator = (*Payment1Code)(&value)
}

func (c *CashOption26) AddIncomeType() *GenericIdentification20 {
	c.IncomeType = new(GenericIdentification20)
	return c.IncomeType
}

func (c *CashOption26) AddAccount() *Account8Choice {
	c.Account = new(Account8Choice)
	return c.Account
}

func (c *CashOption26) AddCashParties() *CashParties10 {
	c.CashParties = new(CashParties10)
	return c.CashParties
}

func (c *CashOption26) AddAmountDetails() *CorporateActionAmounts23 {
	c.AmountDetails = new(CorporateActionAmounts23)
	return c.AmountDetails
}

func (c *CashOption26) AddDateDetails() *CorporateActionDate24 {
	c.DateDetails = new(CorporateActionDate24)
	return c.DateDetails
}

func (c *CashOption26) AddForeignExchangeDetails() *ForeignExchangeTerms11 {
	c.ForeignExchangeDetails = new(ForeignExchangeTerms11)
	return c.ForeignExchangeDetails
}

func (c *CashOption26) AddTaxVoucherDetails() *TaxVoucher2 {
	c.TaxVoucherDetails = new(TaxVoucher2)
	return c.TaxVoucherDetails
}

func (c *CashOption26) AddRateAndAmountDetails() *RateDetails11 {
	c.RateAndAmountDetails = new(RateDetails11)
	return c.RateAndAmountDetails
}

func (c *CashOption26) AddPriceDetails() *PriceDetails11 {
	c.PriceDetails = new(PriceDetails11)
	return c.PriceDetails
}
