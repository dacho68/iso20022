package model

// Characteristics that apply to the debit side of the payment transactions included in the credit transfer initiation.
type PaymentInstruction16 struct {

	// Unique identification, as assigned by a sending party, to unambiguously identify the payment information group within the message.
	PaymentInformationIdentification *Max35Text `xml:"PmtInfId"`

	// Specifies the means of payment that will be used to move the amount of money.
	PaymentMethod *PaymentMethod3Code `xml:"PmtMtd"`

	// Identifies whether a single entry per individual transaction or a batch entry for the sum of the amounts of all transactions within the group of a message is requested.
	// Usage: Batch booking is used to request and not order a possible batch booking.
	BatchBooking *BatchBookingIndicator `xml:"BtchBookg,omitempty"`

	// Number of individual transactions contained in the payment information group.
	NumberOfTransactions *Max15NumericText `xml:"NbOfTxs,omitempty"`

	// Total of all individual amounts included in the group, irrespective of currencies.
	ControlSum *DecimalNumber `xml:"CtrlSum,omitempty"`

	// Set of elements used to further specify the type of transaction.
	PaymentTypeInformation *PaymentTypeInformation19 `xml:"PmtTpInf,omitempty"`

	// Date at which the initiating party requests the clearing agent to process the payment.
	// Usage: This is the date on which the debtor's account is to be debited. If payment by cheque, the date when the cheque must be generated by the bank.
	RequestedExecutionDate *ISODate `xml:"ReqdExctnDt"`

	// Date used for the correction of the value date of a cash pool movement that has been posted with a different value date.
	PoolingAdjustmentDate *ISODate `xml:"PoolgAdjstmntDt,omitempty"`

	// Party that owes an amount of money to the (ultimate) creditor.
	Debtor *PartyIdentification43 `xml:"Dbtr"`

	// Unambiguous identification of the account of the debtor to which a debit entry will be made as a result of the transaction.
	DebtorAccount *CashAccount24 `xml:"DbtrAcct"`

	// Financial institution servicing an account for the debtor.
	DebtorAgent *BranchAndFinancialInstitutionIdentification5 `xml:"DbtrAgt"`

	// Unambiguous identification of the account of the debtor agent at its servicing agent in the payment chain.
	DebtorAgentAccount *CashAccount24 `xml:"DbtrAgtAcct,omitempty"`

	// Further information related to the processing of the payment instruction, that may need to be acted upon by the debtor agent, depending on agreement between debtor and the debtor agent.
	//
	// Usage: when present, then the instructions for the debtor agent apply for all credit transfer transaction information occurrences, present in the payment information.
	InstructionForDebtorAgent *Max140Text `xml:"InstrForDbtrAgt,omitempty"`

	// Ultimate party that owes an amount of money to the (ultimate) creditor.
	UltimateDebtor *PartyIdentification43 `xml:"UltmtDbtr,omitempty"`

	// Specifies which party/parties will bear the charges associated with the processing of the payment transaction.
	ChargeBearer *ChargeBearerType1Code `xml:"ChrgBr,omitempty"`

	// Account used to process charges associated with a transaction.
	//
	// Usage: Charges account should be used when charges have to be booked to an account different from the account identified in debtor's account.
	ChargesAccount *CashAccount24 `xml:"ChrgsAcct,omitempty"`

	// Agent that services a charges account.
	//
	// Usage: Charges account agent should only be used when the charges account agent is different from the debtor agent.
	ChargesAccountAgent *BranchAndFinancialInstitutionIdentification5 `xml:"ChrgsAcctAgt,omitempty"`

	// Provides information on the individual transaction(s) included in the message.
	CreditTransferTransactionInformation []*CreditTransferTransaction20 `xml:"CdtTrfTxInf"`
}

func (p *PaymentInstruction16) SetPaymentInformationIdentification(value string) {
	p.PaymentInformationIdentification = (*Max35Text)(&value)
}

func (p *PaymentInstruction16) SetPaymentMethod(value string) {
	p.PaymentMethod = (*PaymentMethod3Code)(&value)
}

func (p *PaymentInstruction16) SetBatchBooking(value string) {
	p.BatchBooking = (*BatchBookingIndicator)(&value)
}

func (p *PaymentInstruction16) SetNumberOfTransactions(value string) {
	p.NumberOfTransactions = (*Max15NumericText)(&value)
}

func (p *PaymentInstruction16) SetControlSum(value string) {
	p.ControlSum = (*DecimalNumber)(&value)
}

func (p *PaymentInstruction16) AddPaymentTypeInformation() *PaymentTypeInformation19 {
	p.PaymentTypeInformation = new(PaymentTypeInformation19)
	return p.PaymentTypeInformation
}

func (p *PaymentInstruction16) SetRequestedExecutionDate(value string) {
	p.RequestedExecutionDate = (*ISODate)(&value)
}

func (p *PaymentInstruction16) SetPoolingAdjustmentDate(value string) {
	p.PoolingAdjustmentDate = (*ISODate)(&value)
}

func (p *PaymentInstruction16) AddDebtor() *PartyIdentification43 {
	p.Debtor = new(PartyIdentification43)
	return p.Debtor
}

func (p *PaymentInstruction16) AddDebtorAccount() *CashAccount24 {
	p.DebtorAccount = new(CashAccount24)
	return p.DebtorAccount
}

func (p *PaymentInstruction16) AddDebtorAgent() *BranchAndFinancialInstitutionIdentification5 {
	p.DebtorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return p.DebtorAgent
}

func (p *PaymentInstruction16) AddDebtorAgentAccount() *CashAccount24 {
	p.DebtorAgentAccount = new(CashAccount24)
	return p.DebtorAgentAccount
}

func (p *PaymentInstruction16) SetInstructionForDebtorAgent(value string) {
	p.InstructionForDebtorAgent = (*Max140Text)(&value)
}

func (p *PaymentInstruction16) AddUltimateDebtor() *PartyIdentification43 {
	p.UltimateDebtor = new(PartyIdentification43)
	return p.UltimateDebtor
}

func (p *PaymentInstruction16) SetChargeBearer(value string) {
	p.ChargeBearer = (*ChargeBearerType1Code)(&value)
}

func (p *PaymentInstruction16) AddChargesAccount() *CashAccount24 {
	p.ChargesAccount = new(CashAccount24)
	return p.ChargesAccount
}

func (p *PaymentInstruction16) AddChargesAccountAgent() *BranchAndFinancialInstitutionIdentification5 {
	p.ChargesAccountAgent = new(BranchAndFinancialInstitutionIdentification5)
	return p.ChargesAccountAgent
}

func (p *PaymentInstruction16) AddCreditTransferTransactionInformation() *CreditTransferTransaction20 {
	newValue := new(CreditTransferTransaction20)
	p.CreditTransferTransactionInformation = append(p.CreditTransferTransactionInformation, newValue)
	return newValue
}
