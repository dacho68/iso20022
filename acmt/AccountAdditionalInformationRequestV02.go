package acmt

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document01200102 struct {
	XMLName xml.Name                                `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.012.001.02 Document"`
	Message *AccountAdditionalInformationRequestV02 `xml:"AcctAddtlInfReq"`
}

func (d *Document01200102) AddMessage() *AccountAdditionalInformationRequestV02 {
	d.Message = new(AccountAdditionalInformationRequestV02)
	return d.Message
}

// The AccountAdditionalInformationRequest message is sent from a financial institution to an organisation as part of maintenance process. This message is sent in response to a maintenance request message from the organisation, if the business content is valid, but additional information is required.
type AccountAdditionalInformationRequestV02 struct {

	// Set of elements for the identification of the message and related references.
	References *model.References3 `xml:"Refs"`

	// Identifies the business sender of the message, if it is not the account owner or account servicing financial institution.
	From *model.OrganisationIdentification8 `xml:"Fr,omitempty"`

	// Identifier for an organisation.
	OrganisationIdentification *model.OrganisationIdentification8 `xml:"OrgId"`

	// Unique and unambiguous identifier of a financial institution, as assigned under an internationally recognised or proprietary identification scheme.
	AccountServicerIdentification *model.BranchAndFinancialInstitutionIdentification5 `xml:"AcctSvcrId"`

	// Unique and unambiguous identification of the account between the account owner and the account servicer.
	AccountIdentification []*model.AccountForAction1 `xml:"AcctId"`

	// Contains the signature with its components, namely signed info, signature value, key info and the object.
	DigitalSignature []*model.PartyAndSignature2 `xml:"DgtlSgntr,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (a *AccountAdditionalInformationRequestV02) AddReferences() *model.References3 {
	a.References = new(model.References3)
	return a.References
}

func (a *AccountAdditionalInformationRequestV02) AddFrom() *model.OrganisationIdentification8 {
	a.From = new(model.OrganisationIdentification8)
	return a.From
}

func (a *AccountAdditionalInformationRequestV02) AddOrganisationIdentification() *model.OrganisationIdentification8 {
	a.OrganisationIdentification = new(model.OrganisationIdentification8)
	return a.OrganisationIdentification
}

func (a *AccountAdditionalInformationRequestV02) AddAccountServicerIdentification() *model.BranchAndFinancialInstitutionIdentification5 {
	a.AccountServicerIdentification = new(model.BranchAndFinancialInstitutionIdentification5)
	return a.AccountServicerIdentification
}

func (a *AccountAdditionalInformationRequestV02) AddAccountIdentification() *model.AccountForAction1 {
	newValue := new(model.AccountForAction1)
	a.AccountIdentification = append(a.AccountIdentification, newValue)
	return newValue
}

func (a *AccountAdditionalInformationRequestV02) AddDigitalSignature() *model.PartyAndSignature2 {
	newValue := new(model.PartyAndSignature2)
	a.DigitalSignature = append(a.DigitalSignature, newValue)
	return newValue
}

func (a *AccountAdditionalInformationRequestV02) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	a.SupplementaryData = append(a.SupplementaryData, newValue)
	return newValue
}
