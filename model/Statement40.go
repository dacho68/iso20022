package model

// Characteristics of the statement.
type Statement40 struct {

	// Sequential number of the report.
	ReportNumber *Number3Choice `xml:"RptNb,omitempty"`

	// Identification of the SecuritiesStatementQuery message sent to request this statement.
	QueryReference *Max35Text `xml:"QryRef,omitempty"`

	// Reference common to all pages of a statement.
	StatementIdentification *Max35Text `xml:"StmtId,omitempty"`

	// Date and time of the statement.
	StatementDateTime *DateAndDateTimeChoice `xml:"StmtDtTm"`

	// Frequency of the statement.
	Frequency *Frequency25Choice `xml:"Frqcy"`

	// Indicates whether the statement is complete or contains changes only.
	UpdateType *UpdateType15Choice `xml:"UpdTp"`

	// Type of balance on which the statement is prepared.
	StatementBasis *StatementBasis7Choice `xml:"StmtBsis"`

	// Indicates whether there is activity or information update reported in the statement.
	ActivityIndicator *YesNoIndicator `xml:"ActvtyInd"`

	// Indicates whether the statement is audited or not.
	AuditedIndicator *YesNoIndicator `xml:"AudtdInd"`

	// Indicates whether the statement reports holdings at subsafekeeping account level.
	SubAccountIndicator *YesNoIndicator `xml:"SubAcctInd"`

	// Indicates whether the statement contains tax lot details.
	TaxLotIndicator *YesNoIndicator `xml:"TaxLotInd,omitempty"`
}

func (s *Statement40) AddReportNumber() *Number3Choice {
	s.ReportNumber = new(Number3Choice)
	return s.ReportNumber
}

func (s *Statement40) SetQueryReference(value string) {
	s.QueryReference = (*Max35Text)(&value)
}

func (s *Statement40) SetStatementIdentification(value string) {
	s.StatementIdentification = (*Max35Text)(&value)
}

func (s *Statement40) AddStatementDateTime() *DateAndDateTimeChoice {
	s.StatementDateTime = new(DateAndDateTimeChoice)
	return s.StatementDateTime
}

func (s *Statement40) AddFrequency() *Frequency25Choice {
	s.Frequency = new(Frequency25Choice)
	return s.Frequency
}

func (s *Statement40) AddUpdateType() *UpdateType15Choice {
	s.UpdateType = new(UpdateType15Choice)
	return s.UpdateType
}

func (s *Statement40) AddStatementBasis() *StatementBasis7Choice {
	s.StatementBasis = new(StatementBasis7Choice)
	return s.StatementBasis
}

func (s *Statement40) SetActivityIndicator(value string) {
	s.ActivityIndicator = (*YesNoIndicator)(&value)
}

func (s *Statement40) SetAuditedIndicator(value string) {
	s.AuditedIndicator = (*YesNoIndicator)(&value)
}

func (s *Statement40) SetSubAccountIndicator(value string) {
	s.SubAccountIndicator = (*YesNoIndicator)(&value)
}

func (s *Statement40) SetTaxLotIndicator(value string) {
	s.TaxLotIndicator = (*YesNoIndicator)(&value)
}
