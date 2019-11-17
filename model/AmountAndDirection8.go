package model

// Amount of money debited or credited on the books of an account servicer.
type AmountAndDirection8 struct {

	// Amount of money in the cash entry.
	Amount *ActiveCurrencyAndAmount `xml:"Amt"`

	// Indicates whether an entry is a credit or a debit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Posting/settlement amount in its original currency when conversion from/into another currency has occurred.
	OriginalCurrencyAndOrderedAmount *ActiveOrHistoricCurrencyAndAmount `xml:"OrgnlCcyAndOrdrdAmt,omitempty"`
}

func (a *AmountAndDirection8) SetAmount(value, currency string) {
	a.Amount = NewActiveCurrencyAndAmount(value, currency)
}

func (a *AmountAndDirection8) SetCreditDebitIndicator(value string) {
	a.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (a *AmountAndDirection8) SetOriginalCurrencyAndOrderedAmount(value, currency string) {
	a.OriginalCurrencyAndOrderedAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}
