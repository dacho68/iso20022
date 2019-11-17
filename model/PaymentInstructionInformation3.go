package model

// Set of characteristics that apply to the debit side of the payment transactions included in the credit transfer initiation.
type PaymentInstructionInformation3 struct {

	// Unique identification, as assigned by a sending party, to unambiguously identify the payment information group within the message.
	PaymentInformationIdentification *Max35Text `xml:"PmtInfId"`

	// Specifies the means of payment that will be used to move the amount of money.
	PaymentMethod *PaymentMethod3Code `xml:"PmtMtd"`

	// Identifies whether a single entry per individual transaction or a batch entry for the sum of the amounts of all transactions within the group of a message is requested.
	// Usage: Batch booking is used to request and not order a possible batch booking.
	BatchBooking *BatchBookingIndicator `xml:"BtchBookg,omitempty"`

	// Number of individual transactions contained in the paymnet information group.
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
	Debtor *PartyIdentification32 `xml:"Dbtr"`

	// Unambiguous identification of the account of the debtor to which a debit entry will be made as a result of the transaction.
	DebtorAccount *CashAccount16 `xml:"DbtrAcct"`

	// Financial institution servicing an account for the debtor.
	DebtorAgent *BranchAndFinancialInstitutionIdentification4 `xml:"DbtrAgt"`

	// Unambiguous identification of the account of the debtor agent at its servicing agent in the payment chain.
	DebtorAgentAccount *CashAccount16 `xml:"DbtrAgtAcct,omitempty"`

	// Ultimate party that owes an amount of money to the (ultimate) creditor.
	UltimateDebtor *PartyIdentification32 `xml:"UltmtDbtr,omitempty"`

	// Specifies which party/parties will bear the charges associated with the processing of the payment transaction.
	ChargeBearer *ChargeBearerType1Code `xml:"ChrgBr,omitempty"`

	// Account used to process charges associated with a transaction.
	//
	// Usage: Charges account should be used when charges have to be booked to an account different from the account identified in debtor's account.
	ChargesAccount *CashAccount16 `xml:"ChrgsAcct,omitempty"`

	// Agent that services a charges account.
	//
	// Usage: Charges account agent should only be used when the charges account agent is different from the debtor agent.
	ChargesAccountAgent *BranchAndFinancialInstitutionIdentification4 `xml:"ChrgsAcctAgt,omitempty"`

	// Set of elements used to provide information on the individual transaction(s) included in the message.
	CreditTransferTransactionInformation []*CreditTransferTransactionInformation10 `xml:"CdtTrfTxInf"`
}

func (p *PaymentInstructionInformation3) SetPaymentInformationIdentification(value string) {
	p.PaymentInformationIdentification = (*Max35Text)(&value)
}

func (p *PaymentInstructionInformation3) SetPaymentMethod(value string) {
	p.PaymentMethod = (*PaymentMethod3Code)(&value)
}

func (p *PaymentInstructionInformation3) SetBatchBooking(value string) {
	p.BatchBooking = (*BatchBookingIndicator)(&value)
}

func (p *PaymentInstructionInformation3) SetNumberOfTransactions(value string) {
	p.NumberOfTransactions = (*Max15NumericText)(&value)
}

func (p *PaymentInstructionInformation3) SetControlSum(value string) {
	p.ControlSum = (*DecimalNumber)(&value)
}

func (p *PaymentInstructionInformation3) AddPaymentTypeInformation() *PaymentTypeInformation19 {
	p.PaymentTypeInformation = new(PaymentTypeInformation19)
	return p.PaymentTypeInformation
}

func (p *PaymentInstructionInformation3) SetRequestedExecutionDate(value string) {
	p.RequestedExecutionDate = (*ISODate)(&value)
}

func (p *PaymentInstructionInformation3) SetPoolingAdjustmentDate(value string) {
	p.PoolingAdjustmentDate = (*ISODate)(&value)
}

func (p *PaymentInstructionInformation3) AddDebtor() *PartyIdentification32 {
	p.Debtor = new(PartyIdentification32)
	return p.Debtor
}

func (p *PaymentInstructionInformation3) AddDebtorAccount() *CashAccount16 {
	p.DebtorAccount = new(CashAccount16)
	return p.DebtorAccount
}

func (p *PaymentInstructionInformation3) AddDebtorAgent() *BranchAndFinancialInstitutionIdentification4 {
	p.DebtorAgent = new(BranchAndFinancialInstitutionIdentification4)
	return p.DebtorAgent
}

func (p *PaymentInstructionInformation3) AddDebtorAgentAccount() *CashAccount16 {
	p.DebtorAgentAccount = new(CashAccount16)
	return p.DebtorAgentAccount
}

func (p *PaymentInstructionInformation3) AddUltimateDebtor() *PartyIdentification32 {
	p.UltimateDebtor = new(PartyIdentification32)
	return p.UltimateDebtor
}

func (p *PaymentInstructionInformation3) SetChargeBearer(value string) {
	p.ChargeBearer = (*ChargeBearerType1Code)(&value)
}

func (p *PaymentInstructionInformation3) AddChargesAccount() *CashAccount16 {
	p.ChargesAccount = new(CashAccount16)
	return p.ChargesAccount
}

func (p *PaymentInstructionInformation3) AddChargesAccountAgent() *BranchAndFinancialInstitutionIdentification4 {
	p.ChargesAccountAgent = new(BranchAndFinancialInstitutionIdentification4)
	return p.ChargesAccountAgent
}

func (p *PaymentInstructionInformation3) AddCreditTransferTransactionInformation() *CreditTransferTransactionInformation10 {
	newValue := new(CreditTransferTransactionInformation10)
	p.CreditTransferTransactionInformation = append(p.CreditTransferTransactionInformation, newValue)
	return newValue
}
