package seev

import (
	"encoding/xml"

	"github.com/yudaprama/iso20022/model"
)

type Document04100106 struct {
	XMLName xml.Name                                                      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.041.001.06 Document"`
	Message *CorporateActionInstructionCancellationRequestStatusAdviceV06 `xml:"CorpActnInstrCxlReqStsAdvc"`
}

func (d *Document04100106) AddMessage() *CorporateActionInstructionCancellationRequestStatusAdviceV06 {
	d.Message = new(CorporateActionInstructionCancellationRequestStatusAdviceV06)
	return d.Message
}

// Scope
// An account servicer sends the CorporateActionInstructionCancellationRequestStatusAdvice message to an account owner or its designated agent to report status of a previously received CorporateActionInstructionCancellationRequest message sent by the account owner. This will include the acknowledgement/rejection of a request to cancel an outstanding instruction.
// Usage
// The message may also be used to:
// - re-send a message previously sent (the sub-function of the message is Duplicate),
// - provide a third party with a copy of a message for information (the sub-function of the message is Copy),
// - re-send to a third party a copy of a message for information (the sub-function of the message is Copy Duplicate),
// using the relevant elements in the business application header (BAH).
type CorporateActionInstructionCancellationRequestStatusAdviceV06 struct {

	// Identification of a related instruction cancellation request document.
	InstructionCancellationRequestIdentification *model.DocumentIdentification9 `xml:"InstrCxlReqId,omitempty"`

	// Identification of other documents as well as the document number.
	OtherDocumentIdentification []*model.DocumentIdentification33 `xml:"OthrDocId,omitempty"`

	// General information about the corporate action event.
	CorporateActionGeneralInformation *model.CorporateActionGeneralInformation91 `xml:"CorpActnGnlInf"`

	// Provides information about the processing status of the instruction cancellation request.
	InstructionCancellationRequestStatus []*model.InstructionCancellationRequestStatus9Choice `xml:"InstrCxlReqSts"`

	// Information about the corporate action option.
	CorporateActionInstruction *model.CorporateActionOption116 `xml:"CorpActnInstr,omitempty"`

	// Provides additional information.
	AdditionalInformation *model.CorporateActionNarrative10 `xml:"AddtlInf,omitempty"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	SupplementaryData []*model.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (c *CorporateActionInstructionCancellationRequestStatusAdviceV06) AddInstructionCancellationRequestIdentification() *model.DocumentIdentification9 {
	c.InstructionCancellationRequestIdentification = new(model.DocumentIdentification9)
	return c.InstructionCancellationRequestIdentification
}

func (c *CorporateActionInstructionCancellationRequestStatusAdviceV06) AddOtherDocumentIdentification() *model.DocumentIdentification33 {
	newValue := new(model.DocumentIdentification33)
	c.OtherDocumentIdentification = append(c.OtherDocumentIdentification, newValue)
	return newValue
}

func (c *CorporateActionInstructionCancellationRequestStatusAdviceV06) AddCorporateActionGeneralInformation() *model.CorporateActionGeneralInformation91 {
	c.CorporateActionGeneralInformation = new(model.CorporateActionGeneralInformation91)
	return c.CorporateActionGeneralInformation
}

func (c *CorporateActionInstructionCancellationRequestStatusAdviceV06) AddInstructionCancellationRequestStatus() *model.InstructionCancellationRequestStatus9Choice {
	newValue := new(model.InstructionCancellationRequestStatus9Choice)
	c.InstructionCancellationRequestStatus = append(c.InstructionCancellationRequestStatus, newValue)
	return newValue
}

func (c *CorporateActionInstructionCancellationRequestStatusAdviceV06) AddCorporateActionInstruction() *model.CorporateActionOption116 {
	c.CorporateActionInstruction = new(model.CorporateActionOption116)
	return c.CorporateActionInstruction
}

func (c *CorporateActionInstructionCancellationRequestStatusAdviceV06) AddAdditionalInformation() *model.CorporateActionNarrative10 {
	c.AdditionalInformation = new(model.CorporateActionNarrative10)
	return c.AdditionalInformation
}

func (c *CorporateActionInstructionCancellationRequestStatusAdviceV06) AddSupplementaryData() *model.SupplementaryData1 {
	newValue := new(model.SupplementaryData1)
	c.SupplementaryData = append(c.SupplementaryData, newValue)
	return newValue
}
