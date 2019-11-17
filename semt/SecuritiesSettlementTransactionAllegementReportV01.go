package semt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01900101 struct {
	XMLName xml.Name                                            `xml:"urn:iso:std:iso:20022:tech:xsd:semt.019.001.01 Document"`
	Message *SecuritiesSettlementTransactionAllegementReportV01 `xml:"SctiesSttlmTxAllgmtRpt"`
}

func (d *Document01900101) AddMessage() *SecuritiesSettlementTransactionAllegementReportV01 {
	d.Message = new(SecuritiesSettlementTransactionAllegementReportV01)
	return d.Message
}

// Scope
// An account servicer sends a SecuritiesSettlementTransactionAllegementReport to an account owner to provide, at a specified time, the status and details of pending settlement allegements, for all or selected securities in a specified safekeeping account.
// The account servicer/owner relationship may be:
// - a central securities depository or another settlement market infrastructure acting on behalf of their participants
// - an agent (sub-custodian) acting on behalf of their global custodian customer, or
// - a custodian acting on behalf of an investment management institution or a broker/dealer.
// Usage
// The message may also be used to:
// - re-send a message previously sent (the sub-function of the message is Duplicate),
// - provide a third party with a copy of a message for information (the sub-function of the message is Copy),
// - re-send to a third party a copy of a message for information (the sub-function of the message is Copy Duplicate).
// ISO 15022 - 20022 Coexistence
// This ISO 20022 message is reversed engineered from ISO 15022. Both standards will coexist for a certain number of years. Until this coexistence period ends, the usage of certain data types is restricted to ensure interoperability between ISO 15022 and 20022 users. Compliance to these rules is mandatory in a coexistence environment.  The coexistence restrictions are described in a Textual Rule linked to the Message Items they concern. These coexistence textual rules are clearly identified as follows:  “CoexistenceXxxxRule”.
type SecuritiesSettlementTransactionAllegementReportV01 struct {

	// Information that unambiguously identifies a SecuritiesSettlementTransactionAllegementReport message as know by the account servicer.
	Identification *model.DocumentIdentification11 `xml:"Id"`

	// Page number of the message (within a statement) and continuation indicator to indicate that the statement is to continue or that the message is the last page of the statement.
	Pagination *model.Pagination `xml:"Pgntn"`

	// General information related to report.
	StatementGeneralDetails *model.Statement17 `xml:"StmtGnlDtls"`

	// Party that legally owns the account.
	AccountOwner *model.PartyIdentification13Choice `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *model.SecuritiesAccount13 `xml:"SfkpgAcct"`

	// Details of the allegement.
	AllegementDetails []*model.SecuritiesTradeDetails4 `xml:"AllgmtDtls,omitempty"`

	// Party that originated the message, if other than the sender.
	MessageOriginator *model.PartyIdentification10Choice `xml:"MsgOrgtr,omitempty"`

	// Party that is the final destination of the message, if other than the receiver.
	MessageRecipient *model.PartyIdentification10Choice `xml:"MsgRcpt,omitempty"`
}

func (s *SecuritiesSettlementTransactionAllegementReportV01) AddIdentification() *model.DocumentIdentification11 {
	s.Identification = new(model.DocumentIdentification11)
	return s.Identification
}

func (s *SecuritiesSettlementTransactionAllegementReportV01) AddPagination() *model.Pagination {
	s.Pagination = new(model.Pagination)
	return s.Pagination
}

func (s *SecuritiesSettlementTransactionAllegementReportV01) AddStatementGeneralDetails() *model.Statement17 {
	s.StatementGeneralDetails = new(model.Statement17)
	return s.StatementGeneralDetails
}

func (s *SecuritiesSettlementTransactionAllegementReportV01) AddAccountOwner() *model.PartyIdentification13Choice {
	s.AccountOwner = new(model.PartyIdentification13Choice)
	return s.AccountOwner
}

func (s *SecuritiesSettlementTransactionAllegementReportV01) AddSafekeepingAccount() *model.SecuritiesAccount13 {
	s.SafekeepingAccount = new(model.SecuritiesAccount13)
	return s.SafekeepingAccount
}

func (s *SecuritiesSettlementTransactionAllegementReportV01) AddAllegementDetails() *model.SecuritiesTradeDetails4 {
	newValue := new(model.SecuritiesTradeDetails4)
	s.AllegementDetails = append(s.AllegementDetails, newValue)
	return newValue
}

func (s *SecuritiesSettlementTransactionAllegementReportV01) AddMessageOriginator() *model.PartyIdentification10Choice {
	s.MessageOriginator = new(model.PartyIdentification10Choice)
	return s.MessageOriginator
}

func (s *SecuritiesSettlementTransactionAllegementReportV01) AddMessageRecipient() *model.PartyIdentification10Choice {
	s.MessageRecipient = new(model.PartyIdentification10Choice)
	return s.MessageRecipient
}
