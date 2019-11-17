package model

// Provide further details on the requested modifications of the underlying payment instruction.
type RequestedModification6 struct {

	// Unique identification as assigned by an instructing party for an instructed party to unambiguously identify the instruction.
	//
	// Usage: The instruction identification is a point to point reference that can be used between the instructing party and the instructed party to refer to the individual instruction. It can be included in several messages related to the instruction.
	InstructionIdentification *Max35Text `xml:"InstrId,omitempty"`

	// Unique identification, as assigned by the initiating party, to unambiguously identify the transaction. This identification is passed on, unchanged, throughout the entire end-to-end chain.
	//
	// Usage: The end-to-end identification can be used for reconciliation or to link tasks relating to the transaction. It can be included in several messages related to the transaction.
	EndToEndIdentification *Max35Text `xml:"EndToEndId,omitempty"`

	// Unique identification, as assigned by the first instructing agent, to unambiguously identify the transaction that is passed on, unchanged, throughout the entire interbank chain.
	// Usage: The transaction identification can be used for reconciliation, tracking or to link tasks relating to the transaction on the interbank level. The instructing agent has to make sure that the transaction identification is unique for a pre-agreed period.
	TransactionIdentification *Max35Text `xml:"TxId,omitempty"`

	// Set of elements used to further specify the type of transaction.
	PaymentTypeInformation *PaymentTypeInformation25 `xml:"PmtTpInf,omitempty"`

	// Date at which the initiating party requests the clearing agent to process the payment.
	// Usage: This is the date on which the debtor's account is to be debited. If payment by cheque, the date when the cheque must be generated by the bank.
	RequestedExecutionDate *DateAndDateTimeChoice `xml:"ReqdExctnDt,omitempty"`

	// Date and time at which the creditor requests that the amount of money is to be collected from the debtor.
	RequestedCollectionDate *ISODate `xml:"ReqdColltnDt,omitempty"`

	// Date on which the amount of money ceases to be available to the agent that owes it and when the amount of money becomes available to the agent to which it is due.
	InterbankSettlementDate *ISODate `xml:"IntrBkSttlmDt,omitempty"`

	// Amount of money to be moved between the debtor and creditor, before deduction of charges, expressed in the currency as ordered by the initiating party.
	Amount *AmountType4Choice `xml:"Amt,omitempty"`

	// Amount of money moved between the instructing agent and the instructed agent.
	InterbankSettlementAmount *ActiveOrHistoricCurrencyAndAmount `xml:"IntrBkSttlmAmt,omitempty"`

	// Specifies which party/parties will bear the charges associated with the processing of the payment transaction.
	ChargeBearer *ChargeBearerType1Code `xml:"ChrgBr,omitempty"`

	// Ultimate party that owes an amount of money to the (ultimate) creditor.
	UltimateDebtor *PartyIdentification43 `xml:"UltmtDbtr,omitempty"`

	// Party that owes an amount of money to the (ultimate) creditor. In the context of the payment model, the debtor is also the debit account owner.
	Debtor *PartyIdentification43 `xml:"Dbtr,omitempty"`

	// Unambiguous identification of the account of the debtor to which a debit entry will be made as a result of the transaction.
	DebtorAccount *CashAccount24 `xml:"DbtrAcct,omitempty"`

	// Unambiguous identification of the account of the debtor agent at its servicing agent in the payment chain.
	DebtorAgentAccount *CashAccount24 `xml:"DbtrAgtAcct,omitempty"`

	// Instruction between two clearing agents stipulating the cash transfer characteristics between the two parties.
	SettlementInformation *SettlementInstruction3 `xml:"SttlmInf,omitempty"`

	// Unambiguous identification of the account of the creditor agent at its servicing agent to which a credit entry will be made as a result of the payment transaction.
	CreditorAgentAccount *CashAccount24 `xml:"CdtrAgtAcct,omitempty"`

	// Party to which an amount of money is due.
	Creditor *PartyIdentification43 `xml:"Cdtr,omitempty"`

	// Unambiguous identification of the account of the creditor to which a credit entry will be posted as a result of the payment transaction.
	CreditorAccount *CashAccount24 `xml:"CdtrAcct,omitempty"`

	// Ultimate party to which an amount of money is due.
	UltimateCreditor *PartyIdentification43 `xml:"UltmtCdtr,omitempty"`

	// Underlying reason for the payment transaction.
	// Usage: Purpose is used by the end-customers, that is initiating party, (ultimate) debtor, (ultimate) creditor to provide information concerning the nature of the payment. Purpose is a content element, which is not used for processing by any of the agents involved in the payment chain.
	Purpose *Purpose2Choice `xml:"Purp,omitempty"`

	// Further information related to the processing of the payment instruction, provided by the initiating party, and intended for the debtor agent.
	InstructionForDebtorAgent *Max140Text `xml:"InstrForDbtrAgt,omitempty"`

	// Further information related to the processing of the payment instruction that may need to be acted upon by the next agent.
	//
	// Usage: The next agent may not be the creditor agent.
	// The instruction can relate to a level of service, can be an instruction that has to be executed by the agent, or can be information required by the next agent.
	InstructionForNextAgent []*InstructionForNextAgent1 `xml:"InstrForNxtAgt,omitempty"`

	// Further information related to the processing of the payment instruction, provided by the initiating party, and intended for the creditor agent.
	InstructionForCreditorAgent []*InstructionForCreditorAgent1 `xml:"InstrForCdtrAgt,omitempty"`

	// Structured information that enables the matching, that is reconciliation, of a payment with the items that the payment is intended to settle, such as commercial invoices in an account receivable system.
	RemittanceInformation *RemittanceInformation11 `xml:"RmtInf,omitempty"`
}

func (r *RequestedModification6) SetInstructionIdentification(value string) {
	r.InstructionIdentification = (*Max35Text)(&value)
}

func (r *RequestedModification6) SetEndToEndIdentification(value string) {
	r.EndToEndIdentification = (*Max35Text)(&value)
}

func (r *RequestedModification6) SetTransactionIdentification(value string) {
	r.TransactionIdentification = (*Max35Text)(&value)
}

func (r *RequestedModification6) AddPaymentTypeInformation() *PaymentTypeInformation25 {
	r.PaymentTypeInformation = new(PaymentTypeInformation25)
	return r.PaymentTypeInformation
}

func (r *RequestedModification6) AddRequestedExecutionDate() *DateAndDateTimeChoice {
	r.RequestedExecutionDate = new(DateAndDateTimeChoice)
	return r.RequestedExecutionDate
}

func (r *RequestedModification6) SetRequestedCollectionDate(value string) {
	r.RequestedCollectionDate = (*ISODate)(&value)
}

func (r *RequestedModification6) SetInterbankSettlementDate(value string) {
	r.InterbankSettlementDate = (*ISODate)(&value)
}

func (r *RequestedModification6) AddAmount() *AmountType4Choice {
	r.Amount = new(AmountType4Choice)
	return r.Amount
}

func (r *RequestedModification6) SetInterbankSettlementAmount(value, currency string) {
	r.InterbankSettlementAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (r *RequestedModification6) SetChargeBearer(value string) {
	r.ChargeBearer = (*ChargeBearerType1Code)(&value)
}

func (r *RequestedModification6) AddUltimateDebtor() *PartyIdentification43 {
	r.UltimateDebtor = new(PartyIdentification43)
	return r.UltimateDebtor
}

func (r *RequestedModification6) AddDebtor() *PartyIdentification43 {
	r.Debtor = new(PartyIdentification43)
	return r.Debtor
}

func (r *RequestedModification6) AddDebtorAccount() *CashAccount24 {
	r.DebtorAccount = new(CashAccount24)
	return r.DebtorAccount
}

func (r *RequestedModification6) AddDebtorAgentAccount() *CashAccount24 {
	r.DebtorAgentAccount = new(CashAccount24)
	return r.DebtorAgentAccount
}

func (r *RequestedModification6) AddSettlementInformation() *SettlementInstruction3 {
	r.SettlementInformation = new(SettlementInstruction3)
	return r.SettlementInformation
}

func (r *RequestedModification6) AddCreditorAgentAccount() *CashAccount24 {
	r.CreditorAgentAccount = new(CashAccount24)
	return r.CreditorAgentAccount
}

func (r *RequestedModification6) AddCreditor() *PartyIdentification43 {
	r.Creditor = new(PartyIdentification43)
	return r.Creditor
}

func (r *RequestedModification6) AddCreditorAccount() *CashAccount24 {
	r.CreditorAccount = new(CashAccount24)
	return r.CreditorAccount
}

func (r *RequestedModification6) AddUltimateCreditor() *PartyIdentification43 {
	r.UltimateCreditor = new(PartyIdentification43)
	return r.UltimateCreditor
}

func (r *RequestedModification6) AddPurpose() *Purpose2Choice {
	r.Purpose = new(Purpose2Choice)
	return r.Purpose
}

func (r *RequestedModification6) SetInstructionForDebtorAgent(value string) {
	r.InstructionForDebtorAgent = (*Max140Text)(&value)
}

func (r *RequestedModification6) AddInstructionForNextAgent() *InstructionForNextAgent1 {
	newValue := new(InstructionForNextAgent1)
	r.InstructionForNextAgent = append(r.InstructionForNextAgent, newValue)
	return newValue
}

func (r *RequestedModification6) AddInstructionForCreditorAgent() *InstructionForCreditorAgent1 {
	newValue := new(InstructionForCreditorAgent1)
	r.InstructionForCreditorAgent = append(r.InstructionForCreditorAgent, newValue)
	return newValue
}

func (r *RequestedModification6) AddRemittanceInformation() *RemittanceInformation11 {
	r.RemittanceInformation = new(RemittanceInformation11)
	return r.RemittanceInformation
}
