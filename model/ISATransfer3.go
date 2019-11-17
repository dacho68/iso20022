package model

// Describes the type of product and the assets to be transferred.
type ISATransfer3 struct {

	// Unique and unambiguous identifier for a group of individual transfers as assigned by the instructing party. This identifier links the individual transfers together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Identification assigned by the new plan manager to each transfer of asset.
	TransferIdentification *Max35Text `xml:"TrfId"`

	// Unique and unambiguous identifier for a transfer cancellation, as assigned by the instructing party.
	CancellationReference *Max35Text `xml:"CxlRef,omitempty"`

	// Indicates whether there is cash in the account that is awaiting investment.
	ResidualCash *ResidualCash1Code `xml:"RsdlCsh,omitempty"`

	// Specifies portfolio information or government schemes, for example Individual Savings Account (ISA) in the UK.
	Portfolio *ISAPortfolio1Choice `xml:"Prtfl"`

	// Indicator that all remaining assets in a portfolio not listed for transfer should be liquidated and transferred as cash.
	AllOtherCash *YesNoIndicator `xml:"AllOthrCsh"`

	// Specifies the underlying assets for the ISA or portfolio.
	FinancialInstrumentAssetForTransfer []*FinancialInstrument25 `xml:"FinInstrmAsstForTrf,omitempty"`
}

func (i *ISATransfer3) SetMasterReference(value string) {
	i.MasterReference = (*Max35Text)(&value)
}

func (i *ISATransfer3) SetTransferIdentification(value string) {
	i.TransferIdentification = (*Max35Text)(&value)
}

func (i *ISATransfer3) SetCancellationReference(value string) {
	i.CancellationReference = (*Max35Text)(&value)
}

func (i *ISATransfer3) SetResidualCash(value string) {
	i.ResidualCash = (*ResidualCash1Code)(&value)
}

func (i *ISATransfer3) AddPortfolio() *ISAPortfolio1Choice {
	i.Portfolio = new(ISAPortfolio1Choice)
	return i.Portfolio
}

func (i *ISATransfer3) SetAllOtherCash(value string) {
	i.AllOtherCash = (*YesNoIndicator)(&value)
}

func (i *ISATransfer3) AddFinancialInstrumentAssetForTransfer() *FinancialInstrument25 {
	newValue := new(FinancialInstrument25)
	i.FinancialInstrumentAssetForTransfer = append(i.FinancialInstrumentAssetForTransfer, newValue)
	return newValue
}
