package model

// Transaction information in the completion advice message.
type CardPaymentTransaction13 struct {

	// Flag indicating whether the transaction data must be captured or not in addition to the message process.
	TransactionCapture *TrueFalseIndicator `xml:"TxCaptr,omitempty"`

	// Type of transaction being undertaken for the main service.
	TransactionType *CardPaymentServiceType4Code `xml:"TxTp"`

	// Service in addition to the main service.
	AdditionalService []*CardPaymentServiceType2Code `xml:"AddtlSvc,omitempty"`

	// Additional attribute of the service type.
	ServiceAttribute *CardPaymentServiceType3Code `xml:"SvcAttr,omitempty"`

	// Category code conform to ISO 18245, related to the type of services or goods the merchant provides for the transaction.
	MerchantCategoryCode *Min3Max4Text `xml:"MrchntCtgyCd"`

	// Unique identification of the transaction assigned by the POI (Point Of Interaction).
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`

	// Identification of the original transaction.
	OriginalTransaction *CardPaymentTransaction17 `xml:"OrgnlTx,omitempty"`

	// Outcome of the transaction at the acceptor.
	TransactionSuccess *TrueFalseIndicator `xml:"TxSucss"`

	// Notify that a previous transaction has to be reversed if this original transaction has been appoved by the acquirer.
	Reversal *TrueFalseIndicator `xml:"Rvsl,omitempty"`

	// Indicate that the acceptor has forced the transaction in spite of the authorisation result (online or offline), or incident to complete the transaction.
	MerchantOverride *TrueFalseIndicator `xml:"MrchntOvrrd,omitempty"`

	// List of incidents during the transaction.
	FailureReason []*FailureReason2Code `xml:"FailrRsn,omitempty"`

	// Identification of the transaction assigned by the initiating party for the recipient party.
	InitiatorTransactionIdentification *Max35Text `xml:"InitrTxId,omitempty"`

	// Identification of the transaction assigned by the recipient party for the initiating party.
	RecipientTransactionIdentification *Max35Text `xml:"RcptTxId,omitempty"`

	// Unique identification of the reconciliation period between the acceptor and the acquirer. This identification might be linked to the identification of the settlement for further verification by the merchant.
	ReconciliationIdentification *Max35Text `xml:"RcncltnId,omitempty"`

	// Interchange information related to the card scheme.
	InterchangeData *Max35Text `xml:"IntrchngData,omitempty"`

	// Details of the transaction.
	TransactionDetails *CardPaymentTransactionDetails3 `xml:"TxDtls"`

	// Outcome of the authorisation request.
	AuthorisationResult *AuthorisationResult2 `xml:"AuthstnRslt,omitempty"`

	// Result of the verifications performed by the issuer to deliver or decline the authorisation.
	TransactionVerificationResult *TransactionVerificationResult2 `xml:"TxVrfctnRslt,omitempty"`

	// Additional information related to the transaction.
	AdditionalTransactionData []*Max70Text `xml:"AddtlTxData,omitempty"`
}

func (c *CardPaymentTransaction13) SetTransactionCapture(value string) {
	c.TransactionCapture = (*TrueFalseIndicator)(&value)
}

func (c *CardPaymentTransaction13) SetTransactionType(value string) {
	c.TransactionType = (*CardPaymentServiceType4Code)(&value)
}

func (c *CardPaymentTransaction13) AddAdditionalService(value string) {
	c.AdditionalService = append(c.AdditionalService, (*CardPaymentServiceType2Code)(&value))
}

func (c *CardPaymentTransaction13) SetServiceAttribute(value string) {
	c.ServiceAttribute = (*CardPaymentServiceType3Code)(&value)
}

func (c *CardPaymentTransaction13) SetMerchantCategoryCode(value string) {
	c.MerchantCategoryCode = (*Min3Max4Text)(&value)
}

func (c *CardPaymentTransaction13) AddTransactionIdentification() *TransactionIdentifier1 {
	c.TransactionIdentification = new(TransactionIdentifier1)
	return c.TransactionIdentification
}

func (c *CardPaymentTransaction13) AddOriginalTransaction() *CardPaymentTransaction17 {
	c.OriginalTransaction = new(CardPaymentTransaction17)
	return c.OriginalTransaction
}

func (c *CardPaymentTransaction13) SetTransactionSuccess(value string) {
	c.TransactionSuccess = (*TrueFalseIndicator)(&value)
}

func (c *CardPaymentTransaction13) SetReversal(value string) {
	c.Reversal = (*TrueFalseIndicator)(&value)
}

func (c *CardPaymentTransaction13) SetMerchantOverride(value string) {
	c.MerchantOverride = (*TrueFalseIndicator)(&value)
}

func (c *CardPaymentTransaction13) AddFailureReason(value string) {
	c.FailureReason = append(c.FailureReason, (*FailureReason2Code)(&value))
}

func (c *CardPaymentTransaction13) SetInitiatorTransactionIdentification(value string) {
	c.InitiatorTransactionIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction13) SetRecipientTransactionIdentification(value string) {
	c.RecipientTransactionIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction13) SetReconciliationIdentification(value string) {
	c.ReconciliationIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction13) SetInterchangeData(value string) {
	c.InterchangeData = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction13) AddTransactionDetails() *CardPaymentTransactionDetails3 {
	c.TransactionDetails = new(CardPaymentTransactionDetails3)
	return c.TransactionDetails
}

func (c *CardPaymentTransaction13) AddAuthorisationResult() *AuthorisationResult2 {
	c.AuthorisationResult = new(AuthorisationResult2)
	return c.AuthorisationResult
}

func (c *CardPaymentTransaction13) AddTransactionVerificationResult() *TransactionVerificationResult2 {
	c.TransactionVerificationResult = new(TransactionVerificationResult2)
	return c.TransactionVerificationResult
}

func (c *CardPaymentTransaction13) AddAdditionalTransactionData(value string) {
	c.AdditionalTransactionData = append(c.AdditionalTransactionData, (*Max70Text)(&value))
}
